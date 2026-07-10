---
id: QUEST-0096
title: Design Recovery-First Relic Welcome Scroll Login Wiring
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Design how the manual Relic Welcome Scroll could eventually be wired into SSH login safely without implementing shell startup, MOTD, PAM, SSHD, service, boot, GPIO, display, shutdown, sleep, or resume behavior.
related:
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-login-wiring-design.md
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-preview-notes.md
  - ../../03-operations/gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - ../../../scripts/gpi-case2-relic-welcome-scroll.sh
  - 0093-design-relic-welcome-scroll.md
  - 0094-add-relic-welcome-scroll-preview-script.md
  - 0095-record-first-relic-welcome-scroll-preview-run.md
last_updated: 2026-07-10
---

# QUEST-0096 - Design Recovery-First Relic Welcome Scroll Login Wiring

> The scroll can greet the SSH doorway later, but first the Ledger needs the
> recovery map: quiet `scp`, interactive-only light, and one easy disable rune.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation / Design

## Intent

Create a recovery-first design for eventually wiring the Relic Welcome Scroll
into SSH login at `retropi@gpi` without implementing that wiring yet.

The design keeps current field practice intact: the GPi Case 2 does not have
this repository checked out, the preview script is copied by `scp` to
`/home/retropi/`, SSH is optional support rather than the primary handheld UX,
and the Relic should not require an attached keyboard.

## Outcome

- Added
  [GPi Case 2 Relic Welcome Scroll Login Wiring Design](../../03-operations/gpi-case-2-relic-welcome-scroll-login-wiring-design.md).
- Compared `~/.bashrc`, `~/.profile`, `/etc/profile.d/`, `/etc/motd`,
  dynamic MOTD/PAM, and manual-only wiring choices.
- Recommended a future user-scoped `~/.bashrc` hook as the safest first
  automatic wiring path, while keeping this quest design-only.
- Captured strict guard rules for interactive-only output, `scp` silence,
  non-interactive SSH silence, fail-open behavior, timeout/fast-fail behavior,
  easy disable, and plain/no-color fallback.
- Added a recovery plan with shell-startup bypass, hook rename, disable switch,
  side-switch caution, and noisy-output handling.
- Added proposed future install and uninstall commands and marked them as not
  run yet.
- Added candidate acceptance checks for a future implementation quest.
- Linked the wiring design from the Relic Welcome Scroll design, preview
  notes, First Spark / Boot Veil / Welcome Scroll design, and Local Diagnostics
  Bundle Map.

## Boundary

- Documentation/design only.
- No shell startup file edits.
- No MOTD, PAM, `sshd` config, service, boot config, display, GPIO, shutdown,
  sleep, or resume changes.
- No install on the GPi Case 2.
- No attached-keyboard assumption.
- No repository checkout assumption on the GPi Case 2.
- No telemetry, upload, installer, repair, automatic diagnostics bundle, or
  Lantern Dispatch path.
- Preserve `--plain` and `NO_COLOR` as future wiring requirements.

## Acceptance Checks

- [x] A recovery-first login wiring design exists under `docs/03-operations/`.
- [x] The design compares `~/.bashrc`, `~/.profile`, `/etc/profile.d/`,
  `/etc/motd`, dynamic MOTD/PAM, and manual-only/no auto-install.
- [x] The design recommends the safest first wiring approach without
  implementing it.
- [x] The design requires interactive-only output.
- [x] The design requires silence for `scp` and non-interactive SSH commands.
- [x] The design requires fail-open behavior if the scroll fails.
- [x] The design includes timeout or fast-fail guidance.
- [x] The design includes an easy disable switch.
- [x] The design preserves plain/no-color fallback behavior.
- [x] The design includes recovery commands and handheld UX cautions.
- [x] The design includes proposed future install and uninstall commands marked
  as not run yet.
- [x] The design includes candidate acceptance checks for future wiring.
- [x] The Relic Welcome Scroll design links to the wiring design.
- [x] The Relic Welcome Scroll preview notes link to the wiring design.
- [x] The First Spark / Boot Veil / Welcome Scroll design links to the wiring
  design.
- [x] The Local Diagnostics Bundle Map links to the wiring design.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.

## Completion Commit

- This quest is completed by the commit that adds the recovery-first login
  wiring design, links it from the related maps, and adds this quest ledger.
  The final pushed commit hash is reported in the quest handoff because a Git
  commit cannot contain its own final object hash.

## Final Notes

The Welcome Scroll now has a recovery-first login wiring map, but the Relic's
SSH doorway remains unchanged. Future work can test one guarded user-level
hook at a time while keeping `scp`, command-mode SSH, and handheld recovery
quiet.
