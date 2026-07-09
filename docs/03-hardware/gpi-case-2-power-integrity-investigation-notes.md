---
id: HW-GPI-CASE-2-POWER-INTEGRITY-INVESTIGATION-NOTES-001
title: GPi Case 2 Power Integrity Investigation Notes
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record GPi Case 2 undervoltage and throttling evidence and define a safe read-only power-integrity investigation checklist.
related:
  - gpi-case-2-hardware-findings-kms-power-notes.md
  - gpi-case-2-emergency-recovery-research-ledger.md
  - ../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md
  - ../03-operations/gpi-case-2-boot-power-trace-lantern-map.md
  - ../03-operations/gpi-case-2-recovery-first-field-procedure.md
  - ../03-operations/gpi-case-2-replacement-coverage-matrix.md
  - ../03-operations/gpi-case-2-acceptance-checklist.md
  - ../03-operations/safeshutdown-replacement-boundary-map.md
  - ../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md
last_updated: 2026-07-08
---

# GPi Case 2 Power Integrity Investigation Notes

> The battery trail now has a real warning light. Keep the next pass
> observation-only until the power headroom is mapped.

This document is documentation only. It does not change Go code, read GPIO,
write GPIO, execute shutdown, install or activate systemd, alter `rc.local`,
replace `/opt/RetroFlag/SafeShutdown.py`, implement resume, add persistent
state, add telemetry, make project-code network calls, change installers or
packaging, run RetroFlag install scripts, or approve hardware modification.

Nothing here approves cutting battery leads, modifying lithium battery or
charging circuitry, blind soldering, shorting unknown pads, repeated CM4
cartridge/card removal as routine recovery, or treating battery drain as a
recovery plan. Soldering remains future investigation only and is not approved.

## Purpose

Record the new GPi Case 2 diagnostic evidence that undervoltage and throttling
signals appeared even while the device was being treated as battery-powered
during investigation.

This note gives future field work a focused checklist for power-integrity
evidence before power-save, resume, side-switch, or replacement-runtime claims
move forward.

The future startup-timing capture for this evidence is mapped in
[GPi Case 2 Boot Power Trace Lantern Map](../03-operations/gpi-case-2-boot-power-trace-lantern-map.md).
The manual capture path lives in
[GPi Case 2 Boot Power Trace Capture Procedure](../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md).
Use that procedure when the next question is when the undervoltage appears
during boot rather than whether any power warning has ever appeared.

Vocabulary boundaries still apply:

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from a
  configured hardware profile.
- Lanterns are read-only probes and diagnostics.
- Lantern Dispatch remains future optional support infrastructure, not
  implemented.

## Diagnostic Evidence Captured

Recovery diagnostics after the GPi Case 2 power-save, no-network, and
no-side-switch-response investigation showed power warnings:

- Kernel log showed `hwmon hwmon1: Undervoltage detected!`.
- Kernel later showed `Voltage normalised`.
- `vcgencmd get_throttled` returned `0x50000`.
- The device was being treated as battery-powered during this investigation.

These observations do not prove a single cause. They do make power integrity a
real suspect in the incident chain.

## Interpretation

- Undervoltage does not prove GPi Case 2 board design failure by itself.
- Undervoltage is now a real suspect in the RCU stall, no-network,
  no-side-switch-response incident chain.
- Software shutdown cannot be trusted if the kernel is already wedged.
- Arcadia Runtime may need to detect and report power-integrity warnings before
  enabling risky behavior such as power-save/resume testing, service
  replacement, or future shutdown orchestration.

Do not claim power-save or resume support until power integrity is
characterized and an emergency recovery path exists.

## Candidate Causes To Investigate

- Battery sag.
- Boost regulator droop.
- USB audio, controller, LCD, or CM4 transient load.
- Power-save or resume transition load.
- Aging battery.
- Case board power design limitation.
- KMS/display resume sensitivity.

Each candidate remains an investigation target, not a diagnosis.

## Safe Read-Only Follow-Up Captures

These captures are observation only. They do not write GPIO, execute shutdown,
replace scripts, run installers, or modify system configuration.

```sh
vcgencmd get_throttled
vcgencmd measure_volts
dmesg -T | grep -iE "under-voltage|undervoltage|voltage|thrott"
```

Interpret `vcgencmd` narrowly: `get_throttled` is the main firmware clue for
undervoltage and throttling, but it does not report watts, TDP, amps, power
draw, or actual 5V rail voltage. `measure_volts` reports an internal/core rail
reading, not the GPi Case 2 5V input rail.

Capture the same evidence:

- Before and after boot.
- Before and after idle.
- Before and after resume, if testing is resumed.

If `vcgencmd` or the kernel log path is unavailable, record that absence as
evidence instead of installing anything during the capture pass.

## Manual Observation Checklist

Record these facts beside each power reading:

- Battery charge level.
- Whether USB-C power is attached.
- Docked vs handheld.
- Screen on or off.
- Audio device active.
- Controller active.
- Time since boot.
- Time since idle.

## Power Test Matrix

All rows are observation-only. Any power-save or resume row is
unsafe/unverified until an emergency recovery path exists that does not depend
on SSH, ping, side-switch shutdown, battery depletion, or repeated CM4
cartridge/card removal.

| Mode | Condition | Capture | Status |
| --- | --- | --- | --- |
| Battery only | Fresh charge, cold boot | `vcgencmd get_throttled`, `vcgencmd measure_volts` internal/core rail reading, voltage/throttling log search, manual observations | Safe observation-only |
| Battery only | After 10 minutes idle | Same read-only captures and manual observations | Safe observation-only |
| Battery only | After screen power-save, if it can be observed safely | Same read-only captures and manual observations before and after wake | Unsafe/unverified until emergency recovery exists |
| USB-C power attached | Handheld mode | Same read-only captures and manual observations | Safe observation-only |
| Docked mode | Later safe docked pass | Same read-only captures and manual observations | Defer until docked recovery path is safe |
| Audio active vs idle | Handheld or docked mode being tested | Compare captures with audio active and idle | Safe observation-only if no power-save/resume is induced |
| Controller input active vs idle | Handheld or docked mode being tested | Compare captures with controller input active and idle | Safe observation-only if no power-save/resume is induced |

## Risk Impact On Replacement Planning

- Do not claim power-save or resume support until power integrity is
  characterized.
- Do not rely on side-switch shutdown as emergency recovery.
- Keep emergency reset and cutoff investigation open.
- Consider future diagnostics reporting throttling flags read-only.
- Keep replacement-runtime behavior gated until power, display, network,
  side-switch, and recovery evidence agree.

## Explicit Non-Approved Actions

These actions remain not approved:

- Cutting battery leads.
- Modifying lithium charging circuitry.
- Blind soldering.
- Shorting unknown pads.
- Repeated CM4 removal as routine recovery.
- Treating battery drain as a recovery plan.

The next badge is disciplined evidence: voltage flags, context, and recovery
notes gathered without poking the live power path.
