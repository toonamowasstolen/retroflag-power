---
id: PHIL-001
title: Product Philosophy
version: 1.0.0
status: Accepted
owner: Joshua Taft
audience:
  - Everyone
  - AI Assistants
purpose: Apply the engineering manifesto to this specific handheld, and break the tie when two defensible builds are both on the table.
related:
  - WHY.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - ENGINEERING_MANIFESTO.md
last_updated: 2026-07-30
---

# Product Philosophy

> The player should never have to know how hard this was.

This is the project-specific application of
[ENGINEERING_MANIFESTO.md](ENGINEERING_MANIFESTO.md) — how those
cross-project principles express themselves in *this* handheld's design
decisions. The Manifesto is who you are as an engineer; this document is
who this product is.

Where [PROJECT_CHARTER.md](PROJECT_CHARTER.md) settles what the project
promises and [PROJECT_MANIFEST.md](PROJECT_MANIFEST.md) records what it is
made of, this one settles how it should *feel* — the tie-breaker when two
defensible builds are both on the table.

Nothing here is new. It is drawn out of what this project already said it
believed, in Manifest Sections 2, 8 and 9, and written down in one place
so a decision can be checked against it without reading the whole Manifest
first.

## Design values, ranked

When two of these conflict, the higher one wins. That is the entire point
of ranking them.

1. **The player's data survives.** Safe shutdown comes before every
   feature, animation, benchmark and optimization. Nothing else on this
   list is worth a corrupted save.
2. **The illusion holds.** A handheld should feel like a console. Power
   on, brief polished startup, previous game resumes, play. Bootloaders,
   services, GPIO lines and save-state paths matter enormously to the
   implementation and should never dominate the experience.
3. **A failure is understood immediately.** When something breaks, the
   person in front of it finds out what and why, in words they can search
   for. Clarity outranks charm at exactly this moment and nowhere else.
4. **Delight, where it is earned.** A terminal can be beautiful. A
   benchmark can feel like a finish line. This is real work, not
   decoration — but it never comes before the three values above it.
5. **It runs somewhere else too.** Hardware detail stays behind
   boundaries so this project is not permanently married to one case and
   one board.

## What "delightful" means for this specific product

Manifesto principle 10 says to leave moments of delight. Concretely, here:

- The daemon's startup banner — the first thing that says *something is
  awake and glad to be running*.
- A status screen someone would actually choose to look at.
- A benchmark report that feels like a finish line rather than a table.
- The field lantern scripts, which turn a tedious diagnostic capture into
  something with stage labels and a bit of ceremony. They collect evidence
  either way; there is no cost to them being pleasant about it.
- Resume that just works, with no explanation required. The best delight
  here is an absence — the moment the player never notices.

## What this product refuses to be

- It will never risk a save to look impressive.
- It will never make the player think about Linux.
- It will never dress up a failure. A broken thing says so, plainly.
- It will never depend on copyrighted imagery. Inspiration from retro
  gaming culture is welcome; official assets stay original.
- It will never be noisy, childish, or confusing in the name of
  personality.

## Tone and voice

The audience is the one named in [WHY.md](WHY.md) — someone holding a
handheld who wants to play, and occasionally a maintainer at a terminal
trying to work out why it will not.

Helpful, warm, clear, occasionally playful. Not noisy, not childish, not
confusing. The project may sound like it enjoys existing, because it does.

That warmth has exactly one boundary, and it is Manifest Section 9's:
**clear errors beat clever messages.** Personality belongs in success,
status and discovery. A genuine failure gets none of it — direct,
actionable, searchable, first. A calm status line stays calm because the
situation is calm, not because the daemon is forbidden a voice.

## How this product handles failure

Manifesto principle 16 says protect the user from failure. For this
product that means, in order:

1. **Shut down safely anyway.** If the daemon is confused, the write still
   completes and the filesystem is left clean. Confusion is not an excuse
   to lose data.
2. **Say what happened in plain words.** Name the thing that failed and
   the thing to try. No error codes standing alone.
3. **Degrade rather than vanish.** Missing hardware capability means a
   reduced feature, detected and reported — not a crash, and not a silent
   pretence that the feature worked.
4. **Leave evidence.** A failure worth having is a failure someone can
   reconstruct afterwards from the logs and the diagnostics bundle.
