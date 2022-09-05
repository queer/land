#!/usr/bin/env bash

function log() {
    echo "[$(env TZ=UTC date +%Y-%m-%dT%H:%M:%S%z)] $*"
}

function require() {
  if [ -z "${1}" ]; then
    if [ -z "$2" ]; then
      log "\$$1 is not set but must be set."
    else
      log "\$$2 is not set but must be set."
    fi
    exit 1
  fi
}
