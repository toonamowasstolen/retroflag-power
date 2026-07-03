---
id: HARDWARE-GPI-CASE-2-001
title: GPi Case 2 Reference Hardware
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Hardware Porters
  - Future Maintainers
purpose: Document the RetroFlag GPi Case 2 reference platform, known facts, current shutdown behavior, boot context, terminology, assumptions, risks, and validation needs for RetroFlag Power.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - PROJECT_CHARTER.md
  - docs/00-project/requirements.md
  - docs/00-project/roadmap.md
  - docs/04-architecture/system-overview.md
  - docs/13-reference/terminology.md
  - docs/13-reference/glossary.md
last_updated: 2026-07-03
---

# GPi Case 2 Reference Hardware

> Your Honor, we call GPi Case 2 to the stand.

This document records the reference hardware for RetroFlag Power.

The GPi Case 2 is the first witness because it is the hardware that started the project.

The goal of this document is to capture what is known, what is assumed, what must be validated, and what the first implementation must respect.

---

# 1. Witness Identity

## Hardware

```
RetroFlag GPi Case 2
```

## Compute Module

```
Raspberry Pi Compute Module 4 Rev 1.1
```

## Memory

```
4 GB RAM
```

## Storage

```
Samsung EVO Select microSD
```

Known block devices from current system:

```
mmcblk0p1  /boot
mmcblk0p2  /
```

## Operating System

```
RetroPie on Raspberry Pi OS Bullseye 64-bit
```

## Primary Frontend

```
EmulationStation
```

## Primary Emulator Framework

```
RetroArch
```

## Role in Project

```
Reference Platform
```

This is the platform that defines the first working implementation.

Future hardware support should not weaken the focus on this reference platform.

---

# 2. Testimony Summary

The GPi Case 2 currently behaves like a retro handheld, but its power management path still depends on legacy startup behavior.

Known current behavior:

```
/etc/rc.local
      │
      ▼
sudo python /opt/RetroFlag/SafeShutdown.py &
      │
      ▼
RetroFlag safe shutdown behavior
```

RetroFlag Power intends to replace that path with:

```
systemd
  │
  ▼
retroflag-powerd
  │
  ▼
event-driven power handling
  │
  ▼
safe shutdown
```

The first product obligation is not to be clever.

The first product obligation is to be safe.

---

# 3. Confirmed Facts

## FACT-0001 — Current safe shutdown script path

The current RetroFlag shutdown script is located at:

```
/opt/RetroFlag/SafeShutdown.py
```

## FACT-0002 — Current startup path

The current script is started from:

```
/etc/rc.local
```

Known rc.local line:

```
sudo python /opt/RetroFlag/SafeShutdown.py &
```

## FACT-0003 — Current daemon replacement target

RetroFlag Power intends to replace the rc.local-launched script with a systemd-supervised daemon.

Expected future service concept:

```
retroflag-power.service
```

Expected future daemon concept:

```
retroflag-powerd
```

## FACT-0004 — Current boot system

The system uses systemd.

Known boot timing from prior measurement:

```
kernel:    about 2.157s
userspace: about 22.537s
total:     about 24.694s
```

## FACT-0005 — Existing boot optimization work

The following services or features have already been disabled or reduced during boot tuning work:

- Bluetooth
- Samba
- ModemManager
- Avahi
- swap
- noisy kernel boot output

## FACT-0006 — Current EEPROM version

Known current EEPROM version:

```
2022-04-26
```

## FACT-0007 — CM4 EEPROM update note

`rpi-eeprom-update` is not enabled by default on CM4.

The recommended CM4 EEPROM flashing method is `rpiboot`.

This project should not casually recommend EEPROM changes without careful documentation and validation.

## FACT-0008 — Current boot order under consideration

The system has considered simplifying boot order.

Known values discussed:

```
Current / prior: 0xf25641
Potential:       0xf1
Potential:       0xf41
```

These values require validation before being documented as recommended behavior.

## FACT-0009 — Current display configuration

Known current display context includes:

- custom DPI display timing
- 640x480 configuration
- `dtoverlay=vc4-fkms-v3d`
- `max_framebuffers=2`
- 64-bit enabled

This display stack is important to boot, splash, and emulator behavior.

## FACT-0010 — Current shutdown modernization goal

