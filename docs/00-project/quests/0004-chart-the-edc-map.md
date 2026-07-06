---
id: QUEST-0004
title: Chart the EDC Map
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Documentation Authors
purpose: Audit the current RetroFlag Power documentation model, terminology, progress ownership, duplication, and citation patterns before changing the EDC structure.
related:
  - PROJECT_MEMORY.md
  - docs/00-project/documentation-structure-and-governance.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/00-project/awakening-readiness.md
  - docs/00-project/quests/
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
last_updated: 2026-07-06
---

# QUEST-0004 — Chart the EDC Map

> Before moving the signposts, draw the road that already exists.

## Quest Status

Implemented

## Quest Type

Documentation model audit

---

# 1. Audit Boundary

This quest records the current EDC documentation model before any correction is
made.

This quest does not:

- rewrite the EDC model
- rename `Milestone` to `Epoch`
- redefine `Milestone`
- move checkpoint prose
- create a checkpoint ledger
- change production Go code
- change packaging

The findings below describe the repository as it exists on 2026-07-06.

Follow-up note: QUEST-0005 and ADR-0003 adopted the corrected model later that
day. The old terminology below is retained as audit evidence, not current
guidance.

---

# 2. Documents Audited

Primary progress documents:

- `PROJECT_MEMORY.md`
- `docs/00-project/roadmap.md`
- `docs/00-project/awakening-readiness.md`
- `docs/00-project/milestones.md`

Governance and record locations:

- `docs/00-project/documentation-structure-and-governance.md`
- `docs/00-project/quests/`

The audit searched headings, metadata relationships, prose references, progress
checklists, checkpoint notes, and current-state statements.

---

# 3. Where “Milestone” Means the Large Journey

`Milestone` currently means a major project life stage throughout the formal
documentation model.

## 3.1 Seven-stage narrative in `PROJECT_MEMORY.md`

`PROJECT_MEMORY.md`, under `# 5. Narrative Milestones`, defines this seven-stage
journey:

1. Milestone 0 — Dreaming
2. Milestone 1 — Awakening
3. Milestone 2 — Heartbeat
4. Milestone 3 — Memory
5. Milestone 4 — Momentum
6. Milestone 5 — Adventure
7. Milestone 6 — Launch

This is the clearest location where `Milestone` means the older seven-stage
story.

## 3.2 Eleven-stage journey in `milestones.md`

`docs/00-project/milestones.md` expands the journey to eleven stages:

1. Milestone 0 — Dreaming
2. Milestone 1 — Awakening
3. Milestone 2 — Heartbeat
4. Milestone 3 — Power
5. Milestone 4 — Memory
6. Milestone 5 — Resume
7. Milestone 6 — Momentum
8. Milestone 7 — Polish
9. Milestone 8 — Expansion
10. Milestone 9 — Release
11. Milestone 10 — Launch

Each entry has phase-scale purpose, included work, excluded work, exit criteria,
risks, and a victory condition.

## 3.3 Other large-stage uses

- `docs/00-project/roadmap.md` repeats the eleven-stage sequence in its
  medium-term roadmap and explicitly says, “A milestone is the chapter.”
- `docs/00-project/awakening-readiness.md` uses Milestone 0 and Milestone 1 as
  the transition from Dreaming to Awakening.
- The governance guide defines milestones as “large phases” and says they answer
  “What phase are we in?”
- Quest guidance says a quest is smaller than a milestone.

Conclusion: current formal usage consistently treats `Milestone` as a large
project phase, but two incompatible phase lists exist.

---

# 4. Where “milestone” Means a Small Checkpoint

No consistent current usage was found where `milestone` means a small numbered,
verified checkpoint.

Small completed steps are currently called checkpoints:

- Nameplate checkpoint
- Config satchel checkpoint
- Event charms checkpoint
- Dry-run action charm checkpoint

These are not numbered and are not called milestones.

Other appearances of “checkpoint” in the milestone document refer to gameplay
or resume concepts, such as restoring a checkpoint or polishing a checkpoint
screen. They are not project-progress records.

Conclusion: redefining `Milestone` as a numbered verified checkpoint would be a
new EDC rule, not a cleanup of an already-consistent convention.

---

# 5. Duplicated Progress and Checkpoint Prose

Four recent checkpoints are repeated across three files:

| Checkpoint | `PROJECT_MEMORY.md` | `roadmap.md` | `awakening-readiness.md` | `milestones.md` |
| --- | --- | --- | --- | --- |
| Nameplate | Duplicate summary | Duplicate summary | Duplicate summary | Not recorded |
| Config satchel | Duplicate summary | Duplicate summary | Duplicate summary | Not recorded |
| Event charms | Duplicate summary | Duplicate summary | Duplicate summary | Not recorded |
| Dry-run action charm | Duplicate summary | Duplicate summary | Duplicate summary | Not recorded |

The wording varies slightly, but each copy repeats implementation facts,
validation results, and safety exclusions.

This creates four risks:

1. A completed checkpoint must be updated in three places.
2. Validation claims can drift between copies.
3. `PROJECT_MEMORY.md` is becoming a progress log despite governance calling it
   a safety net.
4. `awakening-readiness.md` is becoming an implementation journal even though
   its stated purpose is readiness and entry criteria.

`docs/00-project/milestones.md` has the opposite problem: its Milestone 1 exit
criteria remain unchecked and its “Immediate Next Path” still says the project
is moving through Milestone 0. It does not reflect the duplicated checkpoint
progress found elsewhere.

---

# 6. Current Canonical Ownership

The repository’s intended ownership and actual ownership are not always the
same.

| Concern | Intended source | Actual current behavior | Audit finding |
| --- | --- | --- | --- |
| Project phase | `docs/00-project/milestones.md` | `roadmap.md` says Milestone 1; `milestones.md` still describes the current path through Milestone 0 | Conflicted |
| Current progress | `docs/00-project/roadmap.md` | Progress is repeated in roadmap, readiness, and project memory | Duplicated |
| Verified checkpoints | No dedicated source | Evidence is scattered through checkpoint prose, quest outcomes, and unchecked criteria | Missing canonical ledger |
| Quest records | `docs/00-project/quests/` | Quest files hold scope, status, evidence, and outcomes | Clear and functioning |
| Safety memory | `PROJECT_MEMORY.md` | Safety context also appears in readiness, requirements, architecture, and quests | `PROJECT_MEMORY.md` is the broad safety net; requirements remain normative |

## 6.1 Project phase

Governance assigns large phases to `docs/00-project/milestones.md`, but that file
does not currently provide an accurate current-position statement.

The practical current phase is most accurately stated in
`docs/00-project/roadmap.md`:

```text
Milestone 1 — Awakening
```

## 6.2 Current progress

The roadmap is the best current candidate because governance says it owns the
practical route, next steps, and passed gates. Its checkpoint prose is still
duplicated elsewhere.

## 6.3 Verified checkpoints

There is no authoritative list of numbered checkpoints with:

- stable checkpoint ID
- verification status
- validation evidence
- completion date
- source quest or commit

Quest outcomes contain some of this evidence, but quests describe work records,
not a consolidated verified-progress history.

## 6.4 Quest records

`docs/00-project/quests/` is already the clear canonical home. Governance defines
the lifecycle:

```text
Draft → Active → Implemented → Verified → Archived
```

No correction is proposed for quest ownership.

## 6.5 Safety memory

`PROJECT_MEMORY.md` is explicitly named “The safety net” by governance. It should
retain durable risks, constraints, context, and ideas that have not found a more
specific home.

It should not be the canonical progress ledger. Normative safety requirements
belong in `docs/00-project/requirements.md`; focused validation boundaries
belong in their quests and test plans.

---

# 7. Broken or Fragile Citation Patterns

No missing `docs/00-project/milestones.md` file reference was found. The larger
problem is semantic and structural fragility.

## 7.1 Plain prose references

Most cross-document references say `Milestone 1`, `Awakening`, or “the current
milestone” as plain prose. They do not point to a stable identifier or a
specific canonical section.

Readers must guess whether the authority is:

- `PROJECT_MEMORY.md`
- `docs/00-project/milestones.md`
- `docs/00-project/roadmap.md`
- `docs/00-project/awakening-readiness.md`

## 7.2 File-only metadata relationships

Front matter commonly lists:

```yaml
related:
  - docs/00-project/milestones.md
```

This is useful for document-level discovery, but it cannot cite a specific
phase, checkpoint, criterion, or verified result.

## 7.3 Heading-anchor fragility

