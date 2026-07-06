---
id: ARCHIVE-PROJECT-MILESTONES-001
title: Pre-EDC Project Milestone Journey
version: 0.1.0
status: Archived
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Preserve the pre-EDC eleven-stage milestone journey as historical planning context after large life stages were renamed Epochs.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - ENGINEERING_MANIFESTO.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - docs/00-project/roadmap.md
  - docs/00-project/milestones.md
  - docs/adr/0003-adopt-epoch-milestone-quest-model.md
  - docs/13-reference/terminology.md
last_updated: 2026-07-06
---

# Pre-EDC Project Milestone Journey

> Archived on 2026-07-06. This document preserves the former eleven-stage
> journey for historical context. It is not the canonical Epoch ladder or
> Milestone ledger.

Canonical current sources:

- Epoch ladder and current direction:
  `docs/00-project/roadmap.md`
- Numbered verified Milestones:
  `docs/00-project/milestones.md`
- Vocabulary decision:
  `docs/adr/0003-adopt-epoch-milestone-quest-model.md`

---

> The dream needs a path so it can become real.

This document defines the milestone path for RetroFlag Power.

The project has a story, but milestones are not just story names.

Each milestone has:

- a purpose
- a primary outcome
- included work
- excluded work
- exit criteria
- commit themes
- risks to watch

The goal is to protect momentum.

A milestone should help the project move forward, not become a place to hide from implementation.

---

# Milestone Philosophy

RetroFlag Power should grow through clear, meaningful phases.

Each phase should leave the project more real than it was before.

The project should avoid two traps:

## Trap 1 — Planning forever

Planning is valuable when it creates artifacts.

Planning becomes dangerous when it delays all implementation.

## Trap 2 — Coding without direction

Implementation is valuable when it follows a clear purpose.

Coding becomes dangerous when it creates systems without shared intent.

The milestone path balances both.

```
Dream
  │
  ▼
Design
  │
  ▼
Build
  │
  ▼
Measure
  │
  ▼
Refine
  │
  ▼
Release
```

---

# Milestone 0 — Dreaming

## Theme

Before software comes intention.

## Purpose

Capture the project's identity, values, terminology, architecture direction, and first implementation path before writing production code.

The project exists, but it is not awake yet.

It is dreaming about what it wants to become.

## Primary outcome

The repository contains enough documentation and shared language that future contributors, maintainers, and AI assistants can understand the project without needing the original conversation.

## Included work

- WHY.md
- PROJECT_MEMORY.md
- ENGINEERING_MANIFESTO.md
- PROJECT_MANIFEST.md
- PROJECT_CHARTER.md
- terminology guide
- glossary
- product vision
- roadmap
- high-level requirements
- high-level architecture
- hardware reference notes
- development workflow
- AI collaboration guidance
- ADR/RFC templates

## Excluded work

Milestone 0 should not include production daemon implementation.

It may include repository cleanup, documentation structure, diagrams, and templates.

It may include small tooling improvements if they directly support documentation or repository health.

## Exit criteria

Milestone 0 is complete when:

- [ ] WHY.md exists.
- [ ] PROJECT_MEMORY.md exists and captures the full blueprint.
- [ ] ENGINEERING_MANIFESTO.md exists.
- [ ] PROJECT_MANIFEST.md exists.
- [ ] PROJECT_CHARTER.md exists.
- [ ] Terminology guide exists.
- [ ] Glossary exists.
- [ ] Product vision exists.
- [ ] Roadmap exists.
- [ ] High-level requirements exist.
- [ ] Initial architecture overview exists.
- [ ] Hardware reference platform is documented.
- [ ] Development workflow is documented.
- [ ] AI collaboration guidance exists.
- [ ] ADR template exists.
- [ ] RFC template exists.
- [ ] Milestone 1 has clear entry criteria.
- [ ] The project has enough clarity to begin implementation without inventing the architecture from scratch.

## Commit themes

Possible commit titles:

```
Begin dreaming.
Give the dream a direction.
Give the dream its charter.
Name the world of the dream.
Give the dream a path.
```

## Risks to watch

