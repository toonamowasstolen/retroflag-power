---
id: OPS-GPI-CASE-2-RELIC-WELCOME-SCROLL-DESIGN-001
title: GPi Case 2 Relic Welcome Scroll Design
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Design a read-only, fast, recovery-first SSH welcome scroll for GPi Case 2 operator sessions without implementing shell startup, MOTD, boot, service, GPIO, shutdown, sleep, or resume behavior.
related:
  - ../../scripts/gpi-case2-relic-welcome-scroll.sh
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - gpi-case-2-true-boot-trace-lantern-design.md
  - gpi-case-2-true-boot-trace-field-run-procedure.md
  - gpi-case-2-true-boot-trace-evidence-ledger.md
  - gpi-case-2-session-watch-lantern-design.md
  - gpi-case-2-session-watch-field-run-procedure.md
  - gpi-case-2-session-watch-evidence-ledger.md
  - gpi-case-2-recovery-first-field-procedure.md
last_updated: 2026-07-10
---

# GPi Case 2 Relic Welcome Scroll Design

> When an operator opens an SSH door to `retropi@gpi`, the Relic should greet
> them with a quick map: where they are, what the weather looks like, and which
> Field Lantern paths are nearby.

The first manually runnable preview script now lives at
[`scripts/gpi-case2-relic-welcome-scroll.sh`](../../scripts/gpi-case2-relic-welcome-scroll.sh).
It is a standalone preview only. It is not installed into SSH login, MOTD,
shell startup, services, boot config, GPIO behavior, display behavior,
shutdown behavior, sleep behavior, or resume behavior.

The GPi Case 2 is a handheld Relic. SSH is optional support, not the primary
handheld UX. Do not assume an attached keyboard. Current field practice is
scp-first: copy a small read-only tool to `retropi@gpi:/home/retropi/`, run it
from `/home/retropi/`, then retrieve the artifact with `scp`. The GPi Case 2
does not currently have this repository checked out on the device.

## Purpose

The Relic Welcome Scroll is a friendly, bounded system-info banner for
interactive SSH logins to `retropi@gpi`.

It should help the operator answer:

- Which Relic did I reach?
- Is the system responsive enough for a short support session?
- Are there obvious thermal, throttling, disk, or network clues?
- Where would the latest Field Lantern artifact be, if a future local pointer
  exists?
- Which commands remain read-only unless the scroll explicitly says otherwise?

The scroll is a greeting and orientation map. It is not a diagnostic bundle, a
health verdict, an installer, a repair tool, or a replacement for the handheld
frontend.

## Current Preview Script

Use the preview through the current scp-first field trail:

```sh
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
ssh retropi@gpi 'chmod +x /home/retropi/gpi-case2-relic-welcome-scroll.sh'
ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh'
```

Plain and no-color previews:

```sh
ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh --plain'
ssh retropi@gpi 'NO_COLOR=1 /home/retropi/gpi-case2-relic-welcome-scroll.sh'
```

Warning: this script is not yet installed into SSH login. Keep it manually
invoked until a later recovery-first wiring quest proves an interactive-only,
easy-disable login path.

The preview prints hostname/Relic name, user, uptime, kernel summary,
temperature and throttled raw value when `vcgencmd` is available, root disk
free, load average, memory available, an address hint, the `retropi@gpi` SSH
target reminder, and a read-only Field Lantern reminder. Missing data shows as
`unavailable` or `unknown`; the scroll should not fail as a whole.

## Non-Goals

- Do not wire the preview into automatic login or boot behavior in this quest.
- Do not edit `.profile`, `.bashrc`, `.zshrc`, `/etc/profile`, `/etc/motd`,
  `/etc/update-motd.d/`, PAM, SSHD config, shell startup files, or login
  managers.
- Do not install packages or enable services.
- Do not write GPIO, read GPIO, change `SafeShutdown.py`, alter side-switch
  behavior, or touch top-button sleep/resume behavior.
- Do not change `/boot/config.txt`, `/boot/cmdline.txt`, framebuffer, KMS,
  Plymouth, RetroPie, EmulationStation, or display behavior.
