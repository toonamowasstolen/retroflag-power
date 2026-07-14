---
id: OPS-GPI-CASE-2-SESSION-WATCH-LANTERN-DESIGN-001
title: GPi Case 2 Session Watch Lantern Design
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Design a future read-only GPi Case 2 Session Watch Lantern for handheld runtime observation across menu, play, sleep-risk, and resume-risk field sessions.
related:
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - gpi-case-2-true-boot-trace-lantern-design.md
  - ../../scripts/gpi-case2-session-watch-lantern.sh
  - gpi-case-2-session-watch-field-run-procedure.md
  - gpi-case-2-session-watch-evidence-ledger.md
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-boot-power-trace-capture-procedure.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - human-facing-field-lantern-script-ux-standard.md
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
  - ../04-architecture/arcadia-runtime-migration-path.md
  - ../02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
last_updated: 2026-07-09
---

# GPi Case 2 Session Watch Lantern Design

> The Session Watch Lantern is the field lantern that stays awake with the
> handheld Relic: it watches the menu, game, input, power, and resume trail
> without touching the runes that make the device go.

This started as a documentation-only design. The first read-only foreground
script skeleton now lives at
[`scripts/gpi-case2-session-watch-lantern.sh`](../../scripts/gpi-case2-session-watch-lantern.sh).
That script writes one final text Ledger artifact for a bounded watch. It is
not a service, daemon, Arcadia Runtime Caster, installer, diagnostics bundle,
Lantern Dispatch path, GPIO reader, GPIO writer, shutdown action, suspend
action, resume action, firmware path, or automatic repair.

The GPi Case 2 is a handheld Relic. The primary field procedure must not
assume an attached keyboard. SSH to `retropi@gpi` is optional support for
copying, launching, or retrieving a satchel; it is not the primary handheld
experience.

The practical handheld procedure for running the current script is
[GPi Case 2 Session Watch Field Run Procedure](gpi-case-2-session-watch-field-run-procedure.md).
Use that page for command examples, physical observation notes, artifact
retrieval, and interpretation of clean, display-blank, SSH-lost, and hard
freeze-like runs.

Real handheld results belong in the
[GPi Case 2 Session Watch Evidence Ledger](gpi-case-2-session-watch-evidence-ledger.md).
That Ledger keeps the script artifact, display and LED notes, SSH state,
side-switch behavior, top-button behavior, EmulationStation status, and final
outcome together without turning observations into guesses.

## Purpose

Design a future read-only Session Watch Lantern that records what was visible
during a GPi Case 2 runtime session after boot. The lantern should help answer
what changed while the maintainer used menus, launched games, left the system
idle, or observed a successful post-resume state.

The Session Watch Lantern is different from the current Bundle Collector
Lantern and the future True Boot Trace Lantern:

- Bundle Collector Lantern: current manual post-boot satchel that gathers
  remembered logs and a short sampling window after the handheld is already
  responsive.
- True Boot Trace Lantern: future scp-first startup recorder for boot,
  display handoff, frontend-start, and warning buckets.
- Session Watch Lantern: current foreground skeleton and future fuller runtime
  observer for menu, emulator, game, idle-risk, and post-resume evidence after
  boot.

The first implementation is a small foreground shell script that writes one
final artifact file. It does not require long terminal paste. It follows the
[Human-Facing Field Lantern Script UX Standard](human-facing-field-lantern-script-ux-standard.md)
at skeleton scale: startup banner, bounded duration, progress lines, plain
mode, `NO_COLOR`, safety marker, and exact artifact path.

## Non-Goals

The Session Watch Lantern must not:

- Diagnose every possible GPi Case 2 issue by itself.
- Claim a single root cause from one runtime trace.
- Start automatically at boot in the first quest that implements it.
- Install or activate systemd.
- Read GPIO or write GPIO.
- Execute shutdown, reboot, halt, poweroff, suspend, sleep, or resume.
- Press, simulate, or depend on the top sleep/resume button.
- Replace `/opt/RetroFlag/SafeShutdown.py`.
- Edit `/etc/rc.local`, `/boot/config.txt`, KMS, display, audio, emulator, or
  RetroPie configuration.
