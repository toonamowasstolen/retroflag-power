---
id: OPS-GPI-CASE-2-SAFESHUTDOWN-SCRIPT-BEHAVIOR-MAP-001
title: GPi Case 2 SafeShutdown Script Behavior Map
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map the observed and inferred behavior of the GPi Case 2 SafeShutdown script path before RetroFlag Power replaces any part of it.
related:
  - gpi-case-2-acceptance-checklist.md
  - safeshutdown-replacement-boundary-map.md
  - gpi-case-gpio-probe-ledger.md
  - ../02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../02-hardware/gpi-case-2-emergency-recovery-research-ledger.md
  - ../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md
  - ../04-architecture/arcadia-runtime-migration-path.md
last_updated: 2026-07-08
---

# GPi Case 2 SafeShutdown Script Behavior Map

> This map names what the old GPi Case 2 spell appears to carry so the new
> lantern does not quietly leave a critical behavior behind.

This document is documentation only. It does not change Go code, read GPIO,
write GPIO, execute shutdown, install or activate systemd, alter `rc.local`,
replace `/opt/RetroFlag/SafeShutdown.py`, implement resume, add persistent
state, add telemetry, make project-code network calls, change installers or
packaging, run RetroFlag install scripts, or approve hardware modification.

## Purpose

This map records what the GPi Case 2 SafeShutdown script path appears to own,
what may belong to the GPi Case 2 board or MCU, and what remains unattributed.

The replacement rule is conservative: RetroFlag Power must preserve or
deliberately retire each behavior with evidence before it replaces the stock
path. The side switch, power-enable latch, display/power-save behavior, and
docking behavior must not be reduced to "just shutdown" without proof.

Vocabulary boundaries stay intact:

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from a
  configured hardware profile.
- Lanterns are read-only probes and diagnostics.
- Lantern Dispatch remains future optional support infrastructure, not
  implemented.

## Script Sources Inspected

Public upstream sources inspected:

- RetroFlag GPi Case 2 script repository:
  <https://github.com/RetroFlag/GPiCase2-Script>
- RetroFlag GPi Case 2 RetroPie SafeShutdown script:
  <https://github.com/RetroFlag/GPiCase2-Script/blob/main/retropie_SafeShutdown_gpi2.py>
- RetroFlag GPi Case 2 RetroPie installer:
  <https://github.com/RetroFlag/GPiCase2-Script/blob/main/retropie_install_gpi2.sh>
- RetroFlag GPi Case 2 RetroPie first LCD script:
  <https://github.com/RetroFlag/GPiCase2-Script/blob/main/retropielcdfirst.sh>
- RetroFlag GPi Case 2 RetroPie next LCD script:
  <https://github.com/RetroFlag/GPiCase2-Script/blob/main/retropielcdnext.sh>
- Older RetroFlag Pi case repository, which points GPi Case 2 users to the
  separate GPi Case 2 script repository:
  <https://github.com/RetroFlag/retroflag-picase>

Local `/opt/RetroFlag` files have not yet been captured in this repository.
This map therefore treats the public RetroPie GPi Case 2 scripts as upstream
evidence and the existing EDC field notes as local behavior evidence.

## Local Script Capture Pending

Local script capture should be read-only and performed later on the GPi Case 2
image. Do not run installers, modify files, activate services, or execute
shutdown as part of capture.

Safe read-only commands for later local capture:

```sh
ls -la /opt/RetroFlag
find /opt/RetroFlag -maxdepth 2 -type f -printf '%TY-%Tm-%Td %TH:%TM %p\n' | sort
sed -n '1,260p' /opt/RetroFlag/SafeShutdown.py
grep -RniE "GPIO|power|sleep|lcd|hdmi|shutdown|killall|poweren|powerPin|resetPin" /opt/RetroFlag 2>/dev/null
grep -nE "RetroFlag|SafeShutdown|python" /etc/rc.local
ps aux | grep -Ei "SafeShutdown|RetroFlag|lcd|hdmi" | grep -v grep
```

## Upstream Script Inspection Notes

The public GPi Case 2 README says the CM4 safe-shutdown script automatically
switches between the LCD and HDMI display when using the dock. It also warns
that the GPi Case 2 patch should be installed before the script because
installing in the wrong order can damage `config.txt` enough to prevent boot.

The public RetroPie `retropie_SafeShutdown_gpi2.py` script:

- Imports `RPi.GPIO`, `os`, `time`, and `multiprocessing.Process`.
- Sets BCM numbering.
- Configures `powerPin = 26` as an input with pull-up.
- Configures `powerenPin = 27` as an output and drives it HIGH.
- Waits for a falling edge on GPIO26.
- Runs EmulationStation kill commands, sleeps, then runs Linux shutdown.
- Starts a second process that repeatedly calls `/opt/RetroFlag/lcdnext.sh`
  every second.

The public RetroPie LCD scripts:

- Use GPIO18 as `HDMI_HPD`.
- `lcdfirst.sh` exports GPIO18, sets it as input, checks the value, rewrites
  `/boot/config.txt` between LCD and HDMI config files, and may reboot.
- `lcdnext.sh` reads GPIO18, checks whether `enable_dpi_lcd=1` is present in
  `/boot/config.txt`, and rewrites `/boot/config.txt` between saved LCD and
  HDMI config files.

The public RetroPie installer:

- Creates or uses `/opt/RetroFlag`.
- Downloads the SafeShutdown and LCD scripts.
- Downloads an ALSA modprobe configuration.
- Edits `/etc/rc.local` to start `lcdfirst.sh`, sleep briefly, and start
  `SafeShutdown.py`.
- Reboots.

These installer behaviors are evidence only. This quest does not authorize
running them.

## GPIO Pins Referenced

| BCM pin | Script name | Direction in upstream script | Current interpretation |
| --- | --- | --- | --- |
| 18 | `HDMI_HPD` | Input through `/sys/class/gpio` in LCD scripts | Dock or HDMI hot-plug detection path used by legacy LCD/HDMI switching. |
| 26 | `powerPin` | Input with pull-up | Side-switch shutdown request signal watched by the Linux script. Treat raw observations as `SignalLow`, `SignalHigh`, or `SignalUnverified` until profile interpretation maps them. |
| 27 | `powerenPin` | Output driven HIGH | Power-enable or latch-hold path. Exact board target is unknown. |

Commented references in the upstream SafeShutdown script include `ledPin = 14`
and `resetPin = 2`, but the inspected RetroPie GPi Case 2 script does not
enable those paths.

## Input Signals Observed

Observed or evidenced inputs:

- GPIO26 falling edge: treated by the upstream Python script as the trigger for
  safe shutdown.
- GPIO18 value: treated by the upstream LCD scripts as dock/HDMI presence for
  config switching.
- Top power-save/resume button: observed in field notes as changing visible
  sleep/wake state, but not yet attributed to a GPIO, input event, script path,
  case-board controller, or display power domain.
- Automatic idle power-save: the GPi Case 2 appears to auto-enter
  display/audio power-save after roughly 15-20 minutes idle. Avoiding the top
  power-save button is therefore not enough to avoid the power-save path.

During the 2026-07-08 RCU stall/no-network/no-side-switch-response state, the
device still appeared able to return to a sleep/display-off state with a
blinking indicator LED. This is an inference-bearing observation, not schematic
proof. It suggests that at least some sleep, display, or LED behavior may be
board-controlled or independent of normal Linux userspace.

## Output Pins Driven

The upstream Python script drives BCM GPIO27 HIGH during initialization. The
existing EDC calls this the power-enable latch path, but the exact board-level
target is unknown.

Do not infer that GPIO27 is safe to write from RetroFlag Power. A replacement
must first prove:

- What GPIO27 drives.
- Whether HIGH must be asserted before, during, or after boot.
- Whether LOW or release changes case power state.
- How the line behaves during shutdown, power-save, resume, docked use, and
  failure states.
- Whether a board MCU, latch, regulator, or other circuit also owns the path.

## Commands Run By The Script

The upstream SafeShutdown script runs:

- `sudo killall emulationstation`
- `sudo killall emulationstatio`
- `sudo sleep 5s`
- `sudo shutdown -h now`
- `sh /opt/RetroFlag/lcdnext.sh`

The second kill command appears to target a truncated process name for older
RetroPie behavior. It should be preserved as an observed upstream assumption,
not copied into new runtime behavior without field evidence.

## External Scripts Called

The upstream SafeShutdown script calls:

- `/opt/RetroFlag/lcdnext.sh` once per second from the `lcdrun()` worker.

The upstream `/etc/rc.local` installer path starts:

- `/opt/RetroFlag/lcdfirst.sh`
- `/opt/RetroFlag/SafeShutdown.py`

