---
id: CONTRIB-001
version: 1.0.0
status: Accepted
owner: Joshua Taft
audience:
  - Contributors
related:
  - docs/05-development/ai-collaboration.md
  - docs/00-project/milestones.md
  - PROJECT_CHARTER.md
  - CODE_OF_CONDUCT.md
last_updated: 2026-07-14
---

# Contributing

This project's contributors include Claude Code sessions working directly in this repo, not just a
human engineer — most of this codebase and nearly all of its documentation were written that way.
Every rule below applies the same way regardless of who or what is proposing the change.

## Before you start

Read [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) and
[`docs/05-development/ai-collaboration.md`](docs/05-development/ai-collaboration.md) — the latter is
the single, authoritative source for how an AI assistant should operate in this repo, and applies to
a human contributor's own judgment calls too where it describes project-wide discipline rather than
AI-specific behavior.

## How work is scoped

Bigger-than-one-commit, smaller-than-a-Milestone work gets a Quest
(`docs/00-project/quests/NNNN-slug.md`) — see `documentation-structure-and-governance.md` Section 8.
Not every change needs one; a routine fix or a single-file edit doesn't. When in doubt: if you'd want
to look back later at "why did we do this and what did we skip," make it a Quest.

## Verification, not "should work"

A Milestone only records a checkpoint after real evidence verifies it —
`docs/00-project/milestones.md`'s own rule. Whoever authors a change is responsible for actually
running it (`make check`, `make check-links`, a real command against real hardware where that
applies) and citing the real result, not asserting it should pass. This is the same discipline for a
human contributor and an AI session alike.

## Commit and tone conventions

Follow the existing commit history's own established voice — plain, direct, occasionally playful
titles (see recent `git log`), `Co-Authored-By: Claude` when a Claude Code session materially
authored the change.

## Proposing a significant decision

Write an ADR (`docs/adr/`, copy `docs/10-decisions/adr-template.md`) when a decision would be
expensive to reverse later, or when a future contributor would reasonably ask "why did we do it this
way" about a specific Milestone. For a larger, not-yet-decided proposal, use
`docs/11-rfc/rfc-template.md` instead — most decisions on this project go straight to an ADR once
implemented, rather than a formal RFC first; reach for RFC when a proposal is big or uncertain enough
that it should be written up before the work starts.

## Reporting a problem

Open an issue describing what happened, on real hardware where applicable — include the exact
`retroflag-powerd --version` output and, if it's a field/hardware issue, which `docs/03-operations/`
procedure (if any) was being followed when it occurred.
