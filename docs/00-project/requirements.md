---
id: REQUIREMENTS-001
title: Project Requirements
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Define the initial traceable requirements for RetroFlag Power so implementation, tests, architecture, and documentation can be tied back to clear project needs.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/01-product/vision.md
  - docs/13-reference/terminology.md
last_updated: 2026-07-03
---

# Project Requirements

> The dream now has laws.

This document defines the initial requirements for RetroFlag Power.

Requirements describe what the system must do, what qualities it must preserve, and what constraints future implementation must respect.

They are not implementation details.

They are not a complete test plan.

They are not a promise that every future idea will be built.

They are traceable statements that help keep code, architecture, tests, and documentation aligned.

---

# 1. Requirement Format

Each requirement uses a stable ID.

Example:

```
REQ-0001
The system shall safely shut down when the Power Switch moves to OFF.
```

Requirement IDs should remain stable once referenced by code, tests, ADRs, RFCs, or documentation.

If a requirement changes significantly, prefer updating its text and version history rather than reusing its ID for a different meaning.

---

# 2. Priority Levels

## Critical

Required for the project to be safe, trustworthy, or usable.

## High

Important for the main product experience or maintainability.

## Medium

Valuable, but not required for the first working implementation.

## Low

Useful polish or future improvement.

## Research

Not yet committed as an implementation requirement.

Requires investigation before becoming a product requirement.

---

# 3. Requirement Status Values

## Draft

Initial requirement. May change.

## Accepted

Requirement is currently part of the project direction.

## Implemented

Requirement has an implementation.

## Verified

Requirement has been validated through tests, hardware validation, or documentation review.

## Deferred

Requirement remains valid but is intentionally postponed.

## Superseded

Requirement has been replaced by another requirement.

## Rejected

Requirement is no longer part of the project.

---

# 4. Functional Requirements

## REQ-0001 — Safe shutdown from Power Switch OFF

Priority: Critical  
Status: Draft  
Milestone: Power

The system shall safely shut down the reference platform when the Power Switch moves to OFF.

Rationale:

Safe shutdown is the project's first responsibility and must preserve or improve the behavior of the original RetroFlag SafeShutdown.py script.

Related:

- PROJECT_CHARTER.md
- docs/13-reference/terminology.md
- ADR: systemd replacement
- ADR: libgpiod direction

## REQ-0002 — Preserve original safe shutdown intent

Priority: Critical  
Status: Draft  
Milestone: Power

The system shall not make shutdown behavior less safe than the original RetroFlag SafeShutdown.py flow.

Rationale:

Modernization must not come at the cost of user data or filesystem safety.

## REQ-0003 — Replace rc.local startup

Priority: Critical  
Status: Draft  
Milestone: Awakening / Heartbeat

The power management daemon shall be started by systemd instead of `/etc/rc.local`.

Rationale:

systemd provides service supervision, dependency management, restart policy, journald logging, and standard Linux service behavior.

## REQ-0004 — Run as a supervised daemon

Priority: Critical  
Status: Draft  
Milestone: Awakening / Heartbeat

The system shall provide a long-running daemon suitable for supervision by systemd.

Expected daemon name:

```
retroflag-powerd
```

Rationale:

The project should behave like a modern Linux service rather than a background shell or Python process launched from legacy startup scripts.

## REQ-0005 — Graceful stop

Priority: Critical  
Status: Draft  
Milestone: Awakening

The daemon shall handle SIGINT and SIGTERM and exit cleanly.

Rationale:

systemd and development workflows need predictable shutdown behavior.

## REQ-0006 — Structured startup logging

Priority: High  
Status: Draft  
Milestone: Awakening

The daemon shall emit clear startup logs.

Rationale:

The first implementation milestone should prove that the daemon can speak before it controls hardware.

## REQ-0007 — journald visibility

Priority: High  
Status: Draft  
Milestone: Heartbeat

When running under systemd, daemon logs shall be visible through `journalctl`.

Rationale:

Troubleshooting should use standard Linux service tooling.

## REQ-0008 — Service status visibility

Priority: High  
Status: Draft  
Milestone: Heartbeat

The project shall provide a way to inspect daemon/service health.

This may initially be through:

```
systemctl status retroflag-power
journalctl -u retroflag-power
```

and may later include:

```
retroflag-power status
```

Rationale:

Power users and maintainers need simple operational visibility.

## REQ-0009 — Reset Button handling

Priority: High  
Status: Draft  
Milestone: Power

The system shall model the Reset Button as a momentary button and handle its events according to documented policy.

