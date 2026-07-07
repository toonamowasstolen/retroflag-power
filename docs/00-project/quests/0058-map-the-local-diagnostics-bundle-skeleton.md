---
id: QUEST-0058
title: Map the Local Diagnostics Bundle Skeleton
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Refine the local diagnostics bundle map with a concrete future skeleton while keeping bundle generation, network submission, hardware behavior, and installer work out of scope.
related:
  - docs/03-operations/local-diagnostics-bundle-map.md
  - CLAUDE.md
  - docs/00-project/edc-quest-operating-rules.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/03-operations/gpi-case-2-acceptance-checklist.md
  - docs/03-operations/gpi-case-gpio-probe-ledger.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - docs/00-project/quests/0051-map-the-local-diagnostics-bundle.md
last_updated: 2026-07-07
---

# QUEST-0058 - Map the Local Diagnostics Bundle Skeleton

> Shape the future diagnostics satchel before anyone builds it: local,
> readable, redacted, and quiet enough to leave every wire untouched.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Updated the existing
  [Local Diagnostics Bundle Map](../../03-operations/local-diagnostics-bundle-map.md)
  instead of creating a duplicate diagnostics compass.
- Added a future bundle skeleton covering version output, dry-run startup
  summary, runtime/startup diagnostics, execution status, event breadcrumbs,
  configured hardware profile name, OS/Pi facts, display/KMS/FKMS facts, audio
  facts, SafeShutdown observations, raw GPIO observations, and GPi Case 2 field
  checklist references.
- Separated allowed read-only diagnostics from forbidden active behavior.
- Preserved the vocabulary boundary between Lanterns, Lantern Dispatch, raw
  `SignalLow`/`SignalHigh`/`SignalUnverified` observations, and interpreted
  `SwitchOn`/`SwitchOff`/`SwitchUnknown` meanings.
- Added future optional user-redacted bundle contents.
- Marked future Lantern Dispatch submission ideas as optional and not
  implemented.

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
- No diagnostics bundle generation.

## Milestone Note

The diagnostics trail now has a clearer local skeleton: enough structure for a
future implementation quest to follow, and enough safety rail to keep today's
prototype read-only.
