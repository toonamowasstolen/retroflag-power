---
id: ARCHIVE-EXTERNAL-FINDINGS-001
title: External Findings on This Project's Docs (2026-07-03)
version: 1.0.0
status: Archived
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Preserve the 2026-07-03 findings written about this project from outside it, and record which of their recommendations were implemented.
related:
  - ../00-project/MILESTONES.md
  - ../00-project/quests/0006-gather-the-checkpoints-into-one-ledger.md
  - ../adr/0003-adopt-epoch-milestone-quest-model.md
  - ../../PROJECT_MEMORY.md
last_updated: 2026-07-30
---

# External Findings on This Project's Docs (2026-07-03)

> **Archived 2026-07-30.** This document was written on 2026-07-03 by the
> `control4-platform` / `driverworks-tooling` / `master-driver` side, at a
> time when this project was believed to be an **external clone** — it was
> filed as `_reference_retroflag-power` and the findings were deliberately
> kept as a side-channel note rather than raised here, "since that project
> isn't ours to edit directly" (`master-driver`'s `QUEST-0004` carries the
> same wrong premise, corrected in place there on 2026-07-30 — named rather
> than linked, because a link out of this repository breaks on clone).
>
> The premise was wrong. On 2026-07-14 this turned out to be the same
> operator's own project. It became a full family member with its own EDC,
> and the note it had been generating went on sitting unversioned at the
> share root until today.
>
> **Its central recommendation was implemented.** Section 1 reported the
> four Awakening checkpoints duplicated verbatim across `ROADMAP.md`,
> `awakening-readiness.md` and [`PROJECT_MEMORY.md`](../../PROJECT_MEMORY.md),
> with nothing keeping the copies in sync. They now live once in
> [`MILESTONES.md`](../00-project/MILESTONES.md) as `M-0001`–`M-0004`, and
> all three former copies carry links instead —
> [`QUEST-0006`](../00-project/quests/0006-gather-the-checkpoints-into-one-ledger.md),
> status `Implemented`. The Epoch/Milestone/Quest model the rest of the
> document argues for was adopted here by
> [`ADR-0003`](../adr/0003-adopt-epoch-milestone-quest-model.md).
>
> The document's own header claimed the findings were being folded into
> `claude-tools/docs/11-retroflag-power-gap-analysis.md`. **That file does
> not exist and never did on this host.** The folding happened through the
> quest and ADR above instead, which is why this is archived as history
> rather than carried as open work.
>
> Everything below is the original text, unedited apart from its
> non-standard frontmatter being replaced by the block above.

---

# Findings for retroflag-power

We're on the same EDC documentation standard (originated on this
project, adopted and extended on `control4-platform`/`driverworks-tooling`).
While pulling your repo for reference during our own documentation work,
we found a real structural bug already live in your docs, used it (plus
our own experience) to redesign a genuine gap in the shared standard,
and have now fully migrated both of our live projects to the corrected
model. This document reports what we found and what we changed — not a
proposal, a completed and independently-verified change on our end that
you may want to apply to yours.

## 1. Bug found: the same four checkpoints are duplicated verbatim in three files

As of commit `183353c`, the "Nameplate," "Config Satchel," "Event
Charms," and "Dry-Run Action Charm" checkpoints appear **word-for-word
identical** (each reworded slightly, but describing the exact same facts
and verify evidence — `make check` passed, `--version` output unchanged,
etc.) in:

- `docs/00-project/roadmap.md` (under "Current Awakening Checkpoint")
- `docs/00-project/awakening-readiness.md` (as four `## ... Checkpoint`
  sections)
- `PROJECT_MEMORY.md` (under "Milestone 1 — Awakening," as `#### ...
  checkpoint` sub-sections)

Right now this is cosmetically harmless because all three copies happen
to agree. But nothing keeps them in sync — the first time one gets
updated (a correction, a caveat added, a status change) without the
other two, the docs will silently contradict each other, and a future
reader (human or AI) has no way to know which copy is authoritative.

## 2. Root cause: "record a small, verified increment" has no single home

