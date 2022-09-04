#!/usr/bin/env bash

set -euo pipefail

function log() {
    echo "[$(env TZ=UTC date +%Y-%m-%dT%H:%M:%S%z)] $*"
}

socket="/tmp/firecracker.socket"
rm -fv $socket

log "booting firecracker..."
firecracker --no-api --config-file ./vm_config.json & # TODO: Should we be providing a socket here anyway?
firecracker_pid=$!

log "kill the vm over at pid $firecracker_pid"
