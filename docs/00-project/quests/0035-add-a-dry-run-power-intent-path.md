---
id: QUEST-0035
title: Add a Dry-Run Power Intent Path
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Give the daemon its first visible power-control-shaped behavior while keeping every real hardware and shutdown boundary safe.
related:
  - internal/power
  - internal/planner
  - internal/executor
  - internal/app
  - docs/04-architecture/system-overview.md
last_updated: 2026-07-07
---

# QUEST-0035 - Add a Dry-Run Power Intent Path

> Light the first power lantern, but keep it on the workbench: bright enough to
> show the route, gentle enough that no hardware relic moves yet.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added a small internal `power.Intent` concept.
- Added `power.IntentPowerButtonPressed` as the first supported intent.
- The planner can turn the power intent into a deterministic dry-run plan.
- The planned action remains `noop`.
- The app can process the power intent through the existing planner and
  executor path.
- The executor continues to accept only dry-run/noop plans.
- Focused tests prove the app path, deterministic plan fields, dry-run/noop
  execution, and existing startup behavior.
- The architecture map now names this as the first safe power-intent path, not
  real hardware control.

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

## Milestone Note

The daemon now carries a tiny power-intent charm in its satchel. Pressing the
internal dry-run power button produces a predictable noop plan and noop result,
giving future hardware quests a clear map marker without waking the real power
relics.
