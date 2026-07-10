---
id: OPS-GPI-CASE-2-RELIC-WELCOME-SCROLL-PREVIEW-NOTES-001
title: GPi Case 2 Relic Welcome Scroll Preview Notes
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record real manual Relic Welcome Scroll preview runs for the GPi Case 2 without wiring the preview into login, boot, services, GPIO, display, shutdown, sleep, or resume behavior.
related:
  - ../../scripts/gpi-case2-relic-welcome-scroll.sh
  - gpi-case-2-relic-welcome-scroll-design.md
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - local-diagnostics-bundle-map.md
last_updated: 2026-07-10
---

# GPi Case 2 Relic Welcome Scroll Preview Notes

> The Relic Welcome Scroll is still a hand-held map from the satchel: copied
> with `scp`, invoked by an operator, and kept away from the login doorway
> until a later recovery-first rune earns that path.

This document records manual preview evidence for
[`scripts/gpi-case2-relic-welcome-scroll.sh`](../../scripts/gpi-case2-relic-welcome-scroll.sh).
It is evidence intake only. It does not install, enable, or wire the script
into SSH login, MOTD, shell startup, services, boot config, GPIO behavior,
display behavior, shutdown behavior, sleep behavior, or resume behavior.

The GPi Case 2 is a handheld Relic. SSH to `retropi@gpi` is optional support,
not the primary handheld UX. Do not assume an attached keyboard. Current field
practice remains scp-first: copy a small read-only tool to
`retropi@gpi:/home/retropi/`, run it from `/home/retropi/`, and copy artifacts
back only if needed. The GPi Case 2 does not currently have this repository
checked out on the device.

## Evidence Status Legend

| Status | Meaning |
| --- | --- |
| `observed` | A real manual preview ran on the GPi Case 2 and produced bounded evidence. |
| `needs-repeat` | The run produced useful evidence but left an important field or behavior unclear. |
| `blocked` | The preview could not run or could not be observed. |
| `superseded` | Later evidence replaces this run as the current reference. |

## Run Entry Template

```text
### YYYY-MM-DD - short run name

Evidence status:
Run date/time:
Commands used:
Normal/art output:
--plain output:
NO_COLOR output:
Speed:
Missing or unavailable fields:
Noisy output:
scp-first flow:
Operator notes:
Candidate improvements:
Next rune:
Do not overclaim:
```

## Preview Runs

### 2026-07-09 - First Manual Relic Welcome Scroll Preview

Evidence status: `observed`

Run date/time: `2026-07-09 17:49:59 PDT` from the operator workstation before
the remote command sequence. The GPi output itself did not print a local date
field.

Commands used:

```sh
scp scripts/gpi-case2-relic-welcome-scroll.sh retropi@gpi:/home/retropi/
ssh retropi@gpi 'chmod +x /home/retropi/gpi-case2-relic-welcome-scroll.sh'
ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh'
ssh retropi@gpi '/home/retropi/gpi-case2-relic-welcome-scroll.sh --plain'
ssh retropi@gpi 'NO_COLOR=1 /home/retropi/gpi-case2-relic-welcome-scroll.sh'
```

Normal/art output: ASCII Relic Welcome Scroll art displayed correctly, followed
by the `GPi Case 2 SSH field map`, manual-preview warning, field list, Field
Lantern reminder, and recovery-first wiring reminder.

`--plain` output: Worked. It printed `Relic Welcome Scroll Preview` without
the ASCII art and preserved the same field order and reminders.

`NO_COLOR` output: Worked. `NO_COLOR=1` produced the plain fallback without
the ASCII art and preserved the same field order and reminders.

Speed: Fast. The three remote preview invocations each returned in less than a
second from the operator's SSH command perspective.

Fields observed:

| Field | First normal preview value |
| --- | --- |
| Relic | `raspberrypi` |
| SSH target | `retropi@gpi` |
| User | `retropi` |
| Uptime | `0m` |
| Kernel | `Linux 6.1.21-v8+ aarch64 GNU/Linux` |
| Disk free | `25G free on /` |
| Load avg | `0.81 0.22 0.08` |
| Mem avail | `3597 MB` |
| Temp | `temp=39.4'C` |
| Throttled | `throttled=0x0` |
| Address | `172.16.7.184` |

Missing or unavailable fields: None in this first run. `vcgencmd`, disk,
memory, load, hostname, user, uptime, kernel, and address hints all returned
usable values.

Noisy output: The script output itself was compact and not too noisy for a
future interactive SSH login. The SSH transport printed a local OpenSSH
post-quantum key-exchange warning before every `scp`/`ssh` command:
`connection is not using a post-quantum key exchange algorithm`. Treat that as
transport noise outside the script. It does not come from the Relic Welcome
Scroll and would not be solved by changing banner formatting.

scp-first flow: Clean. The script copied to `/home/retropi/`, became
executable there, and ran from `/home/retropi/` without requiring a repository
checkout on the GPi Case 2. No artifact copy-back was needed because the
preview writes no artifact.

Operator notes:

- The normal output felt warm and readable as a Field Lantern map.
- The manual warning stayed visible: `Manual preview only; not installed into
  SSH login.`
- The read-only reminder stayed visible in all modes.
- The `Relic` value is the current hostname, `raspberrypi`; this is useful
  evidence but may feel less flavorful than `gpi` in a future configured
  Relic.
- The run happened immediately after boot or restart from the evidence shown
  by `Uptime: 0m`; do not treat that as a general runtime baseline.

## Candidate Improvements

- Consider adding a cheap local time field so future preview notes can capture
  the GPi's own clock without a separate command.
- Consider whether the default art should remain hidden for `NO_COLOR` or
  whether `NO_COLOR` should disable color only while leaving ASCII art visible.
  The current behavior matches the QUEST-0094 plain fallback checks.
- Consider a later hostname/display-name mapping if the Relic should greet as
  `gpi` while still preserving the raw hostname somewhere.
- If this is ever wired into login, keep the scp/non-interactive guard strict
  and test `scp` separately before enabling it.

## Next Rune

Repeat the preview after a longer handheld session and after a future local
time field exists, then decide whether a recovery-first login wiring quest is
ready to test interactive-only display without affecting `scp` or
non-interactive SSH commands.

Do not overclaim: this run proves the manually invoked preview works over the
current scp-first SSH trail. It does not prove automatic login safety, MOTD
safety, shell startup safety, long-session performance, sleep/resume behavior,
shutdown behavior, GPIO behavior, display behavior, or boot behavior.
