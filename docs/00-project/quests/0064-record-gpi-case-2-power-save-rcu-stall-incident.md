---
id: QUEST-0064
title: Record GPi Case 2 Power-Save RCU Stall Incident
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record a GPi Case 2 field incident where power-save or resume behavior led to repeated Linux RCU stall messages and loss of normal shutdown recovery.
related:
  - ../../02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../../03-operations/gpi-case-2-acceptance-checklist.md
  - ../../03-operations/safeshutdown-replacement-boundary-map.md
  - 0057-add-gpi-case-2-field-test-checklist-entries.md
  - 0048-map-the-safeshutdown-replacement-boundaries.md
last_updated: 2026-07-08
---

# QUEST-0064 - Record GPi Case 2 Power-Save RCU Stall Incident

> A darker field sign is now on the map: power-save and resume can strand the
> GPi Case 2 beyond ordinary software recovery.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Recorded a high-priority unresolved GPi Case 2 field incident in the hardware
  notes: repeated Linux `rcu: INFO: rcu_preempt detected stalls on CPUs/tasks`
  messages after resume or power-save behavior, network loss, side switch
  shutdown failure, top-button visible-state toggling, and physical CM4
  cartridge/card removal as the only observed stop.
- Preserved that `SafeShutdown.py` was believed to be enabled, so the incident
  must not be dismissed as only a disabled stock script.
- Recorded that the GPi Case 2 appears to auto-enter display/audio power-save
  after roughly 15-20 minutes of no input.
- Marked power-save, resume, and auto power-save behavior unsafe/unverified in
  the acceptance checklist until field-tested.
- Added investigation items for reversible emergency reset or safe power-cut
  paths, board labels, schematics, teardown photos, test pads, CM4
  `RUN`/`GLOBAL_EN`/reset/power-enable paths, and regulator enable lines.
- Added explicit lithium and charging-circuit caution before any hardware
  modification idea is considered.

## Boundary

- Documentation only.
- No Go code changes.
- No GPIO writes.
- No GPIO reads.
- No shutdown execution.
- No systemd activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume implementation.
- No persistent state.
- No telemetry.
- No network calls.
- No installer or packaging changes.
- No hardware modification instructions beyond documenting investigation
  needs and cautions.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Milestone Note

The GPi Case 2 power-save trail now has a clear red flag: recovery must be
proved before any replacement runtime or field workflow depends on software
shutdown remaining available.
