---
id: MEMORY-001
title: Project Memory and Blueprint Capture
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - AI Assistants
  - Future Contributors
purpose: Preserve the full project intent, decisions, philosophy, terminology, and documentation plan so that no important idea is lost while artifacts are created one file at a time.
related:
  - WHY.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - ENGINEERING_MANIFESTO.md
  - docs/13-reference/terminology.md
last_updated: 2026-07-03
---

# Project Memory and Blueprint Capture

> This document exists to protect the dream.

The project is intentionally being built one artifact at a time to preserve momentum, avoid planning paralysis, and make every work session end with something commit-worthy.

However, the project is larger than any single document.

This file captures the full working memory of the project so that important ideas, decisions, terminology, emotional intent, hardware context, architectural direction, and future plans are not lost while the repository grows.

This is a living document.

It should be updated whenever a meaningful idea is discovered that has not yet been promoted into a dedicated artifact.

---

# 1. Project Identity

## Working project name

`retroflag-power`

The name began as a practical description of replacing the RetroFlag SafeShutdown script, but it also carries a nostalgic connection to Nintendo Power magazine and the NES generation.

The name should be allowed to grow into a broader project identity.

## Broader project concept

This is no longer only a SafeShutdown.py replacement.

It is the beginning of a modern handheld experience layer for RetroFlag and Raspberry Pi based retro gaming handhelds.

The core mission is:

> Make a Raspberry Pi handheld feel magical.

The project aims to transform the stock RetroPie GPi Case 2 experience into something that feels closer to a purpose-built consumer handheld like a Nintendo Switch, Steam Deck, Game Boy, or dedicated retro console.

## Emotional goal

The project should make users, maintainers, and contributors feel that someone cared.

It should be technically solid, but also warm, enjoyable, polished, and memorable.

---

# 2. Company and Personal Philosophy

## Taft Consulting slogan

> Never lose today's excitement in tomorrow's implementation.

This phrase should be treated as both a company slogan and a project promise.

It captures the core tension of the project:

- The excitement of discovery.
- The discipline of implementation.
- The danger of losing joy during technical execution.
- The need to protect the original spark.

## Personal engineering manifesto

A personal document should be created:

`ENGINEERING_MANIFESTO.md`

This document belongs to Joshua as much as to any specific repository.

It should be reusable at the start of future projects.

The manifesto should capture values such as:

- Build experiences, not just software.
- Documentation is part of the product.
- Design before implementation.
- Architecture exists to enable implementation, not delay it.
- Respect future contributors.
- Name things according to real-world behavior.
- Make complexity invisible.
- Leave moments of delight.
- Measure before optimizing.
- Portability is a feature.
- Curiosity is a strength.
- Craftsmanship matters.
- Every work session should end with a victory.
- Never lose today's excitement in tomorrow's implementation.

---

# 3. Core Philosophy

## The project is about craft

The project should be built with craftsmanship.

The goal is not only that it works.

The goal is that it feels thoughtfully made.

People should be able to open the repository and think:

> Whoever built this really cared.

## Software should have personality

Interfaces should feel alive enough to be enjoyable, but not gimmicky.

The tone should be:

- Warm
- Thoughtful
- Clear
- Helpful
- Occasionally playful
- Never confusing
- Never annoying
- Never cute at the expense of function

## Leave moments of delight

Small touches matter.

Examples:

- A friendly terminal banner.
- A clean ASCII diagram.
- A satisfying success message.
- A beautiful benchmark report.
- A thoughtful comment explaining why.
- A status screen that feels like a retro handheld is checking in.
- A commit message that tells a story.

## Fun is allowed

A command-line project can be professional and still be fun.

This is especially appropriate because the project serves retro gaming handhelds.

Users spend a lot of time with their machines.

It should be a fun experience.

---

# 4. Planning Discipline

## Planning must produce artifacts

The project should avoid endless planning.

Every planning session must produce something commit-worthy:

- A markdown document.
- A diagram.
- A glossary entry.
- An ADR.
- An RFC.
- A test.
- A prototype.
- A code change.

## One file at a time, but not one idea at a time

The repository should grow one artifact at a time to create momentum.

However, all larger ideas must be captured in this memory file until they are promoted into dedicated documents.

## Version 0.1 is acceptable

Documents should begin as version `0.1.0`.

Perfect is not required.

Momentum matters.

The project should prefer:

> Commit a good first version today.

over:

> Discuss a perfect version forever.

---

# 5. Narrative Milestones

The project should use milestone names that tell a story.

These names should remain meaningful engineering phases, not empty branding.

## Proposed milestone story

### Milestone 0 — Dreaming

The project has no production code yet.

It is discovering its purpose, values, architecture, and identity.

Primary artifacts:

- WHY.md
- PROJECT_MANIFEST.md
- PROJECT_CHARTER.md
- ENGINEERING_MANIFESTO.md
- Requirements
- Terminology
- Architecture drafts

### Milestone 1 — Awakening

The first executable compiles.

The project begins to run.

Primary artifacts:

- First daemon skeleton
- Structured logging
- Graceful signal handling
- systemd service

#### Nameplate checkpoint

Completed during Awakening:

- `retroflag-powerd --version` prints `retroflag-powerd 0.1.0-dev`.
- Normal startup logs include the daemon name and version.
- `make version` runs the daemon version command through Workshop.
- The VS Code task `Workshop: version` runs `make version`.
- `make check` runs tests, build, and version validation.
- GitHub Actions Forge inherits the name check because CI runs `make check`.
- Tests, build, version output, and the Ctrl+C runtime lifecycle passed.
- `make check` passed.

This checkpoint adds identity only. It does not add hardware, shutdown, resume,
or state behavior.

#### Config satchel checkpoint

The daemon now has a minimal internal configuration boundary with safe defaults:

- `AppName` defaults to `retroflag-powerd`.
- `Version` defaults to `0.1.0-dev` from the version package.
- `DryRun` defaults to `true`.
- The app receives the config and logs name, version, and `dry_run=true` at
  startup.
- `--version` output remains unchanged, the lifecycle remains clean, and
  `make check` passed.

The satchel is defaults-only. It does not load config files or environment
variables, and it adds no new CLI flags, GPIO, shutdown execution, service
activation, resume, or state storage.

#### Event charms checkpoint

The daemon now has a minimal internal lifecycle event model:

- `Event` contains only `Type` and `Message`.
- Four lifecycle types describe daemon starting, daemon ready, shutdown signal
  received, and daemon stopped.
- The app logs its lifecycle through these events.
- `--version` remains unchanged, startup still includes name, version, and
  `dry_run=true`, the Ctrl+C lifecycle exits cleanly, and `make check` passed.

These charms are synchronous descriptions only. No event bus, channels, async
processing, persistence, or third-party dependencies were added. They introduce
no GPIO, shutdown execution, service activation, resume, or state storage.

#### Dry-run action charm checkpoint

The daemon now has a standalone internal action model:

- `Action` contains `Type`, `Message`, and `DryRun`.
- `TypeNoop` represents a no-operation action.
- `NewDryRunNoop` creates a noop action with `DryRun: true`.
- No execution path or lifecycle wiring was added.
- `--version` remains unchanged and `make check` passed.

The charm cannot run anything. It adds no GPIO, shutdown execution, command
runner, shell execution, action queue, channels, async processing, persistence,
packaging changes, service activation, resume, or state storage.

### Milestone 2 — Heartbeat

The daemon becomes a real supervised service.

Primary artifacts:

- systemd integration
- journal logging
- health checks
- restart behavior
- first status output

### Milestone 3 — Memory

Resume and state management begin.

Primary artifacts:

- session state tracking
- current game recording
- save-state integration
- restore flow

### Milestone 4 — Momentum

Performance and boot experience become measurable.

Primary artifacts:

- boot profiling
- resume timing
- benchmark CLI
- optimization research

### Milestone 5 — Adventure

The project becomes usable by others.

Primary artifacts:

- installer
- documentation polish
- release artifacts
- hardware validation
- public beta

### Milestone 6 — Launch

The project reaches a stable public release.

Primary artifacts:

- v1.0 release
- tested hardware profiles
- complete documentation
- upgrade path
- troubleshooting guide

---

# 6. Commit Philosophy

Commit history should be technically useful and narratively meaningful.

A commit title may have personality, but the body should clearly explain the technical change.

## Example commit titles

```text
Begin dreaming.
```

```text
Give shape to the dream.
```

```text
The project takes its first breath.
```

```text
Teach the handheld to remember.
```

```text
Measure the path from power to play.
```

## Commit body expectations

Every non-trivial commit should explain:

- What changed.
- Why it changed.
- What documents, requirements, ADRs, or RFCs it relates to.
- Whether it affects behavior, architecture, tests, docs, or tooling.

