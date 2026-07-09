---
id: OPS-GPI-CASE-2-TRUE-BOOT-TRACE-LANTERN-DESIGN-001
title: GPi Case 2 True Boot Trace Lantern Design
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Design a future read-only, scp-first GPi Case 2 Boot Trace Lantern for capturing startup evidence without changing handheld behavior.
related:
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - ../../scripts/gpi-case2-true-boot-trace-lantern.sh
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-boot-power-trace-capture-procedure.md
  - gpi-case-2-session-watch-lantern-design.md
  - gpi-case-2-session-watch-field-run-procedure.md
  - gpi-case-2-session-watch-evidence-ledger.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - human-facing-field-lantern-script-ux-standard.md
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
  - ../04-architecture/arcadia-runtime-migration-path.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
last_updated: 2026-07-09
---

# GPi Case 2 True Boot Trace Lantern Design

> The True Boot Trace Lantern is the startup field lantern: a read-only Relic
> copied by hand, run from `/home/retropi/`, and trusted to bring back one
> careful Boot Trace Ledger without moving the stones under the handheld.

This design now has a first read-only foreground script skeleton:
[`scripts/gpi-case2-true-boot-trace-lantern.sh`](../../scripts/gpi-case2-true-boot-trace-lantern.sh).
It does not install a service, change boot behavior, or approve automatic
startup. The current tool is a small foreground shell script copied to the GPi
Case 2 with `scp`, launched from `/home/retropi/`, and retrieved with `scp`
when the handheld remains responsive.

The GPi Case 2 is a handheld Relic first. Do not assume an attached keyboard.
SSH to `retropi@gpi` is optional support for copying, launching, and retrieving
the Ledger; it is not the primary handheld UX. While the Relic is responsive,
the side switch remains the normal stock shutdown path. The top sleep/resume
button remains suspect unless a future procedure explicitly says otherwise.

## Purpose

Design the true GPi Case 2 Boot Trace Lantern: a read-only field tool that
captures startup evidence after a normal boot command path has reached a
launchable shell, without changing power, display, boot, sleep, resume,
shutdown, GPIO, or service behavior.

The Lantern should help answer whether observed GPi issues appear to cluster:

- During boot or early userspace.
- During KMS, VC4, V3D, framebuffer, or display handoff.
- During USB, audio, controller, xpad, or input initialization.
- Around EmulationStation or another frontend becoming visible.
- Later during idle, sleep-risk, resume-risk, or shutdown edges.

It complements the Session Watch Lantern. Boot Trace maps startup. Session
Watch maps a bounded responsive handheld session after boot.

## Non-Goals

The True Boot Trace Lantern must not:

- Change `/boot/config.txt`, `/boot/cmdline.txt`, `/etc/rc.local`, systemd
  units, RetroPie config, EmulationStation config, audio config, display
  config, or `/opt/RetroFlag/SafeShutdown.py`.
- Read GPIO or write GPIO.
- Execute shutdown, reboot, halt, poweroff, suspend, sleep, or resume.
- Start automatically at boot in this design.
- Install packages, firmware, RetroFlag scripts, services, or Arcadia Runtime
  Casters.
- Run `lcdfirst.sh`, `lcdnext.sh`, RetroFlag installers, firmware tools, or
  repair commands.
- Contact Lantern Dispatch, upload telemetry, submit artifacts, or make
  project-code network calls.
- Collect secrets, private keys, Wi-Fi credentials, broad environment dumps,
  shell history, ROM lists, save data, screenshots, or arbitrary journals.
- Diagnose a single root cause from one trace.

## Safety Boundaries

The Lantern must stay read-only, local, foreground, bounded, and
operator-driven.

Required boundaries:

- Copy the script to `retropi@gpi:/home/retropi/` and run it from
  `/home/retropi/`.
- Do not require a repository checkout on the GPi Case 2.
- Do not require root for the first design path.
- Do not mutate files except for the final Ledger artifact and any temporary
  files inside a clearly named capture path under `/home/retropi/`.
- Do not leave the long-running script silent.
- Respect `--plain` and `NO_COLOR`.
- Keep command collection allowlisted and narrow.
- Treat missing commands, denied permissions, timeouts, and empty outputs as
  evidence, not as reasons to install or repair.
