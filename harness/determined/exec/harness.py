import argparse
import contextlib
import faulthandler
import logging
import sys
from typing import Iterator, cast

import determined as det
from determined import core, horovod, load
from determined.common.api import analytics, certs
from determined.profiler import ProfilerAgent

try:
    from determined import pytorch
except ImportError:
    pytorch = None
    pass


@contextlib.contextmanager
def maybe_periodic_stacktraces(debug_enabled: bool) -> Iterator[None]:
    if debug_enabled:
        faulthandler.dump_traceback_later(30, repeat=True)
    try:
        yield
    finally:
        if debug_enabled:
            faulthandler.cancel_dump_traceback_later()


def main(train_entrypoint: str) -> int:
    info = det.get_cluster_info()
    assert info is not None, "must be run on-cluster"
    assert info.task_type == "TRIAL", f'must be run with task_type="TRIAL", not "{info.task_type}"'

    # TODO: refactor data_layer, and profiling to to not use the cli_cert.
    certs.cli_cert = certs.default_load(info.master_url)

    trial_class = load.trial_class_from_entrypoint(train_entrypoint)

    # TODO: Don't include EnvContext object in the future high-level APIs for PyTorch or Keras.
    # It was natural to create this big-blob-of-config object, but it was a mistake to pass it into
    # the lowest layers of the harness code; it's too large of an object to be easily mockable,
    # which is part of why building local training mode has always been a challenge.
    #
    # A better pattern is to pass in exactly the information that is necessary at each layer.  We
    # will use that pattern for the future high-level APIs, but it's not worth refactoring e.g. the
    # TFKerasTrialController or EstimatorTrialController to add that functionality, so for now we
    # continue with the legacy strategy.

    if isinstance(trial_class, pytorch.PyTorchTrial):
        return _run_pytorch_trial(trial_class, info)

    env = det.EnvContext(
        master_url=info.master_url,
        master_cert_file=info.master_cert_file,
        master_cert_name=info.master_cert_name,
        experiment_config=info.trial._config,
        hparams=info.trial.hparams,
        latest_checkpoint=info.latest_checkpoint,
        steps_completed=info.trial._steps_completed,
        use_gpu=bool(info.gpu_uuids),
        container_gpus=info.gpu_uuids,
        slot_ids=info.slot_ids,
        debug=info.trial._debug,
        det_trial_unique_port_offset=info.trial._unique_port_offset,
        det_trial_id=str(info.trial.trial_id),
        det_experiment_id=str(info.trial.experiment_id),
        det_agent_id=info.agent_id,
        det_cluster_id=info.cluster_id,
        trial_seed=info.trial.trial_seed,
        trial_run_id=info.trial._trial_run_id,
        allocation_id=info.allocation_id,
        managed_training=True,
        test_mode=False,
        on_cluster=True,
    )

    det.common.set_logger(env.debug)
    logging.debug("Starting harness.")

    with maybe_periodic_stacktraces(env.debug):
        # Step 1: Load user code.
        # We can't build a core.Context without rank information, and we can't gather rank
        # information until the distributed backend is initialized, and we can't initialize the
        # correct distributed backend until we know which Trial class the user implemented.
        controller_class = load.get_trial_controller_class(trial_class)
        if info.container_rank == 0:
            try:
                analytics.send_analytics("trial_loaded", analytics.get_trial_analytics(trial_class))
            except Exception as e:
                logging.debug(f"Cannot send analytics: {e}")

        # Step 2: Initialize framework-specific details (dtrain framework, random seeds, etc).
        distributed_backend = det._DistributedBackend()
        controller_class.pre_execute_hook(env, distributed_backend)

        # Step 3: Now that the dtrain framework is initialized, build the DistributedContext object.
        # For harness.py, we only support a fixed set of Determined-provided launch layers, since
        # the TrialControllers only support a fixed set of launch layers.
        distributed = None
        if distributed_backend.use_horovod():
            distributed = core.DistributedContext.from_horovod(horovod.hvd)
        elif distributed_backend.use_deepspeed():
            distributed = core.DistributedContext.from_deepspeed()
        elif distributed_backend.use_torch():
            distributed = core.DistributedContext.from_torch_distributed()
        elif len(info.container_addrs) > 1 or len(info.slot_ids) > 1:
            raise ValueError(
                "In multi-slot tasks, the determined.exec.harness module must not be invoked "
                "directly.  Instead, it must be wrapped in one of the following launch layers: "
                "determined.launch.horovod, determined.launch.deepspeed"
            )

        # Step 4: Let core.init() create the core.Context.
        with core.init(
            distributed=distributed,
            preempt_mode=core.PreemptMode.ChiefOnly,
            tensorboard_mode=core.TensorboardMode.MANUAL,
        ) as core_context:
            trial_context = trial_class.trial_context_class(core_context, env)

            # Step 4: Instantiate the user's Trial.
            trial_inst = trial_class(trial_context)

            # Step 5: Create a TrialController and execute training
            logging.info(f"Creating {controller_class.__name__} with {trial_class.__name__}.")

            controller = controller_class.from_trial(
                trial_inst=trial_inst,
                context=trial_context,
                env=env,
            )

            controller.run()

    return 0


