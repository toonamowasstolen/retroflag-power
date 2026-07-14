---
id: QUEST-0046
title: Add the GPi Case 2 KMS Campfire Map
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Preserve the verified GPi Case 2 KMS, GPIO, input, and power findings as a hardware campfire map for future quests.
related:
  - docs/02-hardware/gpi-case-2.md
  - docs/02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md
  - docs/00-project/quests/0045-add-a-hardware-read-only-gpio-probe-command.md
last_updated: 2026-07-07
---

# QUEST-0046 - Add the GPi Case 2 KMS Campfire Map

> Gather the GPi Case 2 field signs into one warm map: the working KMS path,
> the contested pins, the power-save behavior, and the next lanterns to light.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- Added the GPi Case 2 hardware findings field note at
  [docs/02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md](../../02-hardware/gpi-case-2-hardware-findings-kms-power-notes.md).
- Created the hardware field-note index at
  [docs/02-hardware/gpi-case-2.md](../../02-hardware/gpi-case-2.md).
- Preserved the verified win: the GPi Case 2 LCD works under KMS DPI with
  Broadcom V3D rendering.
- Recorded the known-good display charms:
  `dtoverlay=vc4-kms-v3d,noaudio` and
  `dtoverlay=vc4-kms-dpi-generic,clock-frequency=24000000,rgb666-padhi`.
- Documented that `audremap` conflicts with KMS DPI by claiming GPIO12/13.
- Documented that legacy `SafeShutdown.py` uses GPIO26 for the shutdown switch,
  GPIO27 for power enable, and still calls `lcdnext.sh` unless replaced.
- Documented the GPi Case 2 power latch trap: the side switch does not directly
  cut battery power, the stock script drives the power-enable latch HIGH, and
  disabling the script can remove the side-switch shutdown path.
- Preserved the verified `SafeShutdown.py` process tree and its
  `multiprocessing.Process` workers for `poweroff()` and `lcdrun()`.
- Captured the observed GPi Case 2 power-save behavior: the screen turns off,
  the power indicator flashes, SSH stays alive, and the top button wakes the
  case.
- Warned that sleep/power-save can take Wi-Fi down and strand the SSH recovery
  path.
- Recorded that a safe `retroflag-powerd` replacement must own the latch,
  side-switch shutdown detection, and the `lcdrun()` power-save/resume behavior.
- Captured the input-map clue that one extra button above Select and left of
  the RetroFlag logo was not detected during EmulationStation Xbox 360 gamepad
  mapping.
- Named the next quest trail: KMS-safe power keeper, EmulationStation GPU
  guard, button mapping, and power-port cartography.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation only.
- No code changes.
- No runtime script changes.
- No GPIO behavior changes.
- No service install.
- No systemd activation.
- No `rc.local` changes.
- No `SafeShutdown.py` replacement.
- No `lcdnext.sh` edits.
- No boot config edits.

## Milestone Note

The GPi Case 2 now has a verified KMS campfire map in the project docs. Future
power and input quests can gather around one field note instead of rediscovering
which pins are claimed, which overlays work, and where the old RetroFlag scripts
still carry sharp edges.
