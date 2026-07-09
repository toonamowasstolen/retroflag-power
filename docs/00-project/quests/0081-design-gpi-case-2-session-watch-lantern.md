---
id: QUEST-0081
title: Design the GPi Case 2 Session Watch Lantern
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a documentation-only design for a future read-only GPi Case 2 Session Watch Lantern that observes handheld runtime sessions without changing device behavior.
related:
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-session-watch-field-run-procedure.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
  - ../../03-operations/human-facing-field-lantern-script-ux-standard.md
  - ../../03-operations/common-problems-mage-map.md
  - ../../03-operations/local-diagnostics-bundle-map.md
last_updated: 2026-07-09
---

# QUEST-0081 - Design the GPi Case 2 Session Watch Lantern

> Give the next Lantern a map before anyone lights it: a handheld-first watch
> for menu, play, idle-risk, and post-resume evidence, with one clean satchel
> at the end of the trail.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Intent

Create a design-only operations document for the future GPi Case 2 Session
Watch Lantern. The design must keep the GPi Case 2 as a handheld Relic, avoid
assuming an attached keyboard, treat SSH to `retropi@gpi` as optional support,
and preserve the current Lantern safety trail from QUEST-0076 through
QUEST-0080.

## Outcome

- Added
  [GPi Case 2 Session Watch Lantern Design](../../03-operations/gpi-case-2-session-watch-lantern-design.md).
- Defined purpose, non-goals, field questions, safety boundaries, handheld
  procedure, artifact naming, progress/status output, minimum signals,
  optional SSH-assisted signals, failure-mode recording, and relationships to
  Boot Trace Lantern and future Arcadia Runtime Casters.
- Linked the design from the Bundle Collector procedure, Boot Power Trace
  Lantern map, Field Lantern procedure, local diagnostics map, Common Problems
  Mage map, and README project documentation list.
- Kept the quest documentation-only: no runtime code, no scripts, no GPIO, no
  shutdown behavior, no systemd, no installer, no Lantern Dispatch, and no
  hardware behavior changes.

## Acceptance Checks

- [x] A Session Watch Lantern design exists under `docs/03-operations/`.
- [x] The design is read-only and documentation-only.
- [x] The design explicitly avoids assuming an attached keyboard.
- [x] SSH target `retropi@gpi` is optional support, not the primary handheld
  experience.
- [x] The side switch remains the normal stock shutdown path while responsive.
- [x] The top sleep/resume button remains suspect unless a procedure explicitly
  says otherwise.
- [x] Artifact output is one final local file, not a long terminal paste.
- [x] Long-running output expectations include progress, timing, status, and
  final artifact path.
- [x] `--plain` and `NO_COLOR` are required for future implementation.
- [x] Links from relevant maps and index docs are updated.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Completion Commit

- This quest is completed by the commit that adds this file. The final pushed
  commit hash is reported in the quest handoff because a Git commit cannot
  contain its own final object hash.

## Final Notes

The Session Watch Lantern now has a Spellbook page before it has a script. The
next implementation quest can stay tiny and honest: a read-only handheld
watch, clear progress, and one final satchel that says what it saw and what it
missed.

Follow-on field use is documented in
[QUEST-0083](0083-add-gpi-case-2-session-watch-field-run-procedure.md) and the
[GPi Case 2 Session Watch Field Run Procedure](../../03-operations/gpi-case-2-session-watch-field-run-procedure.md).