- Run `lcdfirst.sh`, `lcdnext.sh`, RetroFlag installers, or firmware tools.
- Upload, submit, or contact Lantern Dispatch.
- Collect ROM names, save data, screenshots, Wi-Fi details, tokens, private
  keys, broad environment dumps, or arbitrary journals.

## Field Questions

The Session Watch Lantern should answer these field questions with evidence
and uncertainty:

| Field question | Useful answer shape |
| --- | --- |
| Was the handheld responsive throughout the watched session? | Start, progress, final status, and any missing-sample gaps. |
| Was EmulationStation, another frontend, or an emulator visible to process checks? | `detected`, `not_detected`, or `uncertain`, with patterns checked. |
| Did `vcgencmd get_throttled` change during menu or play? | First value, last value, distinct values, and timestamped changes. |
| Did temperature, load, memory, or disk pressure climb during play or idle? | Sample table plus high-water marks. |
| Did input, USB audio, display, KMS, MMC, ext4, RCU, watchdog, or power warnings appear? | Narrow matching excerpts with timestamps or monotonic uptime. |
| Was the device approaching idle or power-save risk? | Elapsed time, last observed input hint when available, and operator notes. |
| Did the handheld resume successfully before the lantern was run or while it was already watching? | Clearly separate post-resume observation from transition-time evidence. |
| What was not captured? | Missing command, missing file, permission, timeout, or interrupted-run markers. |

The lantern should preserve `Unknown` and `Inconclusive` rather than turning a
thin trail into a confident tale.

## Safety Boundaries

The Session Watch Lantern must stay read-only, local, foreground, and
operator-driven.

Required safety boundaries:

- The side switch is the normal stock shutdown path while the GPi Case 2 is
  responsive.
- The top sleep/resume button is suspect unless a procedure explicitly says
  otherwise.
- Do not induce sleep or resume as part of the basic watch.
- Do not keep the handheld idle long enough to trigger risky automatic
  power-save during early field work.
- Do not use the watch as emergency recovery during a kernel stall, lost
  network incident, or visible RCU stall.
- If the handheld wedges, stop the watch attempt and follow the
  [GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md).
- Do not treat a successful post-resume artifact as proof of transition-time
  behavior unless the watcher was already running before sleep.
- Do not change display, audio, boot, runtime, SafeShutdown, or RetroFlag files
  during the same pass.

The satchel may record local state, but it must never perform a hardware
action.

## Expected Handheld Procedure

The handheld path should be designed for a maintainer using the GPi Case 2 as
a handheld first.

Expected procedure:

1. Boot the GPi Case 2 normally.
2. Keep the handheld responsive and avoid the top sleep/resume button.
3. If SSH support is available, copy one portable Session Watch Lantern Relic
   to `/home/retropi/` on `retropi@gpi`.
4. Start the watch with an explicit duration, such as 10 to 15 minutes for a
   first menu/play trail, or a shorter smoke test.
5. Use the handheld normally: menu navigation, a known game launch, a brief
   play session, or supervised idle-risk observation.
6. Stop before expected automatic power-save unless a later procedure
   explicitly approves that test.
7. Let the lantern write one final artifact and print its exact path.
8. Retrieve the artifact with SSH support when available, or leave it on the
   device until the next safe retrieval path.
9. Shut down with the side switch only when the system remains responsive and
   field work is complete.

If the future implementation supports an operator note, it should accept a
short text note at start or through an optional file, not require typing on an
attached keyboard during handheld use.

## Artifact Pattern

The current skeleton writes one final timestamped text Ledger by default:

```text
/home/retropi/gpi-case2-session-watch-lantern-YYYYMMDD-HHMMSS.txt
```

Use `--output FILE` to choose a different final Ledger path.

A later fuller implementation may write a timestamped folder and one final
portable archive beside it:

```text
/home/retropi/gpi-case2-session-watch-lantern-YYYYMMDD-HHMMSS/
/home/retropi/gpi-case2-session-watch-lantern-YYYYMMDD-HHMMSS.tar.gz
```

The final terminal output should name the artifact exactly:

```text
Artifact: /home/retropi/gpi-case2-session-watch-lantern-20260709-121500.txt
Duration: 00:10:03
Status:   completed
```

For the current skeleton, the `.txt` Ledger is the single file a maintainer can
pull into the wider Ledger or attach beside a Field Lantern bundle after
review. For a future fuller satchel, the folder should be inspectable before
sharing and the `.tar.gz` should be the single shared file.

The current `.txt` Ledger keeps the per-sample rows intact and adds a final
`Artifact Summary` after the warnings section. That summary is a human map for
field interpretation: completion status, requested and observed duration,
sample count, parsed temperature range when available, frontend detection,
first frontend-detected sample, script-view SSH-side completion, distinct raw
`get_throttled` values, a cautious throttling note, and warning or
missing-evidence count. It must not turn raw throttled values into a power,
battery, charger, or emulator-performance claim by itself.

## Progress And Status Output

Long-running Session Watch Lantern runs must show life while they work:

- Startup banner with the Relic name, duration, and read-only promise.
- Clear `[[n/m]]` stage labels.
- Progress during sampling, such as elapsed time, remaining time, sample count,
  and latest status.
- Timing for setup, sampling, artifact writing, and total duration.
- Final artifact path and folder path.
- Final `Artifact Summary` inside the Ledger so future field readers can
  orient themselves before reading every checkpoint row.
- Final status: `completed`, `completed_with_warnings`, `interrupted`, or
  `failed`.
- A short reminder to inspect and redact before sharing.

The implementation must respect `--plain` and `NO_COLOR`. Plain mode should
use ASCII-only output, no color, no terminal animation, and stable copyable
lines.

The lantern should never require a long terminal paste as its main result. The
terminal is the campfire view; the artifact is the Ledger entry.

## Minimum Signals To Capture

The minimum useful satchel should include:

- UTC start time, local start time when available, end time, and total
  duration.
- Monotonic uptime from `/proc/uptime`.
- Sample interval and intended duration.
- Per-sample `vcgencmd get_throttled`, when available.
- Per-sample `vcgencmd measure_temp`, when available.
- Per-sample `vcgencmd measure_volts`, clearly labeled as internal/core rail
  evidence, not GPi Case 2 5V input rail evidence.
- Load average from `/proc/loadavg`.
- Memory summary from `/proc/meminfo` or `free`, with narrow fields.
- Disk space for relevant writable paths, such as `/home` or `/`.
- Best-effort frontend and emulator process detection, including checked
  patterns and result policy.
- Narrow matching `dmesg` excerpts for voltage, throttle, KMS, DRM, VC4, DPI,
  USB, HID, input, audio, xpad, MMC, ext4, RCU, watchdog, suspend, resume, and
  power-save wording.
- Best-effort journal excerpts with the same narrow matches when `journalctl`
  is available.
- Process snapshots at start and end, plus any first-seen frontend or emulator
  milestone.
- Command availability report.
- Capture manifest with privacy, redaction, and non-goal notes.

If a command is missing, times out, or lacks permission, record that as a
field fact. Do not install packages or broaden the satchel to compensate.

## Optional SSH-Assisted Signals

SSH support may add convenience, but the watch must remain useful without
treating SSH as the main handheld experience.

Optional SSH-assisted signals:

- Copy path and script checksum used for the Relic handoff.
- SSH start time from the workstation perspective.
- Retrieval command printed at the end.
- Operator-supplied note file copied beside the artifact.
- Workstation-side receipt path after `scp`, if the future helper prints one.

SSH-assisted signals should not include private network inventory, broad shell
history, secrets, keys, or workstation facts. The field target remains
`retropi@gpi` as optional support.

## Failure Modes

