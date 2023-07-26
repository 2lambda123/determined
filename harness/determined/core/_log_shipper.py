import atexit
import collections
import datetime
import queue
import sys
import threading
import time
from types import TracebackType
from typing import Any, Callable, Dict, Iterator, List, Optional, TextIO, Union

from determined import core
from determined.common import api


class LogShipper:
    def __init__(
        self,
        *,
        session: api.Session,
        trial_id: int,
        task_id: str,
        distributed: Optional[core.DistributedContext] = None
    ) -> None:
        self._session = session
        self._trial_id = trial_id
        self._task_id = task_id
        self._distributed = distributed

    def start(self) -> "LogShipper":
        return self

    def close(
        self,
        exc_type: Optional[type],
        exc_val: Optional[BaseException],
        exc_tb: Optional[TracebackType],
    ) -> "LogShipper":
        return self

    def __enter__(self) -> "LogShipper":
        return self.start()

    def __exit__(
        self,
        exc_type: Optional[type],
        exc_val: Optional[BaseException],
        exc_tb: Optional[TracebackType],
    ) -> "LogShipper":
        return self.close(exc_type, exc_val, exc_tb)


class ManagedTrialLogShipper(LogShipper):
    """
    Managed trials will ship their logs normally via fluentd.
    """

    pass


class _Interceptor:
    def __init__(self, original_io: TextIO, handler: Callable[[str], None]) -> None:
        self._original_io = original_io
        self._handler = handler

    def write(self, data: str) -> int:
        self._handler(data)
        return self._original_io.write(data)

    def flush(self) -> None:
        self._original_io.flush()

    def __getattr__(self, attr: str) -> Any:
        return getattr(self._original_io, attr)


SHIPPER_FLUSH_INTERVAL = 1
SHIPPER_FAILURE_BACKOFF_SECONDS = 1
LOG_BATCH_MAX_SIZE = 1000
SHIP_QUEUE_MAX_SIZE = 3 * LOG_BATCH_MAX_SIZE


class ShutdownMessage:
    pass


QueueElement = Union[str, ShutdownMessage]


class LogSender(threading.Thread):
    def __init__(self, session: api.Session, logs_metadata: Dict) -> None:
        self._queue = queue.Queue(maxsize=SHIP_QUEUE_MAX_SIZE)  # type: queue.Queue[QueueElement]
        self._logs = collections.deque()  # type: collections.deque[str]
        self._session = session
        self._logs_metadata = logs_metadata
        self._buf = ""

        super().__init__(daemon=True)

    def write(self, data: str) -> None:
        self._queue.put(data)

    def close(self) -> None:
        self._queue.put(ShutdownMessage())

    def _pop_until_deadline(self, deadline: float) -> Iterator[QueueElement]:
        while True:
            timeout = deadline - time.time()
            if timeout <= 0:
                break

            try:
                yield self._queue.get(timeout=timeout)
            except queue.Empty:
                break

    def run(self) -> None:
        while True:
            deadline = time.time() + SHIPPER_FLUSH_INTERVAL
            for m in self._pop_until_deadline(deadline):
                if isinstance(m, ShutdownMessage):
                    self.ship()
                    return

                self._logs.append(m)
                if len(self._logs) >= LOG_BATCH_MAX_SIZE:
                    self.ship()

            self.ship()

    def ship(self) -> None:
        if len(self._logs) == 0:
            return

        msgs = []

        while len(self._logs):
            data = self._logs.popleft()
            self._buf += data
            while "\n" in self._buf:
                idx = self._buf.index("\n") + 1
                line = self._buf[:idx]
                self._buf = self._buf[idx:]

                msg = dict(self._logs_metadata)
                msg["log"] = line
                msgs.append(msg)

            if len(msgs) > LOG_BATCH_MAX_SIZE:
                self._ship(msgs)
                msgs = []

        if len(msgs) > 0:
            self._ship(msgs)

    def _ship(self, msgs: List[Dict]) -> None:
        self._session.post("task-logs", json=msgs)


class UnmanagedTrialLogShipper(LogShipper):
    def start(self) -> "LogShipper":
        self._original_stdout, self._original_stderr = sys.stdout, sys.stderr

        logs_metadata = {
            "task_id": self._task_id,
            "timestamp": datetime.datetime.now(datetime.timezone.utc).isoformat(),
        }

        if self._distributed:
            logs_metadata["rank"] = str(self._distributed.rank)

        self._log_sender = LogSender(session=self._session, logs_metadata=logs_metadata)

        sys.stdout = _Interceptor(sys.stdout, self._log_sender.write)  # type: ignore
        sys.stderr = _Interceptor(sys.stderr, self._log_sender.write)  # type: ignore

        self._log_sender.start()

        atexit.register(self._exit_handler)

        return self

    def _exit_handler(self) -> None:
        self.close()

    def close(
        self,
        exc_type: Optional[type] = None,
        exc_val: Optional[BaseException] = None,
        exc_tb: Optional[TracebackType] = None,
    ) -> "LogShipper":
        atexit.unregister(self._exit_handler)

        sys.stdout, sys.stderr = self._original_stdout, self._original_stderr
        self._log_sender.close()
        self._log_sender.join()

        return self
