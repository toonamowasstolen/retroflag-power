---
id: QUEST-0089
title: Design GPi Case 2 First Spark and Boot Veil UX
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Design a recovery-first GPi Case 2 startup UX plan for earlier power-on feedback, safe boot-text hiding, and an SSH welcome/status scroll without implementing boot config changes.
related:
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-recovery-first-field-procedure.md
last_updated: 2026-07-09
---

# QUEST-0089 - Design GPi Case 2 First Spark and Boot Veil UX

> Give the Relic a quicker sign of life, a safer way to hide rough boot text,
> and an SSH welcome scroll, without taking away the recovery trail.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Design

## Intent

Create a recovery-first startup UX design for the GPi Case 2 that separates:

- First Spark: earliest possible visible or physical feedback after the side
  power switch is flipped.
- Boot Veil: hiding boot text and penguin graphics once the display path is
  active, where safe.
- Relic Welcome Scroll: SSH login banner/status output for `retropi@gpi`.

The design must preserve current field practice and recovery behavior. It must
not implement boot config changes, service changes, GPIO behavior, shutdown
behavior, sleep/resume behavior, or display changes.

## Outcome

- Added the canonical design:
  [GPi Case 2 First Spark / Boot Veil / Relic Welcome Scroll Design](../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md).
- Separated startup UX into First Spark, Boot Veil, and Relic Welcome Scroll.
- Added a startup timeline from side-switch flip through post-boot Session
  Watch.
- Listed candidate approaches: firmware/kernel config, framebuffer splash,
  Plymouth or splash service, LED/status signal, and SSH MOTD/status scroll.
- Included risks and recovery paths for each approach.
- Required True Boot Trace Lantern evidence before any implementation.
- Added do-not-overclaim notes for the roughly 15-second silent-screen
  observation.
- Linked the design from the True Boot Trace design, Session Watch design,
  Boot Power Trace map, README, and hardware notes shelf.

## Boundary

- No implementation of First Spark, Boot Veil, Plymouth, framebuffer splash,
  SSH MOTD, firmware config, kernel command line, systemd unit, RetroPie,
  EmulationStation, or display behavior.
- No edits to `/boot/config.txt`, `/boot/cmdline.txt`, `/etc/rc.local`, or
  `/opt/RetroFlag/SafeShutdown.py`.
- No GPIO reads or writes.
- No shutdown, reboot, halt, poweroff, sleep, suspend, or resume behavior.
- No assumption that the GPi Case 2 has an attached keyboard.
- No reliance on the top sleep/resume button.
- No automatic repair, telemetry, upload, or Lantern Dispatch.

## Acceptance Checks

- [x] Design document added under `docs/03-operations/`.
- [x] Design separates First Spark, Boot Veil, and Relic Welcome Scroll.
- [x] Design includes startup timeline from power switch flip to post-boot
  Session Watch.
- [x] Design lists firmware/kernel config, framebuffer splash, Plymouth or
  splash service, LED/status signal, and SSH MOTD/status scroll candidates.
- [x] Design includes risks and recovery paths for each candidate.
- [x] Design records True Boot Trace evidence needed before implementation.
- [x] Design includes do-not-overclaim notes for the current roughly
  15-second observation.
- [x] Design preserves scp-first field practice and handheld-first assumptions.
- [x] Design preserves side-switch and top-button safety notes.
- [x] Design forbids boot config implementation in this quest.
- [x] True Boot Trace design links to the new startup UX design.
- [x] Session Watch design links to the new startup UX design.
- [x] Relevant README/map docs link to the new startup UX design.
- [x] QUEST-0089 records intent, acceptance checks, validation, and final
  notes.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds the First Spark, Boot Veil,
  and Relic Welcome Scroll design and threads it into the project maps. The
  final pushed commit hash is reported in the quest handoff because a Git
  commit cannot contain its own final object hash.

## Final Notes

The startup UX now has a recovery-first design path. The Relic can pursue an
earlier sign of life, a reversible veil over rough boot text, and a useful SSH
welcome scroll only after the True Boot Trace Lantern proves which window each
layer can safely touch.
