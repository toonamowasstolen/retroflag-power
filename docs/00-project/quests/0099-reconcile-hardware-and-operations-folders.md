---
id: QUEST-0099
title: Reconcile docs/02-hardware, docs/03-hardware, and docs/03-operations
version: 1.0.0
status: Verified
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Fix a real number collision (docs/03-hardware and docs/03-operations both existed) and a real one-README violation (docs/03-hardware/README.md), found in the 2026-07-14 EDC gap analysis.
related:
  - ../milestones.md
  - ../documentation-structure-and-governance.md
  - ../../02-hardware/gpi-case-2.md
last_updated: 2026-07-14
---

# QUEST-0099 - Reconcile docs/02-hardware, docs/03-hardware, and docs/03-operations

> Two shelves shared one number. Now only one does, and it's labeled.

## Quest Status

Verified

## Epoch

Awakening

## Quest Type

Documentation (structural reconciliation)

## Quest Owner

Joshua Taft

---

# 1. Quest Summary

The 2026-07-14 EDC gap analysis found the on-disk `docs/` structure had drifted from the governance
doc's own numbered taxonomy: Section 11 described a `docs/03-software/` that was never created;
`docs/03-hardware/` (a second hardware folder, alongside the canonical `docs/02-hardware/`) and
`docs/03-operations/` both existed at number "03" instead, undocumented. `docs/03-hardware/README.md`
was also a second `README.md` — a direct violation of this project's own one-README rule.

# 2. Quest Objective

Merge `docs/03-hardware`'s 4 real findings docs into the canonical `docs/02-hardware/`, remove the
duplicate README (folding its index content into `gpi-case-2.md` instead), let `docs/03-operations`
own "03" uncontested, and update the governance doc's Section 11 to describe what's actually there.

---

# 3. Scope

## In Scope

- Moved (`git mv`, preserving history) all 4 `docs/03-hardware/*.md` findings docs into
  `docs/02-hardware/`.
- Removed `docs/03-hardware/README.md`; its index content now lives as a "See also" section in
  `docs/02-hardware/gpi-case-2.md`, which also gained `related:` entries for the 4 relocated docs.
- Updated every real link across the repo (`milestones.md`, several `docs/00-project/quests/*.md`,
  several `docs/03-operations/*.md`) from `docs/03-hardware/...` to `docs/02-hardware/...`. One
  historical quest (`0046`) kept its narrative prose describing what was true at the time, but its
  *link targets* were updated so they keep resolving (a link pointing nowhere serves no one, even in
  a historical record) — its README-specific link now points to `gpi-case-2.md`, where that content
  actually lives now.
- Rewrote governance-doc Section 11 (previously `docs/03-software/`, never created) to describe
  `docs/03-operations/` — what it's actually for, based on its real 25-item contents (field
  procedures, evidence ledgers, lantern designs, coverage/behavior maps) — rather than leave the
  taxonomy pointing at a folder that doesn't exist.
- Added a `docs/03-operations/` row to Section 34's canonical-locations list.

## Out of Scope

- Renumbering `docs/03-operations/` itself — it needed no file moves once hardware vacated "03".
- Any content rewrite inside the moved/existing files themselves — this was a structural move, not
  an editorial pass.
- `CODE_OF_CONDUCT.md`/`CONTRIBUTING.md` — separate quest (`0100`).

---

# 4. Acceptance Criteria

This quest is complete when:

- [x] `find . -iname README.md` (excluding vendored dirs) reports exactly one, at the repo root.
  **Done.**
- [x] `docs/03-hardware/` no longer exists; its real content lives in `docs/02-hardware/`. **Done.**
- [x] `python3 scripts/check-markdown-links.py` reports 0 broken links, run *after* every path
  change. **Done** — "Link Lantern checked 280 internal links across 150 Markdown files," exit 0.
- [x] Governance-doc Section 11 describes `docs/03-operations/` accurately. **Done.**
- [x] `docs/00-project/milestones.md` gets a new entry. **Done as M-0008.**

---

# 5. Suggested Commit

Commit title:

```
Reconcile docs/02-hardware, docs/03-hardware, and docs/03-operations
```

Commit body: merges docs/03-hardware's 4 findings docs into the canonical docs/02-hardware, removes
the duplicate README (folded into gpi-case-2.md's new "See also" section), lets docs/03-operations
own "03" uncontested, and rewrites the governance doc's Section 11 to describe what's actually there
instead of a docs/03-software/ that was never created. Verified with the real link checker before
and after, not by eye.

---

# Closing

Landed 2026-07-14 (see `milestones.md#m-0008`). Confirmed via `git status` (renames tracked, not
delete+recreate), a full link-checker pass (0 broken), and a direct `find` for stray READMEs.
