---
id: OPS-GPI-CASE-2-SESSION-WATCH-FIELD-RUN-PROCEDURE-001
title: GPi Case 2 Session Watch Field Run Procedure
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define the handheld-first field procedure for running the read-only GPi Case 2 Session Watch Lantern script during a real responsive session.
related:
  - ../../scripts/gpi-case2-session-watch-lantern.sh
  - gpi-case-2-session-watch-lantern-design.md
  - gpi-case-2-session-watch-evidence-ledger.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - gpi-case-2-boot-power-trace-capture-procedure.md
  - human-facing-field-lantern-script-ux-standard.md
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
last_updated: 2026-07-09
---

# GPi Case 2 Session Watch Field Run Procedure

> This Field Lantern run watches a living handheld session: the Relic in your
> hands, a bounded trail, one final Ledger, and no changes to the path it maps.

This procedure is for the current read-only script:
[`scripts/gpi-case2-session-watch-lantern.sh`](../../scripts/gpi-case2-session-watch-lantern.sh).
It is a foreground Lantern Relic for a responsive GPi Case 2. It writes one
final text Ledger artifact and prints progress, timing, status, and the final
artifact path while it works.

After the run, record the artifact and concise human observations in the
[GPi Case 2 Session Watch Evidence Ledger](gpi-case-2-session-watch-evidence-ledger.md).

The GPi Case 2 remains a handheld Relic first. Do not assume an attached
keyboard. SSH to `retropi@gpi` is optional support for copying, launching, or
retrieving the artifact; it is not the primary user experience. While the
handheld is responsive, the side switch remains the normal stock shutdown
path. The top sleep/resume button remains suspect unless a future procedure
explicitly says otherwise.

## Purpose

Run the Session Watch Lantern during a real menu, game, idle-risk, or
post-resume observation window so a maintainer can preserve what happened
without changing display, power, shutdown, GPIO, SafeShutdown, RetroPie, or
systemd behavior.

The Ledger should help answer:

- Did the handheld remain responsive during the watched window?
- Did the display stay lit, blank, recover, or freeze?
- Did SSH remain alive when optional support was connected?
- Did the LED blink state, side switch behavior, or top button behavior line up
  with the human observation?
- Did the script finish cleanly and print the final artifact path?

## When To Run It

Run this procedure when:

- The GPi Case 2 boots normally and is responsive in handheld use.
- A maintainer wants a short live watch while navigating EmulationStation,
  launching a known game, playing briefly, or observing supervised idle risk.
- A previous sleep/resume or display-blank incident has recovered and the next
  step is read-only evidence, not repair.
- Optional SSH is available for convenience, or the script is already present
  on the handheld and can be launched safely.

