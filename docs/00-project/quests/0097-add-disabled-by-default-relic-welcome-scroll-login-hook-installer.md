---
id: QUEST-0097
title: Add Disabled-by-Default Relic Welcome Scroll Login Hook Installer
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Add a recovery-first, disabled-by-default installer and uninstaller for wiring the Relic Welcome Scroll into interactive SSH login without enabling it during tests or normal script execution.
related:
  - ../../../scripts/gpi-case2-install-relic-welcome-scroll-hook.sh
  - ../../../scripts/gpi-case2-relic-welcome-scroll.sh
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-login-wiring-design.md
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-design.md
  - ../../03-operations/gpi-case-2-relic-welcome-scroll-preview-notes.md
  - ../../03-operations/local-diagnostics-bundle-map.md
  - 0093-design-relic-welcome-scroll.md
  - 0094-add-relic-welcome-scroll-preview-script.md
  - 0095-record-first-relic-welcome-scroll-preview-run.md
  - 0096-design-recovery-first-relic-welcome-scroll-login-wiring.md
last_updated: 2026-07-10
---

# QUEST-0097 - Add Disabled-by-Default Relic Welcome Scroll Login Hook Installer

> The Relic may greet an SSH traveler, but only when the operator opens that
> map: status first, install by hand, quiet `scp`, and an easy uninstall rune.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Script / Documentation

## Intent

Add a recovery-first installer/uninstaller for the Relic Welcome Scroll login
hook. The installer must be disabled by default: it does nothing to login
unless the operator explicitly runs `--install`, and tests must not enable the
real workstation or GPi login hook.

The GPi field trail remains scp-first. The GPi Case 2 does not currently have
this repository checked out on the device, the SSH target is `retropi@gpi`,
and the home path is `/home/retropi/`. SSH is optional support for the
handheld Relic, not the primary UX, and no attached keyboard should be
assumed.

## Outcome

- Added
  [`scripts/gpi-case2-install-relic-welcome-scroll-hook.sh`](../../../scripts/gpi-case2-install-relic-welcome-scroll-hook.sh).
- Supported `--help`, `--install`, `--uninstall`, `--status`, `--dry-run`,
  `--plain`, and `NO_COLOR`.
- Kept install explicit: running the script without an action exits without
  modifying login behavior.
- Installed only a small guarded hook at
  `/home/retropi/.gpi-relic-welcome-scroll-hook.sh` and one marked source
  block in `/home/retropi/.bashrc`.
- Created timestamped `.bashrc` backups before edits and avoided duplicate
  hook blocks.
- Made uninstall remove the marked `.bashrc` block and move the hook aside.
- Guarded hook output to interactive terminal shells only, with silence for
  `SSH_ORIGINAL_COMMAND`, non-interactive sessions, missing scroll scripts,
  disabled flag file, disabled environment variable, and scroll failures.
- Added a two-second `timeout` path when the platform provides `timeout`.
- Preserved the project boundary: no `sshd`, PAM, MOTD, service, boot, GPIO,
  display, shutdown, sleep, or resume changes.
- Added script checks to `make check-scripts`, including help, status,
  dry-run, plain/no-color, and temp-`HOME` install/uninstall checks.
- Updated the login wiring design, Relic Welcome Scroll design, preview notes,
  and Local Diagnostics Bundle Map.

## scp-First Field Instructions

```sh
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
scp scripts/gpi-case2-install-relic-welcome-scroll-hook.sh retropi@gpi:/home/retropi/
ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --status'
ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --install'
ssh retropi@gpi
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/gpi-case2-relic-welcome-scroll.scp-test
ssh retropi@gpi 'echo spellbook'
ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --uninstall'
```

Run `--uninstall` if anything feels noisy, slow, or wrong.

## Boundary

- Disabled by default.
- No automatic install during tests or normal script execution.
- No large banner in shell startup files.
- No duplicate hook blocks.
- No login block if the scroll script is missing, disabled, slow enough to
  time out, or exits with an error.
- No `sshd` config, PAM, MOTD, service, boot config, GPIO, display, shutdown,
  sleep, or resume changes.
- No repository checkout assumption on the GPi Case 2.
- No attached-keyboard assumption.

## Acceptance Checks

- [x] Installer script exists under `scripts/`.
- [x] Installer supports `--help`.
- [x] Installer supports `--install`.
- [x] Installer supports `--uninstall`.
- [x] Installer supports `--status`.
- [x] Installer supports `--plain` and respects `NO_COLOR`.
- [x] Installer supports `--dry-run`.
- [x] Installer does nothing to login unless `--install` or `--uninstall` is
  explicitly selected.
- [x] Installer creates backups before `.bashrc` edits.
- [x] Installer writes a small guarded hook instead of inlining banner output.
- [x] Installer avoids duplicate hook blocks.
- [x] Installer fails safely if the welcome scroll script is missing.
- [x] Hook prints only for interactive terminal shells.
- [x] Hook stays quiet for non-interactive SSH command mode and `scp`-style
  sessions.
- [x] Hook can be disabled by
  `/home/retropi/.gpi-relic-welcome-scroll.disabled`.
- [x] Hook can be disabled by `GPI_RELIC_WELCOME_SCROLL_DISABLED`.
- [x] Hook does not block login if the scroll script errors.
- [x] Hook uses `timeout` when available.
- [x] `make check-scripts` covers the new installer.
- [x] Docs include scp-first field instructions.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.
- [x] `make check-scripts` passed.
- [x] `git diff --check` passed.

## Completion Commit

- This quest is completed by the commit that adds the disabled-by-default
  installer, validation checks, documentation updates, and this quest ledger.
  The final pushed commit hash is reported in the quest handoff because a Git
  commit cannot contain its own final object hash.

## Final Notes

The Relic's SSH doorway now has an operator-run welcome scroll installer, but
the default path remains quiet. The Field Lantern is in the satchel, `--status`
is the first rune, `--install` is deliberate, and `--uninstall` keeps recovery
plain.
