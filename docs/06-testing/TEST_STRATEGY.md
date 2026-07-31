---
id: TEST-001
title: Test Strategy
version: 0.1.0
status: Accepted
owner: Joshua Taft
audience:
  - Everyone
  - AI Assistants
purpose: Describe how this daemon is actually proved to work, which layer proves what, and which numbers mean something other than what they look like.
related:
  - ARCHETYPE_HANDHELD.md
  - ARCHETYPE_HANDHELD_EXAMPLE_retroflag-power.md
  - ../04-architecture/system-overview.md
  - ../../ENGINEERING_MANIFESTO.md
last_updated: 2026-07-30
---

# Test Strategy

This project had 108 test functions across 15 files and no document
explaining any of them for most of its life. That is the wrong way round —
a suite nobody can interpret is a suite nobody can trust, and two of the
numbers it produces are actively misleading if you read them cold.

The archetype this project belongs to is written up in
[ARCHETYPE_HANDHELD.md](ARCHETYPE_HANDHELD.md), and the specific way it was
applied here — including which of the original predictions came true — is in
[ARCHETYPE_HANDHELD_EXAMPLE_retroflag-power.md](ARCHETYPE_HANDHELD_EXAMPLE_retroflag-power.md).
The pipeline being tested is described in
[system-overview.md](../04-architecture/system-overview.md).

## Levels of testing

There are three, and only the first is what most people mean by "the tests".

| Level | Command | What it proves | Required before merge? |
|---|---|---|---|
| Go unit tests | `make test` (`go test ./...`) | The input → planner → executor pipeline, config parsing, logging, diagnostics, status, version | Must |
| Portable shell smoke | `make check-scripts` | The field scripts parse under `sh`, answer `--help`, honour `--plain`/`NO_COLOR`, and survive a real install → verify → uninstall round trip in a temp `HOME` | Must |
| Markdown link check | `make check-links` | Internal documentation links and anchors resolve | Must |
| Hardware-in-the-loop | Manual, on the real device | GPIO actually reads the pin the daemon thinks it reads | Must, before any release build — and see the warning below |

`make check` runs the first three plus a build and a version print. CI
(`.github/workflows/ci.yml`) runs exactly `make check` on `ubuntu-latest`
with the Go version pinned by `go.mod`, so CI and a local workshop run are
the same thing. That is deliberate: there is no CI-only step to be surprised
by.

## The two numbers that mean something other than what they say

Measured 2026-07-30 with `go test -cover ./...`, on Go 1.24:

| Package | Coverage | Reading |
|---|---|---|
| `internal/actions` | 100.0% | |
| `internal/config` | 100.0% | |
| `internal/diagnostics` | 100.0% | |
| `internal/executor` | 100.0% | |
| `internal/logging` | 100.0% | |
| `internal/planner` | 100.0% | |
| `internal/status` | 100.0% | |
| `internal/version` | 100.0% | |
| `internal/app` | 97.4% | |
| `internal/input` | 91.8% | |
| `cmd/retroflag-powerd` | 88.8% | |
| `internal/events` | no statements | declarations only |
| **`internal/gpio`** | **31.9%** | **not neglect — see below** |
| **`internal/power`** | **0.0%** | **nine lines, see below** |

**`internal/gpio` at 31.9% is the archetype boundary showing up as a
number.** `probe.go` and `probe_linux.go` split by build tag, and the linux
half reads `/sys/class/gpio`, `/sys/kernel/debug/gpio`, and shells out to
`gpiod`. None of those exist on a dev host or in a CI container, so most of
that file is unreachable by design. Raising this number by mocking the
filesystem would buy a bigger figure and no more confidence — the thing that
can actually go wrong is whether the pin the daemon reads is the pin the
hardware drives, and only real hardware answers that. Treat 31.9% as "the
host-portable part is covered", not as a gap to close.

**`internal/power` at 0.0% is nine lines.** One string type, one constant,
one `String()` method that nothing calls yet. It is vocabulary waiting for
the real power actions to exist. Worth a test the moment it grows a second
behaviour; not worth one today.

## What is scaffolded, and what the executor refuses to do

The most important thing to know before reading a green test run: **the
executor currently refuses to perform any real action.**

`internal/executor` returns `ErrUnsupportedPlan` for any plan that is not
both dry-run and noop-only, and `internal/actions` defines exactly one
action type — `TypeNoop`. So a passing suite proves the pipeline routes,
plans, and reports correctly. It does not prove anything about powering a
device down, because nothing in this codebase can yet do that.

This is fail-closed on purpose and matches
[`ENGINEERING_MANIFESTO.md`](../../ENGINEERING_MANIFESTO.md)'s
"Protect the User From Failure" — a half-built shutdown path that runs is
far worse than one that declines. But it means the test suite's scope is
narrower than its green tick suggests, and anyone adding the first real
action is also adding the first test that could fail for a reason that
matters to a user.

## What "verified" means for a change here

Passing `make check` is necessary and not sufficient.

- **Pure logic change** (planner, config, status, logging): `make check`
  green is enough.
- **Anything touching `internal/gpio`**: `make check` proves nothing about
  the part you changed. Flash to the real device and confirm the pin reads
  what you expect, per the archetype's hardware gate.
- **The first real action** (whenever `TypeNoop` gains a sibling): both, plus
  a deliberate power-cycle on real hardware, plus an entry in
  [`docs/03-operations/`](../03-operations/) recording what was observed —
  that folder exists precisely because this project's proof lives on a bench,
  not only in a runner.

## Explicitly not tested, and why

- **Real GPIO reads.** No fake sysfs, no simulated `gpiod`. The archetype
  says not to build a hardware simulator to shoehorn hardware testing into
  CI, and that holds: a simulator would only ever confirm its own
  assumptions.
- **Real power actions.** They do not exist yet (see above).
- **Requirement traceability.** `docs/00-project/REQUIREMENTS.md` carries 70
  requirements and no test references any of them. The governance guide asks
  this folder to cover traceability, so this is a real gap and is named here
  rather than quietly skipped. Wiring `REQ-` ids into test names or table
  entries is the obvious next step and has not been done.
