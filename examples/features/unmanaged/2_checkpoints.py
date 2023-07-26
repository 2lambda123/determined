#!/usr/bin/env python3

import random

from determined.experimental import core_v2 as det


def main():
    det.init(
        defaults=det.DefaultConfig(
            name="unmanaged-2-checkpoints",
            # We allow configuring the local checkpoint storage directory.
            # checkpoint_storage="/tmp/determined-cp",
        ),
        unmanaged=det.UnmanagedConfig(
            external_experiment_id="test-unmanaged-2-checkpoints",
            external_trial_id="test-unmanaged-2-checkpoints",
            # e.g., requeued jobs on slurm:
            # external_experiment_id=f"some-prefix-{os.environ[SLURM_JOB_ID}",
            # external_trial_id=f"some-prefix-{os.environ[SLURM_JOB_ID}",
        ),
    )

    latest_checkpoint = det.info.latest_checkpoint
    initial_i = 0
    if latest_checkpoint is not None:
        with det.checkpoint.restore_path(latest_checkpoint) as path:
            with (path / "state").open() as fin:
                ckpt = fin.read()
                print("Checkpoint contents:", ckpt)

                i_str, _ = ckpt.split(",")
                initial_i = int(i_str)

    print("determined experiment id: ", det.info._trial_info.experiment_id)
    print("initial step:", initial_i)
    for i in range(initial_i, initial_i + 100):
        det.train.report_training_metrics(steps_completed=i, metrics={"loss": random.random()})
        if (i + 1) % 10 == 0:
            loss = random.random()
            det.train.report_validation_metrics(steps_completed=i, metrics={"loss": loss})

            with det.checkpoint.store_path({"steps_completed": i}) as (path, uuid):
                with (path / "state").open("w") as fout:
                    fout.write(f"{i},{loss}")

    det.close()


if __name__ == "__main__":
    main()
