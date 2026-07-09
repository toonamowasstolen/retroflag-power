---
id: QUEST-0083
title: Add the GPi Case 2 Session Watch Field Run Procedure
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a handheld-first field procedure for running the GPi Case 2 Session Watch Lantern script created in QUEST-0082.
related:
  - ../../03-operations/gpi-case-2-session-watch-field-run-procedure.md
  - ../../../scripts/gpi-case2-session-watch-lantern.sh
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-field-lantern-capture-procedure.md
  - ../../03-operations/common-problems-mage-map.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - 0081-design-gpi-case-2-session-watch-lantern.md
  - 0082-add-gpi-case-2-session-watch-lantern-script-skeleton.md
last_updated: 2026-07-09
---

# QUEST-0083 - Add the GPi Case 2 Session Watch Field Run Procedure

> Put the new Session Watch Lantern into a real handheld ritual: the Relic in
> hand, optional SSH support, one final Ledger, and clear notes for the human
> things only the screen, LED, and buttons can tell us.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Intent

Create a handheld-first field procedure for running the
[`scripts/gpi-case2-session-watch-lantern.sh`](../../../scripts/gpi-case2-session-watch-lantern.sh)
script added in QUEST-0082. The procedure must treat the GPi Case 2 as a
handheld Relic first, with SSH to `retropi@gpi` only as optional support for
copying, launching, and retrieving the final Ledger.

The procedure should tell a human what to run, what to watch physically, what
not to touch during the run, where the final Ledger appears, how to share it,
and how to interpret clean, display-blank, SSH-lost, and hard-freeze-like
runs. It must preserve the current boundaries: read-only script, one final
artifact, side switch as the normal responsive shutdown path, and the top
sleep/resume button as suspect unless a later procedure explicitly says
otherwise.

## Outcome

- Added
  [GPi Case 2 Session Watch Field Run Procedure](../../03-operations/gpi-case-2-session-watch-field-run-procedure.md).
- Documented purpose, when to run, safety boundaries, prerequisites,
  recommended commands, short watch, normal watch, `--plain`, and `NO_COLOR`
  examples.
- Added human observation guidance for display state, LED blink state, SSH
  alive or lost, handheld controls, side-switch behavior, top-button behavior,
  sound, and game state.
- Documented what not to do during the run, expected artifact location,
  retrieval with `scp`, manual sharing, and pre-sharing inspection.
- Added interpretation guidance for clean runs, sleep/display-blank runs,
  SSH-lost runs, and hard-freeze or RCU-stall-like runs.
- Linked the procedure from README, the QUEST-0081 design trail, the
  QUEST-0082 script quest, the Session Watch design, the Field Lantern capture
  procedure, the Common Problems Mage map, and the Local Diagnostics Bundle
  map.

## Boundary

- No script behavior change.
- No GPi Case 2 runtime behavior change.
- No GPIO reads or writes.
- No shutdown, reboot, halt, poweroff, suspend, sleep, resume, display,
  config, service, installer, firmware, SafeShutdown, or RetroPie mutation.
- No systemd activation.
- No Lantern Dispatch, telemetry, network submission, or automatic upload.
- No assumption that the handheld has an attached keyboard.

## Acceptance Checks

- [x] A Session Watch field run procedure exists under `docs/03-operations/`.
- [x] The procedure is handheld-first and does not require an attached
  keyboard.
- [x] SSH target `retropi@gpi` is optional support, not the primary handheld
  UX.
- [x] The procedure keeps the side switch as the normal stock shutdown path
  while responsive.
- [x] The top sleep/resume button remains suspect and is not part of the
  general watch run.
- [x] The procedure describes read-only operation and one final Ledger
  artifact.
- [x] Command examples include short watch, normal watch, `--plain`, and
  `NO_COLOR`.
- [x] Human observations cover display state, LED blink state, SSH alive or
  lost, side-switch behavior, and top-button behavior.
- [x] Artifact location, retrieval, sharing, and interpretation guidance are
  documented.
- [x] Relevant design, script, map, diagnostics, and index docs link the
  procedure.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds this file. The final pushed
  commit hash is reported in the quest handoff because a Git commit cannot
  contain its own final object hash.

## Final Notes

The Session Watch Lantern now has a field-run Spellbook, not just a design and
a script. A maintainer can run a bounded handheld watch, record the physical
state of the Relic, retrieve one final Ledger, and keep clean evidence beside
the future Common Problems Mage trail without stepping into repair behavior.
