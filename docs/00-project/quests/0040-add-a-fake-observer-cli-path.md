---
id: QUEST-0040
title: Add a Fake Observer CLI Path
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Make the fake power-button observer path runnable from the daemon command line without touching real GPIO.
related:
  - cmd/retroflag-powerd
  - internal/input
  - internal/app
  - internal/events
  - README.md
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0039-add-a-gpio-observer-interface.md
last_updated: 2026-07-07
---

# QUEST-0040 - Add a Fake Observer CLI Path

> Lift the observer lantern onto the command line, so the workshop can press a
> fake power-button charm and read the same dry-run ledger without waking any
> hardware relic.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added `--fake-power-button-observer` to `retroflag-powerd`.
- The flag starts the app lifecycle, waits for the daemon to reach ready, emits
  one fake `power_button_pressed` observer event, and routes it through the app
  input observer path.
- The event follows the existing power intent, config policy, planner,
  executor, and event breadcrumb flow.
- The command prints a compact deterministic noop result plus the event
  breadcrumb ledger.
- The existing `--dry-run-power-button` command still works.
- Added focused command tests for the fake observer success path and unsupported
  policy path.
- Added the safe fake observer command to the README and architecture map.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No real GPIO reads.
- No Raspberry Pi GPIO dependency.
- No shutdown command execution.
- No systemd service activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume behavior.
- No persistent state.
- No packaging changes.
- No debouncing behavior.
- No latching-switch semantics.

## Safe Command

```sh
go run ./cmd/retroflag-powerd --fake-power-button-observer
```

## Milestone Note

The daemon now has a command-line observer charm in its satchel. Developers can
emit one fake power-button event, watch it travel through the same input path as
future GPIO work, and keep the result dry-run, noop-only, and easy to verify.
