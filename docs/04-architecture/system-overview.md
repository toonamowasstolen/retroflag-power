---
id: ARCH-SYSTEM-OVERVIEW-001
title: System Overview
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Hardware Porters
purpose: Define the high-level system architecture for RetroFlag Power so the first daemon, future services, hardware abstraction, state handling, and resume work can be implemented with clear boundaries.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/00-project/requirements.md
  - docs/00-project/quests/0035-add-a-dry-run-power-intent-path.md
  - docs/00-project/quests/0036-add-a-dry-run-power-intent-cli-flag.md
  - docs/00-project/quests/0038-add-a-configurable-dry-run-power-policy.md
  - docs/00-project/quests/0039-add-a-gpio-observer-interface.md
  - docs/01-product/vision.md
  - docs/13-reference/terminology.md
  - docs/13-reference/glossary.md
last_updated: 2026-07-07
---

# System Overview

> The laws now have a courthouse.

This document defines the high-level architecture for RetroFlag Power.

It is not the final design for every package, interface, or command.

It is the first courthouse: the place where the project requirements, product vision, hardware reality, and implementation path can meet without chaos.

The goal is to give Milestone 1 enough structure to begin implementation while keeping future milestones visible.

---

# 1. Architectural Goal

RetroFlag Power should begin as a small, reliable Linux daemon with clear internal boundaries.

The first daemon should prove that the project can:

- build
- run
- log
- stop cleanly
- run under systemd
- remain portable for non-hardware development
- grow into power, state, resume, metrics, and polish features

The architecture should protect the project from two failures:

## Failure 1 — A script with too much responsibility

The original `SafeShutdown.py` behavior solves a real problem, but RetroFlag Power needs a more maintainable foundation than a background script launched from `rc.local`.

## Failure 2 — A platform with too much architecture too early

The project should not build a plugin framework, multi-daemon ecosystem, or complex IPC design before the first daemon breathes.

The first architecture should be strong enough to grow and small enough to build.

---

# 2. Architectural Style

RetroFlag Power should begin as a:

```
Modular monolith daemon
```

That means:

- one main daemon process
- clear internal services
- explicit interfaces
- no unnecessary IPC
- hardware-specific code isolated behind boundaries
- mocks for non-hardware development
- future expansion possible without early complexity

The initial daemon is expected to be:

```
retroflag-powerd
```

The future user-facing CLI may be:

```
retroflag-power
```

---

# 3. System Context

At the highest level, RetroFlag Power sits between Linux, hardware, and the retro gaming stack.

```
Player
  │
  ▼
RetroFlag GPi Case 2 controls
  │
  ▼
Linux GPIO / system services
  │
  ▼
retroflag-powerd
  │
  ├── systemd
  ├── journald
  ├── RetroPie
  ├── EmulationStation
  ├── RetroArch
  └── state storage
```

The player should not experience this complexity.

The player should experience:

```
Power Switch ON
      │
      ▼
Continue the adventure
```

---

# 4. Reference Platform

The first-class reference platform is:

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

This architecture should support the reference platform first.

It should avoid decisions that unnecessarily block future hardware profiles.

Validated support and experimental support must remain clearly distinguished.

---

# 5. Major System Components

The initial architecture contains these major components:

```
systemd
  │
  ▼
retroflag-powerd
  │
  ├── App Lifecycle
  ├── Event Bus
  ├── Power Service
  ├── Hardware Service
  ├── State Service
  ├── Resume Service
  ├── Metrics Service
  ├── Configuration Service
  ├── Frontend Service
  ├── Emulator Service
  ├── Terminal UI Service
  └── Logging
```

Not all components need to be fully implemented in Milestone 1.

Milestone 1 should mostly establish the daemon lifecycle and enough structure for these components to exist later without guesswork.

---

# 6. Component Responsibilities

## 6.1 App Lifecycle

Responsible for:

- daemon startup
- daemon shutdown
- signal handling
- context cancellation
- dependency initialization
- service start order
- service stop order

Milestone 1 focus:

- start
- log
- handle SIGINT/SIGTERM
- stop cleanly

Startup result and diagnostic ordering:

`App.StartupResult()` is the app's small startup-completion badge. It records
whether startup has completed and whether that completion succeeded.
`App.StartupSucceeded()` reads `App.StartupResult().Succeeded`.