---

# 7. Documentation Metadata Standard

Every durable markdown artifact should include front matter.

## Required fields

```yaml
---
id: DOC-000
title: Document Title
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
purpose: Describe why this document exists.
related:
  - OTHER_DOC.md
last_updated: 2026-07-03
---
```

## Status values

- Draft
- Review
- Approved
- Superseded
- Archived

## Why metadata matters

Metadata allows future tooling to:

- Build indexes.
- Track relationships.
- Generate documentation sites.
- Keep AI assistants oriented.
- Prevent important documents from becoming disconnected.

---

# 8. Documentation Structure

The repository should eventually contain both human-facing and AI-facing documentation.

## Root-level documents

- README.md
- WHY.md
- PROJECT_MANIFEST.md
- PROJECT_CHARTER.md
- ENGINEERING_MANIFESTO.md
- LICENSE
- CHANGELOG.md
- CONTRIBUTING.md
- CODE_OF_CONDUCT.md

## docs/

Human-facing engineering documentation.

Proposed structure:

```text
docs/
  00-project/
  01-product/
  02-hardware/
  03-software/
  04-architecture/
  05-development/
  06-testing/
  07-design/
  08-performance/
  09-research/
  10-decisions/
  11-rfc/
  12-api/
  13-reference/
  99-archive/
```

## ai/

AI-assistant guidance and condensed project intent.

Proposed structure:

```text
ai/
  00_PROJECT_CONSTITUTION.md
  01_PROJECT_OVERVIEW.md
  02_TARGET_HARDWARE.md
  03_SYSTEM_OVERVIEW.md
  04_ARCHITECTURE.md
  05_CODING_STANDARDS.md
  06_COMMIT_GUIDELINES.md
  07_REQUIREMENTS.md
  08_MILESTONES.md
  09_ROADMAP.md
  10_DECISIONS.md
  11_OPEN_QUESTIONS.md
  12_RISKS.md
  13_TEST_STRATEGY.md
  14_UI_GUIDE.md
  15_PRODUCT_PHILOSOPHY.md
  PROMPTS.md
```

---

# 9. Required Documentation Artifacts

The following documents should eventually exist.

## Soul documents

- WHY.md
- PROJECT_MANIFEST.md
- PROJECT_CHARTER.md
- PRODUCT_PHILOSOPHY.md
- ENGINEERING_MANIFESTO.md

## Product documents

- Vision
- Goals
- Non-goals
- Personas
- User stories
- Success metrics
- UX requirements

## Requirements documents

- Functional requirements
- Non-functional requirements
- Hardware requirements
- Performance requirements
- UX requirements
- Testing requirements
- Security and safety requirements

## Architecture documents

- System architecture
- Service architecture
- Power architecture
- Resume architecture
- State architecture
- Event bus architecture
- Hardware abstraction
- Filesystem layout
- Configuration strategy
- Logging strategy
- CLI architecture
- UI architecture

## Development documents

- Git workflow
- Commit guidelines
- Branch strategy
- Pull request process
- Coding standards
- Testing guide
- CI/CD guide
- Release guide
- AI collaboration guide

## Reference documents

- Glossary
- Terminology guide
- Hardware profiles reference
- Diagram style guide
- Terminal UI style guide

## Governance documents

- ADR process
- RFC process
- Decision log
- Assumption log
- Risk register
- Open questions

---

# 10. Requirements Tracking

Requirements should have stable IDs.

Example format:

```text
REQ-0001
System shall safely shut down when the Power Switch moves to OFF.
```

Types of requirements:

- Functional
- Non-functional
- Performance
- Hardware
- UX
- Developer experience
- Reliability
- Portability
- Safety

Requirements should be traceable to:

- Architecture
- Tests
- ADRs
- RFCs
- Releases

---

# 11. Decisions, Assumptions, Risks, and Questions

## ADRs

Architecture Decision Records capture decisions already made.

Examples:

- ADR-0001 Use systemd instead of rc.local.
- ADR-0002 Use Go for the main daemon.
- ADR-0003 Use libgpiod instead of RPi.GPIO.
- ADR-0004 Use an event-driven architecture.
- ADR-0005 Use hardware profiles and capabilities.

## RFCs

Requests for Comments capture proposals before they become decisions.

Examples:

- RFC-0001 Resume Manager.
- RFC-0002 Splash Resume Screen.
- RFC-0003 Sleep Mode.
- RFC-0004 Battery Overlay.
- RFC-0005 Companion CLI.

## Assumption log

Assumptions are not decisions.

Examples:

- Future RetroFlag hardware may continue using Linux GPIO.
- RetroPie remains the initial target software stack.
- RetroArch save-state behavior can be controlled reliably enough for resume.
- Day-to-day development should work without Pi hardware.

## Risk register

Examples:

- Resume corrupts save state.
- Power loss occurs during save.
- libgpiod behavior differs across kernels.
- Hardware revisions change GPIO behavior.
- Boot optimization breaks device compatibility.
- Planning continues too long without implementation.

## Open questions

Examples:

- Can room transitions be detected reliably for autosave?
- Can sleep-like behavior work without full hardware suspend?
- Can battery percentage be read from the RetroFlag board?
- How reliable is RetroArch remote control for pausing and saving?
- What is the safest restore behavior when a ROM no longer exists?
- How much boot time can EEPROM changes actually save?

---

# 12. Hardware Context

## Primary reference hardware

Device:

- RetroFlag GPi Case 2

Compute module:

- Raspberry Pi CM4 Rev 1.1
- 4 GB RAM

Storage:

- Samsung EVO Select microSD

Operating system:

- RetroPie
- Raspberry Pi OS Bullseye 64-bit

Current EEPROM:

- 2022-04-26

## Current boot optimization status

Completed:

- Bluetooth disabled
- Samba disabled
- ModemManager disabled
- Avahi disabled
- Swap disabled
- Quiet kernel boot

Considering:

- `BOOT_ORDER=0xf1`

instead of:

- `BOOT_ORDER=0xf25641`

## RetroFlag hardware awareness

The project should focus first on the GPi Case 2, but remain aware of the broader RetroFlag product line.

Future or related targets may include:

- GPi Case 2W
- NESPi cases
- SUPERPi cases
- Future RetroFlag handhelds and console-style cases

## Raspberry Pi focus

Initial support should prioritize Raspberry Pi hardware.

Likely targets:

- Raspberry Pi CM4
- Raspberry Pi 4
- Raspberry Pi 5
- Raspberry Pi Zero 2 W
- Future Compute Module boards

## Other SBC awareness

The architecture should not unnecessarily prevent support for other popular retro gaming SBCs.

Possible future community targets:

- Orange Pi
- Radxa ROCK boards
- ODROID boards
- Banana Pi
- Libre Computer boards
- x64 retro gaming systems
- Steam Deck-like systems for non-GPIO resume features

---

# 13. Terminology Principles

## Name things according to real-world behavior

Terminology should reflect the physical object and user mental model, not the implementation detail.

## Power Switch vs Power Button

The GPi Case 2 and NESPi-style cases use latching switches.

Therefore:

Use:

- Power Switch

Avoid:

- Power Button

unless the hardware actually uses a momentary power button.

## Reset Button

The reset input is momentary.

Therefore:

Use:

- Reset Button
- Reset Button Pressed
- Reset Button Released

## Power terminology

Use:

- Power Switch
- Power Switch ON
- Power Switch OFF
- Power State
- Power Event

Avoid:

- GPIO status
- raw pin value
- power button press

unless specifically discussing GPIO implementation details.

## User-facing language

Prefer gaming-friendly terms where appropriate:

- Resume
- Save Point
- Checkpoint
- Adventure
- Ready
- Continue

But keep technical errors clear and searchable.

---

# 14. Architecture Direction

## Modular monolith first

The initial implementation should likely be one daemon with multiple internal services rather than many independent executables.

Benefits:

- One systemd service.
- One journal stream.
- Simpler install.
- Simpler debugging.
- Shared in-memory state.
- Avoid premature IPC complexity.

Potential internal services:

- Power Service
- Resume Service
- State Service
- Metrics Service
- Hardware Service
- Configuration Service
- Frontend Service
- UI Service
- Event Bus

## Event-driven design

The project should avoid busy polling when event mechanisms are available.

Preferred flow:

```text
Hardware Event
      │
      ▼
Hardware Abstraction Layer
      │
      ▼
Event Bus
      │
      ▼
Service Handler
      │
      ▼
Action
```

## Hardware capabilities

The application should think in capabilities, not board models.

Examples:

- Power Switch capability
- Reset Button capability
- Battery capability
- Backlight capability
- LED capability
- Audio capability

