---
id: VISION-001
title: Product Vision
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Curious Users
purpose: Define the desired product experience for RetroFlag Power so implementation decisions stay aligned with the goal of making Raspberry Pi handhelds feel more like polished gaming consoles.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - docs/00-project/milestones.md
  - docs/13-reference/terminology.md
  - docs/13-reference/glossary.md
last_updated: 2026-07-03
---

# Product Vision

> Flip the Power Switch. Continue the adventure.

RetroFlag Power exists because the RetroFlag GPi Case 2 is already close to something special.

It looks like a dedicated handheld.

It feels like a dedicated handheld.

It invites the player to treat it like a console.

But under the surface, it is still a Raspberry Pi running Linux, RetroPie, EmulationStation, RetroArch, services, scripts, and configuration files.

That is not a problem.

That is the magic trick.

The purpose of RetroFlag Power is to make the technical layers work together so well that the player does not have to think about them.

The handheld should feel less like a tiny Linux computer and more like a purpose-built retro gaming console.

---

# 1. The North Star

The north star experience is simple:

```
Power Switch ON
      │
      ▼
Brief polished startup
      │
      ▼
Previous game resumes
      │
      ▼
Player continues
```

The player should not need to remember what they were playing.

They should not need to browse menus.

They should not need to manually load a save state.

They should not need to understand which service did what.

They should turn on the handheld and return to the adventure.

---

# 2. The Feeling We Are Building

The product should feel like:

- picking up a Game Boy
- opening a Nintendo Switch
- waking a Steam Deck
- returning to a save point
- continuing a game rather than launching software

The project is not trying to copy any of those products.

It is trying to learn from what they make players feel:

- immediacy
- trust
- continuity
- polish
- confidence
- delight

The player should feel:

```
This thing remembers me.
```

Not:

```
I need to manage Linux.
```

---

# 3. The Current Experience

The current stock experience is functional.

A player can turn on the GPi Case 2, boot RetroPie, launch EmulationStation, select a system, choose a ROM, start a game, and load progress.

But each step adds friction.

```
Power on
   │
   ▼
Wait for boot
   │
   ▼
Wait for frontend
   │
   ▼
Navigate systems
   │
   ▼
Find the game
   │
   ▼
Launch emulator
   │
   ▼
Load state or save manually
   │
   ▼
Resume playing
```

This works.

But it feels like operating a computer.

RetroFlag Power aims to make that same path feel closer to using a console.

---

# 4. The Desired Experience

The desired experience is:

```
Power Switch ON
      │
      ▼
RetroFlag Power starts quietly
      │
      ▼
Resume screen appears
      │
      ▼
Previous game launches
      │
      ▼
Save state restores
      │
      ▼
Player continues
```

The user-facing version might feel like:

```
Resuming...

Super Mario World

Launching...
```

Then the game appears exactly where the player left off.

The player should not feel like they started an app.

They should feel like they returned.

---

# 5. The First Product Promise

The first product promise is safe power handling.

Before the project tries to feel magical, it must be trustworthy.

RetroFlag Power must preserve or improve the safety of the original shutdown behavior.

No visual polish, resume feature, benchmark, animation, or convenience feature is worth risking user data.

The first promise is:

```
When the Power Switch moves to OFF, the system shuts down safely.
```

Everything else builds on that trust.

---

# 6. The Second Product Promise

The second product promise is memory.

The handheld should remember what matters.

Eventually, that may include:

- last played system
- ROM path
- emulator
- RetroArch core
- save-state slot
- shader
- overlay
- aspect ratio
- integer scaling
- controller profile
- volume
- brightness
- last successful resume
- last failed resume

The goal is not to collect state for its own sake.

The goal is continuity.

The system remembers so the player can continue.

---

# 7. The Third Product Promise

The third product promise is visibility.

The project should be easy to operate and debug.

A polished console-like experience does not mean hiding everything from power users.

The project should provide clear tools such as:

```
retroflag-power status
retroflag-power doctor
retroflag-power benchmark
retroflag-power resume
retroflag-power clear
```

The player gets simplicity.

The power user gets visibility.

The maintainer gets logs.

The developer gets tests.

The hardware porter gets profiles.

Good product design serves each audience without making one audience carry the burden of another.

---

# 8. The Fourth Product Promise

The fourth product promise is delight.

RetroFlag Power should have personality.

Not noise.

Not gimmicks.

Personality.

The terminal output should feel like it belongs to a retro handheld project.

The project may use:

- ASCII art
- restrained color
- status cards
- progress bars
- friendly success messages
- playful milestone language
- original pixel-inspired concepts

The project should avoid:

- unclear errors
- excessive jokes
- copyrighted branding
- output that cannot be parsed when needed
- color-only meaning
- personality that gets in the way

A good diagnostic message helps.

A great one helps and feels crafted.

---

# 9. The Player Experience

The Player wants:

- power on
- resume game
- avoid lost progress
- avoid menu friction
- trust shutdown
- enjoy the handheld

The Player should not need to know:

- what GPIO is
- what systemd is
- where save states live
- what `libgpiod` does
- how EmulationStation launches games
- how RetroArch cores behave

For the Player, the product should be simple:

```
Turn it on.

Play.
```

---

# 10. The Power User Experience

The Power User wants control and insight.

They should be able to ask:

```
What was my last game?
Is resume enabled?
Did the last resume work?
How long did boot take?
Is the Power Switch detected?
Is RetroArch available?
What hardware profile is active?
```

The system should answer clearly.

Power users should be able to inspect, troubleshoot, and tune without digging through scattered scripts.

---

# 11. The Developer Experience

The Developer should be able to understand the project quickly.

They should find:

- clear documentation
- consistent terminology
- small packages
- readable interfaces
- tests that run without hardware
- mocks for hardware behavior
- architecture that explains itself
- requirements that trace to implementation
- ADRs explaining major decisions

The repository should feel welcoming, not mysterious.

A good developer experience is part of the product because contributors help the product survive.

---

# 12. The Hardware Porter Experience

The Hardware Porter wants to add support for another case or board.

They should not need to rewrite core services.

They should be able to define:

- hardware profile
- capabilities
- Power Switch behavior
- Reset Button behavior
- battery behavior if available
- backlight behavior if available
- validation notes

The architecture should make this possible without pretending every device is the same.

Hardware support should be honest.

Validated support and experimental support must be clearly distinguished.

---

# 13. The Maintainer Experience

The Maintainer needs project health.

They need:

- clear scope
- clean milestones
- requirements
- risk register
- assumption log
- open questions
- decision records
- release process
- contribution guidelines
- troubleshooting guides

The project should not depend on memory alone.

If a decision matters, it should live somewhere durable.

---

# 14. What Success Looks Like

A successful RetroFlag Power experience might look like this:

```
Player flips the Power Switch ON.

A short resume screen appears.

The previous game launches.

The save state restores.

The player continues within 20 seconds.

Later, they flip the Power Switch OFF.

The game state is saved.

The system shuts down safely.

No progress is lost.

No Linux knowledge was required.
```

A successful maintainer experience might look like this:

```
A bug report includes logs from journalctl.

The active hardware profile is visible.

The last power event is recorded.

The resume state is inspectable.

The failure mode is documented.

The fix is understandable.
```

A successful contributor experience might look like this:

```
A new contributor reads the docs.

They understand the terminology.

They run tests locally on macOS.

They add a mock event test.

They open a pull request.

The project still makes sense.
```

---

# 15. Product Principles

## Principle 1 — Safety before magic

The system must be trustworthy before it is delightful.

## Principle 2 — Resume is continuity, not a trick

Resume should feel like the handheld remembered the player.

## Principle 3 — The interface should respect every audience

Players need simplicity.

Power users need visibility.

Developers need structure.

Maintainers need clarity.

Hardware porters need boundaries.

## Principle 4 — The best technology disappears

The player should not experience implementation details.

## Principle 5 — Delight should reinforce clarity

Personality should make the system feel crafted, not harder to understand.

## Principle 6 — Measure what matters

Boot, shutdown, resume, reliability, and failure recovery should be measurable.

## Principle 7 — Do not overpromise

Experimental features should be marked as research.

Validated features should be documented honestly.

---

# 16. Product Non-Goals

RetroFlag Power is not trying to become:

- a full operating system
- a replacement for RetroPie
- a replacement for RetroArch
- a replacement for EmulationStation
- a cloud gaming platform
- a general desktop power manager
- a brand imitation of any commercial gaming company
- a GUI-first configuration suite
- a project that values flash over trust

The product should stay focused:

```
Power.
Resume.
State.
Metrics.
Polish.
Maintainability.
```

---

# 17. Near-Term Product Path

The near-term path is:

```
Dreaming
   │
   ▼
Awakening
   │
   ▼
Heartbeat
   │
   ▼
Power
```

In practical terms:

1. Finish enough Milestone 0 documentation to guide implementation.
2. Build the first daemon skeleton.
3. Run it under systemd.
4. Replace the original shutdown path safely.
5. Validate behavior on the GPi Case 2 reference platform.

Only after the power foundation is reliable should resume become the primary implementation focus.

---

# 18. Long-Term Product Path

The long-term path is:

```
Safe shutdown
      │
      ▼
Session memory
      │
      ▼
Automatic resume
      │
      ▼
Performance metrics
      │
      ▼
Polished terminal tools
      │
      ▼
Splash/resume experience
      │
      ▼
Hardware profiles
      │
      ▼
Community expansion
```

The product should grow naturally from a reliable core.

---

# 19. Product Voice

The project voice should be:

- clear
- warm
- confident
- curious
- crafted
- gently playful

Possible user-facing words:

- Ready
- Resume
- Continue
- Checkpoint
- Save Point
- Adventure
- Heartbeat
- Memory
- Power
- Player One

Avoid overusing them.

A little magic is memorable.

Too much magic becomes clutter.

---

# 20. The Vision in One Sentence

RetroFlag Power makes Raspberry Pi based retro handhelds feel more like dedicated gaming consoles by combining safe power handling, session resume, performance awareness, clear tooling, and a crafted user experience.

---

# Closing

This project began as a shutdown daemon.

The product vision is bigger:

Make the handheld feel remembered.

Make the software feel trustworthy.

Make the terminal feel welcoming.

Make the architecture feel maintainable.

Make the documentation feel alive.

Make the technology disappear just enough that the player can keep playing.

The dream is simple:

```
Flip the Power Switch.

Continue the adventure.
```
