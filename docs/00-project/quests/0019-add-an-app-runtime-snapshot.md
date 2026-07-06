---
id: QUEST-0019
title: Add an App Runtime Snapshot
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the read-only runtime snapshot for internal app diagnostics.
related:
  - internal/app
  - internal/planner
  - internal/executor
  - docs/00-project/quests/0008-thread-the-status-badge-through-the-app-lifecycle.md
  - docs/00-project/quests/0012-add-a-planner-snapshot-to-the-app.md
  - docs/00-project/quests/0016-execute-the-prepared-dry-run-plan-inside-the-app.md
last_updated: 2026-07-06
---

# QUEST-0019 — Add an App Runtime Snapshot

> Gather the badge, plan charm, and execution ledger onto one small app map for
> internal tests and future diagnostic lantern work.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `RuntimeSnapshot()` returns one read-only view of the current app status
  badge, prepared plan summary availability, prepared plan summary, execution
  status, execution summary availability, and execution summary.
- `PlanSummary()` exposes the app-owned plan summary without requiring callers
  to inspect the full prepared plan.
- Newly constructed apps report no plan or execution summaries.
- Startup reports the ready badge, dry-run noop plan summary, successful
  execution status, and successful execution summary.
- Shutdown reports the stopped badge while preserving the dry-run plan and
  execution summaries.
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
