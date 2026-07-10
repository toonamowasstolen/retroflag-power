---
id: OPS-GPI-CASE-2-FIRST-SPARK-BOOT-VEIL-WELCOME-SCROLL-DESIGN-001
title: GPi Case 2 First Spark / Boot Veil / Relic Welcome Scroll Design
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Design a recovery-first GPi Case 2 startup UX plan for earlier operator feedback, safer boot-text hiding, and an SSH welcome/status scroll without implementing boot config changes.
related:
  - gpi-case-2-relic-welcome-scroll-design.md
  - gpi-case-2-true-boot-trace-lantern-design.md
  - gpi-case-2-true-boot-trace-field-run-procedure.md
  - gpi-case-2-true-boot-trace-evidence-ledger.md
  - gpi-case-2-session-watch-lantern-design.md
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-recovery-first-field-procedure.md
  - gpi-case-2-session-watch-field-run-procedure.md
  - gpi-case-2-session-watch-evidence-ledger.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
last_updated: 2026-07-10
---

# GPi Case 2 First Spark / Boot Veil / Relic Welcome Scroll Design

> The handheld should answer the side-switch flip quickly, hide rough boot text
> only where recovery remains easy, and greet SSH operators with useful status
> instead of a blank hallway.

This is a design document only. It does not approve or implement changes to
firmware config, kernel command line, systemd units, Plymouth, framebuffer
services, RetroPie, EmulationStation, `SafeShutdown.py`, GPIO behavior,
shutdown, sleep, or resume.

The GPi Case 2 is a handheld Relic. Do not assume an attached keyboard. The
side switch remains the normal stock shutdown path while the system is
responsive. The top sleep/resume button remains suspect unless a dedicated
procedure explicitly tests it. Current field practice is scp-first: copy a
small tool to `retropi@gpi:/home/retropi/`, run it from there, and retrieve the
artifact with `scp`.

Use the
[GPi Case 2 True Boot Trace Field Run Procedure](gpi-case-2-true-boot-trace-field-run-procedure.md)
to gather the handheld First Spark observations this design depends on:
power-switch time, first visible display time, first SSH availability, first
EmulationStation visibility, and quiet-window notes.
Returned artifacts and observations belong in the
[GPi Case 2 True Boot Trace Evidence Ledger](gpi-case-2-true-boot-trace-evidence-ledger.md)
before this design graduates any startup UX guess into a future implementation
rune.

## Purpose

Design a recovery-first startup UX plan with three separate layers:

| Layer | Goal | First safe output |
| --- | --- | --- |
| First Spark | Earliest possible visible or physical feedback after the side power switch is flipped. | LED, backlight, display flash, or other unmistakable "power path is alive" sign. |
| Boot Veil | Hide rough boot text and penguin graphics after the display path is active, where safe. | Splash, framebuffer image, or service-managed cover that can be disabled. |
| Relic Welcome Scroll | Show useful SSH login/status output when connecting to `retropi@gpi`. | MOTD or login-time status summary with artifact pointers and safety reminders. |

The design should reduce operator uncertainty without reducing recovery
visibility. Earlier feedback is useful only if diagnostics can still reveal
the boot trail again.

## Do Not Overclaim

The operator observed that after flipping the GPi Case 2 side power switch, the
screen appears to do nothing for roughly 15 seconds before the first visible
text appears.

Treat that as a field observation, not a proven subsystem boundary.

- It may align with the display, framebuffer, KMS, or userspace stack becoming
  visible, but this design does not claim that yet.
- The silent window could include firmware work, kernel work, display handoff,
  boot logging not yet visible, panel/backlight behavior, or other case-board
  behavior.
- One stopwatch observation is not enough to identify the root cause or the
  earliest safe place to draw.
- True Boot Trace evidence is required before implementation chooses timing,
  service ordering, or display assumptions.

The first real QUEST-0091 True Boot Trace run adds useful but bounded
evidence: the script completed from `/home/retropi/`, captured post-boot
framebuffer/DRM hints, recorded `frontend_detected_ever: no`, and preserved
raw `throttled=0x50000`. It did not capture the exact side-switch timestamp,
exact first visible screen timestamp, first SSH availability, LED state,
power source, handheld/docked state, or visible frontend state. Those gaps
mean this design still cannot claim an earliest safe drawing point, a KMS or
framebuffer cause, a power cause, or a proven Boot Veil strategy.

## Startup Timeline

The desired UX should be reasoned about as a timeline, not as one splash
setting:

