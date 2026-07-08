---
id: HW-GPI-CASE-2-DEVELOPER-ACCESS-PATHS-001
title: GPi Case 2 Developer Access Paths
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map candidate GPi Case 2 developer diagnosis and recovery access paths without approving hardware modification.
related:
  - gpi-case-2-hardware-findings-kms-power-notes.md
  - gpi-case-2-power-integrity-investigation-notes.md
  - gpi-case-2-emergency-recovery-research-ledger.md
  - ../03-operations/gpi-case-2-recovery-first-field-procedure.md
  - ../03-operations/gpi-case-2-replacement-coverage-matrix.md
  - ../03-operations/gpi-case-2-safeshutdown-script-behavior-map.md
  - ../03-operations/gpi-case-2-acceptance-checklist.md
  - ../03-operations/safeshutdown-replacement-boundary-map.md
  - ../03-operations/gpi-case-gpio-probe-ledger.md
  - ../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md
last_updated: 2026-07-08
---

# GPi Case 2 Developer Access Paths

> This map looks for doors before anyone reaches for tools. Every path below
> is investigation only until evidence, review, and a separate maintainer
> decision open it.

This document is documentation only. It does not change Go code, read GPIO,
write GPIO, execute shutdown, install or activate systemd, alter `rc.local`,
replace `/opt/RetroFlag/SafeShutdown.py`, implement resume, add persistent
state, add telemetry, make project-code network calls, change installers or
packaging, run RetroFlag install scripts, flash firmware, or approve hardware
modification.

Nothing here approves cutting battery leads, modifying lithium battery or
charging circuitry, blind soldering, shorting unknown pads, running unknown
vendor tools, repeated CM4 cartridge/card removal as routine recovery,
treating battery depletion as a recovery plan, or any soldering except as
future investigation after mapping and approval.

## Purpose

Map possible GPi Case 2 developer diagnosis and recovery access paths for
low-level power, display, kernel, and board-control failures.

The map is a readiness tool, not a procedure. It helps future maintainers
decide what evidence must exist before opening the case, probing pads,
considering UART, considering reset or enable access, or considering any
power-cut path.

Vocabulary boundaries remain active:

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from a
  configured hardware profile.
- Lanterns are read-only probes and diagnostics.
- Lantern Dispatch remains future optional support infrastructure, not
  implemented.

## Why This Investigation Exists

The GPi Case 2 KMS display path now works, but the power-save and recovery
trail is not safe enough for deeper replacement work.

The current EDC evidence says:

- The side switch appears to be software-mediated, not a direct battery cutoff.
- `SafeShutdown.py` uses BCM GPIO26 as `powerPin` and drives BCM GPIO27 HIGH as
  `powerenPin`.
- The top power-save/resume button is not proven to be a hard reset or Linux
  input.
- The GPi Case 2 can auto-enter display or audio power-save after roughly
  15-20 minutes idle.
- Recovery diagnostics found undervoltage/throttling evidence while the device
  was being treated as battery-powered.
- Public GPi Case 2 board-level evidence is scarce. No reliable public
  schematic or reset, enable, regulator, lithium charge, or test-pad map is
  recorded in the current EDC.

## Current Field Problem

The 2026-07-08 field incident recorded in
[QUEST-0064](../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md)
showed a dangerous recovery gap:

- Repeated Linux `rcu: INFO: rcu_preempt detected stalls on CPUs/tasks`
  messages after power-save or resume behavior.
- SSH unavailable.
- Ping unavailable.
- Side switch moved off, but the device did not shut down.
- `SafeShutdown.py` was believed to be enabled.
- Top power-save/resume button still changed visible state.
- Visible sleep or display-off behavior with a blinking LED still appeared to
  work during the stall. This suggests some board-controlled behavior may have
  remained active, but that is an inference, not schematic proof.
- The only observed stop was physical CM4 cartridge/card removal.