The LCD scripts are not KMS-safe for the current project direction because
they can rewrite `/boot/config.txt` and switch between saved LCD/HDMI config
files. RetroFlag Power must not call them from its runtime path.

## Timing And Sleep Assumptions

Observed timing assumptions:

- `SafeShutdown.py` waits indefinitely for a GPIO26 falling edge.
- After the GPIO26 edge, it kills EmulationStation, waits about five seconds,
  then calls shutdown.
- `lcdrun()` calls `lcdnext.sh` every second forever.
- The installer inserts a short sleep between `lcdfirst.sh` and
  `SafeShutdown.py` in `/etc/rc.local`.
- The GPi Case 2 appears to auto-enter display/audio power-save after roughly
  15-20 minutes idle.

Replacement work must treat these as behavior clues, not proven requirements.
Timing should be verified on the actual GPi Case 2 image before replacement.

## Process-Kill Assumptions

The upstream script assumes EmulationStation should be killed before Linux
shutdown. It does not appear to perform game-aware shutdown orchestration,
save-state handling, emulator-specific cleanup, user prompts, or retry logic.

RetroFlag Power must not assume that reproducing `killall` is sufficient for a
safe user-facing shutdown. Clean EmulationStation and active-game behavior
remain acceptance-checklist items.

## LCD/HDMI Switching Assumptions

The public scripts assume LCD/HDMI switching can be handled by:

- Reading GPIO18.
- Checking `enable_dpi_lcd=1` in `/boot/config.txt`.
- Copying the current `/boot/config.txt` into an LCD or HDMI backup file.
- Replacing `/boot/config.txt` with the alternate saved file.
- Rebooting from `lcdfirst.sh` when the boot-time display mode differs.

This conflicts with the current KMS direction. The GPi Case 2 LCD now works
through modern KMS DPI with `vc4-kms-dpi-generic` and `rgb666-padhi`, so any
future docking behavior must be KMS-safe and must not blindly restore legacy
FKMS or DPI config files.

## Docking Assumptions

The upstream README frames dock behavior as automatic switching between the
LCD display and HDMI output. The scripts appear to implement that through
GPIO18 and `/boot/config.txt` rewrites.

Unknowns:

- Whether the dock also changes audio behavior, power domains, or wake paths.
- Whether GPIO18 remains the correct dock signal under the current KMS setup.
- Whether a KMS-native transition is possible without reboot or config rewrite.
- Whether docking changes the power-save or RCU-stall recovery path.

## Power-Save/Resume Assumptions

Known and inferred behavior:

- The top power-save/resume button is not a proven hard power or reset button.
- Field notes show it can wake visible state from display-off/power-save.
- The case may auto-enter power-save after roughly 15-20 minutes idle, so
  avoiding the top button is not enough to avoid the behavior.
- In an RCU stall state, normal Linux userspace may not be able to service the
  side-switch shutdown request even when `SafeShutdown.py` is enabled.
- SafeShutdown.py being enabled does not guarantee recovery if Linux is
  already stalled.
- The side switch should be treated as a software-mediated shutdown request
  until proven otherwise, not as a physical power cut.

The field observation that display-off/blinking-LED behavior still appeared
possible during the stalled state is evidence that some behavior may live
outside normal userspace. It is not proof of a particular MCU, GPIO, regulator,
or schematic design.

## Safe Shutdown Path

The upstream safe shutdown path appears to be:

1. Initialize GPIO.
2. Drive GPIO27 HIGH.
3. Start a `poweroff()` worker.
4. Wait for GPIO26 falling edge.
5. Kill EmulationStation processes.
6. Sleep about five seconds.
7. Execute `sudo shutdown -h now`.

This path depends on Linux userspace being alive enough for Python, GPIO edge
handling, shell commands, process signaling, and shutdown execution.

## Power-Enable/Latch Path

The upstream script sets GPIO27 as output and drives it HIGH at startup. The
existing EDC treats this as likely power-enable or latch-hold behavior because
the GPi Case 2 side switch does not appear to directly cut battery power.

Attribution remains incomplete:

- Likely script-owned: asserting GPIO27 HIGH while the script is running.
- Likely board-owned: the actual power latch, regulator, MCU, or circuit that
  reacts to GPIO27.
- Unknown: what happens if the script never starts, exits, crashes, or loses
  scheduler time during a kernel stall.

