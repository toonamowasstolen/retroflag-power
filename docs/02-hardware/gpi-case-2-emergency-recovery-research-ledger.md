---
id: HW-GPI-CASE-2-EMERGENCY-RECOVERY-RESEARCH-LEDGER-001
title: GPi Case 2 Emergency Recovery Research Ledger
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Gather public evidence and open questions for a reversible GPi Case 2 development reset or safe power-cut path without approving physical modification.
related:
  - gpi-case-2-hardware-findings-kms-power-notes.md
  - ../03-operations/gpi-case-2-acceptance-checklist.md
  - ../03-operations/safeshutdown-replacement-boundary-map.md
  - ../03-operations/gpi-case-gpio-probe-ledger.md
  - ../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md
  - ../04-architecture/arcadia-runtime-migration-path.md
last_updated: 2026-07-09
---

# GPi Case 2 Emergency Recovery Research Ledger

> This ledger is a lantern, not a knife. It gathers public clues for a future
> reversible recovery path while keeping the handheld intact and the field work
> honest.

This document is documentation only. It does not change Go code, read GPIO,
write GPIO, execute shutdown, install or activate systemd, alter `rc.local`,
replace `/opt/RetroFlag/SafeShutdown.py`, implement resume, add persistent
state, add telemetry, make project-code network calls, change installers or
packaging, or approve any hardware modification.

No candidate path below is ready-to-do. Any reset, enable, power-cut, solder,
test-pad, or regulator path remains investigation only until public evidence,
board mapping, electrical review, and a separate maintainer decision say
otherwise.

## Purpose

This ledger starts the research trail for a GPi Case 2 emergency development
recovery path after a power-save/resume field incident showed that ordinary
software recovery may disappear.

The goal is narrow:

- Gather public source links.
- Record source findings without inventing a schematic.
- Identify board-level unknowns.
- Name candidate recovery paths as investigation only.
- Mark unsafe paths as not approved.
- Define safe next questions that do not require live GPIO access, shutdown,
  service activation, or physical modification.

## Current Field Incident That Motivates This Research

On 2026-07-08, a GPi Case 2 field incident after power-save or resume behavior
showed repeated Linux `rcu: INFO: rcu_preempt detected stalls on CPUs/tasks`
messages. SSH and ping were unavailable. Moving the side switch off did not
shut down the handheld. The top power-save/resume button still changed the
visible state, but the only observed stop was physical CM4 cartridge/card
removal.

`SafeShutdown.py` was believed to be enabled, so this must not be dismissed as
only a disabled stock-script failure.

Waiting for battery depletion, repeatedly removing the CM4 cartridge/card, or
depending on SSH after power-save are not acceptable development recovery
plans. The recovery trail needs evidence before deeper power-save, resume, or
replacement-runtime tests depend on it.

On 2026-07-09, a later GPi Case 2 Bundle Collector Field Lantern run captured
a successful post-resume session:
`gpi-case2-bundle-collector-field-lantern-20260709-083407.tar.gz`.
That run happened after an unintended sleep followed by successful resume.
EmulationStation was visibly open and detected by the updated script, the
90-sample trace ran for about 102 seconds, `get_throttled` stayed `0x0`,
temperature stayed roughly 58-60 C, and the internal/core voltage sample stayed
around `0.8700V`.

This is a good field win, but it changes the theory carefully: the resume
wedge is intermittent, not guaranteed. Successful resume has now been
observed, while longer sleep duration, battery state, thermal state,
USB/input state, display/KMS timing, and transient power conditions remain
suspects. The post-resume satchel does not prove what happened during the
sleep/resume transition because no watcher was already running before sleep.
A late `xpad` USB `-19` line around uptime 2652 seconds is a trail marker,
not proof of root cause.

## Known Current Behavior

Known from the current EDC ledgers:

- The GPi Case 2 side power switch does not appear to directly cut battery
  power.
- The stock RetroFlag GPi Case 2 script watches BCM GPIO26 as `powerPin`.
- The stock RetroFlag GPi Case 2 script drives BCM GPIO27 HIGH as
  `powerenPin`.
- The stock script starts separate workers for side-switch shutdown and the
  `lcdrun()` display/power-save loop.
- `lcdrun()` repeatedly calls `/opt/RetroFlag/lcdnext.sh`.
- The legacy LCD scripts are not KMS-safe for the current project direction
  because they can rewrite display configuration.
