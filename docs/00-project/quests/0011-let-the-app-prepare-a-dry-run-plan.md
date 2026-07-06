---
id: QUEST-0011
title: Let the App Prepare a Dry-Run Plan
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record startup preparation of the app-owned dry-run plan.
related:
  - internal/app
  - internal/planner
  - docs/00-project/quests/0010-thread-the-planner-through-the-app-lifecycle.md
last_updated: 2026-07-06
---

# QUEST-0011 — Let the App Prepare a Dry-Run Plan

> Check the compass and tuck the noop plan into the satchel before raising the
> ready lantern.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- App startup now asks the app-owned planner to prepare and store a plan before
  the daemon reaches ready.
- The prepared plan remains dry-run only and carries a noop action.
- Focused tests prove plan preparation and the ready, stopping, and stopped
  lifecycle badges.
- Existing lifecycle log output remains unchanged.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No real shutdown action execution.
- No GPIO or hardware behavior.
- No CLI flags or command output changes.
- No normal startup log changes.
- No `rc.local` or `SafeShutdown.py` edits.
- No systemd, packaging, service activation, resume, or state-storage changes.
- No third-party dependencies.

## Milestone Note

This supports [M-0004 — Dry-Run Action Charm](../milestones.md#m-0004) but does
not create a new milestone.