- Planning becomes comfortable and delays implementation.
- Documents become too large before code exists.
- The project tries to define every future feature too early.
- The story becomes more important than the engineering.

## Victory condition

The project knows why it exists, what it values, what language it speaks, and what the first real implementation step is.

---

# Milestone 1 — Awakening

## Theme

The project takes its first breath.

## Purpose

Introduce the first running daemon.

The daemon does not need GPIO yet.

It only needs to prove that the project can build, run, log, and stop cleanly.

## Primary outcome

`retroflag-powerd` exists as a minimal supervised daemon skeleton.

## Included work

- `cmd/retroflag-powerd`
- daemon entry point
- structured logging
- signal handling
- context-based lifecycle
- graceful shutdown
- version command or flag
- basic Makefile targets
- initial systemd service file
- install script
- status/logging instructions
- basic CI build validation

## Excluded work

- real GPIO handling
- actual shutdown actions
- resume manager
- frontend integration
- boot optimization
- complex configuration
- release packaging

## Exit criteria

Milestone 1 is complete when:

- [ ] The daemon builds locally.
- [ ] The daemon runs locally.
- [ ] The daemon prints structured startup logs.
- [ ] The daemon handles SIGINT and SIGTERM.
- [ ] The daemon exits cleanly.
- [ ] `make build` works.
- [ ] `make run` works.
- [ ] CI builds the project.
- [ ] A systemd service file exists.
- [ ] The daemon can be installed on the reference platform.
- [ ] `systemctl status retroflag-power` shows the service running.
- [ ] `journalctl -u retroflag-power` shows readable logs.

## Commit themes

Possible commit titles:

```
The project takes its first breath.
Give the daemon a voice.
Teach the daemon to say goodbye.
Place the heartbeat under systemd.
```

## Risks to watch

- Adding GPIO too soon.
- Overbuilding the daemon framework.
- Creating abstractions before there is behavior.
- Making configuration too flexible too early.

## Victory condition

The project has a living process.

It can start.

It can speak.

It can stop.

---

# Milestone 2 — Heartbeat

## Theme

The daemon becomes a reliable service.

## Purpose

Turn the daemon skeleton into a real supervised service with health, logging, status, and predictable behavior.

## Primary outcome

The daemon can run continuously under systemd and provide useful operational visibility.

## Included work

- service lifecycle refinement
- journald logging polish
- daemon status command
- internal health state
- systemd restart policy
- optional watchdog research
- basic diagnostics
- installation verification
- documentation for service management

## Excluded work

- advanced CLI UI
- resume logic
- save-state handling
- hardware-specific optimizations beyond service needs

## Exit criteria

Milestone 2 is complete when:

- [ ] Service starts reliably at boot.
- [ ] Service restarts on failure.
- [ ] Service can be stopped cleanly.
- [ ] Logs are useful in journald.
- [ ] A status command or equivalent exists.
- [ ] Installation docs are accurate.
- [ ] Troubleshooting docs exist for service startup failures.
- [ ] The daemon can report a basic healthy state.

## Commit themes

Possible commit titles:

```
Give the daemon a heartbeat.
Teach the service to stand back up.
Make the logs worth reading.
```

## Risks to watch

- Premature watchdog complexity.
- Treating status output as final UI too early.
- Allowing service management to distract from power behavior.

## Victory condition

The daemon behaves like a proper Linux service.

---

# Milestone 3 — Power

## Theme

The daemon learns the Power Switch.

## Purpose

Replace the original SafeShutdown.py behavior with modern event-driven hardware handling.

## Primary outcome

The daemon responds correctly to the reference platform's Power Switch and Reset Button while preserving safe shutdown behavior.

## Included work

- hardware profile for GPi Case 2
- GPIO abstraction
- libgpiod implementation
- mock GPIO implementation
- Power Switch capability
- Reset Button capability
- event definitions
- power service
- shutdown manager
- reboot/reset handling
- safety checks
- reference platform validation
- documentation for wiring and behavior

## Excluded work

- automatic game resume
- splash screen
- background autosave
- advanced battery behavior
- support for non-reference hardware beyond design stubs

## Exit criteria

Milestone 3 is complete when:

- [ ] The GPi Case 2 Power Switch is modeled as a latching switch.
- [ ] The Reset Button is modeled as a momentary button.
- [ ] GPIO implementation is isolated behind interfaces.
- [ ] Mock hardware tests can simulate Power Switch and Reset Button events.
- [ ] The daemon responds to Power Switch OFF.
- [ ] The daemon initiates safe shutdown.
- [ ] The daemon handles Reset Button behavior according to documented policy.
- [ ] Logs clearly show power events.
- [ ] The original rc.local SafeShutdown.py path is no longer required.
- [ ] Install docs explain how to disable the original script safely.
- [ ] Hardware validation notes exist.

## Commit themes

Possible commit titles:

```
Teach the daemon the Power Switch.
Give the hardware a profile.
Let events replace polling.
Retire the old shutdown spell.
```

## Risks to watch

- Breaking safe shutdown.
- Misreading active-low/active-high GPIO behavior.
- Confusing Power Switch and Power Button terminology.
- Making the hardware profile too specific or too generic.
- Accidentally depending on Raspberry Pi-only behavior in core services.

## Victory condition

The original SafeShutdown.py responsibility has a modern replacement.

---

# Milestone 4 — Memory

## Theme

The handheld begins to remember.

## Purpose

Introduce game session tracking and the foundation for resume behavior.

## Primary outcome

The system can record what the player was doing and store that state safely.

## Included work

- state service
- session model
- state file format
- atomic state writes
- state validation
- RetroArch research
- EmulationStation integration research
- game launch detection
- game exit detection
- current session recording
- CLI status for session state
- failure recovery rules

## Excluded work

- fully automatic resume
- splash screen
- background autosaves
- multi-slot resume
- battery overlay

## Exit criteria

Milestone 4 is complete when:

- [ ] Session state model exists.
- [ ] State storage path is defined.
- [ ] State writes are atomic.
- [ ] State can be read and validated.
- [ ] Current game/session can be recorded.
- [ ] Missing ROM or invalid state is handled safely.
- [ ] CLI can show current or last known session.
- [ ] Documentation explains the state file.
- [ ] Tests cover state read/write behavior.

## Commit themes

Possible commit titles:

```
Teach the handheld to remember.
Create the first save point.
Record the shape of a session.
```

## Risks to watch

- State file corruption.
- Storing too much too early.
- Hardcoding RetroPie paths too deeply.
- Treating research assumptions as facts.

## Victory condition

The system can remember the previous session safely.

---

# Milestone 5 — Resume

## Theme

The adventure continues.

## Purpose

Restore the previous game session after boot.

## Primary outcome

The system can launch the previous game and restore a save state on the reference platform.

## Included work

- resume service
- RetroArch save-state integration
- emulator/core launch integration
- frontend coordination
- resume fallback behavior
- resume CLI commands
- resume success/failure metrics
- documentation for supported cores
- first compatibility notes

## Excluded work

- polished splash screen
- background autosaves
- per-core advanced behavior
- frontend tile integration
- sleep mode

## Exit criteria

Milestone 5 is complete when:

- [ ] Previous session can be launched.
- [ ] Save state can be restored for at least one validated system/core.
- [ ] Resume failure falls back safely.
- [ ] Resume can be disabled.
- [ ] Resume state can be cleared.
- [ ] CLI can report resume status.
- [ ] Logs clearly explain resume attempts.
- [ ] Documentation explains supported behavior and limitations.
- [ ] At least one full power-off/power-on resume flow is validated on the reference platform.

## Commit themes

Possible commit titles:

```
Continue the adventure.
Restore the checkpoint.
Teach the handheld to return.
```

## Risks to watch

- Save-state corruption.
- Unsupported cores.
- Timing issues during boot.
- Frontend and emulator race conditions.
- User confusion when resume fails.

## Victory condition

The handheld can return to a previous game session.

---

# Milestone 6 — Momentum

## Theme

Measure the path from power to play.

## Purpose

Measure and improve actual and perceived startup/resume performance.

## Primary outcome

The project can benchmark the power-on-to-resume path and identify meaningful improvements.

## Included work

