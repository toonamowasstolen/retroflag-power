---
id: DOCS-GOVERNANCE-001
title: Documentation Structure and Governance Guide
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
  - Future Maintainers
  - Documentation Authors
purpose: Define how RetroFlag Power documentation is organized, how metadata headers work, what each folder and document is for, how ADR/RFC templates should be used, and how progress, status, TODOs, and completion should be tracked.
related:
  - PROJECT_MEMORY.md
  - PROJECT_CHARTER.md
  - docs/00-project/roadmap.md
  - docs/00-project/milestones.md
  - docs/00-project/requirements.md
  - docs/05-development/ai-collaboration.md
  - docs/10-decisions/adr-template.md
  - docs/11-rfc/rfc-template.md
last_updated: 2026-07-03
---

# Documentation Structure and Governance Guide

> A project that remembers where things belong is easier to continue.

This guide explains how RetroFlag Power documentation should be organized and maintained.

It exists because the project is growing quickly, and future maintainers, contributors, and AI assistants need a clear map.

The goal is simple:

```
Every idea has a home.
Every decision has a record.
Every status means something.
Every TODO can be found again.
```

This document should be treated as the primary guide for documentation structure, metadata headers, ADR/RFC usage, progress tracking, and maintenance rules.

---

# 1. Guiding Principles

## 1.1 Documentation is part of the product

RetroFlag Power is not just code.

The project includes:

- purpose
- requirements
- architecture
- hardware knowledge
- decisions
- research
- workflows
- safety rules
- progress history

Documentation helps the project survive beyond one conversation, one machine, or one moment of excitement.

## 1.2 Every durable Markdown file needs metadata

Durable Markdown files should begin with a metadata header.

This makes the repository easier for humans, AI assistants, scripts, and future tooling to understand.

## 1.3 Put information in the narrowest useful home

Do not put every idea in one giant document.

Use the most specific document that fits.

Examples:

- A hardware fact goes in a hardware doc.
- A decision goes in an ADR.
- A proposal goes in an RFC.
- A future idea goes in PROJECT_MEMORY.md until it has a real home.
- A requirement goes in requirements.md.
- A quest goes in a quest file.

## 1.4 Do not let planning hide implementation

Documentation should move the project forward.

A planning session should produce something commit-worthy.

Once a topic is documented well enough to guide implementation, move forward.

## 1.5 Be honest about certainty

Separate:

- facts
- assumptions
- decisions
- research
- aspirations

Do not present guesses as validated hardware behavior.

---

# 2. Metadata Header Standard

Every durable Markdown artifact should start with this front matter block:

```
---
id: DOC-ID-001
title: Document Title
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Project Maintainers
purpose: Explain why this document exists.
related:
  - RELATED_DOC.md
last_updated: 2026-07-03
---
```

This metadata block is required for major documents and strongly recommended for any document expected to survive more than one work session.

---

# 3. Metadata Fields

## 3.1 `id`

The stable document identifier.

Example:

```
id: DOCS-GOVERNANCE-001
```

Use `id` to help people and AI assistants refer to the document without depending only on filenames.

### Rules

- Use uppercase words separated by hyphens.
- End with a three-digit number.
- Keep it stable after the file is committed.
- Do not reuse an ID for a different document.
- If a document is replaced, create a new ID and mark the old one `Superseded`.

### Suggested ID prefixes

```
PROJECT-*
VISION-*
ROADMAP-*
MILESTONES-*
REQUIREMENTS-*
ARCH-*
HARDWARE-*
DEV-*
ADR-*
RFC-*
QUEST-*
REFERENCE-*
DOCS-*
RESEARCH-*
TEST-*
PERF-*
```

## 3.2 `title`

The human-readable title of the document.

Example:

```
title: Documentation Structure and Governance Guide
```

### Rules

- Use clear human language.
- Match the main heading when practical.
- Avoid cute-only titles in metadata.
- Project personality can appear in the document body.

## 3.3 `version`

The document version.

Example:

```
version: 0.1.0
```

### Recommended version meaning

```
0.1.0
  Initial draft or first useful version.

0.2.0
  Meaningful expansion or restructuring.

0.3.0
  More content added before acceptance.

1.0.0
  Accepted stable version.

1.1.0
  Meaningful update after acceptance.

1.1.1
  Minor correction, typo fix, small clarification.
```

### When to update

Update `version` when the document meaningfully changes.

Do not update the version for every tiny typo unless the document is already stable and the project wants strict versioning.

## 3.4 `status`

The document lifecycle state.

Example:

```
status: Draft
```

