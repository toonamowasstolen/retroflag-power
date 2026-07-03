---
id: ADR-0002
title: Use a Small Context-Driven Daemon Lifecycle
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record the decision to begin RetroFlag Power with a small context-driven Go daemon lifecycle using main, internal/app, and internal/logging boundaries before adding hardware, shutdown, resume, or state behavior.
related:
  - docs/00-project/awakening-readiness.md
  - docs/00-project/requirements.md
  - docs/04-architecture/system-overview.md
  - docs/05-development/ai-collaboration.md
  - docs/10-decisions/adr-template.md
last_updated: 2026-07-03
---

# ADR-0002 — Use a Small Context-Driven Daemon Lifecycle

> The first breath should be steady, not heroic.

## Status

Draft

## Date

2026-07-03

## Owner

Joshua Taft

## Related Requirements

- REQ-0004 — Run as a supervised daemon
- REQ-0005 — Graceful stop
- REQ-0006 — Structured startup logging
- REQ-0403 — Clear errors
- REQ-0600 — Build on Linux ARM64
- REQ-0602 — Build on Linux AMD64
- REQ-0603 — Build on macOS ARM64
- REQ-0605 — Hardware-specific build isolation
- REQ-0606 — CI build validation

## Related Documents

- docs/00-project/awakening-readiness.md
- docs/00-project/requirements.md
- docs/04-architecture/system-overview.md
- docs/05-development/ai-collaboration.md

---

# 1. Context

RetroFlag Power is moving from:

```
Milestone 0 — Dreaming
```

into:

```
Milestone 1 — Awakening
```

The project has already defined its purpose, product vision, requirements, architecture overview, hardware reference, collaboration rules, and readiness criteria.

The first implementation goal is intentionally small:

```
Start.
Log.
Wait.
Stop cleanly.
```

The project must not accidentally begin with GPIO, shutdown execution, SafeShutdown.py replacement, rc.local edits, resume behavior, state storage, RetroArch integration, or EmulationStation integration.

The first daemon exists to prove that the project can run as a real process before it controls real hardware.

---

# 2. Decision

RetroFlag Power will begin with a small context-driven Go daemon lifecycle.

The initial implementation uses this structure:

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

The entry point is responsible for:

- creating the root context
- wiring SIGINT and SIGTERM into context cancellation
- constructing the logger
- constructing the app
- calling `App.Run(ctx)`

The app is responsible for:

- logging startup
- logging readiness
- waiting for context cancellation
- logging shutdown signal receipt
- logging clean stop

The logging package is responsible for:

- creating the initial standard logger boundary

No hardware-specific behavior is included in this decision.

No shutdown behavior is included in this decision.

No resume or state behavior is included in this decision.

---

# 3. Rationale

This decision supports Milestone 1 by creating the smallest useful daemon shape.

It gives the project a real process without creating unnecessary architecture.

It keeps the entry point simple.

It creates a natural place for future lifecycle behavior.

It creates a testable app boundary.

It avoids hardware coupling.

It validates that the project can build and run on macOS before Raspberry Pi-specific code is introduced.

It gives future services a place to attach without forcing them to exist before they are needed.

The first breath should prove life, not power.

---

# 4. Alternatives Considered

## Alternative 1 — Keep all lifecycle logic in main.go

Description:

Keep startup logging, signal handling, waiting, and shutdown logging entirely inside `cmd/retroflag-powerd/main.go`.

Pros:

- smallest possible code
- very easy to read
- no internal packages needed

Cons:

- hard to unit test without OS signal behavior
- main.go would grow quickly
- future app lifecycle logic would have no clear home
- encourages command entry point to become application logic

Reason not chosen:

This was acceptable for the first breath, but once the daemon proved it could run, the project needed a small internal boundary to support tests and future growth.

## Alternative 2 — Build a full service framework immediately

Description:

Introduce app lifecycle management, service registration, event bus, config loading, structured logger, health state, and future component interfaces immediately.

Pros:

- prepares for future architecture
- may reduce some later refactoring
- gives many components a place to plug in

Cons:

- overbuilds before behavior exists
- risks planning paralysis in code form
- introduces complexity before GPIO and systemd behavior are validated
- makes the first daemon harder to understand
- violates the Awakening readiness scope

Reason not chosen:

The project intentionally avoids building the whole machine before the first living process is stable.

