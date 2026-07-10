---
id: OPS-GPI-CASE-2-RELIC-WELCOME-SCROLL-LOGIN-WIRING-DESIGN-001
title: GPi Case 2 Relic Welcome Scroll Login Wiring Design
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Design and record the recovery-first, disabled-by-default path for wiring the manual Relic Welcome Scroll into interactive SSH login without changing MOTD, PAM, SSHD, service, boot, GPIO, display, shutdown, sleep, or resume behavior.
related:
  - ../../scripts/gpi-case2-install-relic-welcome-scroll-hook.sh
  - ../../scripts/gpi-case2-relic-welcome-scroll.sh
  - gpi-case-2-relic-welcome-scroll-design.md
  - gpi-case-2-relic-welcome-scroll-preview-notes.md
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - local-diagnostics-bundle-map.md
last_updated: 2026-07-10
---

# GPi Case 2 Relic Welcome Scroll Login Wiring Design

> The Welcome Scroll may greet an SSH traveler at `retropi@gpi`, but only
> after the operator installs the hook, and the doorway must stay quiet for
> `scp`, automation, and recovery paths.

This document records the design and field instructions for the
disabled-by-default installer. Reading this document, running normal checks, or
running the installer without `--install` does not enable login output. The
installer does not modify MOTD files, PAM, `sshd` config, services, boot
config, GPIO, display, shutdown, sleep, or resume behavior.

The current scroll script is
[`scripts/gpi-case2-relic-welcome-scroll.sh`](../../scripts/gpi-case2-relic-welcome-scroll.sh).
The disabled-by-default installer is
[`scripts/gpi-case2-install-relic-welcome-scroll-hook.sh`](../../scripts/gpi-case2-install-relic-welcome-scroll-hook.sh).
Both scripts are copied by `scp` to `retropi@gpi:/home/retropi/` and run from
`/home/retropi/`. The GPi Case 2 does not currently have this repository
checked out on the device.

## Purpose

Record the safest first way to show the Relic Welcome Scroll during normal
interactive SSH login while preserving these field rules:

- SSH is optional support, not the primary handheld UX.
- The GPi Case 2 is a handheld Relic; do not assume an attached keyboard.
- Current GPi field practice is scp-first.
- Login output must not break `scp`, non-interactive SSH commands,
  automation, or recovery access.
- The scroll must respect `--plain` and `NO_COLOR`.
- The tone can stay warm and adventurer-like: Relic, Field Lantern, Ledger,
  Spellbook, SignalMage, satchel, rune, and map.

## Wiring Locations Compared

| Location | What it would do | Recovery fit | Risk |
| --- | --- | --- | --- |
| `~/.bashrc` | Runs for interactive Bash shells and can be guarded with `$-` and terminal checks. | Best user-level first candidate because one small line in `/home/retropi/.bashrc` can call one script in `/home/retropi/`. | Bash-specific; must not print for non-interactive command mode or any shell used by `scp`. |
| `~/.profile` | Runs for login shells before shell-specific interactive setup. | Easy to find in `/home/retropi/`, but broader than needed. | More likely to affect non-interactive SSH command startup depending on shell/session behavior. |
| `/etc/profile.d/` | Runs for users whose login shell sources system profile snippets. | Central and conventional on Debian-like systems. | System-wide, needs elevated edits, and is too broad for first wiring. |
| `/etc/motd` | Static text shown by login stack. | Easy to clear if reachable, but cannot run the dynamic scroll script. | Static only; does not naturally respect `NO_COLOR`, `--plain`, or live Field Lantern clues. |
| Dynamic MOTD / PAM | Can run scripts through PAM/MOTD hooks during login. | Familiar for login banners when carefully managed. | Higher blast radius, root-owned config, possible ordering surprises, and more ways to affect recovery. |
| Manual-only / no auto-install | Keep using `scp` and manual invocation. | Safest current state. Nothing automatic can break login. | No automatic greeting; operator must remember the command. |

## Implemented First Wiring Approach

The first implementation keeps the installed script and hook in
`/home/retropi/` and adds one tiny guarded source block to `~/.bashrc` only.
The operator must run the installer with `--install`; tests and normal script
execution do not enable the hook.

Reasoning:

- It is user-scoped to `retropi`, not system-wide.
- It matches the current `/home/retropi/` scp-first field trail.
- It can be disabled by renaming one hook file or removing one profile line.
- It avoids PAM, `sshd_config`, `/etc/motd`, `/etc/profile.d/`, services, and
  boot paths for the first automatic test.
- It can fail open to the normal shell prompt if the scroll script is missing,
  slow, or broken.

The installer writes the small hook file
`/home/retropi/.gpi-relic-welcome-scroll-hook.sh`, backs up `.bashrc` before
editing it, avoids duplicate source blocks, and moves the hook aside during
`--uninstall`.

## Strict Guard Rules

The hook must obey all of these rules:

- Print only for interactive shells.
- Require terminal stdout, using a guard equivalent to `[ -t 1 ]`.
- Never print during `scp`, non-interactive SSH commands, rsync, automation,
  or scripted copy/retrieve paths.
- Never block login if the scroll script fails.
- Treat a missing scroll script as normal and continue silently.
- Bound the scroll with a short timeout or fast-fail path if the platform has
  a reliable timeout command.
- Keep one obvious disable switch, such as
  `/home/retropi/.gpi-relic-welcome-scroll.disabled`.
