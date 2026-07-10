---
id: OPS-GPI-CASE-2-TRUE-BOOT-TRACE-EVIDENCE-LEDGER-001
title: GPi Case 2 True Boot Trace Evidence Ledger
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record real GPi Case 2 True Boot Trace Lantern field-run evidence from script artifacts and handheld observations without inventing missing data or changing runtime behavior.
related:
  - ../../scripts/gpi-case2-true-boot-trace-lantern.sh
  - gpi-case-2-true-boot-trace-lantern-design.md
  - gpi-case-2-true-boot-trace-field-run-procedure.md
  - gpi-case-2-first-spark-boot-veil-welcome-scroll-design.md
  - gpi-case-2-boot-power-trace-lantern-map.md
  - gpi-case-2-session-watch-evidence-ledger.md
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
last_updated: 2026-07-09
---

# GPi Case 2 True Boot Trace Evidence Ledger

> The Boot Trace Ledger keeps the First Spark honest: script artifact facts in
> one pouch, handheld notes in another, and every unknown rune left unpainted.

This document records real GPi Case 2 True Boot Trace Lantern field runs from
[`scripts/gpi-case2-true-boot-trace-lantern.sh`](../../scripts/gpi-case2-true-boot-trace-lantern.sh).
It is documentation-only evidence intake. It does not add runtime behavior,
change the script, read or write GPIO, alter display or power configuration,
install services, change RetroPie or EmulationStation settings, replace the
stock SafeShutdown path, contact Lantern Dispatch, or run repairs.

The GPi Case 2 is a handheld Relic. Do not assume an attached keyboard. SSH to
`retropi@gpi` is optional support for copying, launching, and retrieving the
final artifact, not the primary handheld UX. Current field practice is
scp-first: copy the script to `retropi@gpi:/home/retropi/`, run it from
`/home/retropi/`, then copy the final Ledger artifact back. The GPi Case 2
does not currently have the repository checked out on the device.

While the Relic is responsive, the side switch remains the normal stock
shutdown path. The top sleep/resume button remains suspect unless a dedicated
procedure explicitly tests it.

## Purpose

Use this Ledger to preserve True Boot Trace evidence in one place after each
field run:

- Start with the final script Boot Trace Ledger artifact.
- Add only concise handheld observations for first visible screen, LED, SSH,
  EmulationStation, side-switch behavior, and top-button behavior.
- Separate artifact facts from human observations.
- Preserve unknown or not-tested fields as `unknown`, `not tested`, or
  `not provided`.
- Keep interpretation buckets as field labels, not root-cause claims.
- Record candidate improvements without turning them into implementation
  approval.

## Evidence Status Legend

| Status | Meaning |
| --- | --- |
| `intake-pending-artifact` | A field run or observation was reported, but the final artifact path and contents have not been supplied or copied locally. |
| `observed` | Seen once in a field run, with artifact facts and human notes attached. |
| `reproduced` | Seen in more than one comparable field run. |
| `contradicted` | Later evidence conflicts with an earlier observation. |
| `resolved` | A later change or finding explains the item and no current follow-up is needed. |

## Run Entry Template

Copy this template for each real run. Leave unknown fields explicit. Do not
invent values from memory after the trail goes cold.

```text
### YYYY-MM-DD - short run name

Evidence status: intake-pending-artifact | observed | reproduced | contradicted | resolved
Remote artifact path:
Local copied artifact path:
Run date/time:
Requested duration:
Observed duration:
Sample count:
Power source:
Handheld/docked state:
First visible screen timing:
First SSH availability timing:
First EmulationStation timing:
Display state:
LED state:
SSH status:
Frontend / EmulationStation status:
Side-switch behavior:
Top-button behavior:
Warnings/missing evidence:
Final outcome:
Human notes:
Next rune:
Do not overclaim:
```

## Interpretation Buckets

Use one or more buckets when summarizing a run. Buckets are labels for
organizing evidence, not root-cause claims.