`App.StartupDiagnostic()` is the captured startup-complete diagnostic snapshot.
It is taken from the runtime diagnostic path after the app reaches ready-state
success. In the current dry-run/noop path, `StartupResult` and
`StartupDiagnostic` are both established during the same successful startup
completion flow.

The startup result is not yet a detailed failure taxonomy or recovery report.
If a future quest lets startup diagnostic capture fail independently, this
ordering may need to become more explicit or carry additional result detail so
the badge and diagnostic lantern do not blur into one signal.

Related requirements:

- REQ-0004
- REQ-0005
- REQ-0006

## 6.2 Event Bus

Responsible for:

- distributing meaningful events between services
- decoupling hardware events from service actions
- supporting mock event testing

Example events:

```
PowerSwitchChanged
ResetButtonPressed
ShutdownRequested
SessionRecorded
ResumeRequested
ResumeFailed
MetricRecorded
```

Milestone 1 focus:

The event bus may be stubbed or lightly designed.

It does not need full behavior until Power and Memory milestones.

Related requirements:

- REQ-0102
- REQ-0607

## 6.3 Power Service

Responsible for:

- handling power-related events
- coordinating safe shutdown
- applying shutdown policy
- ensuring power behavior is logged clearly

Example responsibility:

```
Power Switch OFF event
      │
      ▼
Power Service
      │
      ▼
Shutdown Manager
      │
      ▼
System shutdown
```

Milestone 1 focus:

The first safe power-intent path exists as dry-run/noop behavior only. A small
internal input observer can report a `PowerButtonPressed`-style event into the
app. The app converts that input event into the existing internal
`PowerButtonPressed` intent, asks the planner for a deterministic dry-run plan,
and sends that plan through the executor. The planned action is guided by the
`power_button_action` config policy. The first supported policy value is
`noop`, and executor results remain dry-run/noop only. As it travels, the app
records small deterministic breadcrumb events so future GPIO, shutdown, and
service behavior has a clear ledger to follow.

This is a lantern on the future power trail, not real hardware control. It does
not read GPIO, run shutdown commands, activate systemd services, replace
`rc.local`, replace `SafeShutdown.py`, resume sessions, or store persistent
state.

Developers can invoke this first safe path from the daemon command line:

```sh
go run ./cmd/retroflag-powerd --dry-run-power-button
```

The command starts the app lifecycle, processes the dry-run power-button intent,
prints a deterministic noop result, and exits cleanly.

The explicit policy form is also safe:

```sh
go run ./cmd/retroflag-powerd --dry-run-power-button --power-button-action noop
```

Developers can also exercise the input observer seam without real GPIO:

```sh
go run ./cmd/retroflag-powerd --fake-power-button-observer
```

That command starts the same app lifecycle, emits one fake
`power_button_pressed` observer event, routes it through the app input
processing path, prints a compact noop result, includes deterministic event
breadcrumbs, and exits cleanly.

Unsupported values fail before plan preparation with a deterministic
`power_button_action` error.

Milestone 3 focus:

Core implementation.

Related requirements:

- REQ-0001
- REQ-0002
- REQ-0010
- REQ-0700

## 6.4 Hardware Service

Responsible for:

- loading hardware profile
- exposing capabilities
- connecting hardware adapters
- converting implementation details into domain events

Example capabilities:

```
Power Switch
Reset Button
Battery
Backlight
Status LED
```

Milestone 1 focus:

The daemon now has a safe input observer seam before real hardware input. The
`internal/input` observer interface reports project-level input events rather
than GPIO pins. Its fake observer can emit a power-button-style event for tests,
which then travels through the existing config policy, planner, executor, and
event breadcrumb flow.

The fake observer is also available from the daemon command line as a workshop
field-kit command:

```sh
go run ./cmd/retroflag-powerd --fake-power-button-observer
```

The command is intentionally deterministic: one fake event enters the observer
path, the current noop power policy is honored, stdout records the result and
breadcrumb ledger, and stderr keeps the usual lifecycle logs.

This observer is intentionally a tiny field kit. It does not use Raspberry Pi
GPIO libraries, read real pins, debounce edges, model latching switches, run
shutdown commands, touch systemd, replace `rc.local`, replace `SafeShutdown.py`,
resume sessions, or persist state.

Milestone 3 focus:

GPi Case 2 profile and GPIO-backed implementation.

Related requirements:

- REQ-0100
- REQ-0101
- REQ-0102
- REQ-0105
- REQ-0106
- REQ-0107

## 6.5 State Service

Responsible for:

- durable state
- atomic writes
- validation
- recovery from corrupt or stale state
- state schema evolution

Initial future state examples:

```
last power event
active hardware profile
last known game session
last resume result
last benchmark result
```

Milestone 1 focus:

Not implemented beyond possible interface planning.

Milestone 4 focus:

Core implementation.

Related requirements:

- REQ-0200
- REQ-0201
- REQ-0202
- REQ-0204

## 6.6 Resume Service

Responsible for:

- preparing resume behavior
- validating resume state
- launching previous session
- coordinating emulator/frontend behavior
- handling resume failure safely

Milestone 1 focus:

Not implemented.

Milestone 5 focus:

Core implementation.

Related requirements:

- REQ-0203
- REQ-0205
- REQ-0206
- REQ-0801

## 6.7 Metrics Service

Responsible for:

- recording measurements
- reporting timing
- supporting benchmark output
- distinguishing actual speed from perceived speed

Possible metrics:

```
daemon startup time
service ready time
shutdown response time
boot-to-resume time
resume success/failure
idle CPU
memory usage
```

Milestone 1 focus:

Maybe startup logging only.

Milestone 6 focus:

Core implementation.

Related requirements:

- REQ-0300
- REQ-0302
- REQ-0303
- REQ-0304
- REQ-0305

## 6.8 Configuration Service

Responsible for:

- loading configuration
- validating configuration
- providing defaults
- avoiding configuration sprawl

Milestone 1 focus:

Avoid complex configuration.

Early daemon behavior should prefer simple constants or minimal config until
real variability appears. The dry-run power-button path now carries a tiny
`power_button_action` policy compass; it supports only `noop` today and exists
to mark a safe future seam without adding GPIO, shutdown, service activation,
resume, or persistent state behavior.

Related requirements:

- REQ-0701

## 6.9 Frontend Service

Responsible for:

- abstracting frontend behavior
- integrating with EmulationStation initially
- tracking game launch and exit where practical

Milestone 1 focus:

Not implemented.

Milestone 4 or 5 focus:

Research and integration.

Related requirements:

- REQ-0200
- REQ-0801

## 6.10 Emulator Service

Responsible for:

- abstracting emulator-specific behavior
- integrating with RetroArch initially
- handling save-state and resume behavior

Milestone 1 focus:

Not implemented.

Milestone 5 focus:

Core resume implementation.

Related requirements:

- REQ-0206
- REQ-0801

## 6.11 Terminal UI Service

Responsible for:

- consistent command output
- status display
- diagnostics display
- benchmark display
- tasteful project personality

Milestone 1 focus:

Minimal logs only.

Later focus:

CLI and polished terminal UX.

Related requirements:

- REQ-0400
- REQ-0401
- REQ-0402
- REQ-0403
- REQ-0404
- REQ-0405
- REQ-0406

## 6.12 Logging

Responsible for:

- clear daemon logs
- useful startup and shutdown messages
- future structured fields
- journald compatibility

Milestone 1 focus:

Startup and shutdown logs.

Related requirements:

- REQ-0006
- REQ-0007
- REQ-0403

---

# 7. Boundary Rules

## Rule 1 — Core services do not know GPIO

Core services should not know:

- GPIO chip names
- GPIO line numbers
- active-low implementation
- pull-up or pull-down details
- libgpiod calls

Core services should know:

```
Power Switch changed to OFF.
```

Not:

```
GPIO17 went low.
```

## Rule 2 — Hardware adapters translate reality into events

Hardware-specific code is responsible for converting raw hardware behavior into project events.

```
GPIO edge
  │
  ▼
GPIO Adapter
  │
  ▼
Hardware Service
  │
  ▼
PowerSwitchChanged event
```

## Rule 3 — User-facing language stays human

The player should see:

```
Power Switch OFF detected.
```

Not:

```
gpiochip0 line 17 falling edge detected.
```

Low-level details belong in debug logs or diagnostics.

`RuntimeSnapshotSummary.String()` is a stable internal diagnostic charm. It
should stay compact, deterministic, test-friendly, and machine-ish so future
tests and diagnostic lanterns can compare it without ceremony.

Future user-facing diagnostics should get a separate friendly formatter instead
of changing that stable string. Do not wire the runtime summary string to CLI
output or daemon logs until a later quest deliberately chooses that path.