Hardware profiles map real hardware to these capabilities.

## GPIO isolation

GPIO code must not leak into application logic.

The power service should not know about chip numbers, line offsets, pull-ups, falling edges, or active-low behavior.

Those belong in the hardware/GPIO layer.

---

# 15. Proposed Go Package Direction

Possible future layout:

```text
cmd/
  retroflag-powerd/

internal/
  app/
  power/
  resume/
  state/
  metrics/
  hardware/
  gpio/
  frontend/
  retroarch/
  emulationstation/
  config/
  events/
  logging/
  ui/
  systemd/
```

The exact layout can evolve, but the intent is:

- Keep packages small.
- Keep hardware-specific logic isolated.
- Keep core services testable without Pi hardware.
- Keep user-facing UI reusable.
- Avoid global state.
- Use context-aware service lifecycles.

---

# 16. Current Shutdown System

Current implementation:

```text
/opt/RetroFlag/SafeShutdown.py
```

Started by:

```text
/etc/rc.local
```

Current startup line:

```text
sudo python /opt/RetroFlag/SafeShutdown.py &
```

Problems:

- rc.local
- root Python process
- polling loop
- no structured logging
- no restart policy
- no dependency management
- difficult to extend
- difficult to debug

---

# 17. Project 1 — Power Service Modernization

Original goal:

Replace SafeShutdown.py with a systemd service.

Proposed service:

```ini
[Unit]
Description=RetroFlag Power Manager
After=multi-user.target

[Service]
ExecStart=/usr/local/bin/retroflag-power
Restart=always
User=root

[Install]
WantedBy=multi-user.target
```

Refined direction:

- Use systemd.
- Prefer `Restart=on-failure` unless there is a reason for `always`.
- Use journald logging.
- Support graceful signal handling.
- Avoid rc.local.
- Avoid Python for hardware daemon.
- Event-driven where possible.
- Add watchdog support later.

Benefits:

- automatic restart
- logs in journalctl
- startup ordering
- dependency management
- watchdog support
- easier debugging

---

# 18. Project 2 — SafeShutdown Rewrite

Current script polls GPIO continuously.

Rewrite goals:

- Use `libgpiod` instead of older `RPi.GPIO`.
- Avoid busy polling.
- Use GPIO events/edge detection.
- Preserve safe shutdown behavior.
- Keep power-facing code small and reliable.

Preferred event flow:

```text
GPIO Interrupt
      │
      ▼
Power Switch Event
      │
      ▼
Power Service
      │
      ▼
Shutdown Manager
```

---

# 19. Project 3 — Resume Manager

The Resume Manager is a major feature.

Possible daemon/service name:

- Internal service: Resume Service
- CLI command: resume-related commands under the main CLI
- Avoid prematurely creating separate executables unless needed

State file:

```text
/var/lib/retroresume/state.json
```

Example state:

```json
{
  "system": "nes",
  "rom": "/home/pi/RetroPie/roms/nes/Mario.nes",
  "core": "nestopia",
  "slot": 0,
  "timestamp": "2026-07-02T18:14:00"
}
```

Before shutdown:

- pause RetroArch
- save state
- save current ROM
- save emulator/core
- save shader if needed
- save controller profile
- initiate shutdown

After boot:

```text
Linux boots
      │
      ▼
EmulationStation starts
      │
      ▼
Resume Service waits
      │
      ▼
Previous ROM launches
      │
      ▼
RetroArch loads save state
      │
      ▼
Game resumes
```

User experience:

```text
Power ON
   │
   ▼
15–20 seconds
   │
   ▼
Exactly where they left off
```

---

# 20. Project 4 — Splash Resume Screen

Goal:

Create a polished resume splash before normal frontend interaction.

Example:

```text
GPi

Resuming...

Super Mario World

Launching...
```

Purpose:

- Improve perceived startup.
- Make the experience feel console-like.
- Hide technical boot steps.
- Provide feedback during resume.

---

# 21. Project 5 — Background Save States

Goal:

Reduce lost progress through periodic autosaves.

Possible triggers:

- Every 60 seconds.
- Before shutdown.
- Game exit.
- Room transition if supported.
- Emulator/core-specific hooks if available.

Risks:

- Save-state corruption.
- Performance impact.
- Unsupported cores.
- Interference with user save states.

Mitigations:

- Use dedicated autosave slots.
- Atomic writes.
- Backups.
- Per-core compatibility matrix.
- Disable by default until proven safe.

