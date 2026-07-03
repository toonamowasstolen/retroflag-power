#!/bin/sh
set -eu

BINARY_SOURCE=./retroflag-powerd
SERVICE_SOURCE=./packaging/retroflag-power.service
BINARY_DESTINATION=/usr/local/bin/retroflag-powerd
SERVICE_DESTINATION=/etc/systemd/system/retroflag-power.service

if [ "$(id -u)" -ne 0 ]; then
	echo "error: run this installer as root" >&2
	exit 1
fi

if [ ! -x "$BINARY_SOURCE" ]; then
	echo "error: $BINARY_SOURCE is missing or not executable; run 'make build' first" >&2
	exit 1
fi

if [ ! -f "$SERVICE_SOURCE" ]; then
	echo "error: $SERVICE_SOURCE is missing; run this installer from the repository root" >&2
	exit 1
fi

install -Dm755 "$BINARY_SOURCE" "$BINARY_DESTINATION"
install -Dm644 "$SERVICE_SOURCE" "$SERVICE_DESTINATION"

echo "Installed retroflag-powerd and the preview systemd unit."
echo "The service was not reloaded, enabled, started, or restarted."
