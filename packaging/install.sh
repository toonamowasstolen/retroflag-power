#!/bin/sh
set -e
install -Dm755 retroflag-powerd /usr/local/bin/retroflag-powerd
install -Dm644 packaging/retroflag-power.service /etc/systemd/system/retroflag-power.service
systemctl daemon-reload
systemctl enable retroflag-power
systemctl restart retroflag-power