That stop is not acceptable as routine recovery. A safer development access
map must be earned before deeper power-save, resume, dock, or replacement
runtime testing depends on ordinary software paths surviving.

## What Normal User Paths Cannot Do During A Kernel Stall

When the kernel or userspace is wedged, normal user paths may disappear:

- SSH can be gone.
- Ping can be gone.
- The side switch may not trigger shutdown because the stock script depends on
  Linux userspace.
- The top button may only change display or board-visible state, not recover
  Linux.
- Controller input may not reach a useful process.
- Docking, undocking, or display transitions may add evidence, but they are not
  proven recovery paths.
- Battery depletion and repeated CM4 cartridge/card removal carry filesystem,
  evidence-loss, and hardware-handling risk.

Software shutdown cannot be assumed once RCU stalls or similar kernel lockups
are visible.

## Likely Factory Or Developer Debug Paths

The paths below are likely classes of developer or factory access, but their
presence on the GPi Case 2 carrier or case board is not proven unless marked
by evidence. Treat each item as inference until a public source, board photo,
datasheet, labeled connector, or maintainer-created map confirms it.

| Candidate path | Current status | Why it might matter |
| --- | --- | --- |
| CM4 UART serial console | Inference | Could provide boot and kernel logs when display, network, or userspace fails. |
| CM4 `RUN`/reset path | Inference | Could reset the CM4 when Linux is wedged, but may risk filesystem damage. |
| CM4 `GLOBAL_EN` or carrier-board enable path | Inference from CM4 docs, unmapped on GPi Case 2 | Could affect module power state, but may not reset all case board domains and should follow OS shutdown when possible. |
| Regulator enable path | Inference | Could control a load rail without touching lithium circuitry, but only if the regulator and enable circuit are identified. |
| Current-limited bench power | Future bench investigation only | Could characterize current draw or brownout behavior, but requires a safe injection point and no battery or charge-circuit modification. |
| Battery-disconnect-safe service method | Unknown | A vendor-supported service disconnect, if one exists, could be safer than improvised power interruption. |
| Hidden firmware/update USB port | Observed candidate, inconclusive | May expose an updater, MCU, bridge, or no normal USB device at all. |
| Raspberry Pi `rpiboot`/`usbboot` observation-only mode | Inference | Could reveal whether the CM4 USB boot path is exposed, but this pass must not flash or send payloads. |
| Case-board MCU firmware/update path | Unknown | If a case-board controller exists and has a public updater path, it may explain power-save, LED, dock, or display behavior. |
| Dock or external USB behavior | Known physical path, recovery role unknown | Docked USB, HID, storage, or power behavior may give observation routes or change failure modes. |
| HDMI/docked logging possibilities | Unknown | Docked HDMI or local console may preserve visibility when the handheld LCD path is dark. |

## Candidate Access Paths

### CM4 UART Serial Console

Potential value:

- Capture boot logs, kernel messages, RCU stalls, undervoltage warnings, and
  early panic output without relying on Wi-Fi, SSH, KMS display, or
  EmulationStation.

Evidence needed before use:

- Which CM4 UART pins, if any, are exposed on the GPi Case 2 carrier.
- Voltage level and ground reference.
- Whether the pads are labeled and reachable without blind soldering.
- Whether any case-board controller shares the same lines.
- A non-invasive connection plan and logging-only procedure.

Current stance: future investigation only. No probing, soldering, or console
connection is approved by this map.

### CM4 RUN Or Reset Path

Potential value:

- Reset the CM4 when Linux is wedged and normal software paths are gone.

Evidence needed before use:

- Whether GPi Case 2 exposes a CM4 reset-like signal.
- Exact signal name, voltage behavior, pull-ups, pulse requirements, and
  interaction with the case board.
- Filesystem risk review for resetting while Linux may have open writes.
- Proof that the path is not an unknown test pad or unrelated enable signal.

