---
id: INFRA-LOG-001
title: Host & Infrastructure Event Ledger
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Provide the canonical ledger of host/infrastructure events, outside this repo entirely, that had a real, concrete effect on this project's data, uptime, environment, or deployment.
related:
  - PROJECT_MEMORY.md
  - docs/00-project/milestones.md
  - docs/00-project/documentation-structure-and-governance.md
last_updated: 2026-07-14
---

# Host & Infrastructure Event Ledger

> Milestones prove the distance this project traveled on purpose. This ledger records what happened
> to the ground it stood on while traveling.

This is the canonical source for numbered host/infrastructure events that this project didn't cause
but had to react to — a reboot, a storage rebuild, a backup change, a network change, anything
outside this repo.

---

# Event Rules

## Identity

Events use stable, sequential IDs:

```text
INFRA-0001
INFRA-0002
```

IDs remain stable after publication.

## When to log an entry

Log an entry when host/infra work touched this project's data, uptime, or environment. Don't log
routine, no-impact maintenance elsewhere. Link out to the infra side's own full writeup instead of
duplicating it — record this project's own concrete impact here.

## Stable anchors

Every entry uses an explicit anchor:

```html
<a id="infra-NNNN"></a>
```

followed by:

```markdown
## INFRA-0001 — Example Title
```

for the same reason Milestone anchors exist — GitHub's auto-generated heading slug is unpredictable
on em-dash titles.

## Infra-triggered ADRs

If reacting to an entry produces a real decision (not just an operational fix), create an ADR in
`docs/adr/` and cross-link it — keep the event log and the decision record separate.

---

# Event Entry Template

Do not copy this template into a real entry until it describes something that actually happened.

```markdown
<a id="infra-NNNN"></a>
## INFRA-NNNN — Short Title

**Date:** YYYY-MM-DD
**Host/system:** which box or service this happened on

One paragraph: what happened at the host/infra layer.

**Impact on this project:**
- what broke, what was lost/recovered, downtime — be concrete

**Full detail:** link to the infra side's own full writeup, not duplicated here

**Follow-up:** link to any ADR/Quest raised in response, if any
```

No entries yet — no known host/infrastructure event has touched this project's data, uptime, or
environment as of 2026-07-14.