The first modernization goal is to replace legacy safe shutdown startup with:

- systemd service
- daemon lifecycle
- structured logs
- no rc.local dependency
- safer maintainable implementation path

---

# 4. Terminology Findings

## Power Switch

The GPi Case 2 uses a latching power control.

Project term:

```
Power Switch
```

Avoid:

```
Power Button
```

Reason:

A button is momentary.

The GPi Case 2 power control maintains ON/OFF position.

## Reset Button

The GPi Case 2 reset control should be modeled as a momentary button.

Project term:

```
Reset Button
```

Reason:

It is pressed and released.

## Reference Platform

The GPi Case 2 configuration in this document is the first-class validation target.

Other hardware may be researched later, but this platform is the witness currently on the stand.

---

# 5. Current Software Stack

The reference stack is:

```
Raspberry Pi OS Bullseye 64-bit
RetroPie
EmulationStation
RetroArch
systemd
rc.local
SafeShutdown.py
```

RetroFlag Power will initially interact most directly with:

- systemd
- journald
- Linux GPIO
- the current shutdown path
- later, EmulationStation and RetroArch

The first daemon milestone should not attempt to integrate with the full stack.

---

# 6. Current Boot Context

Known boot time baseline:

```
systemd-analyze total: about 24.694s
```

Known largest services from prior measurements included:

```
nmbd:          about 16.31s
hciuart:       about 9.10s
fstrim:        about 2.10s
raspi-config:  about 1.58s
man-db:        about 1.35s
avahi:         about 1.21s
swapfile:      about 1.13s
rng-tools:     about 1.05s
wpa_supplicant: about 1.01s
ModemManager:  about 0.78s
```

Some of these have already been disabled or reduced.

This matters because the product vision includes a future power-to-resume target.

However, boot optimization must remain measured and cautious.

Related requirement:

```
REQ-0302 — Measure before optimizing
```

---

# 7. Current Shutdown Context

Current shutdown path:

```
/etc/rc.local
      │
      ▼
SafeShutdown.py
      │
      ▼
polling / hardware detection
      │
      ▼
shutdown behavior
```

Target shutdown path:

```
systemd
  │
  ▼
retroflag-powerd
  │
  ▼
hardware adapter
  │
  ▼
PowerSwitchChanged event
  │
  ▼
Power Service
  │
  ▼
safe shutdown
```

The replacement must satisfy:

```
REQ-0001 — Safe shutdown from Power Switch OFF
REQ-0002 — Preserve original safe shutdown intent
REQ-0003 — Replace rc.local startup
REQ-0010 — Disable original shutdown path safely
```

---

# 8. Hardware Abstraction Needs

The GPi Case 2 implementation should not leak raw GPIO details into core services.

Core services should consume events such as:

```
PowerSwitchChanged
ResetButtonPressed
ShutdownRequested
```

Core services should not consume raw details such as:

```
gpiochip name
line offset
active-low state
edge type
pull-up behavior
```

Those details belong in:

```
Hardware Service
GPIO Adapter
Hardware Profile
Validation Notes
```

Related requirements:

```
REQ-0102 — GPIO isolation
REQ-0103 — Event-driven GPIO direction
REQ-0104 — libgpiod direction
REQ-0105 — Mock hardware support
REQ-0106 — Hardware profiles
```

---

# 9. Candidate Hardware Profile Shape

A future GPi Case 2 hardware profile may include:

```
id: gpi-case-2-cm4
name: RetroFlag GPi Case 2 with Raspberry Pi CM4
platform: raspberry-pi
board: cm4
case: retroflag-gpi-case-2
capabilities:
  - power-switch
  - reset-button
  - display
  - audio
```

Future hardware profile details may include:

```
power_switch:
  type: latching
  backend: gpio
  active_low: true/false
  line: TBD

reset_button:
  type: momentary
  backend: gpio
  active_low: true/false
  line: TBD
```

The exact GPIO values must be validated before becoming facts.

---

# 10. Known Unknowns

The following items require validation.

## UNKNOWN-0001 — Exact Power Switch GPIO line

The exact GPIO chip and line used by the GPi Case 2 Power Switch must be confirmed.

## UNKNOWN-0002 — Exact Reset Button GPIO line

The exact GPIO chip and line used by the Reset Button must be confirmed.

