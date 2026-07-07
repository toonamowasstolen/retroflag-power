---
id: QUEST-0041
title: Plan the GPIO Read-Only Path
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Prepare the first real GPi Case input work with a safe read-only hardware plan and a small observer code boundary.
related:
  - docs/03-operations/gpio-read-only-plan.md
  - docs/02-hardware/gpi-case-2.md
  - docs/04-architecture/system-overview.md
  - internal/input
last_updated: 2026-07-07
---

# QUEST-0041 - Plan the GPIO Read-Only Path

> Mark the next hardware trail with a read-only lantern, so the first GPi Case
> input quest can observe the relic without casting a power spell.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Planning

## Outcome

- Added the GPIO read-only operations plan.
- Documented the current safe input loop from fake CLI observer through noop
  execution and event breadcrumbs.
- Recorded why the next GPi Case hardware step must observe only.
- Listed the hardware actions and lifecycle changes that remain out of bounds.
- Named the device facts still needed before a GPIO observer can be trusted.
- Added acceptance criteria for the first real hardware-read-only quest.
- Added `input.PowerButtonPressedEvent()` as the shared event constructor future
  GPIO observers should return from `Observer.NextEvent`.
- Kept the fake observer CLI path unchanged and backed by the same event shape.
- Linked the architecture overview to the read-only plan.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No real GPIO reads.
- No GPIO writes.
- No Raspberry Pi GPIO dependency.
- No shutdown command execution.
- No systemd service activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume behavior.
- No persistent state.
- No packaging changes.
- No debouncing implementation.
- No latching-switch semantics implementation.

## Milestone Note

The map now shows the safe route from fake observer charm to future read-only
hardware observation. The next quest can bring the GPi Case onto the workbench
with a clear rule: watch first, write nothing, keep every power action noop.