Rationale:

The reset control is physically different from the Power Switch and must be named and modeled correctly.

## REQ-0010 — Disable original shutdown path safely

Priority: Critical  
Status: Draft  
Milestone: Power

The project shall document how to safely disable the original `SafeShutdown.py` and `rc.local` startup path when replacing it.

Rationale:

Users must not run two competing shutdown systems at the same time.

---

# 5. Hardware and GPIO Requirements

## REQ-0100 — Power Switch terminology

Priority: Critical  
Status: Draft  
Milestone: Dreaming / Power

The project shall use `Power Switch` for latching power controls such as the GPi Case 2 power control.

Rationale:

The GPi Case 2 uses a latching switch, not a momentary button.

## REQ-0101 — Reset Button terminology

Priority: Critical  
Status: Draft  
Milestone: Dreaming / Power

The project shall use `Reset Button` for momentary reset controls.

Rationale:

Terminology should match real-world hardware behavior.

## REQ-0102 — GPIO isolation

Priority: Critical  
Status: Draft  
Milestone: Power

Raw GPIO details shall not leak into core services.

Core services should consume hardware capabilities and events instead of GPIO pins, line offsets, active-low values, or chip names.

Rationale:

Hardware abstraction keeps the project portable and testable.

## REQ-0103 — Event-driven GPIO direction

Priority: High  
Status: Draft  
Milestone: Power

The project shall prefer event-driven GPIO handling over continuous busy polling when reliable event APIs are available.

Rationale:

The original script polls continuously. Modernizing the system should reduce unnecessary CPU use and improve service behavior.

## REQ-0104 — libgpiod direction

Priority: High  
Status: Draft  
Milestone: Power

The project should use `libgpiod` or a Go GPIO library backed by modern Linux GPIO character-device behavior for Linux GPIO implementation where practical.

Rationale:

`libgpiod` is a modern Linux GPIO userspace direction and avoids older Raspberry Pi-specific GPIO assumptions.

## REQ-0105 — Mock hardware support

Priority: High  
Status: Draft  
Milestone: Power

The project shall provide mock or simulated hardware behavior so core services can be tested without physical Raspberry Pi hardware.

Rationale:

Most development should be possible on macOS or Linux development machines.

## REQ-0106 — Hardware profiles

Priority: High  
Status: Draft  
Milestone: Power / Expansion

The project shall model hardware through profiles that expose capabilities.

Rationale:

The project should support the GPi Case 2 first while leaving a clean path for future RetroFlag cases and other SBCs.

## REQ-0107 — Reference platform documentation

Priority: Critical  
Status: Draft  
Milestone: Dreaming

The project shall document the reference platform.

Reference platform:

- RetroFlag GPi Case 2
- Raspberry Pi CM4 Rev 1.1
- 4 GB RAM
- Samsung EVO Select microSD
- RetroPie
- Raspberry Pi OS Bullseye 64-bit
- EmulationStation
- RetroArch

Rationale:

Implementation and validation need a clear first target.

---

# 6. Resume and State Requirements

## REQ-0200 — Session model

Priority: High  
Status: Draft  
Milestone: Memory

The project shall define a session model capable of representing the player's current or previous game session.

A session may include:

- system
- ROM path
- emulator
- core
- save-state slot
- timestamp

Future session state may include:

- shader
- overlay
- aspect ratio
- integer scaling
- controller profile
- volume
- brightness

Rationale:

Resume requires durable knowledge of what the player was doing.

## REQ-0201 — Durable state storage

Priority: High  
Status: Draft  
Milestone: Memory

The project shall store resume/session state durably.

Initial candidate path:

```
/var/lib/retroresume/state.json
```

This path may change before implementation if architecture decides on a different namespace.

Rationale:

Resume state must survive reboot.

## REQ-0202 — Atomic state writes

Priority: Critical  
Status: Draft  
Milestone: Memory

The project shall write critical state atomically where practical.

Rationale:

Power loss during state writes must not corrupt resume metadata.

## REQ-0203 — Safe resume fallback

Priority: Critical  
Status: Draft  
Milestone: Resume

If resume cannot complete safely, the system shall fall back to a safe state such as returning to the frontend.

Rationale:

A failed resume must not strand the user.

## REQ-0204 — Missing ROM handling

Priority: High  
Status: Draft  
Milestone: Memory / Resume

If stored session state references a missing ROM file, the system shall report the issue clearly and avoid attempting unsafe resume.

Rationale:

Storage may change, ROMs may be moved, and state can become stale.

## REQ-0205 — Resume disable/clear

