---
id: OPS-GPI-CASE-2-ACCEPTANCE-CHECKLIST-001
title: GPi Case 2 Acceptance Checklist
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define the field-test gates required before RetroFlag Power replaces SafeShutdown.py, installs a daemon service, publishes a public installer, or broadens toward Arcadia Runtime.
related:
  - README.md
  - docs/03-operations/safeshutdown-replacement-boundary-map.md
  - docs/03-operations/installer-migration-toolkit-map.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - docs/00-project/quests/0053-add-the-gpi-case-2-acceptance-checklist.md
  - docs/00-project/quests/0057-add-gpi-case-2-field-test-checklist-entries.md
last_updated: 2026-07-07
---

# GPi Case 2 Acceptance Checklist

> This is the field gate ledger for the GPi Case 2 relic: prove power, display,
> dock, audio, sleep, resume, rollback, and user safety before the runtime
> carries the old spellbook out of the satchel.

This document is documentation only. It does not replace
`/opt/RetroFlag/SafeShutdown.py`, install a service, activate a daemon, write
GPIO, execute shutdown, mutate files, mutate configuration, change `rc.local`,
implement diagnostics, implement an installer, submit telemetry, or make
network calls.

RetroFlag Power remains the current GPi Case 2 prototype. Arcadia Runtime
remains the favored future runtime direction, not an active rename. This
checklist is a gate for future confidence, not proof that any replacement,
installer, diagnostics path, or migration has been implemented.

## Purpose

This checklist defines the field tests required before any of these trails
open:

- Replacing the stock RetroFlag `SafeShutdown.py` path.
- Installing or enabling a daemon service.
- Publishing a public installer or migration toolkit.
- Broadening the work toward an Arcadia Runtime release.

The legacy RetroFlag script currently appears to own more than shutdown:
power latch behavior, side-switch shutdown, top-button power-save and resume,
LCD/HDMI docking behavior, and possibly timing assumptions around KMS.

This checklist turns those risks into visible badges. Passing it should mean
maintainers have field evidence across power, display, dock, audio,
sleep/resume, RetroPie, diagnostics, rollback, and user safety before real
behavior changes.

## Current No-Go Status

The checklist is not passed.

- `SafeShutdown.py` is not replaced.
- No service is installed.
- No shutdown execution exists.
- No GPIO writes exist.
- No config mutation exists.
- No installer implementation exists.
- No diagnostics implementation exists.
- No daemon activation exists.
- Power-save, resume, and automatic power-save behavior are unsafe/unverified
  after a field incident with repeated Linux RCU stall messages, lost network
  recovery, side-switch shutdown failure, and physical CM4 cartridge/card
  removal as the only observed stop.

Until this ledger is filled with real observations, the stock script remains
part of the active GPi Case 2 power path.

## Hardware Identity Checks

Record the exact device before interpreting behavior:

- Raspberry Pi model and Compute Module 4 details, when applicable.
- GPi Case 2 hardware notes, including visible revision notes if available.
- Operating system version.
- Kernel version.
- RetroPie and EmulationStation version, if available.
- Display stack: KMS, FKMS, or unknown.
- Dock state during the test: dock present or not present.
- Audio device presence in handheld and docked modes.

Unknown values should remain `Unknown` in the ledger. Do not turn missing
identity details into confidence.

## Power And GPIO Checks

Field tests should prove the power trail before any write path exists:

- Observe GPIO26 side-switch signals under the stock script.
- Document the current understanding of GPIO27 power-enable latch behavior.
- Record raw `SignalLow`, `SignalHigh`, and `SignalUnverified` observations.
- Do not treat any raw signal as switch meaning without profile
  interpretation.
- Confirm side-switch shutdown behavior while the stock script owns the path.
- Test behavior with the stock script disabled only if the procedure is safe,
  reversible, and documented before the test begins.

The raw signal ledger is only a lantern. It does not prove `SwitchOn`,
`SwitchOff`, shutdown intent, or latch policy by itself.

## Display And Dock Checks

Verify display behavior in each physical mode:

- Handheld LCD boot.
- Docked HDMI boot.
- Handheld-to-dock transition, if applicable.
- Dock-to-handheld transition, if applicable.
- KMS display behavior during startup, shutdown, sleep, and resume.
- Any timing or order dependency around boot, dock insertion, wake, or display
  switching.
