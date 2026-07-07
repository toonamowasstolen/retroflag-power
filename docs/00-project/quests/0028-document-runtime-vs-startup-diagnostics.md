---
id: QUEST-0028
title: Document Runtime vs Startup Diagnostics
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Clarify the difference between the current runtime diagnostic and the captured startup diagnostic.
related:
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0022-record-runtime-summary-formatter-guidance.md
  - docs/00-project/quests/0024-add-a-runtime-diagnostic-value.md
  - docs/00-project/quests/0027-add-an-app-startup-diagnostic-snapshot.md
last_updated: 2026-07-06
---

# QUEST-0028 - Document Runtime vs Startup Diagnostics

> Label the two diagnostic charms before future log and CLI lanterns reach into
> the satchel and mistake the current badge for the startup relic.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- `docs/04-architecture/system-overview.md` now records that
  `App.RuntimeDiagnostic()` represents the current runtime state.
- The same architecture note records that `App.StartupDiagnostic()` returns the
  diagnostic captured when startup completed.
- The note makes the shutdown boundary explicit: `StartupDiagnostic` remains
  the startup-complete snapshot after shutdown, while the current runtime
  diagnostic may report the stopped badge.
- Neither diagnostic is wired to daemon logs or CLI output yet.
- `RuntimeSnapshotSummary.String()` remains the stable internal machine-ish
  string, and future user-facing output is still reserved for a separate
  friendly formatter.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation-only change.
- No Go code changes.
- No test changes.
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

The diagnostic map now marks the difference between the live runtime lantern and
the startup-complete relic, keeping future log and CLI quests from mixing their
signals.
