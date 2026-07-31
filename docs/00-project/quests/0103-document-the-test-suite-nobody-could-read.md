---
id: QUEST-0103
title: Document the test suite nobody could read
version: 1.0.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Record why this project had 108 tests and no docs/06-testing/, what the coverage numbers actually mean, and how the 2026-07-03 archetype plan scored against the code that got written.
related:
  - ../MILESTONES.md
  - ../../06-testing/TEST_STRATEGY.md
  - ../../06-testing/ARCHETYPE_HANDHELD_EXAMPLE_retroflag-power.md
  - ../../../ENGINEERING_MANIFESTO.md
last_updated: 2026-07-30
---

# QUEST-0103 — Document the test suite nobody could read

> A green tick that nobody can interpret is not proof. It is a rumour with
> good posture.

## Quest Status

Implemented. `edc-lint` passes at zero errors; `make check` green.

## Epoch

Awakening

## Quest Type

Documentation backfill

## Quest Owner

Joshua Taft

The checkpoint this reports into is [`../MILESTONES.md`](../MILESTONES.md).

---

# 1. Quest Summary

This project had **108 test functions across 15 files**, an active CI
workflow, a three-layer `make check`, and no `docs/06-testing/` at all. The
governance guide has had a section for that folder the whole time.

The gap was not laziness. The template's worked example for this project was
written on 2026-07-03 and says outright that it is *"a planning template to
copy into that project's own `docs/06-testing/` once it exists"* — the copy
was simply never made once the project did exist. The instruction survived;
the follow-through did not.

Three documents now exist: [`TEST_STRATEGY.md`](../../06-testing/TEST_STRATEGY.md),
`ARCHETYPE_HANDHELD.md` carried in from the template, and
[a graded version of the worked example](../../06-testing/ARCHETYPE_HANDHELD_EXAMPLE_retroflag-power.md).

# 2. Two coverage numbers that lie if you read them cold

Measured with `go test -cover ./...` on Go 1.24, in a stock container.

**`internal/gpio` — 31.9%.** This looks like neglect and is not. The linux
half of the package reads `/sys/class/gpio`, `/sys/kernel/debug/gpio` and
shells out to `gpiod`; none of those exist on a dev host or in CI, so most
of the file is unreachable by design. The archetype is explicit that this is
where the permanent manual gate lives. Mocking the filesystem would buy a
bigger number and no more confidence, because the thing that can actually go
wrong — whether the pin the daemon reads is the pin the hardware drives — is
not a thing a mock can answer.

**`internal/power` — 0.0%.** Nine lines. One string type, one constant, one
`String()` method nothing calls yet.

Everything else is between 88.8% and 100%.

# 3. The finding that matters most

**The executor refuses to perform any real action.**
`internal/executor` returns `ErrUnsupportedPlan` for any plan that is not
both dry-run and noop-only, and `internal/actions` defines exactly one
action type, `TypeNoop`.

This is correct and deliberate — a half-built shutdown path that runs is
worse than one that declines, per
[`ENGINEERING_MANIFESTO.md`](../../../ENGINEERING_MANIFESTO.md)'s "Protect
the User From Failure". But it means a fully green suite proves the pipeline
routes, plans and reports, and proves nothing whatsoever about powering a
device down. Anyone reading the tick without that sentence beside it would
draw a conclusion the code does not support. It is now written down in both
new documents.

# 4. Grading the 2026-07-03 plan

Worth doing honestly rather than quietly rewriting the plan to match what
happened.

- **Capability boundary — right principle, wrong shape.** The plan predicted
  a `PowerControl` interface with `hasPowerSwitch()`, `readBattery()`,
  `setBacklight()`. What exists is a single-method `input.Observer` seam plus
  build-tag separation in `internal/gpio`. The plan's shape would have been
  worse: there is **not one mention of battery or backlight anywhere in the
  Go source**, so two of those three methods would have been invented for
  capabilities this daemon still does not have — the exact assumption
  Manifesto principle 13 warns against, committed in the name of following
  it.
- **Host-portable unit tests — correct.** `go test ./...` needs no hardware.
- **CI scope warning — correct and heeded.** CI runs `make check` on
  `ubuntu-latest` and nothing else. No GPIO simulator was ever built.
- **Timing benchmarks — did not happen.** Zero `Benchmark` functions, zero
  `Fuzz` functions, no `docs/08-performance/`. Not a failure: principle 11
  says measure before optimising, and nothing here has been optimised. The
  honest status is "not applicable yet", not "overdue".

# 5. What the archetype had no slot for

- **A portable shell smoke suite.** `make check-scripts` is the largest
  block in the Makefile — six scripts syntax-checked under `sh`, their
  `--help`/`--status`/`--plain`/`NO_COLOR` paths driven, and a full
  install → verify → uninstall round trip against a throwaway `HOME`. For a
  project whose deliverables are partly shell scripts carried onto hardware,
  that is a first-class test level the abstract archetype does not name.
- **Software that deliberately declines to act.** The archetype assumes the
  code under test does the dangerous thing and asks how to prove it safe.

# 6. Named, not fixed

`docs/00-project/REQUIREMENTS.md` carries **70 requirements and no test
references any of them.** The governance guide asks this folder to cover
requirement traceability. Wiring `REQ-` ids into test names is the obvious
next step; it has not been done, and pretending otherwise would have made
the new strategy document exactly the sterile filler it was written to
avoid.

# 7. Acceptance Criteria

- [x] `docs/06-testing/` exists with a strategy grounded in verified facts,
      not projection. **Done.**
- [x] Every number in it produced by an actual run, not estimated.
      **Done** — `go test -cover ./...` in `golang:1.24`.
- [x] The two misleading coverage figures explained where they are quoted.
      **Done.**
- [x] `node scripts/edc-lint.js .` reports zero errors. **Done.**
- [x] The template's stale "isn't on this host yet" claim corrected.
      **Done.**

---

# Closing

The suite was in good shape the whole time. What was missing was the
paragraph explaining that a green run covers a pipeline whose executor is
wired to refuse — and the two coverage figures that mean the opposite of
what they look like. Both were discoverable only by reading the source,
which is precisely the kind of knowledge that should not require reading the
source twice.
