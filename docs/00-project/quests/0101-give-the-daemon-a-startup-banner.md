---
id: QUEST-0101
title: Give the Daemon a Startup Banner
version: 1.0.0
status: Verified
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
purpose: Close the real gap flagged in the 2026-07-14 gap analysis — internal/logging/logging.go was dead code (never called anywhere) and the running daemon had no warmth at all, despite PROJECT_MANIFEST.md's "Terminal Personality" section and PROJECT_MEMORY.md's "ASCII terminal welcome screen" aspiration.
related:
  - ../milestones.md
  - ../../../internal/logging/logging.go
  - ../../../cmd/retroflag-powerd/main.go
last_updated: 2026-07-14
---

# QUEST-0101 - Give the Daemon a Startup Banner

> Personality belongs in success, status, and discovery moments — PROJECT_MANIFEST.md Section 9.

## Quest Status

Verified

## Epoch

Awakening

## Quest Type

Implementation (small, real behavior change)

## Quest Owner

Joshua Taft

---

# 1. Quest Summary

Investigating the "generic logging.go" gap turned up something more specific than the original
finding described: `internal/logging.New()` was never called anywhere in the codebase — the real
daemon logger was constructed inline in `cmd/retroflag-powerd/main.go` via bare
`log.New(stderr, "", log.LstdFlags)`, twice, at two separate call sites. So the actual fix has two
parts: give `logging.New` real callers (fixing the dead code), and add the one-time startup banner
`PROJECT_MEMORY.md`'s "Product and experience ideas" lists as an aspiration.

# 2. Quest Objective

Add `logging.Banner`/`logging.WriteBanner` (plain ASCII, no color, no box-drawing — matches
`PROJECT_MANIFEST.md`'s "Terminal Personality" guidance and the explicit "Good: Everything checks
out. Player 1 is ready." tone example in `PROJECT_MEMORY.md` Section 28, not an elaborate banner),
print it once at real daemon startup only, and wire `logging.New` (now accepting an `io.Writer`) into
both of `main.go`'s logger-construction call sites.

---

# 3. Scope

## In Scope

- `internal/logging/logging.go`: `New(w io.Writer) *log.Logger` (was `New()`, hardcoded to
  `os.Stderr` — changed to accept a writer since the real usage needs to wrap a custom
  `readySignalWriter` at one call site); `Banner(appName, version string) string`;
  `WriteBanner(w io.Writer, appName, version string)`.
- `internal/logging/logging_test.go` — new, covers `New` writes to the given writer, `Banner`
  contains the app name/version, and (explicitly) that `Banner` never contains the literal
  `" ready\n"` substring, since that string is load-bearing elsewhere (`main.go`'s
  `readySignalWriter` watches for it to detect daemon readiness).
- `cmd/retroflag-powerd/main.go`: both `log.New(...)` call sites replaced with `logging.New(...)`;
  `logging.WriteBanner(stdout, cfg.AppName, cfg.Version)` added immediately before the real
  `app.New(...).Run(ctx)` call — and **only** there, not in `runAppAndProcess` (used by
  `--dry-run-power-button`, `--fake-power-button-observer`, `--fake-power-signal`), which has its own
  tested, parseable stdout contract that a banner would clutter.
- Verified for real: `gofmt -l .` (clean), `go build ./...`, `go test ./...` (all packages pass), and
  a real binary run confirming the banner appears on real daemon start and is absent from
  `--dry-run-power-button`/`--version`.

## Out of Scope

- Rewriting every existing log call site's format — the original finding explicitly called this out
  as "a bigger, separate change"; this quest only adds the one-time banner and fixes the dead-code
  logger constructor.
- ANSI color support — `PROJECT_MANIFEST.md` requires the terminal experience to work *without*
  color; adding color is a real future enhancement, not required to close this gap.
- Any change to the `" ready\n"`-detection logic itself in `main.go`'s `readySignalWriter` — verified
  it still fires correctly, didn't touch its mechanism.

---

# 4. Acceptance Criteria

This quest is complete when:

- [x] `gofmt -l .` reports no files. **Done.**
- [x] `go build ./...` succeeds. **Done.**
- [x] `go test ./...` passes, including new tests in `internal/logging`. **Done** — all 12 packages
  pass (one, `internal/power`, has no test files, pre-existing and unrelated).
- [x] A real built binary shows the banner on real startup, and does not show it on
  `--dry-run-power-button`/`--version`. **Done** — verified against a real compiled binary, not `go
  run`'s output alone.
- [x] `docs/00-project/milestones.md` gets a new entry. **Done as M-0010.**

**Toolchain note**: no Go toolchain exists on Ramuh (this project's docs live on a NAS/infra host,
not a dev workstation) — verified instead via a throwaway `node:22-bookworm-slim` + Go 1.24 container
built on Phoenix (`claude-tools-workspace:go1.24-node22`), synced the working tree over, ran the real
checks there, synced nothing back (verification only, no state changed on Phoenix). The image and
sync workspace persist there for reuse on future Go work.

---

# 5. Suggested Commit

Commit title:

```
Give the daemon a startup banner
```

Commit body: internal/logging.New() was dead code, never called anywhere — the real logger was
built inline in main.go, twice. Wires logging.New (now accepting an io.Writer) into both call sites,
and adds a one-time ASCII startup banner (plain, no color, matching PROJECT_MANIFEST.md's "Terminal
Personality" guidance and the "Good" tone example in PROJECT_MEMORY.md) printed only on real daemon
startup, never on the diagnostic/dry-run/fake-event paths. Verified via a real Go 1.24 toolchain
(none exists on Ramuh itself) run in a throwaway container on Phoenix: gofmt clean, build and all
tests pass, and a real compiled binary confirms the banner appears/doesn't appear exactly where
intended.

---

# Closing

Landed 2026-07-14 (see `milestones.md#m-0010`). The "ASCII terminal welcome screen" aspiration in
`PROJECT_MEMORY.md`'s idea list moves from documented to real, at the smallest scope that does that
honestly — no rewrite of the existing log format, no color dependency, no risk to the
`readySignalWriter` mechanism the daemon's own tests and tooling depend on.
