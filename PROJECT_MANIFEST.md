---
id: MANIFEST-001
title: Project Manifest
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Curious Users
purpose: Define the identity, values, design spirit, and guiding promises of RetroFlag Power before production implementation begins.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - ENGINEERING_MANIFESTO.md
  - PROJECT_CHARTER.md
  - docs/13-reference/terminology.md
last_updated: 2026-07-03
---

# Project Manifest

> Modern engineering for timeless games.

RetroFlag Power began with a practical need:

Replace the original RetroFlag SafeShutdown.py script with something modern, reliable, maintainable, and easier to debug.

That need still matters.

But the project has grown into something larger.

This project is about making Raspberry Pi based retro handhelds feel less like small Linux computers and more like purpose-built gaming devices.

It is about shortening the distance between power and play.

It is about preserving the feeling of classic handheld gaming while using modern engineering practices underneath.

It is about making the technology disappear just enough that the player can return to the adventure.

---

# 1. The Mission

RetroFlag Power exists to modernize the software experience around RetroFlag and Raspberry Pi based retro gaming handhelds.

The first target is the RetroFlag GPi Case 2 with a Raspberry Pi CM4.

The first job is safe power management.

The larger mission is to create a polished handheld experience that includes:

- safe shutdown
- fast perceived startup
- automatic game resume
- modern systemd integration
- structured logging
- event-driven hardware handling
- boot and resume performance metrics
- state management
- hardware profiles
- terminal-friendly tools
- clear documentation
- moments of delight

The goal is simple to say and difficult to build:

> Flip the Power Switch and continue the adventure.

---

# 2. What This Project Believes

## A handheld should feel like a console

Users should not have to think about bootloaders, services, GPIO lines, emulator commands, or save-state paths.

Those details matter deeply to the implementation.

They should not dominate the experience.

The ideal experience is:

```text
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

## Engineering and delight belong together

Professional software does not have to feel sterile.

A terminal can be beautiful.

A command can be helpful and fun.

A benchmark can feel like a finish line.

A status screen can make someone smile.

The project should be warm without being gimmicky.

Clear without being cold.

Playful without being confusing.

## Documentation is part of the product

The repository should teach.

A future contributor should not need to reverse engineer the project's purpose.

A future maintainer should not need to guess why a decision was made.

An AI assistant should not need to infer architecture from scattered code.

Important ideas belong in durable artifacts.

## Names matter

Names should reflect real-world behavior.

A latching control is a Power Switch.

A momentary control is a Reset Button.

A persistent condition is a State.

A change in condition is an Event.

The language of the project should help people understand the hardware and software before they read the implementation.

## Momentum matters

Planning is valuable only when it produces artifacts.

Every work session should end with a victory.

A project grows through small completed promises.

The dream matters.

So does building the machine.

---

# 3. The Reference Experience

The reference hardware is:

- RetroFlag GPi Case 2
- Raspberry Pi CM4 Rev 1.1
- 4 GB RAM
- Samsung EVO Select microSD
- RetroPie
- Raspberry Pi OS Bullseye 64-bit
- EmulationStation
- RetroArch

The current experience is functional but imperfect.

The stock system relies on:

```text
/etc/rc.local
      │
      ▼
sudo python /opt/RetroFlag/SafeShutdown.py &
      │
      ▼
GPIO polling loop
```

This works, but it has limits:

- legacy startup mechanism
- root Python process
- polling loop
- limited logging
- no restart policy
- no dependency management
- difficult troubleshooting
- difficult extension path

RetroFlag Power replaces this with a modern service-oriented approach:

```text
systemd
   │
   ▼
retroflag-powerd
   │
   ▼
event-driven hardware handling
   │
   ▼
power, resume, state, and metrics services
```

---

# 4. The Long-Term Experience

The long-term vision is not only safe shutdown.

The long-term vision is a handheld that remembers.

A player turns the system off.

Before shutdown, the software safely records the current session:

- running system
- ROM path
- emulator/core
- save-state slot
- shader
- overlay
- controller profile
- display settings
- volume
- brightness
- timestamp

On the next boot, the system resumes the previous game automatically or presents a polished continue experience.

Possible flow:

```text
Power Switch ON
      │
      ▼
RetroFlag Power starts
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

The user should feel like they are returning to a save point, not managing a Linux session.

