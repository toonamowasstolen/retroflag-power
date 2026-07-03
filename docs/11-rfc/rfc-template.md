---
id: RFC-TEMPLATE-001
title: Request for Comments Template
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Provide a lightweight template for proposing and discussing meaningful changes before they become accepted decisions or implementation work in RetroFlag Power.
related:
  - docs/00-project/requirements.md
  - docs/00-project/roadmap.md
  - docs/04-architecture/system-overview.md
  - docs/05-development/ai-collaboration.md
  - docs/10-decisions/adr-template.md
last_updated: 2026-07-03
---

# RFC-NNNN — Proposal Title

> A proposal is a bridge between idea and decision.

## Status

Draft

Possible values:

- Draft
- Open
- Accepted
- Rejected
- Superseded
- Implemented

## Date

YYYY-MM-DD

## Owner

Joshua Taft

## Related Requirements

- REQ-0000

## Related Documents

- docs/path/document.md

---

# 1. Summary

Provide a short summary of the proposal.

The summary should answer:

```
What is being proposed?
```

---

# 2. Motivation

Explain why this proposal exists.

Include:

- user problem
- maintainer problem
- hardware need
- safety concern
- product experience goal
- technical limitation
- opportunity for improvement

---

# 3. Goals

List what this proposal is trying to accomplish.

- goal

---

# 4. Non-Goals

List what this proposal is intentionally not trying to do.

- non-goal

Non-goals are important because they protect scope.

---

# 5. User Impact

Describe how this proposal affects users.

Consider:

- players
- power users
- developers
- maintainers
- hardware porters
- AI assistants

---

# 6. Proposed Approach

Describe the proposed approach.

Include diagrams, flow examples, or file layouts if helpful.

Example:

```
Power Switch event
      │
      ▼
Hardware Service
      │
      ▼
Event Bus
      │
      ▼
Power Service
```

---

# 7. Technical Details

Describe technical details that matter.

This may include:

- services affected
- package layout
- data structures
- commands
- configuration
- state files
- systemd behavior
- hardware behavior
- testing approach

Do not invent facts.

Mark assumptions clearly.

---

# 8. Alternatives Considered

## Alternative 1 — Name

Description:

Pros:

- item

Cons:

- item

Reason not chosen:

## Alternative 2 — Name

Description:

Pros:

- item

Cons:

- item

Reason not chosen:

---

# 9. Risks

List risks introduced by this proposal.

- risk

For safety-critical proposals, include shutdown, state, and rollback risks explicitly.

---

# 10. Mitigations

List how the project can reduce or manage the risks.

- mitigation

---

# 11. Open Questions

List questions that must be answered before acceptance or implementation.

- question

---

# 12. Acceptance Criteria

Describe what must be true for this proposal to be considered accepted or complete.

- criterion

---

# 13. Implementation Plan

Break the proposal into practical steps.

1. step
2. step
3. step

Keep the plan small enough to become commit packets.

---

# 14. Testing and Validation

Describe how the proposal will be tested or validated.

Possible validation types:

- unit tests
- mock hardware tests
- integration tests
- hardware validation
- boot timing measurement
- shutdown testing
- documentation review

---

# 15. Rollback Plan

Describe how the project can undo or recover from the change if it causes problems.

This is especially important for:

- power behavior
- shutdown behavior
- state writes
- resume behavior
- install scripts
- systemd units

---

# 16. Documentation Plan

List documents that must be updated if this proposal is accepted.

- docs/path/document.md

---

# 17. Decision Outcome

Complete this section when the proposal is accepted, rejected, or superseded.

If accepted, link to the ADR that records the final decision.

Example:

```
Accepted by ADR-0001.
```

---

# Closing

An RFC is not a delay tactic.

It is a safe place to think before building.

Use it when the question is bigger than a commit but not yet ready to become law.
