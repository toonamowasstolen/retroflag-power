---
id: MANIFESTO-001
title: Engineering Manifesto
version: 0.1.0
status: Draft
owner: Joshua Taft
audience:
  - Joshua Taft
  - Taft Consulting
  - Project Maintainers
  - Future Contributors
  - AI Assistants
purpose: Capture the engineering mindset, creative philosophy, and project-starting principles that should guide this and future projects.
related:
  - WHY.md
  - PROJECT_MEMORY.md
  - PROJECT_MANIFEST.md
  - PROJECT_CHARTER.md
last_updated: 2026-07-03
---

# Engineering Manifesto

> Never lose today's excitement in tomorrow's implementation.

This document is not tied to one project.

It is a reminder of how I want to build.

Every project begins with energy.

A spark.

A problem worth solving.

A feeling that something could be better, more useful, more elegant, more joyful, or more alive.

The work of engineering is not only to implement the idea.

It is to protect that spark while turning it into something real.

---

# 1. I Build Experiences, Not Just Software

Software is more than code.

It is the feeling someone has when they use it.

It is the confidence they feel when it behaves reliably.

It is the smile they get from a thoughtful message.

It is the relief they feel when a difficult task becomes simple.

It is the trust they develop when the system is clear, predictable, and well made.

The implementation matters.

The experience matters too.

If someone spends hours, days, or years with something I build, those hours should feel better because I cared.

---

# 2. Craftsmanship Matters

Craftsmanship is doing the small things well even when nobody asked.

It is choosing a good name.

Writing the extra sentence of documentation.

Leaving a helpful error message.

Designing a clear interface.

Making the command output pleasant to read.

Explaining why a decision was made.

Polishing the rough edge before someone else cuts themselves on it.

Craftsmanship is not perfection.

It is care made visible.

---

# 3. Documentation Is Part of the Product

A project is incomplete if people cannot understand it.

Documentation is not a chore that comes after the real work.

Documentation is part of the real work.

Good documentation teaches.

Good documentation welcomes.

Good documentation preserves intent.

Good documentation helps future contributors move with confidence instead of fear.

Every project should explain:

- Why it exists.
- Who it serves.
- What problem it solves.
- What it will not do.
- How it is designed.
- How to contribute safely.
- How to recover when something goes wrong.

If future me has to reverse engineer my own intention, I failed to document enough.

---

# 4. Design Before Implementation

Architecture is easier to change on paper than in production.

Before writing code, I should understand the experience I am trying to create.

Not every detail needs to be solved up front.

Not every interface needs to be perfect.

But the project deserves a direction before implementation begins.

Design is not procrastination when it produces clarity.

Planning becomes dangerous only when it stops producing artifacts.

---

# 5. Architecture Exists to Enable Implementation, Not Delay It

Planning should create momentum.

It should not become a hiding place.

Every planning session should produce something commit-worthy:

- A document.
- A diagram.
- A decision.
- A requirement.
- A test.
- A prototype.
- A small piece of code.

The goal is not to admire the blueprint forever.

The goal is to build from it.

When planning stops moving the project forward, it is time to build.

---

# 6. Every Work Session Ends With a Victory

Momentum matters.

A victory does not have to be large.

It can be one markdown file.

One diagram.

One test.

One command that finally works.

One bug understood.

One commit pushed.

One decision recorded.

A project grows through small, completed promises.

Every session should leave the project better than it was before.

---

# 7. Respect Future Contributors

Every line of code is a conversation with someone I may never meet.

That person might be a contributor.

A user.

A coworker.

A client.

A future version of me.

They deserve clarity.

They deserve names that make sense.

They deserve errors that help.

They deserve documentation that tells the truth.

They deserve architecture that can be understood without archaeology.

Respect shows up in the small things.

---

# 8. Name Things According to Real-World Behavior

Names should reflect what something is, how it behaves, and how people understand it.

Do not name things only after implementation details.

A latching hardware control is a switch.

A momentary hardware control is a button.

A persistent condition is a state.

A change in condition is an event.

Good names reduce complexity before code is even written.

If the name is wrong, the model is probably wrong too.

---

# 9. Make Complexity Invisible

Users should not have to understand the machinery to benefit from it.

A system may contain services, adapters, event buses, hardware profiles, state stores, configuration files, and platform-specific behavior.

But the experience should feel simple.

The user flips a switch.

The game resumes.

The technology disappears.

That is the magic.

---

# 10. Leave Moments of Delight

Useful software can still be enjoyable.

A terminal can have personality.

A status screen can be beautiful.

A success message can make someone smile.

