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

Safe fake raw signal lantern:

```sh
go run ./cmd/retroflag-powerd --fake-power-signal low
```

This starts the daemon app, creates one raw signal event for the configured
power input, interprets it through the configured latching power switch map, and
prints a compact deterministic result with event breadcrumbs. With the default
map, `low` becomes `SwitchOff` and reaches the existing noop power path;
`high` becomes `SwitchOn`, and `unverified` becomes `SwitchUnknown`, both
exiting cleanly without requesting shutdown behavior.

The input lantern now also has a configured latching power switch interpreter:
raw `SignalLow` or `SignalHigh` can become `SwitchOff` or `SwitchOn` only when
`active_signal` and `active_switch_state` are explicit. `SignalUnverified`
becomes `SwitchUnknown`.

Unsupported policy values fail clearly before a plan is prepared. The path
remains noop-only: no GPIO is read, no shutdown command runs, and no hardware
action is taken.

Read-only GPi Case signal lantern:

```sh
go run ./cmd/retroflag-powerd --probe-gpio-signal 4
```

Run this on the GPi Case with a candidate BCM GPIO pin number. The command only
tries to read the raw wire state and prints `SignalLow`, `SignalHigh`, or
`SignalUnverified`. It does not interpret the result as `SwitchOn` or
`SwitchOff`, does not start daemon processing, and does not request shutdown.
On unsupported platforms or uncertain GPIO access, it reports
`SignalUnverified` deterministically.

## Project Documentation

- [Why RetroFlag Power exists](WHY.md)
- [Project Charter](PROJECT_CHARTER.md): the campfire compass for RetroFlag
  Power's working style, quest rhythm, hardware caution, and project voice.
- [EDC Quest Operating Rules](docs/00-project/edc-quest-operating-rules.md):
  the source-of-truth ritual for future Codex, Claude, and human quest work.
- [Save Room Tech and Arcadia Runtime direction](docs/00-project/project-direction-save-room-arcadia.md):
  the north-star map for the current RetroFlag Power prototype, the future
  Save Room Tech umbrella, and the favored Arcadia Runtime path.
- [Arcadia Runtime migration path](docs/04-architecture/arcadia-runtime-migration-path.md):
  the staged compass for growing from the current GPi Case 2 prototype toward
  the broader runtime direction without premature renames or behavior changes.
- [Local diagnostics bundle map](docs/03-operations/local-diagnostics-bundle-map.md):
  the privacy-first compass for future local, redacted, previewable support
  bundles before any optional Lantern Dispatch trail exists.
- [Installer and migration toolkit map](docs/03-operations/installer-migration-toolkit-map.md):
  the reversible compass for future local install, update, uninstall, and
  restore trails before any field kit mutates a device.
- [GPi Case 2 acceptance checklist](docs/03-operations/gpi-case-2-acceptance-checklist.md):
  the field-test gate ledger for power, display, dock, audio, sleep/resume,
  RetroPie, diagnostics, rollback, and public readiness before replacement or
  migration.
- [Project roadmap](docs/00-project/roadmap.md)
- [Project requirements](docs/00-project/requirements.md)
- [GPi Case GPIO probe field ledger](docs/03-operations/gpi-case-gpio-probe-ledger.md)
- [SafeShutdown replacement boundary map](docs/03-operations/safeshutdown-replacement-boundary-map.md)
- [System overview](docs/04-architecture/system-overview.md)
- [AI collaboration guide](docs/05-development/ai-collaboration.md)
