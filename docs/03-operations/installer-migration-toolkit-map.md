---
id: OPS-INSTALLER-MIGRATION-TOOLKIT-MAP-001
title: Installer and Migration Toolkit Map
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map a future local-first, reversible installer and migration toolkit for safely inspecting, backing up, planning, applying, uninstalling, and restoring RetroFlag Power and future Arcadia Runtime field kits.
related:
  - README.md
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/00-project/quests/0052-map-the-installer-and-migration-toolkit.md
last_updated: 2026-07-07
---

# Installer and Migration Toolkit Map

> This is the installer compass for the future field kit: inspect the relic,
> pack the backup satchel, preview the route, and leave a restore ledger before
> any spell changes the machine.

This document is documentation only. It does not implement an installer,
migration command, shell installer, RetroPie scriptmodule, telemetry, network
call, service install, systemd activation, `rc.local` change, file mutation,
config mutation, `SafeShutdown.py` replacement, GPIO behavior change, GPIO
write, shutdown execution, or daemon activation.

RetroFlag Power remains the current GPi Case 2 prototype. Arcadia Runtime
remains the favored future runtime direction, not an active rename. Save Room
Tech remains the future umbrella. Field kits are future installable and
supportable packages or device-specific bundles. Lantern Dispatch remains
future and optional.

## Purpose

A future installer and migration toolkit should help users safely move from
legacy RetroFlag scripts and older Raspberry Pi OS assumptions toward
RetroFlag Power behavior, and later toward future Arcadia Runtime field kits.

The toolkit's first job is not to be clever. Its first job is to be legible:
show what is present, what looks risky, what would be backed up, what would be
changed, and exactly how the user can restore the previous state.

## Local-First And Reversible Rules

The toolkit should keep the user holding the lantern:

- Inspect before changing.
- Back up before replacing.
- Preview proposed changes.
- Require explicit confirmation before applying changes.
- Produce restore instructions.
- Support uninstall and restore.
- Never require network access for a basic local install or rollback.

A field kit should be useful from local files on the device. Update checks,
support submission, and Lantern Dispatch are separate future choices, not
requirements for basic install, uninstall, or restore.

## Future Command Shape

Possible command shapes, all proposed and not implemented:

```sh
retroflag-powerd installer inspect
retroflag-powerd installer plan
retroflag-powerd installer backup
retroflag-powerd installer apply
retroflag-powerd installer restore
save-room fieldkit install arcadia-gpi-case-2
```

These examples are only a map for future design. They do not add commands,
flags, file writers, package installers, service behavior, GPIO behavior,
shutdown behavior, or network behavior.

## Legacy Detection Checklist

A future installer inspect step should look for local evidence such as:

- `/opt/RetroFlag/SafeShutdown.py` presence and checksum or hash, when
  available.
- Related `/opt/RetroFlag` scripts.
- `rc.local` startup entries.
- Existing systemd units that mention RetroFlag Power, RetroFlag scripts, or
  related power behavior.
- `config.txt` display lines.
- FKMS or KMS overlay hints.
- `cmdline.txt` hints.
- RetroPie or EmulationStation presence.
- GPi Case 2 docking and display notes.
- Audio config notes.
- Current `retroflag-powerd` binary or config, if present.

Detection should record uncertainty clearly. A missing path, unreadable file,
unknown checksum, or ambiguous display stack should become a warning in the
ledger, not a silent assumption.

## Backup Plan

Before any future apply step replaces or edits anything, the toolkit should:

- Copy legacy scripts before changes.
- Record file paths.
- Record checksums.
- Record timestamps.
- Record original config snippets.
- Store a human-readable restore ledger.
- Avoid collecting private ROM or library data.

The backup satchel should be narrow and explainable. It should preserve the
files and snippets needed for rollback without sweeping up personal libraries,
ROM names, scraped logs, or unrelated home-directory content.

## Plan And Apply Model

The toolkit should generate a plan before applying anything.

The plan should show:

- Files to back up.
- Files to edit.
- Services to add or disable.
- Config lines to add or remove.
- Exact restore path.
- Risk notes.
- Whether reboot is needed.

The apply step should consume a visible plan and require explicit confirmation.
If a plan cannot prove its restore path, it should stop and ask for a better
ledger instead of pressing forward.

## SafeShutdown.py Replacement Gates

The `SafeShutdown.py` replacement boundary lives in
[SafeShutdown Replacement Boundary Map](safeshutdown-replacement-boundary-map.md).
That map is a prerequisite ledger, not permission to replace the stock script.

Replacement cannot happen until the project understands and tests:

- Power-enable latch behavior.
- Side-switch shutdown behavior.
- Top-button power-save and resume behavior.
- Handheld LCD behavior.
- Docked HDMI behavior.
- KMS timing.
- Audio behavior.
- Rollback behavior.

Until those gates are satisfied, a future installer may inspect and plan around
the stock script, but it must not silently disable, move, overwrite, or replace
`/opt/RetroFlag/SafeShutdown.py`.

## RetroPie Integration Path

The conservative RetroPie trail should start outside RetroPie itself:

- Start with the project's own installer.
- Later draft a RetroPie-Setup-compatible scriptmodule.
- Test externally and community-first.
- Do not imply RetroPie endorsement.
- Only approach RetroPie maintainers after stable behavior, docs, rollback,
  and support workflow exist.

RetroPie integration should be earned through predictable behavior and a clear
support story. The first version of the toolkit should remain local, explicit,
and reversible before asking another project to carry the charm.

## Relationship To Diagnostics

The local diagnostics satchel is mapped in
[Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md).

Future installer inspect and plan output may feed local diagnostics and
field-kit feature requests. Examples include missing display probes, missing
audio checks, unknown `SafeShutdown.py` hashes, or rollback gaps.

No submission path exists yet. This map does not implement Lantern Dispatch,
upload diagnostics, submit feature requests, or make network calls.

## Explicit No-Go List

This installer map does not permit:

- Do not mutate config silently.
- Do not replace `SafeShutdown.py` without backup and restore.
- Do not install services without confirmation.
- Do not require network access.
- Do not collect ROM or library data.
- Do not claim official RetroFlag or RetroPie support.
- Do not skip rollback documentation.
- Do not write GPIO, execute shutdown, activate a daemon, alter `rc.local`, or
  enable systemd units as part of this documentation quest.

The future toolkit should be a reversible compass before it is a wrench: local
inspection first, visible plan second, confirmed apply third, and a readable
restore ledger always.