Priority: High  
Status: Draft  
Milestone: Resume

The project shall provide a way to disable or clear pending resume state.

Possible future commands:

```
retroflag-power resume disable
retroflag-power resume clear
```

Rationale:

Users and maintainers need control when troubleshooting.

## REQ-0206 — Save-state protection

Priority: Critical  
Status: Draft  
Milestone: Resume

Resume features shall avoid overwriting user-managed save states unless explicitly designed and documented.

Rationale:

User progress is more important than convenience.

## REQ-0207 — Background autosave is research

Priority: Research  
Status: Draft  
Milestone: Future

Periodic background save states shall remain a research topic until safety, performance, and core compatibility are understood.

Rationale:

Autosave can reduce lost progress but may introduce corruption or performance risks.

---

# 7. Performance Requirements

## REQ-0300 — Boot-to-resume target

Priority: High  
Status: Draft  
Milestone: Momentum

The project should target Power Switch ON to resumed gameplay in less than 20 seconds on the reference platform.

Rationale:

Fast perceived startup is a primary product goal.

## REQ-0301 — Boot-to-resume stretch target

Priority: Medium  
Status: Draft  
Milestone: Momentum

The project may pursue a stretch target of less than 15 seconds from Power Switch ON to resumed gameplay.

Rationale:

Stretch goals can guide optimization, but must not compromise reliability.

## REQ-0302 — Measure before optimizing

Priority: High  
Status: Draft  
Milestone: Momentum

Performance improvements shall be guided by measured data.

Rationale:

Boot and resume work should distinguish actual speed from perceived speed.

## REQ-0303 — Boot metrics

Priority: High  
Status: Draft  
Milestone: Momentum

The project shall provide or document a way to measure boot and resume timing.

Possible future command:

```
retroflag-power benchmark
```

Rationale:

The project needs objective feedback to evaluate progress toward startup goals.

## REQ-0304 — Idle CPU target

Priority: Medium  
Status: Draft  
Milestone: Heartbeat / Power

The power daemon should use less than 0.1 percent CPU while idle where practical.

Rationale:

Replacing a polling loop should reduce idle resource usage.

## REQ-0305 — Memory target

Priority: Medium  
Status: Draft  
Milestone: Heartbeat

The power daemon should remain lightweight, with a target below 10 MB RSS where practical.

Rationale:

The target hardware is resource constrained compared with desktop systems.

---

# 8. CLI and Terminal UX Requirements

## REQ-0400 — CLI status

Priority: High  
Status: Draft  
Milestone: Heartbeat / Power

The project should provide a status interface for power users and maintainers.

Possible command:

```
retroflag-power status
```

Rationale:

Users need an easy way to inspect service state and hardware detection.

## REQ-0401 — Diagnostics command

Priority: Medium  
Status: Draft  
Milestone: Polish

The project should provide diagnostics that verify key dependencies and state.

Possible command:

```
retroflag-power doctor
```

Rationale:

Troubleshooting should be accessible without manual log spelunking.

## REQ-0402 — Benchmark command

Priority: Medium  
Status: Draft  
Milestone: Momentum

The project should provide a benchmark command for startup and resume timing.

Possible command:

```
retroflag-power benchmark
```

Rationale:

Performance goals should be measurable.

## REQ-0403 — Clear errors

Priority: Critical  
Status: Draft  
Milestone: All

User-facing errors shall be clear, actionable, and searchable.

Rationale:

Personality must never make troubleshooting harder.

## REQ-0404 — Tasteful personality

Priority: Medium  
Status: Draft  
Milestone: Polish

CLI and terminal output may include tasteful personality, ASCII art, color, and celebratory success messages.

Rationale:

The project should be enjoyable to use while remaining professional.

## REQ-0405 — No-color behavior

Priority: Medium  
Status: Draft  
Milestone: Polish

The terminal UI should support no-color or degraded output where practical.

Rationale:

Output should remain readable over SSH, in logs, and in constrained terminals.

## REQ-0406 — Scriptable output

Priority: Medium  
Status: Draft  
Milestone: Polish

Commands that may be used by scripts should eventually support machine-readable output where useful.

Possible format:

```
JSON
```

Rationale:

Power users and automation workflows need reliable parsing.

---

# 9. Documentation Requirements

## REQ-0500 — Metadata headers

Priority: High  
Status: Draft  
Milestone: Dreaming

Durable markdown artifacts shall include metadata front matter.

Rationale:

Metadata supports indexing, traceability, AI assistant context, and long-term maintainability.

