---
id: OPS-GPI-CASE-2-FIELD-LANTERN-CAPTURE-PROCEDURE-001
title: GPi Case 2 Field Lantern Capture Procedure
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define a manual read-only GPi Case 2 Field Lantern capture bundle procedure for troubleshooting evidence without long terminal pastes.
related:
  - gpi-case-2-session-watch-lantern-design.md
  - gpi-case-2-boot-power-trace-capture-procedure.md
  - human-facing-field-lantern-script-ux-standard.md
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-recovery-first-field-procedure.md
  - gpi-case-2-safeshutdown-script-behavior-map.md
  - safeshutdown-replacement-boundary-map.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
  - ../03-hardware/gpi-case-2-developer-access-paths.md
last_updated: 2026-07-08
---

# GPi Case 2 Field Lantern Capture Procedure

> A Field Lantern gathers the room before anyone reaches for a fix: local,
> read-only, inspectable, and quiet enough to leave every GPIO and power path
> untouched.

This document is documentation only. It does not change Go code, add project
tooling, implement diagnostics, generate a bundle from `retroflag-powerd`,
read GPIO, write GPIO, execute shutdown, install or activate systemd, alter
`rc.local`, replace `/opt/RetroFlag/SafeShutdown.py`, implement resume, flash
firmware, submit telemetry, upload data, make project-code network calls,
apply fixes, run RetroFlag installers, or approve hardware modification.

Field Lantern means this manual read-only local capture procedure. Common
Problems Mage means a future classifier or troubleshooter. Lantern Dispatch
means a future optional support, update, or submission layer. Lantern Dispatch
is not implemented.

The future focused startup timing trail is mapped in
[GPi Case 2 Boot Power Trace Lantern Map](gpi-case-2-boot-power-trace-lantern-map.md).
The current manual post-boot bundle collector procedure lives in
[GPi Case 2 Boot Power Trace Capture Procedure](gpi-case-2-boot-power-trace-capture-procedure.md).
That current Bundle Collector Lantern gathers remembered boot logs and samples
current `vcgencmd get_throttled` state after the handheld is responsive. The
script prints a human-facing banner, double-bracket stage lines, sampling
progress, timing notes, and exact artifact/retrieval lines. Use `--duration`
for shorter field retests and `--plain` when copyable ASCII output is better.
The
Boot Power Trace Lantern name is reserved for a future safe boot-time recorder
that starts during boot and samples timestamped state. It may later become one
optional section inside a Field Lantern bundle, but this Field Lantern
procedure does not run it automatically.

The future runtime session trail is mapped in
[GPi Case 2 Session Watch Lantern Design](gpi-case-2-session-watch-lantern-design.md).
That Session Watch Lantern is designed to observe menu, emulator, play,
idle-risk, and post-resume sessions after boot. It is not implemented here and
does not run automatically.

Manual Field Lantern scripts that a person watches in a terminal should use
the
[Human-Facing Field Lantern Script UX Standard](human-facing-field-lantern-script-ux-standard.md)
for warm startup banners, double-bracket stage labels, long-running progress,
artifact paths, timing, `--plain`, and unmistakable safety messages.

## Purpose

Use this procedure to gather a GPi Case 2 troubleshooting bundle after a field
incident or before asking another maintainer to inspect a device state.

The bundle turns many small evidence commands into one local `.tar.gz` file so
the maintainer does not need to paste long terminal output into chat. The
bundle is meant to be pulled to a workstation, inspected, redacted if needed,
and shared manually only when the user chooses.

## When To Use It

Use the Field Lantern when:

- The GPi Case 2 recovered from a power-save, resume, display, dock, audio, or
  SafeShutdown-related incident.
- A maintainer needs local script provenance, boot config, KMS evidence, power
  warnings, input identity, USB identity, audio identity, and process state in
  one place.
- The device is booted and responsive over a local terminal or optional SSH
  support.
- The next step is evidence collection, not repair.

For a wedged device with no SSH, no ping, or visible RCU stall output, follow
[GPi Case 2 Recovery-First Field Procedure](gpi-case-2-recovery-first-field-procedure.md)
first. Capture screen photos and recovery notes before running any post-boot
bundle.

