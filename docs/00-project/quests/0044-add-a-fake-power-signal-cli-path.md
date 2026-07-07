---
id: QUEST-0044
title: Add a Fake Power Signal CLI Path
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Make the configured latching power switch interpreter visible from the daemon command line with fake raw signal input.
related:
  - cmd/retroflag-powerd
  - internal/config
  - internal/input
  - internal/app
  - README.md
  - docs/03-operations/gpio-read-only-plan.md
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0043-add-a-latching-power-switch-interpreter.md
last_updated: 2026-07-07
---

# QUEST-0044 - Add a Fake Power Signal CLI Path

> Place the raw-signal lantern on the command line, then let the configured
> switch charm decide whether the trail reaches the noop power ledger.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added `--fake-power-signal <low|high|unverified>` to `retroflag-powerd`.
- Added default configured power input and latching power switch interpreter
  options to app config.
- Added dev CLI options for the configured latching switch map:
  `--power-switch-active-signal <low|high>` and
  `--power-switch-active-state <off|on>`.
- The fake signal path starts the app lifecycle, creates a raw signal event for
  the configured power input, interprets it through the configured latching
  power switch map, prints a compact deterministic result, and includes event
  breadcrumbs.
- Interpreted `SwitchOff` reaches the existing dry-run/noop power intent path.
- Interpreted `SwitchOn` and `SwitchUnknown` report deterministically and exit
  cleanly without requesting shutdown behavior.
- Invalid fake signal input fails clearly before the app starts.
- Kept the existing `--dry-run-power-button` and
  `--fake-power-button-observer` paths working.
- Updated the README, architecture map, and GPIO read-only plan with the new
  fake raw signal command.

## Validation

- [x] `go test ./...` passed during implementation.
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

## Safe Command

```sh
go run ./cmd/retroflag-powerd --fake-power-signal low
```

## Milestone Note

The daemon now carries a raw-signal charm in the workshop satchel. Developers
can feed low, high, or unverified into the configured latching switch map, see
the interpretation plainly, and verify that only `SwitchOff` walks into the
dry-run/noop power ledger.
