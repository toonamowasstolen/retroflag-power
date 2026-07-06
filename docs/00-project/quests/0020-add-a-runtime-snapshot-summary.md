---
id: QUEST-0020
title: Add a Runtime Snapshot Summary
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the tiny runtime snapshot summary for internal app diagnostics.
related:
  - internal/app
  - internal/planner
  - internal/executor
  - docs/00-project/quests/0008-thread-the-status-badge-through-the-app-lifecycle.md
  - docs/00-project/quests/0016-execute-the-prepared-dry-run-plan-inside-the-app.md
  - docs/00-project/quests/0019-add-an-app-runtime-snapshot.md
last_updated: 2026-07-06
---

# QUEST-0020 — Add a Runtime Snapshot Summary

> Add a small compass note to the runtime snapshot map so future diagnostic
> lanterns can read the key facts without unpacking every field in the satchel.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `RuntimeSnapshot.Summary()` now returns a small `RuntimeSnapshotSummary`
  value for internal app diagnostics.
- The summary exposes the current lifecycle state, plan presence, execution
  completion, execution success, captured execution error presence, and the
  dry-run noop-only plan/execution pairing.
- The helper reuses the existing `PlanSummary`, `ExecutionStatus`, and
  `ExecutionSummary` values already carried by the runtime snapshot.
- Focused app tests prove the summary before startup, after startup, and after
  shutdown.
- Daemon logs, CLI behavior, lifecycle statuses, planner behavior, executor
  behavior, packaging, hardware behavior, and state storage remain unchanged.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No daemon log changes.
- No CLI output changes.
- No lifecycle status changes.
- No executor behavior changes.
- No planner behavior changes.
- No real shutdown action execution.
- No GPIO or hardware behavior.
- No `rc.local` or `SafeShutdown.py` edits.
- No systemd, packaging, service activation, resume, or state-storage changes.
- No third-party dependencies.

## Milestone Note

This supports the dry-run planning and execution path plus the Status Badge
work. It does not create a new milestone.
