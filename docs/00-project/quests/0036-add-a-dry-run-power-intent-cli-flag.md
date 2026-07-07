---
id: QUEST-0036
title: Add a Dry-Run Power Intent CLI Flag
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Make the dry-run power intent path visible from the daemon command line without GPIO or real shutdown behavior.
related:
  - cmd/retroflag-powerd
  - README.md
  - docs/00-project/quests/0035-add-a-dry-run-power-intent-path.md
  - docs/04-architecture/system-overview.md
last_updated: 2026-07-07
---

# QUEST-0036 - Add a Dry-Run Power Intent CLI Flag

> Put the power-button charm on the command-line map, bright enough for
> workshop testing and gentle enough that no real relic stirs.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added `--dry-run-power-button` to `retroflag-powerd`.
- The flag starts the app lifecycle, waits for the daemon to reach ready, and
  processes the existing `power_button_pressed` dry-run intent.
- The intent still travels through the app, planner, and executor path.
- The command prints a compact deterministic stdout line for tests and scripts.
- Lifecycle logs remain on stderr.
- The result states that the run was dry-run/noop-only and that no real
  shutdown or hardware action occurred.
- Added focused command tests for the version flag and dry-run power-button
  flag.
- Added the safe command to the README field kit.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No GPIO reads.
- No shutdown command execution.
- No systemd service activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume behavior.
- No persistent state.
- No packaging changes.
- No third-party dependencies.

## Safe Command

```sh
go run ./cmd/retroflag-powerd --dry-run-power-button
```

## Milestone Note

The daemon now has a small command-line lantern for the first power-shaped
behavior. Developers can press the dry-run power-button charm from the terminal,
watch the planner and executor path light up, and keep every real hardware and
shutdown spell safely asleep.