- The top power-save/resume button and possible automatic display/audio
  power-save behavior remain unsafe/unverified after the RCU stall incident,
  even though one later successful resume has been captured.
- Raw signal vocabulary remains separate from interpretation:
  `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw observations.
  `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from a
  configured hardware profile.
- Lanterns are read-only probes and diagnostics.
- Lantern Dispatch remains future optional support infrastructure, not
  implemented.

## Public Sources Checked

| Source | URL | Checked for | Result |
| --- | --- | --- | --- |
| RetroFlag `retroflag-picase` README | <https://github.com/RetroFlag/retroflag-picase> | Official script entry point and GPi Case 2 link | Points GPi Case 2 users to the separate `GPiCase2-Script` repository. |
| RetroFlag `GPiCase2-Script` README | <https://github.com/RetroFlag/GPiCase2-Script> | Official GPi Case 2 safe shutdown, dock, and patch notes | Confirms the GPi Case 2 script handles safe shutdown and automatic LCD/HDMI switching; warns the patch should be installed before the script. |
| RetroFlag GPi Case 2 RetroPie script | <https://github.com/RetroFlag/GPiCase2-Script/blob/main/retropie_SafeShutdown_gpi2.py> | GPIO ownership and shutdown behavior | Shows BCM GPIO26 as `powerPin`, BCM GPIO27 as `powerenPin`, GPIO27 driven HIGH, side-switch edge wait, shutdown command, and `lcdrun()` calling `lcdnext.sh`. |
| RetroFlag GPi Case 2 RetroPie installer | <https://github.com/RetroFlag/GPiCase2-Script/blob/main/retropie_install_gpi2.sh> | Installation behavior and startup ownership | Downloads `SafeShutdown.py`, `lcdfirst.sh`, `lcdnext.sh`, writes `/etc/modprobe.d/alsa-base.conf`, edits `/etc/rc.local`, and reboots. This ledger does not authorize those actions. |
| RetroFlag GPi Case 2 LCD switching script | <https://github.com/RetroFlag/GPiCase2-Script/blob/main/retropielcdnext.sh> | Display switching behavior | Public script exists; detailed KMS implications remain a local review item before any replacement plan. |
| Raspberry Pi Compute Module 4 datasheet | <https://datasheets.raspberrypi.com/cm4/cm4-datasheet.pdf> | CM4 reset, `GLOBAL_EN`, power sequencing, and power-down cautions | Documents `GLOBAL_EN`, `RUN_PG`, power sequencing, and filesystem consistency cautions before removing power. |
| Public web searches | Search terms: `"GPi Case 2" "teardown"`, `"GPi Case 2" "schematic"`, `"GPi Case 2" "board" "test pad"`, `"GPi Case 2" "GLOBAL_EN"` | Schematics, teardown photos, board labels, test pads, regulator enables | No reliable public schematic or board-level reset/power-enable map found in this pass. Public sources are scarce. |

## Source Findings

RetroFlag's public GPi Case 2 script repository says the GPi Case 2 CM4 safe
shutdown script automatically switches between the LCD display and HDMI output
when using the dock. Its README also warns to install the GPi Case 2 patch
before the script because installing in the wrong order can damage the boot
configuration enough to prevent boot.

The RetroPie GPi Case 2 script defines BCM GPIO26 as the side-switch
`powerPin`, defines BCM GPIO27 as `powerenPin`, configures GPIO27 as an output,
and drives it HIGH. It waits for a falling edge on GPIO26 before killing
EmulationStation, sleeping, and calling `shutdown -h now`. A second worker
loops over `/opt/RetroFlag/lcdnext.sh`.

The RetroPie installer writes into `/opt/RetroFlag`, installs LCD scripts,
writes an ALSA modprobe file, edits `/etc/rc.local`, and reboots. These
behaviors are useful evidence of script ownership, but they are not actions for
this project to perform in this quest.

The CM4 datasheet documents two relevant module-level paths:

- `GLOBAL_EN`: pulling it low puts the CM4 into its lowest power-down state.
  The datasheet recommends pulling it low only after the OS has shut down.
- `RUN_PG`: when high, it signals the CM4 has started. Driving it low resets
  the module, and the datasheet cautions that open filesystem files will not be
  closed.

