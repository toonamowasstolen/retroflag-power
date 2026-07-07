---
id: QUEST-0051
title: Map the Local Diagnostics Bundle
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a documentation-only local diagnostics bundle map for safe, redacted RetroFlag Power and Arcadia Runtime support information without network submission.
related:
  - docs/03-operations/local-diagnostics-bundle-map.md
  - README.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
last_updated: 2026-07-07
---

# QUEST-0051 - Map the Local Diagnostics Bundle

> Set a privacy lantern beside the future support trail: the diagnostics
> satchel must be local, redacted, previewable, and user-controlled before any
> dispatch path exists.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the local diagnostics bundle map at
  [docs/03-operations/local-diagnostics-bundle-map.md](../../03-operations/local-diagnostics-bundle-map.md).
- Framed diagnostics as a local-first field-kit ledger for debugging GPi Case,
  RetroFlag Power, and future Arcadia Runtime behavior without collecting
  private or unnecessary information.
- Recorded the local-first rule: generate locally, preview or summarize before
  sharing, never upload automatically, never submit in the background, and keep
  telemetry off by default.
- Sketched possible future command shapes as proposed examples only, including
  `retroflag-powerd diagnostics --local`,
  `retroflag-powerd diagnostics --bundle`, and
  `save-room dispatch diagnostics retroflag-power --redact`.
- Listed candidate allowlisted bundle sections for version, OS/kernel,
  Raspberry Pi model, Relic or GPi Case profile, display stack hints,
  `SafeShutdown.py` presence, startup references, event breadcrumbs, dry-run
  configuration, user-included raw GPIO observations, display notes, audio
  notes, and relevant project configuration.
- Added redaction rules for ROM names, Wi-Fi SSIDs, usernames, home paths,
  private IPs, tokens, secrets, full environment dumps, and arbitrary log
  scraping.
- Captured user control requirements: preview before submit, save locally,
  remove sections, require explicit consent, and preserve an offline support
  path.
- Explained that Lantern Dispatch may later receive diagnostics, update
  checks, feature requests, support reports, and compatibility findings, but
  this quest only maps the local bundle.
- Connected diagnostics to future field-kit feature requests such as missing
  display, audio, dock-state, SafeShutdown spellbook, and installer dry-run
  checks.
- Added an explicit no-go list for automatic uploads, ROM/library collection,
  secret collection, telemetry, network calls, and requiring Lantern Dispatch
  for local troubleshooting.
- Linked the map from `README.md`,
  `docs/04-architecture/arcadia-runtime-migration-path.md`,
  `docs/00-project/project-direction-save-room-arcadia.md`, and
  `docs/03-operations/safeshutdown-replacement-boundary-map.md`.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No Go code changes.
- No diagnostics implementation.
- No telemetry implementation.
- No network calls.
- No file collection implementation.
- No service install.
- No systemd activation.
- No `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No config mutation.
- No daemon activation.

## Milestone Note

The project now has a privacy-first diagnostics compass. Before Lantern
Dispatch or any support submission path exists, maintainers can see what a
local bundle should collect, redact, preview, and leave under user control.
