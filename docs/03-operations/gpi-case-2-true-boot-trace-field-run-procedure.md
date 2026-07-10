---
id: OPS-GPI-CASE-2-TRUE-BOOT-TRACE-FIELD-RUN-PROCEDURE-001
title: GPi Case 2 True Boot Trace Field Run Procedure
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define the handheld-first, scp-first field procedure for running the read-only True GPi Case 2 Boot Trace Lantern and retrieving its final Boot Trace Ledger.
related:
  - ../../scripts/gpi-case2-true-boot-trace-lantern.sh
  - gpi-case-2-true-boot-trace-lantern-design.md
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-session-watch-field-run-procedure.md
  - gpi-case-2-recovery-first-field-procedure.md
  - human-facing-field-lantern-script-ux-standard.md
last_updated: 2026-07-09
---

# GPi Case 2 True Boot Trace Field Run Procedure

> This Lantern run watches startup from the human side and the shell side:
> side switch, silent window, first visible text, one bounded trace, and one
> final Boot Trace Ledger carried home with `scp`.

This procedure is for the current read-only script:
[`scripts/gpi-case2-true-boot-trace-lantern.sh`](../../scripts/gpi-case2-true-boot-trace-lantern.sh).
The GPi Case 2 does not need the full repository checked out. Copy this one
portable Relic to `retropi@gpi:/home/retropi/`, run it from `/home/retropi/`,
and retrieve the final Ledger artifact with `scp`.

The GPi Case 2 is a handheld Relic first. Do not assume an attached keyboard.
SSH to `retropi@gpi` is optional support for copying, launching, and retrieving
the artifact; it is not the primary handheld UX. While the Relic is
responsive, the side switch remains the normal stock shutdown path. The top
sleep/resume button remains suspect unless a dedicated procedure explicitly
says otherwise.

## Purpose

Run the True Boot Trace Lantern during a normal startup observation so a
maintainer can preserve the startup trail without changing boot, display,
power, shutdown, sleep/resume, GPIO, RetroPie, EmulationStation, or systemd
behavior.

The run should pair two evidence streams:

- Human notes: side-switch time, first visible display time, display state,
  LED state, SSH availability, EmulationStation visibility, and side-switch
  behavior while responsive.
- Boot Trace Ledger: bounded script observations from the shell side,
  including uptime, selected boot hints, display/framebuffer clues, frontend
  detection, `vcgencmd` values when available, warnings, missing evidence, and
  final artifact summary.

## When To Run It

Run this procedure when:

- The next useful question is startup timing or first-visible-display timing.
- The maintainer wants evidence before any First Spark, Boot Veil, Relic
  Welcome Scroll, display, KMS, framebuffer, or boot-text change.
- The GPi Case 2 can be physically watched during boot.
- Optional SSH support is available soon after boot, or the script has already
  been copied to `/home/retropi/`.
- The operator can record simple timestamps without disturbing the handheld.

