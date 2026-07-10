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
purpose: Design a recovery-first future path for wiring the manual Relic Welcome Scroll into SSH login without implementing shell startup, MOTD, PAM, SSHD, service, boot, GPIO, display, shutdown, sleep, or resume behavior.
related:
  - ../../scripts/gpi-case2-relic-welcome-scroll.sh
  - gpi-case-2-relic-welcome-scroll-design.md
  - gpi-case-2-relic-welcome-scroll-preview-notes.md
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - local-diagnostics-bundle-map.md
last_updated: 2026-07-10
---

# GPi Case 2 Relic Welcome Scroll Login Wiring Design

> The Welcome Scroll may someday greet an SSH traveler at `retropi@gpi`, but
> the doorway must stay quiet for `scp`, automation, and recovery paths.

This is a design document only. It does not install, enable, or modify shell
startup files, MOTD files, PAM, `sshd` config, services, boot config, GPIO,
display, shutdown, sleep, or resume behavior.

The current script is
[`scripts/gpi-case2-relic-welcome-scroll.sh`](../../scripts/gpi-case2-relic-welcome-scroll.sh).
It remains a manual, read-only preview copied by `scp` to
`retropi@gpi:/home/retropi/` and run from `/home/retropi/`. The GPi Case 2
does not currently have this repository checked out on the device.

## Purpose

Design the safest future way to show the Relic Welcome Scroll during normal
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

## Recommended First Wiring Approach

The safest first future implementation should keep the installed script and
hook in `/home/retropi/` and add a tiny guarded call from `~/.bashrc` only.

Reasoning:

- It is user-scoped to `retropi`, not system-wide.
- It matches the current `/home/retropi/` scp-first field trail.
- It can be disabled by renaming one hook file or removing one profile line.
- It avoids PAM, `sshd_config`, `/etc/motd`, `/etc/profile.d/`, services, and
  boot paths for the first automatic test.
- It can fail open to the normal shell prompt if the scroll script is missing,
  slow, or broken.

Do not implement this in this quest. A later quest should create the exact hook
script, copy/install flow, tests, and rollback notes before touching the GPi.

## Strict Guard Rules

Any future hook must obey all of these rules:

- Print only for interactive shells.
- Require terminal stdout, for example a future guard equivalent to `[ -t 1 ]`.
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

Candidate guard shape for a future quest, not run yet:

```sh
# Proposed only. Do not install from this document.
case "$-" in
  *i*) ;;
  *) return 0 2>/dev/null || exit 0 ;;
esac

[ -t 1 ] || return 0 2>/dev/null || exit 0
[ -e "$HOME/.gpi-relic-welcome-scroll.disabled" ] && return 0 2>/dev/null || exit 0
[ -x "$HOME/gpi-case2-relic-welcome-scroll.sh" ] || return 0 2>/dev/null || exit 0

if command -v timeout >/dev/null 2>&1; then
  timeout 2s "$HOME/gpi-case2-relic-welcome-scroll.sh" || true
else
  "$HOME/gpi-case2-relic-welcome-scroll.sh" || true
fi
```

The final implementation may need a shell-function wrapper instead of raw
`return`/`exit` lines so that it behaves correctly in the selected startup
file. That belongs in the future implementation quest.

## Proposed Future Install Commands

These commands are proposed for a later quest. They have not been run by this
design quest.

```sh
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
ssh retropi@gpi 'chmod +x /home/retropi/gpi-case2-relic-welcome-scroll.sh'
ssh retropi@gpi 'test -f /home/retropi/.bashrc && cp /home/retropi/.bashrc /home/retropi/.bashrc.relic-welcome-scroll.bak'
ssh retropi@gpi 'printf "%s\n" "[ -f /home/retropi/.gpi-relic-welcome-scroll-hook.sh ] && . /home/retropi/.gpi-relic-welcome-scroll-hook.sh" >> /home/retropi/.bashrc'
scp FUTURE-HOOK-FILE retropi@gpi:/home/retropi/.gpi-relic-welcome-scroll-hook.sh
```

A future quest should replace `FUTURE-HOOK-FILE` with a reviewed hook file and
should test `scp` and non-interactive SSH before leaving the hook enabled.

## Proposed Future Uninstall Commands

These commands are proposed for a later quest. They have not been run by this
design quest.

Disable without editing profile files:

```sh
ssh retropi@gpi 'touch /home/retropi/.gpi-relic-welcome-scroll.disabled'
```

Rename the hook so startup cannot source it:

```sh
ssh retropi@gpi 'mv /home/retropi/.gpi-relic-welcome-scroll-hook.sh /home/retropi/.gpi-relic-welcome-scroll-hook.sh.disabled'
```

Restore the backed-up Bash profile if the future install created one:

```sh
ssh retropi@gpi 'cp /home/retropi/.bashrc.relic-welcome-scroll.bak /home/retropi/.bashrc'
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

## Candidate Acceptance Checks

A future implementation quest should not be accepted until all checks pass:

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

## Boundaries For The Future Quest

The first implementation quest should only wire the SSH support scroll. It
should not combine this with First Spark, Boot Veil, services, boot config,
GPIO, display, shutdown, sleep, resume, installer, telemetry, upload, or
Lantern Dispatch work.

The Relic Welcome Scroll should stay a small greeting from the satchel. If it
cannot be shown safely, manual-only remains the right answer.