Do not use this as emergency recovery. If the handheld is already frozen,
printing repeated kernel stalls, losing both SSH and ping, or stuck at a blank
display with no useful input response, follow the
[GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
first.

## Safety Boundaries

This Field Lantern run must remain read-only and local.

- Do not read GPIO or write GPIO.
- Do not execute shutdown, reboot, halt, poweroff, suspend, sleep, or resume.
- Do not edit `/boot/config.txt`, `/etc/rc.local`, RetroPie files, emulator
  config, systemd units, display config, audio config, or
  `/opt/RetroFlag/SafeShutdown.py`.
- Do not run `lcdfirst.sh`, `lcdnext.sh`, RetroFlag installers, firmware tools,
  or package installers.
- Do not contact Lantern Dispatch, upload telemetry, or auto-submit the
  Ledger.
- Do not press the top sleep/resume button during the run unless a later
  dedicated procedure explicitly instructs that test.
- If the handheld remains responsive and field work is complete, use the side
  switch for the normal stock shutdown path.

## Prerequisites

- A responsive GPi Case 2 with enough battery or stable USB-C power for the
  planned watch length.
- The current script copied onto the device as one portable Relic. The GPi
  Case 2 does not need the full repository checked out for current field use.

```sh
scp scripts/gpi-case2-session-watch-lantern.sh retropi@gpi:/home/retropi/
ssh retropi@gpi
sh /home/retropi/gpi-case2-session-watch-lantern.sh --plain --duration 300 --interval 15
```

- Optional SSH support from a workstation:

```sh
ssh retropi@gpi
```

- A simple human field note surface: paper, phone note, or workstation text
  file. The script records machine observations; the human records the display,
  LED, buttons, and physical state.

## Recommended Commands

Short handheld smoke run, useful before a real pass:

```sh
sh /home/retropi/gpi-case2-session-watch-lantern.sh --duration 60 --interval 10
```

Normal watch run, useful for menu navigation, one game launch, brief play, and
supervised return to the menu:

```sh
sh /home/retropi/gpi-case2-session-watch-lantern.sh --duration 600 --interval 15
```

Plain copyable output:

```sh
sh /home/retropi/gpi-case2-session-watch-lantern.sh --plain --duration 300 --interval 15
```

No-color output through the environment:

```sh
NO_COLOR=1 sh /home/retropi/gpi-case2-session-watch-lantern.sh --duration 300 --interval 15
```

Explicit Ledger path, useful when running several passes:

```sh
sh /home/retropi/gpi-case2-session-watch-lantern.sh \
  --duration 600 \
  --interval 15 \
  --output /home/retropi/session-watch-menu-play-001.txt
```

## Handheld Field Run

1. Boot the GPi Case 2 normally.
2. Confirm the handheld is responsive in EmulationStation or the intended
   frontend.
3. Confirm the display is visible and note the LED blink state before starting.
4. Start the script with one of the recommended commands.
5. Watch the terminal long enough to confirm the Lantern banner, read-only
   safety line, step lines, progress, and artifact destination appear.
6. Use the handheld normally for the planned pass: menu navigation, one known
   game launch, brief play, return to menu, or supervised idle observation.
7. Do not press the top sleep/resume button during the run.
8. If the display blanks, SSH drops, or the handheld appears frozen, write down
   the exact time and physical state before touching anything.
9. Let the script finish if the device remains responsive.
10. Copy the final `Artifact:` path from the terminal into the field note.
11. Retrieve the Ledger when optional SSH is available, then inspect it before
    sharing.
12. When finished and responsive, shut down through the normal side-switch
    path.

## What To Physically Observe

Record these human observations beside the script Ledger:

| Observation | What to write down |
| --- | --- |
| Display state | `lit`, `blank`, `dimmed`, `frozen image`, `console text`, `recovered`, or `unknown`, plus time seen. |
| LED blink state | Blink pattern, steady state, off state, or change from the start state. |
| SSH state | `not used`, `alive`, `lost`, `recovered`, or `unknown`, plus last known time. |
| Handheld controls | Whether D-pad/buttons still move the menu or game. |
| Side switch behavior | Whether the normal side-switch shutdown path worked after the run while responsive. |
| Top button behavior | Usually `not pressed`. If accidentally pressed, record time, display result, and whether SSH stayed alive. |
| Sound or game state | Whether audio continued, game kept running, or emulator/frontend appeared stuck. |

Use simple field-note lines:

```text
12:10 local - Start: display lit, LED blinking normally, SSH alive.
12:14 local - Launched known game; controls responsive.
12:18 local - Display blank; audio continued; SSH alive.
12:20 local - Returned with side switch not used yet; artifact path printed.
12:22 local - Side switch shutdown behaved normally.
Top button: not pressed.
```

## What Not To Do During The Run

- Do not test the top sleep/resume button during this general watch.
- Do not leave the handheld unattended into a risky automatic power-save
  window.
- Do not start a second Lantern or bundle collector at the same time.
- Do not change games, overlays, display settings, audio settings, or scripts
  as part of the same observation.
- Do not turn a blank display into an immediate repair attempt. First record
  display, LED, input, audio, and optional SSH state.
- Do not paste the whole terminal transcript as the main artifact. Use the
  final Ledger file.

## Expected Artifact Location

By default the script writes one final timestamped text Ledger under the
`retropi` home directory:

```text
/home/retropi/gpi-case2-session-watch-lantern-YYYYMMDD-HHMMSS.txt
```

If `--output FILE` is used, the final Ledger is written to that exact path.
The terminal output should include an exact final artifact path. Treat that
path as the map rune for retrieval and sharing.

The Ledger also ends with an `Artifact Summary` section. Use it as the quick
map for completed or interrupted status, requested versus observed duration,
sample count, temperature range when available, first frontend detection,
whether the SSH-side watch completed normally from the script's point of view,
raw throttled values observed, and the count of warnings or missing evidence.
The throttling note is deliberately cautious: preserve raw `vcgencmd
get_throttled` values for later interpretation, but do not claim power source,
battery, charger, or emulator-performance cause from the summary alone.

## Retrieve And Share The Ledger

From a workstation with optional SSH support, pull the exact path printed by
the script:

```sh
scp retropi@gpi:/home/retropi/gpi-case2-session-watch-lantern-YYYYMMDD-HHMMSS.txt .
```

For an explicit output path:

```sh
scp retropi@gpi:/home/retropi/session-watch-menu-play-001.txt .
```

Before sharing:

- Open the Ledger locally.
- Confirm it says `READ-ONLY / NO CHANGES MADE`.
- Check the final status, `Artifact Summary`, warnings, missing evidence, and
  artifact path.
- Add or update the matching entry in the
  [GPi Case 2 Session Watch Evidence Ledger](gpi-case-2-session-watch-evidence-ledger.md).
- Add the human field notes beside it or below it.
- Redact anything private before sending it to another maintainer.

Lantern Dispatch is not implemented and is not part of this procedure. Sharing
is manual and user-controlled.

## Interpret The Run

| Run shape | Meaning | Next read-only step |
| --- | --- | --- |
| Clean run | The script prints progress until completion, the display remains usable, SSH support stays alive if used, and the final Ledger path exists. | Save the Ledger and field notes as the baseline for this menu/play state. |
| Sleep/display-blank run | The display blanks, dims, or freezes while the script may still finish or SSH may stay alive. | Record display, LED, audio, input, and SSH state. Do not call it a sleep transition unless the top button or power-save path is proven. |
| SSH-lost run | Optional SSH disconnects or stops responding, but the handheld may still show a display or react locally. | Record whether local controls and audio still work. Retrieve the Ledger later only after the device is safely responsive again. |
| Hard-freeze or RCU-stall-like run | Display, input, and SSH stop making progress, or visible kernel stall text appears. | Stop the watch attempt as an observation. Follow the recovery-first procedure and preserve photos, times, LED state, and last artifact path if visible. |

A successful post-resume watch proves only that the handheld was observable
after resume. It does not prove the sleep/resume transition itself unless the
Lantern was already running before that transition under a later approved
procedure.

## Pairing With Other Ledgers

The Session Watch Ledger pairs well with:

- The evidence Ledger:
  [GPi Case 2 Session Watch Evidence Ledger](gpi-case-2-session-watch-evidence-ledger.md).
- The design map:
  [GPi Case 2 Session Watch Lantern Design](gpi-case-2-session-watch-lantern-design.md).
- The wider manual bundle trail:
  [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md).
- Future classification notes:
  [Common Problems Mage Map](common-problems-mage-map.md).
- Future local bundle boundaries:
  [Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md).

Keep the machine Ledger and the human note together. The script sees process,
kernel, and system clues; the human sees the Relic's screen, LED, buttons, and
body language.
