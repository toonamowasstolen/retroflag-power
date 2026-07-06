---
id: GLOSSARY-001
title: Glossary
version: 0.2.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Hardware Porters
  - Curious Users
purpose: Provide a quick-reference dictionary for RetroFlag Power terminology so contributors can understand the project's shared language without reading the full terminology guide.
related:
  - docs/13-reference/terminology.md
  - PROJECT_MEMORY.md
  - PROJECT_CHARTER.md
  - PROJECT_MANIFEST.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
last_updated: 2026-07-06
---

# Glossary

> The dream has words. This is where we keep them.

This glossary is the quick-reference companion to the Terminology Guide.

Use this document when you need a short definition.

Use `docs/13-reference/terminology.md` when you need naming rules, examples, or guidance about which term to prefer.

---

# A

## Adapter

A component that connects RetroFlag Power to an external system while keeping external assumptions out of core services.

Examples:

- RetroArch adapter
- EmulationStation adapter
- systemd adapter
- GPIO adapter

## Adventure

Epoch 5 and a project voice term used to describe the player experience of
continuing a game.

Use in user-facing or celebratory language, not as a technical object.

## Awakening

Epoch 1.

The life stage where the first daemon begins to build, run, identify itself,
describe its lifecycle, and stop cleanly without controlling hardware.

## AI Assistant

A contributor-like tool that helps modify, explain, or generate project artifacts.

AI assistants should be supported with clear context, documentation, architecture notes, and coding standards.

## Architecture Decision Record

A durable document that records an important decision the project has made.

Common abbreviation:

```
ADR
```

Use ADRs for decisions that affect architecture, maintainability, portability, safety, or project direction.

## Assumption

A belief the project currently holds but has not fully proven.

Assumptions should be recorded separately from decisions.

---

# B

## Backlight

The display illumination capability as experienced by the user.

Use `Backlight` in user-facing and architecture documents.

Use lower-level terms such as PWM, framebuffer, DRM, KMS, or FKMS only when discussing implementation.

## Battery

The user-facing energy source or energy state.

Related terms:

- Battery Level
- Battery State
- Critical Battery

Do not use `Battery Percentage` unless the hardware can provide meaningful percentage data.

## Benchmark

A measured result that helps evaluate performance.

Examples:

- boot-to-resume time
- daemon memory usage
- daemon CPU usage
- shutdown response time

## Button

A momentary physical control.

A button is pressed and released.

Correct example:

```
Reset Button
```

Incorrect example for GPi Case 2:

```
Power Button
```

because the GPi Case 2 uses a latching Power Switch.

---

# C

## Capability

A feature exposed by hardware, software, or a platform.

Examples:

- Power Switch capability
- Reset Button capability
- Battery capability
- Backlight capability
- Resume capability

Capabilities allow the core application to avoid hardcoding assumptions about specific hardware.

## Charter

The formal project agreement that defines purpose, scope, non-goals, stakeholders, principles, risks, and success criteria.

Primary document:

```
PROJECT_CHARTER.md
```

## Checkpoint

A friendly or informal concept for restored or completed progress.

Use for terminal UX and status messages. A verified project checkpoint becomes
a numbered `Milestone` with evidence.

Do not use as the technical name for emulator save-state files.

## CLI

Command-line interface.

Potential future command:

```
retroflag-power
```

Potential daemon:

```
retroflag-powerd
```

## Code of Conduct

A document that defines expected community behavior.

The project should eventually include one before broader public contribution.

## Commit Packet

A coherent set of files that belong together in one commit.

Commit packets should be meaningful but not enormous.

## Configuration Service

The internal service responsible for loading and validating configuration.

Configuration should not be added before it is needed.

## Continue

A user-facing action that may trigger Resume.

Example:

```
Continue Last Game
```

Use `Resume` for architecture and implementation.

Use `Continue` when speaking to the player.

## Core

A RetroArch-specific emulator module.

Use `Core` only when specifically discussing RetroArch.

For generic architecture, prefer `Emulator`.

---

# D

## Daemon

A long-running background process supervised by the operating system.

Expected primary daemon:

```
retroflag-powerd
```

