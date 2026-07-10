---
id: QUEST-0095
title: Record the First Relic Welcome Scroll Preview Run
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the first real manually invoked GPi Case 2 Relic Welcome Scroll preview run without wiring it into login, boot, services, GPIO, display, shutdown, sleep, or resume behavior.
related:
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-preview-notes.md
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - ../../../scripts/gpi-case2-relic-welcome-scroll.sh
  - 0093-design-relic-welcome-scroll.md
  - 0094-add-relic-welcome-scroll-preview-script.md
last_updated: 2026-07-10
---

# QUEST-0095 - Record the First Relic Welcome Scroll Preview Run

> The first Welcome Scroll was opened by hand on the Relic: art, plain text,
> no-color fallback, and a clear note that the doorway is not wired yet.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Evidence Intake

## Intent

Record the first real GPi Case 2 Relic Welcome Scroll preview run using the
manually invoked script added in QUEST-0094 and the operator's observations.
This is evidence intake only. It must not install or wire the script into SSH
login, MOTD, shell startup, services, boot config, GPIO behavior, display
behavior, shutdown behavior, sleep behavior, or resume behavior.

The current field practice remains scp-first: copy the script to
`retropi@gpi:/home/retropi/`, run it from `/home/retropi/`, and copy artifacts
back only if needed. The GPi Case 2 does not currently have the repository
checked out on the device. SSH is optional support, not the primary handheld
UX.

## Outcome

- Added
  [GPi Case 2 Relic Welcome Scroll Preview Notes](../../03-operations/gpi-case-2-relic-welcome-scroll-preview-notes.md).
- Recorded the first manual scp-first preview run from
  `2026-07-09 17:49:59 PDT`.
- Captured the commands used for copy, chmod, normal preview, `--plain`, and
  `NO_COLOR`.
- Recorded that normal ASCII art output displayed correctly.
- Recorded that `--plain` output worked.
- Recorded that `NO_COLOR=1` output worked.
- Recorded that the preview felt fast, with each remote run returning in less
  than a second from the operator's SSH command perspective.
- Recorded observed fields, missing-field status, transport noise, scp-first
  cleanliness, operator notes, evidence status, candidate improvements, and
  the next rune.
- Linked the preview notes from the Relic Welcome Scroll design, First Spark /
  Boot Veil / Welcome Scroll design, and Local Diagnostics Bundle Map.

## Boundary

- No script behavior change.
- No GPi Case 2 runtime behavior change.
- No SSH login, MOTD, shell startup, PAM, SSHD, profile, service, systemd,
  boot config, display config, GPIO, shutdown, sleep, or resume wiring.
- No attached-keyboard assumption.
- No repository checkout assumption on the GPi Case 2.
- No telemetry, upload, repair, installer, package, or Lantern Dispatch path.
- Do not claim automatic login safety from a manual preview run.

## Acceptance Checks

- [x] A Relic Welcome Scroll preview/evidence notes document exists under
  `docs/03-operations/`.
- [x] The first real preview run records run date/time where available.
- [x] The notes record the command sequence used.
- [x] The notes record whether normal/art output displayed correctly.
- [x] The notes record whether `--plain` output worked.
- [x] The notes record whether `NO_COLOR` output worked.
- [x] The notes record whether output felt fast.
- [x] The notes record whether fields were missing or unavailable.
- [x] The notes record whether any output looked too noisy for SSH login.
- [x] The notes record whether the scp-first flow stayed clean.
- [x] The notes include operator notes, evidence status, and next rune.
- [x] The notes include candidate improvements.
- [x] The notes clearly say the script is not installed into login yet.
- [x] The Relic Welcome Scroll design links to the preview notes.
- [x] The First Spark / Boot Veil / Welcome Scroll design links to the preview
  notes.
- [x] The Local Diagnostics Bundle Map links to the preview notes.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds the preview evidence notes,
  links the design/map docs, and adds this quest ledger. The final pushed
  commit hash is reported in the quest handoff because a Git commit cannot
  contain its own final object hash.

## Final Notes

The first Welcome Scroll preview is in the Ledger. The Relic showed the art
path, plain path, and no-color path quickly over the scp-first trail, while
the script remains a manual satchel tool and not a login, MOTD, service, boot,
GPIO, display, shutdown, sleep, or resume change.