## UNKNOWN-0003 — Active-low or active-high behavior

The project must confirm whether each relevant control is active-low or active-high.

## UNKNOWN-0004 — Edge behavior

The project must confirm whether hardware events are best modeled as rising edge, falling edge, both edges, or state polling plus debounce.

## UNKNOWN-0005 — Debounce requirements

The project must determine whether software debounce is needed and how much.

## UNKNOWN-0006 — Current SafeShutdown.py implementation details

The original script should be reviewed before replacement to understand timing, pin behavior, and safety assumptions.

## UNKNOWN-0007 — Shutdown hold / delay behavior

The project should confirm whether RetroFlag hardware expects a particular timing sequence during shutdown.

## UNKNOWN-0008 — Battery state availability

It is unknown whether this hardware exposes useful battery state to Linux.

Battery features remain research.

## UNKNOWN-0009 — Backlight control availability

It is unknown whether this hardware exposes useful backlight control for future sleep-like or resume polish features.

## UNKNOWN-0010 — Boot-order impact

The project must measure whether boot-order changes meaningfully improve boot time.

---

# 11. Assumptions

The following are assumptions until validated or promoted to facts.

## ASSUMPTION-0001 — systemd replacement is viable

Assumption:

A systemd daemon can replace the current rc.local SafeShutdown.py startup path.

Risk:

If RetroFlag hardware depends on script timing or behavior not yet understood, replacement may require careful compatibility work.

## ASSUMPTION-0002 — event-driven GPIO is viable

Assumption:

The Power Switch and Reset Button can be handled through event-driven GPIO rather than continuous busy polling.

Risk:

Hardware or kernel behavior may require a hybrid approach.

## ASSUMPTION-0003 — libgpiod direction is practical

Assumption:

The GPi Case 2 reference platform can support the planned modern GPIO direction.

Risk:

OS version, permissions, or kernel interfaces may require implementation adjustment.

## ASSUMPTION-0004 — reference platform is stable enough for validation

Assumption:

The current RetroPie/Bullseye environment can serve as the initial validation platform.

Risk:

Older OS packages may affect library availability or GPIO tooling.

---

# 12. Risks

## RISK-0001 — Unsafe shutdown regression

Replacing the original script could accidentally make shutdown less safe.

Mitigation:

- study original script
- validate hardware behavior
- test shutdown repeatedly
- preserve rollback path
- document disabling original script carefully

## RISK-0002 — Double shutdown handlers

Running SafeShutdown.py and RetroFlag Power at the same time could cause competing behavior.

Mitigation:

- document migration clearly
- detect original script where practical
- provide warning in diagnostics
- avoid enabling replacement until ready

## RISK-0003 — GPIO misidentification

Wrong GPIO mapping could cause missed shutdowns or false events.

Mitigation:

- validate with tools
- document exact commands
- include hardware validation checklist
- avoid hardcoding until confirmed

## RISK-0004 — Boot optimization regression

Boot-order, display, or service changes could break compatibility.

Mitigation:

- measure before optimizing
- document rollback
- separate research from recommendations

## RISK-0005 — Resume corruption

Future resume features could risk save-state or progress corruption.

Mitigation:

- safe shutdown first
- state model next
- resume only after validation
- avoid destructive defaults

---

# 13. Validation Checklist

Before replacing the original safe shutdown path, validate:

- [ ] Current SafeShutdown.py contents reviewed.
- [ ] Current rc.local startup line confirmed.
- [ ] Power Switch GPIO line identified.
- [ ] Reset Button GPIO line identified.
- [ ] Active-low/active-high behavior confirmed.
- [ ] Event edge behavior confirmed.
- [ ] Debounce behavior tested.
- [ ] Safe shutdown tested with original script.
- [ ] Safe shutdown tested with RetroFlag Power implementation.
- [ ] systemd service starts at boot.
- [ ] service logs visible in journald.
- [ ] service stops cleanly.
- [ ] original script disabled only after replacement is confirmed.
- [ ] rollback instructions documented.
- [ ] repeated shutdown tests performed.
- [ ] filesystem appears healthy after repeated shutdown tests.
- [ ] boot timing measured before and after replacement.

---

# 14. Suggested Discovery Commands

These commands are candidates for future hardware discovery and validation.

Do not treat output as known until captured and documented.

## Inspect current rc.local