Current stance: investigation only. A reset action is not approved.

### CM4 GLOBAL_EN Or Carrier-Board Enable Path

Potential value:

- Control a CM4 low-power or carrier-board enable state after a clean shutdown,
  if the path is exposed and understood.

Evidence needed before use:

- Whether `GLOBAL_EN` or a carrier-board enable signal is exposed on GPi Case 2.
- Whether it affects only the CM4, the case board, the LCD/backlight, USB
  devices, dock state, or other domains.
- Electrical review, storage-risk review, and interaction with GPIO27
  `powerenPin`.

Current stance: investigation only. Do not pull any enable/reset-like line.

### Regulator Enable Path

Potential value:

- A mapped regulator enable might provide a controlled bench-only way to
  recover from a wedged load state without touching battery leads or charge
  circuitry.

Evidence needed before use:

- Regulator part number, datasheet, input/output rails, enable pin, load
  domains, backfeed risks, current limits, and board-level surrounding circuit.
- Proof that disabling the regulator will not fight another controller or leave
  unsafe partially powered domains.

Current stance: investigation only. No regulator probing or interruption is
approved.

### Current-Limited Bench Power

Potential value:

- Characterize power draw, brownout, and undervoltage behavior during boot,
  idle, display, audio, dock, and power-save states.

Evidence needed before use:

- Safe power entry point that does not bypass, modify, or fight lithium battery
  charging circuitry.
- Voltage and current limits.
- Whether battery must be disconnected through a documented service method.
- A stop condition for undervoltage, heat, boot loops, or storage errors.

Current stance: future bench plan only. No bench-power injection is approved.

### Battery-Disconnect-Safe Service Method

Potential value:

- A documented vendor service method could reduce reliance on battery
  depletion or repeated CM4 cartridge/card removal.

Evidence needed before use:

- Official service documentation, labeled connector evidence, or maintainer
  board map showing a safe disconnect point.
- Proof that disconnecting does not require cutting leads, modifying lithium
  battery circuitry, modifying charging circuitry, or shorting anything.

Current stance: unknown. Do not improvise a battery disconnect.

### Hidden Firmware Or Update USB Port

Potential value:

- A hidden update port may expose a case-board controller, USB bridge, firmware
  updater, recovery mode, or factory diagnostic mode.

Current evidence:

- A first observation attempt did not show an obvious new Mac USB device in
  `system_profiler` or `ioreg`.
- Treat that result as inconclusive. It is not proof the port is useless.

Possible explanations:

- Wrong mode.
- Charge-only cable.
- Device not powered.
- Port only active during firmware or update mode.
- Port not connected as normal USB.
- Board/controller requires a special updater.

Evidence needed before use:

- Repeatable before/after Mac-side device captures.
- Cable identity and power state notes.
- Any public RetroFlag updater documentation or tool provenance.
- Proof that no firmware payload or updater command is being sent.

Current stance: observation-only. Do not run unknown vendor tools or updater
commands.

### Raspberry Pi rpiboot Or usbboot Observation-Only Mode

Potential value:

- Determine whether the CM4 USB boot path is exposed through the GPi Case 2,
  dock, hidden port, or another connector.

Safe boundary:

- Waiting with `rpiboot`/`usbboot` for an attach event is the only mapped
  concept here.
- Do not send boot payloads.
- Do not flash EEPROM, firmware, boot media, or case-board firmware.
- Do not alter boot mode straps or pads.

Evidence needed before use:

- A clear command transcript showing wait-only behavior.
- Exact cable, port, dock, power, LED, and display states.
- Confirmation that no write, boot image, or updater payload was sent.

Current stance: observation-only candidate.

### Case-Board MCU Firmware Or Update Path

Potential value:

- If a case-board MCU exists and controls sleep, LED, dock, audio, display, or
  power behavior, its update or diagnostic path may explain the RCU stall
  surroundings or recovery limits.

