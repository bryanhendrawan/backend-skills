---
name: generate-topics
description: Proposes a batch of new unique topics for the backend-skills content library. Checks topics.md for non-overlap by angle, not just title. Use when planning the next set of days to publish.
argument-hint: <count, default 5>
disable-model-invocation: false
---

Propose N new topic ideas for the content library. Default N=5 if no argument is given.

## Step 1 — Read the canonical index

Read `topics.md` in the repo root. This is the source of truth for every topic that has been published or planned. Note:

- The latest date in the file (so new proposals auto-increment from there).
- Which **angles** are covered, not just which titles exist. Two topics on "database migrations" cover different ground if one is about generating them and one is about reviewing them.
- Which **domains** have been heavy (e.g. AI workflow, project setup, code review) and which are light (e.g. observability, testing, performance, security, deployment, MCP, debugging).
- Which **tools** the existing topics use (`cc` = Claude Code CLI, `web` = Claude.ai or ChatGPT, `both` = either).
- Which **series** are in progress (e.g. `go-service`) and whether continuing them makes sense.

## Step 2 — Identify gaps

Before proposing anything, write a short gap analysis (1-2 sentences) covering:

- Domains under-represented relative to a working backend engineer's day.
- Tool mix — if every recent topic is `cc`, propose at least one `web` or `both`.
- Stack variety — if every code sample so far is Go, consider Python or no-code for variety. Don't force it.

## Step 3 — Draft proposals

Propose exactly N topics as a markdown table with the same columns as `topics.md`:

```
| Date | Slug | Title | Code | Stack | Tool | Series |
```

Rules:

- **Date**: increment from the latest date in `topics.md`, one day per topic. Skip no days.
- **Slug**: kebab-case, ≤5 words, matches the angle not the title verbatim.
- **Title**: ≤8 words, hook not chapter heading. English title is fine here; the slides themselves will be in Indonesian.
- **Code**: `yes` only if the topic genuinely needs a working sample to prove the concept. Mindset/workflow/tool-choice topics should usually be `no`.
- **Stack**: `go` / `python` / `—` if no code. Stay within an in-progress series unless deliberately branching.
- **Tool**: `cc` / `web` / `both`.
- **Series**: empty `—` unless this topic continues an existing series like `go-service`.

## Step 4 — Plan-review step (mandatory)

Output the table and a one-line rationale per row explaining the angle and why it's distinct from existing topics. Then **stop**. Do not append to `topics.md` yet.

Wait for Bryan's approval, redirection, or substitution. Only after explicit approval, append the approved rows to `topics.md` and confirm with a one-line summary.

## Why this skill exists

The content library lives or dies on **uniqueness by angle**. Title-similarity is fine; angle-overlap is not. This skill enforces the gap analysis + review step that prevents the library from drifting into repetition.

See also: `CLAUDE.md` for the project's Content Philosophy section, especially the Uniqueness rule.