- Recovery path if the screen goes dark.

A dark screen recovery path should be written before risky field tests begin.
SSH availability, keyboard access, power-cycle expectations, and rollback
steps should be clear enough to follow under pressure.

## Sleep/Resume Checks

Verify the top-button and wake trail before replacing the legacy owner:

- Top-button power-save behavior.
- Resume behavior.
- Automatic display/audio power-save behavior after roughly 15-20 minutes of
  no input.
- SSH and network behavior during sleep.
- Wake button behavior.
- Risk notes for losing access during field tests.
- Recovery behavior when Linux shows repeated `rcu: INFO: rcu_preempt detected
  stalls on CPUs/tasks` messages.

If sleep can drop Wi-Fi or SSH, the test note should say how the maintainer
will recover without assuming the display or network will return.

Current caution:

- A 2026-07-08 field incident suggests a kernel, hardware, display, or
  power-save stall where software shutdown paths may no longer be reliable.
- `SafeShutdown.py` was believed to be enabled, so do not classify this as only
  a disabled stock-script issue.
- Waiting for battery depletion or repeatedly pulling the CM4 cartridge/card is
  unacceptable for development because of filesystem and data-loss risk.
- Investigate whether a reversible emergency reset or safe power-cut path
  exists before deeper field testing.
- Look for schematics, teardown photos, board labels, test pads, CM4
  `RUN`/`GLOBAL_EN`/reset/power-enable paths, and regulator enable lines.
- Do not cut battery leads or modify lithium battery or charging circuitry
  without mapping the board first.

## Audio Checks

Verify audio after the KMS migration in both modes:

- Handheld audio after KMS migration.
- Docked HDMI audio after KMS migration.
- RetroPie game audio.
- EmulationStation menu audio, if applicable.
- Known unknowns and notes.

Display passing is not enough for public readiness. The audio badge needs its
own evidence.

## RetroPie / EmulationStation Checks

Verify the user-facing RetroPie trail:

- EmulationStation startup.
- Controller detection.
- Xbox 360 gamepad mapping behavior.
- Button above Select and left of the RetroFlag logo, currently undetected or
  requiring further mapping.
- Safe game exit behavior.
- Clean shutdown sequence expectations.

Clean shutdown should include what EmulationStation and a running game are
expected to do before Linux shutdown begins. If the expectation is unknown, the
ledger should say so plainly.

## Diagnostics And Rollback Checks

Future support and installer paths must be reversible before any mutation:

- A local diagnostics bundle can be produced in the future.
- An installer plan can be previewed in the future.
- A backup and restore ledger exists in the future.
- No network submission is required.
- The rollback path is documented before any mutation.

The diagnostics and installer maps are future compasses. This checklist does
not implement bundle creation, installer planning, backup, restore, network
submission, or file changes.

## Public Readiness Gates

Public readiness needs every badge below:

- This checklist passed on at least the maintainer GPi Case 2.
- SafeShutdown boundary prerequisites satisfied.
- KMS handheld and docked behavior stable.
- Audio verified in handheld and docked modes.
- Rollback tested.
- Unofficial and non-affiliation language in place.
- Naming and domain clearance considered before any Arcadia Runtime public
  release.

If any gate is unknown, the public trail remains closed.

## Current Field-Test Entries

Use these rows as the next GPi Case 2 field-test pass. They turn the current
unknowns into visible checkpoints without pretending the old script has been
replaced.

For GPIO and switch notes, preserve the vocabulary boundary:

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from
  the configured hardware profile.
- Do not write interpreted switch meaning into a raw signal result.

