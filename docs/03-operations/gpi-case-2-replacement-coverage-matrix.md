---
id: OPS-GPI-CASE-2-REPLACEMENT-COVERAGE-MATRIX-001
title: GPi Case 2 Replacement Coverage Matrix
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Translate the GPi Case 2 SafeShutdown behavior map into replacement responsibilities for RetroFlag Power and future Arcadia Runtime work.
related:
  - gpi-case-2-safeshutdown-script-behavior-map.md
  - safeshutdown-replacement-boundary-map.md
  - gpi-case-2-acceptance-checklist.md
  - gpi-case-gpio-probe-ledger.md
  - ../02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../02-hardware/gpi-case-2-emergency-recovery-research-ledger.md
  - ../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md
  - ../04-architecture/arcadia-runtime-migration-path.md
last_updated: 2026-07-09
---

# GPi Case 2 Replacement Coverage Matrix

> This matrix turns the old SafeShutdown spellbook into a replacement checklist:
> what the future runtime must preserve, what it may only observe, what remains
> deferred, and what stays forbidden until the field evidence earns it.

This document is documentation only. It does not change Go code, read GPIO,
write GPIO, execute shutdown, install or activate systemd, alter `rc.local`,
replace `/opt/RetroFlag/SafeShutdown.py`, implement resume, add persistent
state, add telemetry, make project-code network calls, change installers or
packaging, run RetroFlag install scripts, or approve hardware modification.

RetroFlag Power remains the current GPi Case 2 field kit. Arcadia Runtime
remains future optional runtime direction, not an active rename or implemented
replacement.

## Vocabulary Boundaries

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from a
  configured hardware profile.
- Lanterns are read-only probes and diagnostics.
- Lantern Dispatch remains future optional support infrastructure, not
  implemented.

## Interpretation Rules To Preserve

- The side switch is software-mediated until proven otherwise, not a direct
  physical power cut.
- `SafeShutdown.py` being enabled does not guarantee recovery if Linux is
  already stalled.
- The blinking LED and display-off behavior during the RCU stall suggest some
  board-controlled behavior, but that remains an inference, not a schematic
  fact.
- Avoiding the top power-save button is not sufficient if automatic idle
  power-save can happen after roughly 15-20 minutes.
- KMS fixed one display path, but KMS timing or display/power-save interaction
  may still be involved in the current power-save/resume risk.
- A 2026-07-09 Bundle Collector Field Lantern run captured a successful
  post-resume session, so the resume wedge is intermittent rather than
  guaranteed. That Relic is post-resume evidence, not proof of transition-time
  behavior.

## Coverage Matrix

