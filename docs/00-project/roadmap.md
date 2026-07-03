---
id: ROADMAP-001
title: Project Roadmap
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Provide a practical near-term roadmap from Milestone 0 documentation through the first running daemon, safe shutdown replacement, resume foundation, and public release path.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - docs/00-project/milestones.md
  - docs/01-product/vision.md
  - docs/13-reference/terminology.md
last_updated: 2026-07-03
---

# Project Roadmap

> A milestone is the chapter. A roadmap is the next road sign.

This roadmap translates the larger milestone story into a practical build path.

The milestone document explains the full journey.

This roadmap answers:

```
What should we do next?
What should wait?
What is enough to move forward?
When do we start writing code?
```

The purpose of the roadmap is momentum.

---

# 1. Roadmap Philosophy

RetroFlag Power should avoid two extremes:

## Do not rush into code without direction

The project needs purpose, terminology, scope, and architecture before implementation.

## Do not document forever

Documentation should guide implementation, not delay it.

The roadmap exists to keep the project moving.

Each phase should produce commit-worthy artifacts.

Each phase should have a clear stopping point.

---

# 2. Current Project State

Current phase:

```
Milestone 0 — Dreaming
```

Current status:

```
The project is defining its identity, language, product vision, and implementation path before production code begins.
```

Artifacts already created or planned in this stage include:

- WHY.md
- PROJECT_MEMORY.md
- ENGINEERING_MANIFESTO.md
- PROJECT_MANIFEST.md
- PROJECT_CHARTER.md
- docs/13-reference/terminology.md
- docs/13-reference/glossary.md
- docs/00-project/milestones.md
- docs/01-product/vision.md
- docs/00-project/roadmap.md

The project now has:

- a spark
- a safety net
- a builder philosophy
- a compass
- a covenant
- a language
- a dictionary
- a path
- a horizon

The next goal is to finish only the minimum useful foundation needed to begin implementation confidently.

---

# 3. Near-Term Route

The immediate route is:

```
Dreaming Foundation
      │
      ▼
Requirements
      │
      ▼
System Overview
      │
      ▼
Hardware Reference
      │
      ▼
Development / AI Guidance
      │
      ▼
ADR + RFC Templates
      │
      ▼
Milestone 1 Entry
      │
      ▼
First Daemon
```

Milestone 0 should end when implementation can begin without inventing the project from scratch.

Milestone 0 should not attempt to fully document every future feature.

---

# 4. Phase A — Finish the Foundation

## Goal

Complete the core documents that define why the project exists, what experience it is creating, what scope it has, and what language it speaks.

## Status

Mostly underway.

## Artifacts

- [x] WHY.md
- [x] PROJECT_MEMORY.md
- [x] ENGINEERING_MANIFESTO.md
- [x] PROJECT_MANIFEST.md
- [x] PROJECT_CHARTER.md
- [x] docs/13-reference/terminology.md
- [x] docs/13-reference/glossary.md
- [x] docs/00-project/milestones.md
- [x] docs/01-product/vision.md
- [x] docs/00-project/roadmap.md

## Done when

The project has enough philosophical and product clarity that remaining work can become requirements and architecture instead of more discovery.

## Risk

The project may keep producing beautiful foundation docs instead of moving toward implementation.

## Guardrail

After this phase, only create new philosophy documents if they directly unblock implementation or protect important context.

---

# 5. Phase B — Requirements

## Goal

Create the first traceable requirements document.

## Artifact

```
docs/00-project/requirements.md
```

## Purpose

Requirements define what the system must do.

They should be practical, traceable, and implementation-guiding.

## Initial requirement groups

- Safe shutdown
- systemd service behavior
- GPIO / hardware abstraction
- terminology and naming
- portability
- logging
- development workflow
- state and resume foundation
- performance targets
- safety and recovery
- terminal UX
- documentation

## Example requirements

```
REQ-0001
The system shall safely shut down when the Power Switch moves to OFF.
```

```
REQ-0002
The daemon shall run under systemd on the reference platform.
```

```
REQ-0003
The project shall compile on Linux ARM64.
```

```
REQ-0004
Hardware-specific GPIO details shall not leak into core services.
```

## Done when

- Requirements have stable IDs.
- Critical requirements are identified.
- Milestone 1 requirements are clear.
- Requirements are not so detailed that they block starting the daemon.

