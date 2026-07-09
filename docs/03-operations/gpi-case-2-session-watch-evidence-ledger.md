---
id: OPS-GPI-CASE-2-SESSION-WATCH-EVIDENCE-LEDGER-001
title: GPi Case 2 Session Watch Evidence Ledger
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Record real GPi Case 2 Session Watch Lantern field-run evidence from handheld sessions without inventing data or changing runtime behavior.
related:
  - ../../scripts/gpi-case2-session-watch-lantern.sh
  - gpi-case-2-session-watch-lantern-design.md
  - gpi-case-2-session-watch-field-run-procedure.md
  - gpi-case-2-field-lantern-capture-procedure.md
  - common-problems-mage-map.md
  - local-diagnostics-bundle-map.md
  - gpi-case-2-recovery-first-field-procedure.md
last_updated: 2026-07-09
---

# GPi Case 2 Session Watch Evidence Ledger

> The Ledger keeps the Field Lantern honest: one script artifact, concise
> human notes, and a clear rune for what the Relic actually showed.

This document records real GPi Case 2 Session Watch Lantern field runs from
[`scripts/gpi-case2-session-watch-lantern.sh`](../../scripts/gpi-case2-session-watch-lantern.sh).
It is a documentation-only evidence ledger. It does not add runtime behavior,
change the script, collect telemetry, contact Lantern Dispatch, read or write
GPIO, alter display or power configuration, or replace the stock
SafeShutdown path.

The GPi Case 2 is a handheld Relic. Do not assume an attached keyboard. SSH to
`retropi@gpi` is optional support for copying, launching, or retrieving the
final artifact, not the primary handheld experience. While responsive, the
side switch remains the normal stock shutdown path. The top sleep/resume
button remains suspect unless a procedure explicitly says otherwise.

## Purpose

Use this Ledger to preserve Session Watch evidence in one place after each
field run:

- Prefer one final script Ledger artifact from the Session Watch Lantern.
- Add only concise human observations for the screen, LED, controls, buttons,
  and physical state.
- Separate observed evidence from guesses.
- Preserve unknowns as `unknown` or `not tested` until a real run fills them.
- Keep clean runs, display-blank trails, SSH loss, side-switch surprises,
  top-button oddities, and hard-freeze or RCU-stall-like outcomes comparable.

## Evidence Status Legend

| Status | Meaning |
| --- | --- |
| `untested` | No real handheld field run has exercised this item yet. |
| `observed` | Seen once in a field run, with artifact or human notes attached. |
| `reproduced` | Seen in more than one comparable field run. |
| `contradicted` | Later evidence conflicts with an earlier observation. |
| `resolved` | A later change or finding explains the item and no current follow-up is needed. |

## Run Entry Template

Copy this template for each real run. Leave unknown fields explicit. Do not
invent values from memory after the trail goes cold.

```text
### YYYY-MM-DD - short run name

Status: untested | observed | reproduced | contradicted | resolved
Script artifact path/name:
Script artifact retrieved: yes | no | pending | not applicable
Run date:
Run duration:
Power source: battery | USB-C | dock | unknown
Docked/handheld state: handheld | docked | moved during run | unknown
Display state:
LED state:
SSH status: not used | alive | lost | recovered | unknown
Side-switch behavior:
Top-button behavior:
EmulationStation status:
Final outcome:
Interpretation bucket:
Human notes:
Next rune:
Do not overclaim:
```

## Interpretation Buckets

Use one or more buckets when summarizing a run. Buckets are labels for
organizing evidence, not root-cause claims.

| Bucket | Use when | Do not overclaim |
| --- | --- | --- |
| `clean-run` | The script completed, the display remained usable, optional SSH stayed alive if used, and the final artifact exists. | A clean short run does not prove all sleep, resume, dock, battery, or long-idle paths are safe. |
| `display-blank-ssh-alive` | The display blanked, dimmed, or froze while optional SSH stayed alive or the script artifact was retrievable. | Do not call this a sleep transition unless the tested path is known. |
| `ssh-lost` | Optional SSH disconnected, stopped responding, or could not retrieve the artifact during the run. | SSH loss alone does not prove the handheld UI froze. |
| `side-switch-failure` | The normal side-switch shutdown path did not behave as expected while the handheld seemed responsive. | Do not assume GPIO, SafeShutdown, or hardware cause without separate evidence. |
| `top-button-oddity` | The top button was pressed accidentally or under an approved procedure and produced surprising display, resume, SSH, LED, or freeze behavior. | The top button remains suspect; do not generalize from accidental presses. |
| `hard-freeze-rcu-stall-like` | Display, input, and SSH stopped making progress, or visible kernel stall text appeared. | This describes the field shape only; preserve photos, times, LED state, and last artifact path before guessing. |

## Do Not Overclaim

Observed evidence is what the script artifact, screen, LED, controls, SSH
state, and human note directly record. Guesses are explanations about why the
field shape happened. Keep them separate.

Good Ledger language:

- `Observed: display blank at 12:18 local; SSH remained alive; artifact
  retrieved.`
- `Guess: display path or power-save behavior may be involved. Needs repeat
  run before classification.`

Avoid claiming:

- A root cause from one run.
- That the top sleep/resume transition is understood unless the Lantern was
  already running under a procedure that explicitly tested it.
- That SSH status alone proves the handheld screen, controls, or kernel state.
- That a post-resume artifact proves transition-time behavior unless the
  watcher covered the transition.

## Ledger Entries

### Awaiting First Real Field Run

Status: `untested`

Script artifact path/name: `awaiting first real field run`

Script artifact retrieved: `pending`

Run date: `pending`

Run duration: `pending`

Power source: `unknown`

Docked/handheld state: `unknown`

Display state: `unknown`

LED state: `unknown`

SSH status: `not used`

Side-switch behavior: `not tested`

Top-button behavior: `not tested`

EmulationStation status: `unknown`

Final outcome: `awaiting first real handheld Session Watch Lantern run`

Interpretation bucket: `untested`

Human notes: No real Session Watch Lantern field run has been recorded in
this Ledger yet. The first entry should come from the actual script artifact
and concise human notes, not reconstructed expectations.

Next rune: Run the
[GPi Case 2 Session Watch Field Run Procedure](gpi-case-2-session-watch-field-run-procedure.md)
on a responsive handheld Relic, retrieve the final artifact when possible, and
replace this placeholder with the first observed entry.

Do not overclaim: This placeholder records readiness only. It is not evidence
that the GPi Case 2 display, SSH, side switch, top button, EmulationStation,
idle, resume, or power paths behaved in any particular way.
