---
id: OPS-HUMAN-FACING-FIELD-LANTERN-SCRIPT-UX-STANDARD-001
title: Human-Facing Field Lantern Script UX Standard
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define the terminal UX standard for human-facing manual Field Lantern scripts so long local captures stay readable, alive, and safe.
related:
  - gpi-case-2-field-lantern-capture-procedure.md
  - gpi-case-2-boot-power-trace-capture-procedure.md
  - gpi-case-2-boot-power-trace-lantern-map.md
  - local-diagnostics-bundle-map.md
  - common-problems-mage-map.md
  - ../00-project/project-direction-save-room-arcadia.md
  - ../00-project/edc-quest-operating-rules.md
  - ../00-project/quests/0077-add-a-human-facing-field-lantern-script-ux-standard.md
last_updated: 2026-07-08
---

# Human-Facing Field Lantern Script UX Standard

> A human-facing Field Lantern should feel alive while it works: warm enough
> to build confidence, plain enough to copy, and serious the instant safety or
> evidence is at stake.

This document is documentation only. It does not change Go code, refactor
scripts, read GPIO, write GPIO, execute shutdown, install or activate systemd,
alter `rc.local`, replace `/opt/RetroFlag/SafeShutdown.py`, implement resume,
flash firmware, run installers, apply automatic fixes, submit telemetry,
upload data, or approve hardware modification.

## Scope

Use this standard for manual scripts that a person runs directly in a
terminal, especially Field Lantern, Bundle Collector Lantern, local diagnostic,
support-bundle, and evidence-capture scripts.

This is not a requirement for machine-only helpers, CI-only commands, or
library internals. It is the preferred pattern when a person is watching the
terminal and deciding whether the Relic is still working.

Keep implementations practical for POSIX-ish shell scripts. Prefer shell
builtins, `printf`, `date`, `awk`, `sed`, `wc`, and ordinary core utilities
over heavy dependencies. A script may use richer terminal behavior when it is
already available, but the plain path must remain useful.

## Startup Banner

Human-facing scripts should start with a friendly ASCII banner that identifies
the Relic, names the safety posture, and gives the user a quick sense of what
will happen.

Prefer:

- A small ASCII lantern or related Field Lantern mark.
- Warm Save Room Tech or Arcadia Runtime wording.
- A plain statement that the script is local and read-only when that is true.
- The expected output artifact shape.
- The expected duration or sample count when known.

Do not let the banner hide the command's real purpose. The first screen should
make a cautious maintainer more confident, not make them hunt for facts.

## Step Line Format

Use double-bracket step markers for visible progress:

```text
[[1/6]] 🏮 Preparing capture satchel...
[[2/6]] 📖 Reading boot scrolls...
[[3/6]] ✨ Sampling power runes...
```

The preferred format is:

```text
[[current/total]] glyph Human-readable stage text...
```

Use a small glyph before the human text when stdout is an interactive terminal
and the terminal can reasonably display UTF-8. Good glyphs include `🏮`, `📖`,
`✨`, `🧭`, `🧰`, `💾`, and `✅`.

When glyphs are not safe, output is redirected, `--plain` is set, or the
terminal environment is uncertain, fall back to ASCII-safe labels:

```text
[[1/6]] [lantern] Preparing capture satchel...
[[2/6]] [scroll] Reading boot scrolls...
[[3/6]] [spark] Sampling power runes...
```

Step text should describe the stage in human terms and avoid implying that a
diagnosis or repair happened unless the script truly did that work.

## Stage Labels

Every meaningful phase should have a clear stage label. Use stages to answer
the user's quiet question: "Is it still doing something?"

Good stage labels:

- Name the kind of work: preparing, reading, sampling, bundling, sealing,
  verifying, reporting.
- Name the artifact or evidence class when helpful: boot scrolls, power runes,
  Ledger notes, satchel, map.
- Stay stable enough that docs and support messages can refer to them.

Avoid stage labels that are only decorative. Warm words are welcome; vague
words are not.

## Long-Running Progress

For long-running interactive scripts, show visible progress when the script
can estimate it.

When stdout is an interactive terminal and progress can be estimated, show a
progress bar near the bottom or as a live updating line. Include:

- Completed and total units when known.
- Elapsed time.
- Estimated time remaining when practical.
- Sample counter or countdown when applicable.

Example live line:

```text
Samples [############--------] 54/90 elapsed 00:54 ETA 00:36
```

Do not use a progress bar when:

- Output is redirected.
- `--plain` is set.
- stdout is not an interactive terminal.
- The terminal cannot reasonably support live updates.
- The work cannot estimate progress well enough for the bar to mean anything.

In those cases, print periodic plain progress lines instead:

```text
[[3/6]] [spark] Sample 54/90, elapsed 00:54, remaining about 00:36.
```

Long-running scripts should emit some visible progress at least every few
seconds during expected waits. A quiet terminal during normal capture can look
like a hang, and that false alarm costs trust.

## Timing And Reports

Human-facing manual scripts should record:

- Start time.
- End time.
- Total duration.
- Per-stage durations when practical.

Generated reports or bundles should include timing when useful. This can be a
small `timing.txt`, a section in `manifest.txt`, or a row in a CSV Ledger.