The CM4 datasheet also says the operating system should be shut down before
power is removed to keep the filesystem consistent. If that cannot be
guaranteed, filesystem strategies such as overlay-oriented or resilient
filesystems may need separate evaluation. That is future architecture research,
not an emergency mod approval.

No public GPi Case 2 schematic, board-level reset map, regulator-enable map,
lithium charge-controller map, or clearly labeled test-pad map was found in
this pass.

## Public GitHub Activity Notes

These notes record public Git metadata only. They do not prove current vendor
support, hardware safety, or replacement readiness.

| Repository | Public commit span checked | Commit count checked | Latest checked commit | Notes |
| --- | --- | ---: | --- | --- |
| <https://github.com/RetroFlag/GPiCase2-Script> | 2021-12-15 to 2023-01-14, about 1 year and 30 days | 20 | `de498e9` on 2023-01-14, `Update README.md` | The GPi Case 2 repository was added with an initial commit on 2021-12-15. The selected RetroPie GPi2 script files checked in this ledger were added on 2021-12-15 and last touched in that file set on 2021-12-30. |
| <https://github.com/RetroFlag/retroflag-picase> | 2018-01-08 to 2021-12-28, about 3 years and 11 months | 58 | `fb69441` on 2021-12-28, `Update README.md` | This older RetroFlag safe-shutdown repository points GPi Case 2 users to the separate `GPiCase2-Script` repository. |

`git ls-remote --tags` returned no public tags for either repository in this
research pass, so this ledger does not record a release or version cadence.

## Board-Level Unknowns

These unknowns must stay unknown until evidence exists:

- Whether the GPi Case 2 exposes CM4 `RUN_PG`, `GLOBAL_EN`, `EXT_nRESET`, or
  another reset-like signal on a labeled pad.
- Whether any visible pad is a reset, enable, power-good, regulator enable,
  battery sense, charge-control, or dock-detect signal.
- Whether BCM GPIO27 controls a GPi Case 2 board latch, a regulator enable, a
  soft power gate, or a higher-level microcontroller input.
- Whether the top power-save/resume button is routed to CM4 GPIO, a case board
  controller, a regulator, the LCD/backlight path, or another power domain.
- Whether the side switch remains readable or actionable during the RCU stall
  state.
- Whether a CM4-level reset would recover from the observed stall or merely
  increase filesystem risk.
- Whether a module-level `GLOBAL_EN` path would fully power-cycle the relevant
  case board domains.
- Whether the GPi Case 2 power board can be interrupted safely upstream or
  downstream of the regulator without touching lithium battery or charge
  circuitry.
- Whether docking state changes the recovery-relevant power or reset paths.
- Whether any public board photo shows usable labels at sufficient resolution.

## Candidate Emergency Recovery Paths To Investigate

All candidate paths in this section are investigation only and not approved.

| Candidate path | Why it might matter | Evidence needed before action |
| --- | --- | --- |
| CM4 `RUN_PG` or reset path | A module reset might recover a wedged Linux state when SSH and side-switch shutdown are unavailable. | Confirm exact CM4 signal access on the GPi Case 2 board, electrical behavior, required pulse shape, filesystem risk, and whether it is exposed without blind soldering. |
| CM4 `GLOBAL_EN` path | The CM4 datasheet identifies `GLOBAL_EN` as a low-power control path after software shutdown. | Confirm whether the GPi Case 2 exposes this signal, whether it affects only the CM4 or the whole case power domain, and how it interacts with the case latch and battery board. |
| Regulator enable path | A board regulator enable might provide a reversible development power interruption without touching the battery. | Identify the regulator part, datasheet, enable pin, surrounding circuit, load domains, and whether disabling it can backfeed or corrupt storage. |
| Case-board power-enable/latch path | GPIO27 appears to hold some power-enable path HIGH under the stock script. | Map what GPIO27 drives, whether a hardware latch exists, and whether there is a safe external recovery point that does not fight the stock script or the case board. |
| External development-only power interruption upstream of the CM4 regulator | A controlled fixture might avoid repeated CM4 cartridge/card removal during lab recovery. | Prove where the CM4 +5V path enters, isolate it from lithium charging, verify current rating and connector safety, and document filesystem-risk boundaries. |
| External development-only power interruption downstream of the battery/charge circuit and upstream of load regulation | A reversible fixture might cut load power while leaving the charging circuit unmodified. | Board map, charge-controller map, current path proof, backfeed analysis, fuse/protection review, and maintainer approval. |

