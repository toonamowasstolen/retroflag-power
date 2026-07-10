---
id: QUEST-0093
title: Design the Relic Welcome Scroll
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Design a read-only, fast, recovery-first SSH welcome scroll for GPi Case 2 operator sessions without implementing shell startup, MOTD, boot, service, GPIO, shutdown, sleep, or resume behavior.
related:
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-true-boot-trace-lantern-design.md
  - ../../03-operations/gpi-case-2-session-watch-lantern-design.md
  - ../../03-operations/gpi-case-2-recovery-first-field-procedure.md
  - 0089-design-gpi-case-2-first-spark-and-boot-veil-ux.md
  - 0091-record-first-true-gpi-case-2-boot-trace-field-run.md
  - 0092-refine-true-boot-trace-first-spark-evidence-capture.md
last_updated: 2026-07-10
---

# QUEST-0093 - Design the Relic Welcome Scroll

> Give optional SSH support a warm little map: fast, read-only, scp-safe, and
> easy to silence if recovery needs a bare prompt.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Design

## Intent

Create a dedicated Relic Welcome Scroll design for GPi Case 2 SSH login to
`retropi@gpi`. The scroll is a future read-only, friendly system-info banner
for interactive SSH sessions. It should show useful status without becoming a
diagnostic bundle, installer, service, shell-risk trap, or primary handheld
UX.

This quest is design/documentation only. It does not implement the scroll and
does not modify shell startup files, MOTD files, boot config, services, GPIO,
display, shutdown, sleep, or resume behavior.

## Boundary

- No shell profile, MOTD, PAM, SSHD, service, boot, display, framebuffer, KMS,
  Plymouth, RetroPie, EmulationStation, or SafeShutdown change.
- No GPIO read or write.
- No shutdown, reboot, halt, poweroff, sleep, suspend, or resume behavior.
- No attached-keyboard assumption.
- No repository checkout assumption on the GPi Case 2.
- No Lantern Dispatch, telemetry, upload, auto-repair, installer, package, or
  network behavior.
- SSH remains optional support; handheld use remains primary.
- Current scp-first field practice remains protected.

## Outcome

- Added the dedicated
  [GPi Case 2 Relic Welcome Scroll Design](../../03-operations/gpi-case-2-relic-welcome-scroll-design.md).
- Defined purpose, non-goals, safety boundaries, welcome fields,
  never-show fields, fast-path requirements, missing-command fallback,
  `--plain` and `NO_COLOR` behavior, future homes, recovery-first rules, and
  risks.
- Proposed a first ASCII banner layout with Relic name, SSH target, uptime,
  temperature, throttled raw value, disk free, address hint, future Lantern
  artifact path, and Field Lantern read-only reminder.
- Linked the dedicated scroll design from the First Spark / Boot Veil / Relic
  Welcome Scroll design, README, and Local Diagnostics Bundle Map.

## Acceptance Checks

- [x] A dedicated Relic Welcome Scroll design exists under
  `docs/03-operations/`.
- [x] The design defines purpose and non-goals.
- [x] The design defines safety boundaries and read-only command candidates.
- [x] The design defines what should be shown on SSH login.
- [x] The design defines what should never be shown.
- [x] The design defines fast-path requirements.
- [x] The design defines fallback behavior when commands are missing.
- [x] The design defines `--plain` and `NO_COLOR` behavior for any future
  script shape.
- [x] The design names possible future homes without implementing them.
- [x] The design includes a proposed first banner layout with ASCII art and
  required fields.
- [x] The design includes risks for slow SSH login, noisy `scp` output, broken
  shell profile lockout, and exposing too much system detail.
- [x] The design includes recovery-first rules: must not block login, must be
  easy to disable, must not affect `scp`, and should detect interactive shells
  before printing.
- [x] First Spark / Boot Veil / Relic Welcome Scroll design links to the
  dedicated scroll design.
- [x] README or operations map docs link to the dedicated scroll design.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds the dedicated Relic Welcome
  Scroll design, links it from the startup and operations maps, and adds this
  quest file. The final pushed commit hash is reported in the quest handoff
  because a Git commit cannot contain its own final object hash.

## Final Notes

The Relic Welcome Scroll is now designed as a support-lane greeting, not a
boot change. Future work has a safe first shape: standalone, read-only,
interactive-only, plain-capable, no-color aware, fast to skip, and friendly to
the scp-first field trail.