- Do not require git, Go, this repository, root, network access, or an attached
  keyboard on the GPi Case 2.
- Do not create Lantern Dispatch, telemetry, uploads, auto-repair, or automatic
  support bundle generation.

## Safety Boundaries

The scroll must be read-only and local. A future implementation may call only
small, allowlisted commands that read already-available state.

Allowed evidence candidates:

- `hostname` or `/etc/hostname`.
- `/proc/uptime`.
- `date` for local time.
- `df` for the root filesystem.
- `vcgencmd measure_temp` if `vcgencmd` exists.
- `vcgencmd get_throttled` if `vcgencmd` exists.
- `ip` or `hostname -I` for a short local address hint.
- A future local pointer file for the last known Field Lantern artifact path,
  if the pointer exists and is cheap to read.

The scroll must never broaden collection to make a missing command available.
Missing commands are normal on field devices and should be shown as
`unavailable`, `unknown`, or omitted.

## What To Show

The first layout should be compact enough to scan before the prompt appears:

```text
       _.-._
    .-'  |  '-.        Relic Welcome Scroll
   /  .--+--.  \       GPi Case 2 Field Lantern
   |  |  *  |  |
   \  '--+--'  /       Relic: gpi
    '-.  |  .-'        SSH: retropi@gpi
       '-'

   Uptime: 2h 14m                 Temp: 58.4 C
   Throttled: 0x50000             Disk free: 18G on /
   Address: 192.168.1.42          Lantern: /home/retropi/gpi-case2-...txt

   Field Lantern tools are read-only unless a command says otherwise.
```

Fields:

| Field | Source idea | Fallback |
| --- | --- | --- |
| Relic name / hostname | `hostname` or `/etc/hostname` | `unknown-relic` |
| SSH target note | Static `retropi@gpi` note | Omit if plain/minimal mode needs space. |
| Uptime | `/proc/uptime` | `unknown` |
| Temperature | `vcgencmd measure_temp` | `unavailable` |
| Throttled raw value | `vcgencmd get_throttled` | `unavailable` |
| Disk free | `df -h /` | `unknown` |
| IP/address hint | `ip -o -4 addr` or `hostname -I` | `unknown` |
| Lantern artifact path | Future local pointer file | `none recorded` |
| Safety reminder | Static text | Always show in interactive rich mode. |

The throttled value should be raw and humble. For example, `0x50000` is a clue,
not a diagnosis. The scroll should not decode sticky throttling flags into a
root cause unless a later design adds a small, tested interpretation table.

## What Never To Show

The scroll must never print:

- Private keys, tokens, passwords, Wi-Fi secrets, or SSH authorized keys.
- Full environment dumps.
- Shell history.
- ROM lists, save files, usernames beyond the login context, or personal paths
  unrelated to Field Lantern artifacts.
- Long `journalctl`, `dmesg`, process, mount, package, or network dumps.
- Public IP discovery, remote geolocation, external network calls, or upload
  status.
- Automatic suggestions to run shutdown, reboot, sleep, resume, firmware,
  installer, update, repair, or GPIO commands.
- Any claim that the handheld is safe, healthy, overheating, undervolting, or
  fixed based only on the banner.

## Fast-Path Requirements

SSH login must stay fast. A future implementation should aim for a sub-second
path on the GPi Case 2 and must never block login.

Rules:

- Use only local reads and tiny commands.
- Avoid `journalctl`, broad `find`, package managers, network probes, DNS
  lookups, `systemctl` trees, and recursive directory scans.
- Bound every optional probe with a short timeout if the shell environment
  provides a portable way to do so.
- Prefer simple files such as `/proc/uptime` over heavier commands.
- Print at most one screen of output.
- Degrade to a shorter scroll if any probe is slow, missing, or denied.
- Never wait for the latest Lantern artifact by searching the whole home
  directory.

If a future scroll cannot gather a field cheaply, it should leave a quiet blank
or print `unknown`. The prompt matters more than the flourish.

