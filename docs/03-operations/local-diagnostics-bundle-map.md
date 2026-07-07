---
id: OPS-LOCAL-DIAGNOSTICS-BUNDLE-MAP-001
title: Local Diagnostics Bundle Map
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map a future local-first diagnostics bundle for safe, redacted RetroFlag Power and Arcadia Runtime support information without submitting anything over the network.
related:
  - README.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/00-project/quests/0051-map-the-local-diagnostics-bundle.md
last_updated: 2026-07-07
---

# Local Diagnostics Bundle Map

> This is the privacy compass for the future support satchel: collect useful
> field notes locally, redact them clearly, preview them before sharing, and
> keep the user holding the lantern.

This document is documentation only. It does not implement diagnostics,
telemetry, network submission, file collection, service install, systemd
activation, `rc.local` changes, `SafeShutdown.py` replacement, configuration
mutation, daemon activation, or Lantern Dispatch.

RetroFlag Power remains the current GPi Case 2 prototype. Arcadia Runtime
remains the favored future runtime direction, not an active rename. Lantern
Dispatch remains future and optional.

## Purpose

A future diagnostics bundle should help users and maintainers debug GPi Case,
RetroFlag Power, and nearby retro hardware behavior without collecting private
or unnecessary information.

The bundle should feel like a readable field-kit ledger, not a sealed box. Its
job is to gather enough local context to explain power, display, dock, audio,
startup, and dry-run behavior while preserving user privacy and keeping
troubleshooting possible offline.

## Local-First Rule

Diagnostics should follow this rule before any support submission exists:

- Generate locally first.
- Show or summarize the contents before sharing.
- Never upload automatically.
- Never submit in the background.
- Keep telemetry off by default.

A local bundle can later become the input to a support trail, but it must be
useful on its own. A user should be able to save it, inspect it, edit it, and
share it manually without Lantern Dispatch.

## Future Command Shape

Possible command shapes, all proposed and not implemented:

```sh
retroflag-powerd diagnostics --local
retroflag-powerd diagnostics --bundle
save-room dispatch diagnostics retroflag-power --redact
```

These examples are only a compass for future design. They do not add CLI
flags, binaries, dispatch behavior, file writers, upload paths, or service
behavior.

## Candidate Bundle Contents

A future local diagnostics bundle may include allowlisted sections such as:

- `retroflag-powerd` version.
- OS and kernel summary.
- Raspberry Pi model, when available.
- GPi Case or Relic profile, when known.
- Current display stack hints, such as KMS or FKMS config lines.
- `SafeShutdown.py` presence and path summary.
- `rc.local` or systemd startup references.
- Recent daemon event breadcrumbs.
- Dry-run configuration.
- Raw GPIO probe observations, only when the user includes them.
- Docked and handheld display test notes.
- Audio test notes.
- Relevant project configuration summary.

The bundle should prefer narrow, explainable fields over broad dumps. Each
section should say why it exists, whether it was detected automatically or
entered by the user, and whether any redaction was applied.

## Redaction Rules

The diagnostics charm should redact by default and avoid collecting private
context that maintainers do not need:

- No ROM names.
- No Wi-Fi SSIDs.
- No usernames or home paths without redaction.
- No private IPs unless the user explicitly includes them.
- No tokens or secrets.
- No full environment dumps.
- No arbitrary log scraping without an allowlist.

Redaction should be visible in the preview. A user should be able to see where
the bundle replaced private values with placeholders before deciding whether to
share it.

## User Control

The user should control the diagnostics trail:

- Preview before submit.
- Save to a local file.
- Remove sections before sharing.
- Require explicit consent for any submission.
- Keep an offline support path possible.

The safest first version is a local ledger that can be pasted into an issue,
attached manually, or kept for the user's own notes. Network submission is a
separate future decision, not a requirement for troubleshooting.

## Relationship To Lantern Dispatch

Lantern Dispatch may later receive diagnostics, update checks, feature
requests, support reports, and compatibility findings. That future layer should
remain optional and consent-driven.

This quest only maps the local bundle. It does not implement Lantern Dispatch,
create a hosted service, add update checks, send support reports, or make
network calls.

## Relationship To Field-Kit Feature Requests

Diagnostics may reveal missing CLI or tooling features needed by a field kit.
Examples include:

- Missing display probe.
- Missing audio probe.
- Missing dock-state probe.
- Missing SafeShutdown spellbook detection.
- Missing installer dry-run checks.

Those findings should become explicit future quests. The diagnostics bundle
should record the gap without pretending the missing probe or installer check
already exists.

## Explicit No-Go List

This diagnostics map does not permit:

- Do not upload automatically.
- Do not collect ROM or library data.
- Do not collect secrets.
- Do not implement telemetry.
- Do not make network calls.
- Do not require Lantern Dispatch for local troubleshooting.
- Do not scrape arbitrary logs outside an allowlist.
- Do not activate a daemon, install a service, change `rc.local`, replace
  `SafeShutdown.py`, or mutate configuration.

The bundle should be a user-readable satchel first: local, redacted,
previewable, and useful before any future dispatch trail exists.