## Decision

A choice the project has intentionally made.

Major decisions should be captured in ADRs or decision logs.

## Delight

A small thoughtful touch that makes the project more enjoyable without reducing clarity.

Examples:

- a friendly success message
- clean ASCII output
- a readable benchmark report
- a polished status screen

Delight should support usefulness, not replace it.

## Diagnostics

Checks that help users and maintainers understand whether the system is healthy.

Possible future command:

```
retroflag-power doctor
```

## Dreaming

Epoch 0.

The stage where the project captures purpose, philosophy, terminology, scope, architecture direction, and first implementation path before production code begins.

---

# E

## Epoch

A large project life stage.

The canonical Epoch ladder is:

```text
Dreaming → Awakening → Heartbeat → Memory → Momentum → Adventure → Launch
```

The Roadmap records the current Epoch and future direction.

## EEPROM

Non-volatile firmware configuration storage on Raspberry Pi hardware.

Relevant to boot behavior and boot-order research.

## EmulationStation

The initial frontend on the RetroPie reference platform.

Use `EmulationStation` only when discussing that specific software.

Use `Frontend` for generic architecture.

## Emulator

Software that runs games.

RetroArch is the initial target emulator framework, but the architecture should not assume every emulator is RetroArch.

## Event

A meaningful change or occurrence.

Examples:

- Power Switch moved to OFF
- Reset Button pressed
- Game launched
- Resume failed
- Shutdown requested

## Event Bus

The internal mechanism for distributing events between services.

It should help services remain decoupled without becoming an overcomplicated framework.

---

# F

## Fact

Verified information.

Example:

```
The current system starts SafeShutdown.py from /etc/rc.local.
```

Facts should be distinguished from assumptions, decisions, research, and aspirations.

## FKMS

Fake Kernel Mode Setting.

A Raspberry Pi graphics stack option relevant to boot and display performance research.

Use only in technical performance or graphics documentation.

## Frontend

The software interface used to browse systems and launch games.

Initial reference frontend:

```
EmulationStation
```

Use `Frontend` in architecture unless discussing EmulationStation specifically.

---

# G

## Game Session

See `Session`.

## GPIO

General Purpose Input/Output.

Used for hardware signaling.

Core services should not depend directly on raw GPIO details.

GPIO behavior belongs behind hardware abstractions and adapters.

## GPi Case 2

The initial reference handheld hardware.

Full reference platform:

- RetroFlag GPi Case 2
- Raspberry Pi CM4 Rev 1.1
- 4 GB RAM
- Samsung EVO Select microSD
- RetroPie
- Raspberry Pi OS Bullseye 64-bit
- EmulationStation
- RetroArch

---

# H

## Hardware Abstraction Layer

A boundary that hides hardware-specific implementation details from core services.

Common abbreviation:

```
HAL
```

## Hardware Porter

A contributor who adds support for another case, board, SBC, or hardware profile.

## Hardware Profile

A description of a real hardware platform and the capabilities it provides.

Examples:

- GPi Case 2 profile
- NESPi-style case profile
- Raspberry Pi CM4 profile

## Hardware Service

The internal service responsible for loading hardware profiles and exposing capabilities to other services.

## Heartbeat

Epoch 2.

The stage where the daemon becomes a reliable supervised service with health, logs, and operational visibility.

## Host Development Platform

A system used to develop the project away from target hardware.

Examples:

- macOS Apple Silicon
- macOS Intel
- Linux x64
- Linux ARM64
- Windows where practical for non-hardware code

---

# I

## Implementation Detail

A low-level mechanism that should not leak into higher-level language unless the document is specifically technical.

Examples:

- GPIO line offset
- active-low signal
- PWM channel
- systemd unit setting
- RetroArch command syntax

## Interface

A defined boundary between components.

Interfaces should represent domain concepts and capabilities, not accidental implementation details.

---

# J

## journald

The logging system used by systemd.

RetroFlag Power should log clearly through journald when running as a service.

---

# K

## KMS

Kernel Mode Setting.

A graphics stack option relevant to boot, display, and performance research.

Use only in technical performance or graphics documentation.

