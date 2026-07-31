---
id: TEST-004-EX
title: Handheld Archetype Applied — RetroFlag Power
version: 0.2.0
status: Verified
owner: Joshua Taft
audience:
  - Everyone
  - AI Assistants
purpose: Check the handheld archetype's predictions for this project against the code that actually got written, and say plainly where they missed.
related:
  - ARCHETYPE_HANDHELD.md
  - TEST_STRATEGY.md
  - ../../ENGINEERING_MANIFESTO.md
last_updated: 2026-07-30
---

# Handheld Archetype Applied: RetroFlag Power

The template's copy of this document was written on 2026-07-03, before any
of this project's code existed on this host. It said so itself: *"this is a
planning template to copy into that project's own `docs/06-testing/` once it
exists."* The project now exists — 108 test functions across 15 files, all
green — and the copy was never made.

So this version is not a plan. It is the same document graded against what
was actually built, which is far more useful than a fresh set of intentions.
The archetype it grades is [ARCHETYPE_HANDHELD.md](ARCHETYPE_HANDHELD.md);
the resulting strategy is [TEST_STRATEGY.md](TEST_STRATEGY.md).

## The predictions, and how they held up

### The capability boundary — right principle, wrong shape

The plan called for *"a `PowerControl` interface with methods like
`hasPowerSwitch()`, `readBattery()`, `setBacklight()`"*, following
[`ENGINEERING_MANIFESTO.md`](../../ENGINEERING_MANIFESTO.md)'s
"Build With Capabilities, Not Assumptions".

No such interface exists. What exists instead is narrower and, on the
evidence, better: a single-method seam in `internal/input`.

```go
type Observer interface {
	NextEvent(context.Context) (Event, error)
}
```

`App.ProcessNextInputEvent(ctx, observer)` takes it as a parameter, so tests
pass a fake and never touch hardware. Platform variation lives one layer
down in `internal/gpio`, split by build tag across `probe.go`,
`probe_linux.go` and `probe_unsupported.go` rather than behind a capability
interface.

The principle survived; the shape did not. A three-method `PowerControl`
would have been speculative — `readBattery()` and `setBacklight()` describe
capabilities this daemon still does not have, and **there is not one mention
of battery or backlight anywhere in the Go source.** Guessing an interface
for them in 2026-07-03 would have produced exactly the assumption the
manifesto principle warns against, in the name of following it.

### Host-portable unit tests — correct

*"Unit tests run against a mock/stub implementation of that interface on any
dev machine."* True. `go test ./...` needs no hardware and no board. It runs
clean in a stock `golang:1.24` container.

### The CI scope warning — correct, and followed

*"CI cannot run a hardware-in-the-loop pass. Scope any future CI pipeline to
the host-portable unit level only — don't build a fake GPIO simulator."*

`.github/workflows/ci.yml` runs `make check` on `ubuntu-latest`, nothing
more. No GPIO simulator was ever built. This one was heeded exactly.

### Timing benchmarks — did not happen

The plan listed benchmarks as a "Should", logged to `docs/08-performance/`.
There are **zero** `Benchmark` functions in the repository, zero `Fuzz`
functions, and no `docs/08-performance/` directory.

That is not a failure. Manifesto principle 11 says measure before optimising
— and nothing here has been optimised, so nothing needed measuring. The
honest status is "not applicable yet", not "overdue". It becomes real the
first time boot or resume time is something anybody complains about.

## What the archetype did not anticipate at all

Two things this project does that the abstract archetype has no slot for:

**A portable shell smoke suite.** `make check-scripts` is the single largest
block in the Makefile. It syntax-checks six field scripts under `sh`, drives
their `--help`, `--status`, `--plain` and `NO_COLOR` paths, and runs a full
install → verify → uninstall round trip against a throwaway `HOME`. For a
project whose real artifacts are partly shell scripts carried onto hardware,
this is a first-class test level, and the archetype should probably grow one.

**An executor that refuses to act.** `internal/executor` returns
`ErrUnsupportedPlan` for anything that is not a dry-run noop, and
`internal/actions` defines only `TypeNoop`. The archetype assumes the
software under test does the dangerous thing and asks how to prove it safe.
Here the dangerous thing is deliberately not implemented yet, which changes
what a green run means — see
[TEST_STRATEGY.md](TEST_STRATEGY.md)'s note on scope.

## The one number to explain to anyone reading coverage

`internal/gpio` sits at 31.9%. That is the archetype's own boundary made
visible: the linux half reads `/sys/class/gpio`, `/sys/kernel/debug/gpio`
and shells out to `gpiod`, none of which exist off the device. The archetype
is explicit that this is where the manual gate lives permanently. The number
is the boundary, not a backlog item.