The checkpoint headings are repeated across three documents. A link such as
`#nameplate-checkpoint` would be locally valid in several files but would not
identify the canonical record.

Generated Markdown anchors also depend on heading text. Renaming a charm or
moving prose can silently break deep links.

## 7.4 Numbering drift

The seven-stage story and eleven-stage journey assign different numbers to
similarly named stages:

- Memory is Milestone 3 in `PROJECT_MEMORY.md`.
- Memory is Milestone 4 in `milestones.md` and `roadmap.md`.
- Momentum is Milestone 4 in `PROJECT_MEMORY.md`.
- Momentum is Milestone 6 in the expanded journey.
- Launch is Milestone 6 in `PROJECT_MEMORY.md`.
- Launch is Milestone 10 in the expanded journey.

A citation such as “Milestone 4” is therefore ambiguous without a document and
stage name.

## 7.5 Stale current-state citations

- `roadmap.md` identifies the current phase as Milestone 1 — Awakening.
- `milestones.md` still labels its immediate path as “Current path through
  Milestone 0.”
- `awakening-readiness.md` contains current implementation checkpoints even
  though its primary job is to decide readiness to enter Awakening.

These statements can all be linked successfully while disagreeing about what
they mean.

---

# 8. Proposed Correction

This audit recommends the following EDC vocabulary and ownership model for a
future correction quest:

## Epoch

An Epoch is a large project life stage.

The current large narrative stages—Dreaming, Awakening, Heartbeat, Power,
Memory, Resume, Momentum, Polish, Expansion, Release, and Launch—should be
evaluated and normalized as Epochs.

## Milestone

A Milestone is a numbered, verified checkpoint.

Each Milestone should have a stable ID, evidence, verification status, date,
and links to the work that established it.

## Quest

A Quest is a task or work record.

Quests retain their current role: define scope, guardrails, acceptance criteria,
implementation outcome, and validation evidence for a focused mission.

## Roadmap

The Roadmap owns current state and direction.

It should identify the current Epoch, the next useful Milestone or Quest, passed
gates, and what must wait. It should summarize rather than duplicate the full
checkpoint ledger.

## PROJECT_MEMORY

`PROJECT_MEMORY.md` remains the safety net, not the progress log.

It should preserve durable context, unresolved ideas, risks, constraints, and
memory that has not yet moved to a more specific canonical artifact.

---

# 9. Recommended Follow-Up Quests

## QUEST-0005 — Correct the EDC Compass

Define and apply the corrected vocabulary and ownership rules:

- Epoch = large project life stage
- Milestone = numbered verified checkpoint
- Quest = task/work record
- Roadmap = current state and direction
- PROJECT_MEMORY = safety net

This quest should decide migration and compatibility rules before performing
broad terminology replacement.

## QUEST-0006 — Gather the Checkpoints into One Ledger

Create the canonical checkpoint ledger, move or summarize duplicated checkpoint
prose, assign stable IDs, and connect verification evidence to quests and
commits.

This quest should preserve history while removing the need to update the same
progress fact in three documents.

---

# 10. Acceptance Criteria

- [x] Large-stage uses of `Milestone` were identified.
- [x] The seven-stage and eleven-stage journeys were compared.
- [x] Small-checkpoint terminology was audited.
- [x] Duplicated checkpoint prose was mapped across all four requested files.
- [x] Current canonical ownership was identified.
- [x] Fragile milestone and checkpoint citation patterns were documented.
- [x] The proposed Epoch/Milestone/Quest/Roadmap/PROJECT_MEMORY correction was
  summarized.
- [x] QUEST-0005 and QUEST-0006 were recommended.
- [x] No source documentation model was rewritten.
- [x] No production code or packaging was changed.

---

# 11. Outcome

The current map is now documented without moving any roads.

The audit found:

- one older seven-stage journey
- one expanded eleven-stage journey
- no established small-milestone convention
- four checkpoints duplicated across three progress documents
- no canonical verified-checkpoint ledger
- clear quest ownership
- conflicted project-phase ownership
- citation patterns that resolve files but not stable semantic records

This quest is Implemented because the audit and recommendations are recorded.
QUEST-0005 adopted the correction through ADR-0003. QUEST-0006 remains future
work for gathering verified checkpoints into the canonical ledger.

---

# Closing

The compass can be corrected now because the old map has been preserved.

Do not move the signposts until the follow-up quests begin.
