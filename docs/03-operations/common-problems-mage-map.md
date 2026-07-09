---
id: OPS-COMMON-PROBLEMS-MAGE-MAP-001
title: Common Problems Mage Map
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map a future read-only troubleshooting helper that can classify common GPi Case 2 and RetroFlag failure modes from local diagnostics.
related:
  - local-diagnostics-bundle-map.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - gpi-case-2-boot-power-trace-capture-procedure.md
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-session-watch-lantern-design.md
  - gpi-case-2-session-watch-field-run-procedure.md
  - gpi-case-2-safeshutdown-script-behavior-map.md
  - safeshutdown-replacement-boundary-map.md
  - gpi-case-2-acceptance-checklist.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
  - ../03-hardware/gpi-case-2-developer-access-paths.md
last_updated: 2026-07-08
---

# Common Problems Mage Map

> This is the future troubleshooting map for naming common GPi Case 2 failure
> shapes without touching the hardware path. The mage reads the room; it does
> not cast shutdowns, rewrite scripts, or flash anything.

This document is documentation only. It does not implement diagnostics,
generate a diagnostics bundle, read GPIO, write GPIO, execute shutdown, install
or activate systemd, alter `rc.local`, replace or modify
`/opt/RetroFlag/SafeShutdown.py`, change display configuration, run RetroFlag
install scripts, flash firmware, submit telemetry, make network calls, or apply
automatic fixes.

The future local diagnostics and bundle boundary is mapped in
[Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md). This
troubleshooting helper should consume that kind of allowlisted local evidence,
not broaden collection beyond it.

The manual GPi Case 2 capture-bundle trail is documented in
[GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md).
That Field Lantern can provide local evidence for future mage classifications,
but it is not a classifier, fixer, uploader, or `retroflag-powerd diagnostics`
implementation.

The future boot-specific power timing trail is mapped in
[GPi Case 2 Boot Power Trace Lantern Map](gpi-case-2-boot-power-trace-lantern-map.md).
The current manual post-boot bundle collector path lives in
[GPi Case 2 Boot Power Trace Capture Procedure](gpi-case-2-boot-power-trace-capture-procedure.md).
The current collector can contribute remembered boot logs and current
throttling flags. Future true Boot Power Trace Lantern buckets can later help
the mage distinguish early boot, KMS/display initialization,
USB/audio/controller initialization, EmulationStation startup, and idle or
power-save risk patterns.

Power integrity evidence and read-only capture boundaries are mapped in
[GPi Case 2 Power Integrity Investigation Notes](../03-hardware/gpi-case-2-power-integrity-investigation-notes.md).
The runtime session trail is mapped in
[GPi Case 2 Session Watch Lantern Design](gpi-case-2-session-watch-lantern-design.md).
The current read-only skeleton lives at
[`scripts/gpi-case2-session-watch-lantern.sh`](../../scripts/gpi-case2-session-watch-lantern.sh)
and can provide one local Ledger artifact for future session-watch buckets.
Real handheld runs are documented in
[GPi Case 2 Session Watch Field Run Procedure](gpi-case-2-session-watch-field-run-procedure.md),
including the human display, LED, SSH, side-switch, and top-button notes that
the future mage should treat as user-supplied field evidence.
Developer recovery and access candidates, including the inconclusive hidden
firmware USB path, are mapped in
[GPi Case 2 Developer Access Paths](../03-hardware/gpi-case-2-developer-access-paths.md).
KMS, legacy LCD switching, controller, USB audio, and SafeShutdown hardware
findings are preserved in
[GPi Case 2 Hardware Findings and KMS Power Notes](../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md).

## Purpose

A future `retroflag-powerd troubleshoot` command should classify common GPi
Case 2 and RetroFlag failure modes from local, read-only diagnostics.

The helper should turn scattered safe observations into a small set of named
issue buckets, confidence levels, evidence lines, and next read-only checks. It
should help maintainers and users say "this looks like undervoltage plus
legacy LCD switching" instead of hand-sorting logs, config snippets, and field
notes every time.

The helper is a classifier, not a repair tool. It should preserve uncertainty
as `Unknown` or `Inconclusive` when evidence is missing.

## Initial Read-Only Scope

The first scope should stay narrow:

- Read allowlisted local diagnostics already produced or safely collected by
  future diagnostics paths.
