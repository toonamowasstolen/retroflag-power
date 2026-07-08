---
id: QUEST-0073
title: Add the GPi Case 2 Field Lantern Capture Procedure
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Document the manual read-only GPi Case 2 Field Lantern capture bundle procedure for support evidence without long terminal pastes.
related:
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
  - ../../03-operations/common-problems-mage-map.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - ../../03-operations/gpi-case-2-recovery-first-field-procedure.md
  - ../../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
last_updated: 2026-07-08
---

# QUEST-0073 - Add the GPi Case 2 Field Lantern Capture Procedure

> Give the field team one quiet satchel for evidence: local, read-only,
> inspectable, and ready to pull from macOS without a wall of terminal paste.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the manual
  [GPi Case 2 Field Lantern Capture Procedure](../../03-operations/gpi-case-2-field-lantern-capture-procedure.md).
- Documented purpose, when to use it, when not to use it, safety boundaries,
  bundle contents, bundle non-goals, privacy and redaction notes, manual
  GPi-side capture, macOS `scp` pull examples, and optional SSH key setup.
- Preserved the Field Lantern, Common Problems Mage, Lantern Dispatch,
  `SignalLow`/`SignalHigh`/`SignalUnverified`, and
  `SwitchOn`/`SwitchOff`/`SwitchUnknown` vocabulary boundaries.
- Connected the procedure to Common Problems Mage, the Local Diagnostics
  Bundle Map, recovery-first field work, and the future migration path toward
  `retroflag-powerd diagnostics --bundle`, `retroflag-powerd troubleshoot`,
  and optional Lantern Dispatch.
- Linked the new procedure from the Common Problems Mage map, Local
  Diagnostics Bundle map, and Recovery-First Field Procedure.

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
- No resume implementation.
- No firmware flashing.
- No telemetry.
- No project-code network calls.
- No automatic fixes.
- No hardware modification instructions as approved.
- No battery lead cutting, lithium battery modification, charging-circuit
  modification, blind soldering, or shorting unknown pads.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Milestone Note

The GPi Case 2 support trail now has a first-class Field Lantern: one local
bundle to carry the evidence, and no new behavior that touches the power path.
