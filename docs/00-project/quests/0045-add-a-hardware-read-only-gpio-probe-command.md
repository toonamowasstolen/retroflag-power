---
id: QUEST-0045
title: Add a Hardware Read-Only GPIO Probe Command
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Give the daemon its first real hardware-facing read-only GPIO lantern for reporting raw signal state only.
related:
  - cmd/retroflag-powerd
  - internal/gpio
  - internal/input
  - README.md
  - docs/03-operations/gpio-read-only-plan.md
  - docs/00-project/quests/0044-add-a-fake-power-signal-cli-path.md
last_updated: 2026-07-07
---

# QUEST-0045 - Add a Hardware Read-Only GPIO Probe Command

> Light one hardware lantern at the edge of the GPi Case trail: read the wire,
> name the raw signal, and leave every power spell untouched.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added `--probe-gpio-signal <pin>` to `retroflag-powerd`.
- Added `internal/gpio` as a small read-only probe boundary.
- On Linux, the probe attempts read-only signal checks from already exposed
  GPIO value sources and `gpioget` input reads, then falls back safely.
- On unsupported platforms or uncertain reads, the probe reports
  `SignalUnverified` deterministically.
- The command prints raw signal vocabulary only: `SignalLow`, `SignalHigh`, or
  `SignalUnverified`.
- The command explicitly avoids switch interpretation; it does not report
  `SwitchOn`, `SwitchOff`, or `SwitchUnknown`.
- Added tests for CLI output, invalid pin handling, raw signal display labels,
  and platform-neutral fallback behavior.
- Documented how to run the probe on the GPi Case during a supervised hardware
  observation session.

## Validation

- [x] `go test ./...` passed during implementation.
- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Read-only only.
- No GPIO writes.
- No shutdown command execution.
- No service install.
- No systemd activation.
- No `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No resume behavior.
- No persistent state.
- No daemon behavior changes beyond the explicit CLI/dev probe command.
- No switch interpretation in probe output.

## Safe Command

```sh
go run ./cmd/retroflag-powerd --probe-gpio-signal 4
```

On the GPi Case, replace `4` with the candidate BCM GPIO pin under observation.
The result is a raw wire report, not a power decision.

## Milestone Note

The daemon now has its first hardware-facing read-only lantern. It can look at
a candidate GPi Case pin, report what the wire appears to say, and keep the
satchel closed on shutdown, services, persistence, and interpretation.