Recommended status values:

```
Draft
Review
Accepted
Active
Implemented
Verified
Superseded
Archived
Rejected
```

See section 4 for detailed status rules.

## 3.5 `owner`

The person responsible for the document.

Current default:

```
owner: Joshua Taft
```

This does not mean only the owner may edit it.

It means the owner is responsible for final judgment.

## 3.6 `audience`

The intended readers.

Example:

```
audience:
  - Project Maintainers
  - Contributors
  - AI Assistants
```

Common audience values:

```
Project Maintainers
Contributors
AI Assistants
Future Maintainers
Hardware Porters
Curious Users
Players
Power Users
Documentation Authors
```

Use only audiences that make sense for the document.

## 3.7 `purpose`

One sentence explaining why the document exists.

Example:

```
purpose: Define how RetroFlag Power documentation is organized and maintained.
```

### Rules

- Keep it concise.
- Explain the document's job.
- Do not use this field for a summary of all content.

## 3.8 `related`

Other files connected to this document.

Example:

```
related:
  - PROJECT_MEMORY.md
  - docs/00-project/requirements.md
```

### Rules

- Use repository-relative paths.
- Include documents that provide context or are affected by this document.
- Update this list when the document becomes connected to new important artifacts.

## 3.9 `last_updated`

The last meaningful update date.

Example:

```
last_updated: 2026-07-03
```

### Rules

- Use `YYYY-MM-DD`.
- Update when content meaningfully changes.
- Do not worry about changing it for every typo unless the project becomes strict about documentation bookkeeping.

## 3.10 Optional metadata fields

Use optional fields only when useful.

### `superseded_by`

Use when status is `Superseded`.

```
superseded_by:
  - docs/adr/0007-new-lifecycle-model.md
```

### `supersedes`

Use when a document replaces older documents.

```
supersedes:
  - docs/adr/0002-old-lifecycle-model.md
```

### `implements`

Use when a quest, implementation note, or completion note implements requirements.

```
implements:
  - REQ-0005
  - REQ-0606
```

### `related_issues`

Use when GitHub issues exist.

```
related_issues:
  - 12
  - 18
```

### `reviewed_on`

Use when status moves to `Review`, `Accepted`, or `Verified`.

```
reviewed_on: 2026-07-10
```

### `validated_on`

Use when status moves to `Verified`.

```
validated_on: 2026-07-11
```

---

# 4. Status Values and When to Use Them

## 4.1 `Draft`

Use when a document is newly created or still flexible.

Most new documents start here.

Examples:

- new vision doc
- new hardware doc
- new ADR
- new quest
- new requirements section
- new architecture proposal

Use `Draft` when:

```
This is useful, but not final.
```

## 4.2 `Review`

Use when a document is ready for human review before becoming accepted guidance.

Use `Review` when:

```
The content is probably right, but we want a deliberate check.
```

Good candidates:

- requirements.md before coding against it heavily
- architecture docs before implementation
- hardware docs before relying on them
- ADRs before accepting them

## 4.3 `Accepted`

Use when the project agrees that the document reflects current direction.

Good candidates:

- ADRs
- charter
- terminology guide
- stable architecture docs
- stable requirements docs
- product vision

Use `Accepted` when:

```
This is the current project decision or direction.
```

## 4.4 `Active`

Use when the document describes an ongoing plan, process, or current work item.

Good candidates:

- PROJECT_MEMORY.md
- roadmap.md
- current quest docs
- current milestone notes
- active research index

Use `Active` when:

```
This document is still alive and guiding current work.
```

## 4.5 `Implemented`

Use when the proposed work has been built, but not necessarily fully validated.

Good candidates:

- quests
- RFCs
- requirements
- implementation notes

Use `Implemented` when:

```
The code or document change exists, but we have not fully validated it yet.
```

## 4.6 `Verified`

Use when the work has been tested, validated, or confirmed.

Good candidates:

- hardware docs after real hardware validation
- requirements after tests pass
- quests after acceptance criteria are met
- release checklists
- performance measurements

Use `Verified` when:

```
This has proof.
```

Examples of proof:

- CI pass
- unit test
- integration test
- hardware test
- repeated shutdown validation
- measured benchmark
- documentation review

## 4.7 `Superseded`

Use when a document has been replaced by a newer one.

Use `Superseded` when:

```
This was once useful, but another document now replaces it.
```

Add:

```
superseded_by:
  - docs/path/new-document.md
```

Do not delete the old document if it explains history.

## 4.8 `Archived`

Use when a document is no longer active but remains useful as history.

Good candidates:

