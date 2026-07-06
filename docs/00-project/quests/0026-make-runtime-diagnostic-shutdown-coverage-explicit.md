---
id: QUEST-0026
title: Make Runtime Diagnostic Shutdown Coverage Explicit
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Make the shutdown boundary for the runtime diagnostic explicit instead of relying on nearby summary lanterns.
related:
  - internal/app
  - docs/00-project/quests/0023-add-an-app-runtime-summary-accessor.md
  - docs/00-project/quests/0024-add-a-runtime-diagnostic-value.md
  - docs/00-project/quests/0025-cover-runtime-diagnostic-across-the-app-lifecycle.md
last_updated: 2026-07-06
---

# QUEST-0026 - Make Runtime Diagnostic Shutdown Coverage Explicit

> Pin the stopped badge to the diagnostic charm so future shutdown map work can
> read the boundary without guessing from a neighboring lantern.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Test Coverage

## Outcome

- Added explicit shutdown-focused coverage proving
  `App.RuntimeDiagnostic().Summary` still matches `App.RuntimeSummary()`.
- The same test proves `App.RuntimeDiagnostic().String()` still matches
  `App.RuntimeSummary().String()`.
- The shutdown diagnostic test now asserts the summary state is
  `status.StateStopped`, making the lifecycle badge visible at the diagnostic
  boundary.
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

Runtime diagnostic shutdown coverage now carries its own stopped-state badge in
the field kit, keeping the summary and diagnostic lanterns aligned for the next
diagnostic spellbook step.
