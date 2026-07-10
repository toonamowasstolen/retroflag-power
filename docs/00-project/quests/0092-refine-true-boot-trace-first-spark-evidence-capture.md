---
id: QUEST-0092
title: Refine True Boot Trace First Spark Evidence Capture
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Refine the True Boot Trace Lantern evidence model and field documentation after the first real QUEST-0091 run, without changing GPi Case 2 runtime behavior.
related:
  - ../../03-operations/gpi-case-2-true-boot-trace-evidence-ledger.md
  - ../../03-operations/gpi-case-2-true-boot-trace-field-run-procedure.md
  - ../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../03-operations/artifacts/true-boot-trace/gpi-case2-true-boot-trace-lantern-20260710-081529.txt
  - ../../../scripts/gpi-case2-true-boot-trace-lantern.sh
  - 0087-design-true-gpi-case-2-boot-trace-lantern.md
  - 0088-add-true-gpi-case-2-boot-trace-lantern-script-skeleton.md
  - 0089-design-gpi-case-2-first-spark-and-boot-veil-ux.md
  - 0090-add-true-boot-trace-field-run-procedure.md
  - 0091-record-first-true-gpi-case-2-boot-trace-field-run.md
last_updated: 2026-07-10
---

# QUEST-0092 - Refine True Boot Trace First Spark Evidence Capture

> The first Boot Trace Ledger came home with real clues and honest gaps. This
> quest sharpens the Lantern map so the next First Spark run records the human
> moment and the script moment side by side.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Read-Only Script Refinement

## Intent

Use the first real True Boot Trace field run from QUEST-0091 to refine the
Boot Trace Lantern evidence model for first-spark timing. The refinement keeps
the current scp-first field practice: copy the script to
`retropi@gpi:/home/retropi/`, run it from `/home/retropi/`, then copy the
final Ledger artifact back. The GPi Case 2 does not need the repository
checked out on the handheld.

The first returned artifact is:
[`docs/03-operations/artifacts/true-boot-trace/gpi-case2-true-boot-trace-lantern-20260710-081529.txt`](../../03-operations/artifacts/true-boot-trace/gpi-case2-true-boot-trace-lantern-20260710-081529.txt).

## Boundary

- No Boot Veil implementation.
- No splash, Plymouth, framebuffer, KMS, firmware, display config, or boot
  config change.
- No GPIO read or write.
- No service, shutdown, sleep, resume, side-switch, top-button, RetroPie, or
  EmulationStation runtime behavior change.
- No Lantern Dispatch, telemetry, upload, repair, installer, or package work.
- SSH remains optional support, not the primary handheld UX.
- The side switch remains the normal stock shutdown path while responsive.
- The top sleep/resume button remains suspect unless explicitly tested later.

## Outcome

- Reviewed the QUEST-0091 artifact and evidence Ledger.
- Added first-run lessons to the True Boot Trace Evidence Ledger.
- Identified missing or ambiguous fields from the first run: exact
  side-switch time, exact first visible screen time, first SSH availability,
  LED state, power source, handheld/docked state, visible frontend state,
  side-switch shutdown outcome, and top-button status.
- Updated the read-only script's final `Artifact Summary` with first display
  hint sample, first systemd timing sample, first journal hint sample, first
  dmesg hint sample, and the explicit
  `first_visible_screen_note: human_observed_not_script_observed`.
- Kept raw evidence rows backward-compatible.
- Updated the True Boot Trace field-run procedure so human First Spark notes
  align with the artifact summary fields.
- Updated the True Boot Trace design and First Spark / Boot Veil / Relic
  Welcome Scroll design with evidence-driven notes from QUEST-0091 without
  claiming KMS, framebuffer, firmware, power, or Boot Veil causes.

## Acceptance Checks

- [x] QUEST-0091 artifact reviewed.
- [x] True Boot Trace Evidence Ledger includes concise lessons from the first
  run.
- [x] Missing, ambiguous, or hard-to-interpret first-run fields are called out.
- [x] Script changes remain read-only and affect only evidence capture and
  final artifact summary clarity.
- [x] Raw sample rows remain backward-compatible.
- [x] Field-run procedure keeps scp-first `/home/retropi/` instructions and
  aligns human First Spark notes with artifact fields.
- [x] First Spark / Boot Veil / Relic Welcome Scroll design gets only
  evidence-driven notes, not implementation claims.
- [x] `--plain` and `NO_COLOR` remain respected by script checks.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that refines the read-only summary
  fields, updates the evidence and design docs, and adds this quest file. The
  final pushed commit hash is reported in the quest handoff because a Git
  commit cannot contain its own final object hash.

## Final Notes

QUEST-0092 keeps the First Spark mystery evidence-led. The next Ledger should
make the human side-switch and first-visible moments explicit while the script
reports first script-observed display, systemd, journal, dmesg, frontend,
warning, missing-evidence, and raw throttled clues from the shell side.
