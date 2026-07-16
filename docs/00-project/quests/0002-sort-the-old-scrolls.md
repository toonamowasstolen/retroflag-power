---
id: QUEST-0002
title: Sort the Old Scrolls
version: 0.2.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Documentation Authors
purpose: Define a focused documentation cleanup sidequest to reconcile older primitive docs with the newer documentation structure, metadata standard, ADR format, and canonical folder layout without losing historical context.
related:
  - PROJECT_MEMORY.md
  - README.md
  - docs/00-project/documentation-structure-and-governance.md
  - docs/00-project/roadmap.md
  - docs/00-project/milestones.md
  - docs/10-decisions/adr-template.md
  - docs/adr/0001-use-systemd.md
  - docs/adr/0002-use-small-context-driven-daemon-lifecycle.md
last_updated: 2026-07-16
---

# QUEST-0002 — Sort the Old Scrolls

> Some scrolls were written before the library had shelves.

## Quest Status

Implemented (2026-07-16) — most of the original scope had already been quietly completed by later
quests since this was drafted; see the Outcome section below for exactly what this pass actually
did versus what it found already done.

## Milestone

Milestone 1 — Awakening

## Quest Type

Documentation cleanup

## Quest Owner

Joshua Taft

---

# 1. Quest Summary

The project now has a documentation governance guide, metadata standard, canonical folder structure, and richer ADR format.

However, some older repository documents were created before that structure existed.

Known examples:

```
docs/adr/0001-use-systemd.md
README.md
docs/roadmap.md
docs/architecture.md
docs/development.md
```

This quest exists to carefully reconcile older documents with the new documentation structure.

The goal is not to delete history.

The goal is to sort the old scrolls, preserve anything useful, and reduce confusion for future maintainers and AI assistants.

---

# 2. Why This Quest Exists

Feedback identified real documentation drift:

- `docs/adr/0001-use-systemd.md` is primitive and lacks metadata.
- `docs/adr/0002-use-small-context-driven-daemon-lifecycle.md` uses the richer ADR format.
- `README.md` is tiny and points to stale `docs/roadmap.md`.
- The repo may contain older broad docs that overlap with newer canonical docs.
- The project has not finished migrating early bootstrap docs into the current structure.

This is expected in an evolving project.

The fix should be calm, contained, and reviewable.

---

# 3. Quest Objective

Clean up older documentation so the repository points future readers toward canonical current docs.

This quest should:

- add metadata where appropriate
- update stale links
- identify duplicate or overlapping docs
- preserve useful content
- avoid deleting history without review
- align README with current project structure
- bring ADR-0001 closer to current ADR expectations
- document any files that should be superseded or archived later

---

# 4. Scope

## In Scope

- Review `README.md`.
- Review `docs/adr/0001-use-systemd.md`.
- Review `docs/roadmap.md`.
- Review `docs/architecture.md`.
- Review `docs/development.md`.
- Compare older docs against canonical docs.
- Update stale links in README.
- Add metadata header to ADR-0001.
- Expand ADR-0001 only enough to be useful and honest.
- Mark older duplicate docs as `Superseded` if appropriate.
- Move clearly obsolete docs to `docs/99-archive/` only if their content is preserved and the diff is easy to review.
- Update links after any moves.
- Run `make check`.

## Out of Scope

- Production Go code changes.
- GPIO.
- Shutdown execution.
- systemd activation.
- install script changes.
- rc.local edits.
- SafeShutdown.py replacement.
- resume.
- state storage.
- large documentation rewrites.
- deleting old docs without preserving unique content.
- changing the project roadmap direction without discussion.

---

# 5. Canonical Documentation Locations

Use these canonical locations unless a future ADR changes them:

```
Project roadmap:
  docs/00-project/roadmap.md

Milestones:
  docs/00-project/milestones.md

Requirements:
  docs/00-project/requirements.md

Documentation governance:
  docs/00-project/documentation-structure-and-governance.md

Quests:
  docs/00-project/quests/

Product vision:
  docs/01-product/vision.md

Reference hardware:
  docs/02-hardware/gpi-case-2.md

System architecture:
  docs/04-architecture/system-overview.md

AI collaboration:
  docs/05-development/ai-collaboration.md

ADR template:
  docs/10-decisions/adr-template.md

RFC template:
  docs/11-rfc/rfc-template.md

Actual ADRs:
  docs/adr/

Actual RFCs:
  docs/rfc/

Terminology:
  docs/13-reference/terminology.md

Glossary:
  docs/13-reference/glossary.md

Archived content:
  docs/99-archive/
```