---

# 5. The Engineering Shape

The project should begin as a modular monolith.

One daemon.

Multiple internal services.

Clean boundaries.

No unnecessary IPC.

Possible internal services:

- Power Service
- Resume Service
- State Service
- Metrics Service
- Hardware Service
- Frontend Service
- Configuration Service
- Terminal UI Service
- Event Bus

This keeps installation and debugging simple while preserving modularity.

If a component eventually needs to become its own process, it can be extracted later behind the interfaces already created.

---

# 6. Hardware Philosophy

The software should understand capabilities, not just models.

Instead of asking:

```text
Am I running on a GPi Case 2?
```

the system should ask:

```text
Does this hardware provide a Power Switch?
Does this hardware provide a Reset Button?
Does this hardware expose battery information?
Can this platform control a backlight?
Can this system resume a game session?
```

Hardware profiles map real devices to capabilities.

This allows the project to focus first on the GPi Case 2 without closing the door to future hardware.

---

# 7. Portability Philosophy

The project should focus first on Raspberry Pi hardware.

But day-to-day development should not require sitting at a Raspberry Pi.

The project should support development and testing on:

- macOS Apple Silicon
- macOS Intel
- Linux x64
- Linux ARM64
- Linux ARMv7
- Windows where practical for non-hardware code

Hardware-specific behavior should be isolated.

Core services should be testable with mocks and simulated events.

The Raspberry Pi should be required for validation, not for every edit.

---

# 8. Terminal Personality

The command-line experience should feel like it belongs to a retro handheld project.

It should support:

- ASCII art
- color when available
- no-color mode
- readable output over SSH
- compact output for scripts
- clear errors
- friendly success messages
- useful status screens
- benchmark reports that feel satisfying

The tone should be:

- helpful
- warm
- clear
- occasionally playful

It should not be:

- noisy
- childish
- confusing
- legally risky
- dependent on copyrighted imagery

The project may take inspiration from retro gaming culture, but official assets should remain original.

---

# 9. What This Project Will Not Compromise

## Safe shutdown comes first

No resume feature, animation, benchmark, or optimization is worth risking user data.

## Clear errors beat clever messages

Personality belongs in success, status, and discovery moments.

Failure messages must be direct, actionable, and searchable.

## Hardware details stay behind boundaries

Core services should not depend on raw GPIO implementation details.

## Documentation must stay close to decisions

If a major decision is made, it should be documented.

If a major assumption exists, it should be recorded.

If a risk is known, it should be named.

## Planning must lead to building

The project values thoughtful design.

It rejects endless planning.

Every planning session should produce something commit-worthy.

---

# 10. The Project Voice

RetroFlag Power should sound like it was built by people who love both good engineering and good games.

Possible project language:

- Dreaming
- Awakening
- Heartbeat
- Memory
- Momentum
- Adventure
- Launch
- Save Point
- Checkpoint
- Player One
- Ready
- Continue
- Power
- Resume
- Spark

These words are not decoration.

They help create a consistent emotional world for the project.

The project should feel alive, but not pretend to be alive.

It should have personality, not gimmicks.

Warmth, not noise.

Craft, not clutter.

---

# 11. The First Promise

This project will not start with code alone.

It starts with intention.

The first real work is to define:

- why the project exists
- what kind of experience it should create
- who it serves
- what principles guide it
- how decisions are recorded
- how terminology is used
- how contributors are welcomed
- how momentum is protected

Then the software can begin.

The daemon should not merely compile.

It should belong to a project that knows what it is trying to become.

---

# 12. The North Star

If the project succeeds, most users will never think about the hard parts.

They will not think about `systemd`.

They will not think about `libgpiod`.

They will not think about save-state paths.

They will not think about boot order.

They will not think about hardware abstraction.

They will simply flip the Power Switch and continue playing.

That is the magic.

That is the work.

That is why the project exists.

---

# Closing

RetroFlag Power is a technical project.

It is also a craft project.

It is a restoration project.

It is a small love letter to retro handhelds, thoughtful engineering, and the joy of making machines feel a little more alive.

Build it carefully.

Document it honestly.

Name things well.

Measure what matters.

Protect the user.

Respect the hardware.

Welcome contributors.

Leave moments of delight.

And above all:

> Never lose today's excitement in tomorrow's implementation.
