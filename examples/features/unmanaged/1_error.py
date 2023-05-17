#!/usr/bin/env python3
#
# Demonstrate an experiment which errors out.

import random

from determined.experimental import core_v2 as det


def main():
    det.init(
        defaults=det.DetConfig(
            name="unmanaged-1-error",
        ),
    )

    for i in range(100):
        det.train.report_training_metrics(steps_completed=i, metrics={"loss": random.random()})
        if i == 15:
            raise ValueError("oops")
        if (i + 1) % 10 == 0:
            det.train.report_validation_metrics(
                steps_completed=i, metrics={"loss": random.random()}
            )

    det.close()


if __name__ == "__main__":
    main()