| Date | Mode | Field-test entry | Expected result | Observed result | Pass/Fail/Unknown | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| YYYY-MM-DD | Handheld | Handheld audio after KMS | Built-in handheld audio path plays RetroPie game audio after KMS boot. | Unknown | Unknown | Note selected ALSA/Pulse device, volume path, emulator tested, and whether `audremap` remains disabled. |
| YYYY-MM-DD | Docked | Docked audio after KMS | Docked HDMI or dock audio path plays RetroPie game audio after KMS boot. | Unknown | Unknown | Note HDMI device presence, dock state at boot, dock insertion order, and whether EmulationStation menu audio also works. |
| YYYY-MM-DD | Handheld | LCD sleep/wake behavior | LCD/backlight can enter power-save and return without losing EmulationStation state. | Unknown | Unknown | Record whether SSH, Wi-Fi, controller input, and display return cleanly after wake. |
| 2026-07-08 | Handheld | Power-save/resume RCU stall incident | Power-save or resume returns without kernel stalls, network loss, or loss of shutdown recovery. | Repeated Linux `rcu: INFO: rcu_preempt detected stalls on CPUs/tasks` console output; SSH unavailable; ping unavailable; side switch moved off but did not shut down; top sleep/resume button still toggled visible state; only physical CM4 cartridge/card removal stopped the device. | Fail | Unresolved high-priority field incident. `SafeShutdown.py` was believed enabled, so do not treat as only disabled stock script. A field photo exists and can be linked later if evidence-asset rules are explicit. |
| YYYY-MM-DD | Handheld | Top-button power-save/resume behavior | Top button triggers the observed case power-save path and wakes the LCD/backlight without kernel stalls, network loss, or shutdown-path loss. | Unsafe/unverified after RCU stall incident | Fail | Treat this as case/power-board behavior unless an input event or GPIO path is proven. Record undervoltage, access-loss, RCU stall, network, side-switch, and recovery events. |
| YYYY-MM-DD | Handheld | Automatic display/audio power-save after idle | After roughly 15-20 minutes of no input, any automatic display/audio power-save path returns without kernel stalls, network loss, or shutdown-path loss. | Unsafe/unverified after field report that the case may auto-enter power-save | Unknown | Avoiding the top power-save button may not avoid this path. Field-test only with a documented recovery plan that does not depend on SSH, ping, side-switch shutdown, battery depletion, or repeated CM4 cartridge/card removal. |
| YYYY-MM-DD | Handheld | Side switch with `SafeShutdown.py` still active | Side switch behavior is observed while the stock script still owns the path. | Unknown | Unknown | Record visible side-switch position, process state, raw GPIO26 probe result if taken, and any interpreted `SwitchOn`/`SwitchOff` meaning only in a separate interpretation note. Do not disable the stock script for this row. |
| YYYY-MM-DD | Handheld to docked | LCD to HDMI switching | Display changes from built-in LCD to docked HDMI without legacy config rewrites. | Unknown | Unknown | Record insertion order, active connector, recovery path, and whether old `lcdnext.sh` or `lcdfirst.sh` ran. |
| YYYY-MM-DD | Docked to handheld | HDMI to LCD switching | Display returns from docked HDMI to built-in LCD without losing the session. | Unknown | Unknown | Record removal order, active connector, recovery path, and whether the KMS DPI display remains present. |
| YYYY-MM-DD | Docked | Docking behavior | Dock attach, controller, display, and audio behavior are stable under KMS. | Unknown | Unknown | Note whether the system was booted docked or docked after boot, and record any timing dependency. |
| YYYY-MM-DD | Handheld | Button above Select / left of RetroFlag logo | Button is identified as gamepad input, GPIO, power-board control, or still undetected. | Unknown | Unknown | During EmulationStation mapping it was not detected like the other controls. Record `/dev/input` or `evtest` evidence when available. |
| YYYY-MM-DD | Handheld | Controller identity | Built-in controller identity is recorded before mapping assumptions are made. | Unknown | Unknown | Current field note observed Microsoft Xbox 360 gamepad / `GBA Pi Case+` / Nuvoton identity. Record exact device names from the test image. |

## Field Ledger Table

Use this compact ledger for each field-test pass:

| Date | Device | OS/kernel | Mode: handheld/docked | Test area | Expected result | Observed result | Pass/Fail/Unknown | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| YYYY-MM-DD | GPi Case 2 / Pi details | OS and kernel | Handheld, docked, or transition | Power, display, dock, audio, sleep/resume, RetroPie, diagnostics, rollback | Expected behavior | Observed behavior | Unknown | Recovery notes, risks, follow-up quests |

The first honest value for an untested row is `Unknown`. The checklist becomes
useful by showing what remains dark as clearly as what has been lit.
