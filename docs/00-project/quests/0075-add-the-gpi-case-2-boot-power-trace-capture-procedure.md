---
id: QUEST-0075
title: Add the GPi Case 2 Post-Boot Power Evidence Capture Procedure
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Document the manual read-only GPi Case 2 post-boot power evidence capture procedure that later needed clearer Bundle Collector naming.
related:
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
  - ../../03-operations/common-problems-mage-map.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - ../../02-hardware/gpi-case-2-power-integrity-investigation-notes.md
  - ../../02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
last_updated: 2026-07-08
---

# QUEST-0075 - Add the GPi Case 2 Post-Boot Power Evidence Capture Procedure

> Gather remembered boot clues into a small scroll after the handheld is
> responsive, without touching the power path.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the manual
  [GPi Case 2 Bundle Collector Lantern Capture Procedure](../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md).
- Documented purpose, when to run, safety notes, captured files, non-goals, a
  copy-paste shell script sample, `scp` pull examples, Field Lantern inclusion,
  and interpretation notes.
- Preserved the Field Lantern, Common Problems Mage, and Lantern Dispatch
  boundaries. Later quest work reserves Boot Power Trace Lantern for a true
  boot-time recorder.
- Linked the procedure from the Boot Power Trace Lantern map, Field Lantern
  capture procedure, Common Problems Mage map, Local Diagnostics Bundle map,
  Power Integrity notes, and KMS Power notes.

## Boundary

- Documentation only.
- No Go code changes.
- No executable project tooling.
- No GPIO writes.
- No GPIO reads.
- No shutdown or reboot execution.
- No systemd activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No firmware flashing.
- No installer execution.
- No automatic fixes.
- No telemetry or automatic upload.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Milestone Note

The GPi Case 2 power trail gained a copy-paste post-boot evidence procedure:
local, timestamped, inspectable, and short enough to stop before idle
power-save risk. Later terminology maps this as a Bundle Collector Lantern,
not a true Boot Power Trace Lantern.
