---
id: QUEST-0076
title: Add a Portable GPi Case 2 Bundle Collector Field Lantern Script
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a single portable read-only GPi Case 2 Bundle Collector Field Lantern script that can be copied to the Pi without a repository checkout.
related:
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../../scripts/gpi-case2-bundle-collector-field-lantern.sh
last_updated: 2026-07-08
---

# QUEST-0076 - Add a Portable GPi Case 2 Bundle Collector Field Lantern Script

> Turn the post-boot evidence collector into a pocket Lantern Relic: one file
> carried to the handheld, one local bundle brought home.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Diagnostics

## Outcome

- Added the portable read-only script at
  [`scripts/gpi-case2-bundle-collector-field-lantern.sh`](../../../scripts/gpi-case2-bundle-collector-field-lantern.sh).
- Updated the
  [Bundle Collector Lantern Capture Procedure](../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md)
  with copy, run, retrieve, and review commands for a Pi without git, Go, a
  repository checkout, or a project install.
- Updated the
  [Boot Power Trace Lantern Map](../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md)
  so the portable Bundle Collector Lantern script is separate from the future
  boot-time recorder path.

## Boundary

- No GPIO writes.
- No GPIO reads.
- No shutdown or reboot execution.
- No systemd activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume implementation.
- No firmware flashing.
- No installer execution.
- No automatic fixes.
- No telemetry or network contact.
- No hardware modification instructions.

## Validation

- [x] Script smoke test with a one-second local capture passed.
- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- `c8cd9f673a54f7c9e68316f10cf089c7f77eda2d` - added the portable GPi Case 2
  Bundle Collector Field Lantern script and its capture trail markers.

## Milestone Note

The GPi Case 2 now has a portable Bundle Collector Lantern: copy one script to
the handheld, run it locally after boot, and retrieve one timestamped `.tar.gz`
bundle for the Ledger.