```bash
cat /etc/rc.local
```

## Inspect original SafeShutdown.py

```bash
sudo sed -n '1,240p' /opt/RetroFlag/SafeShutdown.py
```

## Check service state

```bash
systemctl status rc-local.service
```

## Check boot time

```bash
systemd-analyze
systemd-analyze blame
```

## Check GPIO chips

```bash
gpiodetect
gpioinfo
```

## Watch GPIO events

Example only:

```bash
gpiomon gpiochip0 LINE_NUMBER
```

Replace `LINE_NUMBER` with a validated line.

## Check kernel and OS

```bash
uname -a
cat /etc/os-release
```

## Check boot config

```bash
cat /boot/config.txt
cat /boot/cmdline.txt
```

## Check EEPROM config

```bash
sudo rpi-eeprom-config
```

On CM4, EEPROM flashing should be treated carefully and may require `rpiboot`.

---

# 15. Migration Concept

The migration from original script to RetroFlag Power should eventually look like this:

## Stage 1 — Observe

- Review SafeShutdown.py
- Identify GPIO behavior
- Document current startup
- Do not change behavior yet

## Stage 2 — Build daemon

- Create daemon skeleton
- Add systemd unit
- Add logs
- Validate service lifecycle
- Still do not replace shutdown behavior

## Stage 3 — Simulate hardware

- Add mock hardware events
- Test Power Switch logic without real GPIO
- Test shutdown decision flow without executing shutdown

## Stage 4 — Validate real hardware

- Add GPi Case 2 hardware profile
- Add GPIO adapter
- Log detected events
- Do not execute shutdown until event detection is trusted

## Stage 5 — Controlled replacement

- Disable original script path
- Enable RetroFlag Power shutdown handling
- Test repeatedly
- Preserve rollback path

## Stage 6 — Document release behavior

- Installation instructions
- Uninstall instructions
- Troubleshooting
- Known limitations
- Recovery steps

---

# 16. Rollback Concept

A safe rollback plan is required before replacing the original shutdown path.

A future rollback should be able to restore:

```
/etc/rc.local
/opt/RetroFlag/SafeShutdown.py startup behavior
```

Potential rollback documentation should include:

- how to disable RetroFlag Power service
- how to re-enable original rc.local behavior
- how to verify original script is running
- how to inspect logs
- how to recover over SSH if display/frontend is unavailable

This rollback procedure must be written before public install instructions recommend replacement.

---

# 17. Hardware Support Boundaries

For now:

```
Official focus:
  RetroFlag GPi Case 2 with Raspberry Pi CM4
```

Later:

```
Potential future RetroFlag family:
  GPi Case 2W
  NESPi
  SUPERPi
```

Community / experimental:

```
Orange Pi
Radxa ROCK
ODROID
Banana Pi
Libre Computer
x64 handhelds
Steam Deck-like systems
```

Do not claim official support without validation hardware and documented behavior.

---

# 18. Relationship to Milestones

## Milestone 0 — Dreaming

This document belongs to Milestone 0 because it captures the first hardware witness before implementation begins.

## Milestone 1 — Awakening

The daemon can begin without full GPIO implementation.

This document helps ensure Milestone 1 does not accidentally begin by hardcoding hardware assumptions.

## Milestone 3 — Power

This document becomes critical during Power implementation.

Milestone 3 must convert current unknowns into validated facts.

## Milestone 8 — Expansion

This document becomes the model for future hardware reference documents.

---

# 19. What This Document Does Not Do

This document does not:

- provide final GPIO mappings
- provide final install commands
- provide a final systemd unit
- recommend EEPROM changes
- claim battery support
- claim sleep mode support
- define all future hardware profiles
- replace hardware validation

This document is testimony, not a verdict.

---

# 20. First Witness Statement

The GPi Case 2 says:

```
I am the first machine.

I have a latching Power Switch.

I have a Reset Button.

I currently rely on SafeShutdown.py.

I deserve a modern daemon.

But do not break my shutdown path.

Study me first.
```

The court accepts this testimony.

---

# Closing

The GPi Case 2 is the first witness because it is the first promise.

Before resume, polish, metrics, sleep research, battery overlays, or hardware expansion, the project must honor this machine.

The reference hardware has taken the stand.

The next step is not to guess.

The next step is to validate.
