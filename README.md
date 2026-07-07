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

## Project Documentation

- [Why RetroFlag Power exists](WHY.md)
- [Project Charter](PROJECT_CHARTER.md): the campfire compass for RetroFlag
  Power's working style, quest rhythm, hardware caution, and project voice.
- [Project roadmap](docs/00-project/roadmap.md)
- [Project requirements](docs/00-project/requirements.md)
- [System overview](docs/04-architecture/system-overview.md)
- [AI collaboration guide](docs/05-development/ai-collaboration.md)