| Behavior or responsibility | Current known owner | Current evidence | Replacement stance | Safe next test or evidence needed | Risk |
| --- | --- | --- | --- | --- | --- |
| Side switch observation | Stock script and Linux userspace | Upstream GPi Case 2 script configures GPIO26 as `powerPin` input with pull-up and waits for a falling edge. Hardware notes say the side switch does not directly cut battery power. | Must preserve | Read-only GPIO26 observations while stock script remains active, recorded as `SignalLow`, `SignalHigh`, or `SignalUnverified`; separate interpreted `SwitchOn`/`SwitchOff` mapping only after profile review. | High |
| Side switch interpreted shutdown request | Stock script and configured interpretation, with board involvement unknown | The script treats GPIO26 falling edge as shutdown intent, then kills EmulationStation and calls Linux shutdown. The 2026-07-08 stall showed side-switch off did not recover the device. | Must preserve | Prove edge polarity, debounce, boot timing, docked/handheld behavior, and whether the signal remains actionable when Linux is healthy; document interpretation policy separately from raw signals. | High |
| GPIO26 raw signal handling | Linux userspace under stock script; future lanterns may observe only | GPIO26 is named `powerPin` in upstream script. Probe ledger defines raw results as `SignalLow`, `SignalHigh`, and `SignalUnverified` only. | Observe only | Add repeated read-only probe rows across visible switch positions, modes, and script state without disabling SafeShutdown or treating raw state as meaning. | Medium |
| GPIO27 power-enable/latch behavior | Stock script drives; board or MCU likely consumes | Upstream script configures GPIO27 as `powerenPin` output and drives it HIGH. Boundary map treats it as a power-enable latch path, but exact board target is unknown. | Forbidden until proven | Board-level evidence for what GPIO27 drives; passive observation or source review for required HIGH timing during boot, shutdown, power-save, resume, docked use, and failure states. | Critical |
| GPIO18 LCD/power-save involvement | Stock LCD scripts and Linux userspace; KMS path now avoids legacy rewrite assumptions | Upstream `lcdfirst.sh`/`lcdnext.sh` use GPIO18 as `HDMI_HPD`; KMS notes say modern `rgb666-padhi` avoids GPIO18/19 and GPIO26/27. | Defer | Prove whether GPIO18 still reports dock or HDMI state under KMS without running legacy scripts or rewriting `/boot/config.txt`. | High |
| Top power-save/resume button behavior | Board or MCU likely, with Linux path unknown | Field notes show screen turns off, indicator flashes, face buttons do not wake, and top button wakes visible state. It is not detected like normal EmulationStation controls. | Observe only | Use passive input-device inspection and documented field observations to identify whether the button appears as HID, GPIO, board control, or remains invisible to Linux. | High |
| Automatic idle power-save behavior | Board or MCU, Linux userspace, display/audio domain, or unknown | Field notes say GPi Case 2 may auto-enter display/audio power-save after roughly 15-20 minutes idle. | Forbidden until proven | Reproduce only with a written recovery plan that does not depend on SSH, ping, side-switch shutdown, battery depletion, or repeated CM4 removal. | Critical |
| LCD sleep/wake behavior | Board or MCU, display/backlight path, kernel/KMS, or unknown | Screen can turn off while Linux may stay alive; top button can wake visible state; wake has shown undervoltage; a later incident showed RCU stalls; a 2026-07-09 post-resume Bundle Collector Lantern captured one successful resume with `get_throttled=0x0`. | Defer | Field-test sleep/wake under KMS with power-quality notes, SSH/network status, controller status, dmesg, recovery steps, and a future watcher when transition evidence is needed. | Critical |
| LCD/HDMI switching | Stock LCD scripts currently own legacy path; future KMS behavior unknown | RetroFlag scripts call `lcdnext.sh`, read GPIO18, and rewrite `/boot/config.txt`; current project direction says this is not KMS-safe. | Defer | Design evidence for KMS-native LCD/HDMI behavior that does not call `lcdnext.sh`, `lcdfirst.sh`, or restore legacy FKMS/DPI config files. | High |
| Docking behavior | Stock scripts, Linux display stack, dock hardware, and unknown audio/power paths | Upstream README says script automatically switches LCD/HDMI when docked; behavior map lists unknown dock effects on audio, power domains, wake, and recovery. | Defer | Field checklist entries for boot docked, handheld-to-dock, dock-to-handheld, active connector, audio path, controller path, and whether any legacy script ran. | High |
| Handheld audio after KMS | Linux audio stack, USB audio, possibly board routing; current owner unknown | Hardware notes show `audremap` disabled because it conflicts with KMS DPI; USB audio device is present; audio after KMS is not fully verified. | Defer | Verify RetroPie game audio and menu audio in handheld mode, recording selected device, volume path, emulator, and whether `audremap` remains disabled. | Medium |
| Docked audio after KMS | Linux audio stack, HDMI or dock audio path, dock hardware; current owner unknown | Acceptance checklist marks docked audio after KMS as unknown and asks for HDMI device/dock state notes. | Defer | Verify docked HDMI or dock audio path after KMS boot and after dock transition, including device identity and insertion order. | Medium |
| EmulationStation/controller identity | Linux input stack and EmulationStation | Hardware notes observed Microsoft Xbox 360 gamepad / `GBA Pi Case+` / Nuvoton identity. Acceptance checklist requires exact identity before mapping assumptions. | Must preserve | Capture `/proc/bus/input/devices`, `/dev/input/by-id`, `/dev/input/by-path`, and mapping behavior during field tests. | Medium |
| Button above Select / left of RetroFlag logo | Unknown: HID, GPIO, board control, or unmapped input | Hardware notes say it was not detected by EmulationStation like other controls. | Observe only | Identify with passive input-event tools when available; record whether it is gamepad input, GPIO, power-board control, or still undetected. | Medium |
| Safe shutdown sequencing | Stock script and Linux userspace | Upstream script kills `emulationstation`, kills truncated `emulationstatio`, sleeps about five seconds, then calls `shutdown -h now`. | Must preserve | Define future clean-shutdown policy for EmulationStation, running games, process cleanup, delay, failure handling, and Linux shutdown before any shutdown execution exists. | High |
| Process shutdown assumptions such as EmulationStation kills | Stock script | Behavior map records direct `killall` commands and notes they are observed assumptions, not proof of a good user-facing shutdown. | Defer | Determine expected RetroPie and active-game shutdown behavior; avoid copying `killall` blindly without evidence and user-safety review. | Medium |
| Kernel-stall/RCU-stall emergency recovery | Kernel, power-save/display path, board behavior, or unknown | QUEST-0064 records RCU stall messages, no SSH, no ping, side-switch off failure, top-button visible-state toggling, and physical CM4 removal as the only observed stop. | Forbidden until proven | Find a reversible, mapped, maintainer-approved recovery path through public docs, board photos, electrical review, and separate decision before deeper risky tests. | Critical |
| Session Watch Lantern | Future local read-only watcher only | Post-resume Bundle Collector satchels are useful, but they cannot prove what happened during sleep/resume unless a watcher was already running. | Defer | Map a watcher that records pre-sleep state, records post-resume state when available, tracks `get_throttled`, temperature, frontend, and input hints over time, and avoids telemetry and automatic fixes. | Medium |
| Local diagnostics coverage | Future Linux userspace lanterns only | Local diagnostics bundle map and acceptance checklist describe future local-only support output; no implementation is authorized here. | Defer | Define read-only diagnostics that summarize GPIO vocabulary, script presence, KMS state, audio state, logs, and checklist evidence without network submission. | Low |
| Field checklist coverage | EDC acceptance checklist | GPi Case 2 acceptance checklist has rows for power, GPIO, display, dock, audio, sleep/resume, RetroPie, diagnostics, rollback, and public readiness. | Must preserve | Keep checklist rows current as tests are performed; unknown remains `Unknown` until evidence exists. | Low |

