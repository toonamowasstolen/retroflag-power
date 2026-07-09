---
id: QUEST-0078
title: Retrofit the GPi Case 2 Manual Lantern Script to the Human-Facing UX Standard
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record the completed retrofit that made the GPi Case 2 manual Bundle Collector Lantern script human-facing, timed, and easier to carry in the field.
related:
  - ../../03-operations/human-facing-field-lantern-script-ux-standard.md
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
  - ../../../scripts/gpi-case2-bundle-collector-field-lantern.sh
last_updated: 2026-07-09
---

# QUEST-0078 - Retrofit the GPi Case 2 Manual Lantern Script to the Human-Facing UX Standard

> Polish the hand-carried Lantern Relic so the operator sees the satchel being
> packed, sealed, timed, and ready for the Ledger.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Diagnostics

## Outcome

- Retrofitted the GPi Case 2 Bundle Collector Field Lantern script with the
  human-facing UX pattern: banner, step lines, plain mode, `--duration`,
  `--output-dir`, progress, help text, final artifact path, and duration
  reporting.
- Added timing artifacts such as `timing.txt` and `timing.tsv` so unusually
  slow capture stages can become evidence instead of folklore.
- Improved frontend detection from a brittle EmulationStation-only clue into
  `detected`, `not_detected`, or `uncertain`, with process snapshots for manual
  review.
- Updated the GPi Case 2 Bundle Collector procedure, Lantern map, and Field
  Lantern capture procedure so the docs match the friendlier terminal trail.
- Preserved the Lantern's local, read-only, human-carried shape.

## Boundary

- No GPIO writes.
- No GPIO reads.
- No shutdown or reboot execution.
- No systemd activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume implementation.
- No firmware flashing.
- No installer execution.
- No automatic fixes.
- No telemetry or network contact.
- No hardware modification instructions.

## Validation

- [x] Script smoke test passed for the human-facing output path.
- [x] Duration timing alignment was validated after the retrofit.
- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commits

- `3df81c0480f1a24ad9efb2a5156d45f290b9849e` - retrofitted the GPi Case 2
  Field Lantern UX across the script and operations docs.
- `3a5a83ff2b91fc5995ce9b6b22e3977617cebc90` - aligned Lantern duration
  timing after the retrofit, completing the field-facing timing behavior.

## Milestone Note

The GPi Case 2 Bundle Collector Lantern became a better field companion: it
shows its work, names its satchel, reports its timing, and still keeps every
safety boundary bright on the trail.
