---
id: OPS-GPI-CASE-2-RECOVERY-FIRST-FIELD-PROCEDURE-001
title: GPi Case 2 Recovery-First Field Procedure
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define a conservative recovery-first field procedure for a GPi Case 2 stranded after the power-save/resume RCU stall incident.
related:
  - gpi-case-2-replacement-coverage-matrix.md
  - gpi-case-2-safeshutdown-script-behavior-map.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - gpi-case-2-acceptance-checklist.md
  - safeshutdown-replacement-boundary-map.md
  - gpi-case-gpio-probe-ledger.md
  - ../03-hardware/gpi-case-2-emergency-recovery-research-ledger.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md
last_updated: 2026-07-08
---

# GPi Case 2 Recovery-First Field Procedure

> When the handheld goes dark and the network trail vanishes, this page keeps
> the field work calm: recover first, preserve evidence second, and do not turn
> panic into a new hardware plan.

This document is documentation only. It does not change Go code, read GPIO,
write GPIO, execute shutdown, install or activate systemd, alter `rc.local`,
replace `/opt/RetroFlag/SafeShutdown.py`, implement resume, add persistent
state, add telemetry, make project-code network calls, change installers or
packaging, run RetroFlag install scripts, or approve hardware modification.

This is not a repair guide. It is a recovery-first field procedure for the
known GPi Case 2 power-save/resume failure mode where SSH and ping are gone,
the screen or console may show Linux RCU stall messages, and the side switch
does not recover the handheld.

Any emergency reset or cutoff modification requires separate board evidence,
electrical review, and maintainer approval. Cutting battery leads, modifying
lithium battery circuitry, modifying charging circuitry, blind soldering, and
treating unknown pads as reset or power paths are not approved.

## Scope And Warning

Use this procedure only after a GPi Case 2 appears wedged after power-save,
resume, or automatic idle power-save.

The current evidence says:

- The side switch is software-mediated until proven otherwise.
- If Linux is stalled, `SafeShutdown.py` may not run even if enabled.
- The top power-save/resume button is not a proven reset or hard power button.
- Automatic display/audio power-save after roughly 15-20 minutes means avoiding
  the top button is not enough.
- Repeated CM4 cartridge/card removal and battery depletion are unacceptable as
  normal recovery plans.

Preserve the vocabulary boundary:

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted meanings from a
  configured hardware profile.
- Lanterns are read-only probes and diagnostics.
- Lantern Dispatch remains future optional support infrastructure, not
  implemented.

## Current Known Failure Mode

The 2026-07-08 field incident recorded in
[QUEST-0064](../00-project/quests/0064-record-gpi-case-2-power-save-rcu-stall-incident.md)
showed this pattern:

- Repeated Linux `rcu: INFO: rcu_preempt detected stalls on CPUs/tasks`
  messages after power-save or resume behavior.
- SSH unavailable.
- Ping unavailable.
- Side switch moved off, but the device did not shut down.
- Top power-save/resume button still changed the visible state.
- `SafeShutdown.py` was believed to be enabled.
- The only observed stop was physical CM4 cartridge/card removal.

Treat this as a critical unresolved field failure. Do not classify it as only a
disabled-stock-script problem.

## Immediate Triage While Wedged

When the handheld is already wedged:

1. Stop active testing and write down the exact time, physical mode, dock
   state, charger state, visible screen state, LED behavior, and the last user
   action before the failure.
2. Photograph the screen if RCU stall text or any kernel text is visible.
3. Try one ordinary ping and one ordinary SSH connection from the field
   workstation if those were already part of the test setup.
4. Record whether the top power-save/resume button changes only the visible
   state or appears to restore a working session.
5. Record whether the side switch has any visible effect, without assuming
   `SwitchOn` or `SwitchOff` meaning.
6. If no safe software path returns, stop the test and move to the least-bad
   stop options below.

Do not run GPIO probes, GPIO reads, GPIO writes, shutdown commands, installers,
service activation, or script replacement while triaging this wedged state.

## What Not To Do Repeatedly

Do not normalize any of these as routine recovery:

- Repeated CM4 cartridge/card removal.
- Waiting for battery depletion.
- Repeated side-switch cycling.
- Repeated top-button sleep/wake cycling after RCU stalls appear.
- Blind dock/undock cycling after the device is unresponsive.
- Running RetroFlag install scripts.
- Replacing, editing, or disabling `/opt/RetroFlag/SafeShutdown.py`.
- Replacing or editing `/etc/rc.local`.
- Adding a hardware reset, cutoff, battery, charge, or solder path without a
  separate approved investigation.

Repeated forced stops can hide the evidence and may increase filesystem or
hardware risk.

## Least-Bad Stop Options

If SSH, ping, and the side switch all fail:

- Prefer stopping the field test over trying to prove one more behavior.
- If the display can still wake, capture a photo first.
- If a keyboard, docked input path, or already-attached local console is
  available and responsive, capture visible evidence only. Do not execute
  shutdown from this procedure.
- If no safe software path exists, record that no approved recovery path was
  available and stop the session using the least physically disruptive method
  available to the maintainer.

This procedure does not approve physical CM4 removal as a normal plan. It only
records that a previous incident had no better observed stop. A better
development recovery path must be earned through the
[emergency recovery research ledger](../03-hardware/gpi-case-2-emergency-recovery-research-ledger.md).

## Evidence To Capture After Recovery

Once the device boots again, capture evidence before starting new field tests:

- Date and local time of reboot.
- Whether the device was handheld, docked, charging, or on battery during the
  failure.
- Visible screen or LED behavior from the wedged state.
- Whether ping, SSH, side switch, top button, controller input, or docked input
  responded.
- Whether the device booted normally afterward.
- Any filesystem, mmc, ext4, undervoltage, display, or RCU messages.
- Local `/opt/RetroFlag` script provenance.
- `/etc/rc.local` launcher state.
- Kernel, OS, boot config, and display stack state.

Keep the first post-recovery capture read-only. Do not install, repair, replace
scripts, enable services, or run shutdown commands during this evidence pass.