---

# 22. Project 6 — State Manager

State Manager should eventually remember more than save states.

Possible state:

- volume
- brightness
- shader
- overlay
- integer scaling
- aspect ratio
- controller profile
- emulator/core
- ROM path
- save-state slot
- frontend state
- last successful resume
- last failed resume

---

# 23. Project 7 — Frontend Integration

Initial target:

- EmulationStation on RetroPie

Architecture should avoid hardcoding frontend assumptions.

Use adapters where possible.

Possible event flow:

```text
Game Launch
      │
      ▼
Frontend Adapter
      │
      ▼
Resume Service records current session
```

On exit:

```text
Game Exit
      │
      ▼
Resume Service clears pending resume
      │
      ▼
Return to menu
```

Possible future frontends:

- EmulationStation Desktop Edition
- Batocera frontend environment
- Recalbox frontend environment
- Pegasus
- Other emulator launchers

---

# 24. Project 8 — Sleep Mode Research

Sleep mode may be limited or impossible with complete power removal.

Research goal:

Determine whether a useful sleep-like mode can be implemented.

Possible behavior:

- backlight off
- audio muted
- CPU governor powersave
- pause RetroArch
- disable nonessential services
- preserve session in memory

Wake behavior:

- restore clocks
- backlight on
- audio unmuted
- continue game

This should remain research until hardware behavior is understood.

---

# 25. Project 9 — Performance

Goal:

Power switch to resumed gameplay in less than 20 seconds.

Stretch goal:

Less than 15 seconds.

Research areas:

- EEPROM boot order
- bootloader timing
- systemd critical chain
- RetroPie startup
- EmulationStation startup
- RetroArch launch time
- KMS vs FKMS
- Mesa updates
- firmware updates
- kernel updates
- service disabling
- splash timing
- perceived boot improvements

Measure before optimizing.

---

# 26. Project 10 — Companion CLI

Potential commands:

```text
retroflag-power status
retroflag-power doctor
retroflag-power save
retroflag-power resume
retroflag-power clear
retroflag-power disable
retroflag-power benchmark
retroflag-power profile
retroflag-power version
```

The CLI should be useful, scriptable, and enjoyable.

It should support:

- plain output
- colorful terminal output
- possible JSON output
- no-color mode
- non-interactive mode

---

# 27. Nice Future Ideas

## Future EDC project expansion

These are later ideas and are not part of Milestone 1 — Awakening.

- [ ] Create `docs/00-project/edc-project-structure.md` to generalize the
  RetroFlag Power documentation and project structure into a reusable EDC
  structure guide that older projects can adopt.
- [ ] Create or extract an `edc-project-template/` starter for future and legacy
  projects. It should include root documents, metadata conventions, folder
  structure, ADR and RFC templates, project memory, roadmap, requirements, and
  AI collaboration guidance.

Do not create the guide or template until this work is deliberately promoted
into future scope.

## Product and experience ideas

- Battery percentage overlay if available.
- Graceful shutdown countdown when battery is critically low.
- Multiple resume slots.
- Resume slot per emulator/core.
- “Continue Last Game” tile in EmulationStation.
- Fast startup metrics.
- Historical boot timing tracking.
- Hardware profile wizard.
- Diagnostics/doctor command.
- First-run setup wizard.
- ASCII terminal welcome screen.
- Optional pixel-inspired branding.
- Original project mascot/companion.
- Terminal-friendly status dashboard.

---

# 28. Design and Terminal UX

## Personality

The project should feel retro, warm, and console-like without copying protected branding.

Avoid direct Nintendo, Mario, Zelda, Power Glove, or other trademarked/copyrighted material in official project assets.

Inspiration is acceptable.

Imitation is not.

## Visual style

Preferred:

- ASCII art
- ANSI color where available
- terminal-friendly layouts
- status cards
- progress bars
- playful but clear messages
- original pixel-inspired elements

## Accessibility

Terminal output should degrade gracefully:

- Works without color.
- Works in narrow terminals.
- Works over SSH.
- Works in logs.
- Can be parsed by scripts when needed.
- Avoids essential information being color-only.

## Tone

Success messages can be playful.

Errors must be clear, actionable, and searchable.

Good:

```text
Everything checks out.

Player 1 is ready.
```

Bad:

```text
Oopsie! The goblin stole your GPIO!
```

## Possible style vocabulary

