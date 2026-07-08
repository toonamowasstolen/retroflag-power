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

This document is documentation only. It does not change Go code, add
executable project tooling, read GPIO, write GPIO, execute shutdown, execute
reboot, install or activate systemd, alter `rc.local`, replace
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

The sample script creates a timestamped folder like:

```text
/home/pi/gpi-case2-boot-power-trace-20260708-191500
```

It writes:

- `trace.csv`, one sample per second for 90 seconds.
- `report.txt`, capture context, command availability, and safety reminders.
- `dmesg-power-display-usb.txt`, matching kernel lines seen during the trace.
- `process-milestones.txt`, first-seen milestones such as EmulationStation
  becoming visible to `pgrep`.
- A `.tar.gz` bundle beside the folder.

Each CSV row includes:

- Capture timestamp.
- Uptime from `/proc/uptime`.
- `vcgencmd get_throttled`.
- `vcgencmd measure_volts`.
- `vcgencmd measure_temp`.
- Whether `emulationstation` is running.
- A compact latest matching `dmesg` hint.

If a command is unavailable, the script records that as evidence instead of
installing anything.

## What It Does Not Do

The sample script does not:

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

## Manual Capture On The GPi Case 2

Paste this into an SSH session or local terminal shortly after boot. Keep the
handheld awake while it runs, and stop the session before idle auto power-save
can trigger.

This is docs-only sample text, not project tooling.