| Window | Current operator experience | Desired evidence | Possible future UX |
| --- | --- | --- | --- |
| Power switch flip | Side switch moves to on. | Human timestamp, LED/backlight state, power latch behavior if safely observable. | First Spark physical signal if available without disturbing recovery. |
| Pre-display silent window | Screen appears inactive for about 15 seconds in current observation. | True Boot Trace timing, external video notes if any, LED notes, power warnings after boot. | Avoid claims until evidence identifies what can run this early. |
| First visible framebuffer/text | First boot text becomes visible. | Kernel, DRM/KMS, framebuffer, and frontend milestones from logs. | Boot Veil may start here or shortly after, if reversible. |
| EmulationStation/frontend start | Frontend becomes visible and handheld is usable. | Process and display clues from Boot Trace and human notes. | Fade or remove veil; show normal frontend. |
| Post-boot Session Watch phase | Handheld is responsive for menu, play, idle-risk, or post-resume observation. | Session Watch Ledger, side-switch behavior, top-button caution, SSH state. | Relic Welcome Scroll helps SSH operators orient after login. |

## First Spark

First Spark means the earliest possible feedback that the operator can perceive
after the side switch is flipped. It is not the same as a pretty boot splash.
It may need to happen before Linux can draw anything.

Candidate approaches:

| Approach | What it could do | Risks | Recovery path |
| --- | --- | --- | --- |
| Firmware or kernel config | Reduce blank time, alter console handoff, or show a simple early logo if supported. | Can hide useful early text, break KMS/DPI timing, or make recovery harder if command-line changes are wrong. | Keep an alternate config note, make every change reversible from the boot partition, and preserve a documented "show boot text" profile. |
| Framebuffer splash | Draw a static image once `/dev/fb*` exists. | Starts only after framebuffer exists; may flicker, race KMS, or fail silently. | Service can be disabled over SSH or by removing one unit/config file; boot text profile restores console. |
| Plymouth or splash service | Provide a managed splash once userspace/display is ready. | Package and theme complexity; may mask service failures; may not start early enough for the 15-second window. | Boot with splash disabled, remove `quiet`/splash args, or disable the service. |
| LED/status signal | Use an existing safe physical signal if one is already available and understood. | Unknown GPIO/LED ownership can collide with case power, display, or latch behavior. | Do not use GPIO writes until hardware ownership is proven; keep LED plans read-only until a later approved quest. |

First Spark implementation must wait until the True Boot Trace Lantern records
where the device becomes observable. If evidence shows Linux cannot safely draw
during the silent window, the design should prefer a physical signal or leave
that window honest rather than fake certainty.

## Boot Veil

Boot Veil means hiding rough boot text, penguin graphics, or transient console
noise after the display path is active. It should make startup feel deliberate
without turning diagnostics into a blindfold.

Candidate approaches:

| Approach | What it could do | Risks | Recovery path |
| --- | --- | --- | --- |
| Firmware/kernel config | Hide kernel logos or reduce console verbosity. | Removes clues needed during boot failures; bad cmdline/config can affect display. | Keep a documented diagnostic profile that restores verbose boot text and kernel logos. |
| Framebuffer splash | Paint a low-level bitmap as soon as framebuffer exists. | Depends on framebuffer readiness; can be overwritten by kernel or frontend; may not hide earliest text. | Disable the splash service or remove the image path; do not make it required for boot. |
| Plymouth or splash service | Own a nicer boot veil during userspace startup. | More moving parts and possible package dependency; may obscure stuck services. | Provide an SSH-disable command and a boot-partition recovery note before enabling. |
| EmulationStation/frontend transition | Let the normal frontend take over without extra boot layers. | Does not solve early text; can still show rough console first. | No special recovery needed, but no early veil benefit either. |

Boot Veil must be opt-out and diagnosable. A future implementation should have
at least two profiles:

- Normal profile: hide safe-to-hide boot text after the display path is known.
- Diagnostic profile: show verbose boot text and preserve logs for recovery.

Do not enable `quiet`, splash arguments, or service masking until the True Boot
Trace evidence says which messages are valuable and which are cosmetic.

## Relic Welcome Scroll

Relic Welcome Scroll means useful SSH output when connecting to
`retropi@gpi`. It supports the scp-first workflow. It is not the handheld's
primary UX and must not require a keyboard attached to the GPi Case 2.

The dedicated source-of-truth design for this lane is
[GPi Case 2 Relic Welcome Scroll Design](gpi-case-2-relic-welcome-scroll-design.md).
That document defines the future banner layout, read-only field allowlist,
fast-path behavior, plain/no-color behavior, `scp` safety, missing-command
fallbacks, and recovery-first disable rules.

Candidate approach:

| Approach | What it could do | Risks | Recovery path |
| --- | --- | --- | --- |
| SSH MOTD/status scroll | Show boot status, last artifact paths, safety reminders, recovery links, and current session hints at login. | Login-time commands can be slow, noisy, or leak details if too broad. | Keep output bounded, local, redactable, and disable-able by removing one MOTD script. |

Expected content:

- Relic name and hostname.
- Uptime and local time.
- Reminder that side switch is the normal stock shutdown path while responsive.
- Reminder that the top sleep/resume button is suspect unless a procedure names
  that test.
- Last known Boot Trace and Session Watch artifact paths, when available.
- `vcgencmd get_throttled` value when cheaply available, labeled as a clue and
  not a diagnosis.
- One-line pointer to the recovery-first field procedure.

The scroll should avoid broad environment dumps, Wi-Fi secrets, ROM lists,
shell history, private keys, and long journal excerpts.

## Evidence Needed Before Implementation

The True Boot Trace Lantern should provide the minimum evidence before any
First Spark or Boot Veil implementation quest:

- Timestamped boot timeline from power-on observation to frontend visible.
- First visible text/framebuffer time from human notes.
- `/proc/uptime` at script start and during samples.
- KMS/DRM/VC4/DPI/framebuffer clues from `dmesg`, journal, `/sys/class/drm`,
  and `/sys/class/graphics` when available.
- Script summary fields for first display hint sample, first systemd timing
  sample, first journal hint sample, first dmesg hint sample, and first
  frontend-detected sample.
- Whether `/dev/fb*` exists during the capture and which framebuffer is active.
- EmulationStation or frontend first-detected milestone.
- `vcgencmd get_throttled` values, with sticky-flag caution.
- Temperature, load, memory, and disk context.
- Missing-command and permission-denied markers.
- Explicit note of what the Lantern could not see before it started.
- Human notes for LED, backlight, screen, side switch, and SSH visibility.

The
[True Boot Trace Field Run Procedure](gpi-case-2-true-boot-trace-field-run-procedure.md)
is the field trail for collecting those human notes beside the final Boot
Trace Ledger.

The True Boot Trace evidence should answer these implementation gates:

| Gate | Needed answer |
| --- | --- |
| Can anything safe draw before the observed silent window ends? | Yes, no, or unknown; do not guess. |
| When does framebuffer output become available? | First evidence timestamp or inconclusive. |
| Which boot text is diagnostic versus cosmetic? | Keep diagnostic text visible in recovery profile. |
| Does display handoff coincide with warnings? | Evidence bucket, not root-cause claim. |
| Can SSH still recover if the veil misbehaves? | Confirmed path or no implementation. |

## Recovery Requirements

Any future implementation must include recovery before polish:

- A documented diagnostic mode that restores boot text.
- A clear path to disable the veil over SSH.
- A boot-partition recovery note for config/cmdline changes.
- No dependency on the top sleep/resume button.
- No assumption that an attached keyboard exists.
- No GPIO writes or LED ownership unless a later hardware quest proves the
  signal is safe.
- No automatic repair, telemetry, upload, or Lantern Dispatch.
- No change to `SafeShutdown.py` or side-switch behavior in this quest.

## Relationship To Existing Lanterns

The [True Boot Trace Lantern](gpi-case-2-true-boot-trace-lantern-design.md)
gathers the evidence needed before changing boot UX. It should remain the
source for startup timing and do-not-overclaim notes.

The [Session Watch Lantern](gpi-case-2-session-watch-lantern-design.md)
continues after boot. It should verify that the handheld remains usable once
the frontend is up and that any future veil does not create later runtime
confusion.

The [Relic Welcome Scroll Design](gpi-case-2-relic-welcome-scroll-design.md)
keeps SSH login greeting behavior separate from First Spark and Boot Veil work.
It should remain read-only, fast, interactive-only, and safe for the current
scp-first field practice.

The [Boot Power Trace Lantern Map](gpi-case-2-boot-power-trace-lantern-map.md)
keeps power-warning timing separate from startup cosmetics. A veil must never
hide voltage or throttling evidence from recovery artifacts.

The [Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
remains the operator fallback if a splash, veil, login scroll, or later runtime
change makes the handheld harder to observe.

## Future Quest Boundaries

Later implementation quests may choose one layer at a time:

1. Add a read-only SSH Relic Welcome Scroll.
2. Add a reversible framebuffer splash prototype.
3. Evaluate Plymouth or an equivalent splash service only after evidence shows
   it improves the correct window.
4. Consider firmware/kernel changes only with a tested diagnostic rollback.
5. Consider physical First Spark signaling only after hardware ownership is
   proven and recovery remains intact.

Do not combine all layers in one first implementation. The first useful
startup UX improvement is the one that can be safely backed out in the field.