## When Not To Use It

Do not use this procedure as a repair step or live recovery step.

Do not start a new capture while:

- The device is currently wedged, repeatedly printing kernel stalls, or losing
  both SSH and ping.
- A field test would require shutdown, reboot, suspend, resume, service
  activation, GPIO probing, or script replacement.
- The proposed next step is to run RetroFlag installers, edit `/boot/config.txt`,
  edit `/etc/rc.local`, disable `SafeShutdown.py`, or change power behavior.
- The device needs emergency physical recovery or hardware investigation.

Pause and update the relevant EDC ledger if capture evidence shows RCU stalls,
filesystem or MMC warnings, undervoltage, lost recovery paths, or an unexpected
SafeShutdown/script state.

## Safety Boundaries

The Field Lantern must remain read-only and local.

It must not:

- Fix anything.
- Write GPIO.
- Read GPIO.
- Trigger shutdown, reboot, halt, poweroff, suspend, or resume.
- Install RetroFlag scripts.
- Run RetroFlag installers.
- Start, stop, enable, disable, or install systemd services.
- Modify `/etc/rc.local`.
- Modify or replace `/opt/RetroFlag/SafeShutdown.py`.
- Run `lcdfirst.sh`, `lcdnext.sh`, or any display-switching script.
- Rewrite `/boot/config.txt`, `/boot/cmdline.txt`, KMS, FKMS, display, audio,
  boot, or runtime configuration.
- Flash firmware.
- Contact Lantern Dispatch.
- Upload anything automatically.
- Replace future `retroflag-powerd diagnostics`.
- Treat a capture as proof that hardware is safe.

The side power switch remains the normal stock shutdown control while the
system is responsive. The top sleep/resume button remains suspect and should
be avoided during diagnostics unless a procedure explicitly says otherwise.

It also must not collect or infer raw GPIO observations. Preserve the project
vocabulary boundary:

- `SignalLow`, `SignalHigh`, and `SignalUnverified` are raw signal
  observations.
- `SwitchOn`, `SwitchOff`, and `SwitchUnknown` are interpreted switch meanings.
- This procedure does not create either kind of observation.

## Bundle Contents

The current Field Lantern bundle should contain:

- `report.txt`, a human-readable index and context report.
- Local `/opt/RetroFlag` copies.
- Local `/boot/config.txt`.
- Local `/boot/cmdline.txt`.
- Local `/etc/rc.local`.
- RetroFlag script hashes.
- Upstream RetroFlag files downloaded into the bundle for comparison only.
- Local-vs-upstream diffs.
- Kernel RCU and stall excerpts.
- Voltage and throttling excerpts.
- Display, KMS, and DRM excerpts.
- Filesystem and MMC warning excerpts.
- Input device identity.
- USB device identity.
- Audio device identity.
- Relevant service and process state.

`vcgencmd get_throttled` is the main firmware clue for undervoltage and
throttling, but it does not report watts, TDP, amps, power draw, or actual 5V
rail voltage. `vcgencmd measure_volts`, when captured by a Lantern, reports an
internal/core rail reading, not the GPi Case 2 5V input rail.

Current process detection is useful but not authoritative. The first field
run reported EmulationStation as not running even though the operator observed
it open, so process state should remain a clue rather than a verdict.

The bundle should prefer narrow excerpts and copied allowlisted files over
broad system dumps.

## Bundle Non-Goals

The bundle must not include:

- ROM names, save data, game libraries, screenshots, emulator content, or user
  media.
- Wi-Fi SSIDs, tokens, secrets, private keys, or full environment dumps.
- Usernames, home paths, hostnames, private IP addresses, or network identity
  unless the user deliberately keeps them after review.
- Arbitrary journal output outside the narrow allowlisted searches.
- GPIO reads or GPIO probe output collected by this procedure.
- Any generated fix plan that claims replacement readiness.

## Privacy And Redaction Notes

Before sharing a bundle, unpack it locally and inspect it.

Recommended review points:

