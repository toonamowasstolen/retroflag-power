---
id: QUEST-0090
title: Add the True Boot Trace Field Run Procedure
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add the handheld-first, scp-first field procedure for running the True GPi Case 2 Boot Trace Lantern on the Relic and retrieving its final Ledger artifact.
related:
  - ../../03-operations/gpi-case-2-true-boot-trace-field-run-procedure.md
  - ../../../scripts/gpi-case2-true-boot-trace-lantern.sh
  - ../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
last_updated: 2026-07-09
---

# QUEST-0090 - Add the True Boot Trace Field Run Procedure

> Give the True Boot Trace Lantern a field trail: copy the Relic by hand, watch
> the first spark with human eyes, and bring home one honest Ledger.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Procedure

## Intent

Create the handheld-first, scp-first run procedure for the True GPi Case 2 Boot
Trace Lantern. The procedure must match current field practice:

- Copy `scripts/gpi-case2-true-boot-trace-lantern.sh` to
  `retropi@gpi:/home/retropi/`.
- Run from `/home/retropi/`.
- Retrieve the final Boot Trace Ledger with `scp`.
- Do not assume the repository is checked out on the GPi Case 2.
- Treat SSH as optional support, not the primary handheld UX.
- Preserve the side switch as the normal stock shutdown path while responsive.
- Keep the top sleep/resume button suspect unless a dedicated procedure says
  otherwise.
- Record the roughly 15-second quiet-window observation without claiming KMS,
  framebuffer, firmware, userspace, panel, or case-board cause.

## Outcome

- Added
  [GPi Case 2 True Boot Trace Field Run Procedure](../../03-operations/gpi-case-2-true-boot-trace-field-run-procedure.md).
- Included purpose, when to run it, safety boundaries, prerequisites, scp
  copy, `chmod`, smoke run, normal run, `--plain`, `NO_COLOR`, explicit output,
  and Ledger retrieval examples.
- Added handheld observation guidance for first visible screen time, display
  state, LED state, SSH availability, EmulationStation visibility, side-switch
  behavior, and top-button avoidance.
- Added a dedicated First Spark observation section for power-switch timestamp,
  first visible display timestamp or rough seconds, first SSH availability,
  first EmulationStation visibility, and quiet-window notes.
- Linked the field procedure from the True Boot Trace design, First Spark /
  Boot Veil / Relic Welcome Scroll design, Boot Power Trace Lantern map, and
  README.

## Boundary

- Documentation/procedure only.
- No script changes.
- No boot config changes.
- No service, systemd, or automatic boot activation.
- No GPIO reads or writes.
- No power, display, shutdown, sleep, resume, RetroPie, or EmulationStation
  behavior changes.
- No installer, firmware, package, RetroFlag script, or repair execution.
- No Lantern Dispatch, telemetry, network submission, automatic upload, or
  automatic fix.

## Acceptance Checks

- [x] Field-run procedure added under `docs/03-operations/`.
- [x] Procedure names the current True Boot Trace Lantern script.
- [x] Procedure is handheld-first and does not assume an attached keyboard.
- [x] Procedure is scp-first and runs from `/home/retropi/`.
- [x] Procedure includes copy, `chmod`, smoke run, normal run, `--plain`,
  `NO_COLOR`, and retrieval commands.
- [x] Procedure explains what the human should physically observe during boot.
- [x] Procedure records first visible screen timing and display state choices.
- [x] Procedure records LED state, SSH availability, EmulationStation
  visibility, and side-switch behavior.
- [x] Procedure includes a First Spark observation section.
- [x] Procedure preserves the roughly 15-second quiet-window observation
  without overclaiming cause.
- [x] Procedure states what not to do during the run.
- [x] True Boot Trace design links to the procedure.
- [x] First Spark / Boot Veil / Relic Welcome Scroll design links to the
  procedure.
- [x] Boot Power Trace Lantern map links to the procedure.
- [x] README links to the procedure.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds the True Boot Trace field
  procedure and threads it into the project maps. The final pushed commit hash
  is reported in the quest handoff because a Git commit cannot contain its own
  final object hash.

## Final Notes

The True Boot Trace Lantern now has a real field trail for the handheld Relic:
scp the script, watch the physical startup, record the First Spark timing, run
one bounded read-only trace, and bring home the final Ledger without changing
the device.
