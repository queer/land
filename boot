#!/usr/bin/env bash

set -euo pipefail

source ./lib.sh

log "booting firecracker..."
firecracker --no-api --config-file ./vm_config.json # TODO: Should we be providing a socket here anyway?
