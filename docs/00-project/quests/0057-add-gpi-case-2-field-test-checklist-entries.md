---
id: QUEST-0057
title: Add GPi Case 2 Field Test Checklist Entries
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add practical GPi Case 2 field-test entries for the known KMS, audio, dock, power-save, input, and SafeShutdown unknowns before any replacement work begins.
related:
  - ../../03-operations/gpi-case-2-acceptance-checklist.md
  - ../../03-operations/gpi-case-gpio-probe-ledger.md
  - ../../03-operations/safeshutdown-replacement-boundary-map.md
  - ../../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - 0053-add-the-gpi-case-2-acceptance-checklist.md
  - 0048-map-the-safeshutdown-replacement-boundaries.md
last_updated: 2026-07-07
---

# QUEST-0057 - Add GPi Case 2 Field Test Checklist Entries

> Turn the fresh hardware signs into rows the next maintainer can carry into
> the field: audio, dock, LCD wake, side switch, buttons, and controller
> identity all get their own lantern.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added current field-test entries to the
  [GPi Case 2 Acceptance Checklist](../../03-operations/gpi-case-2-acceptance-checklist.md).
- Added practical pass/fail/unknown rows for handheld audio after KMS, docked
  audio after KMS, LCD sleep/wake behavior, top-button power-save/resume
  behavior, side-switch behavior while `SafeShutdown.py` is still active,
  LCD-to-HDMI switching, HDMI-to-LCD switching, docking behavior, the button
  above Select and left of the RetroFlag logo, and controller identity.
- Preserved the current field note that the built-in controller has been
  observed as a Microsoft Xbox 360 gamepad / `GBA Pi Case+` / Nuvoton device.
- Preserved the current field note that the button above Select and left of
  the RetroFlag logo was not detected during EmulationStation mapping.
- Repeated the vocabulary boundary between raw signal observations
  `SignalLow`, `SignalHigh`, and `SignalUnverified`, and interpreted switch
  meanings `SwitchOn`, `SwitchOff`, and `SwitchUnknown`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No GPIO writes.
- No shutdown execution.
- No systemd activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume implementation.
- No persistent state.
- No telemetry.
- No network calls.
- No installer or packaging changes.

## Milestone Note

The GPi Case 2 acceptance gate now has ready-to-fill field rows for the
specific hardware unknowns discovered during the KMS and power investigation.
