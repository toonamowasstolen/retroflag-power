---
id: MILESTONES-001
title: Verified Milestone Ledger
version: 0.6.0
status: Active
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Provide the canonical ledger of numbered, verified RetroFlag Power checkpoints with stable anchors and evidence.
related:
  - PROJECT_MEMORY.md
  - docs/00-project/roadmap.md
  - docs/00-project/quests/
  - docs/00-project/quests/0006-gather-the-checkpoints-into-one-ledger.md
  - docs/00-project/quests/0007-add-the-link-lantern.md
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
  - docs/13-reference/terminology.md
supersedes:
  - docs/99-archive/project-milestones-pre-edc.md
last_updated: 2026-07-06
---

# Verified Milestone Ledger

> Epochs tell the journey. Milestones prove the distance traveled.

This is the canonical source for numbered, verified project checkpoints.

QUEST-0006 gathered the first four verified Awakening checkpoints into this
ledger without changing their historical evidence.

---

# Milestone Rules

## Identity

Milestones use stable, sequential IDs:

```text
M-0001
M-0002
M-0003
```

IDs remain stable after publication.

## Verification

A Milestone records a checkpoint only after evidence verifies it. Planning work
and implementation without verification remain in Quests, requirements, or the
Roadmap.

Each entry must include:

- stable anchor and ID
- title
- status
- verification date
- concise verified fact
- evidence
- related Quest or Quests
- related commit or revision when available
- related ADR when implementation reasoning needs durable explanation

## Stable anchors

Every entry uses an explicit anchor:

```html
<a id="m-NNNN"></a>
```

followed by:

```markdown
## M-0001 — Example Title
```

Other documents cite the anchored Milestone instead of copying its evidence or
completion prose.

## Milestone-scale ADRs

If a future contributor would reasonably ask why a Milestone was implemented
that way, create an ADR in `docs/adr/` and cross-link the ADR and Milestone.

---

# Milestone Entry Template

Do not copy this template into a real entry until its evidence is ready.

```markdown
<a id="m-NNNN"></a>
## M-NNNN — Verified Checkpoint Title

Status: Verified
Verified on: YYYY-MM-DD

Verified fact:

- concise statement

Evidence:

- command, test, artifact, or hardware result

Related:

- QUEST-NNNN
- docs/adr/NNNN-decision.md, when needed
- commit or revision
```

---

# Verified Milestones

<a id="m-0001"></a>
## M-0001 — Daemon Nameplate

Epoch: Awakening

Status: Verified

Verified on: 2026-07-03

### Summary

The daemon can identify itself consistently during local runs and standard
validation.

### Verified

- `retroflag-powerd --version` prints `retroflag-powerd 0.1.0-dev`.
- Version is `0.1.0-dev`.
- Startup logs include the daemon name and version.
- `make check` passed.

### Not included

- GPIO
- shutdown execution
- service activation
- resume
- state storage

### Evidence

- Revision `939c3a9` — Teach the daemon to say its name.
- Revision `81c3a3e` — Teach Workshop to ask the daemon its name.
- Revision `66a044c` — Teach Workshop check to verify the daemon name.

### Quest

No dedicated Quest record exists for this small checkpoint.

---

<a id="m-0002"></a>
## M-0002 — Config Satchel

Epoch: Awakening

Status: Verified

Verified on: 2026-07-03

### Summary

The daemon has a minimal internal configuration boundary with safe,
defaults-only identity and dry-run settings.

### Verified

- An internal config boundary exists.
- `AppName` defaults to `retroflag-powerd`.
- `Version` defaults to `0.1.0-dev`.
- `DryRun` defaults to `true`.
- The app receives config.
- Startup logs name, version, and `dry_run=true`.
- `make check` passed.

### Not included

- config file loading
- environment variable loading
- new CLI flags
- GPIO
- shutdown execution
- service activation
- resume
- state storage

### Evidence

- Revision `486c15d` — Give the daemon a config satchel.

### Quest

No dedicated Quest record exists for this small checkpoint.

---

<a id="m-0003"></a>
## M-0003 — Event Charms

Epoch: Awakening

Status: Verified

Verified on: 2026-07-03

### Summary

The daemon has a minimal internal model for describing lifecycle events.

### Verified

- `Event` has `Type` and `Message`.
- Lifecycle event types exist for daemon starting, daemon ready, shutdown signal
  received, and daemon stopped.
- The app logs lifecycle messages through the event model.
- `make check` passed.

### Not included

- event bus
- channels
- async processing
- persistence
- third-party dependencies
- GPIO
- shutdown execution
- service activation
- resume
- state storage

### Evidence

- Revision `ee463bb` — Give the daemon event charms.

### Quest

No dedicated Quest record exists for this small checkpoint.

---

<a id="m-0004"></a>
## M-0004 — Dry-Run Action Charm

Epoch: Awakening

Status: Verified

Verified on: 2026-07-03

### Summary

