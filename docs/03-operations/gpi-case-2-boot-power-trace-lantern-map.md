---
id: OPS-GPI-CASE-2-BOOT-POWER-TRACE-LANTERN-MAP-001
title: GPi Case 2 Boot Power Trace Lantern Map
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map a future focused read-only boot power trace for timing GPi Case 2 undervoltage and throttling evidence during startup.
related:
  - gpi-case-2-boot-power-trace-capture-procedure.md
  - ../../scripts/gpi-case2-boot-power-trace-field-lantern.sh
  - gpi-case-2-field-lantern-capture-procedure.md
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
  - ../03-hardware/gpi-case-2-power-integrity-investigation-notes.md
  - ../03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - ../03-hardware/gpi-case-2-developer-access-paths.md
last_updated: 2026-07-08
---

# GPi Case 2 Boot Power Trace Lantern Map

> The Boot Power Trace Lantern is a small timing lantern for the first moments
> of boot: enough light to see when the power warning appears, without touching
> the power path.

This map is paired with one portable read-only script:
[`scripts/gpi-case2-boot-power-trace-field-lantern.sh`](../../scripts/gpi-case2-boot-power-trace-field-lantern.sh).
That script is a hand-carried Field Lantern Relic for the Pi; it is not an
installer, service, daemon, or Arcadia Runtime activation path. It does not
require git, a repository checkout, Go, root-only writes, or project install on
the GPi Case 2.

This map and its script do not change Go code, read GPIO, write GPIO, execute
shutdown, install or activate systemd, alter `rc.local`, replace
`/opt/RetroFlag/SafeShutdown.py`, implement resume, flash firmware, submit
telemetry, upload data, make project-code network calls, apply automatic
fixes, run RetroFlag installers, or approve hardware modification.

Nothing here approves cutting battery leads, modifying lithium battery or
charging circuitry, blind soldering, shorting unknown pads, relying on battery
depletion as recovery, or treating the side switch as a reliable emergency
stop during a kernel stall.

Boot Power Trace Lantern means a focused read-only power and boot timing
capture. Field Lantern means the broader read-only local capture bundle.
Common Problems Mage means a future classifier or troubleshooter. Lantern
Dispatch means a future optional support, update, or submission layer. Lantern
Dispatch is not implemented.

The current manual capture procedure lives in
[GPi Case 2 Boot Power Trace Capture Procedure](gpi-case-2-boot-power-trace-capture-procedure.md).
It provides the copy, run, and retrieve path for the portable shell script. The
script does not run automatically.

## Purpose

Map a future Boot Power Trace Lantern for the GPi Case 2 that records when
undervoltage or throttling appears during startup.

The trace should help separate these candidate timing buckets:

- Early power-rail sag before userspace is fully available.
- KMS, DRM, VC4, or DPI display initialization load.
- USB audio, controller, input, or hub initialization load.
- EmulationStation or runtime startup load.
- Later idle or power-save risk.

The lantern should preserve timing evidence and uncertainty. It should not
claim a root cause from one warning flag.

## Why This Exists

Current power evidence shows undervoltage is real enough to track, but not yet
specific enough to diagnose.

Known evidence:

- Undervoltage appears every boot in the current field observations, whether
  plugged in or on battery.
- `vcgencmd get_throttled` returned `0x50000`.
- Current evidence suggests a transient boot or initialization sag, not
  necessarily constant undervoltage.
- KMS may be an aggravating factor by changing display and DRM initialization
  timing, but this is not proven.

The important next question is not only "did undervoltage happen?" It is:

Did undervoltage appear before userspace, during KMS/DRM display bring-up,
during USB audio/controller bring-up, during EmulationStation startup, or
later during idle or power-save?

## Core Question

The future trace should answer this as narrowly as possible:

| Timing window | Question |
| --- | --- |
| Early boot | Did throttling or undervoltage appear before a clearer subsystem milestone? |
| KMS/display init | Did voltage warnings cluster near VC4, DRM, DPI, framebuffer, or panel initialization? |
| USB/audio/controller init | Did warnings cluster near USB, input, HID, hub, or audio device initialization? |
| EmulationStation startup | Did warnings appear when EmulationStation or the launcher stack became active? |
| Idle/power-save | Did warnings appear only after boot, idle, display-off, or power-save behavior? |
| No observed warning | Did the trace complete with no current or historical throttling flags? |

The trace should keep "current undervoltage" separate from "undervoltage has
occurred since boot" when interpreting `vcgencmd get_throttled`.

## Safe Capture Fields

The Boot Power Trace Lantern should collect only allowlisted, read-only facts:

- Timestamp.
- Monotonic uptime, such as `/proc/uptime` or an equivalent elapsed time.
- `vcgencmd get_throttled`.
- `vcgencmd measure_volts`.
- `vcgencmd measure_temp`.
- Recent `dmesg` lines matching voltage, throttle, KMS, DRM, VC4, DPI, USB,
  audio, input, MMC, or ext4.
- Process milestone observations when available, such as whether
  `emulationstation` is running.
