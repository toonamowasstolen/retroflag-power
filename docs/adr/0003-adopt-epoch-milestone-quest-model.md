---
id: ADR-0003
title: Adopt the Epoch, Milestone, and Quest Model
version: 0.1.0
status: Accepted
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the EDC vocabulary and ownership model for project life stages, verified checkpoints, work records, current direction, and durable project memory.
related:
  - PROJECT_MEMORY.md
  - PROJECT_MANIFEST.md
  - docs/00-project/documentation-structure-and-governance.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/00-project/quests/0004-chart-the-edc-map.md
  - docs/00-project/quests/0005-correct-the-edc-compass.md
  - docs/13-reference/terminology.md
last_updated: 2026-07-06
---

# ADR-0003 — Adopt the Epoch, Milestone, and Quest Model

> Give each kind of progress one name and one home.

## Status

Accepted

## Date

2026-07-06

## Owner

Joshua Taft

## Related Documents

- `docs/00-project/quests/0004-chart-the-edc-map.md`
- `docs/00-project/quests/0005-correct-the-edc-compass.md`
- `docs/00-project/roadmap.md`
- `docs/00-project/milestones.md`
- `docs/00-project/documentation-structure-and-governance.md`
- `PROJECT_MEMORY.md`

---

# 1. Context

RetroFlag Power used `Milestone` for major life stages. Two competing ladders
developed:

- an older seven-stage narrative in `PROJECT_MEMORY.md`
- an expanded eleven-stage journey in the former milestone document and Roadmap

Recent implementation wins were called checkpoints and copied into Project
Memory, the Roadmap, and Awakening readiness. They had no stable IDs or
canonical verified ledger.

QUEST-0004 found that:

- the large-stage vocabulary was inconsistent
- `Milestone` did not currently mean a small verified checkpoint
- Project Memory was becoming a progress log
- current progress facts were duplicated
- citations pointed to files or generated headings rather than stable records

The project needs distinct names for long life stages, verified progress, and
the work that produces progress.

---

# 2. Decision

RetroFlag Power adopts this EDC model:

## Epoch

An Epoch is a large project life stage.

The canonical seven-Epoch ladder is:

```text
Dreaming → Awakening → Heartbeat → Memory → Momentum → Adventure → Launch
```

The Roadmap owns the canonical ladder, current Epoch, and future direction.

## Milestone

A Milestone is a numbered, verified checkpoint.

Milestones use stable IDs:

```text
M-0001
M-0002
M-0003
```

The canonical ledger is `docs/00-project/milestones.md`.

Milestone entries require explicit anchors such as:

```html
<a id="m-0001"></a>
```

## Quest

A Quest is a task or work record.

Quests define scope, safety boundaries, acceptance criteria, implementation
outcomes, and validation evidence. A Quest may contribute to zero, one, or
several Milestones. A Milestone may be supported by several Quests.

## Roadmap

The Roadmap records current state and future direction.

It identifies the current Epoch, next useful work, upcoming gates, deferred
scope, and route themes. It may summarize Milestones but must cite the canonical
ledger rather than duplicate verified evidence.

## Project Memory

`PROJECT_MEMORY.md` preserves origin, principles, safety context, durable risks,
important unresolved ideas, and knowledge that has not yet found a narrower
home.

Project Memory is not the progress log or verified Milestone ledger.

---

# 3. Seven Epochs and the Former Eleven Stages

The seven-stage ladder is preferred because it preserves the original narrative
and keeps life stages memorable:

1. Epoch 0 — Dreaming
2. Epoch 1 — Awakening
3. Epoch 2 — Heartbeat
4. Epoch 3 — Memory
5. Epoch 4 — Momentum
6. Epoch 5 — Adventure
7. Epoch 6 — Launch

Useful detail from the eleven-stage journey remains as Roadmap themes:

- Power work informs Heartbeat.
- Resume work belongs within Memory.
- Polish and Expansion shape Adventure.
- Release preparation leads into Launch.

The full former journey is preserved in
`docs/99-archive/project-milestones-pre-edc.md` as historical context, not as a
competing canonical model.

---

# 4. Milestones and ADRs

A Milestone proves that something is true. An ADR explains why an important
implementation or project decision was made.

If a future contributor would reasonably ask why a Milestone was implemented
that way:

1. create an ADR in `docs/adr/`
2. link the ADR from the Milestone
3. link the Milestone from the ADR

Simple verification does not require an ADR. Durable reasoning does.

Actual ADRs live in `docs/adr/`. The ADR template lives in
`docs/10-decisions/`.

Actual RFCs live in `docs/rfc/`. The RFC template lives in `docs/11-rfc/`.

---

# 5. Consequences

## Positive Consequences

- Large life stages and verified checkpoints no longer share one term.
- The seven-Epoch project story becomes canonical.
- Verified facts gain stable IDs and anchors.
- Quests retain their useful work-record role.
- The Roadmap has clear ownership of current state and direction.
- Project Memory can return to its safety-net role.
- Citations can target stable Milestones instead of duplicated prose.

## Negative Consequences

- Existing documents require a staged vocabulary migration.
- Historical prose will continue using old terms until deliberately updated or
  archived.
- The checkpoint ledger remains empty until QUEST-0006 reviews evidence.

## Neutral Consequences

- Existing checkpoint prose remains in place during this decision.
- No M-0001 through M-0004 entries are created by this ADR.
- Runtime behavior and packaging are unaffected.

---

# 6. Alternatives Considered

## Keep the eleven large Milestones

Rejected because it preserves the overloaded term and competes with the older
seven-stage narrative.

## Keep the seven large Milestones

Rejected because verified checkpoints would still need another stable term and
the original overload would remain.

## Use “Phase” for large stages

Not chosen because `Epoch` better matches the EDC project voice and clearly
distinguishes long life stages from tactical Roadmap phases.

## Create Milestones immediately

Deferred. Existing checkpoint claims must be gathered and verified before
stable M-IDs are assigned.

---

# 7. Migration Plan

1. Adopt vocabulary and ownership rules in governance and reference docs.
2. Make the seven-Epoch ladder canonical in the Roadmap.
3. Preserve the former eleven-stage journey in the archive.
4. Repurpose `docs/00-project/milestones.md` as the verified ledger.
5. Leave duplicated checkpoint prose in place during this Quest.
6. Use QUEST-0006 to create reviewed Milestone entries and remove duplication.
7. Update remaining historical terminology when touching affected documents.

---

# 8. Validation

The decision is correctly adopted when:

- governance defines all five EDC concepts
- the Roadmap contains one canonical Epoch ladder
- the Milestone ledger defines stable M-IDs and anchors
- Project Memory says it is not a progress log
- terminology and glossary entries agree
- old eleven-stage detail is preserved but clearly non-canonical
- actual ADR/RFC and template locations remain explicit
- no runtime or packaging behavior changes

---

# 9. Revision Plan

If this model creates confusion, supersede this ADR with a new decision. Do not
silently reuse Epoch, Milestone, or Quest for different concepts.

Stable Milestone IDs must never be reassigned after publication.
