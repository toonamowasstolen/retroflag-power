---
id: COC-001
version: 1.0.0
status: Accepted
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
related:
  - PROJECT_CHARTER.md
  - CONTRIBUTING.md
  - docs/05-development/ai-collaboration.md
last_updated: 2026-07-14
---

# Code of Conduct

Contributors to this project include Claude Code sessions working directly in the repo, not just
human engineers — this document holds both to the same standard, in every interaction it covers.

## What's expected

- Be direct and honest in every interaction touching this project — in commits, in quests, in issues
  and pull requests, and in what an AI session writes on the project's behalf.
- Disagree about architecture, hardware behavior, and design openly and often — disagree about people
  never.
- Assume good faith on a first pass. A confusing PR or a terse field-test report is much more often a
  busy person than a bad-faith one.
- Credit real work honestly. When a Claude Code session materially authors a change — which is most
  of this project's Quest history — the commit says so (`Co-Authored-By: Claude`, per
  `docs/05-development/COMMIT_STRATEGY.md`) and no session claims a decision it didn't actually make.
- Respect the hardware and the player behind it (`ENGINEERING_MANIFESTO.md` §17, §1) — a shortcut that
  looks fine in a demo but risks real data loss or a bricked device on someone's actual handheld isn't
  acceptable, ever, regardless of how it got there.

## What's not acceptable

- Harassment, personal attacks, or discriminatory language directed at anyone involved with this
  project, in any venue connected to it (issues, PRs, commit messages, any future community channel).
- Deliberately misrepresenting what was tested, verified, or shipped — this project's entire
  discipline rests on Milestones pairing every claim with real evidence
  (`docs/00-project/milestones.md`'s verification rules). A false claim here isn't just rude, it's a
  direct violation of how this project works — the same rule whether a human or an AI session made
  the claim.
- Retaliating against someone for reporting a concern in good faith.

## Enforcement

The project owner (`PROJECT_CHARTER.md`'s stakeholder model) decides violation outcomes. An AI
session working in this repo doesn't get a vote on whether its own conduct crossed a line — that call
belongs to the owner, same as any decision `docs/05-development/ai-collaboration.md` treats as
something AI must not do unilaterally.

**The escalation path:**

1. **Report it directly** — this is currently a small, owner-maintained project; there's no separate
   moderation body yet. If that changes, this section gets a real update, not a stale reference to a
   role that no longer exists.
2. **A substantiated violation gets a real consequence.** For a human contributor: a direct
   conversation for a first, minor issue; removal from the project for anything serious or repeated.
   For an AI session: the specific behavior gets corrected in this doc or `ai-collaboration.md` so it
   can't recur silently — a session doesn't get "fired," but a real mistake gets a real, permanent fix
   to the guidance that let it happen.
3. **Every report gets a response**, even if the response is "this doesn't rise to a violation, here's
   why" — silence is never an acceptable outcome.

This document is a durable artifact — a substantive change to it (not a typo fix) gets logged in
`PROJECT_MEMORY.md`, same as any other foundational doc.
