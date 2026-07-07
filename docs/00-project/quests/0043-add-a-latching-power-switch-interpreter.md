---
id: QUEST-0043
title: Add a Latching Power Switch Interpreter
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add the first configured interpretation layer from raw signal states to latching power switch states without reading real GPIO.
related:
  - internal/input
  - internal/app
  - docs/03-operations/gpio-read-only-plan.md
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0042-separate-raw-signals-from-interpreted-inputs.md
last_updated: 2026-07-07
---

# QUEST-0043 - Add a Latching Power Switch Interpreter

> Fit the raw-signal lantern with its first honest lens: no guessed polarity,
> no hidden hardware spell, just the map the keeper names.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added `input.LatchingPowerSwitchOptions` with explicit `ActiveSignal` and
  `ActiveSwitchState` values.
- Added configured interpretation from `SignalLow`, `SignalHigh`, and
  `SignalUnverified` into `SwitchOn`, `SwitchOff`, or `SwitchUnknown`.
- Kept `SignalUnverified` deterministic by mapping it to `SwitchUnknown`.
- Added clear errors for unsupported active signal, unsupported active switch
  state, unsupported signal state, and unsupported event type.
- Added `input.PowerSwitchEvent(state)` for interpreted latching switch events.
- Wired interpreted `PowerSwitchEvent(SwitchOff)` into the existing dry-run/noop
  power intent path.
- Kept the existing fake observer and CLI paths working.
- Updated the GPIO read-only plan and architecture map to show the current
  route: raw signal, configured latching switch interpretation, power switch
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
- No momentary button interpreter.
- No full latching-switch transition state machine.

## Milestone Note

The input field kit now carries its first configured interpretation charm.
Raw low and high observations still stay plain until a caller declares both the
active signal and the active switch state. Once declared, `SwitchOff` can walk
the dry-run/noop power path and leave breadcrumbs without touching hardware.
