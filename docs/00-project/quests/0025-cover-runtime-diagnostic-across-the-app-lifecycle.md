---
id: QUEST-0025
title: Cover Runtime Diagnostic Across the App Lifecycle
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Prove the runtime diagnostic stays aligned with the app runtime summary across lifecycle checkpoints.
related:
  - internal/app
  - docs/00-project/quests/0020-add-a-runtime-snapshot-summary.md
  - docs/00-project/quests/0023-add-an-app-runtime-summary-accessor.md
  - docs/00-project/quests/0024-add-a-runtime-diagnostic-value.md
last_updated: 2026-07-06
---

# QUEST-0025 - Cover Runtime Diagnostic Across the App Lifecycle

> Keep the runtime diagnostic badge beside the app summary lantern through the
> startup and shutdown map points.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Test Coverage

## Outcome

- Added focused startup coverage proving `RuntimeDiagnostic().Summary` matches
  `App.RuntimeSummary()`.
- Startup coverage also proves `RuntimeDiagnostic().String()` matches
  `App.RuntimeSummary().String()`.
- Existing shutdown coverage proves the same diagnostic alignment after the app
  reaches the stopped state.
- Daemon logs, CLI output, lifecycle behavior, planner behavior, executor
  behavior, summary string formatting, packaging, hardware behavior, and state
  storage remain unchanged.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No production Go code changes.
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

The runtime diagnostic now has lifecycle coverage at startup and shutdown, so
future diagnostic spellbook work can rely on the same summary charm without
guessing where it drifted.