---

# L

## Launch

Epoch 6 and a general product concept.

As an Epoch, `Launch` leads through release preparation to a stable public
release for the reference platform.

As a product concept, launching means moving from startup into playable state.

## LED

A physical indicator light.

Examples:

- Status LED
- Power LED
- Activity LED

Describe what an LED communicates when possible.

## libgpiod

The preferred modern Linux GPIO userspace interface direction for this project.

The project intends to use `libgpiod` rather than older Raspberry Pi specific GPIO libraries where practical.

## Linux

The operating system family underneath the reference platform.

Initial target:

```
Raspberry Pi OS Bullseye 64-bit
```

---

# M

## Manifest

A document that defines project identity, values, mission, and guiding promises.

Primary document:

```
PROJECT_MANIFEST.md
```

## Memory

Epoch 3.

The life stage where the project records durable session context and develops
the route toward restoring a previous game safely.

## Metrics

Measured operational or performance data.

Examples:

- boot time
- resume time
- daemon CPU
- daemon memory
- last resume status

## Metrics Service

The internal service responsible for collecting, storing, or reporting metrics.

## Milestone

A numbered, verified project checkpoint.

Primary document:

```
docs/00-project/milestones.md
```

Milestones use stable IDs such as `M-0001` and cite verification evidence.

## Modular Monolith

A single process with clear internal service boundaries.

RetroFlag Power should begin this way to avoid unnecessary IPC and process complexity.

## Momentum

Epoch 4 and the force created by small completed victories.

Momentum is protected by committing coherent artifacts and avoiding endless planning.

---

# N

## Non-Goal

Something the project intentionally does not aim to do.

Non-goals protect scope.

Examples:

- not a full operating system
- not a replacement for RetroPie
- not a cloud service
- not a GUI dashboard in early Epochs

---

# O

## Open Question

A known uncertainty that needs research, testing, or future decision-making.

Open questions should be recorded so uncertainty does not become hidden assumption.

---

# P

## Performance Target

A measurable goal.

Initial examples:

- Power Switch ON to resumed gameplay under 20 seconds.
- Stretch target under 15 seconds.
- Power daemon idle CPU below 0.1 percent where practical.

## Persona

A representative type of user or contributor.

Examples:

- Player
- Power User
- Developer
- Hardware Porter
- Maintainer
- AI Assistant

## Player

The person using the handheld to play games.

The player should experience simplicity, reliability, and fast return to play.

## Polish

An Adventure Epoch route theme.

The stage where terminal UX, diagnostics, splash presentation, and documentation quality become intentionally crafted.

## Portability

The ability for the project to build, test, or eventually run across more than one platform.

Portability is a feature.

## Power Service

The internal service responsible for reacting to power-related events and coordinating safe shutdown behavior.

## Power State

The persistent state of the Power Switch or power-related domain.

Examples:

```
Power Switch ON
Power Switch OFF
```

## Power Switch

A latching physical control with maintained ON and OFF positions.

The GPi Case 2 has a Power Switch, not a Power Button.

## Power User

A user who tweaks settings, reads logs, runs diagnostics, and wants control.

## Product Vision

A document or idea describing the desired user experience and long-term product direction.

## Project Memory

The project's origin, principles, safety net, important unresolved ideas, and
durable context.

Primary document:

```text
PROJECT_MEMORY.md
```

Project Memory is not the current progress log or verified Milestone ledger.

## Profile

See `Hardware Profile`.

---

# Q

## Quest

A task or work record with focused scope, guardrails, acceptance criteria,
outcome, and validation.

Canonical location:

```text
docs/00-project/quests/
```

---

# R

## Raspberry Pi

The initial SBC family focus for the project.

Reference hardware uses a Raspberry Pi CM4.

## Reference Platform

The first-class platform used for design, validation, and documentation.

Current reference platform:

```
RetroFlag GPi Case 2
Raspberry Pi CM4 Rev 1.1
4 GB RAM
Samsung EVO Select microSD
RetroPie
Raspberry Pi OS Bullseye 64-bit
EmulationStation
RetroArch
```

## Release

A packaged, versioned project delivery.