- boot timing metrics
- resume timing metrics
- benchmark command
- historical metric storage
- systemd timing analysis
- service startup analysis
- documentation of current baseline
- safe boot optimization research
- perceived startup improvements

## Excluded work

- risky firmware changes by default
- forced EEPROM modifications
- unsupported kernel changes
- optimization without measurements

## Exit criteria

Milestone 6 is complete when:

- [ ] Baseline boot-to-resume time is measured.
- [ ] Benchmark command exists.
- [ ] Metrics are stored or displayed clearly.
- [ ] Performance docs explain measurement method.
- [ ] At least one safe improvement is identified or implemented.
- [ ] Research notes distinguish actual speed from perceived speed.
- [ ] The 20-second target is evaluated against real data.

## Commit themes

Possible commit titles:

```
Measure the path from power to play.
Give speed a scoreboard.
Chase the twenty-second dream.
```

## Risks to watch

- Optimizing before measuring.
- Breaking reliability for speed.
- Confusing perceived startup improvements with actual boot reduction.
- Applying hardware-specific tweaks too broadly.

## Victory condition

The project can prove whether it is getting faster.

---

# Milestone 7 — Polish

## Theme

Make the machine feel crafted.

## Purpose

Improve terminal UX, diagnostics, splash/resume presentation, documentation quality, and user-facing polish.

## Primary outcome

The project feels enjoyable to use without sacrificing clarity.

## Included work

- terminal UI style guide
- reusable CLI UI components
- status cards
- diagnostics/doctor command
- no-color mode
- JSON/scriptable mode where useful
- polished benchmark output
- resume splash screen prototype
- accessibility checks
- troubleshooting improvements

## Excluded work

- branding that copies protected properties
- excessive animations
- UI polish that hides errors
- features that reduce reliability

## Exit criteria

Milestone 7 is complete when:

- [ ] Terminal UI guide exists.
- [ ] CLI output is consistent.
- [ ] Errors are clear and actionable.
- [ ] Success/status output has tasteful personality.
- [ ] Diagnostics command exists or is designed.
- [ ] Output degrades gracefully without color.
- [ ] Documentation reflects the UI tone.

## Commit themes

Possible commit titles:

```
Leave moments of delight.
Make the terminal smile.
Polish the checkpoint screen.
```

## Risks to watch

- Personality becoming gimmicky.
- Making output harder to parse.
- Prioritizing polish over reliability.
- Introducing inaccessible color-only meaning.

## Victory condition

The project feels crafted.

---

# Milestone 8 — Expansion

## Theme

The dream learns new machines.

## Purpose

Prepare the project for hardware profiles beyond the GPi Case 2 reference platform.

## Primary outcome

The hardware profile system is documented and usable by contributors.

## Included work

- hardware profile schema
- capability matrix
- profile validation
- contribution guide for hardware porters
- documentation for RetroFlag product line awareness
- Raspberry Pi 4/5 notes
- GPi Case 2W notes
- NESPi/SUPERPi research
- community SBC support notes

## Excluded work

- official support for unvalidated hardware
- pretending community targets are first-class before testing
- complex plugin system unless needed

## Exit criteria

Milestone 8 is complete when:

- [ ] Hardware profile format is documented.
- [ ] Capabilities are documented.
- [ ] GPi Case 2 profile is complete.
- [ ] A contributor guide exists for adding hardware.
- [ ] Compatibility matrix exists.
- [ ] Unsupported hardware can be described honestly.
- [ ] The core architecture remains hardware-agnostic.

## Commit themes

Possible commit titles:

```
Open the map beyond the first kingdom.
Teach the dream about new machines.
Draw the hardware atlas.
```

## Risks to watch

- Supporting hardware without validation.
- Making the profile system too abstract.
- Expanding scope before the reference platform is solid.

## Victory condition

The project can grow beyond one handheld without losing its shape.

---

# Milestone 9 — Release

## Theme

The adventure leaves the workshop.

## Purpose

Prepare the project for public use.

## Primary outcome

RetroFlag Power can be installed, tested, understood, and updated by users outside the original development environment.

## Included work

- release workflow
- versioning
- changelog process
- install documentation
- upgrade documentation
- uninstall documentation
- troubleshooting guide
- known limitations
- release artifacts
- checksums
- GitHub release notes
- package research
- beta validation checklist

