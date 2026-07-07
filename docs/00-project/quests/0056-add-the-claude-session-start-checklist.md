---
id: QUEST-0056
title: Add The Claude Session Start Checklist
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a short Claude and Codex session-start checklist so future agents name the quest, source-of-truth docs, safety boundaries, and intended files before editing.
related:
  - ../../../CLAUDE.md
  - ../agent-prompt-templates.md
  - ../edc-quest-operating-rules.md
  - ../documentation-structure-and-governance.md
  - ../../05-development/ai-collaboration.md
last_updated: 2026-07-07
---

# QUEST-0056 - Add The Claude Session Start Checklist

> Light the lantern before the tools come out: name the quest, name the maps,
> name the safe path, then begin.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added a session-start checklist to [CLAUDE.md](../../../CLAUDE.md).
- Required future agents to begin by reporting the active quest or request,
  EDC source-of-truth docs consulted, hard safety boundaries, intended files,
  smallest safe change, and whether the quest is docs-only, code-only, or
  mixed.
- Added the explicit rule that agents should not edit until the checklist is
  complete, unless the user says to skip planning.
- Updated [Agent Prompt Templates](../agent-prompt-templates.md) so reusable
  starter prompts carry the same checklist in the satchel.
- Refreshed the EDC operating rules reminder block to mention the checklist
  where future agents are most likely to copy it.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No runtime behavior changes.
- No installer implementation.
- No diagnostics implementation.
- No service, `systemd`, or `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No config mutation.
- No telemetry.
- No network calls during implementation.

## Milestone Note

Future sessions now begin at the map table. The agent names the quest, lights
the lantern, checks the hard boundaries, and points to the intended path before
touching the files.
