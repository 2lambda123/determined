# type: ignore
from typing import Optional, Union

from lightning.pytorch.loggers.logger import Logger, rank_zero_experiment
from lightning.pytorch.utilities import rank_zero_only

from determined.experimental import core_v2 as det


# TODO(ilia): Expand the integration.
class DetLogger(Logger):
    def __init__(
        self,
        *,
        defaults: Optional[det.DefaultConfig] = None,
        client: Optional[det.Determined] = None,
        experiment_id: Optional[Union[str, int]] = None,
        trial_id: Optional[Union[str, int]] = None,
    ) -> None:
        self._kwargs = {
            "defaults": defaults,
            "client": client,
            "experiment_id": experiment_id,
            "trial_id": trial_id,
        }
        self._initialized = False

    @property
    @rank_zero_experiment
    def experiment(self) -> None:
        if not self._initialized:
            det.init(**self._kwargs)
            self._initialized = True

    @property
    def name(self):
        return "DetLogger"

    @property
    def version(self):
        # Return the experiment version, int or str.
        return "0.1"

    @rank_zero_only
    def log_hyperparams(self, params):
        # params is an argparse.Namespace
        # your code to record hyperparameters goes here
        pass

    @rank_zero_only
    def log_metrics(self, metrics, step):
        # metrics is a dictionary of metric names and values
        # your code to record metrics goes here
        det.train.report_training_metrics(step, metrics)

    @rank_zero_only
    def save(self):
        # Optional. Any code necessary to save logger data goes here
        pass

    @rank_zero_only
    def finalize(self, status):
        # Optional. Any code that needs to be run after training
        # finishes goes here
        det.close()
