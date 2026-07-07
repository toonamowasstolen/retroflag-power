---
id: ARCH-ARCADIA-RUNTIME-MIGRATION-PATH-001
title: Arcadia Runtime Migration Path
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map the cautious staged path from the current RetroFlag Power GPi Case 2 prototype toward the broader Arcadia Runtime direction without triggering a rename, install, service activation, or hardware behavior change.
related:
  - README.md
  - PROJECT_MANIFEST.md
  - docs/00-project/edc-quest-operating-rules.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/04-architecture/system-overview.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/03-operations/installer-migration-toolkit-map.md
  - docs/03-operations/gpi-case-2-acceptance-checklist.md
  - docs/00-project/quests/0050-map-the-arcadia-runtime-migration-path.md
  - docs/00-project/quests/0051-map-the-local-diagnostics-bundle.md
  - docs/00-project/quests/0052-map-the-installer-and-migration-toolkit.md
  - docs/00-project/quests/0053-add-the-gpi-case-2-acceptance-checklist.md
last_updated: 2026-07-07
---

# Arcadia Runtime Migration Path

> This is the migration compass beside the current lantern: RetroFlag Power
> keeps mapping the GPi Case 2 trail today, while Arcadia Runtime remains the
> wider engine path the project may earn later.

This document is documentation only. It does not rename the repository, rename
the Go module, rename binaries or packages, change imports, claim any domain,
implement telemetry, make network calls, install services, activate daemons,
mutate configuration, alter `rc.local`, replace `/opt/RetroFlag/SafeShutdown.py`,
write GPIO, or execute shutdown.

RetroFlag Power remains the current working repository, prototype name, and GPi
Case 2-focused field kit. Arcadia Runtime is the favored future engine
direction, not an active rename.

Future agents should follow the
[EDC Quest Operating Rules](../00-project/edc-quest-operating-rules.md) before
editing this map: read the relevant EDC docs, name the source-of-truth files,
keep the quest small, record discoveries, and stop when docs and implementation
disagree.

## Why Migration Is Staged

RetroFlag Power is still doing careful GPi Case 2 discovery. The project is
mapping real hardware behavior around power, GPIO signals, display modes,
docking, audio, power-save, resume, and the stock RetroFlag
`SafeShutdown.py` path.

That work needs trust more than speed. Before any broader platform move, public
rename, installer, or replacement path, the current prototype must preserve
working hardware behavior and make every safety boundary visible.

The staged path exists so the team can grow toward Save Room Tech and Arcadia
Runtime without confusing the current GPi Case work or making premature claims
about what the runtime owns.

## Current Phase: RetroFlag Power Prototype

The current phase is the RetroFlag Power prototype.

It includes:

- GPi Case 2-focused discovery.
- Raw GPIO signal vocabulary.
- Interpreted switch vocabulary.
- Dry-run/noop power path.
- Read-only GPIO lantern.
- `SafeShutdown.py` boundary mapping.
- KMS, docking, and display findings.
- No real shutdown, no GPIO writes, and no service replacement.

This phase is about building ledgers before machinery. The project can observe,
interpret, and model safe paths, but it does not yet own real GPi Case 2 power
behavior.

## Future Phase: Arcadia Runtime Engine

Arcadia Runtime is the favored future direction for a broader runtime that
preserves hidden retro hardware behavior on modern systems.

If the project earns that step, Arcadia Runtime may eventually own:

- GPIO and signal observation.
- Switch interpretation.
- Safe shutdown policy.
- Display and dock behavior.
- Resume and power-save behavior.
- Device profiles.
- Diagnostics.
- Migration from legacy vendor scripts.

This future phase should remain grounded in verified ledgers. Arcadia Runtime
should not become a brand, package name, daemon name, or public release promise
until the rename gates in this map are satisfied.

The local-first diagnostics compass lives in
[Local Diagnostics Bundle Map](../03-operations/local-diagnostics-bundle-map.md).
That map keeps future support bundles redacted, previewable, and useful before
any optional Lantern Dispatch path exists.

