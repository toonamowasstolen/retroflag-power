---
id: QUEST-0076
title: Add a Portable GPi Case 2 Boot Power Trace Field Lantern Script
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Add a single portable read-only GPi Case 2 Boot Power Trace Field Lantern script that can be copied to the Pi without a repository checkout.
related:
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../../scripts/gpi-case2-boot-power-trace-field-lantern.sh
last_updated: 2026-07-08
---

# QUEST-0076 - Add a Portable GPi Case 2 Boot Power Trace Field Lantern Script

> Turn the Boot Power Trace from a copied spell into a pocket Lantern Relic:
> one file carried to the handheld, one local bundle brought home.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Diagnostics

## Outcome

- Added the portable read-only script at
  [`scripts/gpi-case2-boot-power-trace-field-lantern.sh`](../../../scripts/gpi-case2-boot-power-trace-field-lantern.sh).
- Updated the
  [Boot Power Trace Capture Procedure](../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md)
  with copy, run, retrieve, and review commands for a Pi without git, Go, a
  repository checkout, or a project install.
- Updated the
  [Boot Power Trace Lantern Map](../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md)
  so the portable Field Lantern script is part of the current path.

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

## Milestone Note

The GPi Case 2 Boot Power Trace now has a portable Field Lantern: copy one
script to the handheld, run it locally, and retrieve one timestamped `.tar.gz`
bundle for the Ledger.
