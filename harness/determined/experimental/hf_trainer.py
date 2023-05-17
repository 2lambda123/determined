# type: ignore

from typing import Any, Optional

from transformers import TrainerCallback
from transformers.trainer_callback import TrainerControl, TrainerState
from transformers.training_args import TrainingArguments


# TODO(ilia): reuse `determined/transformers/_hf_callback.py` instead.
class DetCallback(TrainerCallback):
    def __init__(self):
        try:
            from determined.experimental import core_v2 as det

            self._det = det
        except ImportError:
            raise RuntimeError(
                "DetCallback requires determined to be installed. Run `pip install determined`."
            )
        self._initialized = False

    def setup(self, args: TrainingArguments, state: TrainerState, model: Optional[Any]) -> None:
        if self._initialized or not state.is_world_process_zero:
            return

        self._initialized = True

        det = self._det

        det.init(
            defaults=det.DetConfig(
                name="hf-trainer",
            )
        )

    def on_train_begin(
        self, args: TrainingArguments, state: TrainerState, control: TrainerControl, **kwargs
    ):
        self.setup(args, state, None)
        self._training = True

    def on_log(
        self,
        args: TrainingArguments,
        state: TrainerState,
        control: TrainerControl,
        model=None,
        logs=None,
        **kwargs
    ):
        self.setup(args, state, None)

        if not state.is_world_process_zero:
            return

        det = self._det

        if self._training:
            det.train.report_training_metrics(steps_completed=state.global_step, metrics=logs)
        else:
            det.train.report_validation_metrics(steps_completed=state.global_step, metrics=logs)

    def on_train_end(
        self, args: TrainingArguments, state: TrainerState, control: TrainerControl, **kwargs
    ):
        self.setup(args, state, None)
        self._training = False

        if not state.is_world_process_zero:
            return
