---
id: QUEST-0012
title: Add a Planner Snapshot to the App
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the read-only app snapshot for its prepared dry-run plan.
related:
  - internal/app
  - internal/planner
  - docs/00-project/quests/0011-let-the-app-prepare-a-dry-run-plan.md
last_updated: 2026-07-06
---

# QUEST-0012 — Add a Planner Snapshot to the App

> Open the satchel just far enough to inspect the noop plan without disturbing
> the trail map inside.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `Plan()` reports whether the app has prepared a plan and returns its current
  value as a read-only snapshot.
- A newly constructed app reports that no plan is available.
- Startup makes the dry-run plan available with its noop action and reason.
- Changing a returned snapshot cannot alter the app-owned plan.
- Daemon lifecycle behavior and log output remain unchanged.

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
