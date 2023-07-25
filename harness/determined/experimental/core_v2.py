"""Singleton-style Core API. Familiar, yet revolutionary."""

import dataclasses
import logging
import uuid
from typing import Any, Dict, List, Optional, Union

import determined
from determined import core, get_cluster_info
from determined.common import yaml
from determined.core._context import Context
from determined.experimental import Determined, core_context_v2
from determined.experimental import unmanaged as unmanaged_utils

logger = logging.getLogger("determined.experimental.core_v2")


_context = None  # type: Optional[Context]
_client = None  # type: Optional[Determined]
train = None  # type: Optional[core.TrainContext]
distributed = None  # type: Optional[core.DistributedContext]
preempt = None  # type: Optional[core.PreemptContext]
checkpoint = None  # type: Optional[core.CheckpointContext]
searcher = None  # type: Optional[core.SearcherContext]
info = None  # type: Optional[determined.ClusterInfo]


@dataclasses.dataclass
class DefaultConfig:
    """
    In the future, `DefaultConfig` options will be merged with the experiment config values when
    running in the managed mode.
    """

    name: Optional[str] = None
    hparams: Optional[Dict[str, Any]] = None
    data: Optional[Dict] = None
    description: Optional[str] = None
    labels: Optional[List[str]] = None
    # Also to be added:
    # - hyperparameters: const only
    # - checkpoint_policy: for optional gc
    # For managed mode, workspace and project MUST be present in the exp conf for RBAC reasons.
    # We expose them as separate arguments on the `det.init` call.
    # - workspace
    # - project
    # Unsupported:
    # - bind_mounts
    # - data_layer
    # - debug
    # - entrypoint
    # - environment
    # - internal
    # - max_restarts
    # - optimizations
    # - profiling
    # - reproducibility
    # - security
    # - slurm
    # Searcher:
    # - searcher
    # - min_checkpoint_period
    # - min_validation_period
    # - pbs
    # - perform_initial_validation
    # - records_per_epoch
    # - scheduling_unit
    # Deprecated:
    # - data_layer
    # - pbs
    # - tensorboard_storage
    checkpoint_storage: Optional[Union[str, Dict[str, Any]]] = None
    # Searcher is currently a hack to disambiguate single trial experiments
    # and hp searches in WebUI.
    # Also searcher metric is useful for sorting in the UI and, in the future, checkpoint gc.
    searcher: Optional[Dict[str, Any]] = None
    # TODO(ilia): later to be replaced with:
    # Unmanaged mode only config options:
    # multi_trial_experiment: bool = False
    # metric: Optional[str] = None
    # smaller_is_better: bool = True  # mode?


@dataclasses.dataclass
class UnmanagedConfig:
    """
    `UnmanagedConfig` values are only used in the unmanaged mode.
    """

    # For the managed mode, workspace is critical for RBAC so it cannot be easily
    # merged and patched during the experiment runtime.
    workspace: Optional[str] = None
    project: Optional[str] = None
    # External experiment & trial ids.
    experiment_id: Optional[Union[str, int]] = None
    trial_id: Optional[Union[str, int]] = None


def _set_globals() -> None:
    global train
    global checkpoint
    global distributed
    global preempt
    global searcher
    global info

    assert _context is not None
    train = _context.train
    checkpoint = _context.checkpoint
    distributed = _context.distributed
    preempt = _context.preempt
    searcher = _context.searcher
    info = _context.info


def init(
    *,
    defaults: Optional[DefaultConfig] = None,
    unmanaged: Optional[UnmanagedConfig] = None,
    client: Optional[Determined] = None,
    # Classic core context arguments.
    distributed: Optional[core.DistributedContext] = None,
    checkpoint_storage: Optional[Union[str, Dict[str, Any]]] = None,
    preempt_mode: core.PreemptMode = core.PreemptMode.WorkersAskChief,
    tensorboard_mode: core.TensorboardMode = core.TensorboardMode.AUTO,
    # resume: bool = True  # TODO(ilia): optionally control resume behaviour.
) -> None:
    global _context
    global _client

    if _context is not None:
        _context.close()

    if client is None:
        client = Determined()
    _client = client

    info = get_cluster_info()
    if info is not None and info.task_type == "TRIAL":
        # Managed trials.
        _context = core_context_v2.init(
            distributed=distributed,
            checkpoint_storage=checkpoint_storage,
            preempt_mode=preempt_mode,
            tensorboard_mode=tensorboard_mode,
            client=client,
        )
        _context.start()
        _set_globals()
        return

    # Unmanaged trials.
    if defaults is None:
        raise NotImplementedError(
            "either specify `defaults`, or run as a managed determined experiment"
        )

    # Construct the config.
    defaulted_config = defaults or DefaultConfig()
    unmanaged_config = unmanaged or UnmanagedConfig()
    checkpoint_storage = checkpoint_storage or defaulted_config.checkpoint_storage

    config = {
        "name": defaulted_config.name or f"unmanaged-{uuid.uuid4().hex[:8]}",
        "data": defaulted_config.data,
        "description": defaulted_config.description,
        "labels": defaulted_config.labels,
        "searcher": defaulted_config.searcher
        or {
            "name": "single",
            "metric": "unmanaged",
            "max_length": 100000000,
        },
        "workspace": unmanaged_config.workspace,
        "project": unmanaged_config.project,
    }

    config_text = yaml.dump(config)
    assert config_text is not None

    unmanaged_info = unmanaged_utils.get_or_create_experiment_and_trial(
        client,
        config_text=config_text,
        experiment_id=unmanaged_config.experiment_id,
        trial_id=unmanaged_config.trial_id,
        distributed=distributed,
        hparams=defaulted_config.hparams,
    )

    _context = core_context_v2.init(
        distributed=distributed,
        checkpoint_storage=checkpoint_storage,
        preempt_mode=preempt_mode,
        tensorboard_mode=tensorboard_mode,
        unmanaged_info=unmanaged_info,
        client=client,
    )
    _context.start()
    _set_globals()


def close() -> None:
    global _context
    global train

    if _context is not None:
        _context.close()

    _context = None
    train = None


def url_reverse_webui_exp_view() -> str:
    assert info is not None
    assert info._trial_info is not None
    exp_id = info._trial_info.experiment_id

    assert _client is not None
    return unmanaged_utils.url_reverse_webui_exp_view(_client, exp_id)
