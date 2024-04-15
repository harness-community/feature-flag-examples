from logging import DEBUG
from time import sleep
from os import getenv

from featureflags.evaluations.auth_target import Target
from featureflags.config import with_events_url
from featureflags.config import with_base_url
from featureflags.client import CfClient
from featureflags.util import log

log.setLevel(DEBUG)


def main():
    if not (sdk_key := getenv("FF_SDK_KEY")):
        log.error("Set SDK key with FF_SDK_KEY")

    if relay_proxy_addr := getenv("RELAY_PROXY_ADDRESS"):
        config_addr = relay_proxy_addr
        events_addr = relay_proxy_addr
    else:
        config_addr = "https://config.ff.harness.io/api/1.0"
        events_addr = "https://events.ff.harness.io/api/1.0"

    log.info(f"connecting to ff at {config_addr}")

    client = CfClient(
        sdk_key,
        with_base_url(config_addr),
        with_events_url(events_addr),
    )

    target = Target(
        identifier="connectionCheck",
        name="Connection Check",
        attributes={"sdk": "python"},
    )

    while True:
        log.info(client.bool_variation(getenv("FF_IDENTIFIER", "test"), target, False))
        sleep(10)


if __name__ == "__main__":
    main()