## REQ-0501 — Project memory

Priority: Critical  
Status: Draft  
Milestone: Dreaming

The project shall maintain a memory or blueprint capture document until all major ideas are promoted into dedicated artifacts.

Primary document:

```
PROJECT_MEMORY.md
```

Rationale:

Important context must not be lost while artifacts are created incrementally.

## REQ-0502 — Terminology guide

Priority: Critical  
Status: Draft  
Milestone: Dreaming

The project shall maintain a terminology guide.

Rationale:

Consistent language prevents architecture confusion.

## REQ-0503 — Glossary

Priority: High  
Status: Draft  
Milestone: Dreaming

The project shall maintain a glossary for quick term lookup.

Rationale:

Contributors and AI assistants need a compact vocabulary reference.

## REQ-0504 — Milestone tracking

Priority: High  
Status: Draft  
Milestone: Dreaming

The project shall maintain milestone documentation with exit criteria.

Rationale:

Milestones protect momentum and scope control.

## REQ-0505 — ADR process

Priority: High  
Status: Draft  
Milestone: Dreaming

The project shall define an Architecture Decision Record process before major implementation decisions accumulate.

Rationale:

Future maintainers need to understand why decisions were made.

## REQ-0506 — RFC process

Priority: Medium  
Status: Draft  
Milestone: Dreaming

The project should define an RFC process for major proposals.

Rationale:

Large ideas need a place to be explored before they become decisions.

## REQ-0507 — Clear distinction between facts and assumptions

Priority: High  
Status: Draft  
Milestone: Dreaming / All

Documentation shall distinguish facts, assumptions, decisions, research, and aspirations where practical.

Rationale:

The project should remain honest about what is known, chosen, uncertain, or desired.

---

# 10. Development and Portability Requirements

## REQ-0600 — Build on Linux ARM64

Priority: Critical  
Status: Draft  
Milestone: Awakening

The project shall build for Linux ARM64.

Rationale:

The reference platform uses Raspberry Pi CM4 with a 64-bit OS.

## REQ-0601 — Build on Linux ARMv7

Priority: High  
Status: Draft  
Milestone: Awakening

The project should build for Linux ARMv7 where practical.

Rationale:

Future Raspberry Pi and RetroFlag targets may use 32-bit ARM environments.

## REQ-0602 — Build on Linux AMD64

Priority: High  
Status: Draft  
Milestone: Awakening

The project should build on Linux AMD64.

Rationale:

CI and many development environments use Linux AMD64.

## REQ-0603 — Build on macOS ARM64

Priority: High  
Status: Draft  
Milestone: Awakening

The project should build on macOS ARM64 for non-hardware-specific code.

Rationale:

Day-to-day development may happen on Apple Silicon Macs.

## REQ-0604 — Build on macOS AMD64

Priority: Medium  
Status: Draft  
Milestone: Awakening

The project should build on macOS AMD64 for non-hardware-specific code where practical.

Rationale:

Some contributors may use Intel Macs.

## REQ-0605 — Hardware-specific build isolation

Priority: Critical  
Status: Draft  
Milestone: Awakening / Power

Hardware-specific implementation shall be isolated so non-hardware code can compile and test away from Raspberry Pi hardware.

Rationale:

Portability and development speed depend on clean boundaries.

## REQ-0606 — CI build validation

Priority: High  
Status: Draft  
Milestone: Awakening

The project shall use CI to validate build behavior.

Rationale:

Every commit should provide confidence that the project still builds.

## REQ-0607 — Mockable core logic

Priority: High  
Status: Draft  
Milestone: Power / Memory

Core service logic should be testable with mocks or simulated events.

Rationale:

Hardware should be required for validation, not every unit test.

---

# 11. Safety and Reliability Requirements

## REQ-0700 — Fail safely

Priority: Critical  
Status: Draft  
Milestone: All

When the system encounters an unexpected failure, it shall prefer safe fallback behavior over risky continuation.

Rationale:

The project must protect the user and their data.

## REQ-0701 — Avoid destructive defaults

Priority: Critical  
Status: Draft  
Milestone: All

The project shall avoid destructive behavior by default.

Rationale:

Users should not lose progress because of an aggressive default.

## REQ-0702 — Clear recovery path

Priority: High  
Status: Draft  
Milestone: All

When a recoverable error occurs, the project should provide clear next steps.

Rationale:

Failures should be understandable and actionable.

## REQ-0703 — Service restart policy

Priority: High  
Status: Draft  
Milestone: Heartbeat

The systemd service shall define an appropriate restart policy.