- Classify common GPi Case 2 failure modes into stable issue buckets.
- Explain which evidence caused each bucket to appear.
- Recommend next observation-only checks.
- Support terminal text output by default.
- Support JSON output for tests, issue templates, and future bundle use.
- Optionally include bundle-shaped local evidence when the user explicitly
  asks for it.

The helper should not become a daemon, background collector, updater,
installer, migration tool, service manager, firmware utility, or fix runner.

## Problems It Should Classify

The first GPi Case 2 issue buckets should be:

| Bucket | Meaning | Example evidence |
| --- | --- | --- |
| `power-integrity-warning` | Power headroom is suspicious or needs focused follow-up. | Any undervoltage, throttling, voltage-normalized, brownout-like, or context mismatch evidence from safe captures. |
| `boot-power-trace-bucket` | A focused boot trace suggests a timing bucket for undervoltage or throttling. | Early-boot, KMS/display init, USB/audio/controller init, EmulationStation startup, idle/power-save, or no-observed-warning buckets from a Boot Power Trace Lantern. |
| `session-watch-bucket` | A focused runtime watch suggests a menu, emulator, play, idle-risk, or post-resume pattern. | Session Watch Lantern samples, missing-sample gaps, process milestones, thermal/load trails, throttle changes, and narrow runtime log excerpts. |
| `safeshutdown-not-running` | The stock SafeShutdown path appears absent or inactive when expected. | `/opt/RetroFlag/SafeShutdown.py` exists but no matching process or startup reference is visible. |
| `safeshutdown-script-modified` | The local SafeShutdown script differs from the expected stock or captured baseline. | File hash, timestamp, content markers, or missing expected GPIO26/GPIO27 logic. |
| `legacy-lcd-switching-active` | Old RetroFlag LCD/HDMI switching appears active and may rewrite display configuration. | `lcdfirst.sh`, `lcdnext.sh`, GPIO18 polling, or `rc.local` entries for old LCD scripts. |
| `kms-profile-active` | The modern KMS DPI path appears active. | `vc4-kms-v3d`, `vc4-kms-dpi-generic`, `DPI-1`, or `vc4drmfb` evidence. |
| `undervoltage-detected` | The system has direct undervoltage or throttling evidence. | `Undervoltage detected!`, `Voltage normalised`, or non-zero `vcgencmd get_throttled` flags. |
| `unclean-shutdown-detected` | The last shutdown or recent boot history suggests an unsafe stop. | Filesystem recovery messages, journal boot gaps, kernel warnings, or user-supplied recovery notes. |
| `controller-detected-as-xbox360` | The built-in controller is detected through an Xbox 360 style USB HID path. | Microsoft Xbox 360 pad, GBA Pi Case+, Nuvoton, or equivalent input inventory. |
| `usb-audio-detected` | USB audio hardware is present and may be the active audio route. | GeneralPlus USB Audio Device or equivalent safe audio inventory. |
| `hidden-firmware-usb-inconclusive` | A candidate hidden firmware/update USB path was observed, but its role is not proven. | Developer-access notes or user observations show a candidate path without confirmed updater, MCU, bridge, or normal USB identity. |

Buckets may overlap. For example, a device can have both
`kms-profile-active` and `legacy-lcd-switching-active`, which should be
reported as a risky mixed display state rather than collapsed into one answer.

## Inputs It May Safely Read

Initial inputs should be allowlisted, local, read-only facts:

- Existing `retroflag-powerd diagnostics` text or JSON output, when that future
  command grows beyond the current stub.
- A future local diagnostics bundle described by
  [Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md).
- Safe OS and kernel identity summaries.
- Safe Raspberry Pi model and display stack summaries.
- Read-only `/boot/config.txt` excerpts limited to display, audio, and overlay
  lines needed for classification.
- Read-only process observations for `SafeShutdown.py`, RetroFlag LCD scripts,
  and related startup references.
- Read-only file presence, metadata, or approved hashes for
  `/opt/RetroFlag/SafeShutdown.py`, `lcdfirst.sh`, and `lcdnext.sh`.
- Read-only startup references from `/etc/rc.local` or systemd unit inventory.
- Read-only `vcgencmd get_throttled` and voltage warning summaries.
- Narrow kernel log matches for voltage, throttling, filesystem recovery, RCU
  stalls, and display/audio clues.
