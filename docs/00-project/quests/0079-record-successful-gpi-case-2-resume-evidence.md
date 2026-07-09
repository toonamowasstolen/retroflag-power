---
id: QUEST-0079
title: Record Successful GPi Case 2 Resume Evidence
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record the completed evidence update that captured a successful GPi Case 2 post-resume Lantern run and refined the intermittent resume-wedge theory.
related:
  - ../../02-hardware/gpi-case-2.md
  - ../../03-hardware/gpi-case-2-emergency-recovery-research-ledger.md
  - ../../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
  - ../../03-operations/gpi-case-2-acceptance-checklist.md
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-recovery-first-field-procedure.md
  - ../../03-operations/gpi-case-2-replacement-coverage-matrix.md
last_updated: 2026-07-09
---

# QUEST-0079 - Record Successful GPi Case 2 Resume Evidence

> Place the good resume satchel beside the scary one: the trail now shows an
> intermittent wedge, not a guaranteed curse.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Hardware Evidence

## Outcome

- Recorded the 2026-07-09 Bundle Collector Field Lantern satchel captured after
  unintended sleep followed by successful resume:
  `gpi-case2-bundle-collector-field-lantern-20260709-083407.tar.gz`.
- Updated GPi Case 2 hardware, power integrity, emergency recovery, recovery
  procedure, acceptance checklist, Lantern map, and replacement coverage docs
  to reflect the successful post-resume evidence.
- Revised the resume theory from "resume wedge is guaranteed" to "resume wedge
  is intermittent" while preserving caution around sleep duration, battery
  state, thermal state, USB/input state, display/KMS timing, and transient
  power conditions.
- Marked the post-resume satchel as after-the-fact evidence, not proof of the
  sleep/resume transition itself, because no watcher was already running before
  sleep.
- Added a future Session Watch Lantern direction for pre-sleep and post-resume
  evidence without telemetry or automatic fixes.

## Boundary

- Documentation and evidence recording only.
- No Go code changes.
- No shell script changes.
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

- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- `114b9903979facbf3e523e0618e0cd955b88b7e6` - added successful GPi Case 2
  resume evidence and revised the intermittent resume-wedge theory.

## Milestone Note

The Ledger now holds both the failed resume trail and the successful
post-resume Lantern satchel. That is a verified win: the next Arcadia Runtime
map can move with sharper uncertainty instead of a single frightening story.
