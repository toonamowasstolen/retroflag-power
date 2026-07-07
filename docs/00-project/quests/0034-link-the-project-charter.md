---
id: QUEST-0034
title: Link the Project Charter
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Make the project charter easy to discover from the main documentation paths.
related:
  - README.md
  - PROJECT_CHARTER.md
  - docs/00-project/documentation-structure-and-governance.md
  - docs/00-project/quests/0033-add-the-project-charter.md
last_updated: 2026-07-07
---

# QUEST-0034 - Link the Project Charter

> Place the charter lantern on the main map, so future travelers can find the
> project's working compass before they step onto the trail.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Documentation

## Outcome

- `README.md` now links to `PROJECT_CHARTER.md` from the Project Documentation
  section.
- The README describes the charter as the campfire compass for RetroFlag
  Power's working style, quest rhythm, hardware caution, and project voice.
- `docs/00-project/documentation-structure-and-governance.md` now lists
  `PROJECT_CHARTER.md` in the recommended canonical locations.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- Documentation-only change.
- No Go code changes.
- No test changes.
- No daemon log changes.
- No CLI output changes.
- No lifecycle behavior changes.
- No planner behavior changes.
- No executor behavior changes.
- No GPIO.
- No hardware behavior changes.
- No shutdown execution.
- No `rc.local` edits.
- No `SafeShutdown.py` edits.
- No systemd or service activation.
- No resume.
- No state storage.
- No packaging changes.
- No third-party dependencies.

## Milestone Note

The project charter is now visible from the front README and the documentation
governance map. The working-style badge, hardware caution, and project voice
have a clearer trailhead for future quests.
