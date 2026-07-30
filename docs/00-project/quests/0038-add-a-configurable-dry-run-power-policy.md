---
id: QUEST-0038
title: Add a Configurable Dry-Run Power Policy
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Let configuration guide the dry-run power-button intent while keeping the daemon noop-only and safe.
related:
  - ../../../internal/config
  - ../../../internal/app
  - ../../../internal/planner
  - ../../../internal/executor
  - ../../../cmd/retroflag-powerd
  - ../../../README.md
  - ../../04-architecture/system-overview.md
  - 0037-record-power-intent-events.md
last_updated: 2026-07-07
---

# QUEST-0038 - Add a Configurable Dry-Run Power Policy

> Add a small policy compass to the power-button lantern: enough to point the
> dry-run route, never enough to wake real hardware or shutdown spells.

## Quest Status

Implemented

## Epoch

Awakening

## Quest Type

Implementation

## Outcome

- Added a `PowerButtonAction` [config](../../../internal/config) value for the dry-run power-button path.
- The default config sets `power_button_action` to `noop`.
- The [app](../../../internal/app) validates the configured policy before preparing a dry-run power
  intent plan.
- The [planner](../../../internal/planner) receives the configured action and still prepares only the
  supported noop dry-run plan.
- Unsupported policy values fail with a deterministic
  `unsupported power_button_action` error before [executor](../../../internal/executor) work.
- The dry-run CLI accepts `--power-button-action noop` and uses that configured
  policy for the power-button intent route.
- Tests prove the default noop policy, explicit noop policy, unsupported policy
  failures, deterministic breadcrumbs, and the continued absence of real
  shutdown or hardware behavior.
- README and architecture docs now describe the policy compass as a safe
  dry-run seam for future behavior.

## Validation

- [x] `make check` passed.
- [x] `make check-links` passed.

## Boundary

- No GPIO reads.
- No shutdown command execution.
- No systemd service activation.
- No `rc.local` replacement.
- No `SafeShutdown.py` replacement.
- No resume behavior.
- No persistent state.
- No packaging changes.
- No multiple real actions.

## Safe Commands

```sh
go run ./cmd/retroflag-powerd --dry-run-power-button
go run ./cmd/retroflag-powerd --dry-run-power-button --power-button-action noop
```

## Milestone Note

The daemon now carries a tiny policy compass in its config satchel. It points
the dry-run power-button route at `noop`, rejects unknown paths clearly, and
keeps every real power relic asleep until a future quest deliberately earns
that behavior.

## Related work

It rests on [README](../../../README.md) (the project overview). The work itself was carried out under [Record Power Intent Events](0037-record-power-intent-events.md). It reads alongside [retroflag-powerd](../../../cmd/retroflag-powerd) and [System Overview](../../04-architecture/system-overview.md).