- old research notes
- old milestone notes
- abandoned plans
- old quest records

Use `Archived` when:

```
This is no longer active, but it is worth keeping.
```

## 4.9 `Rejected`

Use for proposals that were considered and intentionally not adopted.

Good candidates:

- RFCs
- alternative designs
- rejected feature proposals

Use `Rejected` when:

```
We considered this and chose not to do it.
```

---

# 5. Recommended Status by Document Type

```
WHY.md
  Draft → Accepted

PROJECT_MEMORY.md
  Active

ENGINEERING_MANIFESTO.md
  Draft → Accepted

PROJECT_MANIFEST.md
  Draft → Accepted

PROJECT_CHARTER.md
  Draft → Accepted

docs/00-project/milestones.md
  Draft → Active or Accepted

docs/00-project/roadmap.md
  Active

docs/00-project/requirements.md
  Draft → Review → Accepted
  Individual requirements may track Implemented or Verified inside the document.

docs/00-project/awakening-readiness.md
  Draft → Verified once criteria are met.

docs/00-project/quests/*.md
  Draft → Active → Implemented → Verified → Archived

docs/01-product/vision.md
  Draft → Accepted

docs/02-hardware/*.md
  Draft → Review → Verified

docs/04-architecture/*.md
  Draft → Review → Accepted
  Superseded if replaced by a newer architecture.

docs/05-development/*.md
  Draft → Accepted or Active

docs/10-decisions/adr-template.md
  Accepted

docs/11-rfc/rfc-template.md
  Accepted

docs/adr/*.md
  Draft → Review → Accepted → Superseded if replaced

docs/rfc/*.md
  Draft → Review/Open → Accepted or Rejected → Implemented if built

docs/13-reference/*.md
  Draft → Accepted
```

---

# 6. Folder Structure

This is the intended documentation structure.

```
docs/
  00-project/
  01-product/
  02-hardware/
  03-software/
  04-architecture/
  05-development/
  06-testing/
  07-design/
  08-performance/
  09-research/
  10-decisions/
  11-rfc/
  12-api/
  13-reference/
  99-archive/
  adr/
  rfc/
```

Some folders may exist before they are heavily used.

Do not create empty folders unless needed.

---

# 7. Folder Purposes

## 7.1 Root project files

Root files are foundational identity and entry-point documents.

### `README.md`

The public project entry point.

`README.md` is the deliberate exception to the metadata header standard. It
should begin with the project title so GitHub presents a clean front door.

Should contain:

- short project description
- current status
- quick start once available
- development commands
- links to key docs
- safety warning if relevant
- what the project is and is not

Should not contain:

- every requirement
- full architecture
- long project history
- deep hardware research

### `WHY.md`

The spark.

Should contain:

- why the project exists
- what started it
- what experience it wants to protect
- emotional/product motivation

Should not become a roadmap or technical spec.

### `PROJECT_MEMORY.md`

The safety net.

Should contain:

- important ideas not yet promoted into dedicated docs
- open threads
- future features
- decisions waiting for ADRs
- context that must not be lost

Should be updated when:

- a major idea appears
- a major decision is made but no ADR exists yet
- a risk is discovered
- a new artifact promotes an idea out of memory

Should eventually point to dedicated docs instead of holding everything forever.

### `ENGINEERING_MANIFESTO.md`

The builder philosophy.

Should contain:

- engineering values
- craftsmanship principles
- collaboration style
- design/implementation balance

Should not contain project-specific implementation details unless used as examples.

### `PROJECT_MANIFEST.md`

The project compass.

Should contain:

- mission
- values
- promises
- product identity
- broad direction

Should not contain detailed requirements or file-by-file implementation.

### `PROJECT_CHARTER.md`

The formal project agreement.

Should contain:

- purpose
- scope
- non-goals
- stakeholders
- constraints
- risks
- success criteria

Should be updated carefully.

---

# 8. `docs/00-project/`

Project management, planning, governance, and progress tracking.

Use for documents that guide the project as a project.

## Expected files

### `docs/00-project/milestones.md`

The major project phases.

Should contain:

- milestone names
- purpose
- included work
- excluded work
- exit criteria
- risks
- victory conditions

Should not contain daily task tracking.

### `docs/00-project/roadmap.md`

The practical route through upcoming work.

Should contain:

- near-term sequence
- next gates
- what should wait
- when to stop planning and start coding

Should be updated when the route changes.

### `docs/00-project/requirements.md`

Traceable project requirements.

Should contain:

- requirement IDs
- requirement text
- priority
- status
- milestone
- rationale
- related docs