- Search for usernames, home paths, hostnames, private IP addresses, Wi-Fi
  names, tokens, or secrets.
- Review `report.txt` and copied config files before sending them anywhere.
- Remove any optional section that is not needed for the support question.
- Share the `.tar.gz` manually only after review.

The safest bundle is useful even after redaction. Unknowns should remain
`Unknown` instead of being replaced with confident guesses.

## Manual Capture On The GPi Case 2

Run this from a local terminal or optional SSH session on the GPi Case 2 after
the device has recovered and is stable enough for read-only inspection. The
handheld path must assume no attached keyboard and no repository checkout on
the device; SSH is support convenience, not the primary UX.

This is docs-only sample text, not project tooling. Keep the script in a
temporary home-directory file on the device and delete it later if desired.

```sh
cat > ~/field-lantern-capture.sh <<'EOF'
#!/bin/sh
set -eu

STAMP="$(date -u +%Y%m%dT%H%M%SZ)"
ROOT="${HOME}/field-lantern-${STAMP}"
BUNDLE="${ROOT}.tar.gz"

mkdir -p "${ROOT}/local/opt" "${ROOT}/local/boot" "${ROOT}/local/etc"
mkdir -p "${ROOT}/upstream" "${ROOT}/diffs" "${ROOT}/logs"
mkdir -p "${ROOT}/identity" "${ROOT}/state"

{
  echo "GPi Case 2 Field Lantern Capture"
  echo "Captured UTC: ${STAMP}"
  echo
  echo "This bundle is read-only evidence. It does not fix, upload, install,"
  echo "flash, read GPIO, write GPIO, or trigger shutdown."
  echo
  echo "Review and redact before sharing."
} > "${ROOT}/report.txt"

uname -a > "${ROOT}/identity/uname.txt" 2>&1 || true
cat /etc/os-release > "${ROOT}/identity/os-release.txt" 2>&1 || true
vcgencmd version > "${ROOT}/identity/vcgencmd-version.txt" 2>&1 || true
vcgencmd get_throttled > "${ROOT}/identity/vcgencmd-get-throttled.txt" 2>&1 || true

if [ -d /opt/RetroFlag ]; then
  cp -a /opt/RetroFlag "${ROOT}/local/opt/RetroFlag" 2>/dev/null || true
  find /opt/RetroFlag -maxdepth 2 -type f -print | sort > "${ROOT}/state/retroflag-files.txt" 2>&1 || true
  sha256sum /opt/RetroFlag/* > "${ROOT}/state/retroflag-sha256sum.txt" 2>&1 || true
fi

cp -a /boot/config.txt "${ROOT}/local/boot/config.txt" 2>/dev/null || true
cp -a /boot/cmdline.txt "${ROOT}/local/boot/cmdline.txt" 2>/dev/null || true
cp -a /etc/rc.local "${ROOT}/local/etc/rc.local" 2>/dev/null || true

grep -nEi "vc4|kms|fkms|dpi|dtoverlay|dtparam|hdmi|display|framebuffer|audio|audremap|dwc2|disable_fw_kms_setup|max_framebuffers" \
  /boot/config.txt > "${ROOT}/state/boot-config-display-audio-lines.txt" 2>&1 || true

dmesg -T | grep -Ei "rcu|stall|hung|blocked" > "${ROOT}/logs/kernel-rcu-stall-excerpts.txt" 2>&1 || true
dmesg -T | grep -Ei "under-voltage|undervoltage|voltage|thrott" > "${ROOT}/logs/kernel-voltage-excerpts.txt" 2>&1 || true
dmesg -T | grep -Ei "vc4|v3d|drm|dpi|kms|framebuffer|panel|mailbox|pinctrl" > "${ROOT}/logs/kernel-display-kms-drm-excerpts.txt" 2>&1 || true
dmesg -T | grep -Ei "mmc|ext4|fsck|filesystem|I/O error|Buffer I/O|read-only|remount" > "${ROOT}/logs/kernel-filesystem-mmc-excerpts.txt" 2>&1 || true

journalctl -b -k --no-pager | grep -Ei "rcu|stall|hung|blocked|mmc|ext4|filesystem|under-voltage|undervoltage|voltage|thrott|vc4|drm|dpi|kms|usb|wlan" \
  > "${ROOT}/logs/journal-current-boot-excerpts.txt" 2>&1 || true
journalctl -b -1 -k --no-pager | grep -Ei "rcu|stall|hung|blocked|mmc|ext4|filesystem|under-voltage|undervoltage|voltage|thrott|vc4|drm|dpi|kms|usb|wlan" \
  > "${ROOT}/logs/journal-previous-boot-excerpts.txt" 2>&1 || true

cat /proc/bus/input/devices > "${ROOT}/identity/input-devices.txt" 2>&1 || true
ls -la /dev/input/by-id /dev/input/by-path > "${ROOT}/identity/input-by-id-path.txt" 2>&1 || true
lsusb > "${ROOT}/identity/usb-devices.txt" 2>&1 || true
aplay -l > "${ROOT}/identity/audio-aplay-l.txt" 2>&1 || true
aplay -L > "${ROOT}/identity/audio-aplay-L.txt" 2>&1 || true

ps aux | grep -Ei "SafeShutdown|RetroFlag|lcdfirst|lcdnext|emulationstation|retroflag-powerd" | grep -v grep \
  > "${ROOT}/state/relevant-processes.txt" 2>&1 || true
systemctl list-units --type=service --no-pager | grep -Ei "retroflag|safeshutdown|emulationstation" \
  > "${ROOT}/state/relevant-service-units.txt" 2>&1 || true
systemctl status retroflag-power.service --no-pager > "${ROOT}/state/retroflag-power-service-status.txt" 2>&1 || true

if command -v curl >/dev/null 2>&1; then
  curl -fsSL https://raw.githubusercontent.com/RetroFlag/GPiCase2-Script/main/retropie_SafeShutdown_gpi2.py \
    -o "${ROOT}/upstream/retropie_SafeShutdown_gpi2.py" || true
  curl -fsSL https://raw.githubusercontent.com/RetroFlag/GPiCase2-Script/main/retropielcdfirst.sh \
    -o "${ROOT}/upstream/retropielcdfirst.sh" || true
  curl -fsSL https://raw.githubusercontent.com/RetroFlag/GPiCase2-Script/main/retropielcdnext.sh \
    -o "${ROOT}/upstream/retropielcdnext.sh" || true
fi

if [ -f "${ROOT}/local/opt/RetroFlag/SafeShutdown.py" ] && [ -f "${ROOT}/upstream/retropie_SafeShutdown_gpi2.py" ]; then
  diff -u "${ROOT}/upstream/retropie_SafeShutdown_gpi2.py" "${ROOT}/local/opt/RetroFlag/SafeShutdown.py" \
    > "${ROOT}/diffs/SafeShutdown.local-vs-upstream.diff" 2>&1 || true
fi
if [ -f "${ROOT}/local/opt/RetroFlag/lcdfirst.sh" ] && [ -f "${ROOT}/upstream/retropielcdfirst.sh" ]; then
  diff -u "${ROOT}/upstream/retropielcdfirst.sh" "${ROOT}/local/opt/RetroFlag/lcdfirst.sh" \
    > "${ROOT}/diffs/lcdfirst.local-vs-upstream.diff" 2>&1 || true
fi
if [ -f "${ROOT}/local/opt/RetroFlag/lcdnext.sh" ] && [ -f "${ROOT}/upstream/retropielcdnext.sh" ]; then
  diff -u "${ROOT}/upstream/retropielcdnext.sh" "${ROOT}/local/opt/RetroFlag/lcdnext.sh" \
    > "${ROOT}/diffs/lcdnext.local-vs-upstream.diff" 2>&1 || true
fi

tar -czf "${BUNDLE}" -C "${HOME}" "$(basename "${ROOT}")"
echo "${BUNDLE}"
EOF

sh ~/field-lantern-capture.sh
```

