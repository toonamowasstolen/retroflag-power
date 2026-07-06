---
id: MILESTONES-001
title: Verified Milestone Ledger
version: 0.5.0
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