The artifact should record failure modes as first-class evidence:

| Failure mode | Artifact record |
| --- | --- |
| Command missing | `command_unavailable` with command name and affected field. |
| Command timeout | `command_timeout` with timeout seconds and partial output path, if any. |
| Permission denied | `permission_denied` with path or command and no retry that mutates state. |
| Sampling interrupted | `interrupted` with completed sample count, elapsed time, and signal name if known. |
| Artifact creation failed | `artifact_failed` with folder path preserved when possible. |
| Disk space low | `low_space` warning with available bytes and path. |
| Time changed during run | `time_shift_detected` with monotonic uptime preserved as the trusted ordering clue. |
| Handheld wedge suspected | `watch_lost_visibility` or `samples_stopped` with last successful sample time. |
| Resume observed only after the fact | `post_resume_observed_only`, not transition captured. |
| Frontend detection mismatch | `frontend_uncertain` with patterns checked and process snapshots. |

The final report should include a concise summary, a warning section, and a
missing-evidence section. Missing evidence is part of the map.

## Relationship To Boot Trace Lantern

The
[True Boot Trace Lantern](gpi-case-2-true-boot-trace-lantern-design.md)
answers startup timing questions. The Session Watch Lantern answers runtime
session questions.

The two lanterns should not duplicate each other:

- Boot Trace Lantern: starts early enough to time boot, KMS/display init,
  USB/audio/controller init, EmulationStation startup, and early power-warning
  buckets.
- Session Watch Lantern: starts after the handheld is responsive and watches
  menu, emulator, play, idle-risk, and post-resume state.

A later Field Lantern bundle may include both satchels as separate sections.
The Common Problems Mage can eventually read both: Boot Trace buckets for
startup timing and Session Watch buckets for runtime behavior.

The
[First Spark / Boot Veil / Relic Welcome Scroll Design](gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md)
uses Session Watch as the post-boot verification companion. If a future veil
or welcome scroll changes startup presentation, Session Watch should help
confirm that the handheld remains understandable and responsive once
EmulationStation or another frontend is visible.

## Relationship To Arcadia Runtime Casters

Future Arcadia Runtime Casters may eventually own long-lived power, display,
dock, resume, or input duties. The Session Watch Lantern is not one of those
Casters.

The Session Watch Lantern should remain a read-only field instrument that
teaches the Spellbook what a real handheld session looks like before a Caster
is trusted with runtime work. It can provide evidence for future Caster
acceptance gates:

- What a normal menu/play session looks like.
- What temperature, load, memory, and throttle trails look like during play.
- What successful post-resume state looks like.
- Which gaps still need safer observation before Caster behavior is proposed.

No future Caster should use this design as approval to mutate GPi Case 2
behavior. Each Caster needs its own quest, tests, safety review, and field
validation.

## Future Implementation Notes

QUEST-0082 started tiny with
[`scripts/gpi-case2-session-watch-lantern.sh`](../../scripts/gpi-case2-session-watch-lantern.sh):

- Portable foreground shell script.
- Explicit `--duration`, `--interval`, and `--output`.
- `--plain` plus `NO_COLOR` support.
- One final `.txt` Ledger artifact.
- Shell syntax and help smoke checks through `make check-scripts`.
- No systemd, no auto-start, no GPIO, no shutdown, no sleep/resume trigger, no
  upload, and no repair advice.

QUEST-0083 adds the handheld-first field run procedure for the current
skeleton:
[GPi Case 2 Session Watch Field Run Procedure](gpi-case-2-session-watch-field-run-procedure.md).
That page is the operator Spellbook for real runs; this design remains the
map of intent, boundaries, and future growth.

A later fuller implementation can grow from that skeleton:

- One final `.tar.gz` artifact.
- Unit tests or shell tests for option parsing, artifact naming, plain output,
  missing-command markers, and failure summaries where practical.

The first badge is not a clever daemon. The first badge is a trustworthy field
satchel that tells the next maintainer exactly what it saw and exactly what it
missed.
