---
id: TEST-004
title: Test Strategy Archetype — Handheld / Embedded Hardware
version: 0.1.0
status: Accepted
owner: Joshua Taft
audience:
  - Everyone
  - AI Assistants
purpose: Give projects that talk to real GPIO, power and battery hardware test levels that respect the device.
related:
  - TEST_STRATEGY.md
  - ../../ENGINEERING_MANIFESTO.md
last_updated: 2026-07-03
---

# Test Strategy Archetype: Handheld / Embedded Hardware

For projects like RetroFlag Power — software that talks to real GPIO/
power/battery hardware on a device (Raspberry Pi class or similar), where
[Manifesto](../../ENGINEERING_MANIFESTO.md) principles 12 ("Portability Is
a Feature") and 17 ("Respect the Hardware") apply directly to how testing
is structured.

The shared philosophy these levels serve lives in
[TEST_STRATEGY.md](TEST_STRATEGY.md); this file is only the concrete
tooling for one kind of project.

## Levels of testing

| Level | Tooling | What it covers | Required before merge? |
|---|---|---|---|
| Unit (host-portable) | Whatever the implementation language's standard test framework is (e.g. `pytest` for Python, `cargo test` for Rust) | Core logic that doesn't touch real GPIO/hardware — state machines, config parsing, capability negotiation | Should |
| Hardware-in-the-loop (HIL) | Manual, on real target hardware (or a dev board with the same SoC/GPIO layout) | Anything behind a hardware capability interface (per Manifesto principle 13, "Build With Capabilities, Not Assumptions") — power switch, reset button, battery read, backlight control, resume-from-suspend | Must, before any release build |
| Timing/performance benchmarks | Manual measurement on real hardware, logged in `docs/08-performance/` | Boot time, resume time, shutdown time — see Manifesto principle 11, "Measure Before Optimizing" | Should, once a baseline exists |

## Core rule: core logic must be testable without hardware

Per Manifesto principle 12: "Core logic should be testable without
special hardware." This means the architecture itself has to put a real
interface boundary between hardware-specific code and everything else —
if a unit test needs a real GPIO pin to run, that's a sign the
abstraction boundary is in the wrong place, not a reason to skip unit
testing.

- Development should be possible on macOS/Linux/Windows (per Manifesto
  principle 12's list) using a mock/stub implementation of the hardware
  capability interface.
- The **real** hardware interface implementation gets its own
  hardware-in-the-loop pass — mocking a GPIO pin proves your state
  machine is correct, not that the pin actually does what you think.

## What "verified" means for a hardware change

Passing the host-portable unit suite is necessary, not sufficient. Before
calling a hardware-facing change done: flash/deploy to real target
hardware (or the closest available dev board) and observe the actual
physical behavior (switch flips, LED changes, resume actually resumes) —
matches the same "verify from source, never assume" discipline used
elsewhere on this host, just with a physical device standing in for "the
running container."

## What NOT to test in CI

Nothing that requires physical hardware presence can run in a generic CI
runner. Structure CI to run the host-portable unit level only, and treat
the hardware-in-the-loop pass as a manual, pre-release gate — don't try
to fake a hardware-in-the-loop step in CI with a software mock; that's
just the unit level again with extra steps.