def _run_pytorch_trial(
    trial_class: pytorch.PyTorchTrial,
    info: det.ClusterInfo,
):
    det.common.set_logger(info.trial._debug)
    logging.debug("Starting harness.")

    with maybe_periodic_stacktraces(info.trial._debug):
        # Step 1: Load user code.
        # We can't build a core.Context without rank information, and we can't gather rank
        # information until the distributed backend is initialized, and we can't initialize the
        # correct distributed backend until we know which Trial class the user implemented.
        if info.container_rank == 0:
            try:
                analytics.send_analytics("trial_loaded", analytics.get_trial_analytics(trial_class))
            except Exception as e:
                logging.debug(f"Cannot send analytics: {e}")

        # Step 2: Initialize framework-specific details (dtrain framework, random seeds, etc).
        distributed_backend = det._DistributedBackend()
        pytorch.PyTorchTrialController.pre_execute_hook(info.trial.trial_seed, distributed_backend)

        # Step 3: Now that the dtrain framework is initialized, build the DistributedContext object.
        # For harness.py, we only support a fixed set of Determined-provided launch layers, since
        # the TrialControllers only support a fixed set of launch layers.
        distributed = None
        if distributed_backend.use_horovod():
            distributed = core.DistributedContext.from_horovod(horovod.hvd)
        elif distributed_backend.use_deepspeed():
            distributed = core.DistributedContext.from_deepspeed()
        elif distributed_backend.use_torch():
            distributed = core.DistributedContext.from_torch_distributed()
        elif len(info.container_addrs) > 1 or len(info.slot_ids) > 1:
            raise ValueError(
                "In multi-slot tasks, the determined.exec.harness module must not be invoked "
                "directly.  Instead, it must be wrapped in one of the following launch layers: "
                "determined.launch.horovod, determined.launch.deepspeed"
            )

        # Step 4: Let core.init() create the core.Context.
        with core.init(
            distributed=distributed,
            preempt_mode=core.PreemptMode.ChiefOnly,
            tensorboard_mode=core.TensorboardMode.MANUAL,
        ) as core_context:
            trial_context = pytorch.PyTorchTrialContext(
                core_context=core_context,
                trial_seed=info.trial.trial_seed,
                hparams=info.trial.hparams,
                slots_per_trial=info.trial._config["resources"]["slots_per_trial"],
                num_gpus=len(info.gpu_uuids),
                exp_conf=info.trial._config,
                aggregation_frequency=info.trial._config["optimizations"]["aggregation_frequency"],
                fp16_compression=info.trial._config["optimizations"]["gradient_compression"],
                average_aggregated_gradients=cast(
                    bool, info.trial._config["optimizations"]["average_aggregated_gradients"]
                ),
                steps_completed=info.trial._steps_completed,
                managed_training=True,
            )

            # Step 4: Instantiate the user's Trial.
            trial_inst = trial_class(trial_context)

            # Step 5: Create a TrialController and execute training
            logging.info(
                f"Creating {pytorch.PyTorchTrialController.__name__} with {trial_class.__name__}."
            )

            profiling_enabled = cast(bool, info.trial._config["profiling"]["enabled"])
            det_profiler = ProfilerAgent(
                trial_id=str(info.trial.trial_id),
                agent_id=info.agent_id,
                master_url=info.master_url,
                profiling_is_enabled=profiling_enabled,
                global_rank=core_context.distributed.rank,
                local_rank=core_context.distributed.local_rank,
                begin_on_batch=profiling_enabled
                and cast(int, info.trial._config["profiling"]["begin_on_batch"])
                or 0,
                end_after_batch=profiling_enabled
                and cast(int, info.trial._config["profiling"]["end_after_batch"])
                or 0,
                sync_timings=cast(bool, info.trial._config["profiling"]["sync_timings"]),
            )

            controller = pytorch.PyTorchTrialController(
                trial_inst=trial_inst,
                context=trial_context,
                min_checkpoint_period=pytorch.TrainUnit._from_values(
                    **info.trial._config["min_checkpoint_period"]
                ),
                min_validation_period=pytorch.TrainUnit._from_values(
                    **info.trial._config["min_validation_period"]
                ),
                average_training_metrics=cast(
                    bool, info.trial._config["optimizations"]["average_training_metrics"]
                ),
                checkpoint_policy=info.trial._config["checkpoint_policy"],
                smaller_is_better=cast(bool, info.trial._config["searcher"]["smaller_is_better"]),
                searcher_metric_name=info.trial._config["searcher"]["metric"],
                local_training=False,
                det_profiler=det_profiler,
                steps_completed=info.trial._steps_completed,
                latest_checkpoint=info.latest_checkpoint,
                debug=info.trial._debug,
                step_zero_validation=info.trial._config["perform_initial_validation"]
            )
            controller.run()
    return 0


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("train_entrypoint")
    args = parser.parse_args()
    sys.exit(main(args.train_entrypoint))