Runtime and startup diagnostics are related, but they answer different
questions. `App.RuntimeDiagnostic()` reports the current runtime state whenever
it is called. `App.StartupDiagnostic()` returns the diagnostic captured when
startup completed, and that startup diagnostic remains the startup-complete
snapshot even after shutdown changes the current runtime badge to stopped.

Neither diagnostic is wired to daemon logs or CLI output yet. Future log and
CLI lanterns should choose deliberately between the current runtime state and
the captured startup snapshot, then format any user-facing output through a
separate friendly formatter rather than changing
`RuntimeSnapshotSummary.String()`.

## Rule 4 — Platform-specific code stays isolated

Raspberry Pi specific behavior must not become a global assumption.

The reference platform is first-class.

It is not the whole universe.

## Rule 5 — Safety decisions are explicit

Any decision that affects shutdown, save-state safety, state writing, or resume fallback should be documented.

Use ADRs where appropriate.

---

# 8. Initial Package Direction

A possible Go package layout:

```
cmd/
  retroflag-powerd/

internal/
  app/
  events/
  power/
  hardware/
  state/
  resume/
  metrics/
  config/
  frontend/
  emulator/
  logging/
  ui/
  systemd/
```

This is a direction, not a final commandment.

Milestone 1 may start smaller.

A first implementation might only need:

```
cmd/
  retroflag-powerd/

internal/
  app/
  logging/
```

Then grow as behavior appears.

---

# 9. Milestone 1 Architecture

Milestone 1 is called:

```
Awakening
```

Its goal:

```
The project takes its first breath.
```

Milestone 1 should implement:

- daemon entry point
- app lifecycle
- signal handling
- startup log
- shutdown log
- basic build command
- basic run command
- initial CI build if practical

Milestone 1 should not implement:

- GPIO
- Power Switch handling
- shutdown execution
- Reset Button handling
- resume
- state storage
- frontend integration
- emulator integration
- splash screen
- complex CLI
- hardware profile loading

Milestone 1 architectural shape:

```
main.go
  │
  ▼
App
  │
  ├── Logger
  ├── Context
  └── Signal Handler
```

This is enough.

The first daemon does not need to be powerful.

It needs to be alive.

---

# 10. Milestone 2 Architecture

Milestone 2 is called:

```
Heartbeat
```

Its goal:

```
The daemon becomes a reliable service.
```

Milestone 2 should introduce:

- systemd unit
- restart policy
- journald behavior
- install instructions
- service status documentation
- possibly a minimal status command or health state

Possible systemd shape:

```
[Unit]
Description=RetroFlag Power daemon
After=multi-user.target

[Service]
Type=simple
ExecStart=/usr/local/bin/retroflag-powerd
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

This file will need implementation review before use.

It is shown here as architectural intent, not final installation guidance.

---

# 11. Milestone 3 Architecture

Milestone 3 is called:

```
Power
```

Its goal:

```
The daemon learns the Power Switch.
```

Milestone 3 should introduce:

- Hardware Service
- hardware profile
- GPIO adapter
- mock GPIO adapter
- Power Service
- Shutdown Manager
- power events
- safe shutdown policy
- original script retirement docs

Power event flow:

```
Power Switch moves to OFF
      │
      ▼
GPIO Adapter
      │
      ▼
Hardware Service
      │
      ▼
PowerSwitchChanged
      │
      ▼
Event Bus
      │
      ▼
Power Service
      │
      ▼
Shutdown Manager
      │
      ▼
Linux shutdown
```

---

# 12. Milestone 4 Architecture

Milestone 4 is called:

```
Memory
```

Its goal:

```
The handheld begins to remember.
```

Milestone 4 should introduce:

- State Service
- session model
- state file
- atomic writes
- validation
- current/last session recording
- frontend/emulator research adapters

Possible state flow:

```
Game launched
    │
    ▼
Frontend / Emulator Adapter
    │
    ▼
SessionRecorded
    │
    ▼
State Service
    │
    ▼
Durable state file
```

---

# 13. Milestone 5 Architecture

Milestone 5 is called:

```
Resume
```

Its goal:

```
The adventure continues.
```

Milestone 5 should introduce:

- Resume Service
- resume validation
- RetroArch integration
- frontend coordination
- fallback behavior
- resume status
- resume clear/disable controls

Resume flow:

```
Boot complete
    │
    ▼
Resume Service
    │
    ▼
State Service reads previous session
    │
    ▼
Validate ROM / emulator / save state
    │
    ▼
Launch previous session
    │
    ▼
Restore save state
    │
    ▼
