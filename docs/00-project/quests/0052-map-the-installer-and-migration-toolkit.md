---
id: QUEST-0052
title: Map the Installer and Migration Toolkit
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a documentation-only map for a future local-first, reversible installer and migration toolkit for RetroFlag Power and future Arcadia Runtime field kits.
related:
  - docs/03-operations/installer-migration-toolkit-map.md
  - README.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
last_updated: 2026-07-07
---

# QUEST-0052 - Map the Installer and Migration Toolkit

> Set the installer compass beside the field kit: inspect first, pack the
> backup satchel, preview the route, apply only with consent, and leave a
> restore ledger behind.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the installer and migration toolkit map at
  [docs/03-operations/installer-migration-toolkit-map.md](../../03-operations/installer-migration-toolkit-map.md).
- Framed the toolkit as a future local-first, reversible path for moving from
  legacy RetroFlag scripts and older Raspberry Pi OS assumptions toward
  RetroFlag Power and future Arcadia Runtime field-kit behavior.
- Recorded reversible rules for inspect-before-change, backup-before-replace,
  preview, explicit confirmation, restore instructions, uninstall and restore,
  and offline local install or rollback.
- Sketched proposed future command shapes, including
  `retroflag-powerd installer inspect`, `retroflag-powerd installer plan`,
  `retroflag-powerd installer backup`, `retroflag-powerd installer apply`,
  `retroflag-powerd installer restore`, and
  `save-room fieldkit install arcadia-gpi-case-2`.
- Added a legacy detection checklist for `/opt/RetroFlag/SafeShutdown.py`,
  related scripts, `rc.local`, systemd units, display config, FKMS/KMS hints,
  `cmdline.txt`, RetroPie or EmulationStation, GPi Case 2 dock/display notes,
  audio config, and existing `retroflag-powerd` binary or config.
- Mapped backup expectations for copying legacy scripts, recording file paths,
  checksums, timestamps, original config snippets, and a human-readable restore
  ledger while avoiding private ROM and library data.
- Captured the plan/apply model: show files to back up, files to edit, services
  to add or disable, config lines to add or remove, exact restore path, risk
  notes, and reboot requirements before applying anything.
- Connected SafeShutdown replacement gates to the existing boundary map for
  latch behavior, side-switch shutdown, top-button power-save/resume, handheld
  LCD, docked HDMI, KMS timing, audio, and rollback.
- Recorded the cautious RetroPie integration path: start with the project's own
  installer, later draft a RetroPie-Setup-compatible scriptmodule, test
  externally and community-first, avoid endorsement claims, and only approach
  maintainers after stable behavior, docs, rollback, and support workflow
  exist.
- Linked installer inspect and plan output to future local diagnostics and
  field-kit feature requests without adding submission, Lantern Dispatch, or
  network behavior.
- Added explicit no-go guidance against silent config mutation, unbacked
  `SafeShutdown.py` replacement, unconfirmed service install, required network
  access, ROM/library collection, official support claims, and skipped rollback
  documentation.
- Linked the map from `README.md`,
  `docs/04-architecture/arcadia-runtime-migration-path.md`,
  `docs/00-project/project-direction-save-room-arcadia.md`,
  `docs/03-operations/local-diagnostics-bundle-map.md`, and
  `docs/03-operations/safeshutdown-replacement-boundary-map.md`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No installer implementation.
- No shell installer implementation.
- No RetroPie scriptmodule implementation.
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

The project now has a safe installer compass. Before any future script mutates
a GPi Case, maintainers can see what the toolkit must detect, back up, preview,
apply, uninstall, and restore.