Should not become an implementation log.

### `docs/00-project/awakening-readiness.md`

The readiness check for entering Milestone 1.

Should contain:

- readiness verdict
- entry criteria
- first breath scope
- safety check
- non-requirements
- test plan

Can be marked `Verified` after Milestone 1 entry criteria are met and validated.

### `docs/00-project/quests/`

Focused work items.

A quest is bigger than a TODO but smaller than a milestone.

Use quests for:

- focused implementation missions
- validation missions
- documentation cleanup missions
- hardware discovery missions

Quest status flow:

```
Draft → Active → Implemented → Verified → Archived
```

Preferred quest names should feel like adventurer tools/items rather than restraint/control.

Good words:

```
kit
toolkit
lantern
compass
satchel
key
charm
relic
badge
spellbook
map
```

Avoid words that imply restraint unless technically appropriate:

```
harness
leash
cage
shackle
```

---

# 9. `docs/01-product/`

Product direction and user experience.

Use for documents that describe what the project should feel like and who it serves.

## Expected files

### `docs/01-product/vision.md`

The horizon.

Should contain:

- product north star
- desired player experience
- product promises
- user groups
- success examples
- non-goals
- product voice

### Future product docs

Possible future files:

```
goals.md
personas.md
user-stories.md
success-metrics.md
```

Use these when the product needs more detail.

---

# 10. `docs/02-hardware/`

Hardware facts, profiles, validation notes, and hardware-specific research.

## Expected files

### `docs/02-hardware/gpi-case-2.md`

Reference hardware witness.

Should contain:

- hardware identity
- known facts
- current shutdown path
- terminology findings
- current boot context
- assumptions
- unknowns
- risks
- validation checklist
- migration concept
- rollback concept

### Future hardware docs

Possible future files:

```
raspberry-pi.md
hardware-profiles.md
gpi-case-2w.md
nespi.md
superpi.md
```

Hardware docs must distinguish:

```
Fact
Assumption
Decision
Research
Aspiration
```

Never invent GPIO mappings.

Never claim support without validation.

---

# 11. `docs/03-software/`

Software stack notes.

Use for external software the project integrates with.

Possible files:

```
retropie-stack.md
retroarch.md
emulationstation.md
systemd.md
journald.md
```

Should contain:

- how the external software is used
- relevant commands
- integration boundaries
- known assumptions
- risks
- validation notes

Should not contain core architecture unless the architecture doc points here.

---

# 12. `docs/04-architecture/`

System design and technical boundaries.

## Expected files

### `docs/04-architecture/system-overview.md`

The courthouse.

Should contain:

- major components
- service responsibilities
- boundaries
- event flow
- package direction
- milestone-specific architecture
- decision points

### Future architecture docs

Possible files:

```
event-bus.md
hardware-abstraction.md
power-service.md
state-service.md
resume-service.md
metrics-service.md
configuration.md
logging.md
cli.md
filesystem-layout.md
```

Use architecture docs to explain system shape, not to record final decisions.

Use ADRs to record decisions.

---

# 13. `docs/05-development/`

Development workflow and contributor guidance.

## Expected files

### `docs/05-development/ai-collaboration.md`

How AI and humans work together.

Should contain:

- AI rules
- commit packet expectations
- no-planning-paralysis rule
- artifact rules
- fact/assumption/decision discipline
- safety rules

### Future development docs

Possible files:

```
workflow.md
git-workflow.md
coding-standards.md
commit-guidelines.md
local-development.md
```

Use this folder for how contributors work.

---

# 14. `docs/06-testing/`

Testing strategy and validation plans.

Possible files:

```
testing-strategy.md
hardware-validation.md
shutdown-validation.md
resume-validation.md
ci.md
```

Should contain:

- test types
- validation commands
- hardware validation checklists
- CI expectations
- requirement traceability

Testing docs should explain how proof is gathered.

---

# 15. `docs/07-design/`

Design, terminal UX, voice, and visual style.

Possible files:

```
terminal-ui.md
style-guide.md
ascii-banners.md
error-messages.md
```

Should contain:

- CLI output style
- color rules
- no-color behavior
- tone
- accessibility
- examples

Personality belongs here when it affects user-facing design.

---

# 16. `docs/08-performance/`

Performance measurements, boot timing, benchmarks, and optimization research.

Possible files:

```
boot-performance.md
benchmarking.md
resume-timing.md
baseline-measurements.md
```

Should contain:

- measurement methods
- current baseline
- benchmark commands
- results
- optimization candidates
- rollback notes