The future installer compass lives in
[Installer and Migration Toolkit Map](../03-operations/installer-migration-toolkit-map.md).
That map keeps local install, update, uninstall, and restore behavior
inspectable and reversible before any field kit mutates a device.

The GPi Case 2 public-readiness gate lives in
[GPi Case 2 Acceptance Checklist](../03-operations/gpi-case-2-acceptance-checklist.md).
That ledger must be passed before the Arcadia Runtime trail becomes a public
release promise.

## Device Profile Model: Relics

Hardware support should eventually move toward Relics: explicit device or
hardware profiles with documented behavior, boundaries, and acceptance notes.

Possible future Relics include:

- GPi Case 2 relic/profile.
- NESPi-style case relic/profile.
- PiBoy-like relic/profile.
- Custom build relic/profile.

This is conceptual only. This map does not introduce a profile schema, load a
profile, install a field kit, or broaden support beyond the current GPi Case 2
prototype.

## Component Model

The current and future naming model should stay clear and useful:

| Term | Migration meaning |
| --- | --- |
| SignalMage | Signal observation and interpretation layer. |
| Casters | Focused daemons or workers such as power, display, dock, and resume. |
| Spellbooks | Legacy scripts and migration recipes, including `SafeShutdown.py` behavior maps. |
| Lanterns | Read-only probes and diagnostics. |
| Ledgers | Field notes, compatibility records, probe observations, and acceptance evidence. |
| Lantern Dispatch | Future optional update, diagnostics, and support submission layer. |

These names are project language, not an instruction to rename Go packages,
exported types, binaries, services, or import paths.

## Migration Stages

The conservative staged list is:

1. Stage 0: Current RetroFlag Power docs and dry-run safety.
2. Stage 1: Complete GPi Case 2 field ledgers.
3. Stage 2: Model `SafeShutdown.py` behavior fully.
4. Stage 3: Add local-only diagnostics bundles.
5. Stage 4: Draft installer/migration toolkit, no mutation by default.
6. Stage 5: Support reversible install/uninstall.
7. Stage 6: Consider repo/module/binary rename only after stability.
8. Stage 7: Introduce optional Lantern Dispatch integration.
9. Stage 8: Consider additional Relics/hardware profiles.

Each stage should leave a visible badge in documentation or tests before the
next stage depends on it.

## Rename Gates

Before any repository, module, package, binary, or service rename, the project
needs objective gates:

- Stable GPi Case 2 behavior.
- Documented rollback.
- Tested KMS handheld and docked behavior.
- Audio verified handheld and docked.
- `SafeShutdown.py` replacement prerequisites understood.
- Installer dry-run complete.
- Public naming clearance done.
- Unofficial/non-affiliation language in place.

Until these gates are met, RetroFlag Power remains the working name and Arcadia
Runtime remains the favored future engine direction.

## No-Go List

The migration path does not permit:

- Do not rename mid-discovery.
- Do not broaden before GPi Case 2 is trustworthy.
- Do not add telemetry before local diagnostics and a consent model.
- Do not imply RetroFlag endorsement.
- Do not replace `SafeShutdown.py` until all boundary prerequisites are met.

If a future quest needs to cross one of these lines, it needs its own acceptance
criteria, rollback notes, and maintainer decision.

## Suggested Next Quest Sequence

Useful future quests:

- Add local diagnostics bundle map.
- Add installer/migration toolkit map.
- Add Relic profile schema map.
- Add Lantern Dispatch privacy model.
- Add naming/domain clearance checklist.
- Fill the GPi Case 2 acceptance checklist for KMS, dock, audio, sleep/resume,
  shutdown, diagnostics, rollback, and public readiness.

The compass for now is simple: keep the current lantern honest, keep the
prototype safe, and let the larger Arcadia Runtime trail open only when the
field ledgers say it is ready.
