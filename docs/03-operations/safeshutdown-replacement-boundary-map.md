---
title: SafeShutdown Replacement Boundary Map
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Map the behavior retroflag-powerd must preserve before it can safely replace the stock RetroFlag SafeShutdown.py path on GPi Case 2 hardware.
related:
  - docs/00-project/project-direction-save-room-arcadia.md
  - docs/04-architecture/arcadia-runtime-migration-path.md
  - docs/03-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - docs/03-operations/gpio-read-only-plan.md
  - docs/03-operations/gpi-case-gpio-probe-ledger.md
  - docs/03-operations/local-diagnostics-bundle-map.md
  - docs/03-operations/installer-migration-toolkit-map.md
  - docs/03-operations/gpi-case-2-acceptance-checklist.md
  - docs/00-project/quests/0048-map-the-safeshutdown-replacement-boundaries.md
  - docs/00-project/quests/0051-map-the-local-diagnostics-bundle.md
  - docs/00-project/quests/0052-map-the-installer-and-migration-toolkit.md
  - docs/00-project/quests/0053-add-the-gpi-case-2-acceptance-checklist.md
last_updated: 2026-07-07
---

# SafeShutdown Replacement Boundary Map

> This is the compass page for the old RetroFlag power spell. Before
> `retroflag-powerd` carries the satchel alone, every behavior below needs a
> verified place on the map.

This document is documentation only. It does not authorize service install,
daemon activation, GPIO writes, shutdown execution, `rc.local` edits, or
replacement of `/opt/RetroFlag/SafeShutdown.py`.

The larger naming and platform compass lives in
[Save Room Tech and Arcadia Runtime Direction](../00-project/project-direction-save-room-arcadia.md).
This boundary map remains focused on the current RetroFlag Power prototype and
the GPi Case 2 SafeShutdown replacement trail.

The staged migration compass toward the broader future engine lives in
[Arcadia Runtime Migration Path](../04-architecture/arcadia-runtime-migration-path.md).
It treats this SafeShutdown boundary map as one prerequisite ledger, not as
permission to replace the stock script.

The future local support satchel is mapped in
[Local Diagnostics Bundle Map](local-diagnostics-bundle-map.md). That
diagnostics compass may summarize this boundary ledger later, but it does not
collect files, upload data, or replace any part of the stock script.

The future reversible installer compass is mapped in
[Installer and Migration Toolkit Map](installer-migration-toolkit-map.md). That
toolkit map may inspect, plan, back up, uninstall, and restore around this
boundary later, but it does not authorize replacing the stock script.

The concrete field-test gate lives in
[GPi Case 2 Acceptance Checklist](gpi-case-2-acceptance-checklist.md). That
ledger must be filled before this boundary map becomes a replacement plan.

## What SafeShutdown.py Currently Owns

Verified GPi Case 2 findings:

- The GPi Case 2 side power switch does not directly cut battery power.
- The stock RetroFlag `/opt/RetroFlag/SafeShutdown.py` script is part of the
  active power-control path.
- The script appears to watch BCM GPIO26 as the side-switch shutdown signal.
- The script drives BCM GPIO27 HIGH as a power-enable latch.
- The script uses separate `multiprocessing.Process` workers for `poweroff()`
  and `lcdrun()`.
- The top power-save/resume behavior is currently part of the legacy
  `lcdrun()` path.
- The RetroFlag legacy script path also appears to participate in docking
  behavior, including switching between the built-in LCD and HDMI when docked.
- Sleep or power-save can make SSH recovery harder if Wi-Fi goes down.
- A 2026-07-08 field incident after resume or power-save behavior showed
  repeated Linux `rcu: INFO: rcu_preempt detected stalls on CPUs/tasks`
  messages, loss of SSH and ping, side-switch off failing to shut down, and
  only physical CM4 cartridge/card removal stopping the device.
- `SafeShutdown.py` was believed to be enabled during that incident, so the
  failure must not be treated as only a disabled stock-script path.
- The GPi Case 2 may auto-enter display/audio power-save after roughly 15-20
  minutes of no input, so avoiding the top power-save button may not avoid the
  risky path.

Observed ownership from the legacy script:

| Behavior | Current owner | Why it matters |
| --- | --- | --- |
| Power-enable latch | `SafeShutdown.py` drives GPIO27 HIGH | The case appears to need this latch held for the power path to remain alive. |
| Side-switch shutdown detection | `poweroff()` watches GPIO26 | The side switch signals software instead of simply cutting battery power. |
| Shutdown sequence | `poweroff()` kills EmulationStation, waits, then calls shutdown | A replacement must preserve clean EmulationStation and Linux shutdown behavior. |
| Top-button power-save/resume path | `lcdrun()` loop and old LCD scripts | The top button wakes the case from power-save, and that behavior cannot vanish silently. |
| Legacy LCD/HDMI switching | `lcdrun()` calls `lcdnext.sh` | This appears to participate in docked HDMI behavior and transitions between the built-in LCD and HDMI. It is not KMS-safe because old scripts can rewrite display configuration. |
| Power-save stall recovery | Unknown: kernel, hardware, display, or power-save path | A field incident suggests software shutdown paths may no longer be reliable after some power-save or resume failures. |

The old script therefore owns more than "shutdown on button press." It is part
of the power latch, side-switch, display/power-save, docking display, and
clean-shutdown trail.

## What retroflag-powerd Already Owns

`retroflag-powerd` currently owns safe, deterministic software paths only:

- Daemon lifecycle startup, signal wait, and clean exit.
- Dry-run power intent processing.
- Noop-only execution policy.
- Fake observer and fake raw-signal CLI paths.
- Raw signal vocabulary: `SignalLow`, `SignalHigh`, and `SignalUnverified`.
- Configured latching power switch interpretation from trusted raw signals.
- A read-only GPIO probe command for one candidate BCM pin.
- Breadcrumb-style diagnostic output for dry-run paths.

These are lanterns and ledgers, not replacement machinery. The project can
observe and model, but it does not yet own real GPi Case power behavior.

## What retroflag-powerd Does Not Own Yet

`retroflag-powerd` does not yet own:

- GPIO27 power-enable latch policy or any GPIO output behavior.
- GPIO26 edge monitoring as a live side-switch shutdown source.
- Debounce, repeat, and boot-state behavior for the side switch.
- Top-button power-save/resume behavior.
- KMS-safe display behavior during top-button power-save/resume.
- Handheld built-in LCD behavior.
- Docked HDMI behavior.
- Transitions between the built-in LCD and HDMI.
- Timing and ordering dependencies around KMS display setup.
- Fully verified audio behavior after the FKMS-to-KMS update.
- Field-tested behavior after automatic display/audio power-save around 15-20
  minutes of no input.
- A reversible emergency reset or safe power-cut recovery path for development.
- Clean EmulationStation shutdown orchestration on real hardware.
- Linux shutdown execution.
- systemd unit installation, enablement, restart policy, or ordering.
- `rc.local` migration.
- Safe coexistence or handoff with the stock RetroFlag script.
- Persistent state across boot, shutdown, power-save, or resume.
- A field-tested replacement plan for `/opt/RetroFlag/SafeShutdown.py`.

Until these are mapped and tested, the stock script remains part of the active
hardware power path.

## Replacement Prerequisites

A safe replacement must preserve these behaviors before any install or
replacement quest begins:

1. Power-enable latch behavior.
2. Side-switch shutdown detection.
3. Top-button power-save/resume behavior.
4. Clean EmulationStation and Linux shutdown sequencing.
5. KMS-safe display behavior.
6. Handheld built-in LCD behavior.
7. Docked HDMI behavior.
8. Transitions between the built-in LCD and HDMI.
9. Any timing or ordering dependencies around KMS display setup.

Each prerequisite needs evidence in the project ledgers:

- Read-only observations for the relevant GPIO or control behavior.
- A documented interpretation policy that separates raw signal from meaning.
- A tested dry-run model before hardware writes or shutdown calls exist.
- A clear service plan that can be reviewed before installation.
- A hardware field checklist with rollback and recovery notes.
- Explicit confirmation that KMS display configuration is not rewritten by old
  LCD/HDMI scripts.
- Field evidence that handheld LCD, docked HDMI, and LCD-to-HDMI or
  HDMI-to-LCD transitions behave correctly under the replacement plan.

The power-enable latch deserves special caution. A future quest may model latch
policy, but the first implementation steps should stay fake or dry-run until
the team has proved when the line must be driven, when it must be released, and
how the device behaves during boot, shutdown, power-save, and resume.

Audio caution: audio after the FKMS-to-KMS update has not been fully verified
yet. Display was the priority fix, so future field testing should include audio
checks in both handheld and docked modes before any replacement plan is
considered complete.

## Explicit No-Go List

This boundary map does not permit:

- GPIO writes.
- Shutdown command execution.
- Service install.
- systemd activation.
- `rc.local` changes.
- `SafeShutdown.py` replacement.
- Persistent state.
- Daemon activation.
- KMS display rewrites.
- `lcdnext.sh` or `lcdfirst.sh` execution from `retroflag-powerd`.
- Assuming handheld LCD, docked HDMI, display transitions, or audio behavior
  are safe without field evidence.
- Assuming power-save, resume, or automatic power-save behavior is safe without
  field evidence.
- Relying on battery depletion or repeated physical CM4 cartridge/card removal
  as a development recovery path.
- Cutting battery leads or modifying lithium battery or charging circuitry
  before the board is mapped.
- Treating one raw probe result as a complete switch map.
- Disabling the stock RetroFlag script on real GPi Case 2 hardware.

If a future quest needs one of these actions, it needs its own acceptance
criteria, hardware checklist, rollback path, and explicit maintainer decision.

## Suggested Future Quest Sequence

The conservative trail should look like this:

1. Read-only probe observations for GPIO26, GPIO27 visibility, top-button
   behavior, and any related case controls.
2. Read-only latch state documentation that records what GPIO27 appears to do
   while the stock script is running, booting, shutting down, and power-saving.
3. Fake or tested latch policy modeling that never writes GPIO on real
   hardware.
4. Controlled dry-run service plan that documents process ownership,
   coexistence with `SafeShutdown.py`, startup order, and failure behavior.
5. systemd unit draft, not installed, with ordering and rollback notes.
6. Hardware field test checklist covering handheld and docked states, LCD and
   HDMI behavior, LCD-to-HDMI and HDMI-to-LCD transitions, KMS display timing,
   audio in both modes, battery and charger conditions, SSH recovery, top-button
   wake, side-switch shutdown, and power-quality observations.
7. Board-recovery investigation for a reversible emergency reset or safe
   power-cut path, including schematics, teardown photos, board labels, test
   pads, CM4 `RUN`/`GLOBAL_EN`/reset/power-enable paths, and regulator enable
   lines.
8. Replacement planning only after the latch, side switch, top-button
   power-save/resume, clean shutdown, KMS-safe display, docking display, display
   transition, and audio-check behaviors are all verified.

This map keeps the quest pointed at replacement readiness without pretending
the relic is ready to leave the satchel today.
