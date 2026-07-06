---
id: QUEST-0023
title: Add an App Runtime Summary Accessor
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a small app-level runtime summary accessor for internal diagnostics.
related:
  - internal/app
  - docs/00-project/quests/0019-add-an-app-runtime-snapshot.md
  - docs/00-project/quests/0020-add-a-runtime-snapshot-summary.md
  - docs/00-project/quests/0021-add-a-runtime-summary-string.md
last_updated: 2026-07-06
---

# QUEST-0023 - Add an App Runtime Summary Accessor

> Add a small lantern hook to the app runtime map so callers can read the
> current summary charm without unpacking the whole snapshot satchel first.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `App.RuntimeSummary()` now returns the current `RuntimeSnapshotSummary`.
- The accessor reuses the existing `RuntimeSnapshot()` and
  `RuntimeSnapshot.Summary()` behavior.
- Focused app tests prove the accessor matches `RuntimeSnapshot().Summary()`
  before startup, after startup, and after shutdown.
- Daemon logs, CLI output, lifecycle behavior, planner behavior, executor
  behavior, summary string formatting, packaging, hardware behavior, and state
  storage remain unchanged.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No daemon log changes.
- No CLI output changes.
- No lifecycle behavior changes.
- No planner behavior changes.
- No executor behavior changes.
- No `RuntimeSnapshotSummary.String()` changes.
- No friendly formatter.
- No GPIO.
- No shutdown execution.
- No `rc.local` edits.
- No `SafeShutdown.py` edits.
- No systemd or service activation.
- No resume.
- No state storage.
- No packaging changes.
- No third-party dependencies.

## Milestone Note

This gives future internal diagnostic lanterns a shorter route to the runtime
summary. It does not add a new formatter or print the summary anywhere yet.
