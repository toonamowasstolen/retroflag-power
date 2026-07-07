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
  - docs/03-operations/installer-migration-toolkit-map.md
  - docs/03-operations/gpi-case-2-acceptance-checklist.md
  - docs/00-project/quests/0051-map-the-local-diagnostics-bundle.md
  - docs/00-project/quests/0058-map-the-local-diagnostics-bundle-skeleton.md
  - docs/00-project/quests/0052-map-the-installer-and-migration-toolkit.md
  - docs/00-project/quests/0053-add-the-gpi-case-2-acceptance-checklist.md
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

The future installer and migration compass is mapped in
[Installer and Migration Toolkit Map](installer-migration-toolkit-map.md). Its
inspect and plan output may eventually feed local diagnostics, but it does not
submit anything, call the network, or mutate a device.

The GPi Case 2 field-test gate is tracked in
[GPi Case 2 Acceptance Checklist](gpi-case-2-acceptance-checklist.md). Future
diagnostics may help fill that ledger, but this map does not collect or submit
those notes yet.

## Purpose

A future diagnostics bundle should help users and maintainers debug GPi Case,
RetroFlag Power, and nearby retro hardware behavior without collecting private
or unnecessary information.

The bundle should feel like a readable field-kit ledger, not a sealed box. Its
job is to gather enough local context to explain power, display, dock, audio,
startup, and dry-run behavior while preserving user privacy and keeping
troubleshooting possible offline.

This map is a skeleton only. It names future bundle sections and boundaries so
the project can agree on the satchel before any command creates one.

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
- Dry-run startup summary.
- Runtime and startup diagnostics summary.
- Execution status summary.
- Recent event breadcrumbs.
- Configured hardware profile name, if available.
- OS, kernel, and Raspberry Pi model summary.
- Current display stack hints, such as KMS or FKMS config lines.
- KMS, FKMS, display connector, and display configuration facts that can be
  collected read-only.
- Audio device facts that can be collected read-only.
- `SafeShutdown.py` presence, process, and status observations, read-only only.
- `rc.local` or systemd startup references, read-only only.
- Raw GPIO probe observations, only when explicitly run and included by the
  user.
- Docked and handheld display test notes.
- Audio test notes.
- Field checklist references for GPi Case 2 testing.
- Relevant project configuration summary.

The bundle should prefer narrow, explainable fields over broad dumps. Each
section should say why it exists, whether it was detected automatically or
entered by the user, and whether any redaction was applied.

## Bundle Skeleton

The first future bundle shape should be small and sectioned. Each section
should be independently previewable and removable before sharing.

| Section | Allowed source | Future contents | Boundary |
| --- | --- | --- | --- |
| Version | Read-only command or build metadata | `retroflag-powerd` version, build label, and command mode. | No update check, network call, or package mutation. |
| Startup | Existing dry-run/app output | Dry-run startup summary, startup diagnostics, and startup result. | No service activation or daemon install. |
| Runtime | Existing diagnostic snapshots | Runtime diagnostics summary and execution status summary. | No persistent state or background collector. |
| Breadcrumbs | Existing in-process event trail | Recent power-intent, dry-run, startup, or diagnostic breadcrumbs. | No broad journal scraping. |
| Hardware profile | Local configuration, if present | Configured hardware profile or Relic name, such as GPi Case 2, when available. | Do not infer `SwitchOn` or `SwitchOff` from raw GPIO alone. |
| GPIO observations | User-supplied probe output | Raw `SignalLow`, `SignalHigh`, or `SignalUnverified` observations copied from explicit user-run probes. | Do not run probes automatically and do not perform GPIO writes. |
| OS and Pi facts | Read-only OS files or commands | OS version, kernel version, Pi model, and relevant CPU/device identity. | No username, home path, network identity, or environment dump. |
| Display facts | Read-only config and display probes | KMS/FKMS status, relevant display config lines, connector names, and handheld/docked notes. | No config rewrite, no `lcdnext.sh`, no `lcdfirst.sh`, and no display mode mutation. |
| Audio facts | Read-only audio inventory | Safe device names and selected audio path notes needed for handheld/docked triage. | No mixer changes and no audio remap changes. |
| SafeShutdown boundary | Read-only file/process/status observations | Whether `/opt/RetroFlag/SafeShutdown.py` exists, whether matching processes are visible, and any startup references. | No process kill, disable, replacement, or `rc.local` edit. |
| Field checklist | EDC references and user notes | Links or references to GPi Case 2 acceptance rows and manual test notes. | No claim that checklist rows passed without user evidence. |

