---
id: QUEST-0032
title: Document Startup Result and Diagnostic Ordering
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Mark the ordering between the startup result badge and startup diagnostic snapshot before future failure paths add more map detail.
related:
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0027-add-an-app-startup-diagnostic-snapshot.md
  - docs/00-project/quests/0028-document-runtime-vs-startup-diagnostics.md
  - docs/00-project/quests/0029-add-an-app-startup-result-accessor.md
  - docs/00-project/quests/0030-document-startup-success-meaning.md
  - docs/00-project/quests/0031-add-a-tiny-startup-result-type.md
last_updated: 2026-07-06
---

# QUEST-0032 - Document Startup Result and Diagnostic Ordering

> Keep the startup badge and diagnostic lantern labeled in the satchel, so
> future failure quests can add detail without mixing their signals.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- `docs/04-architecture/system-overview.md` now records that
  `App.StartupResult()` is the small startup-completion badge.
- The architecture note clarifies that `App.StartupDiagnostic()` is the
  captured startup-complete diagnostic snapshot.
- The note records that `App.StartupSucceeded()` reads
  `App.StartupResult().Succeeded`.
- The current dry-run/noop path is documented as establishing
  `StartupResult` and `StartupDiagnostic` during the same successful startup
  completion flow.
- The future map marker is explicit: if diagnostic capture can fail
  independently, startup ordering may need to carry more detail.

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
- No new failure taxonomy.
- No recovery report.
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

The startup result badge and diagnostic lantern now have their ordering written
in the architecture spellbook. Future startup failure work has a clearer place
to add detail if diagnostic capture ever becomes its own fallible step.