- Respect `NO_COLOR` whenever it is set.
- Provide a plain/no-color fallback by calling the script with `--plain` when
  color or terminal capability is uncertain.
- Redirect expected hook errors away from the prompt unless a deliberate debug
  mode is enabled.

Current guard shape:

```sh
case "$-" in *i*) ;; *) return 0 ;; esac
[ -t 1 ] || return 0
[ -z "${SSH_ORIGINAL_COMMAND:-}" ] || return 0
[ -z "${GPI_RELIC_WELCOME_SCROLL_DISABLED:-}" ] || return 0
[ ! -e "$HOME/.gpi-relic-welcome-scroll.disabled" ] || return 0
[ -x "$HOME/gpi-case2-relic-welcome-scroll.sh" ] || return 0

if command -v timeout >/dev/null 2>&1; then
  timeout 2s "$HOME/gpi-case2-relic-welcome-scroll.sh" 2>/dev/null || true
else
  "$HOME/gpi-case2-relic-welcome-scroll.sh" 2>/dev/null || true
fi
```

The generated hook wraps this guard in a shell function so it can be safely
sourced by `.bashrc`.

## scp-First Field Install Commands

Copy both scripts to the GPi home directory first. Do not assume the repository
is checked out on the handheld Relic. Run `--status` before `--install`.

```sh
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
scp scripts/gpi-case2-install-relic-welcome-scroll-hook.sh retropi@gpi:/home/retropi/
ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --status'
ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --install'
```

After install, open a normal interactive login and confirm the scroll appears:

```sh
ssh retropi@gpi
```

Then confirm copy and command paths stay quiet:

```sh
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/gpi-case2-relic-welcome-scroll.scp-test
ssh retropi@gpi 'true'
ssh retropi@gpi 'echo spellbook'
```

If anything feels wrong, uninstall immediately:

```sh
ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --uninstall'
```

## Uninstall And Disable Commands

Simple uninstall:

```sh
ssh retropi@gpi 'sh /home/retropi/gpi-case2-install-relic-welcome-scroll-hook.sh --uninstall'
```

Disable without editing profile files:

```sh
ssh retropi@gpi 'touch /home/retropi/.gpi-relic-welcome-scroll.disabled'
```

Rename the hook so startup cannot source it:

```sh
ssh retropi@gpi 'mv /home/retropi/.gpi-relic-welcome-scroll-hook.sh /home/retropi/.gpi-relic-welcome-scroll-hook.sh.disabled'
```

Restore a backed-up Bash profile only if a manual recovery needs it:

```sh
ssh retropi@gpi 'ls -1 /home/retropi/.bashrc.relic-welcome-scroll.bak-*'
ssh retropi@gpi 'cp /home/retropi/.bashrc.relic-welcome-scroll.bak-YYYYMMDD-HHMMSS /home/retropi/.bashrc'
```

Remove the preview script only after confirming no manual field run needs it:

```sh
ssh retropi@gpi 'rm -f /home/retropi/gpi-case2-relic-welcome-scroll.sh'
```

## Recovery Plan

If login output becomes noisy, slow, or broken, keep recovery boring and use
the least invasive path that still works.

Bypass shell startup if possible:

```sh
ssh -t retropi@gpi 'bash --noprofile --norc'
```

Run a remote command that renames the hook without opening a full interactive
shell:

```sh
ssh retropi@gpi 'mv /home/retropi/.gpi-relic-welcome-scroll-hook.sh /home/retropi/.gpi-relic-welcome-scroll-hook.sh.disabled'
```

Use the disable switch:

```sh
ssh retropi@gpi 'touch /home/retropi/.gpi-relic-welcome-scroll.disabled'
```

If command mode is also noisy or broken, try the handheld path first: use the
side switch only as the normal stock shutdown path while the system is
responsive, avoid relying on the top sleep/resume button unless a dedicated
procedure names that test, then boot again and retry SSH. Do not assume an
attached keyboard on the Relic.

If `scp` output becomes noisy, disable the hook before collecting artifacts.
The scroll is a map, not a toll at the doorway; artifact copy paths must stay
quiet.

## Acceptance Checks

The implementation quest should not be accepted until all checks pass:

- Normal interactive `ssh retropi@gpi` shows the Relic Welcome Scroll once.
- `scp` to `/home/retropi/` stays silent except for transport-level SSH noise.
- `scp` from `/home/retropi/` stays silent except for transport-level SSH noise.
- `ssh retropi@gpi 'true'` does not show the banner.
- `ssh retropi@gpi 'echo spellbook'` prints only `spellbook`.
- `ssh -t retropi@gpi 'bash -i -c exit'` may show the scroll because it is
  explicitly interactive.
- `NO_COLOR=1` remains honored by the script path.
- `--plain` remains available for manual preview and fallback use.
- Touching `/home/retropi/.gpi-relic-welcome-scroll.disabled` restores normal
  login output.
- Renaming the hook restores normal login output.
- A missing or non-executable scroll script does not print shell errors and
  does not block login.
- A slow scroll probe times out or fast-fails without delaying recovery.

## Boundaries For This Lane

This lane should only wire the SSH support scroll. It should not combine this
with First Spark, Boot Veil, services, boot config, GPIO, display, shutdown,
sleep, resume, telemetry, upload, or Lantern Dispatch work.

The Relic Welcome Scroll should stay a small greeting from the satchel. If it
cannot be shown safely, manual-only remains the right answer.
