---
id: MILESTONES-001
title: Verified Milestone Ledger
version: 0.2.0
status: Active
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Provide the canonical ledger of numbered, verified RetroFlag Power checkpoints with stable anchors and evidence.
related:
  - PROJECT_MEMORY.md
  - docs/00-project/roadmap.md
  - docs/00-project/quests/
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
  - docs/13-reference/terminology.md
supersedes:
  - docs/99-archive/project-milestones-pre-edc.md
last_updated: 2026-07-06
---

# Verified Milestone Ledger

> Epochs tell the journey. Milestones prove the distance traveled.

This is the canonical source for numbered, verified project checkpoints.

No Milestones are created by this vocabulary migration. QUEST-0006 will gather
existing checkpoint evidence, define the first entries, and remove duplicated
checkpoint prose only after review.

---

# Milestone Rules

## Identity

Milestones use stable, sequential IDs:

```text
M-0001
M-0002
M-0003
```

IDs remain stable after publication.

## Verification

A Milestone records a checkpoint only after evidence verifies it. Planning work
and implementation without verification remain in Quests, requirements, or the
Roadmap.

Each entry must include:

- stable anchor and ID
- title
- status
- verification date
- concise verified fact
- evidence
- related Quest or Quests
- related commit or revision when available
- related ADR when implementation reasoning needs durable explanation

## Stable anchors

Every entry uses an explicit anchor:

```html
<a id="m-0001"></a>
```

followed by:

```markdown
## M-0001 — Example Title
```

Other documents cite the anchored Milestone instead of copying its evidence or
completion prose.

## Milestone-scale ADRs

If a future contributor would reasonably ask why a Milestone was implemented
that way, create an ADR in `docs/adr/` and cross-link the ADR and Milestone.

---

# Milestone Entry Template

Do not copy this template into a real entry until its evidence is ready.

```markdown
<a id="m-NNNN"></a>
## M-NNNN — Verified Checkpoint Title

Status: Verified
Verified on: YYYY-MM-DD

Verified fact:

- concise statement

Evidence:

- command, test, artifact, or hardware result

Related:

- QUEST-NNNN
- docs/adr/NNNN-decision.md, when needed
- commit or revision
```

---

# Verified Milestones

No entries yet.

QUEST-0006 — Gather the Checkpoints into One Ledger will propose the initial
entries without changing their historical evidence.
