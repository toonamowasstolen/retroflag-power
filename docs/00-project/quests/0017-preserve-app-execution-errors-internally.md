---
id: QUEST-0017
title: Preserve App Execution Errors Internally
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the app startup execution error badge kept for internal inspection.
related:
  - internal/app
  - internal/executor
  - docs/00-project/quests/0016-execute-the-prepared-dry-run-plan-inside-the-app.md
last_updated: 2026-07-06
---

# QUEST-0017 — Preserve App Execution Errors Internally

> Keep the startup execution ledger honest: when the app lifts the dry-run
> lantern, it now keeps the executor error badge too, even when that badge is
> empty.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Startup execution now stores the executor error internally instead of
  discarding it.
- `ExecutionStatus()` exposes a read-only badge showing whether execution
  completed and whether an execution error was captured.
- Successful noop dry-run startup still has the same `ExecutionSummary()`.
- Focused app tests prove successful startup completes without an execution
  error and that an internally stored unsupported execution error is visible as
  captured.
- Daemon logs, CLI behavior, lifecycle statuses, packaging, hardware behavior,
  and state storage remain unchanged.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No real shutdown action execution.
- No GPIO or hardware behavior.
- No CLI flags, command output, or startup log changes.
- No lifecycle status changes.
- No `rc.local` or `SafeShutdown.py` edits.
- No systemd, packaging, service activation, resume, or state-storage changes.
- No third-party dependencies.

## Milestone Note

This supports [M-0004 — Dry-Run Action Charm](../milestones.md#m-0004) but does
not create a new milestone.
