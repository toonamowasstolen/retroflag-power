---
id: PROJECT-DIRECTION-SAVE-ROOM-ARCADIA-001
title: Save Room Tech and Arcadia Runtime Direction
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record the long-term naming, platform, privacy, and migration direction around Save Room Tech, Arcadia Runtime, and the current RetroFlag Power prototype.
related:
  - README.md
  - PROJECT_MANIFEST.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/04-architecture/system-overview.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/00-project/quests/0049-record-the-save-room-tech-and-arcadia-runtime-direction.md
last_updated: 2026-07-07
---

# Save Room Tech and Arcadia Runtime Direction

> This is the north-star map for the wider adventure: RetroFlag Power remains
> the GPi Case 2 lantern in hand, while Save Room Tech and Arcadia Runtime mark
> the larger trail the project may walk later.

This document is documentation only. It does not rename the repository, rename
packages or binaries, claim any domain, implement telemetry, make network
calls, install services, activate daemons, mutate configuration, alter
`rc.local`, or replace `/opt/RetroFlag/SafeShutdown.py`.

The staged path from this prototype toward the favored future engine direction
is tracked in
[Arcadia Runtime Migration Path](../04-architecture/arcadia-runtime-migration-path.md).
That map keeps the current GPi Case 2 work separate from future rename,
installer, diagnostics, and field-kit questions.

## Current Project State

RetroFlag Power remains the current repository and working prototype name.

For now, the project is focused on GPi Case 2 power, signal, display, docking,
and runtime discovery. Its immediate duty is still careful field mapping around
the stock RetroFlag behavior before any replacement trail begins.

No rename happens as part of this direction record. The current Go module,
binary names, package names, documentation paths, and local development
commands remain unchanged.

RetroFlag Power should continue to be described as an unofficial community
project when context requires it. It should not imply affiliation with,
endorsement by, or ownership by RetroFlag.

## Future Umbrella

Save Room Tech is the favored future umbrella and home for hobby/community
tools around old machines and new adventures.

Taft Consulting remains the professional and client-facing side.

`toonamo` remains the personal creator and maintainer handle.

The future split should keep those identities clear:

| Name | Role |
| --- | --- |
| Save Room Tech | Hobby/community umbrella for retro hardware tools, docs, field kits, and support trails. |
| Taft Consulting | Professional/client-facing consulting identity. |
| toonamo | Personal creator and maintainer handle. |
| RetroFlag Power | Current repo and GPi Case-focused prototype. |
| Arcadia Runtime | Preferred future broader runtime direction. |

## Future Engine Direction

Arcadia Runtime is the preferred future project direction for the broader
engine.

Its purpose should be:

> A modern runtime for preserving hidden retro hardware behavior on newer
> systems.

Arcadia Runtime may eventually cover more than RetroFlag hardware. It can grow
toward NESPi-style cases, PiBoy-like handhelds, and custom builds if the project
earns that breadth through verified ledgers, hardware profiles, cautious
install paths, and privacy-first support tools.

The name is favored, not final. Before a public v1.0 or broad release, Arcadia
Runtime should receive a proper name, domain, and trademark clearance pass.
This record does not claim ownership of any name or domain.

## Naming Taxonomy

These names are direction markers for product, documentation, and project
language. They should not be forced into Go package names, exported types,
filenames, or low-level technical identifiers when plain engineering names are
clearer.

| Term | Meaning |
| --- | --- |
| Save Room Tech | The umbrella home for old machines and new adventures. |
| Arcadia Runtime | The broader runtime for preserving hidden retro hardware behavior on modern systems. |
| Relics | Supported devices or hardware profiles, such as GPi Case 2, NESPi-style cases, PiBoy-like handhelds, and custom builds. |
| Casters | Focused runtime workers or daemons that perform specific duties, such as a power caster, display caster, dock caster, or resume caster. |
| Spellbooks | Legacy scripts, migration recipes, and documented behavior maps, such as the SafeShutdown.py replacement boundary map. |
| SignalMage | The GPIO and signal observation and interpretation layer. |
| Lanterns | Read-only probes and diagnostics, such as the GPIO signal probe. |
| Ledgers | Field notes, compatibility records, probe observations, and test results. |
| Lantern Dispatch | A future optional update, diagnostics, issue-reporting, and support-submission layer. |
| Field kits | Installable and supportable packages or device-specific bundles. |

## Future Platform Path

Save Room Tech may eventually need its own EDC/project to support a broader
field-kit ecosystem.

That future platform could include:

- marketing and documentation site
- field kit registry
- update metadata
- diagnostics intake
- compatibility ledgers
- feature request intake from field kits
- issue submission flow
- privacy-first opt-in diagnostics

This platform does not exist yet. RetroFlag Power should not pretend it does,
and docs should avoid references that imply domain ownership, production
service availability, telemetry, or hosted infrastructure before those things
are deliberately created.

## Project-To-Platform Feedback

Arcadia Runtime, and the current RetroFlag Power prototype beneath it, should
eventually be able to produce local findings and feature needs that Lantern
Dispatch can collect with user consent.

Examples include:

- requested CLI features
- missing diagnostic fields
- hardware compatibility findings
- update-channel needs
- support bundle gaps

The first version of this feedback trail should stay local-first. A user should
be able to inspect the satchel before anything leaves the device.

## Privacy And Trust Principles

Trust is part of the runtime. The project should keep these promises as the
field-kit trail grows:

- No telemetry by default.
- Local diagnostics first.
- Redacted support bundles.
- Explicit user consent before any network submission.
- No ROM names.
- No Wi-Fi SSIDs.
- No private IPs unless explicitly included by the user.
- No usernames or home paths without redaction.

Future diagnostics should favor useful, plain-language ledgers over secretive
collection. A support bundle should feel like a field note the user can read,
not a sealed box.

## Rename Caution

RetroFlag Power may need a rename before public v1.0 if it expands beyond
RetroFlag hardware or risks looking affiliated with RetroFlag.

No rename is happening now.

Until the project deliberately chooses a public name, maintainers should:

- preserve or add an unofficial community-project disclaimer where appropriate
- avoid claims of RetroFlag affiliation or endorsement
- keep code package names boring and clear
- keep public-facing names reviewable before release
- treat Arcadia Runtime as the favored direction, not a cleared brand

The current compass is simple: build carefully under RetroFlag Power, map the
field honestly, and let Save Room Tech and Arcadia Runtime remain the larger
trail markers until the project is ready for them.
