---
id: QUEST-0006
title: Gather the Checkpoints into One Ledger
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Documentation Authors
purpose: Gather duplicated Awakening checkpoint evidence into the canonical verified Milestone ledger.
related:
  - PROJECT_MEMORY.md
  - docs/00-project/awakening-readiness.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/00-project/quests/0004-chart-the-edc-map.md
  - docs/00-project/quests/0005-correct-the-edc-compass.md
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
last_updated: 2026-07-06
---

# QUEST-0006 — Gather the Checkpoints into One Ledger

> Keep one ledger of proof, then let the rest of the map point to it.

## Quest Status

Implemented

## Quest Type

Documentation consolidation and Milestone migration

---

# 1. Objective

Gather duplicated Awakening checkpoint prose into
`docs/00-project/milestones.md`, the canonical verified Milestone ledger, and
replace the duplicate summaries with stable citations.

---

# 2. Scope

## In Scope

- create M-0001 through M-0004 from evidence already recorded in project docs
  and local revision history
- assign explicit, stable anchors to each Milestone
- preserve known validation dates and evidence
- record relevant deferred or absent features
- replace duplicated checkpoint prose in Project Memory, the Roadmap, and
  Awakening readiness with citations
- preserve each destination document's purpose and safety framing

## Out of Scope

- changing production Go code
- changing packaging
- adding status badge work or a link checker
- rewriting `CODE_OF_CONDUCT.md` or `CONTRIBUTING.md`
- adding a runtime ASCII banner
- activating systemd
- adding GPIO
- executing shutdown
- editing `rc.local`
- replacing `SafeShutdown.py`
- adding resume or state storage

---

# 3. Canonical Milestones

This Quest established:

- [M-0001 — Daemon Nameplate](../milestones.md#m-0001)
- [M-0002 — Config Satchel](../milestones.md#m-0002)
- [M-0003 — Event Charms](../milestones.md#m-0003)
- [M-0004 — Dry-Run Action Charm](../milestones.md#m-0004)

No dedicated Quest records existed for these small checkpoints. The Milestone
entries say so rather than inventing historical work records.

---

# 4. Acceptance Criteria

- [x] M-0001 through M-0004 exist in the canonical ledger.
- [x] Each Milestone uses an explicit stable anchor.
- [x] Each Milestone identifies the Awakening Epoch and Verified status.
- [x] Dates and evidence come from existing docs or local revision history.
- [x] Each Milestone records verified facts and relevant exclusions.
- [x] Missing dedicated Quest records are stated plainly.
- [x] Project Memory cites the ledger and retains its safety role.
- [x] The Roadmap cites the ledger and retains current direction.
- [x] Awakening readiness cites the ledger and retains readiness framing.
- [x] Full checkpoint summaries are no longer duplicated across those files.
- [x] No production code or packaging changed.

---

# 5. Outcome

The first four verified Awakening checkpoints now have one canonical home.
Project Memory, the Roadmap, and Awakening readiness cite stable M-ID anchors
instead of maintaining parallel evidence summaries.

This preserves the corrected EDC model: Milestones hold verified checkpoint
facts, the Roadmap holds direction, Project Memory holds durable safety memory,
and Quests record work.

---

# 6. Validation

Completed before commit:

- [x] `git diff --check` passed.
- [x] `make check` passed.
- [x] Changed Markdown links and explicit anchors were inspected.
- [x] Duplicate checkpoint headings and prose were searched.
- [x] `git diff --stat` was reviewed.
- [x] `git diff` was reviewed.

---

# Closing

One ledger holds the proof.

The rest of the map stays light enough to carry.