Do not recommend optimizations without measurements.

---

# 17. `docs/09-research/`

Research topics that are not yet decisions.

Possible files:

```
index.md
battery-state.md
sleep-mode.md
boot-order.md
kms-fkms.md
sbc-support.md
```

Should contain:

- question being researched
- findings
- sources or commands
- assumptions
- open questions
- recommendation
- whether an ADR/RFC is needed

Research is not commitment.

---

# 18. `docs/10-decisions/`

Decision process and templates.

## Expected file

### `docs/10-decisions/adr-template.md`

The ADR template.

Should contain:

- status
- context
- decision
- rationale
- alternatives
- consequences
- risks
- validation
- rollback/revision plan
- outcome

This is a template, not a decision record.

Actual ADRs currently live in:

```
docs/adr/
```

---

# 19. `docs/11-rfc/`

Proposal process and templates.

## Expected file

### `docs/11-rfc/rfc-template.md`

The RFC template.

Should contain:

- summary
- motivation
- goals
- non-goals
- user impact
- technical approach
- alternatives
- risks
- mitigations
- open questions
- acceptance criteria
- implementation plan
- testing
- rollback
- documentation plan
- decision outcome

This is a template, not a proposal record.

Actual RFCs should live in:

```
docs/rfc/
```

Note:

If someone says `AFC`, confirm whether they mean `RFC`.

The project standard term is `RFC` for Request for Comments.

---

# 20. `docs/12-api/`

Future API, command, event, or data contracts.

Possible files:

```
events.md
cli.md
state-schema.md
hardware-profile-schema.md
```

Use this folder when interfaces become stable enough to document as contracts.

Do not create API docs too early.

---

# 21. `docs/13-reference/`

Stable reference material.

## Expected files

### `docs/13-reference/terminology.md`

Rules of speech.

Should contain:

- naming rules
- preferred terms
- terms to avoid
- examples
- physical behavior naming rules

### `docs/13-reference/glossary.md`

Dictionary.

Should contain:

- short definitions
- quick lookup entries
- cross-references

Reference docs should be concise and stable.

---

# 22. `docs/99-archive/`

Historical or inactive material.

Use for:

- old research
- old plans
- old milestone summaries
- superseded drafts that should not clutter active folders

Do not use archive as a dumping ground.

Archive intentionally.

---

# 23. `docs/adr/`

Actual Architecture Decision Records.

Existing example:

```
docs/adr/0001-use-systemd.md
```

New ADR example:

```
docs/adr/0002-use-small-context-driven-daemon-lifecycle.md
```

Use `docs/adr/` for decisions that have been made or are being considered as decisions.

File naming:

```
0001-short-decision-title.md
0002-short-decision-title.md
```

Status flow:

```
Draft → Review → Accepted → Superseded
```

Rejected ADRs are possible but less common. If a topic is still being debated, prefer an RFC first.

---

# 24. `docs/rfc/`

Actual Request for Comments documents.

Use this folder for major proposals before they become decisions.

File naming:

```
0001-short-proposal-title.md
0002-short-proposal-title.md
```

Status flow:

```
Draft → Review/Open → Accepted or Rejected → Implemented
```

If accepted, an RFC may lead to:

- one or more ADRs
- requirements updates
- implementation quests
- architecture docs
- code changes

---

# 25. How to Use ADRs

ADR means:

```
Architecture Decision Record
```

Use an ADR when the project has made, or is about to make, a meaningful decision.

## Use ADRs for

- daemon lifecycle model
- systemd strategy
- GPIO abstraction
- hardware profile format
- state storage path
- CLI binary strategy
- logging strategy
- event bus design
- resume safety model
- install/uninstall strategy

## Do not use ADRs for

- tiny typo fixes
- routine refactors
- temporary notes
- ideas not yet ready for decision
- task tracking

## ADR lifecycle

1. Copy the template.
2. Name the file with the next number.
3. Fill in context and decision.
4. Mark status `Draft`.
5. Review.
6. Mark `Accepted` when agreed.
7. Mark `Superseded` if replaced later.

## Creating a new ADR

Example:

```
cp docs/10-decisions/adr-template.md docs/adr/0003-use-systemd-service-strategy.md
```

Then update:

- metadata `id`
- metadata `title`
- metadata `status`
- date
- related requirements
- context
- decision
- alternatives
- consequences
- risks
- validation plan
- rollback plan
- outcome

## Updating an ADR

Update an ADR when:

- status changes
- validation occurs
- a consequence is discovered
- the decision is superseded
- outcome is known

Do not rewrite history to pretend the project always knew the answer.