For a bundle-shaped post-recovery capture that avoids long terminal pastes,
use the manual
[GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
after the device is booted and responsive. The Field Lantern is read-only and
local; it does not fix the device, read or write GPIO, execute shutdown, upload
data, or replace future `retroflag-powerd diagnostics`.

## Read-Only Log Searches After Recovery

Search for RCU stalls and neighboring power/display/kernel messages:

```sh
dmesg -T | grep -Ei "rcu|stall|hung|blocked|mmc|ext4|filesystem|under-voltage|voltage|vc4|drm|dpi|kms|usb|wlan"
journalctl -b -k --no-pager | grep -Ei "rcu|stall|hung|blocked|mmc|ext4|filesystem|under-voltage|voltage|vc4|drm|dpi|kms|usb|wlan"
journalctl -b -1 -k --no-pager | grep -Ei "rcu|stall|hung|blocked|mmc|ext4|filesystem|under-voltage|voltage|vc4|drm|dpi|kms|usb|wlan"
grep -RniE "rcu|stall|hung|blocked|mmc|ext4|filesystem|under-voltage|voltage|vc4|drm|dpi|kms|usb|wlan" /var/log/syslog* /var/log/kern.log* 2>/dev/null
```

If previous-boot journals are unavailable, record that absence. Do not change
journald settings as part of this procedure.

## Filesystem Health Checks After Recovery

Use read-only status checks first:

```sh
findmnt -no TARGET,SOURCE,FSTYPE,OPTIONS /
findmnt -no TARGET,SOURCE,FSTYPE,OPTIONS /boot
lsblk -f
df -hT
dmesg -T | grep -Ei "mmc|ext4|fsck|filesystem|I/O error|Buffer I/O|read-only|remount"
journalctl -b -k --no-pager | grep -Ei "mmc|ext4|fsck|filesystem|I/O error|Buffer I/O|read-only|remount"
journalctl -b -1 -k --no-pager | grep -Ei "mmc|ext4|fsck|filesystem|I/O error|Buffer I/O|read-only|remount"
```

Do not run repair modes against mounted filesystems. If a future recovery plan
needs offline filesystem checks, it needs its own procedure and approval.

## Script Provenance After Recovery

Capture the local RetroFlag script state without running installers or
replacing files:

```sh
ls -la /opt/RetroFlag
find /opt/RetroFlag -maxdepth 2 -type f -printf "%TY-%Tm-%Td %TH:%TM %s %p\n" | sort
sha256sum /opt/RetroFlag/* 2>/dev/null
grep -RniE "GPIO|powerPin|powerenPin|lcdrun|lcdnext|lcdfirst|shutdown|killall|wait_for_edge|multiprocessing" /opt/RetroFlag 2>/dev/null
```

Capture the `rc.local` launcher state:

```sh
ls -la /etc/rc.local
sed -n "1,220p" /etc/rc.local
grep -nE "RetroFlag|SafeShutdown|lcdfirst|lcdnext|python|shutdown" /etc/rc.local
```

Capture whether the stock process is running after the recovered boot:

```sh
ps aux | grep -Ei "SafeShutdown|RetroFlag|lcdfirst|lcdnext" | grep -v grep
pgrep -af "SafeShutdown|RetroFlag|lcdfirst|lcdnext"
```

These commands inspect provenance and process state only. They do not prove the
side switch is safe, do not prove GPIO27 latch behavior, and do not authorize
replacement.

## Kernel, OS, And Boot Config After Recovery

Capture kernel, OS, package, and boot configuration context:

```sh
uname -a
cat /etc/os-release
dpkg -l "raspberrypi-kernel*" "raspi-firmware*" "libraspberrypi*" 2>/dev/null
vcgencmd version
vcgencmd get_throttled
```

Capture the KMS, FKMS, display, DPI, and audio lines in `/boot/config.txt`:

```sh
grep -nEi "vc4|kms|fkms|dpi|dtoverlay|dtparam|hdmi|display|framebuffer|audio|audremap|dwc2|disable_fw_kms_setup|max_framebuffers" /boot/config.txt
```

If `vcgencmd` or package commands are unavailable, record that as evidence
rather than installing anything during the capture pass.

## Power-Save Prevention Ideas To Investigate

These ideas are not yet proven and are not approved as implementation steps:

- Determine whether automatic display/audio power-save can be configured off
  without breaking handheld LCD, docked HDMI, audio, or stock SafeShutdown
  behavior.
- Determine whether the top power-save/resume path appears as an input event,
  a GPIO signal, a display/backlight event, or board-controlled behavior.
- Determine whether a local diagnostics lantern can summarize logs, script
  provenance, KMS config, and filesystem warnings after recovery.
- Determine whether a reversible development recovery path exists without
  touching lithium battery or charging circuitry.
- Determine whether field tests should use a prepared local console, docked
  keyboard, or other non-network observation path before triggering
  power-save.

None of these ideas permit GPIO writes, GPIO reads, shutdown execution,
systemd activation, `rc.local` edits, stock script replacement, installer work,
or hardware modification in this procedure.

## Pause Criteria For Further Field Tests

Pause further power-save, resume, dock-transition, or replacement-runtime
field tests when any of these occur:

- Any RCU stall, hung task, repeated kernel warning, mmc, ext4, filesystem, or
  I/O error appears.
- SSH and ping are both lost after power-save or resume.
- The side switch fails to recover or shut down the device.
- The top button changes visible state but does not restore a working session.
- Recovery depends on battery depletion or CM4 cartridge/card removal.
- The same failure occurs twice without new evidence or a safer recovery plan.
- A proposed next step requires hardware modification, GPIO writes, shutdown
  execution, service activation, script replacement, or installer changes.

When paused, update the relevant ledger before resuming:

- [GPi Case 2 Replacement Coverage Matrix](gpi-case-2-replacement-coverage-matrix.md)
- [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
- [GPi Case 2 Emergency Recovery Research Ledger](../03-hardware/gpi-case-2-emergency-recovery-research-ledger.md)
- [GPi Case 2 SafeShutdown Script Behavior Map](gpi-case-2-safeshutdown-script-behavior-map.md)
- [GPi Case 2 Acceptance Checklist](gpi-case-2-acceptance-checklist.md)

The field win here is not forcing the device back into testing. The win is
keeping enough evidence alive that the next quest can move with a steadier
lantern.