The daemon has a standalone dry-run action model that can describe a planned
no-operation action without executing it.

### Verified

- `Action` has `Type`, `Message`, and `DryRun`.
- `TypeNoop` exists.
- A dry-run noop helper exists.
- The noop action has `DryRun: true`.
- No execution path exists.
- No lifecycle wiring exists.
- `make check` passed.

### Not included

- command runner
- shell execution
- action queue
- channels
- async processing
- persistence
- packaging changes
- GPIO
- shutdown execution
- service activation
- resume
- state storage

### Evidence

- Revision `5ed148f` — Give the daemon a dry-run action charm.

### Quest

No dedicated Quest record exists for this small checkpoint.

---

<a id="m-0005"></a>
## M-0005 — Link Lantern

Epoch: Awakening

Status: Verified

Verified on: 2026-07-06

### Summary

The project has a local Markdown link checker exposed through
`make check-links`. It verifies internal Markdown links and explicit stable
anchors, including `m-0001` through `m-0004`, using Python 3 with no third-party
dependencies. The standard `make check` workflow remains unchanged.

### Verified

- `make check-links` passed.
- `make check` passed.
- `git status --short` was clean before the Milestone ledger update.

### Not included

- production Go changes
- packaging changes
- systemd activation
- GPIO
- shutdown execution
- `SafeShutdown.py` replacement
- `rc.local` edits
- resume
- state storage
- Node or npm dependencies

### Evidence

- Revision `a664052` — Add the Link Lantern.

### Quest

- [QUEST-0007 — Add the Link Lantern](quests/0007-add-the-link-lantern.md)

---

<a id="m-0006"></a>
## M-0006 — Status Badge

Epoch: Awakening

Status: Verified

Verified on: 2026-07-06

### Summary

The daemon has an `internal/status` package with a small `Status` model
containing `AppName`, `Version`, `DryRun`, and `State`. Its lifecycle states are
starting, ready, stopping, and stopped. A `New(config, state)` helper and focused
unit tests describe the model without adding runtime wiring or an external
status interface.

### Verified

- `gofmt` passed.
- `go test ./...` passed.
- `go build ./cmd/retroflag-powerd` passed.
- `retroflag-powerd --version` remained exactly
  `retroflag-powerd 0.1.0-dev`.
- `make check-links` passed.
- `make check` passed.

### Not included

- runtime wiring
- HTTP
- status server
- persistence
- channels
- async processing
- GPIO
- shutdown execution
- command runner
- shell execution
- packaging changes
- service activation
- `SafeShutdown.py` replacement
- `rc.local` edits
- resume
- state storage

### Evidence

- Revision `aa277a3` — Give the daemon a status badge.

### Quest

No dedicated Quest record exists for this small checkpoint.

---

<a id="m-0007"></a>
## M-0007 — Infrastructure Ledger and Regular-Project Conversion

Epoch: Awakening

Status: Verified

Verified on: 2026-07-14

### Summary

This repo converted from an external, read-only reference clone
(`_reference_retroflag-power`, as seen by the sibling EDC projects) into a regular, owned project —
it's the same owner's own repo, and the *origin* the shared EDC standard was retroactively
generalized from. Gained the same `docs/14-infrastructure/host-events.md` category (a numbered,
append-only ledger for host/infrastructure events outside this repo with a real effect on it) just
added to the template and the three sibling projects, adapted to this project's own voice.

### Verified

- `docs/14-infrastructure/host-events.md` exists.
- `git log` confirms Revision `03d2a61` landed and is pushed to `origin/main`.
- `find . -iname README.md` still reports the pre-existing two (root + `docs/02-hardware/`) —
  unrelated to this milestone, tracked separately.

### Not included

- Any real `INFRA-NNNN` entry — no host/infra event has affected this project yet.
- The `docs/02-hardware`/`docs/03-operations` reconciliation, `CODE_OF_CONDUCT.md`/`CONTRIBUTING.md`
  port, or any code-level change — separate work.

### Evidence

- Revision `03d2a61` — Open a filing slot for the world outside the repo.

### Quest

[0098-backfill-infrastructure-ledger-and-anchor-rule-sync](quests/0098-backfill-infrastructure-ledger-and-anchor-rule-sync.md)

---

<a id="m-0008"></a>
## M-0008 — Reconciled docs/02-hardware, docs/03-hardware, and docs/03-operations

Epoch: Awakening

Status: Verified

Verified on: 2026-07-14

### Summary

The governance doc's Section 11 described a `docs/03-software/` that was never created;
`docs/03-hardware/` (a second hardware folder) and `docs/03-operations/` both existed at number "03"
instead, undocumented — and `docs/03-hardware/README.md` was a real second `README.md`. Merged
`docs/03-hardware`'s 4 findings docs into the canonical `docs/02-hardware/` (via `git mv`, preserving
history), folded the README's index content into `gpi-case-2.md`'s new "See also" section, let
`docs/03-operations/` own "03" uncontested, and rewrote Section 11 to describe `docs/03-operations/`
accurately.