Do not use this as emergency recovery. If the handheld is already frozen,
printing repeated kernel stalls, losing both SSH and ping, or stuck at a blank
display with no useful input response, follow the
[GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
first.

## Safety Boundaries

This Boot Trace run must remain read-only, local, foreground, and bounded.

- Do not edit `/boot/config.txt`, `/boot/cmdline.txt`, `/etc/rc.local`,
  RetroPie files, emulator config, systemd units, display config, audio
  config, or `/opt/RetroFlag/SafeShutdown.py`.
- Do not read GPIO or write GPIO.
- Do not execute shutdown, reboot, halt, poweroff, suspend, sleep, or resume
  during the capture.
- Do not run `lcdfirst.sh`, `lcdnext.sh`, RetroFlag installers, firmware
  tools, package installers, or repair commands.
- Do not contact Lantern Dispatch, upload telemetry, or auto-submit the
  Ledger.
- Do not press the top sleep/resume button during the run.
- Do not treat the side switch as a guaranteed emergency path during a kernel
  stall. Use it as the normal stock shutdown path only while responsive.
- If the handheld wedges, stop interpreting and preserve what was last seen.

## Prerequisites

- Local repo checkout on the workstation.
- Network path to `retropi@gpi`.
- GPi Case 2 battery or USB-C power suitable for one short supervised run.
- A human timing surface: stopwatch, phone note, paper note, or workstation
  note file.
- The script copied to the handheld as one portable Relic.

Copy from the local repo to the GPi Case 2:

```sh
scp scripts/gpi-case2-true-boot-trace-lantern.sh retropi@gpi:/home/retropi/
```

Make the copied Relic executable for convenience:

```sh
ssh retropi@gpi chmod +x /home/retropi/gpi-case2-true-boot-trace-lantern.sh
```

Open optional SSH support when the Relic is responsive:

```sh
ssh retropi@gpi
```

## Recommended Commands

Short smoke run, useful after copying the script:

```sh
sh /home/retropi/gpi-case2-true-boot-trace-lantern.sh --duration 30 --interval 5
```

Normal boot-trace run:

```sh
sh /home/retropi/gpi-case2-true-boot-trace-lantern.sh --duration 120 --interval 5
```

Plain copyable output:

```sh
sh /home/retropi/gpi-case2-true-boot-trace-lantern.sh --plain --duration 120 --interval 5
```

No-color output through the environment:

```sh
NO_COLOR=1 sh /home/retropi/gpi-case2-true-boot-trace-lantern.sh --duration 120 --interval 5
```

Explicit Ledger path, useful when running several passes:

```sh
sh /home/retropi/gpi-case2-true-boot-trace-lantern.sh \
  --duration 120 \
  --interval 5 \
  --output /home/retropi/true-boot-trace-first-spark-001.txt
```

## Handheld Field Run

1. Keep the GPi Case 2 physically visible.
2. Start a stopwatch or write down the local time.
3. Flip the side power switch on.
4. Record the exact power-switch timestamp if available.
5. Watch the display without pressing the top sleep/resume button.
6. Record the first visible display timestamp or rough seconds after the
   switch.
7. Record whether the first visible display is blank/backlit, text, splash or
   art, EmulationStation, or unknown.
8. When optional SSH becomes available, connect to `retropi@gpi`.
9. Start the Boot Trace Lantern from `/home/retropi/`.
10. Watch for the banner, read-only safety line, progress lines, and artifact
    destination.
11. Keep the handheld in view until the bounded capture finishes.
12. Record whether EmulationStation becomes visible and whether controls appear
    responsive.
13. Record LED state, SSH availability, and any side-switch behavior observed.
14. Copy the final `Artifact:` path from the terminal into the field note.
15. Retrieve the Ledger with `scp`.
16. If field work is complete and the system is responsive, shut down through
    the normal side-switch path.

## First Spark Observation

Use this section for the startup moment itself. Do not assume the cause of any
blank or quiet window.

```text
Run label:
Power source:
Power switch flipped:
First visible display:
First visible display rough seconds after switch:
First visible display state: blank | text | splash/art | EmulationStation | unknown
First SSH availability:
First EmulationStation visible:
Quiet window notes:
```

The operator has observed a roughly 15-second silent or blank window after
flipping the side power switch before first visible text. Record whether this
run shows a similar window, a shorter window, a longer window, or no clear
measurement. Do not claim KMS, framebuffer, firmware, userspace, panel, or
case-board cause from this observation alone.

Example:

```text
Power switch flipped: 12:10:00 local
First visible display: about 12:10:15 local
First visible display rough seconds after switch: about 15 seconds
First visible display state: text
First SSH availability: 12:10:42 local
First EmulationStation visible: 12:10:58 local
Quiet window notes: screen appeared blank or inactive until first text; cause unknown.
```

## What To Physically Observe

Record these human observations beside the Boot Trace Ledger:

| Observation | What to write down |
| --- | --- |
| First visible screen time | Exact local timestamp or rough seconds after side-switch flip. |
| Display state | `blank`, `text`, `splash/art`, `EmulationStation`, `frozen image`, `recovered`, or `unknown`, plus time seen. |
| LED state | Blink pattern, steady state, off state, or change from the start state. |
| SSH state | `not used`, `not yet available`, `alive`, `lost`, `recovered`, or `unknown`, plus first known time. |
| EmulationStation state | First visible time, not visible, or unknown. |
| Handheld controls | Whether D-pad/buttons appear responsive once the frontend is visible. |
| Side switch behavior | Whether the normal side-switch shutdown path worked after the run while responsive. |
| Top button behavior | Usually `not pressed`. If accidentally pressed, record time, display result, and whether SSH stayed alive. |

Use simple field-note lines:

```text
12:10:00 local - Side switch flipped on; LED state unknown; display blank.
12:10:15 local - First visible text; about 15 seconds after switch.
12:10:42 local - SSH available; launched Boot Trace Lantern.
12:10:58 local - EmulationStation visible; controls responsive.
12:12:44 local - Lantern completed; artifact path printed.
12:13:10 local - Side-switch shutdown behaved normally while responsive.
Top button: not pressed.
```

## What Not To Do During The Run

- Do not press the top sleep/resume button.
- Do not attach assumptions to the quiet window while recording it.
- Do not start a second Lantern or bundle collector at the same time.
- Do not change boot, display, audio, power, RetroPie, or EmulationStation
  settings as part of the same observation.
- Do not run installers, firmware tools, package installers, or repair
  commands.
- Do not leave the handheld unattended into a risky automatic power-save
  window.
- Do not treat a blank display as an immediate repair attempt. First record
  display, LED, input, audio if relevant, and optional SSH state.
- Do not paste the whole terminal transcript as the main artifact. Use the
  final Boot Trace Ledger file.

## Expected Artifact Location

By default the script writes one final timestamped text Ledger under the
`retropi` home directory:

```text
/home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS.txt
```

If `--output FILE` is used, the final Ledger is written to that exact path.
The terminal output should include an exact final artifact path:

```text
Artifact: /home/retropi/gpi-case2-true-boot-trace-lantern-20260709-121500.txt
```

Treat that `Artifact:` line as the source path for retrieval.

## Retrieve The Ledger

From the workstation, pull the exact path printed by the script:

```sh
scp retropi@gpi:/home/retropi/gpi-case2-true-boot-trace-lantern-YYYYMMDD-HHMMSS.txt .
```

For an explicit output path:

```sh
scp retropi@gpi:/home/retropi/true-boot-trace-first-spark-001.txt .
```

Before sharing:

- Open the Ledger locally.
- Confirm it says `READ-ONLY / NO CHANGES MADE`.
- Check the final status, `Artifact Summary`, warnings, missing evidence, and
  artifact path.
- Keep the human First Spark notes beside the Ledger. The script cannot see
  the exact power-switch moment or the first visible display moment by itself.
