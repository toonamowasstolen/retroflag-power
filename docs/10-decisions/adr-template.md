---
id: ADR-TEMPLATE-001
title: Architecture Decision Record Template
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Provide a lightweight template for recording important architecture and project decisions in RetroFlag Power.
related:
  - docs/00-project/requirements.md
  - docs/00-project/roadmap.md
  - docs/04-architecture/system-overview.md
  - docs/05-development/ai-collaboration.md
last_updated: 2026-07-03
---

# ADR-NNNN — Decision Title

> Decisions become memory when they are written down.

## Status

Draft

Possible values:

- Draft
- Proposed
- Accepted
- Superseded
- Rejected
- Deprecated

## Date

YYYY-MM-DD

## Owner

Joshua Taft

## Related Requirements

- REQ-0000

## Related Documents

- docs/path/document.md

---

# 1. Context

Describe the situation that led to this decision.

Include:

- what problem exists
- what constraints matter
- what requirements apply
- what risks are involved
- what is known
- what remains uncertain

Keep facts, assumptions, and research separate where practical.

---

# 2. Decision

State the decision clearly.

Example:

```
RetroFlag Power will begin as a modular monolith daemon supervised by systemd.
```

The decision should be specific enough that future contributors can tell whether later work follows it.

---

# 3. Rationale

Explain why this decision was made.

Include the reasoning that matters most.

Good rationale should help a future maintainer understand why this path made sense at the time.

---

# 4. Alternatives Considered

## Alternative 1 — Name

Description:

Pros:

- item

Cons:

- item

Reason rejected or deferred:

## Alternative 2 — Name

Description:

Pros:

- item

Cons:

- item

Reason rejected or deferred:

---

# 5. Consequences

## Positive Consequences

- item

## Negative Consequences

- item

## Neutral Consequences

- item

---

# 6. Risks

List any risks introduced or affected by this decision.

Example:

- This decision may make early implementation easier but require future refactoring if hardware support expands.

---

# 7. Validation Plan

Describe how the project will know the decision is working.

Examples:

- tests
- hardware validation
- boot measurements
- successful install
- successful shutdown behavior
- contributor feedback

---

# 8. Rollback or Revision Plan

Describe how the project could reverse, revise, or supersede this decision if needed.

Not every decision needs an easy rollback, but safety-critical decisions should have one.

---

# 9. Notes

Add any additional notes, links, commands, or context.

---

# 10. Outcome

Complete this section after the decision has been implemented or tested.

Possible prompts:

- Was the decision successful?
- Did the project learn anything unexpected?
- Should this ADR be updated, superseded, or kept as-is?

---

# Closing

An ADR is not bureaucracy.

It is a memory crystal.

Use it when future maintainers will ask:

```
Why did we do it this way?
```
