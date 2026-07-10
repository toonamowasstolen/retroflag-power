---
id: QUEST-0094
title: Add the Relic Welcome Scroll Preview Script
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a manually runnable, read-only Relic Welcome Scroll preview script for GPi Case 2 SSH support sessions without wiring it into login, boot, services, GPIO, display, shutdown, sleep, or resume behavior.
related:
  - ../../../scripts/gpi-case2-relic-welcome-scroll.sh
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - 0093-design-relic-welcome-scroll.md
last_updated: 2026-07-10
---

# QUEST-0094 - Add the Relic Welcome Scroll Preview Script

> Give `retropi@gpi` a friendly manual welcome scroll: a quick Field Lantern
> map for SSH support, still read-only, still scp-first, still easy to leave
> asleep until recovery-first wiring is earned.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Script / Documentation

## Intent

Add a standalone Relic Welcome Scroll preview script for GPi Case 2 SSH/system
info greeting output. The script is manually runnable only. It should help an
operator who connects over SSH see the Relic name, user, uptime, kernel, disk,
load, memory, optional Raspberry Pi temperature/throttling clues, and the
`retropi@gpi` support target without modifying the handheld.

This quest must not install the scroll into SSH login, MOTD, shell startup,
services, boot config, GPIO, display, shutdown, sleep, or resume behavior.

## Boundary

- Manual preview only.
- No edits to profile files, MOTD files, PAM, SSHD config, services, boot
  config, display config, GPIO paths, shutdown, sleep, or resume behavior.
- No attached-keyboard assumption.
- No repository checkout assumption on the GPi Case 2.
- No telemetry, upload, installer, package, repair, or Lantern Dispatch path.
- Missing tools such as `vcgencmd`, `hostname -I`, `df`, or `/proc` fields
  must show friendly fallback values instead of failing the whole scroll.
- `--plain` and `NO_COLOR` must disable color and decorative banner output.

## Outcome

- Added
  [`scripts/gpi-case2-relic-welcome-scroll.sh`](../../../scripts/gpi-case2-relic-welcome-scroll.sh).
- Supported `--help`, `--plain`, `--no-art`, and `--compact`.
- Printed a fast read-only SSH field map with ASCII art in normal mode and a
  plain fallback for `--plain` and `NO_COLOR`.
- Included Relic hostname, SSH target reminder, user, uptime, kernel summary,
  root disk free, load average, memory available, optional `vcgencmd`
  temperature, optional `vcgencmd get_throttled`, and an address hint.
- Added a clear warning that the preview is not installed into SSH login and
  should remain manually invoked until a later recovery-first wiring quest.
- Added syntax, help, normal, `--plain`, and `NO_COLOR` smoke checks to
  `make check-scripts`.
- Updated the Relic Welcome Scroll design, First Spark / Boot Veil / Welcome
  Scroll design, and Local Diagnostics Bundle Map with script links and
  scp-first preview examples.

## Manual Preview Commands

```sh
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
ssh retropi@gpi 'chmod +x /home/retropi/gpi-case2-relic-welcome-scroll.sh'
ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh'
```

Plain preview:

```sh
ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh --plain'
ssh retropi@gpi 'NO_COLOR=1 /home/retropi/gpi-case2-relic-welcome-scroll.sh'
```

## Acceptance Checks

- [x] A Relic Welcome Scroll preview script exists under `scripts/`.
- [x] The script is manually runnable and standalone.
- [x] The script supports `--help`.
- [x] The script supports `--plain`.
- [x] The script respects `NO_COLOR`.
- [x] The script supports optional `--no-art` and `--compact` flags.
- [x] Normal mode prints ASCII art / banner output.
- [x] Plain and `NO_COLOR` modes print a plain text fallback.
- [x] The scroll prints hostname / Relic name, user, uptime, kernel summary,
  disk free, load average, memory available, SSH target reminder, and a
  read-only Field Lantern reminder.
- [x] Temperature and throttled raw value are printed when `vcgencmd` is
  available, otherwise they show as `unavailable`.
- [x] Missing data is tolerated and does not fail the whole scroll.
- [x] The script does not write artifacts or mutate system state.
- [x] No shell startup, MOTD, SSHD, service, boot, GPIO, display, shutdown,
  sleep, or resume wiring was added.
- [x] `make check-scripts` includes syntax, help, normal, `--plain`, and
  `NO_COLOR` smoke checks for the new script.
- [x] Docs include scp-first preview run examples.
- [x] Docs warn that the script is not yet installed into SSH login.

## Validation

- [x] `sh -n scripts/gpi-case2-relic-welcome-scroll.sh` passed.
- [x] `sh scripts/gpi-case2-relic-welcome-scroll.sh --help` passed.
- [x] `NO_COLOR= sh scripts/gpi-case2-relic-welcome-scroll.sh` passed.
- [x] `sh scripts/gpi-case2-relic-welcome-scroll.sh --plain` passed.
- [x] `NO_COLOR=1 sh scripts/gpi-case2-relic-welcome-scroll.sh` passed.
- [x] `make check-scripts` passed.
- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- This quest is completed by the commit that adds the manual Relic Welcome
  Scroll preview script, validation checks, design updates, and this quest
  ledger. The final pushed commit hash is reported in the quest handoff because
  a Git commit cannot contain its own final object hash.

## Final Notes

The Relic now has a welcome scroll preview in the satchel, but it has not been
stitched into the doorway. Operators can copy it to `/home/retropi/`, run it by
hand, and keep SSH login boring until a later recovery-first wiring quest.