Initial preference:

```
Restart=on-failure
```

unless a later ADR chooses otherwise.

Rationale:

Crashes should not leave the power management service unavailable.

## REQ-0704 — Watchdog research

Priority: Research  
Status: Draft  
Milestone: Heartbeat / Future

systemd watchdog support should be researched after the daemon lifecycle is stable.

Rationale:

Watchdog support may improve reliability but should not complicate the first daemon milestone.

---

# 12. Product Experience Requirements

## REQ-0800 — Console-like experience

Priority: High  
Status: Draft  
Milestone: All

The project should move the handheld experience toward a console-like model.

Rationale:

The product vision is to reduce the feeling of managing a Linux computer.

## REQ-0801 — Resume continuity

Priority: High  
Status: Draft  
Milestone: Resume

The project should support returning to the previous game session where practical.

Rationale:

Continuity is a core product promise.

## REQ-0802 — Visibility for power users

Priority: High  
Status: Draft  
Milestone: Heartbeat / Power / Memory

The project should provide enough visibility for power users to inspect and troubleshoot behavior.

Rationale:

A polished user experience should not make the system opaque.

## REQ-0803 — Contributor clarity

Priority: High  
Status: Draft  
Milestone: All

The repository should explain itself well enough that future contributors can understand the project direction.

Rationale:

Maintainability depends on shared understanding.

## REQ-0804 — Moments of delight

Priority: Medium  
Status: Draft  
Milestone: Polish

The project should include small moments of delight where they improve the experience without reducing clarity.

Rationale:

The project intentionally values warmth and craftsmanship.

---

# 13. Research Requirements

## REQ-0900 — Sleep mode research

Priority: Research  
Status: Draft  
Milestone: Future

The project should research whether useful sleep-like behavior is possible on the reference hardware.

Rationale:

Full suspend may be limited by hardware, but partial pause behavior may still be useful.

## REQ-0901 — Battery capability research

Priority: Research  
Status: Draft  
Milestone: Future

The project should research whether the RetroFlag power board exposes useful battery state.

Rationale:

Battery overlays and critical shutdown features depend on available data.

## REQ-0902 — Boot order research

Priority: Research  
Status: Draft  
Milestone: Momentum

The project should research boot-order changes such as `BOOT_ORDER=0xf1` and their actual effect on boot time.

Rationale:

Firmware changes should be measured before recommendation.

## REQ-0903 — KMS/FKMS research

Priority: Research  
Status: Draft  
Milestone: Momentum

The project should research whether KMS/FKMS changes affect boot time, display behavior, or emulator performance.

Rationale:

Graphics stack changes can have compatibility and performance implications.

## REQ-0904 — Other SBC awareness

Priority: Research  
Status: Draft  
Milestone: Expansion

The project should remain aware of popular retro gaming SBCs beyond Raspberry Pi.

Examples:

- Orange Pi
- Radxa ROCK
- ODROID
- Banana Pi
- Libre Computer

Rationale:

Architecture should not unnecessarily block future community support.

---

# 14. Milestone 1 Minimum Requirements

Milestone 1 — Awakening should focus only on the first breath.

Minimum requirement set:

- REQ-0004 — Run as a supervised daemon
- REQ-0005 — Graceful stop
- REQ-0006 — Structured startup logging
- REQ-0500 — Metadata headers
- REQ-0600 — Build on Linux ARM64
- REQ-0602 — Build on Linux AMD64
- REQ-0603 — Build on macOS ARM64
- REQ-0605 — Hardware-specific build isolation
- REQ-0606 — CI build validation

Milestone 1 should not attempt:

- GPIO
- shutdown execution
- resume
- state storage
- splash screens
- sleep mode
- battery features

The first daemon should be simple.

---

# 15. Traceability Expectations

Future work should reference requirements where practical.

Examples:

```
Implements: REQ-0005
Related: REQ-0605
```

Pull requests and commits do not need to reference every requirement, but major work should connect back to the requirement that justifies it.

Tests may eventually reference requirement IDs.

ADRs should reference related requirements.

RFCs should identify requirements they propose to satisfy or create.

---

# 16. Requirement Review

This document should be reviewed:

- before Milestone 1 begins
- before GPIO implementation begins
- before resume implementation begins
- before public beta
- before v1.0

Requirements may evolve as the project learns.

Changes should be intentional.

---

# Closing

Requirements are not chains.

They are rails.

They keep the project moving toward the experience it promised.

The first laws are now written:

Protect the user.

Respect the hardware.

Preserve the dream.

Build the machine.