- If the handheld wedges, stop interpreting and follow the recovery-first
  field procedure.

The side switch is the normal stock shutdown path while the system is
responsive. Do not treat the side switch as a guaranteed emergency path during
a kernel stall. Do not press the top sleep/resume button during a boot trace
unless a later dedicated procedure explicitly names that test.

## Scp-First Field Flow

Expected field flow for the current skeleton:

1. Boot the GPi Case 2 normally and keep the handheld physically observable.
2. From the workstation, copy the one portable Lantern Relic:

```sh
scp scripts/gpi-case2-true-boot-trace-lantern.sh retropi@gpi:/home/retropi/
```

3. Start it through optional SSH support:

```sh
ssh retropi@gpi
sh /home/retropi/gpi-case2-true-boot-trace-lantern.sh --plain --duration 120 --interval 5
```

4. Watch for the startup banner, read-only safety line, progress lines, and
   planned artifact path.
5. Keep the handheld in view. Record display state, LED state, SSH state,
   input responsiveness, and whether EmulationStation appears.
6. Avoid the top sleep/resume button.
7. Let the bounded capture finish.
8. Copy the final `Artifact:` path back to the workstation:

```sh
scp retropi@gpi:/home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS.txt .
```

9. Inspect and redact the Ledger before sharing.
10. If field work is complete and the system is responsive, shut down with the
    normal side-switch path.

This scp-first flow acknowledges current field reality: the GPi Case 2 does
not have the repository checked out on the device.

## Artifact Pattern

The first true Boot Trace skeleton produces one final text Ledger:

```text
/home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS.txt
```

If a later fuller version needs a satchel folder, it may add:

```text
/home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS/
/home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS.tar.gz
```

The terminal must always end with copyable lines:

```text
Artifact: /home/retropi/gpi-case2-true-boot-trace-lantern-20260709-121500.txt
Duration: 00:02:03
Status:   completed
```

The final Ledger is the source of truth. Terminal output is the campfire view.

## Progress And Status Output

Manual Field Lantern scripts must show progress, timing, status, and final
artifact path. A boot trace must never sit silently.

Expected output behavior:

- Startup banner naming the Relic, duration, interval, and read-only promise.
- `[[n/m]]` stage labels for setup, initial snapshot, sampling, artifact
  writing, and summary.
- Progress during sampling with elapsed time, remaining time, sample count,
  latest `get_throttled` value when available, frontend detection state, and
  latest warning count.
- Per-stage timing and total duration.
- Final status: `completed`, `completed_with_warnings`, `interrupted`, or
  `failed`.
- Final artifact path in an `Artifact:` line.
- Short inspect-and-redact reminder.

`--plain` and `NO_COLOR` must disable color and any terminal decoration. Plain
mode should be ASCII-only, stable, and easy to paste into a Ledger note.

## Recommended Bounded Duration

The first field-safe default should be short:

- Recommended default: 120 seconds.
- Smoke test: 30 to 60 seconds.
- First real startup pass: 120 to 180 seconds.
- Upper bound for early field work: 300 seconds.

Longer idle, sleep-risk, resume-risk, or shutdown-edge observation belongs to
the Session Watch Lantern or a later dedicated procedure. Do not let a boot
trace drift into unattended idle power-save testing.

## Minimum Captured Signals

The minimum Boot Trace Ledger should capture:

- Script start timestamp, end timestamp, elapsed duration, and status.
- UTC timestamp and local timestamp when available.
- Monotonic uptime from `/proc/uptime`.
- Kernel version from `uname -a`.
- Device model from `/proc/device-tree/model` when readable.
- Command availability and permission report.
- `vcgencmd get_throttled`, when available.
- `vcgencmd measure_temp`, when available.
- `vcgencmd measure_volts`, clearly labeled as internal/core rail evidence,
  not GPi Case 2 5V input rail evidence.
- Load average from `/proc/loadavg`.
- Memory summary from `/proc/meminfo` or `free`.
- Disk space for `/` and `/home` from `df`.
- Frontend detection for EmulationStation or known frontend process patterns.
- Framebuffer/display hints from safe read-only sources such as `/sys/class/graphics`,
  `/dev/fb*` existence checks, and narrow command output when available.
