---
id: QUEST-0091
title: Record the First True GPi Case 2 Boot Trace Field Run
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the first real True GPi Case 2 Boot Trace Lantern field result from the real artifact and operator handheld observations.
related:
  - ../../03-operations/gpi-case-2-true-boot-trace-evidence-ledger.md
  - ../../03-operations/gpi-case-2-true-boot-trace-field-run-procedure.md
  - ../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../../scripts/gpi-case2-true-boot-trace-lantern.sh
  - 0087-design-true-gpi-case-2-boot-trace-lantern.md
  - 0088-add-true-gpi-case-2-boot-trace-lantern-script-skeleton.md
  - 0089-design-gpi-case-2-first-spark-and-boot-veil-ux.md
  - 0090-add-true-boot-trace-field-run-procedure.md
last_updated: 2026-07-09
---

# QUEST-0091 - Record the First True GPi Case 2 Boot Trace Field Run

> The first True Boot Trace Ledger came home from the handheld Relic: one
> honest startup artifact, one First Spark note, and the unknown runes left
> visible for the next pass.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Evidence Intake

## Intent

Record the first real GPi Case 2 True Boot Trace Lantern field result using
the real artifact and the operator's handheld observations. This quest is
evidence intake only: no behavior-changing runtime code, no GPIO work, no
power, display, shutdown, sleep, resume, service, installer, or Lantern
Dispatch change.

The current field practice is scp-first: copy
[`scripts/gpi-case2-true-boot-trace-lantern.sh`](../../../scripts/gpi-case2-true-boot-trace-lantern.sh)
to `retropi@gpi:/home/retropi/`, run it from `/home/retropi/`, then `scp`
the final Boot Trace Ledger artifact back. The GPi Case 2 does not currently
need the full repository checked out on the handheld.

## Outcome

- Added the
  [GPi Case 2 True Boot Trace Evidence Ledger](../../03-operations/gpi-case-2-true-boot-trace-evidence-ledger.md).
- Ran the True Boot Trace Lantern from `/home/retropi/` over optional SSH
  support and retrieved the final artifact:
  `/home/retropi/gpi-case2-true-boot-trace-lantern-20260710-081529.txt`.
- Copied the artifact locally under
  `docs/03-operations/artifacts/true-boot-trace/`.
- Recorded the first observed entry with artifact facts, human observations,
  unknown or not-tested fields, interpretation buckets, candidate
  improvements, and next rune.
- Preserved the operator's roughly 15-second blank/silent First Spark window
  observation without claiming KMS, framebuffer, firmware, userspace, panel,
  case-board, or power cause.
- Linked the new evidence Ledger from the True Boot Trace design, True Boot
  Trace field-run procedure, First Spark / Boot Veil / Relic Welcome Scroll
  design, Boot Power Trace Lantern map, and README.

## Boundary

- No script behavior change.
- No GPi Case 2 runtime behavior change.
- No invented field evidence.
- No GPIO reads or writes.
- No shutdown, reboot, halt, poweroff, suspend, sleep, resume, display,
  config, service, installer, firmware, SafeShutdown, or RetroPie mutation.
- No systemd activation.
- No Lantern Dispatch, telemetry, network submission, automatic upload, or
  automatic diagnostics bundle generation.

## Acceptance Checks

- [x] A True Boot Trace evidence Ledger exists under `docs/03-operations/`.
- [x] The first real field-run entry records the remote artifact path.
- [x] The first real field-run entry records the local copied artifact path.
- [x] Run date/time, requested duration, observed duration, sample count, SSH
  status, frontend status, display hints, warning count, missing-evidence
  count, final outcome, human notes, evidence status, and next rune are
  recorded where available.
- [x] Unknown or not-tested fields remain explicit for power source,
  handheld/docked state, exact first-visible timing, first SSH timing, first
  EmulationStation timing, LED state, side-switch behavior, and top-button
  behavior.
- [x] Artifact facts, human observations, unknown or not-tested fields,
  interpretation, and candidate improvements are separate.
- [x] Interpretation buckets include first-spark delay observed,
  display blank before first visible output, SSH timing unknown, and frontend
  unknown.
- [x] The entry does not overclaim clean boot trace, frontend reached,
  frontend failed, no hard-freeze observed, or hard-freeze-like outcome
  observed without complete handheld evidence.
- [x] The entry includes a do-not-overclaim note for first-visible delay and
  KMS/framebuffer interpretation.
- [x] The True Boot Trace design, True Boot Trace field-run procedure, First
  Spark / Boot Veil / Relic Welcome Scroll design, Boot Power Trace Lantern
  map, and README link to the Ledger.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds the True Boot Trace
  evidence Ledger, records the first returned artifact, links the maps, and
  adds this quest file. The final pushed commit hash is reported in the quest
  handoff because a Git commit cannot contain its own final object hash.

## Final Notes

The first True Boot Trace Field Lantern result is now in the Ledger: the Relic
accepted the scp-first script, completed a 120-second requested trace in 127
seconds, wrote 25 samples, preserved `throttled=0x50000` as a cautious clue,
and left the next rune clear: repeat with exact handheld timestamps, LED
state, visible frontend state, and side-switch notes.
