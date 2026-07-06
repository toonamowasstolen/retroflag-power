---
id: QUEST-0005
title: Correct the EDC Compass
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Documentation Authors
purpose: Adopt the Epoch, Milestone, Quest, Roadmap, and Project Memory vocabulary and ownership model after the QUEST-0004 audit.
related:
  - PROJECT_MEMORY.md
  - PROJECT_MANIFEST.md
  - docs/00-project/documentation-structure-and-governance.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/00-project/quests/0004-chart-the-edc-map.md
  - docs/13-reference/terminology.md
  - docs/13-reference/glossary.md
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
last_updated: 2026-07-06
---

# QUEST-0005 — Correct the EDC Compass

> The map was charted. Now make every needle point the same way.

## Quest Status

Implemented

## Quest Type

Documentation governance and vocabulary migration

---

# 1. Objective

Adopt one EDC model:

- Epoch = large project life stage
- Milestone = numbered verified checkpoint
- Quest = task or work record
- Roadmap = current state and future direction
- Project Memory = origin, principles, safety net, and important memory; not a
  progress log

---

# 2. Scope

## In Scope

- adopt the seven-Epoch ladder
- preserve the former eleven-stage journey as historical detail
- create the verified Milestone ledger scaffold
- record the decision in ADR-0003
- update governance and reference terminology
- define stable future Milestone anchors
- define Milestone-scale ADR guidance
- confirm ADR/RFC record and template locations
- cross-link affected project documents

## Out of Scope

- moving duplicated checkpoint prose
- assigning M-0001 through M-0004
- rewriting every historical progress document
- adding a link checker
- changing production code
- changing packaging
- activating systemd
- adding GPIO, shutdown execution, resume, or state storage

---

# 3. Adopted Model

## Epoch ladder

```text
Dreaming → Awakening → Heartbeat → Memory → Momentum → Adventure → Launch
```

The Roadmap is canonical for the ladder, current Epoch, and direction.

## Milestone ledger

`docs/00-project/milestones.md` is canonical for numbered, verified checkpoint
facts.

No Milestones were created by this Quest. QUEST-0006 will review existing
evidence before assigning stable IDs.

## Historical preservation

The former eleven-stage document is preserved at:

```text
docs/99-archive/project-milestones-pre-edc.md
```

Power, Resume, Polish, Expansion, and Release remain useful Roadmap themes
within the seven Epochs.

---

# 4. Migration Notes

Existing checkpoint summaries remain in Project Memory, the Roadmap, and
Awakening readiness until QUEST-0006 moves or summarizes them.

Historical readiness documents, requirement fields, ADRs, and Quests may still
use the former large-stage meaning of `Milestone`. Those references are
pre-ADR vocabulary, not a second canonical journey. Update them deliberately
when their owning documents are next revised.

Do not interpret old `Milestone: Power` or similar planning labels as verified
M-IDs.

---

# 5. Acceptance Criteria

- [x] ADR-0003 records the vocabulary decision.
- [x] One seven-Epoch ladder is canonical.
- [x] The eleven-stage journey is preserved but non-canonical.
- [x] `docs/00-project/milestones.md` is a verified ledger scaffold.
- [x] No M-0001 through M-0004 entries were created.
- [x] Governance defines all five EDC concepts.
- [x] Governance requires stable Milestone anchors.
- [x] Governance requires cross-linked ADRs for durable Milestone reasoning.
- [x] ADR and RFC record/template locations are explicit.
- [x] Terminology and glossary definitions agree.
- [x] Project Memory is explicitly not the progress log.
- [x] Existing checkpoint prose was not moved.
- [x] No production code or packaging changed.

---

# 6. Outcome

The EDC compass now points to one model.

The vocabulary conflict is resolved without erasing history or prematurely
claiming verified Milestones. QUEST-0006 remains responsible for gathering
checkpoint evidence into the ledger and removing duplicated progress prose.

---

# 7. Validation

Completed before commit:

- [x] `git diff --check` passed.
- [x] `make check` passed.
- [x] Changed Markdown links and metadata targets were inspected.
- [x] `git diff --stat` was reviewed.
- [x] `git diff` was reviewed.

---

# Closing

Epochs tell the journey.

Milestones prove progress.

Quests do the work.
