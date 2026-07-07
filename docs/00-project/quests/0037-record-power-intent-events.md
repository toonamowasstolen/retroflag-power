---
id: QUEST-0037
title: Record Power Intent Events
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Leave deterministic internal breadcrumbs as the dry-run power intent moves through the app path.
related:
  - internal/events
  - internal/app
  - internal/planner
  - internal/executor
  - docs/04-architecture/system-overview.md
last_updated: 2026-07-07
---

# QUEST-0037 - Record Power Intent Events

> Add a tiny ledger to the power-button charm, bright enough to trace the route
> and quiet enough that no real shutdown spell wakes.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Reused the internal `events` package for power-intent breadcrumbs.
- Added stable event types for:
  - `power.intent_received`
  - `power.dry_run_plan_prepared`
  - `power.noop_execution_completed`
- The app records the dry-run power intent route as it receives the intent,
  prepares the noop plan, and completes noop execution.
- `App.Events()` exposes a read-only snapshot for tests and future diagnostics.
- Existing lifecycle logs are also retained internally through the same small
  event ledger.
- CLI dry-run power-button stdout remains compact and deterministic.
- Tests prove the expected event sequence, deterministic events, snapshot
  safety, dry-run/noop execution, and existing startup behavior.

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
- No logging-system redesign.

## Milestone Note

The first power-intent path now leaves three small lantern marks in the app's
internal ledger: intent received, dry-run plan prepared, and noop execution
completed. These breadcrumbs give future GPIO, shutdown, and service quests a
stable map without changing today's safe dry-run boundary.
