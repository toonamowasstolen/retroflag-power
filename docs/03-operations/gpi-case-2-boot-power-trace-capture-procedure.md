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
purpose: Define the current manual read-only GPi Case 2 bundle collector procedure and separate it from the future Boot Power Trace Lantern.
related:
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - human-facing-field-lantern-script-ux-standard.md
  - ../../scripts/gpi-case2-bundle-collector-field-lantern.sh
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
  - ../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
last_updated: 2026-07-08
---

# GPi Case 2 Bundle Collector Lantern Capture Procedure

> A Bundle Collector Lantern gathers remembered boot clues after the handheld
> has reached a responsive campfire. It is a satchel, not a time machine.

This procedure uses one portable read-only script:
[`scripts/gpi-case2-bundle-collector-field-lantern.sh`](../../scripts/gpi-case2-bundle-collector-field-lantern.sh).
The script is copied to the GPi Case 2 by hand and run from the user's home
directory. It does not require git, a repository checkout, Go, a project
install, root-only writes, or Arcadia Runtime services on the Pi.

The script does not read GPIO, write GPIO, execute shutdown, execute reboot,
install or activate systemd, alter `rc.local`, replace
`/opt/RetroFlag/SafeShutdown.py`, implement resume, flash firmware, run
RetroFlag installers, submit telemetry, upload data, apply automatic fixes, or
contact the network.

Human-facing terminal UX for this and future manual Lantern scripts should
follow the
[Human-Facing Field Lantern Script UX Standard](human-facing-field-lantern-script-ux-standard.md):
visible startup, double-bracket stage lines, progress during long sampling,
exact artifact paths, duration reporting, and plain safety messages.

Bundle Collector Lantern means this manual post-boot evidence and `.tar.gz`
collector. Boot Power Trace Lantern is reserved for a future safe recorder
that starts during boot and samples timestamped state from early startup.
Session Watch Lantern means a later runtime observation lantern. Field Lantern
means the broader family of local read-only evidence procedures. Common
Problems Mage means a future classifier or troubleshooter. Lantern Dispatch
means a future optional support, update, or submission layer. Lantern Dispatch
is not implemented.

## Purpose

Use this procedure to collect a post-boot evidence bundle and, when requested,
sample the current power/throttle state for about 90 seconds.

The goal is to gather remembered boot logs, current `vcgencmd get_throttled`
state, thermal clues, display clues, USB/audio/controller clues, and process
state into one local bundle. It can show that firmware or kernel evidence was
present after boot. It cannot prove the exact second of early boot
undervoltage unless a Boot Power Trace Lantern was already running during
boot.

The bundle preserves evidence. It does not diagnose the whole power path,
claim that KMS is the cause, repair the device, or prove the hardware is safe.

## When To Run

Run this after a clean manual boot when the GPi Case 2 is responsive enough for
a local terminal or optional SSH support. The normal handheld path assumes no
attached keyboard and no repository checkout on the device; copy the one Relic
script onto the handheld only when a support path is available.

It is useful when:

- `vcgencmd get_throttled` or kernel logs show undervoltage or throttling
  somewhere during boot.
- The maintainer needs a short timing trail before changing any power,
  display, audio, or startup behavior.
- The next step is post-boot evidence collection for the
  [GPi Case 2 Boot Power Trace Lantern Map](gpi-case-2-boot-power-trace-lantern-map.md),
  not repair.
