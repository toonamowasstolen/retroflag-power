---
id: QUEST-0102
title: Join the family EDC standard — 879 lint findings to zero
version: 1.0.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record this project's first pass under the shared EDC lint tooling, the four root artifacts it was missing, and the path convention it had been following since before the standard existed.
related:
  - ../MILESTONES.md
  - ../../../PRODUCT_PHILOSOPHY.md
  - ../../../EPOCHS.md
  - ../../adr/0003-adopt-epoch-milestone-quest-model.md
last_updated: 2026-07-30
---

# QUEST-0102 — Join the family EDC standard

> This project wrote the standard the others follow, and was the last to be
> measured against it.

## Quest Status

Implemented. `edc-lint` passes at zero errors; merged to `main` 2026-07-30.

## Epoch

Awakening

## Quest Type

Documentation cleanup

## Quest Owner

Joshua Taft

The checkpoint this reports into is [`../MILESTONES.md`](../MILESTONES.md).

---

# 1. Quest Summary

The EDC lint scripts had never run here. First measurement was 879
findings; the count now is zero, and internal links checked rose from 294
to 1,081.

That number was never 879 pieces of writing. **522 were dead `related:`
paths written repository-root-relative** — the convention this project
adopted before the standard settled on file-relative, and the one that
`documentation-structure-and-governance.md` section 3.8 still wrongly
described family-wide until this pass corrected it. Each was converted by
exact resolution against the repository root rather than by matching
basenames, which matters here: this project has many files sharing a name,
and a guess would have quietly mis-pointed them.

Another 115 were promotable — bare paths sitting in readiness checklists,
and references written as a document's title rather than its filename. The
remaining **569 across 136 documents genuinely had no mention at all**, and
those got a `## Related work` section naming each document by role rather
than as a bare see-also list.

That last decision was deliberate and is worth recording. These lists
average 6.7 entries against roughly 3 elsewhere in the family, so pruning
them was on the table. They were checked for copy-paste boilerplate and
there is none — the cross-references are curated and specific to this
project's field work. Keeping every one of them, and paying for them with a
visible section, was chosen over quietly deciding which of someone else's
references no longer mattered.

# 2. The four missing root artifacts

Every sibling project carried these; this one did not, being the ancestor
the standard was generalized from rather than a project grown out of it.

- [`PRODUCT_PHILOSOPHY.md`](../../../PRODUCT_PHILOSOPHY.md) — written, not
  templated. Nothing in it is new; it is drawn out of what
  `PROJECT_MANIFEST.md` already said in Sections 2, 8 and 9 and put where a
  decision can be checked against it. The design values are ranked, because
  a ranking that never breaks a tie is decoration: the player's data
  survives, then the illusion holds, then a failure is understood
  immediately, then delight, then portability.
- [`EPOCHS.md`](../../../EPOCHS.md) — the family's seven-Epoch ladder, which
  is not an imposition:
  [`ADR-0003`](../../adr/0003-adopt-epoch-milestone-quest-model.md) already
  adopted exactly that ladder here after auditing an older eleven-stage
  journey. The file defines vocabulary only and says so, because that ADR
  gives the Roadmap ownership of current position.
- `LICENSE_POLICY.md` — records what was already true, MIT and public, and
  why Apache-2.0 was the plausible alternative that lost.
- `METADATA_HEADER.md` — distributed from the template.

# 3. Also fixed

- Seven canonical documents raised to UPPERCASE via `git mv`, matching every
  other project; 234 references updated, history following the files.
- `gpi-case-2-hardware-findings-kms-power-notes.md` predated the metadata
  standard and carried a schema of its own invention — `status:
  field-notes`, plus `created`/`project`/`device`/`tone` keys.
- Two `related:` entries held git commit SHAs rather than documents.
- The README said "Milestone 1 — Awakening", the noun ADR-0003 retired, and
  understated the daemon: it claimed no GPIO capability when
  `--probe-gpio-signal` had already shipped.

# 4. Acceptance Criteria

- [x] `node scripts/edc-lint.js .` reports zero errors. **Done.**
- [x] No file deleted; the seven renames tracked as renames. **Done.**
- [x] Root-artifact parity with the rest of the family. **Done.**

---

# Closing

Still outstanding, deliberately: `docs/03-operations/` collides on number
with every sibling's `docs/03-software/`, which makes "see `docs/03-`"
ambiguous forever. That is a directory rename touching many paths and it
deserves its own quest rather than being smuggled into a documentation
pass.