A benchmark can feel like crossing a finish line.

A project can be professional without being sterile.

Moments of delight should never obscure important information.

They should support clarity, not replace it.

Errors should be direct, actionable, and searchable.

Success can celebrate.

---

# 11. Measure Before Optimizing

Performance work should be guided by evidence.

Guessing can inspire investigation, but data should guide decisions.

Measure:

- Boot time.
- Startup time.
- Resume time.
- Shutdown time.
- CPU usage.
- Memory usage.
- Reliability.
- Failure modes.

A good optimization is one that can be proven.

---

# 12. Portability Is a Feature

Hardware changes.

Platforms evolve.

Operating systems move forward.

Good architecture adapts.

Hardware-specific code should live behind interfaces.

Core logic should be testable without special hardware.

Development should happen wherever possible:

- macOS
- Linux
- Windows where practical
- Raspberry Pi hardware for validation

The goal is not to support everything immediately.

The goal is to avoid closing doors unnecessarily.

---

# 13. Build With Capabilities, Not Assumptions

A system should ask what a platform can do.

Not just what model it is.

Does it have a Power Switch?

Does it have a Reset Button?

Does it expose battery information?

Can it control a backlight?

Can it resume a game session?

Capabilities make software portable.

Assumptions make software brittle.

---

# 14. Curiosity Is a Strength

Research is part of engineering.

Before replacing a system, understand why it exists.

Before optimizing boot, measure where time is spent.

Before choosing a library, understand the alternatives.

Before naming an abstraction, understand the real thing it represents.

Curiosity prevents shallow solutions.

---

# 15. Do Not Be Clever When Clear Is Better

Clever code can be satisfying.

Clear code is useful.

The best solution is usually the one a future maintainer can understand quickly, modify safely, and trust.

Prefer readable architecture over impressive architecture.

Prefer direct language over fashionable language.

Prefer boring reliability over fragile brilliance.

---

# 16. Protect the User From Failure

Failure should be expected.

Power may be lost.

Files may be missing.

Save states may fail.

Hardware may behave differently than expected.

A good system fails safely.

It explains what happened.

It preserves user data when possible.

It offers a recovery path.

It does not punish the user for trusting it.

---

# 17. Respect the Hardware

Hardware has a personality.

It has limits.

It has quirks.

It has history.

Do not fight the hardware without understanding it.

Model the physical world accurately.

A switch is not a button.

A battery is not just a number.

A handheld is not just a computer.

Respecting hardware leads to better software.

---

# 18. Build Things I Would Be Excited to Discover

Before beginning a project, ask:

> If someone else had built this, would I be excited to find it?

If the answer is yes, build toward that feeling.

If the answer is no, find the missing spark before going further.

A project worth building should feel worth discovering.

---

# 19. Preserve the Dream, Then Build the Machine

Every meaningful project has a dream stage.

That stage matters.

It reveals the values, the experience, the tone, and the reason the project exists.

But the dream must eventually become artifacts.

Artifacts become commits.

Commits become software.

Software becomes experience.

Experience becomes memory.

Do not skip the dream.

Do not live there forever.

---

# 20. Never Lose Today's Excitement in Tomorrow's Implementation

The beginning of a project is full of energy.

The middle is full of details.

The end is full of pressure.

The danger is forgetting why the work mattered in the first place.

When the implementation gets difficult, return to the spark.

When the bug is frustrating, return to the user.

When the architecture feels heavy, return to the experience.

When motivation fades, return to the promise:

> Never lose today's excitement in tomorrow's implementation.

---

# Working Principles Checklist

At the start of a project, ask:

- [ ] Why does this project exist?
- [ ] What experience should it create?
- [ ] Who is it for?
- [ ] What problem is it solving?
- [ ] What will it intentionally not do?
- [ ] What would make this project delightful to discover?
- [ ] What should be documented before code is written?
- [ ] What is the smallest meaningful first artifact?
- [ ] What would count as today's victory?
- [ ] How do we protect momentum?
- [ ] How do we avoid planning paralysis?
- [ ] How do we respect future contributors?
- [ ] What names need to be chosen carefully?
- [ ] What assumptions need to be recorded?
- [ ] What risks need to be acknowledged?
- [ ] What can be measured?
- [ ] What must remain portable?
- [ ] What should make someone smile?

---

# Lessons Learned

This section should grow over time.

Every meaningful project should leave behind at least one lesson.

## 2026-07-03 — RetroFlag Power

A simple technical project can become something much more meaningful when the experience, philosophy, and craft are taken seriously before implementation begins.

The excitement is not a distraction.

It is design material.

Protect it.
