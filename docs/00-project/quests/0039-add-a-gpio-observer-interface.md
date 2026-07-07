---
id: QUEST-0039
title: Add a GPIO Observer Interface
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add the safe observer seam that lets future GPIO input become a dry-run power intent without touching real hardware yet.
related:
  - internal/input
  - internal/app
  - internal/config
  - internal/planner
  - internal/executor
  - internal/events
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0038-add-a-configurable-dry-run-power-policy.md
last_updated: 2026-07-07
---

# QUEST-0039 - Add a GPIO Observer Interface

> Place a small input lantern before the hardware gate: bright enough for tests,
> quiet enough that every real GPIO relic still sleeps.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added an `internal/input` observer interface for project-level input events.
- Added a fake observer that can emit a `PowerButtonPressed`-style event.
- Added app processing for input events that maps the fake power-button event
  into the existing `PowerButtonPressed` intent.
- Reused the existing config policy, planner, executor, and event breadcrumb
  flow.
- Tests prove the fake observer event becomes the existing power intent, the
  configured noop policy is honored, the deterministic dry-run plan and
  execution path still works, and event breadcrumbs are recorded.
- Architecture docs now describe the observer as the safe seam before real
  hardware input.

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

## Milestone Note

The daemon now carries a tiny observer charm in its field kit. It can hear a
fake power-button event during tests, translate that event into the existing
dry-run power intent, and leave the same clear breadcrumb ledger behind. The
real hardware map remains folded for a future quest.

