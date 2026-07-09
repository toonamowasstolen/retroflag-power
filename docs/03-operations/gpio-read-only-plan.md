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
  - docs/00-project/quests/0042-separate-raw-signals-from-interpreted-inputs.md
  - docs/00-project/quests/0043-add-a-latching-power-switch-interpreter.md
  - docs/00-project/quests/0045-add-a-hardware-read-only-gpio-probe-command.md
  - docs/03-operations/gpi-case-gpio-probe-ledger.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - internal/input
last_updated: 2026-07-07
---

# GPIO Read-Only Plan

> The next hardware lantern should only watch the trail. It should not move a
> single stone.

For the larger GPi Case 2 replacement compass, see the
[SafeShutdown replacement boundary map](safeshutdown-replacement-boundary-map.md)
before planning any service install, GPIO write, shutdown execution,
`rc.local` change, or `SafeShutdown.py` replacement.

## Current Safe Input Loop

RetroFlag Power already has a deterministic dry-run path for a power-button
observer event:

```text
CLI fake observer
  -> input observer
  -> raw signal
  -> configured latching switch interpretation
  -> power switch event
  -> power intent
  -> config policy
  -> deterministic plan
  -> noop execution
  -> event breadcrumbs
```

The command-line field kit is:

```sh
go run ./cmd/retroflag-powerd --fake-power-button-observer
go run ./cmd/retroflag-powerd --fake-power-signal low
```

The fake observer command emits one fake `power_button_pressed` input event,
maps it to the existing power-button intent, applies the current
`power_button_action` policy, builds a deterministic plan, executes only the
noop action, and prints the breadcrumb ledger. The fake raw signal command
creates one configured power input signal, interprets it through the configured
latching switch map, and only sends interpreted `SwitchOff` into that same
noop power path.

The code boundary is `internal/input.Observer`. Future GPIO-backed observers
should first report raw observations with `input.SignalEvent(name, state)`,
where `state` is `SignalLow`, `SignalHigh`, or `SignalUnverified`. Those names
mean only what the wire or logical input appears to be doing.

The first interpretation layer is deliberately configured. A latching power
switch interpreter accepts an explicit `active_signal` of `low` or `high` and an
explicit `active_switch_state` of `off` or `on`. It maps raw `SignalLow` and
`SignalHigh` observations to `SwitchOff` or `SwitchOn` only by that declared
map. `SignalUnverified` becomes `SwitchUnknown`, keeping the lantern honest
when the raw observation has not earned meaning yet.

The first hardware-facing lantern is a read-only probe command. In local
development it can be run from a repository checkout:

```sh
go run ./cmd/retroflag-powerd --probe-gpio-signal 4
```

For GPi Case field work, do not assume the repository exists on the handheld;
use a copied binary or a later explicit development-checkout procedure. Run it
with a candidate BCM GPIO pin number. The command reports only raw signal
vocabulary:

```text
SignalLow
SignalHigh
SignalUnverified
```

It does not map the wire to `SwitchOn`, `SwitchOff`, or `SwitchUnknown`; those
belong to the later interpretation layer. If the platform is unsupported, the
pin is not readable, the GPIO tools are missing, or the output cannot be trusted
as a raw low/high value, the command reports `SignalUnverified` instead of
guessing.

Only after that interpretation should the app receive project-level meaning
such as `input.PowerSwitchEvent(input.SwitchOff)` and turn it into a power
intent. The current dry-run/noop route maps interpreted `SwitchOff` to the same
safe power intent path already used by the fake power-button observer. Core
app, planner, config, and executor code should continue to receive interpreted
input events and power intents, not GPIO chips, line offsets, edge names, or
electrical polarity.

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

The probe command also stays outside daemon behavior. It does not start the app
lifecycle, execute shutdown, install services, touch `rc.local`, replace
`SafeShutdown.py`, write GPIO, or persist state. It is a field-kit lantern for
looking at one candidate wire.

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

If a later hardware quest provides a copied `retroflag-powerd` binary or an
explicit development checkout on the device, the project probe can be used
during the same supervised observation session:

```sh
./retroflag-powerd --probe-gpio-signal 4
```

Record the command, pin number, case control position, and raw result exactly as
printed in the
[GPi Case GPIO probe field ledger](gpi-case-gpio-probe-ledger.md). Treat
`SignalLow` and `SignalHigh` as electrical observations only, not as
power-switch meaning.

## Device Facts Still Needed

The project still needs facts from the device before a real observer can be
trusted:

- Candidate GPIO chips and line offsets for the power switch.
- Candidate GPIO chips and line offsets for the reset button, if included in
  the first input quest.
- Active-high or active-low behavior for each candidate input.
- Whether each control behaves like a latching switch or a momentary button.
- Which configured interpretation maps each raw signal state to a switch or
  button meaning.
- Which latching power switch config is correct: `active_signal` low or high,
  and `active_switch_state` off or on.
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
- The fake raw signal CLI path can report `low`, `high`, and `unverified`
  observations deterministically without reading GPIO.
- The read-only observer can emit raw `SignalLow`, `SignalHigh`, and
  `SignalUnverified` observations without claiming button or switch meaning too
  early.
- A configured interpretation step maps trusted raw observations into either a
  latching switch event or a momentary button event before any power intent is
  produced.
- The latching power switch interpretation remains explicit and deterministic:
  raw signal, configured latching switch interpretation, power switch event,
  power intent, policy, plan, noop execution, and breadcrumbs.
- Candidate pins, polarity, latching or momentary behavior, and bounce/noise
  notes are documented with device context.
- The resulting plan and execution remain deterministic and noop-only.
