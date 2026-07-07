---
id: QUEST-0042
title: Separate Raw Signals from Interpreted Inputs
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Keep raw input observations distinct from interpreted switch and button meaning before future GPIO work begins.
related:
  - internal/input
  - docs/03-operations/gpio-read-only-plan.md
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0041-plan-the-gpio-read-only-path.md
last_updated: 2026-07-07
---

# QUEST-0042 - Separate Raw Signals from Interpreted Inputs

> Give the input lantern two lenses: one for the wire's raw glimmer, and one
> for the meaning the configured map will name later.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added `input.SignalState` with `SignalLow`, `SignalHigh`, and
  `SignalUnverified`.
- Added `input.SignalEvent(name, state)` for deterministic raw signal
  observations.
- Added small validation on `SignalState` so unknown raw signal states are clear
  to tests and future diagnostics.
- Reserved interpreted latching switch vocabulary with `SwitchOn`, `SwitchOff`,
  and `SwitchUnknown`.
- Reserved interpreted momentary button vocabulary with `ButtonPressed`,
  `ButtonReleased`, and `ButtonUnknown`.
- Kept the existing fake power-button observer and dry-run power intent CLI
  behavior unchanged.
- Updated the GPIO read-only plan and architecture map to show the layered route:
  raw signal, configured interpretation, latching switch or momentary button
  event, power intent, policy, plan, noop execution, and breadcrumbs.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No real GPIO reads.
- No GPIO writes.
- No Raspberry Pi GPIO dependency.
- No shutdown command execution.
- No systemd service activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume behavior.
- No persistent state.
- No packaging changes.
- No debouncing implementation.
- No press-and-hold, double-tap, or triple-tap gesture implementation.
- No full latching-switch transition state machine.

## Milestone Note

The input field kit now has a clean first badge for raw observations. Future
GPIO work can record whether a line appears low, high, or unverified without
pretending it already knows whether the player toggled a latching switch or
pressed a momentary button. The power path remains dry-run, noop-only, and
quietly deterministic.
