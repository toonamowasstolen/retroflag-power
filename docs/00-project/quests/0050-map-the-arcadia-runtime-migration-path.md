---
id: QUEST-0050
title: Map the Arcadia Runtime Migration Path
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a documentation-only staged migration map for moving from the current RetroFlag Power GPi Case 2 prototype toward the broader Arcadia Runtime direction.
related:
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - README.md
  - PROJECT_MANIFEST.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/04-architecture/system-overview.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
last_updated: 2026-07-07
---

# QUEST-0050 - Map the Arcadia Runtime Migration Path

> Set the migration compass beside the lantern: RetroFlag Power keeps the GPi
> Case 2 field trail steady today, while Arcadia Runtime marks the broader
> engine path for a later epoch.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the staged migration map at
  [docs/04-architecture/arcadia-runtime-migration-path.md](../../04-architecture/arcadia-runtime-migration-path.md).
- Recorded why the migration must stay staged while RetroFlag Power continues
  careful GPi Case 2 discovery.
- Captured the current RetroFlag Power prototype phase: raw signal vocabulary,
  interpreted switch vocabulary, dry-run/noop power path, read-only GPIO
  lantern, SafeShutdown boundary mapping, KMS/docking/display findings, and no
  real hardware mutation.
- Captured Arcadia Runtime as the favored future engine direction, not an
  active rename.
- Mapped future Relic profile ideas conceptually without introducing schemas,
  loaders, installers, or support claims.
- Mapped SignalMage, Casters, Spellbooks, Lanterns, Ledgers, and Lantern
  Dispatch as project language without forcing package, binary, service, or
  import-path renames.
- Added conservative migration stages from current docs and dry-run safety
  through local diagnostics, reversible installer planning, possible future
  rename gates, optional Lantern Dispatch, and later hardware profiles.
- Added objective rename gates around stable GPi Case 2 behavior, rollback,
  KMS, audio, SafeShutdown replacement prerequisites, installer dry-run, naming
  clearance, and unofficial/non-affiliation language.
- Added explicit no-go guidance for mid-discovery renames, premature hardware
  broadening, telemetry before consent, implied RetroFlag endorsement, and
  unsafe SafeShutdown replacement.
- Linked the migration path from `README.md`, `PROJECT_MANIFEST.md`,
  `docs/00-project/project-direction-save-room-arcadia.md`,
  `docs/04-architecture/system-overview.md`, and
  `docs/03-operations/safeshutdown-replacement-boundary-map.md`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No repo rename.
- No module rename.
- No binary or package rename.
- No import path changes.
- No domain claims.
- No telemetry implementation.
- No network calls.
- No installer implementation.
- No service install.
- No systemd activation.
- No `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No config mutation.
- No daemon activation.

## Milestone Note

The project now has a cautious migration compass. RetroFlag Power stays focused
on trustworthy GPi Case 2 discovery, while Arcadia Runtime remains visible as a
future engine direction that must be earned stage by stage.
