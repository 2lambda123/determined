import json
import os
import subprocess

import pytest

from tests import config as conf
from tests import ray_utils

EXAMPLES_ROOT = conf.EXAMPLES_PATH / "features" / "unmanaged"


@pytest.mark.e2e_cpu
def test_unmanaged() -> None:
    master_url = conf.make_master_url()
    exp_path = EXAMPLES_ROOT / "1_singleton.py"
    env = os.environ.copy()
    env["DET_MASTER"] = master_url
    subprocess.run(["python", exp_path], env=env, check=True)


@pytest.mark.e2e_cpu
def test_unmanaged_ray_hp_search() -> None:
    master_url = conf.make_master_url()
    exp_path = EXAMPLES_ROOT / "ray"
    runtime_env = {
        "env_vars": {
            "DET_MASTER": master_url,
        }
    }

    try:
        subprocess.run(["ray", "start", "--head"], check=True)
        ray_utils.ray_job_submit(
            exp_path,
            ["python", "ray_hp_search.py"],
            submit_args=[
                "--runtime-env-json",
                json.dumps(runtime_env),
            ],
        )
    finally:
        subprocess.run(["ray", "stop"], check=True)
