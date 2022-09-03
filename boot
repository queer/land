#!/usr/bin/env bash

set -euo pipefail

function log() {
    echo "[$(env TZ=UTC date +%Y-%m-%dT%H:%M:%S%z)] $*"
}

socket="/tmp/firecracker.socket"
rm -fv $socket

log "booting firecracker..."
firecracker --api-sock /tmp/firecracker.socket &
firecracker_pid=$!

log "configuring kernel..."
curl --unix-socket $socket -i \
             -X PUT 'http://localhost/boot-source'   \
             -H 'Accept: application/json'           \
             -H 'Content-Type: application/json'     \
             -d "{
               \"kernel_image_path\": \"./kernel.bin\",
               \"boot_args\": \"console=ttyS0 reboot=k panic=1 pci=off\" 1>/dev/null
          }"

log "configuring rootfs..."
curl --unix-socket $socket -i \
         -X PUT 'http://localhost/drives/rootfs' \
         -H 'Accept: application/json'           \
         -H 'Content-Type: application/json'     \
         -d "{
           \"drive_id\": \"rootfs\",
           \"path_on_host\": \"./rootfs.ext4\",
           \"is_root_device\": true,
           \"is_read_only\": false
      }" 1>/dev/null

log "booting vm..."
 curl --unix-socket $socket -i \
         -X PUT 'http://localhost/actions'       \
         -H  'Accept: application/json'          \
         -H  'Content-Type: application/json'    \
         -d '{
         "action_type": "InstanceStart"
      }' 1>/dev/null

log "kill the vm over at pid $firecracker_pid"