## Risk

Over-specifying features that are still research.

## Guardrail

Mark uncertain items as research or future requirements.

---

# 6. Phase C — System Overview

## Goal

Create the first architecture overview.

## Artifact

```
docs/04-architecture/system-overview.md
```

## Purpose

Define the shape of the system before writing code.

This should include:

- daemon
- internal services
- event bus
- hardware profiles
- state storage
- CLI
- systemd
- frontend/emulator adapters
- boundaries between core and platform-specific code

## Initial architecture shape

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

## Done when

- A contributor can understand the intended system shape.
- The first daemon package layout is clear enough to implement.
- Hardware and software adapters are conceptually separated.
- The architecture is useful but not overdesigned.

## Risk

Designing a perfect architecture before behavior exists.

## Guardrail

The overview should guide Milestone 1 only deeply enough to build the daemon skeleton.

---

# 7. Phase D — Hardware Reference

## Goal

Document the reference hardware and known current state.

## Artifact

```
docs/02-hardware/gpi-case-2.md
```

## Purpose

Capture the known facts about the reference platform.

## Include

- RetroFlag GPi Case 2
- Raspberry Pi CM4 Rev 1.1
- 4 GB RAM
- Samsung EVO Select microSD
- RetroPie
- Raspberry Pi OS Bullseye 64-bit
- current EEPROM 2022-04-26
- current SafeShutdown.py path
- current rc.local startup
- current boot optimizations
- Power Switch terminology
- Reset Button terminology
- research questions
- validation checklist

## Done when

- The reference platform is clear.
- Known facts are separated from assumptions.
- Current boot and shutdown state is documented.
- Future hardware profile work has a starting point.

## Risk

Trying to research every RetroFlag device before supporting the first one.

## Guardrail

Focus on GPi Case 2 now.

Mention broader hardware awareness, but do not expand beyond what is needed for the first implementation.

---

# 8. Phase E — Development and AI Guidance

## Goal

Document how humans and AI assistants should work in the repository.

## Artifacts

```
docs/05-development/workflow.md
docs/05-development/ai-collaboration.md
```

Possible later:

```
ai/00_PROJECT_CONSTITUTION.md
ai/01_PROJECT_OVERVIEW.md
```

## Purpose

Make the repository easier to work in from VS Code, Codex, ChatGPT, and future tooling.

## Include

- one artifact per work session
- commit packet strategy
- metadata header requirement
- how to use PROJECT_MEMORY.md
- how to decide when to create ADRs/RFCs
- coding standards pointer
- testing expectation
- no-planning-paralysis rule
- copy/paste command preference
- artifact download workflow
- AI assistant project expectations

## Done when

- A future AI assistant can read the repo and understand how to help.
- A future human contributor can understand the workflow.
- The project has explicit guardrails against endless planning.

## Risk

Creating too much process too early.

## Guardrail

Keep the workflow practical.

The goal is to help contributors move, not slow them down.

---

# 9. Phase F — ADR and RFC Templates

## Goal

Create templates for decisions and proposals.

## Artifacts

```
docs/10-decisions/adr-template.md
docs/11-rfc/rfc-template.md
```

## Purpose

Give future decisions and proposals a consistent structure.

## ADR template should include

- status
- context
- decision
- alternatives considered
- consequences
- related requirements
- related risks

## RFC template should include

- proposal
- motivation
- user impact
- technical approach
- alternatives
- risks
- open questions
- acceptance criteria

## Done when

- The project can record decisions consistently.
- Future major proposals have a place to start.

## Risk

Bureaucracy.

## Guardrail

Templates should be lightweight.

Use ADRs/RFCs for meaningful decisions, not every tiny change.

---

# 10. Phase G — Milestone 1 Entry

## Goal

Stop documenting and start the first implementation milestone.

## Entry criteria for Milestone 1 — Awakening

Milestone 1 can begin when:

- [ ] WHY.md exists.
- [ ] PROJECT_MEMORY.md exists.
- [ ] ENGINEERING_MANIFESTO.md exists.
- [ ] PROJECT_MANIFEST.md exists.
- [ ] PROJECT_CHARTER.md exists.
- [ ] Terminology guide exists.
- [ ] Glossary exists.
- [ ] Milestones exist.
- [ ] Product vision exists.
- [ ] Roadmap exists.
- [ ] Requirements exist.
- [ ] System overview exists.
- [ ] Reference hardware doc exists.
- [ ] Development / AI workflow exists.
- [ ] ADR template exists.
- [ ] RFC template exists.

