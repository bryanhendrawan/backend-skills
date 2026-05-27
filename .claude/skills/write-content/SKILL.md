---
name: Write Content
description: Writes the slides zone of a day file in Bahasa Indonesia, following all CLAUDE.md slide constraints. Use after the topic exists in topics.md and before /caption and /design-prompt.
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Generate the carousel slides for the day identified by `$ARGUMENTS`. Output writes to `days/$ARGUMENTS.md` in the **slides zone** (everything above the first `---` separator).

## Step 1 — Read context

- `topics.md` — find the row for `$ARGUMENTS` to confirm angle, code flag, stack, series.
- `CLAUDE.md` — slides-zone constraints (these are non-negotiable; re-read them every run).
- Existing `days/$ARGUMENTS.md` if it exists — preserve the post zone (caption + hashtags) and design prompt zone untouched.
- If the topic is in a series (e.g. `go-service`), skim the prior day's file in that series so the language and references stay consistent.

## Step 2 — Outline first (plan-review step)

Before writing the file, draft the slide outline as a numbered list:

1. **Cover** — title (≤8 words) + subtitle (≤15 words).
2. **The Problem** — 1–2 sentences, specific pain.
3. **Step 1, 2, 3 (or 4)** — each ≤15 words.
4. **Key Takeaway** — one bold statement, ≤12 words.
5. **Try It Today** — one specific action takeable in the next 10–20 minutes.

Present the outline to Bryan and **stop**. Wait for approval or redirection before writing the file. Skip this step only if Bryan explicitly says "skip outline" or "just write it".

## Step 3 — Write the slides

After approval, write the slides zone using this structure (matches the Day 1 template):

```markdown
# [Title]
<!-- Slide 1 -->

**Theme**: [one-line subtitle]

## The Problem
<!-- Slide 2 -->

[1–2 sentences]

## How It Works

1. [Step one]
<!-- Slide 3 -->

2. [Step two]
<!-- Slide 4 -->

3. [Step three]
<!-- Slide 5 -->

## Key Takeaway
<!-- Slide 6 -->

**[one bold statement]**

## Try It Today
<!-- Slide 7 -->

[specific action]
```

The `<!-- Slide N -->` markers let `/design-prompt` parse slide boundaries.

## Constraints (non-negotiable)

- **Language**: Bahasa Indonesia, `kamu/aku` address — professional but casual, not `lo/gue` (too streetwise), not lecture-style.
- **Technical terms in English**: `backend`, `AI`, `Claude Code`, `prompt`, `code review`, `pull request`, `migration`, `MCP`, etc. Don't force translation.
- **Title**: ≤8 words. Hook, not chapter heading.
- **Each slide body**: ≤2 sentences OR ≤3 bullets. One idea per slide.
- **Each step**: ≤15 words. Split a long step into two rather than going over.
- **No em-dashes** (—) anywhere. Use a period or comma. If splitting helps, do that.
- **No tables, no nested lists, no blockquotes** — they don't survive the design step.
- **No "Yuk!", "Mari kita...", "Hai teman-teman", "kamu pasti bisa!"** — too tutorial / influencer.
- **Inline code**: ≤8 lines, illustrative only. Full code lives in `code/`.
- **Try It Today is mandatory** on every day.

## Step 4 — Replace the slides zone in the file

Read `days/$ARGUMENTS.md`. If the file exists, find the first `---` separator at column 0 and replace everything **above** it with the new slides zone. If the file does not exist, create it with just the slides zone (no `---` separator at the end yet — that's `/caption`'s job).

Never touch:
- The caption block (between the first and second `---`).
- The `## Design Prompt` section (after the second `---`).

## Why this skill exists

Producing a day's carousel by hand is fine for Day 1; doing it for Day 30 means slipping on constraints. This skill bakes the constraints in and forces the plan-review step that catches misfires before the file is written.

See also: `caption-guidelines.md` for tone reference, and the existing `days/2026-05-27_why-delegate-to-ai.md` as the canonical template.