## Alternative 3 — Start with systemd service behavior first

Description:

Focus immediately on service installation, unit files, and running under systemd before refining the internal daemon lifecycle.

Pros:

- validates the Linux service path early
- moves toward Heartbeat milestone
- aligns with the original replacement goal

Cons:

- risks touching deployment before daemon lifecycle is shaped
- may distract from testability
- may create install confidence before behavior is ready
- not useful on macOS development machine

Reason not chosen:

systemd integration matters, but the daemon should have a clean lifecycle before the service wrapper becomes important.

---

# 5. Consequences

## Positive Consequences

- The daemon has a small, understandable structure.
- `main.go` remains an entry point instead of application logic.
- `internal/app` can be tested without OS signals.
- `internal/logging` provides a simple future extension point.
- The project remains hardware-free in Milestone 1.
- The project can build and test on macOS.
- The project has a clean foundation for future systemd work.

## Negative Consequences

- There is a small amount of package structure before much behavior exists.
- `App.Run` currently has no real error path.
- Logging is still basic standard library logging.
- The lifecycle is intentionally minimal and will need to evolve.

## Neutral Consequences

- Future services may change the app constructor.
- Future logging may become structured.
- Future lifecycle behavior may return errors.
- The current structure may be refined as real behavior appears.

---

# 6. Risks

## RISK-0001 — Premature package growth

The existence of internal packages could tempt the project to add many abstractions too early.

Mitigation:

Keep internal boundaries minimal until behavior requires more structure.

## RISK-0002 — Logger boundary grows too soon

The logging package could become a premature framework.

Mitigation:

Use the standard library logger until real logging needs justify more complexity.

## RISK-0003 — App lifecycle hides errors later

`App.Run(ctx)` currently does not return an error.

Mitigation:

Add an error return when real startup, service, hardware, or state initialization failures exist.

## RISK-0004 — Future signal handling assumptions

Signal handling currently lives in `main`.

Mitigation:

Keep OS process concerns at the entry point and unit test application behavior through context cancellation.

---

# 7. Validation Plan

This decision has been validated through the current Awakening chain:

- daemon builds locally
- daemon runs locally
- daemon logs startup
- daemon logs ready
- daemon handles Ctrl+C
- daemon logs shutdown signal receipt
- daemon logs stopped
- daemon exits cleanly
- lifecycle unit test passes
- Makefile Workshop `make check` passes
- VS Code Workshop tasks exist
- GitHub Actions Forge runs `make check`
- latest CI run passed

Current green chain:

```
The project takes its first breath.
Teach the module its true name.
Teach the daemon to breathe through its bones.
Teach the breath to prove itself.
Summon Workshop as the dream's companion.
Open a second door to Workshop.
Teach the forge to use Workshop.
```

---

# 8. Rollback or Revision Plan

If this structure becomes unhelpful, the project can revise it before hardware behavior is added.

Possible revisions:

- collapse `internal/logging` back into `main` if it remains unnecessary
- update `App.Run(ctx)` to return `error`
- introduce a richer logger only when justified
- introduce service registration only when multiple services exist
- add lifecycle hooks only when real startup/shutdown ordering exists

No rollback affects hardware behavior because no hardware behavior has been added yet.

---

# 9. Notes

The first tested lifecycle currently verifies the expected log sequence:

```
retroflag-powerd starting
retroflag-powerd ready
shutdown signal received
retroflag-powerd stopped
```

The test uses:

- canceled context
- in-memory log buffer
- no OS signals
- no sleeps
- no hardware
- no third-party dependencies

This matches the Awakening rule:

```
No GPIO.
No shutdown.
No resume.
Just life.
```

With one clarification:

A little magic belongs in the game.

But not in the shutdown path until the heartbeat is stable.

---

# 10. Outcome

Pending.

This ADR should be updated when Milestone 1 completes or when the daemon lifecycle changes meaningfully.

Possible future updates:

- mark Accepted once reviewed
- reference first lifecycle test commit
- reference CI pass
- supersede if lifecycle architecture changes before Milestone 2
- update if `App.Run` begins returning errors

---

# Closing

The first daemon does not need to be powerful.

It needs to be alive.

This decision gives RetroFlag Power a small body, a clean breath, and enough bones to keep growing.

The dream has entered the waking world.

Now its first breath has a memory crystal.