## Disallowed Or Unsafe Paths

These paths are explicitly not approved:

- Cutting battery leads.
- Modifying lithium battery circuitry.
- Modifying lithium charging circuitry.
- Blind soldering.
- Shorting unknown pads.
- Treating any unlabeled pad as reset, `RUN_PG`, `GLOBAL_EN`, power enable, or
  regulator enable without proof.
- Soldering as a recommendation. Any soldering mention belongs only in future
  investigation notes after mapping and review.
- Relying on repeated CM4 cartridge/card removal.
- Waiting for battery depletion as a recovery plan.
- Assuming a CM4 reset is safe because it is convenient.
- Assuming a power cut is safe while Linux may have open filesystem writes.
- Running the RetroFlag installer or stock scripts as part of this research
  quest.

## Evidence Needed Before Any Development Mod

Before any development-only mod is even proposed, the ledger needs:

- A public schematic, official board documentation, or maintainer-created board
  map from visual inspection with clear evidence-asset rules.
- High-resolution photos of both sides of the GPi Case 2 CM4 carrier and power
  board, with visible labels and component markings.
- Identified regulator, charge-controller, latch, and any microcontroller or
  power-management components.
- Exact mapping from CM4 pins to GPi Case 2 board nets for `GLOBAL_EN`,
  `RUN_PG`, `EXT_nRESET`, +5V, CM4 3.3V, CM4 1.8V, BCM GPIO26, and BCM GPIO27.
- A written electrical review of voltage levels, pull-ups, pull-downs, current
  limits, backfeed risk, and storage-corruption risk.
- A reversible bench-only procedure that does not touch battery leads or charge
  circuitry.
- A recovery decision tree that says when to prefer software shutdown, when to
  stop testing, and when no safe recovery path exists.
- Maintainer approval in a separate quest before any physical change.

## Open Questions

- Does any official RetroFlag GPi Case 2 schematic, service manual, or board
  pinout exist outside the public GitHub scripts?
- Are there trustworthy teardown photos that show board labels, test pads,
  regulator part numbers, or a case-board controller?
- Is BCM GPIO27 a direct latch enable, a signal into another controller, or
  something else?
- Does the top power-save/resume button interact with Linux, the display path,
  a microcontroller, or pure power hardware?
- During the RCU stall state, was storage still active, read-only, idle, or
  unknown?
- Would a CM4 `RUN_PG` reset recover display/network state, and what filesystem
  damage risk would it carry?
- Would a `GLOBAL_EN` event reset only the CM4 or also enough of the case board
  to recover the visible power-save state?
- Can a future development fixture interrupt CM4 load power without modifying
  lithium battery or charging circuitry?
- Does the dock expose any safer recovery or power path than handheld mode?

## Next Safe Field Steps

These steps are safe because they are documentation, source review, or passive
visual inspection only:

- Continue public-source research for official RetroFlag manuals, support
  downloads, board photos, and any GPi Case 2 schematic or pinout.
- Review RetroFlag GPi Case 2 scripts locally as text only; do not run them.
- If evidence-asset rules allow, add maintainer-taken board photos to the
  appropriate evidence area and link them from this ledger.
- Record visible board markings, component labels, connector labels, and test
  pad labels from photos without probing them.
- Compare any visible component markings against datasheets before assigning
  meaning.
- Keep the acceptance checklist marked no-go for power-save/resume replacement
  until a recovery plan exists.
- Use the
  [GPi Case 2 Bundle Collector Lantern Capture Procedure](../03-operations/gpi-case-2-boot-power-trace-capture-procedure.md)
  for post-resume satchels when the device is already responsive, while
  recording that post-resume captures do not prove transition-time behavior.
- Keep any future read-only GPIO lantern work separate from this ledger and
  preserve the `SignalLow`/`SignalHigh`/`SignalUnverified` vocabulary boundary.
- Map a future Session Watch Lantern that records pre-sleep state, records
  post-resume state when available, tracks `get_throttled`, temperature,
  frontend, and input hints over time, and avoids telemetry and automatic
  fixes.

The next badge is not a mod. The next badge is a clearer map.