- A later
  [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
  bundle needs an optional collector bundle attached by hand.

Do not run it during a wedged RCU stall, lost-network incident, or active
recovery. Follow the
[GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
first, then capture after the device is booted again.

## Safety Notes

- Keep the GPi Case 2 active and stop before idle auto power-save can trigger.
  Current field evidence suggests idle display/audio power-save can appear
  after roughly 15-20 minutes of no input.
- Do not test resume yet.
- Avoid the top sleep/resume button during diagnostics unless a procedure
  explicitly says otherwise.
- Do not use this trace as a power-save or wake test.
- The side power switch is the normal stock shutdown control while the system
  is responsive, but do not rely on it as emergency recovery during a kernel
  stall.
- Do not change KMS, display, audio, boot, runtime, `rc.local`, or
  `SafeShutdown.py` configuration during this capture pass.
- If the trace shows RCU stalls, MMC/ext4 warnings, repeated undervoltage, or
  lost recovery clues, record that in the relevant EDC ledger before deeper
  testing.

## What It Captures

The portable Bundle Collector Lantern script creates a timestamped folder like:

```text
/home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500
```

It writes:

- `trace.csv`, one post-boot sample per second for 90 seconds by default.
- `report.txt`, capture context, command availability, and safety reminders.
- `manifest.txt`, portability and privacy notes for the capture.
- `dmesg-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt`, matching kernel
  lines for voltage, throttle, display, USB, controller, RCU, watchdog, MMC,
  and ext4 clues.
- `journal-power-display-usb-xpad-rcu-watchdog-mmc-ext4.txt`, matching
  `journalctl -b` lines when `journalctl` is available.
- `process-milestones.txt`, first-seen milestones such as EmulationStation
  becoming visible to `pgrep`. Treat this as helpful but unreliable; the first
  field run reported EmulationStation as not running even though the operator
  observed it open.
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
- `vcgencmd measure_volts`, which reports an internal/core rail reading, not
  the GPi Case 2 5V input rail.
- `vcgencmd measure_temp`.
- Whether `emulationstation` is running.
- A compact latest matching `dmesg` hint.

If a command or file is unavailable, the script records that as evidence
instead of installing anything or broadening the capture.

What this proves:

- It can gather remembered boot logs from `dmesg`, `journalctl -b`, and
  allowlisted boot files after the handheld is responsive.
- It can sample current and sticky firmware throttling flags from
  `vcgencmd get_throttled`.
- It can show whether undervoltage or throttling evidence is visible in the
  bundle.

What this cannot prove:

- It cannot determine the exact second of early boot undervoltage unless a
  boot-time recorder was already running.
- It cannot report watts, TDP, amps, power draw, or actual 5V rail voltage.
- It cannot make process detection authoritative. EmulationStation visibility
  through `pgrep` is a clue, not a verdict.

## What It Does Not Do

The portable Bundle Collector Lantern script does not:

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

From the Mac, copy the single script to the Pi. The field target is
`retropi@gpi`.

```sh
scp scripts/gpi-case2-bundle-collector-field-lantern.sh \
  retropi@gpi:/home/retropi/
```

This copy step is the whole Field Lantern Relic handoff. The Pi does not need
the repository, git, Go, a service, or an installer.

## Run The Lantern On The GPi Case 2

If SSH is available, SSH to the Pi shortly after a clean boot, keep the
handheld awake, and run the script from the home directory. SSH is optional
support, not the primary handheld UX. The default sample window is 90 seconds.

```sh
ssh retropi@gpi
sh /home/retropi/gpi-case2-bundle-collector-field-lantern.sh
```

The script prints the local bundle path as its final line, for example:

```text
/home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz
```

If a shorter smoke test is needed, pass a duration in seconds:

```sh
sh /home/retropi/gpi-case2-bundle-collector-field-lantern.sh 15
```

## Pull The Output With `scp`

From the Mac, pull the bundle with `scp`. Replace the timestamp with the value
printed by the test device.

```sh
mkdir -p ~/Desktop/gpi-case-2-boot-power-traces
scp retropi@gpi:/home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz \
  ~/Desktop/gpi-case-2-boot-power-traces/
```

Inspect before sharing:

```sh
mkdir -p ~/Desktop/gpi-case-2-boot-power-traces/review
tar -xzf ~/Desktop/gpi-case-2-boot-power-traces/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz \
  -C ~/Desktop/gpi-case-2-boot-power-traces/review
```

## Include In A Future Field Lantern Bundle

Today, this Bundle Collector Lantern output is pulled and reviewed separately.

To include it with a later Field Lantern bundle:

1. Run this Bundle Collector Lantern shortly after boot.
2. Run the
   [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
   after the device is stable and before any repair attempt.
3. Pull both `.tar.gz` files to the Mac.
4. Unpack and inspect both locally.
5. Attach the collector folder or `.tar.gz` beside the Field Lantern
   bundle only after redaction review.

Do not edit the Field Lantern script to run this automatically at boot. That
belongs to a later quest with explicit scope and validation.

## Interpretation Notes

Use the bundle as remembered-log and current-state evidence, not as a
single-root-cause verdict.

Early boot:

- A non-zero throttling value or voltage line already present in the first
  samples only proves the evidence was already visible when the script ran.
- The bundle may miss events that happened before userspace became available
  or before the script was manually started.

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

- The first `emulationstation` process milestone may help separate boot
  services from launcher/runtime startup, but current process detection is
  known to be unreliable.
- If `get_throttled` changes or warnings cluster after this milestone, treat
  runtime startup load as a candidate timing bucket.

Idle risk:

- This 90-second capture is meant to finish before idle auto power-save.
- Do not extend the test into idle power-save or resume yet.
- If warnings appear only after a later idle period, use the
  [GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
  boundaries before planning another capture.

## Next Lantern Direction

The future Boot Power Trace Lantern should be a safe boot-time recorder:

- Read-only.
- Local file output.
- Timestamped samples from early boot.
- No GPIO.
- No shutdown or reboot.
- No systemd activation yet in this quest.

It should start early enough to make boot-time capture real instead of
post-boot recollection, but that startup wiring is deliberately not
implemented here.

## Future Session Watch Lantern

A later Session Watch Lantern can watch runtime play and menu sessions:

- `vcgencmd get_throttled` flags over time.
- Temperature, load, and memory.
- Frontend, emulator, and game detection where possible.
- Recent `dmesg` and journal warnings.
- No telemetry by default.

The strongest verified win is a small one: a timestamped local bundle that can
be inspected without touching GPIO, shutdown, installers, or the power path.
