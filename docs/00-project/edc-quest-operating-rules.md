---
id: EDC-QUEST-OPERATING-RULES-001
title: EDC Quest Operating Rules
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Define the source-of-truth ritual for future RetroFlag Power quest work so agents keep the Engineering Documentation Codex current, safe, and warm.
related:
  - PROJECT_MANIFEST.md
  - README.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/03-operations/installer-migration-toolkit-map.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/03-operations/gpi-case-2-acceptance-checklist.md
  - docs/03-operations/gpi-case-gpio-probe-ledger.md
  - docs/00-project/quests/0054-add-the-edc-quest-operating-rules.md
  - docs/00-project/quests/0056-add-the-claude-session-start-checklist.md
last_updated: 2026-07-07
---

# EDC Quest Operating Rules

> The Engineering Documentation Codex is the camp ledger, compass, and source
> of truth for the adventure. Keep it close before the next lantern is lit.

This document is documentation only. It does not change Go code, implement
automation, alter CI, install services, activate daemons, mutate configuration,
alter `rc.local`, replace `/opt/RetroFlag/SafeShutdown.py`, write GPIO, or
execute shutdown.

## Purpose

The EDC is not decorative.

It is the project memory, safety rail, source of truth, and camp ledger for
RetroFlag Power. It carries the reasons behind the code, the current hardware
unknowns, the safety boundaries, the accepted vocabulary, and the verified
badges earned along the way.

Future Codex, Claude, and human quest work should treat the EDC as part of the
implementation surface. If the docs are stale, the project is partly stale. If
the docs and implementation disagree, the quest must pause before the project
intent is silently rewritten.

## Required Pre-Flight For Every Quest

Before editing, every future quest agent must:

- Read the quest file or user request.
- Read the relevant EDC docs.
- Identify the source-of-truth documents before editing.
- Preserve all hard safety boundaries named by the quest and the EDC.
- Keep the change small, reviewable, and easy to validate.

The pre-flight is not ceremony for ceremony's sake. It is how the project keeps
its compass aligned while many small quests add up to a larger machine.

## Source-Of-Truth Rule

When code and EDC docs disagree:

- Do not guess.
- Stop and report the conflict.
- Name the files or behavior that disagree.
- Propose whether code or docs should change.
- Do not silently rewrite project intent.

If the conflict touches power behavior, GPIO behavior, display behavior,
docking behavior, audio behavior, install behavior, diagnostics, service
activation, or `SafeShutdown.py`, treat the stop as a safety badge, not a
failure.

## Quest Report Rule

Every quest report should include:

- Pushed commit hash.
- Validation results.
- Diff stat.
- Branch status.
- EDC docs consulted.
- Docs updated.
- New findings or follow-up quests.

The report should name the victory, not only the diff. A small verified win is
still a real badge for the ledger.

## Discovery Rule

If a quest discovers hardware behavior, install behavior, display behavior,
docking behavior, audio behavior, GPIO behavior, or user-support behavior:

- Record it in the appropriate EDC ledger or map.
- Do not leave it only in chat.
- Do not bury it only in code comments.

Discovery belongs where future maintainers and future agents will look first.
The chat can carry the spark; the EDC must carry the durable lantern.

## Small-Step Coding Rule

Prefer:

- Narrow quests.
- Testable seams.
- Dry-run behavior.
- Read-only probes before mutation.
- Fake paths before hardware paths.
- Local diagnostics before network submission.
- Docs before installers.

The project should earn each risky step with a smaller verified step before it.
That is especially important for real shutdown behavior, real GPIO paths,
installer trails, config mutation, and service activation.

## Project Voice Rule

Preserve the warm retro/adventurer tone in:

- Quests.
- Docs.
- Milestones.
- Reports.
- Handoffs.

Celebrate small verified wins. Name the victory, not only the diff.

Keep Go code, exported types, package names, filenames, and low-level
technical identifiers clear and maintainable. Do not force RPG naming into
code when plain engineering names are clearer.

## Preferred And Avoided Language

Preferred project artifact words:

- toolkit
- lantern
- badge
- satchel
- charm
- ledger
- compass
- quest
- epoch
- map
- relic
- spellbook
- field kit

Avoid restraint or control metaphors unless discussing language to avoid:

- harness
- cage
- leash
- lockstep

Warmth should make the project easier to understand. It should never obscure a
safety boundary, validation result, command, error, or implementation detail.

## Agent Reminder Block

Copy this into future Codex or Claude prompts when starting a quest:

```
Before editing, read the relevant EDC docs and list the source-of-truth files you used. Start by naming the active quest or request, hard safety boundaries, intended files, smallest safe change, and whether the quest is docs-only, code-only, or mixed. Do not edit until that checklist is complete, unless the user explicitly says to skip planning. Keep the quest small, update the EDC with any discoveries, and keep the RetroFlag Power warm retro/adventurer voice in docs and reports. If code and EDC disagree, stop and report the conflict.
```

## Link Map

Likely source-of-truth docs for future agents:

- [Project Manifest](../../PROJECT_MANIFEST.md)
- [README](../../README.md)
- [Save Room Tech and Arcadia Runtime Direction](project-direction-save-room-arcadia.md)
- [Arcadia Runtime Migration Path](../04-architecture/arcadia-runtime-migration-path.md)
- [SafeShutdown Replacement Boundary Map](../03-operations/safeshutdown-replacement-boundary-map.md)
- [Installer and Migration Toolkit Map](../03-operations/installer-migration-toolkit-map.md)
- [Local Diagnostics Bundle Map](../03-operations/local-diagnostics-bundle-map.md)
- [GPi Case 2 Acceptance Checklist](../03-operations/gpi-case-2-acceptance-checklist.md)
- [GPi Case GPIO Probe Field Ledger](../03-operations/gpi-case-gpio-probe-ledger.md)
