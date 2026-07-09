---
id: QUEST-0088
title: Add the True GPi Case 2 Boot Trace Lantern Script Skeleton
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add the first read-only True GPi Case 2 Boot Trace Lantern script skeleton based on the QUEST-0087 design.
related:
  - ../../../scripts/gpi-case2-true-boot-trace-lantern.sh
  - ../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-session-watch-field-run-procedure.md
  - ../../03-operations/human-facing-field-lantern-script-ux-standard.md
last_updated: 2026-07-09
---

# QUEST-0088 - Add the True GPi Case 2 Boot Trace Lantern Script Skeleton

> Give startup its first pocket Lantern: copied by hand, run from
> `/home/retropi/`, bright while it works, and honest about every missing sign.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Diagnostics

## Intent

Implement the first read-only True Boot Trace Lantern script skeleton for GPi
Case 2 based on the QUEST-0087 design. The script supports the current
scp-first field practice:

- Copy the script to `retropi@gpi:/home/retropi/`.
- Run it from `/home/retropi/`.
- Retrieve the final Boot Trace Ledger with `scp`.
- Do not require the repository to be checked out on the handheld Relic.
- Treat SSH as optional support, not the primary handheld UX.
- Keep the side switch as the normal stock shutdown path while responsive.
- Keep the top sleep/resume button suspect unless a procedure explicitly says
  otherwise.

The skeleton must stay foreground, bounded, human-facing, and read-only. It
must not add behavior-changing runtime code.

## Outcome

- Added
  [`scripts/gpi-case2-true-boot-trace-lantern.sh`](../../../scripts/gpi-case2-true-boot-trace-lantern.sh).
- Supported `--help`, `--plain`, `--duration`, `--interval`, and `--output`.
- Added `NO_COLOR` handling and plain ASCII output for copyable field runs.
- Printed startup, stage, progress, status, timing, warning count, missing
  evidence count, and final artifact path while running.
- Wrote one final `.txt` Boot Trace Ledger artifact with status, read-only
  marker, artifact path, start/end timestamps, requested and observed
  duration, host metadata, uptime/proc uptime, systemd timing where available,
  selected journal and dmesg boot hints where available, frontend detection,
  framebuffer/display hints, vcgencmd temperature/voltage/throttled values
  where available, memory, disk, load, warnings, missing evidence, and a final
  `Artifact Summary`.
- Kept missing optional tools as Ledger evidence instead of fatal errors.
- Added syntax, help, `--plain`, and `NO_COLOR` smoke checks to
  `make check-scripts`.
- Updated the True Boot Trace design, Boot Power Trace Lantern map, README,
  and Field Lantern capture procedure to link and describe the current
  skeleton.

## Boundary

- No boot config changes.
- No service, systemd, or automatic boot activation.
- No GPIO reads.
- No GPIO writes.
- No power, display, shutdown, sleep, resume, RetroPie, or EmulationStation
  behavior changes.
- No shutdown, reboot, halt, suspend, sleep, or resume command.
- No installer, firmware, RetroFlag script, package, or repair execution.
- No Lantern Dispatch, telemetry, network submission, automatic upload, or
  automatic fix.
- No assumption that the repository exists on the GPi Case 2.

## Acceptance Checks

- [x] A True Boot Trace Lantern script exists under `scripts/`.
- [x] The script is read-only and writes only the requested final artifact
  file.
- [x] `--help` documents purpose, safety boundaries, default output, and
  scp-first examples.
- [x] `--plain` disables color, glyphs, and live terminal behavior.
- [x] `NO_COLOR` disables color.
- [x] `--duration` provides a bounded capture length.
- [x] `--output` provides an explicit final artifact path, with a documented
  `$HOME` timestamp default.
- [x] The script prints visible progress/status while running.
- [x] The final Ledger includes status, read-only marker, artifact path,
  started/ended timestamps, requested and observed duration, host metadata,
  uptime/proc uptime, systemd timing, journal/dmesg hints, frontend detection,
  framebuffer/display hints, vcgencmd values, memory, disk, load, warnings,
  missing evidence, and final `Artifact Summary`.
- [x] Missing optional tools are recorded, not treated as fatal unless the
  script cannot write the artifact.
- [x] `make check-scripts` includes syntax, help, `--plain`, and `NO_COLOR`
  smoke checks for the new script.
- [x] Relevant design, map, README, and field-run docs link the script.

## Validation

- [x] `sh scripts/gpi-case2-true-boot-trace-lantern.sh --help` passed.
- [x] `sh scripts/gpi-case2-true-boot-trace-lantern.sh --plain --duration 1 --interval 1 --output /tmp/gpi-case2-true-boot-trace-lantern-plain-smoke.txt` passed.
- [x] `NO_COLOR=1 sh scripts/gpi-case2-true-boot-trace-lantern.sh --duration 1 --interval 1 --output /tmp/gpi-case2-true-boot-trace-lantern-nocolor-smoke.txt` passed.
- [x] `make check-scripts` passed.
- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- This quest is completed by the commit that adds the True Boot Trace Lantern
  skeleton and threads it into the project maps. The final pushed commit hash
  is reported in the quest handoff because a Git commit cannot contain its own
  final object hash.

## Final Notes

The True Boot Trace Lantern now has its first handheld-safe Relic. It is still
not a boot service and not a repair spell. It gives startup a bounded,
scp-first Ledger with warm progress, missing-evidence honesty, and a final map
for later SignalMage or Caster work.
