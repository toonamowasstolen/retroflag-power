---
id: AWAKENING-READINESS-001
title: Awakening Readiness Check
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Confirm that RetroFlag Power is ready to move from Milestone 0 — Dreaming into Milestone 1 — Awakening without losing safety, scope control, or implementation momentum.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - ENGINEERING_MANIFESTO.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/00-project/requirements.md
  - docs/01-product/vision.md
  - docs/02-hardware/gpi-case-2.md
  - docs/04-architecture/system-overview.md
  - docs/05-development/ai-collaboration.md
  - docs/10-decisions/adr-template.md
  - docs/11-rfc/rfc-template.md
last_updated: 2026-07-03
---

# Awakening Readiness Check

> The project is exiting REM sleep. Make sure it can survive its first breath.

This document confirms whether RetroFlag Power is ready to move from:

```
Milestone 0 — Dreaming
```

to:

```
Milestone 1 — Awakening
```

Dreaming gave the project its identity, language, laws, witness testimony, architecture, and working rules.

Awakening gives the project its first living process.

The first breath must be small.

The first breath must be safe.

The first breath must not accidentally become the whole machine.

---

# 1. Readiness Verdict

## Current verdict

```
Ready to prepare Milestone 1.
```

The project is ready to begin the first implementation milestone after this readiness check is committed.

This does not mean the project is ready to replace SafeShutdown.py.

This does not mean the project is ready to touch GPIO.

This does not mean the project is ready to control power.

It means the project is ready to create the first daemon skeleton.

---

# 2. What Awakening Means

Milestone 1 — Awakening means:

```
The project takes its first breath.
```

The goal is to prove that RetroFlag Power can exist as a real process.

The first daemon should:

- compile
- run
- log startup
- wait for stop signal
- handle SIGINT
- handle SIGTERM
- log shutdown
- exit cleanly

That is enough.

A living process is the first victory.

---

# 3. What Awakening Does Not Mean

Awakening does not mean:

- GPIO support
- Power Switch handling
- Reset Button handling
- safe shutdown replacement
- disabling SafeShutdown.py
- editing rc.local
- resume support
- state storage
- RetroArch integration
- EmulationStation integration
- boot optimization
- splash screen
- hardware profile loading
- battery support
- sleep mode
- release packaging

Those belong to later milestones.

The first breath should not carry the whole dream on its lungs.

---

# 4. Completed Foundation Artifacts

The following foundation artifacts exist or are planned as committed project context.

## Root project artifacts

- [x] WHY.md
- [x] PROJECT_MEMORY.md
- [x] ENGINEERING_MANIFESTO.md
- [x] PROJECT_MANIFEST.md
- [x] PROJECT_CHARTER.md

## Project planning artifacts

- [x] docs/00-project/milestones.md
- [x] docs/00-project/roadmap.md
- [x] docs/00-project/requirements.md

## Product artifacts

- [x] docs/01-product/vision.md

## Hardware artifacts

- [x] docs/02-hardware/gpi-case-2.md

## Architecture artifacts

- [x] docs/04-architecture/system-overview.md

## Development artifacts

- [x] docs/05-development/ai-collaboration.md

## Decision and proposal artifacts

- [x] docs/10-decisions/adr-template.md
- [x] docs/11-rfc/rfc-template.md

## Reference artifacts

- [x] docs/13-reference/terminology.md
- [x] docs/13-reference/glossary.md

This is enough foundation for the first daemon.

Not every future document needs to exist before implementation starts.

---

# 5. Milestone 1 Entry Criteria

Milestone 1 can begin when the following are true:

- [x] The project purpose is documented.
- [x] The product vision is documented.
- [x] The milestone path is documented.
- [x] The practical roadmap is documented.
- [x] Initial requirements are documented.
- [x] The system overview is documented.
- [x] The reference hardware is documented.
- [x] Terminology rules are documented.
- [x] Glossary exists.
- [x] AI collaboration guidance exists.
- [x] ADR template exists.
- [x] RFC template exists.
- [x] The first daemon scope is intentionally small.
- [x] The project knows what not to implement in the first breath.

Milestone 1 entry criteria are satisfied.

---

# 6. First Breath Scope

The first code commit should add only the minimum needed to prove the daemon can live.

Recommended first implementation:

```
cmd/
  retroflag-powerd/
    main.go

internal/
  app/
    app.go

internal/
  logging/
    logging.go
```

Possible supporting files:

```
go.mod
Makefile
README update if needed
```

The first behavior:

```
start
  │
  ▼
create logger
  │
  ▼
create app context
  │
  ▼
log startup
  │
  ▼
wait for SIGINT/SIGTERM
  │
  ▼
log shutdown
  │
  ▼
exit cleanly
```

The first daemon should be boring in the best possible way.

---

# 7. First Breath Requirements

The first implementation should focus on this requirement set.

## REQ-0004 — Run as a supervised daemon

The daemon should be suitable for future systemd supervision.

## REQ-0005 — Graceful stop

The daemon should handle SIGINT and SIGTERM.

## REQ-0006 — Structured startup logging

The daemon should log startup clearly.

## REQ-0403 — Clear errors

Errors should be understandable.

## REQ-0600 — Build on Linux ARM64

The project should be able to target the reference platform.

## REQ-0602 — Build on Linux AMD64

The project should be able to build in common CI environments.

## REQ-0603 — Build on macOS ARM64

The project should support development on Apple Silicon Macs for non-hardware code.

## REQ-0605 — Hardware-specific build isolation

The first daemon must not depend on Raspberry Pi-only GPIO behavior.

## REQ-0606 — CI build validation

CI should eventually validate basic build behavior.

This may be added during or shortly after the first daemon commit.

---

# 8. First Breath Non-Requirements

The first implementation should explicitly avoid these requirements for now:

## REQ-0001 — Safe shutdown from Power Switch OFF

Not yet.

## REQ-0102 — GPIO isolation

Architecture should respect it, but no GPIO implementation yet.

## REQ-0103 — Event-driven GPIO direction

Not yet.

## REQ-0104 — libgpiod direction

Not yet.

## REQ-0200 — Session model

Not yet.

## REQ-0201 — Durable state storage

Not yet.

## REQ-0300 — Boot-to-resume target

Not yet.

## REQ-0400 — CLI status

Not yet.

The first daemon should prepare the ground, not build every room.

---

# 9. Safety Check

Before the first daemon is created, confirm:

- [x] No original shutdown behavior will be changed.
- [x] No GPIO behavior will be assumed.
- [x] No systemd unit will replace the original script yet.
- [x] No rc.local edits will be made.
- [x] No EEPROM changes will be made.
- [x] No boot-order changes will be made.
- [x] No save-state behavior will be touched.
- [x] No hardware claims will be invented.
- [x] No release or install instructions will imply production readiness.

The first daemon is safe because it does not control anything yet.

It only breathes.

---

# 10. First Commit Shape

Suggested first code commit title:

```
The project takes its first breath.
```

Suggested commit body:

```
Add the initial retroflag-powerd daemon skeleton.

Create a minimal app lifecycle with startup logging, signal handling, graceful shutdown, and a small logging boundary.

Keep the first implementation intentionally hardware-free so Milestone 1 can validate the daemon lifecycle before GPIO, systemd replacement, power handling, or resume behavior are introduced.
```

Possible files:

```
go.mod
cmd/retroflag-powerd/main.go
internal/app/app.go
internal/logging/logging.go
Makefile
```

Do not include:

```
GPIO
system shutdown execution
SafeShutdown.py replacement
hardware profile loading
resume state
RetroArch integration
EmulationStation integration
```

---

# 11. First Breath Test Plan

The first daemon should be testable with simple local commands.

Possible commands:

```bash
go run ./cmd/retroflag-powerd
```

In another terminal, send interrupt:

```bash
Ctrl+C
```

Expected behavior:

```
startup log appears
daemon waits
shutdown log appears
process exits cleanly
```

Possible build command:

```bash
go build ./cmd/retroflag-powerd
```

Possible Makefile commands:

```bash
make build
make run
```

Minimum validation:

- [ ] daemon builds
- [ ] daemon runs
- [ ] startup log appears
- [ ] Ctrl+C stops it
- [ ] shutdown log appears
- [ ] exit code is clean
- [ ] code does not import GPIO libraries
- [ ] code does not require Raspberry Pi hardware

## Nameplate Checkpoint

The first-breath daemon now identifies itself without expanding its role:

- `retroflag-powerd --version` prints `retroflag-powerd 0.1.0-dev`.
- The startup log includes the daemon name and version.
- `make version` runs the identity command through Workshop.
- The VS Code task `Workshop: version` runs `make version`.
- `make check` runs tests, build, and version validation.
- CI/Forge receives the same identity validation by running `make check`.
- Tests, build, version output, and the Ctrl+C runtime lifecycle passed.
- `make check` passed.

This is an Awakening progress note, not a new hardware or service capability.