- Network/SSH reachability from the script perspective: whether the script is
  running under an SSH session when safely detectable, and the local hostname.
- Selected `dmesg` warning/error hints when safely available.
- Selected `journalctl -b` snippets when safely available.
- Known GPi Case 2 warning hints where available: `gpio12`, `vc4`, `v3d`,
  `audio`, `xpad`, `RCU`, `watchdog`, `mmc`, `ext4`, `voltage`, `throttle`,
  `under-voltage`, `drm`, `kms`, `dpi`, `framebuffer`, `usb`, `hid`, and
  `input`.

If `journalctl`, `dmesg`, `vcgencmd`, `systemd-analyze`, `free`, or another
allowlisted command is unavailable, the Ledger should record
`command_unavailable` and continue.

## Optional Captured Signals

Optional signals may be added only if they stay read-only, bounded, and
privacy-aware:

- `systemd-analyze time`, `blame`, or `critical-chain` when safely available.
- Boot ID from `/proc/sys/kernel/random/boot_id`.
- Focused process snapshots at start and end.
- First-seen frontend milestone during the capture.
- Narrow `/sys/class/drm` and `/sys/class/graphics` readable attributes.
- Narrow `lsusb` output for controller, audio, hub, or input clues.
- Short operator note supplied at start through `--note` or `--note-file`.
- Script checksum and copied path for field reproducibility.

Optional does not mean broad. The Lantern should prefer a smaller trustworthy
map over a sprawling satchel.

## Artifact Summary Expectations

Every final Ledger should include an `Artifact Summary` near the end.

The summary should report:

- Final status.
- Requested duration and observed duration.
- Sample count and interval.
- Start and end timestamps.
- First and last `/proc/uptime` values.
- Kernel version.
- Whether systemd boot timing was captured, unavailable, or permission-denied.
- Whether journal snippets were captured, unavailable, or permission-denied.
- Whether dmesg hints were captured, unavailable, or permission-denied.
- Frontend detection result: `detected`, `not_detected`, or `uncertain`.
- First frontend-detected sample when available.
- Display hint summary.
- Distinct raw `vcgencmd get_throttled` values.
- Temperature range when safely parsed.
- Memory, disk, and load high-water notes.
- SSH-side/script perspective: likely SSH session, local terminal, or
  unknown.
- Warning count and missing-evidence count.
- A cautious timing bucket if supported by the data, such as
  `early-userspace`, `display-handoff`, `frontend-start`, or `inconclusive`.

The summary must not replace raw rows. It is a map in the front pocket of the
satchel, not the whole satchel.

## Failure Modes

Failure modes should be recorded as field evidence:

| Failure mode | Ledger record |
| --- | --- |
| Command missing | `command_unavailable` with command name and affected field. |
| Permission denied | `permission_denied` with command or path and no mutating retry. |
| Command timeout | `command_timeout` with timeout seconds and affected field. |
| Empty output | `empty_output` with command name and timestamp. |
| Sampling interrupted | `interrupted` with signal name if known, elapsed time, and completed sample count. |
| Artifact write failed | `artifact_failed` with intended path and nearest preserved temporary path. |
| Disk space low | `low_space` with path and available space. |
| Clock shift | `time_shift_detected` with `/proc/uptime` preserved as ordering evidence. |
| SSH lost from operator view | Human note only unless the script can safely infer session loss; do not overclaim. |
| Frontend mismatch | `frontend_uncertain` with patterns checked and process snapshot references. |
| Handheld wedge suspected | Last successful sample, last progress line time, and final status if the script can write one. |

An interrupted or partial Ledger is still valuable if it names what was
captured and what went missing.

## Do Not Overclaim

Boot evidence is not later session evidence.

- A post-boot trace can show what the system reported after startup. It cannot
  prove what happened before the script started.
- `vcgencmd get_throttled` contains sticky flags. It can show that a condition
  has occurred since boot, but one value alone does not prove the exact moment,
  cause, power supply quality, battery health, or emulator performance impact.
