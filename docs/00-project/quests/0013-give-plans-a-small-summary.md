---
id: QUEST-0013
title: Give Plans a Small Summary
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the small planner summary added for prepared dry-run plans.
related:
  - internal/planner
  - internal/app
  - docs/00-project/quests/0012-add-a-planner-snapshot-to-the-app.md
last_updated: 2026-07-06
---

# QUEST-0013 — Give Plans a Small Summary

> Add a tiny ledger badge for the prepared plan, just enough to read the map
> without casting any shutdown spells.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `Plan.Summary()` now returns a small `PlanSummary` value for focused tests and
  future app code.
- The summary reports whether the plan is dry-run, how many actions it carries,
  and whether it is noop-only.
- The dry-run planner still prepares exactly one noop action.
- The zero-value plan has a boring empty summary.
- App lifecycle tests can inspect the prepared plan summary without changing
  daemon behavior or log output.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No real shutdown action execution.
- No GPIO or hardware behavior.
- No CLI flags, command output, or log output changes.
- No `rc.local` or `SafeShutdown.py` edits.
- No systemd, packaging, service activation, resume, or state-storage changes.
- No third-party dependencies.

## Milestone Note

This supports [M-0004 — Dry-Run Action Charm](../milestones.md#m-0004) but does
not create a new milestone.
