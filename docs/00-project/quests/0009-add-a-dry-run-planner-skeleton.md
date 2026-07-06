---
id: QUEST-0009
title: Add a Dry-Run Planner Skeleton
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the small dry-run planner boundary added for future action planning.
related:
  - internal/planner
  - docs/00-project/milestones.md#m-0004
  - b0a84c4 Add a dry-run planner skeleton
last_updated: 2026-07-06
---

# QUEST-0009 — Add a Dry-Run Planner Skeleton

> Pack the plan before taking the trail.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Related Commit

`b0a84c4` — Add a dry-run planner skeleton

## Outcome

- Added the `internal/planner` package.
- Added focused planner tests.
- The planner creates dry-run plans with a noop action and a supplied reason.
- The planner is not wired into the app yet.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No app wiring.
- No CLI expansion.
- No GPIO.
- No shutdown execution.
- No `rc.local` edits.
- No `SafeShutdown.py` replacement.
- No service activation.
- No resume.
- No state storage.
- No third-party dependencies.

## Milestone Note

This supports [M-0004 — Dry-Run Action Charm](../milestones.md#m-0004) but does
not create a new milestone.
