---
id: QUEST-0082
title: Add the GPi Case 2 Session Watch Lantern Script Skeleton
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add the first read-only GPi Case 2 Session Watch Lantern script skeleton based on the QUEST-0081 design.
related:
  - ../../../scripts/gpi-case2-session-watch-lantern.sh
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - ../../03-operations/common-problems-mage-map.md
  - ../../03-operations/human-facing-field-lantern-script-ux-standard.md
last_updated: 2026-07-09
---

# QUEST-0082 - Add the GPi Case 2 Session Watch Lantern Script Skeleton

> Light the first handheld Session Watch Lantern carefully: one foreground
> Relic, one bounded watch, one final Ledger, and no changes to the path it
> observes.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Diagnostics

## Intent

Implement the first read-only Session Watch Lantern script skeleton for GPi
Case 2 based on the QUEST-0081 design. The skeleton must support a handheld
Relic workflow where SSH to `retropi@gpi` is optional support rather than the
primary UX, the side switch remains the normal stock shutdown path while the
system is responsive, and the top sleep/resume button remains suspect unless a
future procedure explicitly says otherwise.

The skeleton should be small and honest: bounded runtime observation,
human-facing progress, and one final artifact file that marks itself as
read-only evidence.

## Outcome

- Added
  [`scripts/gpi-case2-session-watch-lantern.sh`](../../../scripts/gpi-case2-session-watch-lantern.sh).
- Supported `--help`, `--plain`, `--duration`, `--interval`, and `--output`.
- Added `NO_COLOR` handling and plain ASCII output for copyable field runs.
- Printed startup, stage, progress, status, timing, and final artifact path
  lines while running.
- Wrote one final `.txt` Ledger artifact with timestamp/session metadata,
  hostname, user, kernel, uptime, watch duration, observed checkpoints,
  warnings or missing evidence, a clear `READ-ONLY / NO CHANGES MADE` marker,
  and the final artifact path.
- Added `make check-scripts` and threaded it into `make check` for portable
  shell syntax plus a help smoke check.
- Updated the QUEST-0081 design and relevant Lantern maps/procedures to link
  the script and describe the current skeleton versus future fuller satchels.

## Boundary

- No GPIO reads.
- No GPIO writes.
- No shutdown, reboot, halt, suspend, sleep, resume, or power action.
- No display, KMS, audio, boot, emulator, service, `rc.local`, or
  SafeShutdown behavior changes.
- No systemd activation.
- No installer or firmware execution.
- No telemetry, upload, Lantern Dispatch, or automatic fix.
- No requirement for an attached keyboard on the GPi Case 2.

## Acceptance Checks

- [x] A Session Watch Lantern script exists under `scripts/`.
- [x] The script is read-only and writes only the requested final artifact
  file.
- [x] `--help` documents purpose, safety boundaries, default output, and
  examples.
- [x] `--plain` disables color, glyphs, and live terminal behavior.
- [x] `NO_COLOR` disables color.
- [x] `--duration` provides a bounded watch length.
- [x] `--output` provides an explicit final artifact path, with a documented
  `$HOME` timestamp default.
- [x] The script prints periodic progress/status while running.
- [x] The final artifact includes timestamp/session metadata, hostname, user,
  kernel, uptime where available, duration, checkpoints, read-only marker, and
  final artifact path.
- [x] The docs link the script from the QUEST-0081 design and relevant Lantern
  maps/procedures.

## Validation

- [x] `make check-scripts` passed.
- [x] One-second plain smoke run passed with a final Ledger artifact under
  `/private/tmp`.
- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- This quest is completed by the commit that adds this file. The final pushed
  commit hash is reported in the quest handoff because a Git commit cannot
  contain its own final object hash.

## Final Notes

The Session Watch Lantern now has its first pocketable Relic. It is not yet a
full diagnostics satchel, and it does not pretend to be one. It watches for a
bounded duration, shows the operator that the Lantern is alive, and seals one
Ledger file that says exactly what it saw and what it did not change.