Unusually slow captures are evidence. Treat them as potentially useful
diagnostic material rather than only a UX problem. If a capture takes far
longer than expected, the final report should preserve that fact plainly.

## Final Artifact Line

At the end, print the exact final artifact path. If the script creates both a
folder and an archive, print both.

Prefer:

```text
[[6/6]] ✅ Lantern sealed.
Artifact: /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz
Folder:   /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500
Duration: 00:01:34
```

When retrieval is relevant, include the next manual command using the real
target style for the procedure:

```text
Retrieve from your workstation:
scp retropi@gpi:/home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz .
```

Do not upload, dispatch, submit, or repair automatically.

## Color

Use color for interactive terminals when it helps scanning, but never make
color required for understanding.

Expectations:

- Respect `NO_COLOR`.
- Provide `--plain` for manual scripts where practical.
- Disable color when output is redirected.
- Keep warnings and errors copyable without color.
- Never let color or decorative output hide an error, safety boundary, path,
  command, or validation result.

## Tone

Normal operation can be playful, warm, and adventurous. Field Lanterns may
carry Save Room Tech and Arcadia Runtime language: satchel, scroll, spark,
rune, map, Ledger, Spellbook, Caster, SignalMage, and Relic are welcome when
they clarify the journey.

Warnings should be clear and easy to copy:

```text
WARNING: vcgencmd is unavailable; power sampling will be recorded as unknown.
```

Errors should be serious, direct, and actionable:

```text
ERROR: Could not create output directory: /home/retropi/field-lantern-...
Action: Check free space and permissions, then rerun the command.
```

Safety boundary messages should be plain and unmistakable:

```text
SAFETY: This script will not read GPIO, write GPIO, shut down, reboot, install
services, replace rc.local, or modify SafeShutdown.py.
```

Do not use playful language to soften a safety boundary or an error. The
adventure tone is for confidence; the safety line is for clarity.

## `--help` Expectations

Human-facing scripts should provide useful `--help` output. Include:

- What the script does.
- What the script does not do.
- Safety boundaries.
- Output artifacts.
- Expected duration.
- Important options such as `--plain`, `--output-dir`, `--samples`, or
  `--interval` when present.
- Examples using the real GPi target style when relevant, such as
  `retropi@gpi`.

The help text should be enough for a maintainer to decide whether running the
script is safe for the current moment.

## Example Interactive Transcript

```text
$ ./gpi-case2-bundle-collector-field-lantern.sh

        .-.
       (   )
        | |
      __| |__
     /  | |  \
    /___| |___\
        |_|

GPi Case 2 Bundle Collector Field Lantern
Save Room Tech: local, read-only, human-carried.
SAFETY: no GPIO reads, no GPIO writes, no shutdown, no reboot, no service changes.
Output satchel: /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500
Expected duration: about 90 seconds

[[1/6]] 🏮 Preparing capture satchel...
[[2/6]] 📖 Reading boot scrolls...
[[3/6]] ✨ Sampling power runes...
Samples [############--------] 54/90 elapsed 00:54 ETA 00:36
Samples [####################] 90/90 elapsed 01:30 ETA 00:00
[[4/6]] 🧭 Writing Ledger notes...
[[5/6]] 💾 Sealing archive...
[[6/6]] ✅ Lantern sealed.

Artifact: /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz
Folder:   /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500
Duration: 00:01:34

Retrieve from your workstation:
scp retropi@gpi:/home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz .
```

## Example Plain Transcript

```text
$ ./gpi-case2-bundle-collector-field-lantern.sh --plain

GPi Case 2 Bundle Collector Field Lantern
Save Room Tech: local, read-only, human-carried.
SAFETY: no GPIO reads, no GPIO writes, no shutdown, no reboot, no service changes.
Output satchel: /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500
Expected duration: about 90 seconds

[[1/6]] [lantern] Preparing capture satchel...
[[2/6]] [scroll] Reading boot scrolls...
[[3/6]] [spark] Sampling power runes...
[[3/6]] [spark] Sample 54/90, elapsed 00:54, remaining about 00:36.
[[3/6]] [spark] Sample 90/90, elapsed 01:30, remaining about 00:00.
[[4/6]] [map] Writing Ledger notes...
[[5/6]] [disk] Sealing archive...
[[6/6]] [done] Lantern sealed.

Artifact: /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz
Folder:   /home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500
Duration: 00:01:34

Retrieve from your workstation:
scp retropi@gpi:/home/retropi/gpi-case2-bundle-collector-field-lantern-20260708-191500.tar.gz .
```

## Implementation Notes For Shell Scripts

For POSIX-ish shell scripts:

- Detect interactivity with `[ -t 1 ]`.
- Treat `NO_COLOR` and `--plain` as reasons to avoid color and live terminal
  control.
- Prefer `printf` over `echo` for predictable output.
- Use simple elapsed-time helpers based on epoch seconds when available.
- Keep live progress optional; plain periodic lines are the reliable fallback.
- Avoid dependencies that make the Relic harder to hand-carry to a Pi.

The standard is intentionally small. A human-facing Lantern does not need a
terminal UI framework. It needs to tell the person watching it: "I started, I
am still working, here is what I made, and here is what you can safely do
next."
