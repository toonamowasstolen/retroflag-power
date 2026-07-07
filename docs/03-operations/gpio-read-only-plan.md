---
title: GPIO Read-Only Plan
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Plan the first safe RetroFlag Power step toward real GPi Case input without changing hardware state.
related:
  - docs/02-hardware/gpi-case-2.md
  - docs/04-architecture/system-overview.md
  - docs/00-project/quests/0041-plan-the-gpio-read-only-path.md
  - internal/input
last_updated: 2026-07-07
---

# GPIO Read-Only Plan

> The next hardware lantern should only watch the trail. It should not move a
> single stone.

## Current Safe Input Loop

RetroFlag Power already has a deterministic dry-run path for a power-button
observer event:

```text
CLI fake observer
  -> input observer
  -> power intent
  -> config policy
  -> deterministic plan
  -> noop execution
  -> event breadcrumbs
```

The command-line field kit is:

```sh
go run ./cmd/retroflag-powerd --fake-power-button-observer
```

That command emits one fake `power_button_pressed` input event, maps it to the
existing power-button intent, applies the current `power_button_action` policy,
builds a deterministic plan, executes only the noop action, and prints the
breadcrumb ledger.

The code boundary is `internal/input.Observer`. Future GPIO-backed observers
should emit the same project-level event as the fake observer by returning
`input.PowerButtonPressedEvent()` from `NextEvent`. Core app, planner, config,
and executor code should continue to receive input events and power intents,
not GPIO chips, line offsets, edge names, or electrical polarity.

## Why Read-Only Comes Next

The GPi Case path has real hardware unknowns. The first device quest should
prove that RetroFlag Power can observe candidate inputs safely before it tries
to affect power, boot, shutdown, services, or state.

Read-only observation lets the project gather evidence while preserving the
existing RetroFlag setup. It keeps the lantern bright enough to see real button
behavior and dim enough to avoid changing the device.

## Not Yet

The first hardware-readiness quest must not add or perform:

- GPIO writes.
- Shutdown command execution.
- systemd service activation.
- `rc.local` replacement.
- `SafeShutdown.py` replacement.
- Resume behavior.
- Persistent state.
- Raspberry Pi GPIO dependencies.
- Packaging changes.
- Debouncing behavior beyond recorded observations.
- Latching-switch semantics beyond recorded observations.

## Later GPi Case Test Approach

When hardware work begins, approach the GPi Case test as a supervised
observation session:

- Start from a known-good device with the existing RetroFlag behavior intact.
- Record the case model, Raspberry Pi model, OS image, kernel version, and any
  installed RetroFlag scripts or services.
- Inspect available GPIO chips and lines with read-only tools.
- Watch candidate lines one at a time while pressing or toggling only the
  intended case controls.
- Record every observation in the hardware ledger before turning it into code.
- Keep RetroFlag Power in a dry-run/noop mode while observations are collected.
- Stop if a candidate line appears to affect shutdown, reset, backlight, fan, or
  any other hardware output.

Useful manual probes are documented in the GPi Case 2 hardware notes, including
`gpiodetect`, `gpioinfo`, and `gpiomon` examples.

## Device Facts Still Needed

The project still needs facts from the device before a real observer can be
trusted:

- Candidate GPIO chips and line offsets for the power switch.
- Candidate GPIO chips and line offsets for the reset button, if included in
  the first input quest.
- Active-high or active-low behavior for each candidate input.
- Whether each control behaves like a latching switch or a momentary button.
- Bounce, duplicate edge, or noise observations during normal presses and
  toggles.
- Whether observed behavior changes across boot, shutdown, suspend-like states,
  or RetroFlag script activity.
- Whether any candidate input is already consumed by another service.

These facts should remain observations until repeated on the target device and
linked back to the hardware note.

## First Hardware Read-Only Quest Acceptance

The first real hardware-read-only quest is complete only when:

- The daemon can run a read-only observer mode without GPIO writes.
- No shutdown command is executed.
- No systemd service is installed, enabled, disabled, or restarted.
- Existing `rc.local` and `SafeShutdown.py` behavior is not replaced.
- No resume or persistent state behavior is added.
- No Raspberry Pi-specific GPIO dependency is introduced.
- Any future modern Linux GPIO dependency is chosen explicitly and remains
  read-only.
- The fake observer CLI path still passes its existing tests.
- The read-only observer emits the same `power_button_pressed` project event as
  `input.PowerButtonPressedEvent()`.
- Candidate pins, polarity, latching or momentary behavior, and bounce/noise
  notes are documented with device context.
- The resulting plan and execution remain deterministic and noop-only.
