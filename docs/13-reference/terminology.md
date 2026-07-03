---
id: TERM-001
title: Terminology Guide
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Hardware Porters
purpose: Define the official vocabulary for RetroFlag Power so documentation, code, events, interfaces, and user-facing messages model the hardware and software consistently.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - ENGINEERING_MANIFESTO.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
last_updated: 2026-07-03
---

# Terminology Guide

> Good names reduce complexity before code is even written.

This guide defines the official language of RetroFlag Power.

Terminology is not decoration.

Terminology shapes architecture.

If the project names things well, the code becomes easier to read, the documentation becomes easier to understand, and contributors are less likely to build the wrong mental model.

The guiding rule is simple:

> Name things according to their real-world behavior, not just their implementation detail.

---

# 1. Why Terminology Matters

RetroFlag Power sits between physical hardware, Linux services, emulator software, and the player experience.

That means a single concept may exist at several layers:

```
Physical hardware
      │
      ▼
GPIO signal
      │
      ▼
Hardware capability
      │
      ▼
Application event
      │
      ▼
User-facing action
```

If the project uses vague names, these layers become confused.

For example, calling a latching hardware switch a `Power Button` may seem harmless, but it teaches the wrong model.

A button is momentary.

A switch has a maintained state.

That distinction affects:

- event names
- interface design
- hardware profiles
- tests
- documentation
- user-facing messages

The project should use language that mirrors reality.

---

# 2. Terminology Principles

## 2.1 Use real-world behavior first

Prefer names that describe how the user experiences the hardware.

Use:

```
Power Switch
```

when the physical control latches between ON and OFF.

Use:

```
Reset Button
```

when the physical control is momentary.

## 2.2 Separate state from event

A state persists.

An event happens.

Examples:

```
Power Switch is ON
```

is a state.

```
Power Switch moved to OFF
```

is an event.

## 2.3 Separate hardware from implementation

A user-facing or architecture-facing term should describe the capability, not the GPIO detail.

Use:

```
Power Switch capability
```

not:

```
GPIO 3 shutdown pin
```

unless the document is specifically discussing the GPIO implementation.

## 2.4 Use specific product names only when needed

Use generic terms when discussing architecture.

Use product-specific names when discussing actual hardware.

Example:

Use:

```
Frontend
```

for architecture.

Use:

```
EmulationStation
```

when discussing the RetroPie reference platform.

## 2.5 Keep errors boring and useful

Personality belongs in success, status, and discovery moments.

Errors must be clear, actionable, and searchable.

## 2.6 Avoid legally risky references

The project may be inspired by retro gaming culture, but official terminology and assets should remain original.

Do not rely on Nintendo, Sega, Valve, RetroPie, or other protected branding as part of the project's identity.

---

# 3. Hardware Terms

## Power Switch

A latching physical control with maintained ON and OFF positions.

Examples:

- RetroFlag GPi Case 2 side Power Switch
- NESPi-style case Power Switch

Use:

```
Power Switch
Power Switch ON
Power Switch OFF
Power Switch moved to OFF
Power Switch State
Power Switch Event
```

Avoid:

```
Power Button
Power button pressed
Power button released
```

unless the actual hardware uses a momentary power button.

## Reset Button

A momentary physical control that is pressed and released.

Examples:

- GPi Case 2 Reset Button
- NESPi Reset Button

Use:

```
Reset Button
Reset Button pressed
Reset Button released
Reset Button Event
```

Avoid:

```
Reset Switch
```

unless the physical hardware is actually latching.

## Button

A momentary hardware control.

A button does not maintain ON/OFF state after the user releases it.

Correct examples:

```
Reset Button
Menu Button
Hotkey Button
```

Incorrect example for GPi Case 2:

```
Power Button
```

## Switch

A hardware control with a maintained state.

Correct examples:

```
Power Switch
Mode Switch
```

## Backlight

The display illumination capability.

Use:

```
Backlight
Backlight Level
Backlight OFF
Backlight restored
```

Avoid describing this as a PWM channel unless discussing implementation details.

## Battery

The user-facing energy source or reported energy state.

Use:

```
Battery
Battery Level
Battery State
Critical Battery
```

Avoid:

```
ADC value
raw sensor value
```

unless discussing low-level hardware measurement.

## LED

A physical indicator light.

Use:

```
LED
Status LED
Power LED
Activity LED
```

When possible, describe what the LED communicates rather than only how it is wired.

## Display

The visible screen as experienced by the user.

Use:

```
Display
Screen
Resume Screen
Splash Screen
```

Use implementation terms such as framebuffer, DRM, KMS, or FKMS only in technical documents.

## Audio

The user-facing sound output capability.

Use:

```
Audio
Volume
Muted
```

Use ALSA, PulseAudio, or PipeWire only when referring to the actual software stack.

---

# 4. State and Event Terms

## State

A persistent condition at a point in time.

Examples:

```
Power State
Resume State
Session State
Battery State
Frontend State
```

States can be read, stored, compared, and restored.

## Event

A meaningful change or occurrence.