## Excluded work

- claiming v1.0 stability too early
- unsupported hardware claims
- hiding limitations

## Exit criteria

Milestone 9 is complete when:

- [ ] Release process is documented.
- [ ] CI builds release artifacts.
- [ ] Installation instructions are tested.
- [ ] Upgrade path is documented.
- [ ] Known limitations are documented.
- [ ] Troubleshooting guide exists.
- [ ] A tagged pre-release can be created.
- [ ] Users can install without reading source code.

## Commit themes

Possible commit titles:

```
Prepare the adventure for travelers.
Pack the first release satchel.
Open the gate to beta testers.
```

## Risks to watch

- Releasing before safe shutdown is reliable.
- Incomplete install/uninstall instructions.
- Users misunderstanding beta limitations.
- Artifacts that do not match documentation.

## Victory condition

Someone else can try the project safely.

---

# Milestone 10 — Launch

## Theme

The project becomes real for others.

## Purpose

Deliver a stable release for the reference platform.

## Primary outcome

A documented, validated, stable release exists for the RetroFlag GPi Case 2 reference platform.

## Included work

- v1.0 release candidate
- validation checklist
- reference platform test report
- documentation audit
- troubleshooting audit
- compatibility matrix
- known limitations
- support process
- release notes
- archive of milestone decisions

## Excluded work

- every possible hardware target
- every possible emulator/core
- every future feature
- unvalidated claims

## Exit criteria

Milestone 10 is complete when:

- [ ] Reference platform safe shutdown is validated.
- [ ] Resume behavior is validated for documented supported cases.
- [ ] Installation and uninstall are validated.
- [ ] Documentation is accurate for v1.0.
- [ ] Known limitations are clearly stated.
- [ ] Recovery paths are documented.
- [ ] Release artifacts are available.
- [ ] The project can be recommended to GPi Case 2 users with confidence.

## Commit themes

Possible commit titles:

```
Launch the adventure.
Open the cartridge slot.
Send Player One onward.
```

## Risks to watch

- Overpromising support.
- Treating beta behavior as stable.
- Skipping documentation audit.
- Neglecting recovery paths.

## Victory condition

RetroFlag Power becomes a real, usable project for users beyond the original builder.

---

# Cross-Milestone Rules

## Rule 1 — Safety before delight

Delight is important.

Safe shutdown and user data are more important.

## Rule 2 — Documentation tracks decisions

If a major decision is made, record it.

## Rule 3 — Research is not commitment

Research topics should be clearly marked until validated.

## Rule 4 — Hardware support requires honesty

Do not claim official support without validation.

## Rule 5 — Personality supports clarity

Warmth is good.

Confusion is not.

## Rule 6 — Every session ends with a victory

Even if the victory is small.

Especially if the victory is small.

---

# Immediate Next Path

Current path through Milestone 0:

- [x] WHY.md
- [x] PROJECT_MEMORY.md
- [x] ENGINEERING_MANIFESTO.md
- [x] PROJECT_MANIFEST.md
- [x] PROJECT_CHARTER.md
- [x] docs/13-reference/terminology.md
- [ ] docs/13-reference/glossary.md
- [ ] docs/01-product/vision.md
- [ ] docs/00-project/roadmap.md
- [ ] docs/00-project/requirements.md
- [ ] docs/04-architecture/system-overview.md
- [ ] docs/02-hardware/gpi-case-2.md
- [ ] docs/05-development/ai-collaboration.md
- [ ] docs/10-decisions/adr-template.md
- [ ] docs/11-rfc/rfc-template.md

Milestone 0 should not grow forever.

When the remaining foundation documents are complete enough to guide implementation, the project should move to Milestone 1 — Awakening.

---

# Closing

A dream without a path stays a dream.

A path without a dream becomes chores.

RetroFlag Power needs both.

Milestone 0 gives the project its purpose.

Milestone 1 gives it breath.

Milestone 2 gives it rhythm.

Milestone 3 gives it power.

Milestone 4 gives it memory.

Milestone 5 lets the adventure continue.

The path is now visible.

Walk it one victory at a time.
