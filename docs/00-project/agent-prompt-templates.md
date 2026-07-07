---
id: AGENT-PROMPT-TEMPLATES-001
title: Agent Prompt Templates
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Provide reusable prompt blocks that keep future Claude and Codex sessions aligned with the Engineering Documentation Codex source-of-truth ritual.
related:
  - ../../CLAUDE.md
  - ../../PROJECT_MANIFEST.md
  - ../../PROJECT_MEMORY.md
  - docs/00-project/edc-quest-operating-rules.md
  - docs/00-project/documentation-structure-and-governance.md
  - docs/05-development/ai-collaboration.md
last_updated: 2026-07-07
---

# Agent Prompt Templates

> Hand the next helper a compass before handing them the wrench.

These templates are reusable starting points for future Claude, Codex, and
other AI-assisted RetroFlag Power sessions. They are documentation only. They
do not change code, automation, CI, installers, services, GPIO behavior,
shutdown behavior, or project configuration.

## EDC Quest Starter

Copy this into a future agent prompt when starting a quest:

```
Before editing, treat the Engineering Documentation Codex as the source of truth.

Start by reporting the session-start checklist: active quest or user request, EDC source-of-truth docs consulted, hard safety boundaries, intended files to edit, smallest safe change, and whether the quest is docs-only, code-only, or mixed. Do not edit until this checklist is complete, unless the user explicitly says to skip planning.

Read the quest or request, then read the relevant EDC docs before changing files. Name the source-of-truth docs you consulted in your final report. Do not append everything to PROJECT_MEMORY.md; store discoveries in the most specific ledger, map, hardware note, architecture doc, operations doc, or quest record that fits.

If code and docs disagree, stop and report the conflict before editing either side. Preserve the warm retro/adventurer voice in docs and reports, but keep code names plain, boring, and maintainable when clarity asks for it.
```

## Documentation-Only Quest Starter

Use this when the quest is explicitly docs/configuration only:

```
This is a documentation/configuration-only RetroFlag Power quest. Do not change Go code or runtime behavior.

Start by reporting the session-start checklist: active quest or user request, EDC source-of-truth docs consulted, hard safety boundaries, intended files to edit, smallest safe change, and whether this quest is docs-only. Do not edit until this checklist is complete, unless the user explicitly says to skip planning.

Before editing, read the relevant EDC docs and identify the source-of-truth files. Keep the change small and reviewable. Update the most specific EDC document for any durable discovery; do not use PROJECT_MEMORY.md as a dumping ground. If code and docs disagree, stop and report the conflict.

In the report, include consulted docs, docs updated, validation commands, branch status, diff stat, and the pushed commit hash if pushed. Keep the report warm, clear, and adventure-flavored without hiding technical facts.
```

## Report Checklist

Future quest reports should name:

- Pushed commit hash, when a push happened.
- Validation results.
- Branch status.
- Diff stat.
- EDC docs consulted.
- Docs updated.
- Any conflict between code and docs.
- Any new findings or follow-up quests.

The victory should be named plainly. A small verified badge still belongs in
the ledger.