Releases should include documentation, changelog, artifacts, limitations, and validation notes.

## Request for Comments

A proposal document used before a major decision is made.

Common abbreviation:

```
RFC
```

## Requirement

A traceable statement describing something the project must do or satisfy.

Requirements should eventually use stable IDs such as:

```
REQ-0001
```

## Reset Button

A momentary physical control that is pressed and released.

Use this term for the GPi Case 2 reset control.

## Resume

The act of returning to a previously recorded game session.

Use `Resume` for architecture and implementation.

Use `Continue` for player-facing actions when appropriate.

## Resume Service

The internal service responsible for recording, preparing, and restoring game sessions.

## RetroArch

The initial emulator framework target.

Architecture should use adapters so core services do not depend directly on RetroArch behavior.

## RetroPie

The initial retro gaming software distribution target.

The project should support RetroPie first, while avoiding unnecessary assumptions that block future adapters.

## Risk

A known possible problem that could affect the project.

Risks should be recorded and mitigated.

## Roadmap

A record of current project state and future direction.

The Roadmap owns the current Epoch, near-term route, upcoming gates, route
themes, and deferred scope. It cites verified Milestones rather than duplicating
their evidence.

---

# S

## Safe Shutdown

A shutdown flow designed to preserve user data and avoid filesystem or save-state corruption.

Safe shutdown is the first critical responsibility of the project.

## Save Point

A friendly user-facing phrase for a place where progress is preserved.

Do not use as the technical name for emulator save-state files.

## Save State

An emulator-managed snapshot of game execution.

Use when referring to actual emulator save-state behavior or files.

## SBC

Single-board computer.

Examples:

- Raspberry Pi
- Orange Pi
- ODROID
- Radxa ROCK
- Banana Pi
- Libre Computer

## Session

The complete set of information needed to understand or restore what the player was doing.

A session may include:

- system
- ROM path
- emulator
- core
- save-state slot
- shader
- overlay
- controller profile
- display settings
- audio settings
- timestamp

## Sleep Mode

A research topic unless true hardware sleep behavior is proven.

Prefer `sleep-like mode` or `low-power pause` when describing partial behavior.

## Splash Screen

A polished screen shown during startup or resume.

Possible future use:

```
Resuming...
Super Mario World
Launching...
```

## State

A persistent condition at a point in time.

Examples:

- Power State
- Resume State
- Session State
- Battery State

## State Service

The internal service responsible for durable state and state validation.

## Switch

A physical control with maintained state.

Correct example:

```
Power Switch
```

## systemd

The Linux service manager used to run and supervise the daemon.

RetroFlag Power should use systemd instead of rc.local.

---

# T

## Target Platform

A platform the project intends to support.

The reference platform is a target platform, but not all target platforms are reference platforms.

## Terminal UI

The designed command-line user experience.

It includes:

- layout
- color
- banners
- status tables
- diagnostics
- errors
- benchmark output
- accessibility behavior

## Troubleshooting

Documentation and commands that help users recover from problems.

Troubleshooting should be direct, searchable, and practical.

---

# U

## User Experience

The total experience of using the system.

In this project, user experience includes:

- power behavior
- startup time
- resume behavior
- terminal commands
- logs
- documentation
- failure recovery
- emotional polish

Common abbreviation:

```
UX
```

---

# V

## Validation Hardware

Physical hardware used to confirm real behavior.

Mock tests can validate logic.

Validation hardware validates reality.

## Victory

A small completed outcome that moves the project forward.

Every work session should end with a victory.

---

# W

## Watchdog

A service health mechanism that can restart or recover a daemon if it becomes unhealthy.

Watchdog support is a future enhancement after the service foundation is stable.

## WHY.md

The root document that explains why the project exists and preserves the original spark.

---

# X

No project terms currently defined.

---

# Y

No project terms currently defined.

---

# Z

No project terms currently defined.

---

# Closing

This glossary should grow as the project grows.

When a new term appears in code, documentation, events, interfaces, or user-facing output, decide whether it belongs here.

If the word shapes how people understand the project, define it.

The dream has words.

Keep them clear.
