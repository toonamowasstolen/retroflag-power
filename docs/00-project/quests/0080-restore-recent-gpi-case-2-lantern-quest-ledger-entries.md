---
id: QUEST-0080
title: Restore Recent GPi Case 2 Lantern Quest Ledger Entries
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Restore the quest ledger entries for recent GPi Case 2 and Arcadia Runtime Lantern work so completed missions have numbered trail markers again.
related:
  - 0076-add-a-portable-gpi-case-2-boot-power-trace-field-lantern-script.md
  - 0077-add-a-human-facing-field-lantern-script-ux-standard.md
  - 0078-retrofit-the-gpi-case-2-manual-lantern-script-to-the-human-facing-ux-standard.md
  - 0079-record-successful-gpi-case-2-resume-evidence.md
  - ../edc-quest-operating-rules.md
last_updated: 2026-07-09
---

# QUEST-0080 - Restore Recent GPi Case 2 Lantern Quest Ledger Entries

> Repair the map without moving the trail: every recent Lantern win gets its
> quest number, validation record, and commit trail marker back in the satchel.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Confirmed that QUEST-0076 already existed for the portable GPi Case 2 Bundle
  Collector Field Lantern script.
- Confirmed that QUEST-0077 already existed for the Human-Facing Field Lantern
  Script UX Standard and preserved it.
- Added QUEST-0078 for the GPi Case 2 manual Lantern script retrofit to the
  human-facing UX standard.
- Added QUEST-0079 for the successful resume evidence and updated intermittent
  resume-wedge theory.
- Added completion commit trail markers to the restored and confirmed recent
  quest files.
- Updated the EDC quest operating rules so future Codex missions must carry an
  explicit quest number, quest file, validation results, and commit hash.

## Boundary

- Documentation only.
- No runtime code changes.
- No shell script changes.
- No GPIO behavior changes.
- No systemd changes.
- No shutdown behavior changes.
- No `SafeShutdown.py` changes.
- No installer changes.
- No telemetry changes.
- No hardware instruction changes.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- This restoration quest is completed by the commit that adds this file. The
  final pushed commit hash is reported in the quest handoff because a Git
  commit cannot contain its own final object hash.

## Milestone Note

The Ledger is bright again from QUEST-0076 through QUEST-0080. The older map
still has historic numbering gaps, but the recent GPi Case 2 Lantern run now
has a clean row of trail markers.
