---
id: QUEST-0024
title: Add a Runtime Diagnostic Value
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a small runtime diagnostic value as a future internal formatting boundary.
related:
  - internal/app
  - docs/00-project/quests/0020-add-a-runtime-snapshot-summary.md
  - docs/00-project/quests/0021-add-a-runtime-summary-string.md
  - docs/00-project/quests/0022-record-runtime-summary-formatter-guidance.md
  - docs/00-project/quests/0023-add-an-app-runtime-summary-accessor.md
last_updated: 2026-07-06
---

# QUEST-0024 - Add a Runtime Diagnostic Value

> Add a small diagnostic satchel for the runtime summary charm so future
> lanterns have a clear place to decide how diagnostics should look.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- `RuntimeDiagnostic` now carries the existing `RuntimeSnapshotSummary`.
- `App.RuntimeDiagnostic()` returns a diagnostic built from
  `App.RuntimeSummary()`.
- `RuntimeDiagnostic.String()` delegates directly to
  `RuntimeSnapshotSummary.String()`.
- Focused app tests prove the diagnostic summary matches `App.RuntimeSummary()`
  and that the diagnostic string matches the summary string.
- Daemon logs, CLI output, lifecycle behavior, planner behavior, executor
  behavior, friendly formatting, packaging, hardware behavior, and state
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

This places a small diagnostic badge beside the runtime summary without
lighting it up in user-facing output yet.
