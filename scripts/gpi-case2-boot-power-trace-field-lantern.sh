#!/bin/sh
set -eu

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
exec sh "${SCRIPT_DIR}/gpi-case2-bundle-collector-field-lantern.sh" "$@"