## Config Satchel Checkpoint

Awakening now includes a tiny defaults-only configuration boundary:

- `AppName` is `retroflag-powerd`.
- `Version` is `0.1.0-dev`, sourced from the version package.
- `DryRun` is `true`.
- The app receives the config.
- Startup logging includes name, version, and `dry_run=true`.
- `--version` output remains unchanged.
- Lifecycle validation and `make check` passed.

The daemon does not load config files or environment variables and has no new
CLI flags. This checkpoint adds no GPIO, shutdown execution, service activation,
resume, or state storage.

## Event Charms Checkpoint

Awakening now includes a minimal internal lifecycle event model:

- `Event` has `Type` and `Message` fields.
- Lifecycle types cover daemon starting, daemon ready, shutdown signal received,
  and daemon stopped.
- The app logs lifecycle messages through the event model.
- `--version` is unchanged.
- Startup still logs name, version, and `dry_run=true`.
- The Ctrl+C lifecycle exits cleanly and `make check` passed.

This is not an event bus. It adds no channels, async processing, persistence,
third-party dependencies, GPIO, shutdown execution, service activation, resume,
or state storage.

---

# 12. First Breath Design Notes

The first daemon should use Go's standard building blocks where practical:

- `context`
- `os/signal`
- `syscall`
- standard logging or a small logging wrapper
- simple app struct
- clean return paths

Avoid:

- framework sprawl
- dependency-heavy logging
- daemonization libraries
- global mutable state
- early configuration complexity
- premature event bus implementation
- premature package explosion

The first breath is not the place to prove cleverness.

It is the place to prove life.

---

# 13. Suggested Initial App Shape

The app boundary may look conceptually like this:

```
main
  │
  ▼
logging.New()
  │
  ▼
app.New(logger)
  │
  ▼
app.Run(ctx)
```

The app should return errors instead of exiting from deep inside package code.

`main` should be responsible for process exit behavior.

This keeps future testing easier.

---

# 14. Suggested Initial Log Messages

Startup:

```
retroflag-powerd starting
```

Ready / waiting:

```
retroflag-powerd ready
```

Signal received:

```
shutdown signal received
```

Stopped:

```
retroflag-powerd stopped
```

Avoid too much personality in daemon logs early.

Save the banners for CLI output later.

Daemon logs should be boring, clear, and useful.

---

# 15. Future Work Immediately After First Breath

After the first daemon works, the next implementation steps should be:

## Step 1 — Makefile

If not included in the first commit, add:

```bash
make build
make run
make test
```

## Step 2 — Basic tests

Add tests for app lifecycle where practical.

## Step 3 — systemd unit draft

Add a service file, but do not replace the existing shutdown path yet.

## Step 4 — journald validation

Run service under systemd and confirm logs.

## Step 5 — first ADRs

Create ADRs for:

- daemon lifecycle
- systemd service strategy
- hardware abstraction/GPIO direction

## Step 6 — original script review

Before power implementation, review `SafeShutdown.py`.

Do not skip this.

---

# 16. What Would Make Us Not Ready

The project should pause before Milestone 1 if:

- the user wants more foundational docs before code
- current repo state is unknown or broken
- the bootstrap code skeleton conflicts with this plan
- the language/runtime decision is not actually settled
- there are uncommitted docs that should be saved first
- the user is not ready to touch code yet

Otherwise, the project is ready.

---

# 17. Pre-Awakening Repo Checklist

Before the first code commit:

- [ ] `git status` is clean or understood.
- [ ] Foundation docs are committed.
- [ ] Current branch is correct.
- [ ] Existing skeleton files are reviewed.
- [ ] `.DS_Store` is ignored.
- [ ] No generated artifacts are accidentally staged.
- [ ] First code commit scope is agreed.
- [ ] No hardware behavior is changed.

---

# 18. First Breath Motto

```
No GPIO.
No shutdown.
No resume.
No magic.

Just life.
```

---

# 19. Readiness Summary

The project is ready to begin Milestone 1 because:

- the spark is documented
- the vision is documented
- the path is documented
- the laws are documented
- the architecture has a first shape
- the reference hardware has testified
- the builders have rules
- decisions and proposals have chambers
- the first daemon scope is intentionally small

The dream is no longer only dreaming.

It is waking.

---

# Closing

Milestone 0 gave the project a mind.

Milestone 1 gives it breath.

Do not make the first breath heroic.

Make it steady.

The first daemon does not need to save the world.

It needs to inhale, exhale, and prove it can keep going.