Evidence needed before use:

- Public evidence of the controller, firmware path, updater, protocol,
  supported hardware revision, and failure recovery mode.
- Tool provenance and a no-write inspection mode.
- Separate maintainer approval before running any updater.

Current stance: unknown and not approved.

### Dock Or External USB Behavior

Potential value:

- Docked USB or external devices may provide local keyboard, serial adapter,
  storage, logging, or a different power/display state during failures.

Evidence needed before use:

- Whether docked mode changes power domains, USB topology, controller identity,
  HDMI state, charging state, or recovery behavior.
- Safe observation captures before and after dock attach.
- Stop criteria if dock transitions trigger stalls, display loss, or
  filesystem errors.

Current stance: observation-only if no risky power-save/resume transition is
induced.

### HDMI Or Docked Logging Possibilities

Potential value:

- A docked HDMI console, local keyboard, or already-attached display path may
  reveal kernel logs when handheld LCD, Wi-Fi, or SSH is gone.

Evidence needed before use:

- Whether docked HDMI is KMS-safe without legacy `lcdnext.sh` or `lcdfirst.sh`
  rewrites.
- Whether a keyboard is available before the failure starts.
- Whether console output remains visible during power-save, wake, and stall.

Current stance: observation-only. Do not use docking as a forced recovery
ritual until transitions are safer.

## Evidence Needed Before Use

Before any candidate path becomes a procedure, collect:

- Public schematic, official service document, high-resolution board photos, or
  a maintainer-created board map.
- Clear labels or component markings for connectors, pads, regulators,
  controllers, charge components, and power-path parts.
- Exact net mapping for CM4 UART, `RUN`, `GLOBAL_EN`, `EXT_nRESET`, GPIO26,
  GPIO27, +5V, 3.3V, 1.8V, ground, battery, charge, dock, and USB paths.
- Voltage levels, current limits, pull-ups, pull-downs, backfeed risk, and
  storage-corruption risk.
- A written observation plan with stop criteria.
- A separate maintainer decision before physical modification, probing live
  pads, flashing, or changing any runtime behavior.

## Safe Observation-Only Tests

These tests are safe only if they remain observation-only and do not send
firmware payloads, updater commands, GPIO operations, shutdown commands, or
configuration changes:

- Mac-side `system_profiler` before and after hidden USB connection.
- Mac-side `ioreg` capture before and after hidden USB connection.
- Raspberry Pi `usbboot`/`rpiboot` waiting test only, with no flashing and no
  boot payload.
- Repeat hidden USB observation with a known data cable.
- Repeat hidden USB observation with powered, off, on, docked, and handheld
  states explicitly noted.
- Capture exact LED, display, charger, dock, and power state during each
  attempt.
- Note absence of a device as evidence, not as proof of no path.

Do not run firmware updaters, unknown vendor tools, or commands that write to
the device.

## Explicitly Disallowed Actions

This map does not approve:

- Cutting battery leads.
- Modifying lithium battery circuitry.
- Modifying lithium charging circuitry.
- Blind soldering.
- Shorting unknown pads.
- Flashing firmware.
- Running unknown vendor tools.
- Sending firmware payloads or updater commands.
- Treating unlabeled pads as UART, reset, enable, regulator, charge, or battery
  paths.
- Repeated CM4 cartridge/card removal as routine recovery.
- Treating battery depletion as a recovery plan.
- Using docking, top-button cycling, or side-switch cycling as a stress ritual
  after RCU stalls appear.
- Replacing `SafeShutdown.py`, editing `rc.local`, activating systemd, writing
  GPIO, reading GPIO, or executing shutdown from this investigation map.

## Developer Readiness Gates

### Before Opening The Case

Know:

- Why opening is necessary and why external observation is insufficient.
- What evidence will be collected and where it will be recorded.
- Whether warranty, connector, ribbon, battery, or static risks are acceptable.
- How photos will be captured without moving, cutting, soldering, or probing
  anything.