The skeleton should keep the raw-signal vocabulary boundary visible:

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from
  the configured hardware profile.
- A diagnostics bundle may carry both only when the source and interpretation
  are clearly separated.

## Allowed Read-Only Diagnostics

Allowed future diagnostics are lanterns: read-only probes and diagnostics that
describe the local field state without changing it.

Allowed examples include:

- Printing daemon version information.
- Summarizing dry-run startup results.
- Summarizing runtime, startup, and execution diagnostics already available
  from the app.
- Listing event breadcrumbs already recorded by a foreground dry-run or
  diagnostics command.
- Naming the configured hardware profile, when a profile is available.
- Including raw GPIO observations only when the user explicitly ran the probe
  and chose to include the result.
- Reading OS, kernel, Raspberry Pi model, display, KMS/FKMS, and audio facts
  through narrow allowlisted paths.
- Observing `SafeShutdown.py` file presence, process presence, and startup
  references without changing them.
- Referring to GPi Case 2 acceptance checklist rows and field ledger entries.

Lanterns should explain uncertainty. Unknown values should remain `Unknown`
instead of becoming a confident-looking guess.

## Forbidden Active Behavior

The diagnostics bundle must not become a hidden installer, replacement path, or
hardware actor.

Forbidden behavior includes:

- GPIO writes.
- Automatic GPIO probing.
- Shutdown execution.
- Resume or power-save implementation.
- Persistent state.
- Telemetry.
- Network calls.
- Installer or packaging changes.
- systemd activation.
- `rc.local` replacement or edits.
- `SafeShutdown.py` replacement, disablement, process termination, or mutation.
- Running `lcdnext.sh`, `lcdfirst.sh`, or any display-switching script.
- Rewriting KMS, FKMS, display, audio, boot, or runtime configuration.
- Generating a diagnostics bundle before a future implementation quest exists.

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

## Future User-Redacted Contents

A future bundle may allow the user to add optional, explicitly redacted
sections after preview. These should remain opt-in and locally visible:

- Manual GPi Case 2 field-test notes.
- Photos or descriptions of hardware revision markings.
- Selected command output copied by the user.
- Private IP, hostname, or path details only when the user chooses to keep or
  reveal them.
- Longer excerpts from local logs only when chosen from a narrow allowlist and
  previewed before sharing.

The bundle should support removing any optional section. Redaction should be a
user-visible step, not a promise hidden behind a command name.

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

Future Lantern Dispatch ideas, explicitly not implemented here:

- Submit a user-previewed diagnostics bundle to a support issue.
- Attach a redacted bundle to a compatibility report.
- Check whether a known GPi Case 2 field issue has matching local symptoms.
- Offer an update or migration recommendation after local preview and consent.
- Send a support request that includes only sections the user kept.

Lantern Dispatch must remain future optional update, diagnostics,
issue-reporting, and support-submission infrastructure. Local diagnostics must
remain useful without it.

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
- Do not generate a diagnostics bundle yet.
- Do not run raw GPIO probes unless the user explicitly runs them.
- Do not activate a daemon, install a service, change `rc.local`, replace
  `SafeShutdown.py`, or mutate configuration.

The bundle should be a user-readable satchel first: local, redacted,
previewable, and useful before any future dispatch trail exists.
