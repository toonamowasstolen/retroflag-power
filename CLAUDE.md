---
id: CLAUDE-INSTRUCTIONS-001
title: Claude Code Instructions
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - AI Assistants
  - Project Maintainers
purpose: Give Claude Code sessions a repo-root source-of-truth ritual for RetroFlag Power work.
related:
  - PROJECT_MANIFEST.md
  - PROJECT_MEMORY.md
  - docs/00-project/edc-quest-operating-rules.md
  - docs/00-project/documentation-structure-and-governance.md
  - docs/00-project/agent-prompt-templates.md
  - docs/05-development/ai-collaboration.md
last_updated: 2026-07-07
---

# Claude Code Instructions

> Read the camp ledger before lifting the tools.

The Engineering Documentation Codex is the source of truth for RetroFlag
Power. Treat the EDC as part of the implementation surface, not as decorative
prose beside the code.

## Session Start Checklist

Begin each quest by reporting:

- Active quest or user request.
- EDC source-of-truth docs consulted.
- Hard safety boundaries for this quest.
- Intended files to edit.
- Smallest safe change.
- Whether the quest is docs-only, code-only, or mixed.

Do not edit until this checklist is complete, unless the user explicitly says
to skip planning.

## Required Ritual

Before editing:

- Read the user request or quest record.
- Read the relevant EDC docs before editing.
- Identify the source-of-truth docs for the work.
- Stop if code and docs disagree.
- Report the disagreement clearly before changing either side.

After editing:

- Name the consulted docs in the report.
- Name the docs updated in the report.
- Include validation results and branch status.
- Preserve the warm retro/adventurer voice in docs and reports.

## Memory Rule

Do not append everything to `PROJECT_MEMORY.md`.

Store durable discoveries in the most specific EDC home:

- Hardware findings belong in hardware docs or hardware ledgers.
- GPIO findings belong in GPIO ledgers or maps.
- Installer and migration findings belong in operations maps.
- Architecture direction belongs in architecture docs.
- Quest outcomes belong in quest records.
- Project origin, principles, and uncategorized important memory may live in
  `PROJECT_MEMORY.md` until a more specific home exists.

The chat can carry the spark. The EDC must carry the durable lantern.

## Voice And Naming

Keep documentation, quest records, and reports warm, clear, and lightly
adventurous.

Keep code names plain and maintainable. Do not force retro/adventurer language
into package names, exported types, filenames, commands, or low-level technical
identifiers when plain engineering names are clearer.

Clarity wins whenever personality and precision compete.
