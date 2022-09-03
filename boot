#!/usr/bin/env bash

set -euo pipefail

function log() {
    echo "[$(env TZ=UTC date +%Y-%m-%dT%H:%M:%S%z)] $*"
}

firecracker --api-sock /tmp/firecracker.socket &
firecracker_pid=$!

curl --unix-socket /tmp/firecracker.socket -i \
             -X PUT 'http://localhost/boot-source'   \
             -H 'Accept: application/json'           \
             -H 'Content-Type: application/json'     \
             -d "{
               \"kernel_image_path\": \"./kernel.bin\",
               \"boot_args\": \"console=ttyS0 reboot=k panic=1 pci=off\"
          }"
curl --unix-socket /tmp/firecracker.socket -i \
         -X PUT 'http://localhost/drives/rootfs' \
         -H 'Accept: application/json'           \
         -H 'Content-Type: application/json'     \
         -d "{
           \"drive_id\": \"rootfs\",
           \"path_on_host\": \"./rootfs.ext4\",
           \"is_root_device\": true,
           \"is_read_only\": false
      }"
 curl --unix-socket /tmp/firecracker.socket -i \
         -X PUT 'http://localhost/actions'       \
         -H  'Accept: application/json'          \
         -H  'Content-Type: application/json'    \
         -d '{
         "action_type": "InstanceStart"
      }'

echo "kill the vm over at pid $firecracker_pid"