Examples:

```
Power Switch moved to OFF
Reset Button pressed
Game launched
Game exited
Resume started
Resume failed
Shutdown requested
```

Events may be produced by hardware, software, timers, or user commands.

## Capability

A feature exposed by a hardware profile, software adapter, or platform.

Examples:

```
Power Switch capability
Reset Button capability
Battery capability
Backlight capability
Resume capability
Frontend event capability
```

Capabilities allow the project to support different hardware without hardcoding model-specific assumptions into core services.

## Profile

A description of a hardware or software environment.

Examples:

```
GPi Case 2 profile
Raspberry Pi CM4 profile
RetroPie profile
EmulationStation adapter profile
```

Profiles map real-world platforms to capabilities.

## Adapter

A component that connects RetroFlag Power to an external system.

Examples:

```
RetroArch adapter
EmulationStation adapter
systemd adapter
GPIO adapter
```

Adapters isolate external assumptions from core services.

---

# 5. Software Terms

## Daemon

A long-running background process supervised by the operating system.

Primary future daemon:

```
retroflag-powerd
```

The daemon is responsible for running internal services such as power, resume, state, metrics, and hardware handling.

## Service

An internal responsibility within the daemon.

Examples:

```
Power Service
Resume Service
State Service
Metrics Service
Hardware Service
Frontend Service
Configuration Service
Terminal UI Service
```

Avoid creating separate executables until there is a clear architectural need.

## Power Service

The internal service responsible for reacting to power-related events and coordinating safe shutdown behavior.

The Power Service should think in terms of Power Switch events and shutdown requests, not raw GPIO values.

## Resume Service

The internal service responsible for recording, preparing, and restoring game sessions.

The Resume Service should think in terms of sessions, frontends, emulators, and state.

## State Service

The internal service responsible for durable state.

Examples of state:

- current session
- last known ROM
- emulator/core
- save-state slot
- resume timestamp
- display settings
- audio settings
- resume failure data

## Metrics Service

The internal service responsible for measuring and reporting timing and health data.

Examples:

- boot time
- resume time
- shutdown time
- daemon memory usage
- daemon CPU usage
- last resume success/failure

## Hardware Service

The internal service responsible for loading hardware profiles and exposing capabilities to other services.

## Event Bus

The internal mechanism for distributing events between services.

The Event Bus should help services remain decoupled.

It should not become an overcomplicated framework.

## Configuration Service

The internal service responsible for loading and validating configuration.

Configuration should not be added before it is needed.

Hardcoded defaults are acceptable in early milestones.

---

# 6. Gaming and Resume Terms

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
- aspect ratio
- integer scaling
- volume
- brightness
- timestamp

## Resume

The act of returning to a previously recorded session.

Use:

```
Resume
Resume started
Resume completed
Resume failed
Resume State
Resume Service
```

## Continue

A user-facing action that may trigger resume.

Use `Continue` for player-facing UI when it feels more natural.

Example:

```
Continue Last Game
```

Use `Resume` in architecture and implementation.

## Save State

An emulator-managed snapshot of game execution.

Use when referring specifically to RetroArch or emulator save-state files.

## Save Point

A more playful user-facing phrase.

Use sparingly in terminal UX or splash messaging.

Do not use `Save Point` when discussing exact emulator implementation.

## Checkpoint

A friendly user-facing concept for restored progress.

Use for success messages, progress reports, or milestone language.

Do not use as the technical name for save-state files.

## Frontend

The software interface used to browse systems and launch games.

Initial frontend:

```
EmulationStation
```

Use `Frontend` in architecture unless the document specifically discusses EmulationStation.

## Emulator

The software that runs the game.

RetroArch is the initial target emulator framework, but architecture should not assume every emulator is RetroArch.

## Core

A RetroArch-specific emulator module.

Use `core` only when specifically discussing RetroArch behavior.

---

# 7. Platform Terms

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

## Target Platform

A platform the project intends to support.

The reference platform is always a target platform, but not every target platform is the reference platform.

## Host Development Platform

A system used to develop the project.

Examples:

- macOS Apple Silicon
- macOS Intel
- Linux x64
- Linux ARM64
- Linux ARMv7
- Windows where practical

## Validation Hardware

Physical hardware required to confirm real behavior.

Mock tests can prove logic.

Validation hardware proves reality.

## SBC

Single-board computer.

Examples:

- Raspberry Pi
- Orange Pi
- ODROID
- Radxa ROCK
- Banana Pi
- Libre Computer

Use this term when discussing the broader hardware landscape.

---

# 8. User-Facing Tone

## Success

Success messages may be warm and playful.

Example:

```
Everything checks out.

Player 1 is ready.
```

## Warning

Warnings should be calm and clear.

Example:

```
Resume state exists, but the ROM file was not found.

The system will return to the frontend instead.
```

## Error

Errors should be direct and actionable.

Example:

```
ERROR

RetroArch executable not found.

Expected:
/usr/bin/retroarch

Next step:
Install RetroArch or update the configured emulator path.
```

Avoid jokes in errors.