The final line prints the local bundle path, for example:

```text
/home/retropi/field-lantern-20260708T191500Z.tar.gz
```

If upstream downloads fail, keep the bundle. The local evidence is still
useful, and the failure itself can be noted in `report.txt` or inferred from
missing `upstream/` files.

## Pull The Bundle From macOS

From the Mac, pull the bundle with `scp`. Replace the timestamp with the value
printed by the test device.

```sh
mkdir -p ~/Desktop/gpi-case-2-field-lanterns
scp retropi@gpi:/home/retropi/field-lantern-20260708T191500Z.tar.gz \
  ~/Desktop/gpi-case-2-field-lanterns/
```

To inspect before sharing:

```sh
mkdir -p ~/Desktop/gpi-case-2-field-lanterns/review
tar -xzf ~/Desktop/gpi-case-2-field-lanterns/field-lantern-20260708T191500Z.tar.gz \
  -C ~/Desktop/gpi-case-2-field-lanterns/review
```

## Optional SSH Key Setup From macOS

Passwordless SSH can make repeated captures less fiddly, but it is optional.
Use it only for a trusted local device.

Create a key if needed:

```sh
ssh-keygen -t ed25519 -f ~/.ssh/gpi-case-2-field-lantern -C "gpi-case-2-field-lantern"
```