## Behaviors Likely Owned By Linux Script

- GPIO26 edge wait for side-switch shutdown request.
- GPIO27 HIGH assertion during script initialization.
- EmulationStation process-kill commands.
- Five-second delay before shutdown.
- Linux `shutdown -h now` execution.
- Starting the LCD polling worker.
- Repeated calls to `/opt/RetroFlag/lcdnext.sh`.
- Boot-time startup through `/etc/rc.local`, when installed by the upstream
  installer.

## Behaviors Likely Owned By GPi Case 2 Board/MCU

These are inferences, not proven schematic facts:

- The physical power latch or hold circuit that makes GPIO27 matter.
- Some part of the top-button power-save/resume behavior.
- Some part of the display-off and blinking-indicator state.
- Some automatic idle power-save behavior after roughly 15-20 minutes.
- Any behavior that continued during the RCU stall state when network and
  side-switch shutdown recovery were unavailable.

Do not treat this section as proof of an MCU or specific circuit. It only marks
where Linux userspace ownership appears insufficient to explain the field
behavior.

## Behaviors Not Yet Attributable

- Exact top power-save/resume signal path.
- Whether top-button behavior reaches Linux as GPIO, input event, script
  state, display event, or not at all.
- Exact meaning of GPIO27 on the case board.
- Whether GPIO26 is still readable or actionable during the RCU stall state.
- Whether the blinking indicator LED is controlled by Linux, firmware, board
  logic, charger logic, or another controller.
- Whether automatic idle power-save is script-timed, board-timed, display
  domain behavior, audio behavior, or a combination.
- Whether docking changes power-save timing or shutdown reliability.
- Whether local `/opt/RetroFlag` files differ from upstream scripts.

## Risks During Kernel Stall/RCU Stall

During the 2026-07-08 field incident, the device showed repeated Linux
`rcu: INFO: rcu_preempt detected stalls on CPUs/tasks` messages, no SSH, no
ping, no side-switch shutdown response, and no normal software recovery path.
Physical CM4 cartridge/card removal was the only observed stop.

Risks:

- `SafeShutdown.py` may be installed and still fail to recover the device if
  Linux userspace or scheduling is already wedged.
- GPIO26 edge handling may not run.
- Shell commands may not execute.
- `shutdown -h now` may not run.
- SSH and ping may not be available as recovery channels.
- Display-off or blinking LED state can hide that Linux remains stalled.
- Repeated physical removal risks filesystem and hardware stress.

## Behaviors RetroFlag Power Must Preserve Before Replacement

Before any replacement quest, RetroFlag Power must preserve, replace with
evidence, or explicitly retire:

- GPIO27 power-enable/latch behavior.
- GPIO26 side-switch shutdown request handling.
- Clean shutdown sequencing for EmulationStation and Linux.
- Timing expectations around side-switch action, process cleanup, and shutdown.
- KMS-safe LCD behavior.
- Docked HDMI behavior.
- LCD-to-HDMI and HDMI-to-LCD behavior.
- Top-button power-save/resume behavior or a documented coexistence plan.
- Automatic idle power-save behavior after roughly 15-20 minutes.
- Recovery planning for RCU stall/no-network/no-side-switch-response states.
- Local script migration, backup, restore, and rollback behavior.

Until these are proven, the stock script remains part of the active GPi Case 2
power path.

## Open Questions For Field Testing

- What are the exact local contents and timestamps under `/opt/RetroFlag`?
- Does the local `SafeShutdown.py` match upstream
  `retropie_SafeShutdown_gpi2.py`?
- What exact command starts the local script from `/etc/rc.local`?
- Is GPIO26 observable as `SignalLow`, `SignalHigh`, or `SignalUnverified`
  across side-switch positions while preserving vocabulary boundaries?
- What is GPIO27's observed state while the stock script is running, sleeping,
  shutting down, power-saving, or stalled?
- Does the top button produce any Linux input event or readable GPIO change?
- Does automatic idle power-save occur consistently after roughly 15-20
  minutes across handheld and docked modes?
- Does the blinking indicator LED continue when Linux userspace is stalled?
- Does docking alter GPIO18, audio, display, power-save timing, or recovery?
- Can a KMS-safe docking path be proven without legacy `/boot/config.txt`
  rewrites?
- What recovery path exists when Linux stalls and software shutdown is gone?

The map stays humble until the field evidence earns a brighter badge.
