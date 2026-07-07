---
id: QUEST-0053
title: Add the GPi Case 2 Acceptance Checklist
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a documentation-only acceptance checklist for GPi Case 2 readiness before SafeShutdown.py replacement, service install, public installer release, or Arcadia Runtime migration.
related:
  - docs/03-operations/gpi-case-2-acceptance-checklist.md
  - README.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/03-operations/installer-migration-toolkit-map.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
last_updated: 2026-07-07
---

# QUEST-0053 - Add the GPi Case 2 Acceptance Checklist

> Set the acceptance lantern at the gate: before the runtime touches real GPi
> Case 2 behavior, the field ledger must prove power, display, dock, audio,
> sleep, resume, rollback, and user safety.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the GPi Case 2 acceptance checklist at
  [docs/03-operations/gpi-case-2-acceptance-checklist.md](../../03-operations/gpi-case-2-acceptance-checklist.md).
- Framed the checklist as a field-test gate before replacing
  `SafeShutdown.py`, installing or enabling a daemon service, publishing a
  public installer, or broadening toward Arcadia Runtime.
- Recorded the current no-go status: `SafeShutdown.py` is not replaced, no
  service is installed, no shutdown execution exists, no GPIO writes exist, no
  config mutation exists, and the checklist is not yet passed.
- Added hardware identity checks for Raspberry Pi or CM4 details, GPi Case 2
  notes, OS and kernel versions, RetroPie or EmulationStation versions, KMS or
  FKMS display stack, dock state, and audio device presence.
- Added power and GPIO checks for GPIO26 side-switch observations, GPIO27
  latch understanding, raw `SignalLow`, `SignalHigh`, and `SignalUnverified`
  observations, profile interpretation boundaries, stock script behavior, and
  safe documentation before any stock-script-disabled test.
- Added display and dock checks for handheld LCD boot, docked HDMI boot,
  transitions, KMS behavior, timing dependencies, and recovery if the screen
  goes dark.
- Added sleep/resume checks for top-button power-save, resume, SSH/network
  behavior during sleep, wake behavior, and access-loss risk notes.
- Added audio checks for handheld and docked audio after KMS migration,
  RetroPie game audio, EmulationStation menu audio, and known unknowns.
- Added RetroPie and EmulationStation checks for startup, controller
  detection, Xbox 360 gamepad mapping, the currently undetected button above
  Select and left of the RetroFlag logo, safe game exit, and clean shutdown
  expectations.
- Added diagnostics and rollback checks for future local diagnostics bundles,
  future installer previews, future backup and restore ledgers, no required
  network submission, and rollback documentation before mutation.
- Added public readiness gates for maintainer GPi Case 2 checklist pass,
  SafeShutdown boundary prerequisites, stable KMS handheld/docked behavior,
  audio verification, rollback testing, unofficial/non-affiliation language,
  and naming/domain clearance before public Arcadia Runtime release.
- Added a compact field ledger table template for date, device, OS/kernel,
  mode, test area, expected result, observed result, pass/fail/unknown, and
  notes.
- Linked the checklist from `README.md`,
  `docs/03-operations/safeshutdown-replacement-boundary-map.md`,
  `docs/03-operations/installer-migration-toolkit-map.md`,
  `docs/03-operations/local-diagnostics-bundle-map.md`, and
  `docs/04-architecture/arcadia-runtime-migration-path.md`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No installer implementation.
- No diagnostics implementation.
- No file mutation implementation.
- No config mutation implementation.
- No service install.
- No systemd activation.
- No `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No GPIO behavior changes.
- No GPIO writes.
- No shutdown execution.
- No telemetry implementation.
- No network calls.
- No daemon activation.

## Milestone Note

The project now has a concrete GPi Case 2 acceptance gate. Before future
runtime work touches real behavior, maintainers can see which field badges must
be earned and which unknowns still need a lantern.
