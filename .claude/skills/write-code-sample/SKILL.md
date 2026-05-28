---
name: write-code-sample
description: Generates a runnable, self-contained code sample for a day that has code=yes in topics.md. Supports series continuity by extending the previous day's code in the same series.
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Generate the code sample for the day identified by `$ARGUMENTS`. Output goes to `code/$ARGUMENTS/`.

## Step 1 — Validate

Read `topics.md` and find the row for `$ARGUMENTS`.

- If `Code` is `no`, abort with: `"Topic $ARGUMENTS has Code=no in topics.md. Nothing to write."`
- Note the `Stack` column (`go` / `python`) and `Series` column.

## Step 2 — Series continuity

If `Series` is set (e.g. `go-service`):

1. Find the most recent prior topic in the same series by scanning `topics.md`.
2. Read that day's code directory `code/<prior-slug>/` to understand current state.
3. **Extend it**, don't rewrite. Copy the prior code into `code/$ARGUMENTS/` as the starting point, then add only what this day teaches.
4. The diff between consecutive days in a series should read as a single coherent change, not a fresh start.

If `Series` is empty `—`, this is a standalone sample. Start from scratch.

## Step 3 — Write the sample

Write into `code/$ARGUMENTS/`:

- **Entry point**: `main.go` for Go, `main.py` for Python.
- **Module/deps file**: `go.mod` (Go) or `requirements.txt` (Python).
- **README.md**: include **only** if running the sample requires non-obvious steps (env vars, Docker, DB setup, MCP config). Skip it for trivial samples that run with one command.
- **Supporting files**: SQL migrations, configs, fixtures — only what's needed to make the sample run.

## Constraints

- **Self-contained**: zero external API keys. Zero hidden config. If a DB is needed, use SQLite (file-based, zero install) unless the topic specifically requires Postgres or MySQL.
- **Runnable with one command**: `go run .` or `python main.py`. If Docker is needed, document it in README.
- **Self-evident code**: well-named functions, no clever tricks. The reader is a backend engineer learning AI workflow, not Go/Python syntax.
- **No comments explaining what the code does** — names should do that. Comments are reserved for explaining a non-obvious *why* (a workaround, a constraint).
- **Proves the concept**: the code must demonstrate the topic, not be a toy snippet that hand-waves around it. A topic on migrations means a real migration that runs. A topic on CRUD means real endpoints you can curl.

## Step 4 — Confirm

After writing, output a one-line summary: which files were created, the command to run the sample, and (if series) what was extended from the prior day.

## Why this skill exists

Code samples in this library exist to prove that the AI workflow being taught actually works in real code, not as theory. The bar is: a backend engineer can clone the repo, `cd code/<day>`, run one command, and see the topic demonstrated. This skill enforces that bar.

See also: existing `code/` samples (once any exist) as reference, and `CLAUDE.md` for the production flow that integrates this skill with `/write-content`, `/caption`, and `/design-prompt`.
