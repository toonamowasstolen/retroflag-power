---
id: ADR-0001
title: Use systemd Instead of rc.local
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Preserve the early direction to use systemd rather than rc.local while the decision is reviewed and expanded into the current ADR format.
related:
  - docs/00-project/requirements.md
  - docs/04-architecture/system-overview.md
  - docs/10-decisions/adr-template.md
last_updated: 2026-07-03
---

# ADR-0001 — Use systemd Instead of rc.local

> Legacy decision note: archive candidate pending review against the current ADR
> template. This note records direction only; it does not authorize service
> activation or replacement of the existing shutdown path.

## Status

Draft

## Decision

Use systemd instead of `rc.local`.

## Review Needed

Before this ADR is accepted, document its context, alternatives, consequences,
validation plan, and safe migration boundary using the
[ADR template](../10-decisions/adr-template.md).
