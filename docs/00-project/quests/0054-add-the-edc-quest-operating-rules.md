---
id: QUEST-0054
title: Add the EDC Quest Operating Rules
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a documentation-only rulebook that makes the Engineering Documentation Codex the source of truth for future Codex and Claude quest work.
related:
  - docs/00-project/edc-quest-operating-rules.md
  - README.md
  - PROJECT_MANIFEST.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
last_updated: 2026-07-07
---

# QUEST-0054 - Add the EDC Quest Operating Rules

> Set the operating compass beside the camp ledger: future agents read the EDC,
> name their source-of-truth files, take small safe steps, and pause when docs
> and implementation disagree.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the EDC quest operating rules at
  [docs/00-project/edc-quest-operating-rules.md](../edc-quest-operating-rules.md).
- Framed the Engineering Documentation Codex as project memory, safety rail,
  source of truth, and camp ledger rather than decorative prose.
- Added the required pre-flight ritual for future quests: read the quest, read
  relevant EDC docs, identify source-of-truth files, preserve hard safety
  boundaries, and keep the change small and reviewable.
- Added the source-of-truth rule: if code and docs disagree, stop, report the
  conflict, propose whether code or docs should change, and do not silently
  rewrite project intent.
- Added the quest report rule for pushed commit hash, validation, diff stat,
  branch status, EDC docs consulted, docs updated, and new findings.
- Added the discovery rule for hardware, install, display, docking, audio,
  GPIO, and user-support behavior so findings land in EDC ledgers or maps.
- Added small-step coding guidance favoring narrow quests, testable seams,
  dry-run behavior, read-only probes, fake paths, local diagnostics, and docs
  before installers.
- Added the project voice rule, preferred artifact words, avoided language,
  and a copy-paste agent reminder block for future Codex and Claude prompts.
- Added a source-doc link map for the manifest, README, direction maps,
  SafeShutdown boundary map, installer map, diagnostics map, acceptance
  checklist, and GPIO probe ledger.
- Linked the new operating rules from `README.md`, `PROJECT_MANIFEST.md`,
  `docs/00-project/project-direction-save-room-arcadia.md`, and
  `docs/04-architecture/arcadia-runtime-migration-path.md`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No automation implementation.
- No CI changes.
- No installer implementation.
- No diagnostics implementation.
- No service install.
- No systemd activation.
- No `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No config mutation.
- No daemon activation.

## Milestone Note

The project now has an EDC operating compass. Future agents have a clear
ritual for reading, naming, updating, validating, reporting, and stopping when
the implementation and project ledger disagree.