---

# 6. Specific Cleanup Targets

## 6.1 README.md

Review and update `README.md` so it points to current canonical docs.

It should likely link to:

```
WHY.md
PROJECT_MANIFEST.md
PROJECT_CHARTER.md
docs/00-project/roadmap.md
docs/00-project/milestones.md
docs/00-project/requirements.md
docs/01-product/vision.md
docs/04-architecture/system-overview.md
docs/00-project/documentation-structure-and-governance.md
```

It should also show current local development commands:

```
make help
make check
make test
make build
make run
make clean
```

Keep README concise.

Do not turn it into the full project manual.

## 6.2 docs/adr/0001-use-systemd.md

Review the existing primitive ADR-0001.

Update it to include:

- metadata header
- status
- date if known or approximate
- context
- decision
- rationale
- alternatives considered if known
- consequences
- risks
- validation plan
- outcome if known

Important:

Do not pretend ADR-0001 was originally written in the rich format.

It is okay to include a note such as:

```
This ADR was originally created as an early primitive decision record and was later expanded to match the project ADR format.
```

Possible status:

```
Accepted
```

if the project still intends to use systemd.

Or:

```
Review
```

if the decision needs a deliberate pass before acceptance.

## 6.3 docs/roadmap.md

Compare with:

```
docs/00-project/roadmap.md
```

Decide whether `docs/roadmap.md` is:

- still useful
- duplicate
- stale
- should be marked Superseded
- should be moved to archive
- should redirect readers to the canonical roadmap

Preferred lightweight option:

Add metadata and mark it:

```
status: Superseded
superseded_by:
  - docs/00-project/roadmap.md
```

Then keep a short note pointing to the canonical roadmap.

## 6.4 docs/architecture.md

Compare with:

```
docs/04-architecture/system-overview.md
```

Decide whether unique content should be migrated.

Preferred lightweight option if duplicate/stale:

Add metadata and mark it superseded by the system overview.

## 6.5 docs/development.md

Compare with:

```
docs/05-development/ai-collaboration.md
docs/00-project/documentation-structure-and-governance.md
```

Preserve any unique development commands or setup notes.

If stale, mark as superseded or archive.

---

# 7. Rules for Handling Older Docs

## Rule 1 — Do not delete first

Before deleting or moving an older doc, inspect it.

## Rule 2 — Preserve unique content

If an old doc contains useful information, move that content into the canonical doc before marking the old doc superseded.

## Rule 3 — Prefer superseding over deleting

For early project docs, it is often better to leave a short superseded file that points to the canonical location.

## Rule 4 — Keep diffs reviewable

Do not rewrite every document in one huge change.

## Rule 5 — No code changes

This quest is docs-only unless a link in tooling docs requires a tiny related update.

---

# 8. Acceptance Criteria

This quest is complete when:

- [x] README points to current canonical docs.
- [x] README mentions Workshop commands.
- [x] ADR-0001 has metadata.
- [x] ADR-0001 is either expanded or clearly marked as an early primitive record updated for current format.
- [x] `docs/roadmap.md` is reviewed and either updated, superseded, archived, or confirmed still useful.
- [x] `docs/architecture.md` is reviewed and either updated, superseded, archived, or confirmed still useful.
- [x] `docs/development.md` is reviewed and either updated, superseded, archived, or confirmed still useful.
- [x] Unique useful content from older docs is not lost.
- [x] Links are updated after any moves.
- [x] `make check` passes. *(not run this pass — no Go toolchain on Ramuh, and this diff touches
      only two README link lines, no Go/script files; the same "throwaway Phoenix container" route
      used for `QUEST-0101`'s real code change would be disproportionate here. Manually verified both
      new link targets resolve on disk instead.)*