| Bucket | Use when | Do not overclaim |
| --- | --- | --- |
| `clean-boot-trace` | The script completed, the final artifact exists, and no display, SSH, frontend, or hard-freeze-like failure was observed during the bounded run. | A clean short boot trace does not prove sleep, resume, shutdown, dock, battery, long idle, or later emulator paths are safe. |
| `first-spark-delay-observed` | The operator observed a measurable or rough delay between side-switch flip and first visible output. | Do not claim the cause or the earliest safe drawing point from the delay alone. |
| `display-blank-before-first-visible-output` | The display appeared blank, inactive, or silent before first visible output. | Do not claim KMS, framebuffer, firmware, userspace, panel, or case-board cause without matching artifact evidence. |
| `ssh-unavailable-until-later` | Human notes show SSH was not available at first and became available later. | SSH availability does not prove what the display, controls, or kernel showed before SSH arrived. |
| `ssh-timing-unknown` | SSH was used or expected but first availability timing was not captured. | Do not derive SSH timing from script start unless the field note says that is the first reachable moment. |
| `frontend-reached` | Human notes or artifact facts show EmulationStation/frontend reached a visible or detected state. | Process detection is not proof of the handheld display unless human notes confirm visibility. |
| `frontend-unknown` | Frontend status was not captured by artifact facts or human notes. | Do not treat unknown frontend status as failed frontend. |
| `frontend-failed` | Artifact facts or human notes show the frontend did not appear or failed during the run. | Keep process, display, and operator observations separate before assigning cause. |
| `no-hard-freeze-observed` | No field evidence showed display, input, and SSH all stopped making progress, and no stall text was reported. | Absence of a freeze in one short run does not clear later idle, sleep, resume, or power paths. |
| `hard-freeze-like-outcome-observed` | Display, input, and SSH stopped making progress, or visible kernel stall text appeared. | This describes the field shape only; preserve last sample, photos, LED state, and recovery trail before guessing. |

## Do Not Overclaim

Observed evidence is what the script artifact, screen, LED, controls, SSH
state, and human note directly record. Guesses are explanations about why the
field shape happened. Keep them separate.

The operator has observed a roughly 15-second blank or silent window after
flipping the GPi Case 2 side power switch before first visible text. Record
that as a handheld observation when it appears. Do not claim KMS,
framebuffer, firmware, userspace, panel, case-board, or power-root cause from
that window unless a later artifact and procedure support the claim.

Avoid claiming:

- A root cause from one run.
- First-visible-screen timing from script start unless the field note says so.
- First SSH availability from artifact start unless the field note says that
  artifact launch was the first successful SSH moment.
- Display state from frontend process detection alone.
- That `vcgencmd get_throttled` proves the exact moment, cause, power supply
  quality, battery health, charger state, or emulator impact.
- That a clean boot trace proves sleep, resume, shutdown, or idle paths.

## Ledger Entries

### 2026-07-10 - First True Boot Trace Field Run

Evidence status: `observed`

Remote artifact path:
`/home/retropi/gpi-case2-true-boot-trace-lantern-20260710-081529.txt`

Local copied artifact path:
[`docs/03-operations/artifacts/true-boot-trace/gpi-case2-true-boot-trace-lantern-20260710-081529.txt`](artifacts/true-boot-trace/gpi-case2-true-boot-trace-lantern-20260710-081529.txt)

Run date/time: `2026-07-10T08:15:29+0800` local on the GPi Case 2;
`2026-07-10T00:15:29Z` UTC in the artifact.

Requested duration: `120` seconds

Observed duration: `127` seconds / `02:07`

Sample count: `25`

Power source: `unknown`

Handheld/docked state: `unknown`

First visible screen timing: Operator has observed a roughly 15-second
blank/silent window after flipping the side power switch before first visible
text. This note is preserved as a handheld observation, not as artifact fact.
No exact first-visible timestamp was supplied for this specific run.

First SSH availability timing: `unknown`; artifact launch proves SSH was
available by `2026-07-10T08:15:29+0800`, but does not prove first availability.

First EmulationStation timing: `unknown`

Display state: Human observation records blank/silent before first visible
text. The artifact records framebuffer/display hints as
`graphics=fb0 fbcon graphics drm=card0 card1 card1-DPI-1 card1-HDMI-A-1 card1-HDMI-A-2 card1-Writeback-1 drm renderD128 fbdev=/dev/fb0`.

LED state: `unknown`

SSH status: `alive for script run`; artifact records `ssh_context:
likely_ssh_session` and the SSH-launched script completed.

Frontend / EmulationStation status: `frontend unknown from handheld view`;
artifact process detection reports `frontend_detected_ever: no`.

Side-switch behavior: `not tested or not provided`

Top-button behavior: `not tested or not provided`

Warnings/missing evidence: Artifact records `warnings_count: 25`,
`missing_evidence_count: 0`. Every warning is the script's cautious
`boot power or throttling hint present` note tied to raw
`throttled=0x50000`.

Final outcome: Boot Trace Lantern completed and wrote the final Ledger
artifact. The script made no runtime changes.

Human notes:

- The GPi Case 2 has shown a roughly 15-second blank or silent window after
  the side power switch is flipped before first visible text appears.