Add notes or outcome updates instead.

---

# 26. How to Use RFCs

RFC means:

```
Request for Comments
```

Use an RFC when the project needs to explore a meaningful proposal before deciding.

## Use RFCs for

- sleep-like mode proposal
- battery overlay proposal
- hardware profile schema proposal
- multi-slot resume proposal
- plugin architecture proposal
- release packaging proposal
- public beta process proposal
- major CLI redesign

## Do not use RFCs for

- obvious small code cleanup
- already-made decisions
- simple documentation fixes
- one-line Makefile improvements

## RFC lifecycle

1. Copy the template.
2. Name the file with the next number.
3. Fill in proposal details.
4. Mark status `Draft`.
5. Move to `Review` or `Open` when ready for discussion.
6. Mark `Accepted` or `Rejected`.
7. If accepted and built, mark `Implemented`.
8. Link to ADRs if the proposal creates decisions.

## Creating a new RFC

Example:

```
mkdir -p docs/rfc
cp docs/11-rfc/rfc-template.md docs/rfc/0001-sleep-like-mode.md
```

Then update:

- metadata `id`
- metadata `title`
- metadata `status`
- summary
- motivation
- goals
- non-goals
- user impact
- technical approach
- risks
- mitigations
- open questions
- acceptance criteria
- implementation plan
- documentation plan

## RFC to ADR flow

An accepted RFC may produce one or more ADRs.

Example:

```
RFC-0003 — Hardware Profile Schema
  Accepted
    ↓
ADR-0008 — Use YAML Hardware Profiles
ADR-0009 — Keep GPIO Details Out of Core Services
```

RFCs explore.

ADRs decide.

---

# 27. How to Track Progress

Progress should be tracked at multiple levels.

## 27.1 Milestones

Use:

```
docs/00-project/milestones.md
```

For large phases:

- Dreaming
- Awakening
- Heartbeat
- Power
- Memory
- Resume
- Momentum
- Polish
- Expansion
- Release
- Launch

Milestones answer:

```
What phase are we in?
What does done mean?
What should not be included?
```

Update milestones when:

- exit criteria change
- a milestone is completed
- a new milestone is added
- scope boundaries need clarification

## 27.2 Roadmap

Use:

```
docs/00-project/roadmap.md
```

For the practical route.

Roadmap answers:

```
What comes next?
What should wait?
When do we stop planning?
```

Update roadmap when:

- next steps change
- a gate is passed
- a route is no longer accurate
- implementation changes the sequence

## 27.3 Requirements

Use:

```
docs/00-project/requirements.md
```

For traceable needs.

Requirements answer:

```
What shall the system do?
```

Update requirements when:

- a new requirement is discovered
- a requirement changes priority
- a requirement status changes
- an implementation satisfies a requirement
- validation proves a requirement

Recommended requirement status flow:

```
Draft → Accepted → Implemented → Verified
```

## 27.4 Quests

Use:

```
docs/00-project/quests/
```

For focused work.

Quests answer:

```
What focused mission are we doing now?
```

Quest status flow:

```
Draft → Active → Implemented → Verified → Archived
```

Update quest docs when:

- work begins
- implementation is completed
- validation passes
- scope changes
- quest is abandoned or replaced

## 27.5 GitHub Issues

Use GitHub Issues when external tracking becomes useful.

Good issue types:

- bug
- feature
- research
- hardware validation
- documentation cleanup
- release task

If GitHub Issues are used, docs should link to them with `related_issues`.

## 27.6 Pull Requests

Use PRs when multiple contributors or review workflows become useful.

PRs should reference:

- requirements
- quests
- ADRs
- RFCs
- validation commands

---

# 28. TODO Tracking Rules

TODOs should be easy to find and not scattered randomly.

## Use inline TODOs only for local code reminders

Example:

```
TODO: Replace placeholder logger when structured logging ADR is accepted.
```

Rules:

- include enough context
- avoid vague TODOs
- do not use TODOs for big features
- promote large TODOs into requirements, quests, issues, or PROJECT_MEMORY.md

## Use PROJECT_MEMORY.md for ideas that must not be lost

Good for:

- future feature ideas
- open design threads
- research topics
- remembered user preferences
- scope reminders

## Use requirements.md for things the system must do

If a TODO says “the system shall,” it probably belongs in requirements.

## Use quests for focused next work

If a TODO is a contained mission, make it a quest.

## Use ADR/RFC for decisions/proposals

If a TODO asks “which approach should we choose,” it probably needs an ADR or RFC.

---

# 29. Completion Tracking Rules

