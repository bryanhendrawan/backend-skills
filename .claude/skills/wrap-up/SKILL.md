---
name: wrap-up
description: End-of-session cleanup. Updates CLAUDE.md and README.md with any new patterns or skills added this session, writes session memory, commits all changes, and pushes to GitHub. Invoke when Bryan says "wrap up", "wrap things up", "end session", "close out", or similar phrases.
disable-model-invocation: false
---

Run this at the end of every session to close out cleanly.

## Step 1 — Detect what changed

Run `git diff --stat HEAD` and `git status` to see which files were modified this session. Identify:

- New skills added to `.claude/skills/` — need to be added to CLAUDE.md Skills table + README skills reference
- New slide components added to `write-content/SKILL.md` or `design-guidelines.md` — need to be reflected in CLAUDE.md component markers table
- New rules or constraints added anywhere — check CLAUDE.md constraints sections for gaps
- New day files added to `days/` — no CLAUDE.md update needed, just commit
- Changes to `topics.md` — no CLAUDE.md update needed, just commit

## Step 2 — Update CLAUDE.md

Only update sections that have actual gaps — do not rewrite sections already accurate.

- If new slide components were added: update the **Slide component markers** table.
- If new skills were added: update both the **Skills** table in CLAUDE.md and the **Skills reference** table in README.md.
- If new constraints or rules were added: update the relevant constraints section.
- If the repository structure changed materially (new top-level dirs, new skill dirs): update the **Repository Structure** block.

## Step 3 — Update README.md

If new skills were added, add them to the Skills reference table.

## Step 4 — Write session memory

Check `/Users/bryan.bryan/.claude/projects/-Users-bryan-bryan-go-src-backend-skills/memory/` for memories to create or update based on this session:

- Feedback Bryan gave (corrections, confirmations of non-obvious choices)
- Project decisions that aren't obvious from the code
- User preferences discovered

Write new memory files or update existing ones. Then update `MEMORY.md` index.

## Step 5 — Commit and push

Stage all modified and new tracked files. Write a concise commit message:

```
Session: [one-line summary of main changes — e.g. "add process-flow + chat-bubble components; reframe days 29–30 to company scenarios"]

Co-Authored-By: Claude Sonnet 4.6 <noreply@anthropic.com>
```

Push to origin. Confirm with: "Session wrapped — [N] files committed, pushed to origin."

## What NOT to do

- Do not amend existing commits — always create a new one.
- Do not push if there are no changes to commit — confirm and exit cleanly.
- Do not rewrite CLAUDE.md sections that are already accurate — targeted updates only.