- `vcgencmd measure_volts` is internal/core rail evidence, not actual GPi Case
  2 5V input rail voltage.
- Process detection is a clue. It is not proof of what the display showed.
- Journal and dmesg snippets are selected clues. They are not a complete
  kernel or service narrative.
- A clean boot trace does not prove sleep, resume, shutdown, or idle paths are
  safe.
- A successful post-resume artifact proves the system was responsive when the
  artifact was captured. It does not prove transition-time behavior unless a
  watcher was already running before sleep.

Use the Session Watch Lantern for later runtime evidence. Use human field
notes for display, LED, side switch, top button, audio, and control behavior.

## Relationship To Session Watch Lantern

The True Boot Trace Lantern and Session Watch Lantern should remain separate
tools:

- True Boot Trace Lantern: short startup-focused capture after normal boot,
  aimed at early userspace, display handoff, frontend start, and boot warning
  clues.
- Session Watch Lantern: bounded runtime observation after the handheld is
  responsive, aimed at menu, emulator, play, idle-risk, and post-resume state.

A future Field Lantern bundle may include both artifacts. The Common Problems
Mage can eventually compare the two Ledgers: boot buckets from the Boot Trace
Lantern and runtime buckets from the Session Watch Lantern. Neither Lantern
should apply fixes.

## Relationship To First Spark And Boot Veil

The
[First Spark / Boot Veil / Relic Welcome Scroll Design](gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md)
depends on True Boot Trace evidence before any startup UX behavior changes.
The observed roughly 15-second silent-screen window after the side switch is a
field observation, not yet proof of a KMS, framebuffer, firmware, or userspace
boundary.

The True Boot Trace Lantern should provide the timing, display, framebuffer,
frontend, warning, and missing-evidence clues needed before a future quest
chooses firmware/kernel config, framebuffer splash, Plymouth, LED/status
signal, or SSH welcome/status output. Startup polish must not hide diagnostic
boot text unless a documented recovery profile can reveal it again.

## Relationship To Future Arcadia Runtime Casters

Future Arcadia Runtime Casters may eventually own long-lived display, input,
power, dock, sleep, resume, shutdown, or diagnostics behavior. The True Boot
Trace Lantern is not a Caster.

Its role is evidence-gathering before trust:

- What startup looks like before any runtime Caster is installed.
- Which warnings appear during boot versus during later handheld use.
- Which commands are safely available on the actual Relic.
- Which startup signals should become acceptance gates for future Caster work.

No future Caster should treat this design as approval to mutate boot behavior.
Each Caster needs its own quest, tests, safety review, and field validation.

## Lantern Comparison

| Lantern | Current status | Primary window | Artifact | Main caution |
| --- | --- | --- | --- | --- |
| Bundle Collector Lantern | Current script | Post-boot responsive state | `.tar.gz` satchel plus report files | Gathers remembered boot clues after the fact; not a true startup recorder. |
| Boot Power Trace Lantern | Mapped future direction | Startup power-warning timing | Future timestamped trace | Focuses on undervoltage/throttling buckets; does not prove root cause. |
| Session Watch Lantern | Current foreground skeleton | Runtime menu, play, idle-risk, post-resume state | One final `.txt` Ledger | Complements boot evidence; does not prove transition-time behavior unless already running. |
| True Boot Trace Lantern | Current foreground skeleton | Short startup and frontend-start trail | One final `.txt` Boot Trace Ledger | Captures startup evidence without changing boot behavior; does not replace Session Watch. |

## Future Implementation Notes

QUEST-0088 adds the first tiny portable script skeleton. Later implementation
quests may deepen its evidence map only if they stay:

- POSIX-shell friendly where practical.
- `--duration`, `--interval`, `--output`, `--plain`, `--help`.
- Local-only, read-only, foreground.
- One final `.txt` Ledger.
- Visible progress every sample or at least every 10 seconds.
- Shell syntax checks and local smoke tests through `make check-scripts`.

Do not implement automatic boot start, systemd activation, GPIO capture,
shutdown behavior, sleep/resume behavior, repair suggestions, or Lantern
Dispatch in the first implementation. The first useful rune is a quiet,
honest Ledger.