- Manual context notes when supplied by the maintainer, such as battery,
  USB-C power, docked or handheld state, visible display state, and whether the
  device is being kept active to avoid idle power-save.

If a command is unavailable, the trace should record that absence as evidence.
It should not install packages, change configuration, enable services, or
broaden collection to make the command available.

## Suggested Trace Duration

The first future capture should stay short:

- Capture the first 90 seconds after boot.
- Optionally repeat after 5 minutes idle, as long as the maintainer keeps the
  device from entering risky automatic power-save during the test.
- Optionally repeat before and after power-save only after the emergency
  recovery path is better understood.

Do not keep the GPi Case 2 idle long enough to trigger automatic power-save
during risky testing. Do not test power-save or resume until the recovery path
is improved and documented.

## Output Shape

The output should be local, previewable, and easy to include inside later
diagnostics:

- A timestamped report file.
- A CSV-style table for repeated samples.
- A short human-readable summary naming the strongest timing bucket.
- Unavailable command and file markers instead of capture failure.
- A `.tar.gz` bundle beside the timestamped folder.
- No automatic upload.
- No Lantern Dispatch contact.
- No automatic fix recommendations.
- Optional inclusion inside future Field Lantern bundles.

Example table shape:

| captured_at | uptime_seconds | throttled | volts | temp_c | emulationstation | recent_bucket_hint |
| --- | --- | --- | --- | --- | --- | --- |
| `2026-07-08T00:00:05Z` | `5.2` | `0x0` | `1.2000V` | `48.2` | `not-running` | `early-boot` |
| `2026-07-08T00:00:22Z` | `22.1` | `0x50000` | `1.2000V` | `50.1` | `not-running` | `kms-display-init` |

The table is an output sketch only. It is not implemented by this document.

## Interpretation Buckets

Future tooling or manual review may classify traces into these buckets:

| Bucket | Meaning | Caution |
| --- | --- | --- |
| `early-boot-undervoltage` | Voltage or throttling evidence appears before a clearer subsystem milestone. | May require earlier boot logs or serial evidence to narrow further. |
| `kms-display-init-undervoltage` | Evidence clusters near VC4, DRM, DPI, framebuffer, or panel initialization. | KMS may be timing-related, but this does not prove KMS is the root cause. |
| `usb-audio-controller-init-undervoltage` | Evidence clusters near USB hub, HID/controller, input, or USB audio initialization. | USB load may overlap with display and runtime startup. |
| `emulationstation-startup-undervoltage` | Evidence appears when EmulationStation or launcher processes start. | Confirm renderer state and avoid confusing historic flags with current sag. |
| `idle-power-save-risk` | Evidence appears after boot during idle, display-off, wake, or power-save paths. | Do not keep testing power-save without a safer recovery path. |
| `no-undervoltage-observed` | No current or historical undervoltage is visible during the capture window. | Absence in one trace does not prove power headroom is safe. |

Buckets may overlap. The future Common Problems Mage should report overlapping
buckets as evidence, not collapse them into a single confident answer.

## Safety Notes

- Keep the capture read-only and local.
- Do not run GPIO probes as part of this lantern.
- Do not keep the GPi idle long enough to trigger automatic power-save during
  risky testing.
- Do not test power-save or resume until recovery is improved.
- Do not treat side-switch shutdown as reliable emergency recovery during a
  kernel stall.
- Do not treat `0x50000` or any single throttling value as a complete root
  cause without timing context.
- Do not rewrite KMS, display, audio, boot, or runtime configuration during the
  capture pass.

If a trace finds MMC, ext4, RCU stall, repeated undervoltage, or lost recovery
evidence, record it in the relevant EDC map or ledger before deeper tests.

## Relationship To Other Lanterns

The Boot Power Trace Lantern is narrower than the Field Lantern.

The [GPi Case 2 Field Lantern Capture Procedure](gpi-case-2-field-lantern-capture-procedure.md)
collects broader post-recovery evidence: boot config, RetroFlag script
provenance, logs, process state, input identity, USB identity, audio identity,
and local-vs-upstream script comparison.

The Boot Power Trace Lantern should eventually become one optional section
inside Field Lantern bundles, focused only on startup timing and power-warning
evidence.

The [Common Problems Mage Map](common-problems-mage-map.md) can later consume
the trace buckets to classify power and boot timing patterns. The mage must
keep unknowns visible and must not apply fixes.

The [Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md) remains the
privacy and bundle boundary for any future `retroflag-powerd diagnostics`
integration.

## Future Migration Path

Expected path:

1. Done: docs-only map.
2. Done: portable read-only Field Lantern script and copy/run/retrieve
   procedure.
3. Later: Field Lantern bundle includes a boot power trace section.
4. Later: Common Problems Mage classifies power buckets.
5. Later: `retroflag-powerd diagnostics --bundle` can include the local trace
   when explicitly requested.
6. Later: `retroflag-powerd troubleshoot` can classify trace evidence without
   changing the device.

Each step needs its own quest, review, and validation. This page only lights
the map.