## Missing Command Fallbacks

The scroll should treat missing commands as ordinary field evidence:

| Missing or denied command | Expected behavior |
| --- | --- |
| `vcgencmd` missing | Temperature and throttled fields show `unavailable` or are omitted. |
| `df` missing | Disk free shows `unknown`; login continues. |
| `ip` missing | Try `hostname -I`; if unavailable, show `unknown`. |
| `/proc/uptime` unreadable | Uptime shows `unknown`. |
| Pointer file missing | Lantern artifact shows `none recorded`. |
| Any probe times out | Mark that single field `timed out`; continue login. |

Failures should never spill shell errors into the operator's prompt. Redirect
or capture expected command failures and print the friendly fallback instead.

## Plain And No-Color Behavior

A future script shape should respect both `--plain` and `NO_COLOR`.

Plain mode:

- No ANSI color.
- No box drawing or non-ASCII characters.
- Minimal ASCII art or no art if space is tight.
- Stable labels for copy/paste into an issue or Ledger note.

`NO_COLOR` behavior:

- Disable ANSI color automatically when `NO_COLOR` is set to any non-empty
  value.
- Keep the same fields and ordering where possible.
- Do not use color as the only carrier of warnings or status.

The default rich mode may use warm retro/adventurer styling, but the text must
still be readable in a basic terminal.

## Future Home

Possible future implementation homes, in order of caution:

| Home | Why it might fit | Boundary |
| --- | --- | --- |
| Standalone script in `/home/retropi/` | Easy to copy with `scp`, run manually, and remove. | Does not print automatically until a later quest wires it. |
| User shell profile call | Can target only `retropi` interactive SSH sessions. | Must detect interactive shells and be easy to disable. |
| `/etc/update-motd.d/` or MOTD hook | Familiar login-banner path on Debian-like systems. | Requires more system-level caution and a clean disable path. |

The standalone preview now exists under `scripts/` and is copied manually to
`/home/retropi/` when needed. A future implementation quest may wire automatic
printing only after interactive detection, `scp` safety, and disable rules are
validated.

## Recovery-First Rules

- Must not block login.
- Must be easy to disable by renaming, removing, or bypassing one clearly named
  file or profile line.
- Must not affect `scp`, non-interactive SSH commands, rsync, or automation.
- Must detect an interactive shell before printing.
- Must fail open to a normal shell prompt.
- Must keep a plain fallback that can be pasted into recovery notes.
- Must not depend on the top sleep/resume button, attached keyboard, or local
  repository checkout.

Interactive detection should check for an interactive shell and a terminal
before printing. Future implementation notes may include guards such as testing
for `case "$-" in *i*) ...` and `[ -t 1 ]`, but this document does not install
those guards anywhere.

## Risks

| Risk | Why it matters | Design response |
| --- | --- | --- |
| Slow SSH login | The operator may need the shell during recovery. | Keep probes local, tiny, optional, and timeout-bounded. |
| Noisy output for `scp` or non-interactive sessions | The current field practice is scp-first and artifacts must copy cleanly. | Print only for interactive terminal sessions. |
| Broken shell profile locking out operator | A syntax error in login startup can block support access. | Prefer standalone first; if wired later, keep one-line disable and fail-open guards. |
| Exposing too much system detail | A friendly banner can accidentally leak private context. | Use a narrow allowlist and never print secrets, histories, ROM lists, or broad dumps. |
| Misleading health claims | A raw temperature or throttled value can be overread. | Label fields as clues and avoid diagnosis in the scroll. |

## Relationship To Startup UX

The broader startup UX map lives in
[GPi Case 2 First Spark / Boot Veil / Relic Welcome Scroll Design](gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md).
That document keeps First Spark, Boot Veil, and SSH welcome output separate.
This document is the source-of-truth design for the SSH welcome lane only.

True Boot Trace and Session Watch Lantern artifacts may eventually feed a
single `last known Lantern artifact` pointer, but the welcome scroll should
not search for those artifacts at login time. The scroll lights the map; it
does not empty the whole satchel onto the floor.