Completion should be based on evidence.

## A task is not complete because it was written down

Documentation is progress, but implementation and validation need proof.

## Use these completion states

```
Planned
Active
Implemented
Verified
Archived
```

## Evidence examples

### Implemented evidence

- code exists
- docs exist
- script exists
- workflow exists
- test exists

### Verified evidence

- test passed
- CI passed
- command output captured
- hardware validation completed
- repeated safe shutdown confirmed
- benchmark measured
- review completed

## Example completion update

Quest before work:

```
status: Active
```

Quest after code exists:

```
status: Implemented
```

Quest after validation:

```
status: Verified
validated_on: 2026-07-10
```

Then add a short outcome section:

```
Outcome:
- make check passed
- GitHub Actions CI passed
- no hardware behavior changed
```

---

# 30. Progress Update Examples

## 30.1 Updating a quest

Before:

```
status: Active
```

After implementation:

```
status: Implemented
```

After validation:

```
status: Verified
validated_on: 2026-07-03
```

Add:

```
# Outcome

- Makefile updated.
- VS Code tasks added.
- make check passed.
- CI passed.
```

## 30.2 Updating a requirement

Before:

```
Status: Draft
```

After implementation:

```
Status: Implemented
```

After validation:

```
Status: Verified
Evidence:
- internal/app/app_test.go
- GitHub Actions CI #27
```

## 30.3 Updating an ADR

Before:

```
status: Draft
```

After review:

```
status: Accepted
reviewed_on: 2026-07-03
```

After replaced:

```
status: Superseded
superseded_by:
  - docs/adr/0007-new-daemon-lifecycle.md
```

## 30.4 Updating roadmap

Add a short note:

```
## Current Position

Milestone 1 — Awakening

Completed:
- first daemon
- lifecycle test
- Workshop
- VS Code tasks
- CI make check

Next:
- prepare Heartbeat toolkit
```

---

# 31. Document Update Checklist

Before committing documentation, check:

- [ ] Metadata header exists.
- [ ] `id` is unique.
- [ ] `title` is clear.
- [ ] `status` is appropriate.
- [ ] `last_updated` is current.
- [ ] `related` links are useful.
- [ ] Facts are separated from assumptions.
- [ ] Decisions are not hidden in random docs.
- [ ] Proposals are not treated as decisions.
- [ ] TODOs have a proper home.
- [ ] Completion claims include evidence.
- [ ] File is in the correct folder.
- [ ] Filename is clear and lowercase where practical.
- [ ] Markdown is readable.
- [ ] The document does not duplicate another document unnecessarily.

---

# 32. AI Assistant Instructions

When an AI assistant works in this repository, it should:

1. Read this guide.
2. Read PROJECT_MEMORY.md.
3. Read the current roadmap.
4. Read current milestone docs.
5. Check existing folder structure before creating files.
6. Use metadata headers.
7. Put files in the right folder.
8. Avoid inventing new structure unless needed.
9. Ask before reorganizing major docs.
10. Prefer updating existing docs over creating duplicates.
11. Keep commit packets coherent.
12. Preserve project voice.
13. Avoid planning paralysis.
14. Never invent hardware facts.
15. Keep safety-critical claims honest.

AI assistants should not create duplicate documents such as:

```
docs/roadmap.md
docs/00-project/roadmap.md
```

without explaining why both exist.

If duplicates already exist, suggest consolidation rather than adding more.

---

# 33. Duplicate Document Cleanup Rule

If duplicate or overlapping documents exist:

1. Identify both files.
2. Decide which one is canonical.
3. Move unique content into the canonical file.
4. Mark the older file `Superseded` or move it to archive.
5. Update links.
6. Commit the cleanup as a coherent packet.

Do not delete old content without review.

Example:

```
docs/roadmap.md
docs/00-project/roadmap.md
```

Possible action:

- keep `docs/00-project/roadmap.md` as canonical
- move unique content from `docs/roadmap.md`
- mark `docs/roadmap.md` Superseded or archive it

---

# 34. Recommended Canonical Locations

Use these canonical locations unless an ADR changes them.

```
Project roadmap:
  docs/00-project/roadmap.md

Milestones:
  docs/00-project/milestones.md

Requirements:
  docs/00-project/requirements.md

Quests:
  docs/00-project/quests/

Product vision:
  docs/01-product/vision.md

Reference hardware:
  docs/02-hardware/gpi-case-2.md

System architecture:
  docs/04-architecture/system-overview.md

AI collaboration:
  docs/05-development/ai-collaboration.md

ADR template:
  docs/10-decisions/adr-template.md

RFC template:
  docs/11-rfc/rfc-template.md

Actual ADRs:
  docs/adr/

Actual RFCs:
  docs/rfc/

Terminology:
  docs/13-reference/terminology.md

Glossary:
  docs/13-reference/glossary.md

Archived content:
  docs/99-archive/
```

