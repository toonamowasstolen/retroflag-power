---
id: EPOCH-001
title: Epoch Vocabulary
version: 0.3.0
status: Accepted
owner: Joshua Taft
audience:
  - Everyone
purpose: Define the shared vocabulary for the 7-stage project journey (formerly called "milestones"), used across all projects instead of generic phase numbers.
related:
  - PROJECT_CHARTER.md
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
  - docs/00-project/ROADMAP.md
  - docs/00-project/MILESTONES.md
last_updated: 2026-07-30
---

# Epoch Vocabulary

> Renamed from "Milestone Vocabulary" 2026-07-03. The word "milestone"
> now means something more granular — a numbered, concrete checkpoint
> within the current epoch (see
> [docs/00-project/MILESTONES.md](docs/00-project/MILESTONES.md)) — after
> real usage on `control4-platform` showed a single project needs more
> than 7 waypoints, and conflating "which big-picture stage am I in" with
> "what specific thing did I just finish" made both harder to track
> honestly. This file now defines the big-picture stage only; it lives in
> `docs/00-project/ROADMAP.md` going forward, not a separate
> `MILESTONES.md`.

Epochs should have personality. A project's
[docs/00-project/ROADMAP.md](docs/00-project/ROADMAP.md)
should describe where the project is using these names, not "Phase 3" or
"Sprint 12." The name should evoke what the project actually feels like
at that stage.

This vocabulary originated on `retroflag-power` and is defined here by
**what the software can actually do**, not by which documentation phase
is finished. Documentation work (Phase A–D) mostly happens *during*
Dreaming, but the stages after that are about real, observable
capability — a project doesn't get to claim "Momentum" because its docs
are thorough; it gets there because the thing actually does something.

```
Dreaming
   ↓
Awakening
   ↓
Heartbeat
   ↓
Memory
   ↓
Momentum
   ↓
Adventure
   ↓
Launch
```

## Who owns what, in this project

This file defines the vocabulary. It does not track where the project
currently stands — per
[ADR-0003](docs/adr/0003-adopt-epoch-milestone-quest-model.md), which
adopted this exact seven-Epoch ladder after auditing an older eleven-stage
journey, [docs/00-project/ROADMAP.md](docs/00-project/ROADMAP.md) owns the
current Epoch and the direction of travel, and
[docs/00-project/MILESTONES.md](docs/00-project/MILESTONES.md) owns the
verified checkpoints inside it. Names here, position there. Anything you
read in this file that claims to know what stage the project is in has
drifted and should be deleted rather than corrected.

Stage names appearing in older documents — Power, Resume, Polish,
Expansion, Release — belong to that superseded eleven-stage journey. They
are history, not vocabulary; do not reintroduce them.

## What each epoch means

- **Dreaming** — No production code exists yet. The project is
  discovering its purpose, values, architecture, and identity. Primary
  artifacts: `WHY.md`, `PROJECT_MANIFEST.md`,
  [`PROJECT_CHARTER.md`](PROJECT_CHARTER.md),
  `PRODUCT_PHILOSOPHY.md`, requirements, terminology, and a first
  architecture direction (Phase A, and most of Phase B/C's docs). The
  spark is being protected and written down before it can be forgotten
  or diluted.

- **Awakening** — The first executable compiles and runs. Primary
  artifacts: a first skeleton (a daemon, a server, an app shell —
  whatever "running" means for this project), structured logging,
  graceful startup/shutdown handling. Still minimal behavior — this
  stage proves the project can take a breath, not that it does anything
  useful yet.

- **Heartbeat** — The skeleton becomes a real, supervised thing. Primary
  artifacts: process supervision (systemd/Docker/PM2/whatever fits),
  structured log visibility, health checks, restart behavior, first
  status/diagnostic output. It has a pulse you can actually check.

- **Memory** — The project starts remembering things *at runtime*, not
  just in its docs — session/state persistence, durable storage, restore
  flows. (This is also, not coincidentally, when `PROJECT_MEMORY.md`'s
  own graduation checklist and the first ADRs tend to mature — the
  project's documentation and its runtime both start "remembering" around
  the same time.)

- **Momentum** — Performance and real behavior become *measurable*. Boot/
  startup/response timing, benchmarks, resource usage — the write code →
  test it → commit it → document it loop is running for real and small
  victories accumulate (see [Manifesto](ENGINEERING_MANIFESTO.md)
  principle 6). This epoch can last a long time — most of a project's
  life is spent here, and it's typically where the most **milestones**
  (see `docs/00-project/MILESTONES.md`) accumulate, since "measurable
  real behavior" usually means many distinct features shipping and being
  verified one at a time.

- **Adventure** — The project is usable by others, even roughly — a
  pilot user, real hardware in someone else's hands, a real dataset, a
  real dealer. Surprises happen here that Dreaming and Awakening couldn't
  have predicted. Expect the architecture docs to get revised; that's
  the point of this stage.

- **Launch** — A stable public release. Public site goes live, hardware
  ships, package gets published, `CHANGELOG.md` gets its first tagged
  version entry, and there's a real upgrade/troubleshooting path for
  people who aren't the project's own maintainer.

## Epochs vs. milestones vs. quests

Three different granularities, each with its own file, none duplicating
another:

- **Epoch** (this file, distilled per-project in
  `docs/00-project/ROADMAP.md`) — the big-picture "what stage of life is
  this project in" journey above. A project spends a long time in one
  epoch.
- **Milestone** (`docs/00-project/MILESTONES.md`) — a numbered, concrete
  checkpoint within the *current* epoch — e.g. "ticket system built and
  verified," "password + 2FA built and verified." Several milestones are
  needed to cross from one epoch into the next. This is new content, not
  a renaming of the old per-epoch Theme/Purpose/Included/Excluded
  structure (that content moved to `ROADMAP.md`).
- **Quest** (`docs/00-project/quests/`) — the task-level build-and-test
  log for one milestone's individual pieces of work: what was built,
  what was verified, what was explicitly not included. Smaller and more
  frequent than a milestone.

Each real unit of finished work should be documented **exactly once**, at
whichever of these three granularities actually fits it, and cited (not
copy-pasted) from anywhere else that needs to reference it.

## Per-project overrides

A project may define its own, more specific version of these epochs in
its own `docs/00-project/ROADMAP.md` if the generic definitions above
don't fit its shape well (e.g. a project with no "daemon" concept at
all). Do this the same way `LICENSE_POLICY.md` works: copy the relevant
section into the project's own `ROADMAP.md` and edit it there — don't
silently reinterpret the shared file, and don't leave the override
undocumented. Note *why* the override exists at the top of the section.

## How to use this in a project

- `docs/00-project/ROADMAP.md`'s "Current project state" section states
  which epoch the project is currently in and what would need to be true
  to move to the next one; `docs/00-project/MILESTONES.md` tracks the
  specific numbered checkpoints toward that.
- Don't skip epochs to look further along than reality — a project that
  jumps straight to claiming "Momentum" without a Heartbeat is exactly
  the "planning stops producing artifacts" failure mode the Manifesto
  warns against, just in the other direction. If nothing has been
  deployed/run yet, the project is not past Heartbeat, no matter how much
  documentation exists.
