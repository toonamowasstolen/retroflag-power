# RetroFlag Power

RetroFlag Power is a modern Go daemon project for evolving RetroFlag power
management safely and deliberately.

## Current Status

The project is in Milestone 1 — Awakening. The daemon can build, run, log its
lifecycle, wait for SIGINT or SIGTERM, and exit cleanly.

It does not yet control GPIO, execute shutdown, replace `SafeShutdown.py`, or
activate a systemd service.

## Goals

- Native Go daemon
- systemd integration
- Structured logging
- Event-driven architecture
- Cross-platform development

## Development

Workshop is the local development companion:

```sh
make help
make check
make run
```

Safe dry-run power intent lantern:

```sh
go run ./cmd/retroflag-powerd --dry-run-power-button
```

This processes the dry-run `power_button_pressed` intent through the daemon app,
planner, and executor path. The current `power_button_action` policy supports
only `noop`, including the explicit CLI form:

```sh
go run ./cmd/retroflag-powerd --dry-run-power-button --power-button-action noop
```

Safe fake observer lantern:

```sh
go run ./cmd/retroflag-powerd --fake-power-button-observer
```

This starts the daemon app, emits one fake power-button observer event, routes it
through the same input observer path used by tests, prints the noop result and a
small deterministic event breadcrumb ledger, then exits cleanly.

The input lantern now also has a configured latching power switch interpreter:
raw `SignalLow` or `SignalHigh` can become `SwitchOff` or `SwitchOn` only when
`active_signal` and `active_switch_state` are explicit. `SignalUnverified`
becomes `SwitchUnknown`.

Unsupported policy values fail clearly before a plan is prepared. The path
remains noop-only: no GPIO is read, no shutdown command runs, and no hardware
action is taken.

## Project Documentation

- [Why RetroFlag Power exists](WHY.md)
- [Project Charter](PROJECT_CHARTER.md): the campfire compass for RetroFlag
  Power's working style, quest rhythm, hardware caution, and project voice.
- [Project roadmap](docs/00-project/roadmap.md)
- [Project requirements](docs/00-project/requirements.md)
- [System overview](docs/04-architecture/system-overview.md)
- [AI collaboration guide](docs/05-development/ai-collaboration.md)
