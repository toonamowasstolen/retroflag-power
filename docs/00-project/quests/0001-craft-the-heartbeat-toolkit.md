---
id: QUEST-0001
title: Craft the Heartbeat Toolkit
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Define the first implementation quest after the daemon's first breath, focused on reviewing and preparing the future systemd heartbeat path without activating shutdown behavior or touching hardware.
related:
  - docs/00-project/awakening-readiness.md
  - docs/00-project/milestones.md
  - docs/00-project/roadmap.md
  - docs/00-project/requirements.md
  - docs/04-architecture/system-overview.md
  - docs/adr/0002-use-small-context-driven-daemon-lifecycle.md
last_updated: 2026-07-03
---

# QUEST-0001 — Craft the Heartbeat Toolkit

> The dream breathes. Now craft the toolkit that will one day help it keep breathing after boot.

## Quest Status

Draft

## Milestone

Milestone 1 — Awakening

Bridges toward:

Milestone 2 — Heartbeat

## Quest Type

Implementation preparation

## Quest Owner

Joshua Taft

---

# 1. Quest Summary

RetroFlag Power now has a minimal daemon that can:

- build
- run
- log startup
- log readiness
- wait for cancellation
- log shutdown
- exit cleanly
- pass tests locally
- pass CI through Workshop

The next quest is to review and prepare the future systemd service path without activating it yet.

This quest should inspect and refine:

```
packaging/retroflag-power.service
packaging/install.sh
```

The goal is not to install or enable the service.

The goal is to make sure the future Heartbeat path is understandable, safe, documented, and aligned with the current daemon.

---

# 2. Quest Objective

Prepare the systemd heartbeat toolkit for future validation.

The service and install script should be reviewed for:

- correct binary name
- correct service name
- clear description
- safe defaults
- no premature shutdown behavior
- no rc.local changes
- no SafeShutdown.py replacement
- no automatic enable/start unless explicitly intended
- understandable comments
- compatibility with the current daemon path
- future Raspberry Pi deployment expectations

This quest does not make the daemon a production service yet.

It prepares the toolkit.

---

# 3. Why This Quest Comes First

The project has already proven:

```
daemon breathes locally
daemon proves itself in tests
Workshop runs local rituals
VS Code can summon Workshop
GitHub Actions Forge runs Workshop
```

The next natural risk boundary is systemd.

The original project goal was to replace a legacy rc.local script with a modern supervised service.

But touching systemd too quickly could create risk.

So this quest focuses on preparation, not activation.

---

# 4. Scope

## In Scope

- Review `packaging/retroflag-power.service`
- Review `packaging/install.sh`
- Align names with the current daemon
- Add comments if helpful
- Make install behavior safe and explicit
- Ensure the service does not claim production readiness
- Ensure the install script does not disable SafeShutdown.py
- Ensure the install script does not edit rc.local
- Ensure the install script does not start or enable the service unless intentionally gated
- Update documentation comments if needed
- Run local validation where possible
- Keep changes small

## Out of Scope

- GPIO
- Power Switch handling
- Reset Button handling
- executing shutdown
- disabling SafeShutdown.py
- editing `/etc/rc.local`
- installing the service on the Raspberry Pi
- enabling the service
- starting the service
- daemonizing manually
- resume behavior
- state storage
- RetroArch integration
- EmulationStation integration
- boot optimization
- release packaging

---

# 5. Safety Rules

The quest must obey these rules:

```
No GPIO.
No shutdown execution.
No rc.local edits.
No SafeShutdown.py replacement.
No service activation.
No resume.
No state storage.
```

The service toolkit may be prepared.

The service must not take control of power behavior yet.

---

# 6. Acceptance Criteria

This quest is complete when:

- [ ] `packaging/retroflag-power.service` has been reviewed.
- [ ] `packaging/install.sh` has been reviewed.
- [ ] Service name and binary path expectations are clear.
- [ ] The service file does not imply shutdown behavior exists yet.
- [ ] The install script does not disable the original shutdown path.
- [ ] The install script does not edit rc.local.
- [ ] The install script does not start or enable the service unless explicitly designed as a later gated step.
- [ ] Any changes are small and focused.
- [ ] `make check` passes.
- [ ] No production daemon behavior changes are made unless required for naming/path alignment.
- [ ] The diff is reviewed before commit.

---

# 7. Files to Inspect

Primary files:

```
packaging/retroflag-power.service
packaging/install.sh
```

Supporting files:

```
Makefile
cmd/retroflag-powerd/main.go
internal/app/app.go
internal/logging/logging.go
docs/00-project/requirements.md
docs/04-architecture/system-overview.md
docs/00-project/awakening-readiness.md
```

---

# 8. Suggested Codex Mission

Use this prompt for Codex:

```
Milestone 1 — Awakening / Quest 0001: Craft the Heartbeat Toolkit.

Goal:
Review and safely refine the future systemd service toolkit for retroflag-powerd without activating it or expanding daemon behavior.

Primary files:
- packaging/retroflag-power.service
- packaging/install.sh

Requirements:
- Keep scope limited to service packaging preparation.
- Do not add GPIO.
- Do not execute shutdown.
- Do not edit rc.local.
- Do not disable or replace SafeShutdown.py.
- Do not start or enable the systemd service.
- Do not add resume.
- Do not add state storage.
- Do not add RetroArch or EmulationStation integration.
- Do not add release packaging complexity.

Review for:
- correct binary name: retroflag-powerd
- safe service description
- clear systemd behavior
- appropriate install path expectations
- safe install script behavior
- comments or warnings that prevent premature activation
- alignment with current Makefile and daemon output

Validation:
- make check
- shellcheck packaging/install.sh if shellcheck is already available; do not add it as a required dependency yet
- git diff

Show the diff before committing.
```

---

# 9. Suggested Commit

Commit title:

```
Craft the heartbeat toolkit.
```

Commit body:

```
Review and refine the future systemd service toolkit for RetroFlag Power.

Align the service and install script with the current retroflag-powerd daemon while keeping the quest limited to safe preparation.

Do not activate the service, edit rc.local, disable SafeShutdown.py, add GPIO, execute shutdown behavior, add resume, or introduce state storage.
```

---

# 10. Validation Commands

Run locally:

```bash
make check
```

Inspect files:

```bash
sed -n '1,220p' packaging/retroflag-power.service
sed -n '1,260p' packaging/install.sh
```

Optional if already installed:

```bash
shellcheck packaging/install.sh
```

Do not install shellcheck just for this quest unless the maintainer chooses to add it later as a development dependency.

---

# 11. Review Checklist

Before committing, confirm:

- [ ] Diff does not touch GPIO.
- [ ] Diff does not touch shutdown execution.
- [ ] Diff does not edit rc.local.
- [ ] Diff does not disable SafeShutdown.py.
- [ ] Diff does not activate systemd service.
- [ ] Diff does not add resume or state behavior.
- [ ] Diff keeps install behavior honest.
- [ ] Diff makes future systemd validation easier.
- [ ] `make check` passes.

---

# 12. Quest Reward

Completing this quest earns:

```
Heartbeat Toolkit Crafted
```

This does not complete Milestone 2.

It opens the gate toward Milestone 2.

---

# Closing

The first quest is not to seize power.

The first quest is to craft the toolkit safely.

A daemon that breathes is alive.

A daemon that survives service supervision will have a heartbeat.

Prepare the toolkit.

Do not pull the lever yet.