Install the public key:

```sh
ssh-copy-id -i ~/.ssh/gpi-case-2-field-lantern.pub retropi@gpi
```

Pull with that key:

```sh
scp -i ~/.ssh/gpi-case-2-field-lantern \
  retropi@gpi:/home/retropi/field-lantern-20260708T191500Z.tar.gz \
  ~/Desktop/gpi-case-2-field-lanterns/
```

If `ssh-copy-id` is unavailable, use the normal password-based `scp` examples
instead.

## Connection To Common Problems Mage

The Field Lantern is a manual evidence source for the future
[Common Problems Mage Map](common-problems-mage-map.md).

The future mage may classify bundle evidence into issue buckets such as
`power-integrity-warning`, `legacy-lcd-switching-active`,
`kms-profile-active`, `undervoltage-detected`, `unclean-shutdown-detected`,
`controller-detected-as-xbox360`, or `usb-audio-detected`.

Today, no Common Problems Mage is implemented. A human maintainer reads the
bundle and preserves uncertainty where evidence is missing.

## Connection To Local Diagnostics Bundle Map

This Field Lantern is an early manual cousin of the future
[Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md).

The shared rules are:

- Generate locally first.
- Keep the bundle useful offline.
- Preview and redact before sharing.
- Never upload automatically.
- Never require Lantern Dispatch.
- Keep raw observations separate from interpreted meanings.

This page does not implement the future `retroflag-powerd diagnostics`
bundle. It documents a manual field procedure that may later inform that
implementation.

## Connection To Future Lantern Dispatch

Lantern Dispatch remains future optional support, update, issue-reporting, and
submission infrastructure. It is not implemented.

This procedure does not contact Lantern Dispatch and does not prepare an
automatic submission. A future Dispatch trail may accept a user-previewed and
redacted bundle only after a separate implementation quest, consent model, and
submission boundary exist.

## Future Migration Path

The intended trail is staged:

| Stage | Meaning | Status |
| --- | --- | --- |
| Today | Manual documented Field Lantern capture procedure. | This page. |
| Done | Manual Bundle Collector Lantern procedure for post-boot power-integrity evidence. | [Bundle Collector Lantern Capture Procedure](gpi-case-2-boot-power-trace-capture-procedure.md). |
| Later | Safe Boot Power Trace Lantern starts during boot and writes timestamped local samples. | Future quest; no boot startup or systemd activation here. |
| Later | Session Watch Lantern observes runtime throttling, temp, load, memory, frontend/emulator/game clues, and recent warnings. | Future quest; no telemetry by default. |
| Later | `retroflag-powerd diagnostics --bundle`. | Future local-only implementation quest. |
| Later | `retroflag-powerd troubleshoot`. | Future Common Problems Mage classifier quest. |
| Future optional | Lantern Dispatch submission. | Explicitly out of scope and not implemented. |

The win for today is a smaller support trail: one inspected `.tar.gz`, no long
terminal paste, and no accidental step into repair behavior.