- The GPi was booted and reachable over SSH for this field run.
- No exact LED state, first SSH time, first EmulationStation time, side-switch
  result, or top-button result was provided.

#### Artifact Facts

- `status`: `completed`
- `read_only_marker`: `READ-ONLY / NO CHANGES MADE`
- `artifact_path`:
  `/home/retropi/gpi-case2-true-boot-trace-lantern-20260710-081529.txt`
- `started_utc`: `2026-07-10T00:15:29Z`
- `started_local`: `2026-07-10T08:15:29+0800`
- `ended_utc`: `2026-07-10T00:17:36Z`
- `ended_local`: `2026-07-10T08:17:36+0800`
- `requested_duration_seconds`: `120`
- `observed_duration_seconds`: `127`
- `observed_duration`: `02:07`
- `sample_interval_seconds`: `5`
- `sample_count`: `25`
- `start_proc_uptime_seconds`: `97.68`
- `end_proc_uptime_seconds`: `224.82`
- `hostname`: `raspberrypi`
- `user`: `retropi`
- `kernel`: `Linux raspberrypi 6.1.21-v8+ #1642 SMP PREEMPT Mon Apr  3 17:24:16 BST 2023 aarch64 GNU/Linux`
- `device_model`: `Raspberry Pi Compute Module 4 Rev 1.1`
- `ssh_context`: `likely_ssh_session`
- `systemd_boot_timing_status`: `captured`
- `journal_boot_hints_status`: `captured`
- `dmesg_boot_hints_status`: `captured`
- `frontend_detected_ever`: `no`
- `frontend_first_detected_sample`: `unavailable`
- `display_hint_summary`:
  `graphics=fb0 fbcon graphics drm=card0 card1 card1-DPI-1 card1-HDMI-A-1 card1-HDMI-A-2 card1-Writeback-1 drm renderD128 fbdev=/dev/fb0`
- `throttled_raw_values_observed`: `throttled=0x50000`
- `temperature_min`: `38.9'C`
- `temperature_max`: `45.2'C`
- `warnings_count`: `25`
- `missing_evidence_count`: `0`
- `cautious_timing_bucket`: `display-handoff-or-framebuffer-clues`

#### Human Observations

- Roughly 15-second blank/silent window before first visible text after the
  side switch is flipped.
- SSH was available by the time the script was launched and stayed alive long
  enough for the run to complete and the artifact to be copied back.

#### Unknown Or Not Tested

- Exact power-switch timestamp.
- Exact first visible screen timestamp.
- Whether the first visible screen for this exact run was text, splash/art,
  EmulationStation, or another state.
- First SSH availability timing.
- First EmulationStation visible timing.
- LED state.
- Power source.
- Handheld or docked state.
- Whether EmulationStation was visibly reached on the handheld.
- Whether the side-switch shutdown path was tested after the run.
- Whether the top sleep/resume button was avoided or tested.
- Whether any hard-freeze-like outcome occurred from the handheld view.

#### Interpretation

- `first-spark-delay-observed`
- `display-blank-before-first-visible-output`
- `ssh-timing-unknown`
- `frontend-unknown`

The artifact completed cleanly from the script and SSH perspective, but this
entry does not assign `clean-boot-trace` because frontend visibility, handheld
display state after first text, side-switch behavior, and hard-freeze absence
were not fully observed. It does not assign `frontend-failed` because process
detection alone is not proof of visible frontend failure.

Do not overclaim: The roughly 15-second first-visible delay and the artifact's
`display-handoff-or-framebuffer-clues` timing bucket are not enough to claim
KMS, framebuffer, firmware, userspace, panel, case-board, or power cause.
`throttled=0x50000` is a real raw clue, not a root-cause verdict, battery
verdict, charger verdict, or emulator-performance verdict.

#### Candidate Improvements

- Record exact local timestamps for side-switch flip, first visible output,
  first SSH availability, and first EmulationStation visibility.
- Record power source, charging state if known, handheld/docked state, display
  state, LED state, and whether the top sleep/resume button was avoided.
- Add an operator note before launch when the frontend is visibly present or
  absent so `frontend_detected_ever: no` can be compared with human display
  state.
- If the Relic remains responsive after the run, record whether normal
  side-switch shutdown behaved as expected.
- Keep the next run scp-first from `/home/retropi/` and evidence-only.

Next rune: Repeat one True Boot Trace pass with exact handheld timestamps and
LED/display/frontend notes beside the script artifact, then compare whether
`frontend_detected_ever: no` matches the actual screen.
