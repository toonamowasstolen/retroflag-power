---
id: QUEST-0087
title: Design the True GPi Case 2 Boot Trace Lantern
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Design the true GPi Case 2 Boot Trace Lantern as a read-only, scp-first field tool for startup evidence without implementing the full script.
related:
  - ../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-session-watch-field-run-procedure.md
  - ../../03-operations/human-facing-field-lantern-script-ux-standard.md
last_updated: 2026-07-09
---

# QUEST-0087 - Design the True GPi Case 2 Boot Trace Lantern

> Give startup its own Lantern: read-only, scp-first, honest about what it saw,
> and careful not to pretend the boot trail is the whole journey.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Design

## Intent

Create the design for a true GPi Case 2 Boot Trace Lantern that captures
startup evidence without changing boot config, services, GPIO, power behavior,
display behavior, shutdown behavior, sleep, or resume.

The design preserves current field practice:

- Copy the future field script to `retropi@gpi:/home/retropi/`.
- Run it from `/home/retropi/`.
- Retrieve the final Boot Trace Ledger with `scp`.
- Do not require a repository checkout on the handheld Relic.
- Treat SSH as optional support, not the primary handheld UX.
- Keep the side switch as the normal stock shutdown path while responsive.
- Keep the top sleep/resume button suspect unless a procedure explicitly says
  otherwise.

The True Boot Trace Lantern complements the Session Watch Lantern. It is meant
to help separate boot, frontend-start, display-handoff, idle/sleep,
shutdown/resume-edge, and later runtime questions without overclaiming one
artifact.

## Outcome

- Added the canonical design:
  [GPi Case 2 True Boot Trace Lantern Design](../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md).
- Defined purpose, non-goals, safety boundaries, scp-first field flow,
  artifact pattern, progress output, bounded duration, minimum and optional
  signal lists, Artifact Summary expectations, failure modes, and
  do-not-overclaim guidance.
- Included known GPi Case 2 warning hints such as `gpio12`, `vc4`, `v3d`,
  `audio`, `xpad`, `RCU`, display, USB, MMC, ext4, voltage, and throttling
  clues.
- Added relationship sections for Session Watch Lantern and future Arcadia
  Runtime Casters.
- Added a comparison table covering Bundle Collector Lantern, Boot Power Trace
  Lantern, Session Watch Lantern, and True Boot Trace Lantern.
- Linked the design from the project README and existing lantern maps.

## Boundary

- No full script implementation.
- No placeholder script needed for this quest.
- No GPi Case 2 runtime behavior change.
- No power, display, GPIO, config, service, shutdown, sleep, or resume
  behavior change.
- No GPIO reads or writes.
- No installer, firmware, RetroPie, emulator, SafeShutdown, or systemd
  mutation.
- No Lantern Dispatch, telemetry, automatic upload, or automatic repair path.
- No assumption that the repository exists on the GPi Case 2.

## Acceptance Checks

- [x] Design document added under `docs/03-operations/`.
- [x] Design defines purpose.
- [x] Design defines non-goals.
- [x] Design defines safety boundaries.
- [x] Design defines scp-first field flow.
- [x] Design defines expected artifact filename and location pattern.
- [x] Design defines progress and status output expectations.
- [x] Design recommends a bounded capture duration.
- [x] Design lists minimum captured signals.
- [x] Design lists optional captured signals.
- [x] Design defines final Artifact Summary expectations.
- [x] Design records failure modes and artifact behavior.
- [x] Design explains relationship to Session Watch Lantern.
- [x] Design explains relationship to future Arcadia Runtime Casters.
- [x] Design includes known GPi Case 2 warning hints.
- [x] Design includes a do-not-overclaim section.
- [x] Design includes the required four-lantern comparison table.
- [x] Relevant README/map/docs link to the design.
- [x] QUEST-0087 records intent, acceptance checks, validation, and final
  notes.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds the True Boot Trace Lantern
  design and threads it into the project maps. The final pushed commit hash is
  reported in the quest handoff because a Git commit cannot contain its own
  final object hash.

## Final Notes

The Spellbook now has a startup-specific Lantern design that is honest about
its reach. It captures the boot trail, names what went missing, keeps Session
Watch as the later-runtime companion, and leaves all behavior-changing runes
for future quests with their own safety gates.