- Read-only display inventory such as framebuffer, connector, KMS/FKMS, and
  DPI evidence.
- Read-only audio and USB device inventory needed for GPi Case 2 triage.
- Read-only input device inventory needed to identify the built-in controller.
- User-supplied field notes, such as docked vs handheld, battery/USB-C power
  state, sleep/resume behavior, and recovery observations.

It should avoid broad dumps. It must not collect ROM lists, save data,
screenshots, secrets, Wi-Fi details, private keys, full environment dumps,
private paths, hostnames, private IP addresses, or arbitrary journal excerpts.

## Outputs It Should Produce

Default text output should be short and human-readable:

- Overall troubleshooting status.
- Matched issue buckets.
- Confidence for each bucket: `High`, `Medium`, `Low`, or `Inconclusive`.
- Evidence lines for each bucket.
- Missing evidence that would improve confidence.
- Next read-only checks.
- Explicit safety reminders when a bucket touches power, shutdown, display
  switching, firmware, or recovery paths.

JSON output should carry the same information in a stable shape:

- `status`
- `buckets`
- `confidence`
- `evidence`
- `missing_evidence`
- `next_read_only_checks`
- `non_goals`
- `source_sections`

Bundle output should remain local and previewable. A future bundle should make
the classifier result one section inside a user-controlled diagnostics bundle,
not an automatic upload or repair report.

Manual Field Lantern bundles may later become one input to the classifier, but
the mage must treat missing sections, failed upstream downloads, and redacted
fields as explicit uncertainty rather than proof of safety or failure.

## Explicit Non-Goals

The troubleshooting helper must not:

- Write GPIO.
- Toggle GPIO27 or any power-enable latch.
- Execute shutdown, reboot, halt, poweroff, suspend, or resume.
- Kill or restart `SafeShutdown.py`.
- Replace, edit, chmod, download, or reinstall `SafeShutdown.py`.
- Edit `/etc/rc.local`, systemd units, `/boot/config.txt`, or RetroFlag files.
- Run `lcdfirst.sh`, `lcdnext.sh`, or RetroFlag installers.
- Flash firmware.
- Send USB boot payloads.
- Apply automatic fixes.
- Generate or upload telemetry.
- Contact the network.
- Treat a classification as proof that hardware is safe.

## Future Command Sketch

Possible future command forms:

```sh
retroflag-powerd troubleshoot
retroflag-powerd troubleshoot --bundle
retroflag-powerd troubleshoot --format json
```

Expected behavior:

- `retroflag-powerd troubleshoot` prints local read-only classification output.
- `retroflag-powerd troubleshoot --bundle` includes the classification in a
  local diagnostics bundle flow, if a later quest implements bundle creation.
- `retroflag-powerd troubleshoot --format json` prints deterministic
  machine-readable classification output.

All forms should be foreground, local-only, and explicit about unknowns.

## Relationship To Existing Maps

This map depends on these existing ledgers:

- [Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md) for the
  allowlisted local evidence model and bundle boundaries.
- [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
  for the current manual read-only capture bundle used before an implemented
  diagnostics bundle exists.
- [GPi Case 2 Session Watch Field Run Procedure](gpi-case-2-session-watch-field-run-procedure.md)
  for real handheld Session Watch runs and the human field notes that pair
  with one final Ledger artifact.
- [GPi Case 2 SafeShutdown Script Behavior Map](gpi-case-2-safeshutdown-script-behavior-map.md)
  for the stock script, GPIO26, GPIO27, GPIO18, and legacy LCD behavior.
- [SafeShutdown Replacement Boundary Map](safeshutdown-replacement-boundary-map.md)
  for the replacement safety boundary and the reminder that the stock script
  owns more than shutdown.
- [GPi Case 2 Power Integrity Investigation Notes](../03-hardware/gpi-case-2-power-integrity-investigation-notes.md)
  for undervoltage and throttling evidence.
- [GPi Case 2 Developer Access Paths](../03-hardware/gpi-case-2-developer-access-paths.md)
  for developer access candidates and the inconclusive hidden firmware USB
  trail.
- [GPi Case 2 Hardware Findings and KMS Power Notes](../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md)
  for KMS DPI, controller, USB audio, SafeShutdown, and legacy display
  findings.

The mage should stay humble: classify what the local evidence supports, say
what is missing, and leave every hardware-affecting action for a separate,
reviewed quest.
