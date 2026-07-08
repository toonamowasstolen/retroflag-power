---
id: OPS-GPI-CASE-2-BOOT-POWER-TRACE-CAPTURE-PROCEDURE-001
title: GPi Case 2 Boot Power Trace Capture Procedure
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define a manual read-only GPi Case 2 Boot Power Trace Lantern capture procedure for the first boot window.
related:
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - ../../scripts/gpi-case2-boot-power-trace-field-lantern.sh
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
  - ../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
last_updated: 2026-07-08
---

# GPi Case 2 Boot Power Trace Capture Procedure

> A Boot Power Trace Lantern watches the first boot window with a pocket watch:
> power flags, display clues, USB chatter, and process milestones, all local
> and quiet.

This procedure uses one portable read-only script:
[`scripts/gpi-case2-boot-power-trace-field-lantern.sh`](../../scripts/gpi-case2-boot-power-trace-field-lantern.sh).
The script is copied to the GPi Case 2 by hand and run from the user's home
directory. It does not require git, a repository checkout, Go, a project
install, root-only writes, or Arcadia Runtime services on the Pi.

The script does not read GPIO, write GPIO, execute shutdown, execute reboot,
install or activate systemd, alter `rc.local`, replace
`/opt/RetroFlag/SafeShutdown.py`, implement resume, flash firmware, run
RetroFlag installers, submit telemetry, upload data, apply automatic fixes, or
contact the network.

Boot Power Trace Lantern means this focused manual startup timing capture.
Field Lantern means the broader read-only local capture bundle. Common
Problems Mage means a future classifier or troubleshooter. Lantern Dispatch
means a future optional support, update, or submission layer. Lantern Dispatch
is not implemented.

## Purpose

Use this procedure to capture about 90 seconds of GPi Case 2 boot-time power,
thermal, display, USB, audio, controller, and EmulationStation clues.

The goal is to answer a narrow timing question: did undervoltage or throttling
evidence appear during early boot, KMS/display initialization,
USB/audio/controller initialization, EmulationStation startup, or later idle
risk?

The trace preserves evidence. It does not diagnose the whole power path, claim
that KMS is the cause, repair the device, or prove the hardware is safe.

## When To Run

Run this after a clean manual boot when the GPi Case 2 is responsive enough for
SSH or a local terminal.

It is useful when:

- `vcgencmd get_throttled` or kernel logs show undervoltage or throttling
  somewhere during boot.
- The maintainer needs a short timing trail before changing any power,
  display, audio, or startup behavior.
- The next step is evidence collection for the
  [GPi Case 2 Boot Power Trace Lantern Map](gpi-case-2-boot-power-trace-lantern-map.md),
  not repair.
