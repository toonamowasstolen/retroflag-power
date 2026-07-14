---
id: QUEST-0100
title: Add CODE_OF_CONDUCT.md and Expand CONTRIBUTING.md
version: 1.0.0
status: Verified
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Close a real governance-scaffolding gap found in the 2026-07-14 EDC gap analysis — this project had no CODE_OF_CONDUCT.md at all and a 3-line CONTRIBUTING.md stub, thinner than its three sibling EDC projects.
related:
  - ../milestones.md
  - ../../../CODE_OF_CONDUCT.md
  - ../../../CONTRIBUTING.md
last_updated: 2026-07-14
---

# QUEST-0100 - Add CODE_OF_CONDUCT.md and Expand CONTRIBUTING.md

## Quest Status

Verified

## Epoch

Awakening

## Quest Type

Documentation (governance scaffolding)

## Quest Owner

Joshua Taft

---

# 1. Quest Summary

The 2026-07-14 EDC gap analysis found this project had no `CODE_OF_CONDUCT.md` at all, and
`CONTRIBUTING.md` was a 3-line stub — notably thinner than `control4-platform`/`driverworks-tooling`/
`master-driver`, which each got real governance scaffolding in an earlier pass. Ported the pattern
over, adapted to this project's actual context — a solo/AI-collaborative hardware project, not a
client-facing business — rather than copying the siblings' corporate framing verbatim.

# 2. Quest Objective

Write a real `CODE_OF_CONDUCT.md` and expand `CONTRIBUTING.md`, grounded in what actually exists here
(Quests, Milestones, ADRs/RFCs, `ai-collaboration.md`) rather than referencing sibling-project files
(`GIT_WORKFLOW.md`, `ISSUE_TRIAGE.md`) that don't exist in this repo.

---

# 3. Scope

## In Scope

- `CODE_OF_CONDUCT.md` — what's expected/not acceptable, an enforcement path scoped to this project's
  actual size (owner-maintained, no separate moderation body yet — not a stale reference to a role
  that doesn't exist here).
- `CONTRIBUTING.md` — expanded from 3 lines to cover: before-you-start reading, how work is scoped
  (Quests), verification discipline (Milestones), commit/tone conventions, when to write an ADR vs.
  an RFC, and how to report a problem.

## Out of Scope

- Inventing a branch-naming or PR-review process this project doesn't actually have — checked `git
  log` first: all real commits land directly on `main`, no evidence of a feature-branch/PR workflow,
  so `CONTRIBUTING.md` doesn't claim one.
- The sibling projects' own `CODE_OF_CONDUCT.md`/`CONTRIBUTING.md` — already real and complete, not
  touched.

---

# 4. Acceptance Criteria

This quest is complete when:

- [x] `CODE_OF_CONDUCT.md` exists at repo root. **Done.**
- [x] `CONTRIBUTING.md` expanded beyond the 3-line stub, referencing only files that actually exist in
  this repo. **Done.**
- [x] `python3 scripts/check-markdown-links.py` reports 0 broken links. **Done.**
- [x] `docs/00-project/milestones.md` gets a new entry. **Done as M-0009.**

---

# 5. Suggested Commit

Commit title:

```
Add CODE_OF_CONDUCT.md, expand CONTRIBUTING.md
```

Commit body: closes a real governance-scaffolding gap from the 2026-07-14 gap analysis — adapted
from the sibling EDC projects' pattern but grounded in what actually exists here (Quests, Milestones,
ai-collaboration.md), not their corporate/client-facing framing or files this repo doesn't have.

---

# Closing

Landed 2026-07-14 (see `milestones.md#m-0009`). Verified via the real link checker, and via a direct
check of `git log` before claiming anything about this project's actual review/branch workflow.