## Diagnostics

Diagnostics should be readable and useful.

Example:

```
Power Switch      OK
Reset Button      OK
RetroArch         Found
Resume State      Ready
Last Resume       Successful
```

## Benchmarks

Benchmarks may feel celebratory.

Example:

```
Boot to Resume: 18.4s

Target: <20s
Result: Excellent
```

---

# 9. Terms to Prefer

| Prefer | Avoid | Reason |
|---|---|---|
| Power Switch | Power Button | GPi Case 2 uses a latching switch |
| Reset Button | Reset Switch | Reset is momentary |
| Power State | GPIO Status | State is user/domain level |
| Power Event | Pin Change | Event is architecture level |
| Resume | Continue | Resume is technical; Continue can be UI |
| Session | Game Info | Session captures full context |
| Frontend | EmulationStation | Frontend is portable architecture |
| Emulator | RetroArch | Emulator is generic; RetroArch is adapter-specific |
| Capability | Feature flag | Capability maps to hardware/platform behavior |
| Hardware Profile | Board config | Profile captures real device behavior |
| Backlight | PWM channel | Backlight is user-facing |
| Battery Level | ADC value | Battery level is user-facing |
| Metrics | Timings | Metrics covers broader measurement |
| Terminal UI | Console output | UI is intentionally designed |

---

# 10. Terms That Need Care

## Safe Shutdown

Use this when referring to preserving data and shutting down cleanly.

Do not use it to imply that all shutdown situations are risk-free.

## Sleep Mode

Use only for research unless true hardware sleep behavior is proven.

For partial behavior, prefer:

```
Sleep-like mode
Low-power pause
Pause mode
```

## Automatic Resume

Use carefully.

Automatic resume should always have safe fallback behavior.

## Fast Boot

Distinguish between:

- actual boot time
- perceived boot time
- time to frontend
- time to resumed gameplay

## Battery Percentage

Use only if hardware can provide real percentage data.

Otherwise use:

```
Battery status
Battery estimate
Battery research
```

---

# 11. Naming Patterns for Future Code

These are not final APIs, but they show the preferred naming model.

## Event names

Prefer:

```
PowerSwitchOn
PowerSwitchOff
PowerSwitchChanged
ResetButtonPressed
ResetButtonReleased
ResumeStarted
ResumeCompleted
ResumeFailed
ShutdownRequested
```

Avoid:

```
PowerButtonPressed
GPIOHigh
GPIOLow
PinChanged
```

unless working inside a low-level GPIO package.

## Interface names

Possible examples:

```go
type PowerSwitch interface {
    State() PowerState
    Events() <-chan PowerSwitchEvent
}
```

```go
type ResetButton interface {
    Events() <-chan ResetButtonEvent
}
```

```go
type HardwareProfile interface {
    Capabilities() Capabilities
}
```

## Package names

Prefer domain names:

```
power
resume
state
metrics
hardware
frontend
events
ui
```

Use implementation names only where appropriate:

```
gpio
retroarch
emulationstation
systemd
```

---

# 12. Documentation Language

Documentation should distinguish between:

## Facts

Verified information.

Example:

```
The current system starts SafeShutdown.py from /etc/rc.local.
```

## Decisions

Choices the project has made.

Example:

```
The project uses systemd instead of rc.local.
```

## Assumptions

Beliefs that may need future validation.

Example:

```
Future RetroFlag hardware may continue exposing Linux-accessible GPIO controls.
```

## Research

Open investigation areas.

Example:

```
Sleep mode requires additional hardware testing.
```

## Aspirations

Desired future outcomes.

Example:

```
The handheld should resume gameplay in under 20 seconds.
```

Keeping these separate helps the project stay honest.

---

# 13. Glossary Seeds

These terms should eventually be promoted into a full glossary.

- Adapter
- Backlight
- Battery
- Button
- Capability
- Checkpoint
- CLI
- Core
- Daemon
- EmulationStation
- Emulator
- Event
- Event Bus
- Frontend
- GPIO
- Hardware Profile
- Host Development Platform
- LED
- Metrics
- Power Service
- Power State
- Power Switch
- Reference Platform
- Reset Button
- Resume
- Resume Service
- RetroArch
- RetroPie
- Safe Shutdown
- Save Point
- Save State
- SBC
- Session
- State
- State Service
- Switch
- systemd
- Target Platform
- Terminal UI
- Validation Hardware

---

# 14. Review Rules

This terminology guide should be reviewed whenever:

- A new hardware profile is added.
- A new service is introduced.
- User-facing CLI language is added.
- An event type is created.
- A public interface is named.
- A contributor proposes terminology that conflicts with the guide.
- The project expands beyond the reference platform.

Terminology may evolve, but changes should be intentional.

---

# Closing

Language is part of the architecture.

RetroFlag Power should speak in a way that reflects the hardware, respects the user, and helps contributors understand the system before they read the code.

A latching switch is a Power Switch.

A momentary control is a Reset Button.

A restored game is a Resume.

A polished experience is the goal.

Name things well.

Build from there.