- Stop criteria if the board layout, battery path, or connectors are unclear.

### Before Probing Pads

Know:

- Pad identity from public docs, visible labels, or board mapping.
- Expected voltage, direction, owner, and safe measurement reference.
- Whether the pad is connected to CM4, case-board MCU, regulator, battery,
  charge, dock, display, or USB circuitry.
- Probe type, current limits, ESD precautions, and backfeed risks.
- Why passive photo and continuity-free research is not enough.

No pad probing is approved by this map.

### Before Considering UART

Know:

- Exact UART pins and ground.
- Voltage level and adapter compatibility.
- Whether the path is console-only and receive/transmit direction is safe.
- Whether connecting a serial adapter could backfeed the case board.
- How logs will be captured without changing boot or firmware state.

UART remains future investigation only.

### Before Considering Reset Or Enable Access

Know:

- Whether the signal is CM4 `RUN`, `GLOBAL_EN`, `EXT_nRESET`, a regulator
  enable, a case-board MCU input, or something else.
- Required electrical behavior and owner.
- Filesystem risk if Linux is wedged or writing.
- Whether it resets only the CM4 or also enough of the case board to recover
  display, USB, dock, LED, and power domains.
- Whether using it could fight GPIO27 `powerenPin` or another latch.

No reset or enable action is approved by this map.

### Before Considering Any Power-Cut Path

Know:

- The full current path from battery, charge circuit, regulators, CM4, display,
  USB, dock, and case-board controller.
- Which domains remain powered and whether backfeed can occur.
- Whether the path touches lithium battery or charging circuitry. If it does,
  stop.
- Filesystem and data-loss risks.
- Current rating, connector rating, fuse/protection behavior, heat behavior,
  and recovery procedure.
- Separate maintainer approval for a reversible bench-only procedure.

No power-cut path is approved by this map.

## Open Questions

- Does RetroFlag publish a GPi Case 2 service manual, schematic, board pinout,
  updater guide, or hidden-port description outside the currently checked
  public script repositories?
- Is the hidden USB/update connector connected to the CM4, a case-board MCU, a
  USB bridge, charging hardware, or another controller?
- Does any board label identify UART, reset, boot, enable, firmware, dock, or
  regulator signals?
- Does the GPi Case 2 expose CM4 `RUN`, `GLOBAL_EN`, `EXT_nRESET`, or USB boot
  mode in a documented way?
- What exactly does BCM GPIO27 hold enabled?
- Does the blinking LED during the RCU stall come from a case-board controller,
  Linux, a power circuit, or another domain?
- Does docked mode provide any safer local logging path than handheld mode?
- Is there a documented battery-disconnect-safe service method that does not
  cut leads or modify charging circuitry?
- Can a current-limited bench setup ever be used without bypassing or fighting
  lithium charging behavior?

## Next Field Steps

Keep the next pass calm and observational:

- Repeat hidden USB observation with a known data cable and exact Mac-side
  `system_profiler`/`ioreg` before-and-after captures.
- If `rpiboot`/`usbboot` is tested, use wait-only observation and document that
  no payload or flash action was sent.
- Record powered, off, on, docked, handheld, LED, display, charger, and cable
  state for every hidden-port attempt.
- Search public RetroFlag support pages, manuals, updater downloads, teardown
  photos, and board photos for GPi Case 2 developer or service access clues.
- Add board photos only under clear evidence-asset rules, and record visible
  labels without assigning meaning to unknown pads.
- Keep the recovery-first field procedure active for future RCU stall or
  no-network incidents.
- Keep power-save, resume, automatic idle power-save, and replacement-runtime
  work paused until the emergency recovery path is stronger than SSH,
  side-switch shutdown, battery depletion, or CM4 cartridge/card removal.

The next badge is not opening the handheld. The next badge is knowing which
door, if any, is real.
