---
id: QUEST-0084
title: Add the GPi Case 2 Session Watch Evidence Ledger
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add the first documentation-only evidence ledger for real GPi Case 2 Session Watch Lantern field runs.
related:
  - ../../03-operations/gpi-case-2-session-watch-evidence-ledger.md
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-session-watch-field-run-procedure.md
  - ../../../scripts/gpi-case2-session-watch-lantern.sh
  - ../../03-operations/common-problems-mage-map.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - 0081-design-gpi-case-2-session-watch-lantern.md
  - 0082-add-gpi-case-2-session-watch-lantern-script-skeleton.md
  - 0083-add-gpi-case-2-session-watch-field-run-procedure.md
last_updated: 2026-07-09
---

# QUEST-0084 - Add the GPi Case 2 Session Watch Evidence Ledger

> Give the Session Watch Lantern a proper Ledger page before the first field
> artifact arrives: one map for the machine trail, human notes, and next rune.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Intent

Create the first evidence-ledger structure for GPi Case 2 Session Watch
Lantern field runs. The Ledger should prepare maintainers to record real
handheld results from
[`scripts/gpi-case2-session-watch-lantern.sh`](../../../scripts/gpi-case2-session-watch-lantern.sh)
without adding behavior-changing runtime code.

The Ledger must keep the GPi Case 2 handheld-first, treat SSH to `retropi@gpi`
as optional support, preserve the side switch as the normal responsive
shutdown path, keep the top sleep/resume button suspect unless a procedure
explicitly says otherwise, and separate observed evidence from guesses.

## Outcome

- Added the
  [GPi Case 2 Session Watch Evidence Ledger](../../03-operations/gpi-case-2-session-watch-evidence-ledger.md).
- Documented purpose, evidence status legend, run-entry template, script
  artifact fields, date, duration, power source, docked/handheld state,
  display state, LED state, SSH status, side-switch behavior, top-button
  behavior, EmulationStation status, final outcome, human notes, and next rune.
- Added interpretation buckets for `clean-run`, `display-blank-ssh-alive`,
  `ssh-lost`, `side-switch-failure`, `top-button-oddity`, and
  `hard-freeze-rcu-stall-like`.
- Added a do-not-overclaim note that separates observed evidence from guesses.
- Added a placeholder first entry marked `awaiting first real field run`
  instead of inventing handheld data.
- Linked the Ledger from the Session Watch design, Session Watch field-run
  procedure, README project map, Common Problems Mage map, and Local
  Diagnostics Bundle map.

## Boundary

- No script behavior change.
- No GPi Case 2 runtime behavior change.
- No GPIO reads or writes.
- No shutdown, reboot, halt, poweroff, suspend, sleep, resume, display,
  config, service, installer, firmware, SafeShutdown, or RetroPie mutation.
- No systemd activation.
- No Lantern Dispatch, telemetry, network submission, automatic upload, or
  automatic diagnostics bundle generation.
- No invented field evidence.

## Acceptance Checks

- [x] A Session Watch evidence Ledger exists under `docs/03-operations/`.
- [x] The Ledger includes a purpose section.
- [x] The Ledger includes an evidence status legend covering untested,
  observed, reproduced, contradicted, and resolved.
- [x] The Ledger includes a run-entry template.
- [x] The template includes script artifact path/name fields.
- [x] The template includes date, duration, power source, docked/handheld
  state, display state, LED state, SSH status, side-switch behavior,
  top-button behavior, EmulationStation status, and final outcome.
- [x] The template includes human notes and next rune fields.
- [x] The Ledger includes interpretation buckets for clean run,
  display blank with SSH alive, SSH lost, side-switch failure, top-button
  oddity, and hard-freeze or RCU-stall-like outcome.
- [x] The Ledger includes a do-not-overclaim note separating observed evidence
  from guesses.
- [x] The first entry is a placeholder marked as awaiting the first real field
  run.
- [x] Relevant Session Watch, diagnostics, Common Problems Mage, and README
  map docs link the Ledger.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds this file. The final pushed
  commit hash is reported in the quest handoff because a Git commit cannot
  contain its own final object hash.

## Final Notes

The Session Watch Lantern now has a Ledger ready for real handheld evidence.
The first entry waits for a true field run: one final script artifact, concise
human observations, a humble interpretation bucket, and the next rune for the
trail.