Record resume result
```

---

# 14. Data and State Direction

State should be:

- durable
- atomic where critical
- validated before use
- recoverable
- human-inspectable where practical
- documented

Potential future storage areas:

```
/var/lib/retroflag-power/
/var/lib/retroresume/
```

The final path should be decided before implementation.

Current requirement document mentions a candidate:

```
/var/lib/retroresume/state.json
```

This is not final.

An ADR should decide naming and state path before state implementation begins.

---

# 15. Configuration Direction

Early configuration should be minimal.

Avoid building a large config system before the project knows what actually varies.

Initial configuration may eventually include:

- power button action policy
- active hardware profile
- resume enabled
- log level
- state path
- hardware adapter settings
- terminal color behavior
- debug mode

Configuration should be:

- documented
- validated
- safe by default
- easy to inspect

---

# 16. Logging Direction

Logs should be clear first, structured second, and charming only when appropriate.

Startup example:

```
retroflag-powerd starting
version=0.1.0
profile=gpi-case-2
mode=daemon
```

Shutdown example:

```
retroflag-powerd stopping
reason=signal
signal=SIGTERM
```

Power event example:

```
Power Switch OFF detected
action=safe-shutdown-requested
```

Errors should be actionable.

Bad:

```
failed
```

Better:

```
Failed to open GPIO chip gpiochip0. Check hardware profile and permissions.
```

---

# 17. CLI Direction

The future CLI may include:

```
retroflag-power status
retroflag-power doctor
retroflag-power benchmark
retroflag-power resume
retroflag-power clear
retroflag-power version
```

The CLI may be the same binary as the daemon or a separate command.

This should be decided later.

Do not block Milestone 1 on the CLI design.

---

# 18. Testing Direction

The project should support tests without physical hardware.

Testing layers:

## Unit tests

Test pure logic.

## Mock hardware tests

Simulate Power Switch and Reset Button events.

## Integration tests

Test service wiring and state behavior.

## Hardware validation

Run on the actual GPi Case 2 reference platform.

Mock tests validate logic.

Hardware validation validates reality.

---

# 19. Risk Areas

## Safe shutdown risk

The project must not make shutdown less safe than the original script.

## GPIO behavior risk

Active-low/active-high details may be easy to misread.

## Resume risk

Save-state and ROM handling can risk user progress if implemented carelessly.

## Architecture risk

The project can overdesign itself before implementation.

## Scope risk

The dream can keep generating paths.

The roadmap should keep the first route visible.

---

# 20. Decisions Needed Soon

Before or during early implementation, the project should decide:

## ADR — Daemon language and lifecycle

Confirm Go daemon structure, lifecycle, and signal behavior.

## ADR — systemd service strategy

Define service name, unit file, install location, restart behavior, and logging expectations.

## ADR — Hardware abstraction and GPIO direction

Define how GPIO adapters, capabilities, and hardware profiles interact.

## ADR — State storage namespace

Decide whether state belongs under:

```
/var/lib/retroflag-power/
```

or:

```
/var/lib/retroresume/
```

or another project namespace.

## ADR — CLI binary strategy

Decide whether daemon and CLI are one binary with subcommands or separate binaries.

This can wait until CLI work is closer.

---

# 21. First Implementation Shape

The first implementation should be intentionally small.

Possible file shape:

```
cmd/
  retroflag-powerd/
    main.go

internal/
  app/
    app.go
  logging/
    logging.go
```

First daemon behavior:

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

First daemon should not know about:

- GPi Case 2
- GPIO
- RetroPie
- RetroArch
- EmulationStation
- save states
- boot metrics
- terminal banners

The first breath should be calm.

---

# 22. Architecture Summary

The system begins as:

```
systemd-supervised daemon
```

It grows into:

```
event-driven handheld power and resume platform
```

The core architectural promises are:

- safe shutdown comes first
- core services do not know GPIO
- hardware becomes capabilities
- capabilities produce events
- events drive services
- state is durable and validated
- resume fails safely
- logs are readable
- tests do not require hardware unless validating hardware
- delight supports clarity

---

# Closing

The laws have a courthouse now.

The court is not finished.

The walls are not painted.

The clerk has not found the coffee pot.

But the rooms are named.

The doors exist.

The first case can be heard.

Next stop:

```
Milestone 1 — Awakening
```

But before the daemon breathes, the reference hardware should be documented.

The courthouse has been built.

Now bring in the first witness: the GPi Case 2.
