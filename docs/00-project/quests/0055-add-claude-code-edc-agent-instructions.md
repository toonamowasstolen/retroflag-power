---
id: QUEST-0055
title: Add Claude Code EDC Agent Instructions
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add repo-root Claude Code instructions and reusable agent prompt templates so future assistant sessions follow the EDC source-of-truth ritual.
related:
  - ../../../CLAUDE.md
  - ../agent-prompt-templates.md
  - ../edc-quest-operating-rules.md
  - ../documentation-structure-and-governance.md
  - ../../05-development/ai-collaboration.md
last_updated: 2026-07-07
---

# QUEST-0055 - Add Claude Code EDC Agent Instructions

> Put the compass at the repo gate: future Claude and Codex sessions read the
> EDC first, name the maps they used, and keep discoveries in the right ledger.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added repo-root Claude Code instructions at [CLAUDE.md](../../../CLAUDE.md).
- Added reusable agent prompt templates at
  [docs/00-project/agent-prompt-templates.md](../agent-prompt-templates.md).
- Reaffirmed that the Engineering Documentation Codex is the source of truth.
- Required future agents to read relevant docs before editing and name
  consulted docs in their reports.
- Made `PROJECT_MEMORY.md` a safety net for origin, principles, and
  uncategorized important memory rather than a dumping ground.
- Directed durable discoveries to the most specific ledger, map, quest,
  hardware note, operations doc, or architecture doc.
- Required agents to stop when code and docs disagree.
- Preserved the warm retro/adventurer voice for docs and reports while keeping
  code names plain and maintainable.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation and configuration only.
- No Go code changes.
- No runtime behavior changes.
- No automation implementation.
- No CI changes.
- No installer implementation.
- No service activation.
- No config mutation beyond adding repo documentation files.

## Milestone Note

The repo now has a Claude-facing gate sign and a reusable prompt satchel. Future
assistant sessions have a clearer path back to the EDC before they touch the
workbench.
