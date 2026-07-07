---
id: QUEST-0048
title: Map the SafeShutdown Replacement Boundaries
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a documentation-only replacement boundary map for the behavior retroflag-powerd must preserve before it can safely replace the stock RetroFlag SafeShutdown.py path on GPi Case 2 hardware.
related:
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - docs/03-operations/gpio-read-only-plan.md
  - docs/03-operations/gpi-case-gpio-probe-ledger.md
last_updated: 2026-07-07
---

# QUEST-0048 - Map the SafeShutdown Replacement Boundaries

> Put a compass beside the old RetroFlag power relic so future quests know
> which behaviors must be preserved before the satchel changes hands.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the SafeShutdown replacement boundary map at
  [docs/03-operations/safeshutdown-replacement-boundary-map.md](../../03-operations/safeshutdown-replacement-boundary-map.md).
- Preserved the verified GPi Case 2 findings that the side switch does not
  directly cut battery power and that the stock `SafeShutdown.py` script is
  part of the active power-control path.
- Recorded the observed GPIO responsibilities: GPIO26 as the side-switch
  shutdown signal and GPIO27 driven HIGH as a power-enable latch.
- Recorded that the legacy script uses separate multiprocessing workers for
  `poweroff()` and `lcdrun()`.
- Marked top-button power-save/resume behavior as still part of the legacy
  `lcdrun()` path.
- Added the verified finding that the RetroFlag legacy script path also appears
  to participate in docking behavior, including switching between the built-in
  LCD and HDMI when docked.
- Called out the recovery risk that sleep or power-save can make SSH harder if
  Wi-Fi goes down.
- Listed the replacement prerequisites: power-enable latch behavior,
  side-switch shutdown detection, top-button power-save/resume behavior, clean
  EmulationStation and Linux shutdown sequencing, and KMS-safe display
  behavior, handheld LCD behavior, docked HDMI behavior, LCD/HDMI transitions,
  and KMS display timing or ordering dependencies.
- Added the caution that audio after the FKMS-to-KMS update has not been fully
  verified yet, so future field testing must include audio checks in both
  handheld and docked modes before any replacement plan is complete.
- Added an explicit no-go list for GPIO writes, shutdown execution, service
  install, systemd activation, `rc.local` changes, `SafeShutdown.py`
  replacement, persistent state, and daemon activation.
- Suggested a conservative future quest sequence from read-only observations to
  latch documentation, fake policy modeling, dry-run service planning, unit
  drafting, hardware field testing, and only then replacement planning.
- Linked the map from `README.md`, the GPi Case 2 hardware notes, the GPIO
  read-only plan, and the GPi Case GPIO probe field ledger.

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

## Milestone Note

The project now has a safety compass for the GPi Case 2 replacement trail.
Future maintainers can see the behaviors `retroflag-powerd` must preserve
before the stock RetroFlag script is touched.
