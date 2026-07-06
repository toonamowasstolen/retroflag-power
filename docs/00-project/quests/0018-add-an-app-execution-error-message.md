---
id: QUEST-0018
title: Add an App Execution Error Message
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the small execution status message exposed for internal inspection.
related:
  - internal/app
  - internal/executor
  - docs/00-project/quests/0017-preserve-app-execution-errors-internally.md
last_updated: 2026-07-06
---

# QUEST-0018 — Add an App Execution Error Message

> Add a small label to the execution badge: when the dry-run lantern finds an
> executor error, the app can now show the plain error message in its internal
> satchel.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `ExecutionStatus` now includes an `ErrorMessage` field for internal
  inspection.
- Successful startup leaves the execution error message empty.
- Captured execution errors store the error string alongside the existing
  completion and error-captured badge values.
- Focused app tests prove both the successful startup message and the existing
  unsupported execution error seam.
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
- No executor behavior changes.
- No planner behavior changes.
- No `rc.local` or `SafeShutdown.py` edits.
- No systemd, packaging, service activation, resume, or state-storage changes.
- No third-party dependencies.

## Milestone Note

This supports [M-0004 — Dry-Run Action Charm](../milestones.md#m-0004) but does
not create a new milestone.
