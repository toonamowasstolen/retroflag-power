---
id: ARCH-LEGACY-001
title: Legacy Architecture Summary
version: 0.1.0
status: Superseded
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - Future Maintainers
purpose: Preserve the original one-line architecture sketch as project history.
related:
  - docs/00-project/requirements.md
superseded_by:
  - docs/04-architecture/system-overview.md
last_updated: 2026-07-03
---

# Legacy Architecture Summary

> Superseded by the canonical
> [System Overview](04-architecture/system-overview.md). The original sketch
> remains below for historical context.

```text
systemd -> daemon -> event bus -> handlers
```
