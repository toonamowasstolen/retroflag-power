---
id: QUEST-0022
title: Record Runtime Summary Formatter Guidance
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the formatter boundary for the runtime summary string before user-facing diagnostic lanterns arrive.
related:
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0020-add-a-runtime-snapshot-summary.md
  - docs/00-project/quests/0021-add-a-runtime-summary-string.md
last_updated: 2026-07-06
---

# QUEST-0022 - Record Runtime Summary Formatter Guidance

> Mark the edge of the runtime summary charm before a future lantern tries to
> make it prettier than its job requires.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- `docs/04-architecture/system-overview.md` now records that
  `RuntimeSnapshotSummary.String()` is stable, compact, deterministic, and
  internal.
- The guidance keeps the string test-friendly and machine-ish for internal
  diagnostics.
- Future user-facing diagnostics are directed toward a separate friendly
  formatter instead of changing the stable internal string.
- CLI output, daemon logs, lifecycle behavior, planner behavior, executor
  behavior, packaging, hardware behavior, and state storage remain unchanged.

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

This preserves the runtime map's internal formatter boundary. It does not wire
the summary string into any user-facing lantern.