### Verified

- `find . -iname README.md` (excluding vendored dirs) now reports exactly one, at the repo root.
- `python3 scripts/check-markdown-links.py` — "Link Lantern checked 280 internal links across 150
  Markdown files," exit 0, run after every path change.
- `git status` shows the 4 files as renames, not delete+recreate — history preserved.

### Not included

- Any content rewrite inside the moved files themselves.
- `CODE_OF_CONDUCT.md`/`CONTRIBUTING.md` — separate quest.

### Evidence

- Working-tree change, committed this session (see quest for the exact file list).

### Quest

[0099-reconcile-hardware-and-operations-folders](quests/0099-reconcile-hardware-and-operations-folders.md)

---

<a id="m-0009"></a>
## M-0009 — Added CODE_OF_CONDUCT.md, expanded CONTRIBUTING.md

Epoch: Awakening

Status: Verified

Verified on: 2026-07-14

### Summary

This project had no `CODE_OF_CONDUCT.md` at all and a 3-line `CONTRIBUTING.md` stub — thinner than
its three sibling EDC projects. Ported the pattern over, adapted to this project's actual context
(solo/AI-collaborative hardware project, Quests/Milestones/ADRs, no branch/PR workflow evidenced in
`git log`) rather than copying the siblings' corporate framing verbatim.

### Verified

- `CODE_OF_CONDUCT.md` exists at repo root.
- `CONTRIBUTING.md` expanded, references only files that actually exist in this repo.
- `python3 scripts/check-markdown-links.py` — 0 broken links, exit 0.

### Not included

- A branch-naming/PR-review process — this project's real `git log` shows none, so none was invented.

### Evidence

- Working-tree change, committed this session (see quest for the exact file list).

### Quest

[0100-add-code-of-conduct-and-expand-contributing](quests/0100-add-code-of-conduct-and-expand-contributing.md)

---

<a id="m-0010"></a>
## M-0010 — Gave the daemon a startup banner, fixed the dead-code logger

Epoch: Awakening

Status: Verified

Verified on: 2026-07-14

### Summary

`internal/logging.New()` was dead code — never called anywhere; the real logger was built inline in
`main.go`, twice. Wired `logging.New` (now accepting an `io.Writer`) into both call sites, and added
a one-time ASCII startup banner (plain, no color) printed only on real daemon startup — moving the
"ASCII terminal welcome screen" aspiration from `PROJECT_MEMORY.md`'s idea list to real, at the
smallest scope that does so honestly.

### Verified

- `gofmt -l .` — no files.
- `go build ./...` — succeeds.
- `go test ./...` — all 12 tested packages pass (`internal/power` has no test files, pre-existing).
- A real compiled binary shows the banner on real startup; `--dry-run-power-button` and `--version`
  show no banner, output unchanged.
- No Go toolchain exists on Ramuh itself — verified via a throwaway Go 1.24 + Node container built on
  Phoenix, working tree synced over, checked there directly (not assumed).

### Not included

- Rewriting existing log-line formats.
- ANSI color support.

### Evidence

- Working-tree change, committed this session (see quest for the exact file list).

### Quest

[0101-give-the-daemon-a-startup-banner](quests/0101-give-the-daemon-a-startup-banner.md)

<a id="m-0011"></a>

## M-0011 — Sorted the old scrolls (closed a Draft quest parked since 2026-07-03)

Epoch: Awakening

Status: Verified

Verified on: 2026-07-16

### Summary

Checked each target file's real current state before touching anything, rather than trusting the
quest's 2026-07-03 description of them. Four of five were already fully done — `docs/roadmap.md`,
`docs/architecture.md`, and `docs/development.md` each already had metadata, `status: Superseded`,
a `superseded_by` pointer, and preserved original content; `docs/adr/0001-use-systemd.md` already
had metadata and an honest `Draft`/"Review Needed" treatment. Almost certainly handled incidentally
by later EDC-standardization quests without this one ever being explicitly closed. `README.md` had
one real gap — missing links to `docs/00-project/milestones.md` and
`docs/00-project/documentation-structure-and-governance.md`, both confirmed still-current — added
both.

### Verified

- Both new README link targets confirmed to exist on disk.
- `make check` not run — no Go toolchain on Ramuh, and the diff touches only two README link lines,
  no Go/script files; the throwaway-Phoenix-container route used for `M-0010`'s real code change
  would be disproportionate for a two-line prose addition.

### Not included

- Any code change — this was docs-only per the quest's own Rule 5.
- Archiving any of the four already-Superseded/Draft files further — they're already in the right
  state, moving them to `docs/99-archive/` wasn't warranted.

### Evidence

- Working-tree change, committed this session (see quest for the exact file list).

### Quest

[0002-sort-the-old-scrolls](quests/0002-sort-the-old-scrolls.md)
