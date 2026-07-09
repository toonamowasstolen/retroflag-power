---
id: QUEST-0077
title: Add a Human-Facing Field Lantern Script UX Standard
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a reusable terminal UX standard for human-facing manual Field Lantern scripts.
related:
  - ../../03-operations/human-facing-field-lantern-script-ux-standard.md
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
last_updated: 2026-07-08
---

# QUEST-0077 - Add a Human-Facing Field Lantern Script UX Standard

> Give manual Lantern Relics a visible heartbeat: warm startup, clear steps,
> honest timing, and exact artifacts without touching the hardware path.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the reusable
  [Human-Facing Field Lantern Script UX Standard](../../03-operations/human-facing-field-lantern-script-ux-standard.md).
- Defined banner, step line, glyph, plain fallback, progress bar, timing,
  color, tone, final artifact, retrieval, and `--help` expectations.
- Linked the standard from the GPi Case 2 Bundle Collector Lantern procedure,
  Boot Power Trace Lantern map, and Field Lantern capture procedure without
  rewriting scripts.

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
- No telemetry.
- No hardware modification instructions.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- `63bfd9f3d20f9a79c4c08efced40fb498cff5163` - added the shared
  human-facing Field Lantern Script UX Standard.

## Milestone Note

Human-facing manual Lanterns now have a shared terminal UX Spellbook. Long
captures should show their work, name their satchel, report their duration,
and keep safety messages plain enough for the Ledger.
