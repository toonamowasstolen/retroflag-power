---
id: QUEST-0003
title: Field-Test the Heartbeat Toolkit
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Validators
purpose: Define a future reference-hardware validation quest for manually testing the prepared systemd service toolkit on the GPi Case 2 without replacing the existing shutdown path.
related:
  - docs/00-project/quests/0001-craft-the-heartbeat-toolkit.md
  - docs/00-project/requirements.md
  - docs/02-hardware/gpi-case-2.md
  - docs/04-architecture/system-overview.md
  - docs/adr/0001-use-systemd.md
  - docs/adr/0002-use-small-context-driven-daemon-lifecycle.md
last_updated: 2026-07-03
---

# QUEST-0003 — Field-Test the Heartbeat Toolkit

> The toolkit is packed. Its first field test should be deliberate, observable,
> and easy to undo.

## Quest Status

Draft

This is a future validation quest. It must remain Draft and inactive until the
project is ready to test deliberately on the GPi Case 2 reference platform.

Operating rule:

> Prepare and observe only. Do not replace the current shutdown path.

## Milestone

Bridges Milestone 1 — Awakening toward Milestone 2 — Heartbeat.

## Quest Type

Reference-hardware validation

---

# 1. Quest Summary

The Heartbeat Toolkit prepared:

- `packaging/retroflag-power.service`
- `packaging/install.sh`

Local checks prove that the daemon builds and that the packaging is passive.
They do not prove that the service behaves correctly under systemd on the GPi
Case 2 reference platform.

This future quest closes that evidence gap without replacing or disabling the
existing `SafeShutdown.py` path.

---

# 2. Objective

Safely field-test the preview service on the GPi Case 2 by:

- copying the built `retroflag-powerd` daemon and preview service to the
  reference platform
- inspecting the exact files before installation
- running `systemd-analyze verify` when available
- starting and stopping the service manually only after the safety gate passes
- confirming startup and shutdown logs through journald
- confirming the daemon does not execute shutdown behavior
- preserving a simple rollback path

The service must remain disabled at boot.

---

# 3. Safety Gate

Do not begin manual service testing unless:

- the operator has console or SSH access that will survive stopping the preview
  service
- the device is in a safe state for a controlled daemon lifecycle test
- the built daemon and unit match the reviewed repository versions
- `SafeShutdown.py` and its current startup path will remain unchanged
- no GPIO or shutdown behavior exists in the preview daemon
- rollback commands have been reviewed before the service is started

If any condition is uncertain, stop the quest and gather more evidence.

---

# 4. Scope

## In Scope

- Build or obtain a Linux ARM64 `retroflag-powerd` binary.
- Copy the binary and preview unit to a staging directory on the GPi Case 2.
- Inspect the binary path, unit name, service description, and service behavior.
- Run `systemd-analyze verify` if the command is available.
- Install the reviewed preview files manually for this test.
- Run `systemctl daemon-reload` only when ready for the controlled field test.
- Confirm the service is not enabled.
- Start `retroflag-power.service` manually.
- Inspect `systemctl status` and journald logs.
- Stop the service manually.
- Confirm the daemon logs its complete lifecycle and exits cleanly.
- Confirm the GPi Case 2 remains running and no shutdown command is executed.
- Roll back the preview files after validation.

## Out of Scope

- disabling, replacing, editing, or moving `SafeShutdown.py`
- editing `/etc/rc.local`
- enabling the service at boot
- starting the service automatically
- GPIO handling
- Power Switch or Reset Button handling
- shutdown or reboot execution
- resume behavior
- state storage
- RetroArch or EmulationStation integration
- production installation or release packaging

---

# 5. Validation Procedure

The commands below are a future operator checklist, not authorization to run
them during Milestone 1.

## 5.1 Local preflight

```sh
make check
sh -n packaging/install.sh
git diff --check
```

Build for the reference platform from a suitable Go development host:

```sh
GOOS=linux GOARCH=arm64 go build -o retroflag-powerd ./cmd/retroflag-powerd
```

Copy the binary and unit to a staging directory on the GPi Case 2 using the
operator's preferred secure transfer method.

## 5.2 Reference-platform inspection

From the staged repository or transfer directory:

```sh
sed -n '1,220p' packaging/retroflag-power.service
file ./retroflag-powerd
```

If available:

```sh
systemd-analyze verify packaging/retroflag-power.service
```

Record when `systemd-analyze` is unavailable; do not install new tooling merely
to satisfy this optional check.

## 5.3 Manual service test

Only after the safety gate passes:

```sh
sudo install -Dm755 ./retroflag-powerd /usr/local/bin/retroflag-powerd
sudo install -Dm644 packaging/retroflag-power.service /etc/systemd/system/retroflag-power.service
sudo systemctl daemon-reload
systemctl is-enabled retroflag-power.service
sudo systemctl start retroflag-power.service
systemctl status retroflag-power.service
journalctl -u retroflag-power.service --no-pager
sudo systemctl stop retroflag-power.service
journalctl -u retroflag-power.service --no-pager
```

`systemctl is-enabled` must report that the service is not enabled. Do not run
`systemctl enable`.

Expected logs include:

```text
retroflag-powerd starting
retroflag-powerd ready
shutdown signal received
retroflag-powerd stopped
```

The device must remain powered on throughout the test.

---

# 6. Acceptance Criteria

This quest can move from Draft to Verified only when:

- [ ] The tested device is the documented GPi Case 2 reference platform.
- [ ] The tested binary and service unit are traceable to a reviewed revision.
- [ ] The staged binary and service unit were inspected before installation.
- [ ] `systemd-analyze verify` passed, or its absence was recorded.
- [ ] The service was confirmed disabled before and after the test.
- [ ] The service was started manually and reached an active state.
- [ ] Journald showed the startup and ready logs.
- [ ] The service was stopped manually and exited cleanly.
- [ ] Journald showed the shutdown and stopped logs.
- [ ] No shutdown or reboot command was executed.
- [ ] The GPi Case 2 remained running throughout the test.
- [ ] `SafeShutdown.py` and `/etc/rc.local` were not changed.
- [ ] No GPIO, resume, or state behavior was introduced.
- [ ] Rollback completed successfully.
- [ ] Validation evidence and any deviations were recorded in this quest.

---

# 7. Rollback Notes

Rollback removes only the preview files installed by this field test. Before
removing anything, confirm the paths still refer to the test artifacts.

```sh
sudo systemctl stop retroflag-power.service
sudo rm -f /etc/systemd/system/retroflag-power.service
sudo rm -f /usr/local/bin/retroflag-powerd
sudo systemctl daemon-reload
sudo systemctl reset-failed retroflag-power.service
```

After rollback:

- confirm `retroflag-power.service` is not active or enabled
- confirm the preview files are absent
- confirm the original `SafeShutdown.py` path is unchanged
- confirm `/etc/rc.local` is unchanged
- retain captured logs with the validation record

Rollback must not modify or restart the original shutdown mechanism.

---

# 8. Evidence to Record

Record:

- date and operator
- hardware and OS revision
- source revision or commit
- binary architecture
- `systemd-analyze verify` result or unavailability
- service enabled state before and after testing
- `systemctl status` output
- relevant journald logs
- start and stop results
- confirmation that the device remained running
- rollback result
- unexpected behavior and follow-up work

---

# Closing

The toolkit is ready for a future field test.

It is not yet permission to replace the old path.

Test the heartbeat. Keep the safety net.