```sh
sh <<'EOF'
#!/bin/sh
set -eu

DURATION_SECONDS="90"
STAMP="$(date +%Y%m%d-%H%M%S)"
ROOT="${HOME}/gpi-case2-boot-power-trace-${STAMP}"
BUNDLE="${ROOT}.tar.gz"
DMESG_PATTERN="under-voltage|undervoltage|voltage|thrott|vc4|v3d|drm|dpi|kms|framebuffer|panel|mailbox|usb|hid|input|audio|snd|mmc|ext4|filesystem|I/O error|rcu|stall|hung|blocked"

mkdir -p "${ROOT}"

csv_escape() {
  printf '%s' "$1" | sed 's/"/""/g; s/^/"/; s/$/"/'
}

command_value() {
  if command -v "$1" >/dev/null 2>&1; then
    "$@" 2>&1 || true
  else
    printf 'command-unavailable:%s' "$1"
  fi
}

latest_dmesg_match() {
  dmesg 2>&1 | grep -Ei "${DMESG_PATTERN}" | tail -n 1 || true
}

{
  echo "GPi Case 2 Boot Power Trace Lantern"
  echo "Captured local time: ${STAMP}"
  echo "Duration seconds: ${DURATION_SECONDS}"
  echo "Output folder: ${ROOT}"
  echo
  echo "This trace is local and read-only."
  echo "It does not read GPIO, write GPIO, upload, install, fix, shutdown, or reboot."
  echo "Keep the device active and stop before idle auto power-save."
  echo "Do not test resume yet."
  echo
  echo "Command availability:"
  for cmd in date dmesg grep tail sed awk pgrep tar vcgencmd; do
    if command -v "${cmd}" >/dev/null 2>&1; then
      echo "- ${cmd}: available"
    else
      echo "- ${cmd}: unavailable"
    fi
  done
} > "${ROOT}/report.txt"

printf '%s\n' \
  'captured_at,uptime_seconds,throttled,volts,temp,emulationstation_running,latest_dmesg_match' \
  > "${ROOT}/trace.csv"

{
  echo "GPi Case 2 Boot Power Trace process milestones"
  echo "Captured local time: ${STAMP}"
  echo
} > "${ROOT}/process-milestones.txt"

{
  echo "GPi Case 2 Boot Power Trace matching dmesg lines"
  echo "Captured local time: ${STAMP}"
  echo "Pattern: ${DMESG_PATTERN}"
  echo
} > "${ROOT}/dmesg-power-display-usb.txt"

ES_WAS_RUNNING=0
i=0
while [ "${i}" -lt "${DURATION_SECONDS}" ]; do
  NOW="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
  UPTIME="$(awk '{print $1}' /proc/uptime 2>/dev/null || printf 'unknown')"
  THROTTLED="$(command_value vcgencmd get_throttled | tr '\n' ' ' | sed 's/[[:space:]]*$//')"
  VOLTS="$(command_value vcgencmd measure_volts | tr '\n' ' ' | sed 's/[[:space:]]*$//')"
  TEMP="$(command_value vcgencmd measure_temp | tr '\n' ' ' | sed 's/[[:space:]]*$//')"

  if pgrep -x emulationstation >/dev/null 2>&1; then
    ES_RUNNING="yes"
    if [ "${ES_WAS_RUNNING}" -eq 0 ]; then
      echo "${NOW} uptime=${UPTIME} emulationstation first observed running" \
        >> "${ROOT}/process-milestones.txt"
      ES_WAS_RUNNING=1
    fi
  else
    ES_RUNNING="no"
  fi

  LATEST_DMESG="$(latest_dmesg_match | tr '\n' ' ' | sed 's/[[:space:]]*$//')"

  {
    csv_escape "${NOW}"; printf ','
    csv_escape "${UPTIME}"; printf ','
    csv_escape "${THROTTLED}"; printf ','
    csv_escape "${VOLTS}"; printf ','
    csv_escape "${TEMP}"; printf ','
    csv_escape "${ES_RUNNING}"; printf ','
    csv_escape "${LATEST_DMESG}"; printf '\n'
  } >> "${ROOT}/trace.csv"

  {
    echo "----- ${NOW} uptime=${UPTIME} -----"
    dmesg 2>&1 | grep -Ei "${DMESG_PATTERN}" | tail -n 20 || true
    echo
  } >> "${ROOT}/dmesg-power-display-usb.txt"

  i=$((i + 1))
  sleep 1
done

{
  echo
  echo "Capture completed: $(date -u +%Y-%m-%dT%H:%M:%SZ)"
  echo "Final throttled: $(command_value vcgencmd get_throttled | tr '\n' ' ' | sed 's/[[:space:]]*$//')"
  echo "Final volts: $(command_value vcgencmd measure_volts | tr '\n' ' ' | sed 's/[[:space:]]*$//')"
  echo "Final temp: $(command_value vcgencmd measure_temp | tr '\n' ' ' | sed 's/[[:space:]]*$//')"
  echo
  echo "Interpret with the EDC procedure:"
  echo "- early boot"
  echo "- KMS/display initialization"
  echo "- USB/audio/controller initialization"
  echo "- EmulationStation startup"
  echo "- idle risk"
} >> "${ROOT}/report.txt"

tar -czf "${BUNDLE}" -C "${HOME}" "$(basename "${ROOT}")"
echo "${BUNDLE}"
EOF
```

The final line prints the local bundle path, for example:

```text
/home/pi/gpi-case2-boot-power-trace-20260708-191500.tar.gz
```

## Pull The Output With `scp`

From the Mac, pull the bundle with `scp`. Replace `pi`, `retropie.local`, and
the timestamp with the values for the test device.

```sh
mkdir -p ~/Desktop/gpi-case-2-boot-power-traces
scp pi@retropie.local:/home/pi/gpi-case2-boot-power-trace-20260708-191500.tar.gz \
  ~/Desktop/gpi-case-2-boot-power-traces/
```

If `.local` name resolution is unreliable, use the device IP address:

```sh
scp pi@192.168.1.50:/home/pi/gpi-case2-boot-power-trace-20260708-191500.tar.gz \
  ~/Desktop/gpi-case-2-boot-power-traces/
```

Inspect before sharing:

```sh
mkdir -p ~/Desktop/gpi-case-2-boot-power-traces/review
tar -xzf ~/Desktop/gpi-case-2-boot-power-traces/gpi-case2-boot-power-trace-20260708-191500.tar.gz \
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