---

# 35. Naming Conventions

## Folders

Use lowercase folders.

Use numeric prefixes for major doc categories.

Example:

```
00-project
01-product
02-hardware
```

## Files

Use lowercase kebab-case.

Good:

```
system-overview.md
gpi-case-2.md
ai-collaboration.md
```

Avoid:

```
SystemOverview.md
system_overview.md
New Doc.md
```

## ADR files

Use number plus short title:

```
0002-use-small-context-driven-daemon-lifecycle.md
```

## RFC files

Use number plus short title:

```
0001-sleep-like-mode.md
```

## Quest files

Use number plus short title:

```
0001-craft-the-heartbeat-toolkit.md
```

Prefer quest titles that feel like adventurer tools/items when appropriate.

---

# 36. When to Create a New Document

Create a new document when:

- the topic has a different audience
- the topic will grow independently
- the topic needs its own lifecycle/status
- the topic is a decision
- the topic is a proposal
- the topic is a focused quest
- the topic is hardware-specific
- the topic is stable reference material

Do not create a new document when:

- an existing document already has the right home
- the idea is a small update
- the idea is not yet formed
- the document would duplicate another file
- the goal is just to avoid editing a large file

When unsure, update PROJECT_MEMORY.md first.

---

# 37. When to Update Existing Documents

Update existing documents when:

- a related decision is made
- a requirement is implemented or verified
- a roadmap step changes
- a quest changes status
- hardware facts are validated
- architecture changes
- a duplicate needs consolidation
- progress needs to be reflected

Do not leave stale documents behind.

A wrong document is worse than no document because it misleads future builders.

---

# 38. Commit Packet Guidance for Docs

Documentation commits should be coherent.

Good commit packet:

```
docs/00-project/requirements.md
docs/04-architecture/system-overview.md
```

Only if both changes are related.

Better to split if they are separate topics.

Commit messages should be useful and project-flavored.

Examples:

```
Give the dream its first quest.
Preserve the first breath in crystal.
Open a second door to Workshop.
Teach the forge to use Workshop.
```

Commit body should explain:

- what changed
- why it changed
- what scope was avoided
- what validation occurred if applicable

---

# 39. Current Cleanup Opportunities

The current repository may contain some older broad docs alongside newer structured docs.

Known examples from prior file listing:

```
docs/architecture.md
docs/development.md
docs/roadmap.md
docs/adr/0001-use-systemd.md
```

These should be reviewed later.

Do not delete them casually.

Recommended cleanup quest:

```
QUEST-0002 — Sort the Old Scrolls
```

Possible goal:

- identify canonical docs
- compare older docs to newer docs
- preserve unique content
- archive or supersede duplicates
- update links

## Future EDC extraction

Later, outside Milestone 1 — Awakening:

- create `docs/00-project/edc-project-structure.md` as a reusable guide for
  applying this documentation and project structure to older projects
- create or extract an `edc-project-template/` with the reusable root documents,
  metadata conventions, folder structure, ADR/RFC templates, project memory,
  roadmap, requirements, and AI collaboration guidance

These are TODO captures only. Do not create either artifact as part of the
current cleanup quest.

---

# 40. Closing Rule

When in doubt, ask:

```
What kind of knowledge is this?
```

Then place it accordingly:

```
Why?
  WHY.md

Project memory?
  PROJECT_MEMORY.md

Plan?
  docs/00-project/roadmap.md

Phase?
  docs/00-project/milestones.md

Requirement?
  docs/00-project/requirements.md

Focused work?
  docs/00-project/quests/

Product experience?
  docs/01-product/

Hardware?
  docs/02-hardware/

Software stack?
  docs/03-software/

Architecture?
  docs/04-architecture/

Development workflow?
  docs/05-development/

Testing?
  docs/06-testing/

Design/UX?
  docs/07-design/

Performance?
  docs/08-performance/

Research?
  docs/09-research/

Decision template?
  docs/10-decisions/

Proposal template?
  docs/11-rfc/

Actual decision?
  docs/adr/

Actual proposal?
  docs/rfc/

Reference term?
  docs/13-reference/

Old but valuable?
  docs/99-archive/
```

The filing cabinet is now labeled.

Future builders should not need to guess where the scrolls go.