This is enough.

Not everything needs to be perfect.

Milestone 1 should begin after this even if future docs remain incomplete.

## First implementation goal

Build a minimal daemon that:

- compiles
- runs
- logs
- handles SIGINT/SIGTERM
- exits cleanly
- can later be supervised by systemd

## First daemon should not include

- GPIO
- shutdown execution
- resume
- frontend integration
- advanced configuration
- complex UI

The first daemon exists to prove the project can breathe.

---

# 11. Medium-Term Roadmap

After Milestone 1 begins, the practical path is:

## Milestone 1 — Awakening

Build the first daemon skeleton.

## Milestone 2 — Heartbeat

Run reliably under systemd with useful logs and status.

## Milestone 3 — Power

Replace SafeShutdown.py behavior with event-driven Power Switch handling.

## Milestone 4 — Memory

Record session state safely.

## Milestone 5 — Resume

Restore a previous game session.

## Milestone 6 — Momentum

Measure and improve boot-to-resume performance.

## Milestone 7 — Polish

Improve CLI, diagnostics, terminal UX, and splash/resume experience.

## Milestone 8 — Expansion

Prepare hardware profiles and community portability.

## Milestone 9 — Release

Package, document, and prepare public beta.

## Milestone 10 — Launch

Stable v1.0 for the reference platform.

---

# 12. Features by Priority

## Critical

- safe shutdown
- systemd service
- logging
- reference hardware validation
- clear install/uninstall path

## High

- event-driven GPIO
- hardware abstraction
- mock hardware tests
- session state tracking
- CLI status
- boot/resume metrics

## Medium

- automatic resume
- resume fallback handling
- diagnostics command
- terminal UI polish
- compatibility matrix
- hardware profile contribution guide

## Research

- sleep-like mode
- battery overlay
- critical battery shutdown countdown
- room-transition autosave
- EEPROM boot-order tuning
- KMS/FKMS changes
- non-Raspberry Pi SBC support

## Future

- multiple resume slots
- frontend tile integration
- packaged installer
- Debian package
- broader hardware support
- public project website

---

# 13. What Should Wait

The following should wait until the foundation is real:

- complex configuration system
- plugin architecture
- multiple daemons
- advanced UI framework
- cross-hardware official support
- battery features
- sleep mode
- release packaging
- automatic update system
- emulator/core compatibility database

These may become important later.

They are not needed before the first daemon breathes.

---

# 14. When to Stop Milestone 0

Milestone 0 should stop when the project has enough clarity to build the first daemon.

Do not wait until:

- every research question is answered
- every future feature is specified
- every hardware platform is documented
- every design idea is polished
- every possible risk is solved

The goal of Dreaming is not certainty.

The goal is direction.

When direction exists, begin Awakening.

---

# 15. The Next Five Artifacts

Recommended next artifacts after this roadmap:

1. `docs/00-project/requirements.md`
2. `docs/04-architecture/system-overview.md`
3. `docs/02-hardware/gpi-case-2.md`
4. `docs/05-development/ai-collaboration.md`
5. `docs/10-decisions/adr-template.md` and `docs/11-rfc/rfc-template.md`

After these, strongly consider entering Milestone 1.

---

# 16. The First Code Commit

The first production code commit should be intentionally small.

Possible commit title:

```
The project takes its first breath.
```

It should add:

- `cmd/retroflag-powerd/main.go`
- minimal daemon lifecycle
- structured logging
- signal handling
- Makefile updates if needed
- basic tests if appropriate

It should not add GPIO.

The first breath should be simple.

---

# 17. Roadmap Review

This roadmap should be reviewed:

- at the end of Milestone 0
- before Milestone 1 begins
- after the first daemon runs
- before GPIO implementation begins
- before resume implementation begins
- before any public release

The roadmap should change when reality teaches us something.

It should not change just because the dream gets distracted.

---

# Closing

The project now has a practical route.

Do not rush.

Do not stall.

Build the next artifact.

Commit the next victory.

Move toward the first breath.

The dream has a path.

Now it has road signs.
