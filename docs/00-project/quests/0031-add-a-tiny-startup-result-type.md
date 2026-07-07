---
id: QUEST-0031
title: Add a Tiny Startup Result Type
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Give startup success its own small badge before future failure paths add more map detail.
related:
  - internal/app
  - docs/00-project/quests/0027-add-an-app-startup-diagnostic-snapshot.md
  - docs/00-project/quests/0029-add-an-app-startup-result-accessor.md
  - docs/00-project/quests/0030-document-startup-success-meaning.md
last_updated: 2026-07-07
---

# QUEST-0031 - Add a Tiny Startup Result Type

> Move the startup success charm onto its own tiny badge, so future quests can
> expand the ledger without asking the diagnostic lantern to carry every clue.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added a small internal `StartupResult` value with `Completed` and
  `Succeeded` fields.
- `App` now stores the startup result separately from startup diagnostic
  availability.
- `App.StartupResult()` returns the current startup result badge.
- `App.StartupSucceeded()` now reads `StartupResult().Succeeded`.
- Startup diagnostic availability remains unchanged.
- Successful startup sets the result to completed and succeeded at the ready
  point.
- Focused app tests prove the zero result before startup, the successful result
  after startup, the preserved successful result after shutdown, and the
  `StartupSucceeded()` match against `StartupResult().Succeeded`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No daemon log changes.
- No CLI output changes.
- No lifecycle behavior changes.
- No planner behavior changes.
- No executor behavior changes.
- No failure taxonomy.
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

Startup success now has its own tiny badge in the app satchel. The diagnostic
lantern still shines the same way, while the map has a cleaner place for future
startup result detail.
