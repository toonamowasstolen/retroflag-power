---
id: CHARTER-001
title: Project Charter
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Curious Users
purpose: Define the formal scope, objectives, stakeholders, principles, constraints, risks, and success criteria for RetroFlag Power.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - ENGINEERING_MANIFESTO.md
  - PROJECT_MANIFEST.md
  - docs/00-project/quests/0033-add-the-project-charter.md
  - docs/01-product/vision.md
  - docs/13-reference/terminology.md
last_updated: 2026-07-07
---

# Project Charter

> A good project begins with intention.

This charter defines what RetroFlag Power is, why it exists, who it serves, and how success will be measured.

It is not a detailed implementation plan.

It is the agreement that keeps future implementation aligned with the original purpose.

---

# Working Spirit

RetroFlag Power should feel like restoring a handheld console with a tiny
adventurer's toolkit: careful, warm, practical, and a little magical. The work
should protect the device, respect the player, and leave the repository easier
to navigate for the next traveler with a lantern.

These principles guide the daily craft:

1. Celebrate small verified wins.
2. Move in short quests.
3. Keep code plain and maintainable.
4. Keep flavor in docs, quests, milestones, and project framing, not forced into
   low-level identifiers.
5. Do not rush hardware behavior.
6. Challenge weak ideas kindly.
7. Reviews should name the victory, not only the diff.
8. Handoffs should carry both state and spirit.

The project voice belongs in the map, the ledger, the quest records, and the
field notes. Go package names, exported types, filenames, and low-level
technical identifiers should stay clear and ordinary when plain engineering
language is the better compass.

---

# 1. Project Name

## Primary repository name

`retroflag-power`

## Working product name

RetroFlag Power

## Project stage

Milestone 0 — Dreaming

The project is currently defining its purpose, philosophy, architecture, documentation structure, and development practices before production implementation begins.

---

# 2. Project Purpose

RetroFlag Power exists to modernize the software experience around RetroFlag and Raspberry Pi based retro gaming handhelds.

The initial purpose is to replace the original RetroFlag `SafeShutdown.py` implementation with a modern, reliable, event-driven, systemd-managed service.

The broader purpose is to make Raspberry Pi based retro handhelds feel more like polished consumer gaming devices.

The project should reduce the friction between powering on a handheld and continuing a game.

---

# 3. Problem Statement

The current RetroFlag GPi Case 2 software experience works, but it has several limitations.

The original shutdown flow depends on:

```
/etc/rc.local
      │
      ▼
sudo python /opt/RetroFlag/SafeShutdown.py &
      │
      ▼
GPIO polling loop
```

This approach has problems:

- It depends on legacy startup behavior.
- It runs a root Python process.
- It continuously polls GPIO.
- It has limited structured logging.
- It has no restart policy.
- It has weak dependency management.
- It is difficult to debug.
- It is difficult to extend.
- It does not support a modern console-like resume experience.

At the user experience level, powering on the handheld still feels like booting a Linux computer:

```
Power on
   │
   ▼
Wait for boot
   │
   ▼
Navigate frontend
   │
   ▼
Find game
   │
   ▼
Launch game
   │
   ▼
Load state manually
```

RetroFlag Power aims to turn this into:

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

---

# 4. Mission Statement

RetroFlag Power makes retro handhelds feel more like consoles by combining safe power handling, game session resume, fast startup work, thoughtful terminal tools, and maintainable engineering.

The project should make Linux disappear just enough that the player can continue the adventure.

---

# 5. Reference Platform

The first-class reference platform is:

- Device: RetroFlag GPi Case 2
- Compute Module: Raspberry Pi CM4 Rev 1.1
- Memory: 4 GB RAM
- Storage: Samsung EVO Select microSD
- Operating System: RetroPie on Raspberry Pi OS Bullseye 64-bit
- Frontend: EmulationStation
- Emulator Framework: RetroArch
- Current EEPROM: 2022-04-26

This reference platform is the first target for design, testing, and validation.

---

# 6. Scope

## In scope

The project includes:

- Safe shutdown replacement for RetroFlag GPi Case 2.
- systemd service integration.
- Event-driven Power Switch handling.
- Reset Button handling.
- Structured logging through journald.
- Hardware abstraction through capabilities and profiles.
- Game session tracking.
- Automatic save-state and resume research.
- Resume Manager design and implementation.
- Boot and resume performance measurement.
- Terminal-friendly companion CLI.
- Documentation-first engineering process.
- AI-assistant-friendly repository context.
- Development workflows that support macOS and Linux hosts.
- Cross-compilation for Raspberry Pi targets.
- Future hardware profile support.

## Out of scope for the initial milestone

The following are not part of the first implementation milestone:

- Full multi-hardware support.
- Complete battery integration.
- Kernel-level suspend or hibernation.
- GUI configuration application.
- Web dashboard.
- Cloud services.
- Network account systems.
- Official replacement of RetroPie, EmulationStation, or RetroArch.
- Unsupported copyrighted branding or artwork.
- Features that risk data loss before they are proven safe.

## Research-only for now

The following are valid research areas, but should not be treated as committed features until proven:

- Sleep mode.
- Battery percentage overlay.
- Critical battery shutdown countdown.
- Room-transition autosave.
- Per-core autosave intelligence.
- KMS/FKMS performance changes.
- EEPROM boot-order optimization.
- Support for non-Raspberry Pi SBCs.
- Support for x64 handheld systems.

---

# 7. Primary Objectives

## Objective 1 — Modernize safe shutdown

Replace the current `SafeShutdown.py` and `rc.local` flow with a supervised systemd service.

Expected benefits:

- automatic restart
- structured logs
- startup ordering
- dependency management
- watchdog support
- easier debugging
- cleaner installation

## Objective 2 — Eliminate legacy polling scripts

Move away from continuous polling loops where event-driven APIs are available.

The preferred GPIO direction is `libgpiod`.

## Objective 3 — Preserve user data

Safe shutdown and state preservation must never become less reliable than the original behavior.

Resume features must not endanger save files or user progress.

## Objective 4 — Make startup feel faster

The project should improve both measured and perceived startup time.

The current target is:

- Power Switch ON to resumed gameplay in less than 20 seconds.
- Stretch goal: less than 15 seconds.

## Objective 5 — Resume the previous game

The long-term experience should allow the user to power on the handheld and continue the previous game with minimal or no menu navigation.

## Objective 6 — Build a maintainable platform

The project should be modular, documented, testable, portable, and approachable for future contributors.

## Objective 7 — Make the interface enjoyable

Terminal output, status screens, benchmarks, and diagnostics should be useful and enjoyable without sacrificing clarity.

---

# 8. Non-Goals

RetroFlag Power is not intended to be:

- A full operating system.
- A replacement for RetroPie.
- A replacement for RetroArch.
- A replacement for EmulationStation.
- A general-purpose desktop power manager.
- A cloud-connected gaming platform.
- A commercial clone of Nintendo, Sega, Valve, or RetroPie branding.
- A project that values cleverness over clarity.
- A project that plans forever without implementation.

---

# 9. Stakeholders and Personas

## Player

The player wants the handheld to feel like a console.

They care about:

- fast startup
- reliable shutdown
- returning to the last game
- not losing progress
- minimal technical friction

They do not want to think about Linux internals.

## Power User

The power user tweaks systems, cores, shaders, overlays, controllers, and boot behavior.

They care about:

- diagnostics
- status commands
- configuration
- logs
- benchmarks
- control

## Developer

The developer contributes code, documentation, tests, or tooling.

They care about:

- clear architecture
- good naming
- stable interfaces
- local development
- tests
- CI
- documentation

## Hardware Porter

The hardware porter wants to support another case, board, or SBC.

They care about:

- hardware profiles
- capability abstraction
- GPIO mapping
- platform documentation
- test strategy

## Maintainer

The maintainer cares about project health over time.

They need:

- decision records
- requirements
- risk register
- issue templates
- release process
- clean commit history
- sustainable scope

## AI Assistant

AI assistants are treated as future contributors that need context.

They need:

- project memory
- architecture overview
- coding standards
- terminology
- requirements
- explicit constraints
- clear file organization

---

# 10. Guiding Principles

The project is guided by these principles:

1. Build experiences, not just software.
2. Safe shutdown comes first.
3. Documentation is part of the product.
4. Design before implementation.
5. Architecture exists to enable implementation, not delay it.
6. Every work session ends with a victory.
7. Respect future contributors.
8. Name things according to real-world behavior.
9. Make complexity invisible to users.
10. Leave moments of delight.
11. Measure before optimizing.
12. Portability is a feature.
13. Hardware-specific code lives behind interfaces.
14. Use capabilities, not assumptions.
15. No busy polling where reliable events are available.
16. Keep errors clear, actionable, and searchable.
17. Keep personality warm, not gimmicky.
18. Never lose today's excitement in tomorrow's implementation.

---

# 11. Terminology Rules

The project must use precise terminology.

## Power Switch

Use `Power Switch` for latching ON/OFF hardware controls such as the GPi Case 2 and NESPi-style case power controls.

Do not call these controls `Power Buttons` unless the hardware is actually momentary.

## Reset Button

Use `Reset Button` for momentary reset controls.

## State

Use `State` for persistent conditions.

Examples:

- Power State
- Resume State
- Session State

## Event

Use `Event` for a change in condition.

Examples:

- Power Switch moved to OFF
- Reset Button pressed
- Game launched
- Resume failed

## Capability

Use `Capability` for a feature exposed by a hardware profile or platform.

Examples:

- Power Switch capability
- Reset Button capability
- Battery capability
- Backlight capability

---

# 12. High-Level Architecture

The initial architecture should be a modular monolith.

```
systemd
   │
   ▼
retroflag-powerd
   │
   ├── Power Service
   ├── Resume Service
   ├── State Service
   ├── Metrics Service
   ├── Hardware Service
   ├── Frontend Service
   ├── Configuration Service
   ├── Terminal UI Service
   └── Event Bus
```

This approach provides:

- one daemon
- one service to supervise
- one log stream
- one installation path
- simple debugging
- clean internal boundaries
- future extraction paths if separate processes become necessary

---

# 13. Hardware Strategy

The project should support hardware through profiles and capabilities.

The core application should not depend on raw GPIO details.

Preferred flow:

```
Hardware Profile
      │
      ▼
Hardware Abstraction Layer
      │
      ▼
Capability Interface
      │
      ▼
Event Bus
      │
      ▼
Service
```

Initial hardware support focuses on the GPi Case 2 and Raspberry Pi CM4.

The design should remain aware of:

- GPi Case 2W
- NESPi-style cases
- SUPERPi-style cases
- Raspberry Pi 4
- Raspberry Pi 5
- Raspberry Pi Zero 2 W
- future Compute Modules
- community-supported SBCs

---

# 14. Software Strategy

Initial software stack:

- Raspberry Pi OS Bullseye 64-bit
- RetroPie
- EmulationStation
- RetroArch
- systemd
- journald
- libgpiod

The project should avoid hardcoding assumptions that prevent future adapters for:

- Batocera
- Recalbox
- EmulationStation Desktop Edition
- Pegasus
- other RetroArch-based systems

RetroPie remains the first target.

Adapters make future support possible.

---

# 15. Development Strategy

Day-to-day development should be possible without target hardware when practical.

Development hosts:

- macOS Apple Silicon
- macOS Intel
- Linux x64
- Linux ARM64
- Linux ARMv7
- Windows where practical for non-hardware code

Build targets:

- linux/amd64
- linux/arm64
- linux/arm/v7
- darwin/amd64
- darwin/arm64

The project should use:

- mocks
- simulated events
- interface boundaries
- unit tests
- CI builds
- cross-compilation
- hardware validation for final behavior

---

# 16. Documentation Strategy

Documentation should be treated as an engineering artifact.

Durable markdown documents should include metadata headers.

Major topics should eventually include:

- WHY
- Project Manifest
- Project Charter
- Engineering Manifesto
- Requirements
- Terminology
- Glossary
- Hardware reference
- Architecture
- Development handbook
- Testing strategy
- Terminal UI guide
- Performance research
- ADRs
- RFCs
- Risk register
- Open questions
- Assumption log

The repository should support both human contributors and AI assistants.

---

# 17. Safety and Reliability

The project must protect user data and system integrity.

Safety priorities:

1. Never make shutdown less safe than the original behavior.
2. Avoid corrupting save states.
3. Write state atomically.
4. Prefer recoverable failure modes.
5. Back up critical state where practical.
6. Provide clear recovery steps.
7. Avoid destructive actions without explicit intent.
8. Treat power loss as expected, not exceptional.

Resume features should be introduced cautiously.

A failed resume should fall back to the frontend rather than trapping the user.