## Replacement Readiness Gates

Before replacing `SafeShutdown.py`, prove:

- GPIO27 latch behavior is understood enough to preserve power safely without
  guessing.
- GPIO26 side-switch behavior is observed, interpreted, debounced, and modeled
  without collapsing raw `Signal*` observations into `Switch*` meanings.
- Clean EmulationStation, active-game, and Linux shutdown sequencing is defined
  and tested first in dry-run or fake paths.
- Top-button power-save/resume and automatic idle power-save have a documented
  coexistence or preservation plan.
- KMS-safe handheld LCD, docked HDMI, display transition, and audio behavior
  have field evidence.
- Recovery from dark-screen, no-network, no-side-switch-response states has a
  safer plan than battery depletion or repeated CM4 cartridge/card removal.

Before touching GPIO27 behavior, prove:

- What GPIO27 drives on the GPi Case 2 board.
- Whether HIGH must be asserted before, during, or after boot.
- Whether LOW, release, process exit, or scheduler stalls change case power
  state.
- How the line behaves during shutdown, power-save, resume, docked mode, and
  failure states.
- Whether a board MCU, latch, regulator, or other circuit also owns the path.

Before enabling any systemd service, prove:

- Startup order, coexistence, and handoff with `/opt/RetroFlag/SafeShutdown.py`
  and `/etc/rc.local` are documented.
- The service has a dry-run plan, rollback plan, and failure behavior that does
  not strand the case.
- No GPIO writes, shutdown execution, config mutation, or legacy LCD script
  calls occur before their own acceptance gates pass.

Before claiming power-save/resume support, prove:

- Top-button power-save/resume works under KMS without RCU stalls, network
  loss, shutdown-path loss, or hidden filesystem risk.
- Automatic idle power-save after roughly 15-20 minutes is tested with a
  documented non-SSH-dependent recovery plan.
- Successful and failed resumes are both recorded as separate Field Lantern
  evidence, with post-resume captures treated as after-the-fact satchels unless
  a Session Watch Lantern was already running.
- Display, audio, controller, side-switch, and dock behavior return cleanly in
  handheld and docked modes.
- Any blinking LED or display-off behavior is attributed only to evidence, not
  to an assumed MCU or schematic.

Before any emergency reset or cutoff mod is considered, prove:

- Public schematic evidence, official board documentation, or maintainer-created
  board maps identify the relevant nets and components.
- CM4 `RUN_PG`, `GLOBAL_EN`, `EXT_nRESET`, regulator enable, GPIO26, GPIO27,
  battery, and charge paths are mapped well enough for review.
- Voltage levels, current limits, backfeed risk, filesystem risk, and battery
  safety have written review.
- The proposal is reversible, bench-only, does not touch lithium battery or
  charging circuitry, and has separate maintainer approval.

## Current Standing

This matrix does not make RetroFlag Power replacement-ready. It names the
coverage the future runtime must earn before the stock script leaves the
satchel.
