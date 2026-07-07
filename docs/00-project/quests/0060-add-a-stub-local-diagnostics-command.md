---
id: QUEST-0060
title: Add a Stub Local Diagnostics Command
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add the smallest safe CLI doorway for future local diagnostics without collecting diagnostics yet.
related:
  - CLAUDE.md
  - docs/00-project/edc-quest-operating-rules.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/03-operations/gpi-case-2-acceptance-checklist.md
  - docs/03-operations/gpi-case-gpio-probe-ledger.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
last_updated: 2026-07-07
---

# QUEST-0060 - Add a Stub Local Diagnostics Command

> The diagnostics lantern now has a doorway, but it does not gather the field
> notes yet.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Mixed

## Outcome

- Added `retroflag-powerd diagnostics` as a stub-only CLI path.
- The command prints a short local-only/read-only diagnostics message.
- The command clearly says diagnostics collection is not implemented yet.
- The command exits successfully without starting runtime behavior.
- Added deterministic, hardware-free CLI test coverage for the stub output.
- Updated the local diagnostics map to record that the command exists only as
  a stub doorway.

## Boundary

- No GPIO writes.
- No GPIO reads.
- No shutdown execution.
- No systemd activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume implementation.
- No persistent state.
- No telemetry.
- No network calls.
- No installer or packaging changes.
- No diagnostics bundle generation.
- No OS, display, audio, or process fact collection.
- No file writing by the command.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Milestone Note

The local diagnostics trail now has a safe command doorway, with every larger
diagnostics behavior still waiting behind future quests.
