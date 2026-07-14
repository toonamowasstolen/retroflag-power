---
id: QUEST-0047
title: Add a GPi Case GPIO Probe Field Ledger
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a small documentation-only ledger for safe manual read-only GPIO probe observations on the GPi Case.
related:
  - docs/03-operations/gpi-case-gpio-probe-ledger.md
  - docs/03-operations/gpio-read-only-plan.md
  - docs/02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - docs/00-project/quests/0045-add-a-hardware-read-only-gpio-probe-command.md
last_updated: 2026-07-07
---

# QUEST-0047 - Add a GPi Case GPIO Probe Field Ledger

> Give the new hardware lantern a field ledger: candidate pin, case control,
> raw signal, and no borrowed meaning before the map is reviewed.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the GPi Case GPIO probe field ledger at
  [docs/03-operations/gpi-case-gpio-probe-ledger.md](../../03-operations/gpi-case-gpio-probe-ledger.md).
- Included a simple observation table for date, device or image, command, BCM
  pin, case state or control position, docked or handheld context, raw result,
  and notes.
- Recorded the original local-development read-only command:
  `go run ./cmd/retroflag-powerd --probe-gpio-signal <pin>`.
- Made the raw vocabulary explicit: `SignalLow`, `SignalHigh`, and
  `SignalUnverified`.
- Kept interpreted switch vocabulary separate: `SwitchOn`, `SwitchOff`, and
  `SwitchUnknown`.
- Warned that `SignalLow` and `SignalHigh` must not be labeled as power on,
  power off, switch on, or switch off meaning in the raw ledger.
- Added the "Do Not Do Yet" field-kit boundary for services, shutdown wiring,
  `SafeShutdown.py`, GPIO writes, `rc.local`, systemd activation, persistence,
  and daemon activation.
- Linked the ledger from the GPIO read-only plan and the project documentation
  list in `README.md`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No GPIO behavior changes.
- No GPIO writes.
- No shutdown execution.
- No service install.
- No systemd activation.
- No `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No persistent state.
- No daemon activation.
- No raw-signal-to-switch-meaning inference.

## Milestone Note

The GPi Case hardware lantern now has a small field ledger. Future probe
sessions can record candidate pins, case controls, and raw signal states in one
place before any quest turns those observations into interpreted switch meaning.
