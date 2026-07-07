---
id: QUEST-0027
title: Add an App Startup Diagnostic Snapshot
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Store a tiny internal startup diagnostic snapshot for future log and CLI lanterns.
related:
  - internal/app
  - docs/00-project/quests/0023-add-an-app-runtime-summary-accessor.md
  - docs/00-project/quests/0024-add-a-runtime-diagnostic-value.md
  - docs/00-project/quests/0025-cover-runtime-diagnostic-across-the-app-lifecycle.md
  - docs/00-project/quests/0026-make-runtime-diagnostic-shutdown-coverage-explicit.md
last_updated: 2026-07-06
---

# QUEST-0027 - Add an App Startup Diagnostic Snapshot

> Tuck the ready-time diagnostic charm into the app satchel so future lanterns
> can read the startup badge without waking any user-facing output yet.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `App` now preserves the `RuntimeDiagnostic` captured when startup reaches
  the ready state.
- `App.StartupDiagnostic()` returns the stored diagnostic with an availability
  badge, making the pre-startup absence explicit.
- Startup capture follows the existing `RuntimeDiagnostic` and
  `RuntimeSummary` path.
- Focused app tests prove the startup diagnostic is unavailable before startup,
  matches the runtime diagnostic at startup, and remains the startup snapshot
  after shutdown.
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

The app now carries one startup diagnostic relic in its field kit, ready for the
next map step that wants to show or format startup diagnostics deliberately.