- Power
- Resume
- Adventure
- Journey
- Checkpoint
- Save Point
- Ready
- Continue
- Heartbeat
- Memory
- Momentum
- Launch
- Player One
- Arcade
- Spark
- Dreaming
- Awakening

---

# 29. Mascot / Companion Direction

The project may eventually have a subtle original companion.

Ideas:

- pixel battery
- tiny handheld
- small spark
- little robot technician
- magic cartridge
- save-point icon

Guidelines:

- Original.
- Simple.
- Terminal-friendly.
- Works in monochrome.
- Does not copy Nintendo or RetroPie branding.
- Used sparingly.
- Enhances delight without distracting from function.

---

# 30. Engineering Principles

Draft principles:

1. Build experiences, not just software.
2. Documentation is part of the product.
3. Design before implementation.
4. Architecture exists to enable implementation, not delay it.
5. Every work session ends with a victory.
6. Respect future contributors.
7. Name things according to real-world behavior.
8. Make complexity invisible.
9. Leave moments of delight.
10. Measure before optimizing.
11. Portability is a feature.
12. Curiosity is a strength.
13. Craftsmanship matters.
14. Hardware-specific code lives behind interfaces.
15. No polling if events are available.
16. Systemd is the service manager.
17. Prefer one daemon until complexity proves otherwise.
18. Keep code testable without Raspberry Pi hardware.
19. Use data to guide performance work.
20. Never lose today's excitement in tomorrow's implementation.

---

# 31. Development and Portability

The project should support development away from the target hardware.

## Development hosts

- macOS Intel
- macOS Apple Silicon
- Linux x64
- Linux ARM64
- Linux ARMv7
- Windows where practical for non-GPIO code

## Build targets

- linux/amd64
- linux/arm64
- linux/arm/v7
- darwin/amd64
- darwin/arm64

Possible future:

- windows/amd64 for non-hardware code

## Testing strategy

Most code should compile and test without GPIO hardware.

Use:

- interfaces
- mock hardware
- fake GPIO events
- simulated frontend events
- temporary state stores
- CI cross-compilation

Hardware validation remains required for real GPIO, power, and boot timing behavior.

---

# 32. GitHub Project Direction

The repository should feel like a professional open-source project.

Expected files and systems:

- GitHub Actions CI
- Release workflow
- Issue templates
- Pull request template
- CODEOWNERS
- CONTRIBUTING
- CODE_OF_CONDUCT
- ADR templates
- RFC templates
- documentation metadata
- milestones
- project board
- semantic versioning

## Versioning

Use semantic versioning:

```text
v0.1.0
v0.2.0
v1.0.0
```

Releases may also have codenames.

Example:

```text
v0.1.0 — Dreaming
v0.2.0 — Awakening
v0.3.0 — Heartbeat
v0.4.0 — Memory
```

---

# 33. Success Metrics

Possible success metrics:

## Boot and resume

- Target: power switch to resumed gameplay in less than 20 seconds.
- Stretch: less than 15 seconds.

## Power daemon

- Idle CPU usage below 0.1%.
- Memory below 10 MB RSS if practical.

## Reliability

- Safe shutdown remains reliable.
- Simulated power events run repeatedly without failure.
- Resume failure does not strand the user.
- Save states are protected from corruption.

## Developer experience

- Project builds on macOS and Linux.
- Most tests run without Raspberry Pi hardware.
- New contributors can understand architecture from docs.
- AI assistants can work effectively from repository context.

## Experience

- The handheld feels more like a console.
- CLI output is useful and enjoyable.
- Documentation feels welcoming.
- The repository feels thoughtfully crafted.

---

# 34. Compatibility Strategy

Use support tiers.

## Tier 1 — Reference platform

- RetroFlag GPi Case 2
- Raspberry Pi CM4
- Raspberry Pi OS Bullseye 64-bit
- RetroPie
- EmulationStation
- RetroArch

## Tier 2 — Future official targets

- Raspberry Pi 4
- Raspberry Pi 5
- Raspberry Pi Zero 2 W
- Future Compute Modules
- GPi Case 2W
- NESPi-style cases
- SUPERPi-style cases

## Tier 3 — Community targets

- Orange Pi
- Radxa ROCK
- ODROID
- Banana Pi
- Libre Computer
- Other Linux SBC retro gaming systems

## Tier 4 — Experimental / partial targets