- A later
  [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
  bundle needs an optional startup trace attached by hand.

Do not run it during a wedged RCU stall, lost-network incident, or active
recovery. Follow the
[GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
first, then capture after the device is booted again.

## Safety Notes

- Keep the GPi Case 2 active and stop before idle auto power-save can trigger.
  Current field evidence suggests idle display/audio power-save can appear
  after roughly 15-20 minutes of no input.
- Do not test resume yet.
- Do not use this trace as a power-save or wake test.
- Do not rely on the side switch as emergency recovery during a kernel stall.
- Do not change KMS, display, audio, boot, runtime, `rc.local`, or
  `SafeShutdown.py` configuration during this capture pass.
- If the trace shows RCU stalls, MMC/ext4 warnings, repeated undervoltage, or
  lost recovery clues, record that in the relevant EDC ledger before deeper
  testing.

## What It Captures

The portable Field Lantern script creates a timestamped folder like:

```text
/home/pi/gpi-case2-boot-power-trace-field-lantern-20260708-191500
```

It writes:

- `trace.csv`, one sample per second for 90 seconds.
- `report.txt`, capture context, command availability, and safety reminders.
- `manifest.txt`, portability and privacy notes for the capture.
- `dmesg-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt`, matching kernel
  lines for voltage, throttle, display, USB, controller, RCU, watchdog, MMC,
  and ext4 clues.
- `journal-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt`, matching
  `journalctl -b` lines when `journalctl` is available.
- `process-milestones.txt`, first-seen milestones such as EmulationStation
  becoming visible to `pgrep`.
- Command outputs for `uname`, `uptime`, `mount`, `df`, `free`, `lsusb`,
  `systemd-analyze blame`, `systemd-analyze critical-chain`, and safe
  `vcgencmd` readings when those commands are available.
- Readable boot context files such as `/proc/cmdline`, `/proc/uptime`,
  `/proc/device-tree/model`, `/boot/config.txt`, `/boot/cmdline.txt`,
  `/boot/firmware/config.txt`, and `/boot/firmware/cmdline.txt`.
- A `.tar.gz` bundle beside the folder.

Each CSV row includes:

- Capture timestamp.
- Uptime from `/proc/uptime`.
- `vcgencmd get_throttled`.
- `vcgencmd measure_volts`.
- `vcgencmd measure_temp`.
- Whether `emulationstation` is running.
- A compact latest matching `dmesg` hint.

If a command or file is unavailable, the script records that as evidence
instead of installing anything or broadening the capture.

## What It Does Not Do

The portable Field Lantern script does not:

- Run `curl` or `wget`.
- Modify files outside its own timestamped folder, except for the final
  `.tar.gz` bundle created beside it.
- Read GPIO or write GPIO.
- Execute shutdown, reboot, halt, poweroff, suspend, or resume.
- Start, stop, enable, disable, or install systemd services.
- Modify `/etc/rc.local`.
- Modify or replace `/opt/RetroFlag/SafeShutdown.py`.
- Run `lcdfirst.sh`, `lcdnext.sh`, or any display-switching script.
- Rewrite `/boot/config.txt`, `/boot/cmdline.txt`, KMS, display, audio, boot,
  or runtime configuration.
- Flash firmware.
- Upload telemetry or contact Lantern Dispatch.
- Apply fixes or generate a repair plan.

## Copy The Lantern To The GPi Case 2

From the Mac, copy the single script to the Pi. Replace `pi` and
`retropie.local` with the user and host for the test device.

```sh
scp scripts/gpi-case2-boot-power-trace-field-lantern.sh \
  pi@retropie.local:/home/pi/
```

If `.local` name resolution is unreliable, use the device IP address:

```sh
scp scripts/gpi-case2-boot-power-trace-field-lantern.sh \
  pi@192.168.1.50:/home/pi/
```

This copy step is the whole Field Lantern Relic handoff. The Pi does not need
the repository, git, Go, a service, or an installer.

## Run The Lantern On The GPi Case 2

SSH to the Pi shortly after a clean boot, keep the handheld awake, and run the
script from the home directory. The default trace is 90 seconds.

```sh
ssh pi@retropie.local
sh /home/pi/gpi-case2-boot-power-trace-field-lantern.sh
```

The script prints the local bundle path as its final line, for example:

```text
/home/pi/gpi-case2-boot-power-trace-field-lantern-20260708-191500.tar.gz
```

If a shorter smoke test is needed, pass a duration in seconds:

```sh
sh /home/pi/gpi-case2-boot-power-trace-field-lantern.sh 15
```

## Pull The Output With `scp`

From the Mac, pull the bundle with `scp`. Replace `pi`, `retropie.local`, and
the timestamp with the values for the test device.

```sh
mkdir -p ~/Desktop/gpi-case-2-boot-power-traces
scp pi@retropie.local:/home/pi/gpi-case2-boot-power-trace-field-lantern-20260708-191500.tar.gz \
  ~/Desktop/gpi-case-2-boot-power-traces/
```

If `.local` name resolution is unreliable, use the device IP address:

```sh
scp pi@192.168.1.50:/home/pi/gpi-case2-boot-power-trace-field-lantern-20260708-191500.tar.gz \
  ~/Desktop/gpi-case-2-boot-power-traces/
```

Inspect before sharing:

```sh
mkdir -p ~/Desktop/gpi-case-2-boot-power-traces/review
tar -xzf ~/Desktop/gpi-case-2-boot-power-traces/gpi-case2-boot-power-trace-field-lantern-20260708-191500.tar.gz \
  -C ~/Desktop/gpi-case-2-boot-power-traces/review
```

## Include In A Future Field Lantern Bundle

Today, the Boot Power Trace bundle is pulled and reviewed separately.

To include it with a later Field Lantern bundle:

1. Run this Boot Power Trace capture shortly after boot.
2. Run the
   [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
   after the device is stable and before any repair attempt.
3. Pull both `.tar.gz` files to the Mac.
4. Unpack and inspect both locally.
5. Attach the Boot Power Trace folder or `.tar.gz` beside the Field Lantern
   bundle only after redaction review.

Do not edit the Field Lantern script to run this trace automatically yet. That
belongs to a later quest with explicit scope and validation.

## Interpretation Notes

Use the trace as timing evidence, not as a single-root-cause verdict.

Early boot:

- A non-zero throttling value or voltage line already present in the first few
  samples suggests the sag happened before the trace had a cleaner subsystem
  milestone.
- The trace may still miss events that happened before userspace became
  available.

KMS/display init:

- Lines containing `vc4`, `v3d`, `drm`, `dpi`, `kms`, `framebuffer`, `panel`,
  or `mailbox` near the first voltage warning may point toward display
  initialization timing.
- This does not prove KMS is the root cause. It only places the warning near a
  display milestone.

USB/audio/controller init:

- Lines containing `usb`, `hid`, `input`, `audio`, or `snd` near voltage or
  throttle changes may point toward internal hub, controller, or USB audio
  load.
- USB and display initialization can overlap, so keep both buckets visible.

EmulationStation startup:

- The first `emulationstation` process milestone helps separate boot services
  from launcher/runtime startup.
- If `get_throttled` changes or warnings cluster after this milestone, treat
  runtime startup load as a candidate timing bucket.

Idle risk:

- This 90-second capture is meant to finish before idle auto power-save.
- Do not extend the test into idle power-save or resume yet.
- If warnings appear only after a later idle period, use the
  [GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
  boundaries before planning another capture.

The strongest verified win is a small one: a timestamped local trace that can
be inspected without touching GPIO, shutdown, installers, or the power path.