---

# 18. Performance Targets

Initial performance targets:

- Power daemon idle CPU: below 0.1% where practical.
- Power daemon memory: below 10 MB RSS where practical.
- Power Switch ON to resumed gameplay: under 20 seconds.
- Stretch resume target: under 15 seconds.
- Shutdown response: fast enough to preserve user trust and data.
- CLI commands: responsive enough for SSH and troubleshooting.

Performance improvements should be measured.

Research should distinguish:

- actual boot speed
- perceived startup speed
- frontend launch time
- emulator launch time
- resume restore time

---

# 19. Constraints

Known constraints:

- GPi Case 2 hardware behavior may limit true sleep mode.
- Full suspend/hibernate may not be possible or reliable.
- RetroArch save-state behavior may vary by core.
- RetroPie and Raspberry Pi OS versions may affect available APIs.
- libgpiod behavior may vary across kernel versions.
- Battery data may not be available from all hardware.
- Boot optimization may conflict with hardware compatibility.
- Some behavior requires physical hardware validation.

---

# 20. Risks

Initial risks:

| Risk | Impact | Mitigation |
|---|---:|---|
| Resume corrupts save state | High | Atomic writes, backups, cautious rollout |
| Power loss during save | High | Save early, reduce write windows, fail safely |
| GPIO behavior differs by hardware revision | Medium | Hardware profiles, validation notes |
| libgpiod behavior differs by OS/kernel | Medium | Abstraction layer, compatibility testing |
| Boot optimization breaks reliability | Medium | Measure, document, make optional |
| Planning continues too long | Medium | One artifact per session, commit packets |
| Personality becomes gimmicky | Low | Style guide, clear error rules |
| Scope grows too quickly | High | Milestones, RFCs, non-goals |

---

# 21. Success Criteria

The project is successful when:

- Safe shutdown works reliably under systemd.
- GPIO polling is replaced with event-driven handling where possible.
- The project logs clearly through journald.
- The project can be built and tested away from Raspberry Pi hardware.
- The GPi Case 2 reference platform is documented and validated.
- Resume behavior is safe, understandable, and recoverable.
- Boot and resume performance can be measured.
- The CLI provides useful diagnostics and status.
- The repository explains itself to humans and AI assistants.
- Contributors can understand how to add hardware profiles.
- The terminal experience feels polished and enjoyable.
- Users feel like the handheld is closer to a console.

---

# 22. Initial Milestone Definition

## Milestone 0 — Dreaming

Purpose:

Capture the project's purpose, philosophy, language, constraints, and blueprint before production code begins.

Key artifacts:

- WHY.md
- PROJECT_MEMORY.md
- ENGINEERING_MANIFESTO.md
- PROJECT_MANIFEST.md
- PROJECT_CHARTER.md
- Terminology
- Glossary
- Requirements
- Roadmap
- Initial architecture
- Development workflow
- AI collaboration guide

Exit criteria:

- The project has a documented purpose.
- The project has a documented philosophy.
- The project has documented terminology.
- The project has documented scope and non-goals.
- The project has a first architecture direction.
- The project has a clear next implementation milestone.

## Milestone 1 — Awakening

Purpose:

Introduce the first running daemon.

Exit criteria:

- `retroflag-powerd` builds.
- The daemon starts.
- The daemon logs.
- The daemon handles signals.
- The daemon can run under systemd.
- The daemon can stop cleanly.

---

# 23. Governance

Major decisions should be captured in ADRs.

Major proposals should begin as RFCs.

Known uncertainties should be captured in an Open Questions document.

Assumptions should be captured in an Assumption Log.

Risks should be captured in a Risk Register.

The project should prefer visible decision-making over hidden tribal knowledge.

---

# 24. Charter Review

This charter should be reviewed when:

- Milestone 0 completes.
- The first daemon is implemented.
- Resume work begins.
- Hardware support expands.
- A major scope change is proposed.
- v1.0 planning begins.

The charter may evolve, but changes should be intentional.

---

# Closing

RetroFlag Power is a project about power, resume, performance, and handheld polish.

It is also a project about craft.

The project should be reliable enough to trust, clear enough to maintain, portable enough to grow, and warm enough to make someone smile.

The charter is simple:

Build the machine.

Protect the dream.

Never lose today's excitement in tomorrow's implementation.