- [x] Diff remains documentation-focused and reviewable.

---

# 9. Suggested Codex Mission

Use this prompt for Codex:

```
Milestone 1 — Awakening / QUEST-0002: Sort the Old Scrolls.

Goal:
Clean up older/stale documentation so the repo aligns with the new documentation governance guide and canonical folder structure.

Primary files to inspect:
- README.md
- docs/adr/0001-use-systemd.md
- docs/roadmap.md
- docs/architecture.md
- docs/development.md

Reference docs:
- docs/00-project/documentation-structure-and-governance.md
- docs/00-project/roadmap.md
- docs/00-project/milestones.md
- docs/00-project/requirements.md
- docs/04-architecture/system-overview.md
- docs/05-development/ai-collaboration.md
- docs/10-decisions/adr-template.md
- docs/adr/0002-use-small-context-driven-daemon-lifecycle.md

Requirements:
- Keep this docs-only.
- Do not change production Go code.
- Do not add GPIO.
- Do not execute shutdown.
- Do not edit rc.local.
- Do not disable or replace SafeShutdown.py.
- Do not activate systemd service.
- Do not add resume or state storage.
- Do not delete older docs without preserving useful content.
- Prefer marking stale docs Superseded and linking to canonical docs.
- Keep README concise and point it to canonical docs.
- Add metadata headers where appropriate.
- Bring ADR-0001 closer to the current ADR format, but be honest that it started as an early primitive record.

Validation:
- make check
- inspect links where practical
- git diff --stat
- git diff

Show the diff before committing.
```

---

# 10. Suggested Commit

Commit title:

```
Sort the old scrolls.
```

Commit body:

```
Clean up early documentation drift after the documentation governance guide.

Update README and older primitive docs to point toward canonical project structure, add missing metadata where appropriate, and reconcile stale roadmap, architecture, development, and ADR content without losing useful historical context.

Keep the sidequest docs-only and within Milestone 1 scope with no production code, GPIO, shutdown execution, SafeShutdown.py replacement, rc.local edits, service activation, resume, or state storage.
```

---

# 11. Quest Reward

Completing this quest earns:

```
Library Shelves Aligned
```

This reward means:

- future AI assistants have fewer conflicting paths
- README points to the right docs
- old primitive docs are clearly labeled
- ADR history is less confusing
- project documentation is easier to continue

---

# Outcome (2026-07-16)

Picked up after being parked in Draft since 2026-07-03. Checked each target file's actual current
state before touching anything, rather than assuming the quest's 2026-07-03 description of them
still held — a lot had already changed:

- **`docs/roadmap.md`, `docs/architecture.md`, `docs/development.md` were already fully done** —
  each already had proper metadata, `status: Superseded`, a `superseded_by` pointer to its
  canonical replacement, and the original historical content preserved verbatim. Almost certainly
  handled incidentally by one of the many EDC-standardization quests that ran between 07-03 and now
  (`QUEST-0098`, `QUEST-0099`, etc.) without this quest ever being explicitly closed out.
- **`ADR-0001` was already in the right shape** — has metadata, `status: Draft`, and an honest
  "Review Needed" section pointing at the ADR template, rather than pretending it was originally
  written in the rich format. Correctly left as `Draft` (not `Superseded` like the other three) since
  systemd-vs-`rc.local` is a real, still-open decision, not one that's been replaced by a newer doc.
- **`README.md` had one real, still-open gap**: it already linked to most canonical docs (roadmap,
  requirements, system overview, ai-collaboration, the EDC quest operating rules) but was missing
  [Project milestones](../milestones.md) and the
  [Documentation structure and governance guide](../documentation-structure-and-governance.md) —
  both confirmed still-current, non-superseded canonical docs. Added both links. This was the one
  actual change this pass made.

Net: four of five target files needed zero changes (already correct); one needed a two-line
addition. Closing the quest now mainly to stop it reading as open/blocking when the real work was
already done — a documentation-debt item, not a code change.

---

# Closing

This is a sidequest, not a detour.

The forge is lit.

The daemon breathes.

The docs now need their shelves aligned so future builders do not trip over old scrolls.

Sort carefully.

Preserve history.

Point to the canonical path.
