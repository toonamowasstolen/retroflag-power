---
id: QUEST-0007
title: Add the Link Lantern
version: 0.1.0
status: Implemented
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Documentation Authors
purpose: Add a small local checker for internal Markdown links and anchors.
related:
  - Makefile
  - scripts/check-markdown-links.py
  - docs/00-project/milestones.md
  - docs/00-project/quests/0006-gather-the-checkpoints-into-one-ledger.md
last_updated: 2026-07-06
---

# QUEST-0007 — Add the Link Lantern

> Carry a small light. Notice when the paths between documents disappear.

## Quest Status

Implemented

## Quest Type

Documentation validation and local developer workflow

---

# 1. Objective

Add a small Link Lantern that checks internal Markdown file links and anchors,
including the explicit stable Milestone anchors established by QUEST-0006.

---

# 2. Scope

## In Scope

- scan Markdown files across the repository
- verify relative internal file destinations
- verify explicit HTML `id` and `name` anchors
- verify common GitHub-style Markdown heading anchors
- report source files and line numbers for broken links
- expose the checker through `make check-links`
- keep the checker independent from `make check` for now
- repair broken internal documentation links found during validation

## Out of Scope

- checking external URLs or network availability
- fully reproducing every GitHub heading-slug edge case
- adding third-party packages
- changing production Go code or packaging
- activating systemd
- adding GPIO or executing shutdown
- editing `rc.local` or replacing `SafeShutdown.py`
- adding resume or state storage

---

# 3. Implementation

The Link Lantern lives at `scripts/check-markdown-links.py` and uses only the
Python 3 standard library. No package metadata or third-party dependency is
required.

Workshop exposes it as:

```text
make check-links
```

It is intentionally separate from `make check`. A fresh checkout needs Python
3 available before the link check can run, and the standard Go validation path
does not yet declare or install that runtime.

---

# 4. Known Limitations

- external and protocol-based links are skipped
- inline Markdown links are checked; reference-style link definitions are not
  currently parsed
- heading slugs cover common GitHub-style headings and duplicate suffixes, but
  punctuation, Unicode, embedded formatting, and unusual HTML can differ from
  GitHub's exact slugger behavior
- anchors in non-Markdown files are not verified

Explicit HTML anchors such as `<a id="m-0001"></a>` are checked directly and do
not depend on heading-slug approximation.

---

# 5. Acceptance Criteria

- [x] A small repository-local link checker exists.
- [x] Internal relative Markdown file links are verified.
- [x] Explicit HTML anchors are verified.
- [x] Common Markdown heading anchors are verified.
- [x] Stable M-0001 through M-0004 citations pass.
- [x] `make check-links` is documented in Workshop help.
- [x] `make check` remains unchanged.
- [x] No third-party dependency was added.
- [x] No production Go code or packaging changed.

---

# 6. Outcome

RetroFlag Power now carries a Link Lantern for local documentation validation.
Broken internal paths and anchors fail with a source location, while valid
stable Milestone citations remain visible and testable.

---

# 7. Validation

Completed before commit:

- [x] `make clean` passed.
- [x] `git status --short` was reviewed.
- [x] `make check-links` passed.
- [x] `make check` passed.
- [x] `git diff --check` passed.
- [x] `git diff --stat` was reviewed.
- [x] `git diff` was reviewed.

---

# Closing

The map has paths.

The Lantern makes sure they still lead somewhere.
