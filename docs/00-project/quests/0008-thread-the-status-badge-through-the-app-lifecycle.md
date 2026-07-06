---
id: QUEST-0008
title: Thread the Status Badge Through the App Lifecycle
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
epoch: Awakening
quest_type: Implementation
related:
  - 5af99d2 Thread status through the app lifecycle
last_updated: 2026-07-06
---

# QUEST-0008 — Thread the Status Badge Through the App Lifecycle

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Related Commit

`5af99d2` — Thread status through the app lifecycle

## Outcome

- `App` now owns an internal `status.Status`.
- New apps begin with `StateStarting`.
- `Run` transitions starting → ready → stopping → stopped.
- `Status()` exposes the current internal badge.
- Existing lifecycle log output stayed unchanged.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `--version` still outputs exactly: `retroflag-powerd 0.1.0-dev`

## Boundary

- No CLI expansion.
- No GPIO.
- No shutdown execution.
- No `rc.local` edits.
- No `SafeShutdown.py` replacement.
- No service activation.
- No resume.
- No state storage.

## Milestone Note

This supports M-0006 — Status Badge but does not create a new milestone.