Your structure has three files that can each plausibly claim ownership
of this job: `roadmap.md`, `awakening-readiness.md`, and
`PROJECT_MEMORY.md`. The governance doc's Section 27 explains what
Milestones/Roadmap/Requirements/Quests each answer, but in your current
vocabulary "milestone" still means the whole 7-stage journey
(Dreaming→Launch) — there's no granularity *below* that for a checkpoint
the size of "the daemon now prints its version" to live in without
picking one of three overlapping homes.

## 3. What we changed on our side (2026-07-03, shipped and verified)

We hit the identical gap on `control4-platform`, whose `STATUS.md` had
independently accumulated 13 dated build-log sections (bigger than your
Nameplate/Config-Satchel scale, but the same underlying problem). We
split the vocabulary into four layers, each answering one question, each
cited from exactly one canonical place:

- **Epoch** — renamed from what your repo still calls "Milestone":
  Dreaming → Awakening → Heartbeat → Memory → Momentum → Adventure →
  Launch. The big-picture "what stage of life is the project in." Lives
  in a root file `EPOCHS.md` (was `MILESTONES.md`), distilled per-project
  in `docs/00-project/ROADMAP.md`, which now holds the full per-epoch
  Theme/Purpose/Included/Excluded/Exit-criteria detail (this content used
  to live in a separate `MILESTONES.md`; that file's *old* job moved into
  `ROADMAP.md` entirely).
- **Milestone** — brand new concept, not what the word meant in our docs
  a day earlier. A numbered, concrete, verified checkpoint within the
  *current* epoch: `M-0001`, `M-0002`, ... sequential across the whole
  project's life, never reset per epoch. Lives in
  `docs/00-project/MILESTONES.md` (repurposed — this filename now means
  something different than it used to for us too). Each entry: **Epoch**
  / **Status** (Built/Verified) / **Date** / one paragraph of what
  happened / **Verified:** real commands and their real output /
  **Not included in this milestone:** explicit exclusions / **Quest:**
  link to the work record, if one exists.
- **Quest** (your concept, unchanged) — the task-level build-and-test
  *work record* for one milestone's actual work: scope, in/out, real
  acceptance criteria, suggested commit, Outcome. Not every milestone
  needs a quest file — small, single-commit work is fully covered by its
  milestone entry alone.
- **`PROJECT_MEMORY.md`** — unchanged role on our end too: origin
  concept/brainstorm capture and a "Things that must not be lost" safety
  net, never a progress log.

We retired our `STATUS.md` files entirely (moved to `_trash/`, not
deleted) — their whole job (fact-checked "what's true right now," one
verify command per claim) is now covered with zero functional loss by
`ROADMAP.md`'s "Current project state" section plus `MILESTONES.md`'s
checkpoint log.

### Concretely, for your four duplicated checkpoints

Each one is written **exactly once** now — as a milestone entry — and
cited by number from wherever else needs to reference it, instead of
being re-narrated in full. If you adopt this: your Nameplate/Config-
Satchel/Event-Charms/Dry-Run-Action checkpoints become `M-0001` through
`M-0004` in a new `docs/00-project/milestones.md`, and
`roadmap.md`/`awakening-readiness.md`/`PROJECT_MEMORY.md` each shrink to
a one-line citation ("see M-0001 through M-0004") instead of three full
copies of the same prose.

## 4. Two structural moves worth calling out (in case you adopt the ADR/RFC split next)

Separately from the epoch/milestone split, we also finished migrating to
your repo's own real structure for ADRs/RFCs: real, numbered decision
records live in `docs/adr/` and `docs/rfc/`; `docs/10-decisions/` and
`docs/11-rfc/` hold only their template file (`ADR-TEMPLATE.md`/
`RFC-TEMPLATE.md`), nothing else. This matches what your repo already
does (`docs/adr/0001-use-systemd.md` etc.) — we'd drifted from it
earlier and corrected it as part of this same pass. No action needed on
your side; just confirming we're now consistent with your actual
practice, not just your governance doc's prose.

## 5. Three gaps we found in the governance doc itself, now fixed

While auditing our own migration for completeness (not prompted by
anything in your repo — this is us checking our own copy of the shared
doc), we found three real gaps in
`docs/00-project/documentation-structure-and-governance.md` that your
copy would likely hit too, once you have milestones/epoch-readiness
docs of your own:

1. **No rule for milestone-scale ADRs.** The doc only said "write an ADR
   when an *epoch's* foundational decision is made" — but in practice
   (our `control4-platform` ADR-0003, written for a single milestone deep
   into the Momentum epoch, not for the epoch itself), real ADR-worthy
   decisions get made at milestone scale too, and the doc gave no
   guidance on when that should happen. Fixed: added a rule — "if a
   future contributor would reasonably ask 'why did we do it this way'
   about *this milestone specifically*, not just the epoch it belongs
   to, it needs its own ADR," with the milestone and ADR cross-citing
   each other.
2. **Epoch-readiness docs never cited milestones as evidence.** The
   `EPOCH_READINESS.md` template's exit-criteria checklist had no field
   for "which milestone already satisfies this criterion" — so even
   though `MILESTONES.md` might have real evidence directly relevant to
   a readiness verdict, nothing pointed a reader to it. Fixed: the
   template now asks, for each exit criterion, to cite the specific
   `M-000N` entry that satisfies it, or say plainly that none does yet.
3. **`EPOCH_READINESS_*.md` was missing from the doc-type status table**
   (Section 5) that every other durable file type appears in. Fixed:
   added it (`Draft -> Accepted` once all exit criteria are actually
   met).

All three are now fixed in our shared template and backported to both
live projects, independently verified with a fresh link-checker (zero
broken links across all three) after each change.

## 6. Update (2026-07-04): a real anchor-link bug in the corrected model itself, now fixed

Section 3 above described the numbered `M-0001`, `M-0002`, ... milestone
entries we moved to. Once we'd actually accumulated 16+ of them in
`control4-platform` (and one each in two smaller sibling projects), we
found a genuine bug in the pattern itself — worth flagging since it will
bite you the moment you adopt the numbered-milestone model, not before.

**The bug:** every internal link we wrote used the short, obvious form —
`docs/00-project/MILESTONES.md#m-0001` — assuming that's what a heading
like `## M-0001 — Initial skeleton: Next.js + Prisma + local boot` would
auto-generate as its anchor. It isn't. We checked with the real
`github-slugger` npm package (the same slugger GitHub itself uses, not a
hand-rolled guess) and confirmed GitHub's actual generated anchor for
that heading is
`#m-0001--initial-skeleton-nextjs--prisma--local-boot` — note "nextjs"
with no hyphen, and doubled hyphens around the em dash. Any heading with
an em dash, colon, or period in its title produces a long, genuinely
non-obvious slug, not the short form every one of our links assumed.
This wasn't a typo or a one-off — it was systemic across every
milestone-link we'd written (~58 broken links total across our three
live projects by the time we found and fixed it).

**The fix we chose, and why:** rather than rewrite every link to the
long, fragile, real slug (which would also silently re-break the moment
any milestone's title text ever changed, since the slug is derived from
it), we added an explicit stable anchor immediately above each heading:

```
<a id="m-0001"></a>
## M-0001 — Initial skeleton: Next.js + Prisma + local boot
```

GitHub's Markdown renderer honors an explicit `<a id="...">` as a real
anchor target, so `#m-0001` now resolves correctly regardless of what
the heading's own auto-slug would otherwise be — and it survives title
edits forever, since `M-0001` (not the full title) is the part meant to
be a stable, permanent citation in the first place. We updated our
shared `_template/docs/00-project/MILESTONES.md`'s entry template to
require this anchor on every milestone entry going forward, and added an
explicit rule + the reasoning above to
`documentation-structure-and-governance.md` Section 8, so this doesn't
have to be independently rediscovered by whoever adopts the numbered-
milestone model next.

**Also shipped:** a small, durable link-checker
(`_template/scripts/check-links.js`, requires `npm install
github-slugger` once) that verifies every internal Markdown link —
including anchors — using the real slugger, not an approximation. Copied
into all four of our repos (`control4-platform`, `driverworks-tooling`,
`master-driver`, and `_template` itself). If useful, this is a
zero-dependency-beyond-`github-slugger` script you could drop into your
own repo's `scripts/` folder as-is; it doesn't assume anything specific
to our project structure beyond the shared `MILESTONES.md`/EDC
conventions we already share.

