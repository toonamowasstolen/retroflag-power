---
title: GPi Case GPIO Probe Field Ledger
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record manual read-only GPi Case GPIO probe observations before any raw signal is interpreted as switch meaning.
related:
  - docs/03-operations/gpio-read-only-plan.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - docs/00-project/quests/0047-add-a-gpi-case-gpio-probe-field-ledger.md
last_updated: 2026-07-07
---

# GPi Case GPIO Probe Field Ledger

> A lantern can show the wire without naming the spell. This ledger keeps the
> GPi Case trail honest until the observations have earned a map.

This field ledger records manual, read-only GPIO probe observations on the GPi
Case. Each row should capture the candidate BCM pin, what the case control was
doing, whether the device was docked or handheld, and the exact raw signal
reported by the probe command.

This page is for raw observations only. Do not interpret `SignalLow` or
`SignalHigh` as power on, power off, switch on, or switch off meaning here.
The broader replacement compass lives in the
[SafeShutdown replacement boundary map](safeshutdown-replacement-boundary-map.md).

## Safe Probe Command

Run the probe on the GPi Case with one candidate BCM pin only after a later
hardware quest provides a copied `retroflag-powerd` binary or an explicit
development checkout on the device:

```sh
./retroflag-powerd --probe-gpio-signal <pin>
```

The command output is raw wire state only:

```text
SignalLow
SignalHigh
SignalUnverified
```

`SignalLow`, `SignalHigh`, and `SignalUnverified` are not interpreted switch
states. Interpreted switch states use different vocabulary:

```text
SwitchOn
SwitchOff
SwitchUnknown
```

Do not write `SwitchOn`, `SwitchOff`, or `SwitchUnknown` in the raw result
column unless a later interpretation quest explicitly asks for a separate
interpreted ledger.

## Observation Ledger

| Date | Device / image | Command run | BCM pin | Case state / control position | Docked or handheld | Raw result | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| YYYY-MM-DD | GPi Case model, Pi model, OS image | `./retroflag-powerd --probe-gpio-signal <pin>` | `<pin>` | Describe only what the case control was doing | Docked / handheld | `SignalUnverified` | Note uncertainty, repeated readings, or nearby services |

## Field Notes

- Record the command exactly as run.
- Record the BCM pin number exactly as tested.
- Record the visible or physical case control position without assigning power
  meaning to the raw signal.
- Repeat observations when the same control appears to change state.
- Prefer `SignalUnverified` over guesses when the command reports uncertainty.
- Link any larger hardware findings back to the GPi Case hardware notes.

## Do Not Do Yet

- Do not install as a service.
- Do not replace `SafeShutdown.py`.
- Do not wire this into shutdown.
- Do not interpret raw signal as switch meaning until the observations are
  reviewed.
- Do not perform GPIO writes.
- Do not activate systemd units.
- Do not change `rc.local`.
- Do not add persistent state.
- Do not activate daemon behavior from this ledger.

## Review Gate

Before any later quest maps raw signal states into switch states, review the
ledger for repeated observations across the same device, image, command, pin,
case control position, and docked or handheld context. The compass point should
come from evidence, not from the first low or high value the lantern sees.