- x64 retro gaming systems
- Steam Deck-like systems
- Batocera appliances
- non-GPIO resume/metrics use cases

---

# 35. Compatibility Matrix Draft

| Feature | GPi Case 2 + CM4 | GPi Case 2W | Pi 4/5 Case | Other SBC | x64 |
|---|---:|---:|---:|---:|---:|
| Safe shutdown | Target | Future | Future | Community | N/A |
| Power switch GPIO | Target | Future | Depends | Depends | N/A |
| Reset button GPIO | Target | Future | Depends | Depends | N/A |
| Resume last game | Target | Future | Future | Possible | Possible |
| Boot metrics | Target | Future | Future | Possible | Possible |
| CLI status | Target | Future | Future | Possible | Possible |
| Battery support | Research | Research | Depends | Depends | N/A |
| Sleep mode | Research | Research | Research | Research | Partial |
| Splash resume screen | Target | Future | Future | Possible | Possible |

---

# 36. Immediate Artifact Plan

To avoid losing context while moving one file at a time, create artifacts in this approximate order:

1. WHY.md
2. PROJECT_MEMORY.md
3. ENGINEERING_MANIFESTO.md
4. PROJECT_MANIFEST.md
5. PROJECT_CHARTER.md
6. docs/13-reference/terminology.md
7. docs/01-product/vision.md
8. docs/01-product/goals.md
9. docs/01-product/personas.md
10. docs/01-product/user-stories.md
11. docs/01-product/success-metrics.md
12. docs/02-hardware/gpi-case-2.md
13. docs/02-hardware/raspberry-pi.md
14. docs/02-hardware/hardware-profiles.md
15. docs/03-software/retropie-stack.md
16. docs/04-architecture/system-overview.md
17. docs/04-architecture/event-bus.md
18. docs/04-architecture/hardware-abstraction.md
19. docs/04-architecture/power-service.md
20. docs/04-architecture/resume-service.md
21. docs/05-development/git-workflow.md
22. docs/05-development/ai-collaboration.md
23. docs/05-development/coding-standards.md
24. docs/06-testing/testing-strategy.md
25. docs/07-design/terminal-ui.md
26. docs/07-design/style-guide.md
27. docs/08-performance/boot-performance.md
28. docs/09-research/index.md
29. docs/10-decisions/adr-template.md
30. docs/11-rfc/rfc-template.md

This order may change, but the memory document should remain the safety net.

---

# 37. Things That Must Not Be Lost

This checklist should be reviewed periodically.

- [ ] Root-level WHY.md exists.
- [ ] Personal ENGINEERING_MANIFESTO.md exists.
- [ ] Metadata header standard is applied.
- [ ] Power Switch vs Reset Button terminology is documented.
- [ ] Project philosophy includes delight and craftsmanship.
- [ ] Taft Consulting slogan is recorded.
- [ ] Hardware reference platform is documented.
- [ ] RetroFlag broader hardware line is acknowledged.
- [ ] Raspberry Pi focus is documented.
- [ ] Other SBC awareness is documented.
- [ ] SafeShutdown.py replacement is documented.
- [ ] systemd replacement plan is documented.
- [ ] libgpiod direction is documented.
- [ ] Resume Manager vision is documented.
- [ ] Splash resume screen idea is documented.
- [ ] Background save-state idea is documented.
- [ ] State Manager scope is documented.
- [ ] Frontend integration is documented.
- [ ] Sleep mode is marked as research.
- [ ] Boot performance goals are documented.
- [ ] Companion CLI commands are documented.
- [ ] Terminal UX style is documented.
- [ ] ASCII and color personality is documented.
- [ ] Mascot/companion idea is documented.
- [ ] Planning discipline is documented.
- [ ] One-artifact-per-session rule is documented.
- [ ] No-planning-paralysis risk is documented.
- [ ] Development off-Pi is documented.
- [ ] Cross-compilation goals are documented.
- [ ] Compatibility tiers are documented.
- [ ] Success metrics are documented.
- [ ] ADR/RFC process is documented.
- [ ] Requirements ID system is documented.
- [ ] Risk register is created.
- [ ] Open questions are created.
- [ ] Assumption log is created.
- [ ] Documentation index is created.

---

# 38. Closing Note

This file is not the final documentation.

It is the safety net.

When an idea graduates into a proper document, it should remain referenced here until the project has a stable documentation index.

The purpose of this file is simple:

> Make sure nothing important from the dream is lost before the project wakes up.