**Not yet done on our side either:** none of our four repos has CI set
up yet, so this checker runs manually, not automatically on every
commit/PR — worth flagging as a shared gap, not just yours to close.

## 7. Update (2026-07-05): tone/governance gaps found while applying your own Sections 8/9 more broadly

We spent this pass correcting a real mistake in one of our own projects
(`master-driver`): its docs described warmth as a documentation-only
concern, confined away from the driver's actual runtime log voice, on
the theory that a background/invisible service earns a clinical tone.
That's not what your own `PROJECT_MANIFEST.md` actually says — Section 8
("Terminal Personality") wants warmth in the runtime output itself
(logs, CLI, terminal), narrowed only by Section 9 ("Clear errors beat
clever messages") for genuine failure messages specifically. We'd
generalized your narrow failure-message exception into a blanket "stay
cold by default" rule, backwards from what you actually wrote. Fixed
across our side; while doing that pass, we found two things in your own
repo worth reporting back, in the same spirit as Sections 1-6 above —
not proposals, just what we noticed using your repo as reference.

### 7a. Your own runtime code doesn't yet do what your own docs describe

`internal/logging/logging.go` — the actual logger your running
`retroflag-powerd` daemon uses — is a completely generic, unpersonalized
`log.New(os.Stderr, "", log.LstdFlags)`. No warmth, no ASCII, nothing
that reflects Section 8's "helpful, warm, clear, occasionally playful"
intent or the "ASCII terminal welcome screen" your own `PROJECT_MEMORY.md`
lists as a real aspiration. This isn't a criticism — early-stage
projects reasonably build the daemon skeleton before its voice — but
it's worth naming precisely since it's the same category of gap
(docs promise a personality the running code hasn't caught up to yet)
we just found and fixed in `master-driver`. If useful: the smallest,
most direct fix would be a one-time ASCII banner printed at
`retroflag-powerd`'s boot, in whichever of `cmd/retroflag-powerd/
main.go` or `internal/app/app.go` already owns startup sequencing —
moving that aspiration from documented to real without touching every
log call site's format, which is a bigger, separate change.

### 7b. Your governance scaffolding is thinner than ours was before this pass

We noticed, only because we were adding this exact kind of doc to our
own three projects: your repo has no `CODE_OF_CONDUCT.md` at all, and
your `CONTRIBUTING.md` is a 3-line stub ("Please open issues and pull
requests"). That's notable mostly because your repo is otherwise the
most mature of the four we compared (real Go code, real tests, real CI
concepts sketched in your governance doc's own "TODO Tracking Rules"
section) — the governance layer is the one place our (much younger)
projects had more scaffolding than yours, before this pass. Not
something we're fixing in your repo, obviously; just a real, specific
gap worth naming the same way we'd want a peer project to name one of
ours.

### What we did about both, on our own side

Rewrote all three of our projects' `CODE_OF_CONDUCT.md`/`CONTRIBUTING.md`/
`AI_GUIDANCE.md` to (a) actually name the AI session as a current,
accountable contributor rather than a hypothetical future one, (b) speak
as "the team"/"Taft Consulting" in body text rather than naming an
individual by name (only the `owner:` metadata field keeps a name), and
(c) state every rule in a way that doesn't depend on current headcount —
no "today, solo... once a second person exists" scaffolding that would
go stale and need rewriting the moment a real hire happened. Added a new
`docs/05-development/ISSUE_TRIAGE.md` to each project. Full record in
each project's own `PROJECT_MEMORY.md` and its QUEST-0018 (control4-
platform) / QUEST-0004 (driverworks-tooling, master-driver).

## Offer

Happy to share our actual `_template/EPOCHS.md`,
`_template/docs/00-project/MILESTONES.md`,
`_template/docs/00-project/ROADMAP.md`,
`_template/docs/00-project/EPOCH_READINESS.md`, and (as of the
2026-07-04 update above) `_template/scripts/check-links.js` directly if
useful as a starting point for your own migration — they're generalized
from patterns your repo originated in the first place (the governance
doc itself, the Fact/Assumption/Decision/Research/Aspiration taxonomy,
the runtime-capability epoch definitions), so this is very much a
"here's what came back around, improved," not us inventing something
unrelated to hand you.
