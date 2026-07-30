---
id: TEMPLATE-META
title: Metadata Header Convention
version: 0.2.0
status: Accepted
owner: Joshua Taft
audience:
  - Everyone
purpose: Define the required frontmatter every project document must carry.
related:
  - docs/00-project/documentation-structure-and-governance.md
last_updated: 2026-07-23
---

# Metadata Header Convention

Every document in `docs/` and the Phase A root files carries this YAML
frontmatter as its first lines:

```yaml
---
id: WHY-001
title: Why This Project Exists
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Everyone
purpose: One sentence explaining the document's job.
related:
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
last_updated: 2026-07-03
---
```

## Field meanings

- **id** — stable identifier, `<CATEGORY>-<NNN>`. Category matches the doc's
  role (`WHY`, `CHARTER`, `MANIFEST`, `ADR`, `RFC`, `ARCH`, `REQ`, etc.).
  Never reused once assigned, even if the doc is archived — move it to
  `99-archive/` with its id intact rather than recycling the number.
- **title** — human-readable title, matching the main heading when
  practical. Required for major documents; strongly recommended for
  anything expected to survive more than one work session (every ADR file
  is named `ADR-0007.md`, so `title` is how a reader learns the subject
  without opening it).
- **version** — semantic version of the *document*, not the software.
  Bump patch for wording fixes, minor for added sections, major for a
  reversal of an earlier decision.
- **status** — one of `Draft`, `Review`, `Accepted`, `Active`,
  `Implemented`, `Verified`, `Superseded`, `Archived`, `Rejected`. Most
  docs move `Draft → Accepted`; RFCs/ADRs use
  `Draft → Review → Accepted/Rejected → Superseded`; quests/requirements
  use `Draft → Active → Implemented → Verified → Archived`. Full
  definitions of each value live in
  [`docs/00-project/documentation-structure-and-governance.md`](docs/00-project/documentation-structure-and-governance.md)
  section 4 — that document is canonical for status semantics; this file
  just needs to keep the same list.
- **owner** — the person accountable for the doc staying true, not
  necessarily who wrote it.
- **audience** — who this doc is written for (`Everyone`, `Contributors`,
  `AI Assistants`, `Operators`, etc.) — controls tone and assumed
  background knowledge.
- **purpose** — one sentence explaining the document's job, not a summary
  of all its content. Helps a reader (or an AI assistant) decide whether
  to open the file at all.
- **related** — other docs this one depends on, supersedes, or is
  superseded by, as paths **relative to this file's own directory** —
  the same convention as a normal Markdown link, not relative to the
  project root (e.g. a file in `docs/adr/` links a doc in
  `docs/00-project/` as `../00-project/MILESTONES.md`, not
  `docs/00-project/MILESTONES.md`). Must match what's actually linked in
  the document body — see the governance doc's "the metadata chain
  rabbit-holes, the body does not repeat it." This is how the doc graph
  stays navigable without a separate index.
- **last_updated** — ISO date, bumped every time content (not just the
  header) changes.

## Why this exists

See [Manifesto](ENGINEERING_MANIFESTO.md) principle 3 — "Documentation
Is Part of the Product." A header like this makes every doc individually
addressable (`id`), lets a reader judge trust level at a glance (`status`),
and keeps the doc graph traceable (`related`) without a hand-maintained
index that drifts out of sync.

## Staying in sync with the governance doc

This file is the quick, standalone reference for starting a brand-new
project from `_template/` before its own `docs/00-project/` tree exists
Once a project has copied in
`documentation-structure-and-governance.md`, that document is the
authoritative source for status values and field rules (per its own
"When any other doc in this tree... conflicts with this one, treat that
as drift to fix") — this file should be corrected to match it, not the
other way around, if the two ever disagree.
