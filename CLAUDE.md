# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Purpose

This repo is an **ongoing, ever-growing content library** for social media and mini ebook format, targeted at backend software engineers. The central theme is: *"How backend software engineers use AI in their daily work."*

Each day = one self-contained topic. Content runs indefinitely — the first 30 days are the initial batch, but new topics are added continuously. The audience is practicing backend engineers, not beginners to programming.

## Content Philosophy

- **Audience assumption**: readers already know how to code; they need to know *how AI changes or accelerates their existing workflow*.
- **Progression**: not strictly linear — topics can be ordered by what fits best; deep-dives can follow foundational pieces and vice versa.
- **Uniqueness**: the most critical constraint. Every topic must cover a distinct angle. Before proposing any new topic, scan `topics.md` to confirm no overlap in angle, not just title.
- **Tone**: professional but casual — like a thoughtful conversation between two senior engineers who respect each other. Think pair programming or a Slack DM, not a tutorial blog or hype thread.
- **Code samples**: topics that benefit from working code include one. Code must be self-contained, runnable, and prove the concept — not a toy snippet.

## Repository Structure

```
topics.md                              # Master topic index — source of truth for all topics
days/
  YYYY-MM-DD_topic-name.md            # One file per day: slides + caption + hashtags
code/
  YYYY-MM-DD_topic-name/             # Code samples for days that need them (optional)
    main.py / main.go / ...           # Runnable, self-contained example
    README.md                         # How to run it (only if non-obvious)
.claude/
  agents/
    backend-engineer.md               # Subagent with backend + AI domain knowledge
  agent-memory/
    backend-engineer/
      MEMORY.md                       # Auto-written by the agent; tracks topic coverage
  skills/
    generate-topics/SKILL.md          # Skill: propose next batch of unique topics
    write-content/SKILL.md            # Skill: write a day's carousel slides
    write-code-sample/SKILL.md        # Skill: generate the code sample for a day
    design-prompt/                    # Skill: emit Claude design prompt for slides
      SKILL.md
      design-guidelines.md            # Visual style (edit to change look)
    caption/                          # Skill: write Indonesian caption + hashtags
      SKILL.md
      caption-guidelines.md           # Voice, SEO/GEO rules, hashtag mix
```

- `topics.md` is the canonical record of every topic planned or published. Always check it before proposing new topics.
- Day files use `YYYY-MM-DD` prefix so they sort chronologically and the date records intended publish order.
- The `code/` directory mirrors the `days/` naming — only created when a topic needs working code.

## Session Flow

Every session in this repo follows this interaction pattern:

1. **Greet** — the agent opens with a casual check-in ("how's your day?") to set a collaborative tone.
2. **Orient** — the agent reads `topics.md` and its own memory to understand what has already been covered and what gaps exist.
3. **Topic discussion** — one of two paths:
   - Agent proactively suggests 2–3 topic ideas based on gaps and recent coverage, and asks Bryan to pick or redirect.
   - Bryan has a topic in mind already — agent validates it for uniqueness and fit.
4. **Plan** — agent lays out the plan: topic angle, whether code is needed, rough outline of the content.
5. **Review** — agent presents the plan and waits for Bryan's approval or feedback before writing anything.
6. **Generate** — only after approval: write the day file, write the code sample if needed, update `topics.md`.

The agent never jumps straight to generating content. The plan-review step is mandatory.

## Agent: `backend-engineer`

Defined at `.claude/agents/backend-engineer.md`. Format follows the Claude docs standard:

```markdown
---
name: backend-engineer
description: Backend engineering domain expert. Opens each session with a check-in, suggests topics, plans content, and waits for approval before generating anything.
tools: Read, Write, Edit, Glob, Grep, Bash
memory: project
---

[System prompt: the agent's full instructions, knowledge scope, and how it uses its memory]
```

The agent uses `memory: project` so it accumulates knowledge in `.claude/agent-memory/backend-engineer/MEMORY.md` — tracking which topics have been used, which code patterns have been shown, and coverage gaps. Claude writes and maintains this file automatically.

The agent's memory should track:
- Every topic covered (slug + angle, not just title) — for uniqueness checking.
- Which backend domains have been covered heavily vs. lightly (e.g., databases, APIs, observability, CI/CD).
- Code languages used in samples — to vary them over time.
- Any feedback Bryan gave on tone, depth, or direction.

## Skills

Each skill lives in `.claude/skills/<name>/SKILL.md`. Format:

```markdown
---
description: What the skill does and when Claude should invoke it.
disable-model-invocation: true   # include only for user-only workflows
argument-hint: <hint>            # include only if the skill takes arguments
---

Skill instructions here.
```

| Skill | Command | When to use |
|---|---|---|
| `generate-topics` | `/generate-topics` | Propose the next batch of N unique topics, checked against `topics.md` |
| `write-content` | `/write-content` | Write the carousel slides for a given topic |
| `write-code-sample` | `/write-code-sample` | Generate the `code/` sample for a given day |
| `design-prompt` | `/design-prompt <date_slug>` | Emit a paste-ready Claude design prompt for the slides |
| `caption` | `/caption <date_slug>` | Generate Indonesian caption + hashtags (SEO + GEO) and append to the day file |

## Day File Format

Each `days/YYYY-MM-DD_topic-name.md` has **two zones**, separated by a horizontal rule:

1. **Slides zone** (top) — the carousel script fed into Claude design.
2. **Post zone** (below `---`) — the Indonesian caption and hashtags that go in the IG/TikTok post body.

The `/design-prompt` skill reads only the slides zone. The `/caption` skill writes/replaces the post zone.

```markdown
# [Topic Title]
<!-- Slide 1 cover. ≤8 words. Scroll-stopping — the only thing seen before swiping. -->

**Theme**: [one-line hook]
<!-- Slide 1 subtitle. ≤15 words. Answers "why should I care?" -->

## The Problem
[1–2 sentences max. ≤20 words each. Make the pain specific and real.]

## How It Works
1. [Step one. ≤15 words.]
2. [Step two. ≤15 words.]
3. [Step three. ≤15 words.]
<!-- Each numbered step = its own slide. 3–5 steps max. -->

## Key Takeaway
[One bold statement. ≤12 words. The thing they remember after closing the app.]

## Try It Today
[One specific action they can take in the next 10 minutes. ≤15 words.]
<!-- Mandatory CTA slide — every day must end with this. -->

---

## Caption
[Indonesian caption ~400 chars. 3–4 lines: hook → insight → CTA.
 Written by /caption skill following caption-guidelines.md.]

## Hashtags
[10–15 hashtags, space-separated. Mix of broad, niche tech, and Indonesia-specific.]
```

**Slides-zone constraints** (non-negotiable):
- **Language**: Bahasa Indonesia, `kamu/aku` voice — professional but casual, not too streetwise. Technical terms stay in English: `backend`, `AI`, `Claude Code`, `prompt`, `code review`, `PR`, `migration`, etc. — don't force translation.
- **Title**: ≤8 words. Hook, not a chapter heading.
- **Slide body**: ≤2 sentences OR ≤3 bullets per slide — one idea per slide.
- **Steps**: ≤15 words each. Split a long step into two rather than going over.
- **No tables, no nested lists, no blockquotes** — don't survive the design step.
- **Inline code**: ≤8 lines, one short illustrative snippet only. Full code lives in `code/`.
- **No jargon without a one-word gloss** inline — e.g. "N+1 query (loop database tersembunyi)".
- **"Try It Today" slide is mandatory** on every day file.

**Post-zone constraints**:
- Language: Bahasa Indonesia, `kamu/aku` tone — see `caption-guidelines.md`.
- Caption ~400 chars total, first line ≤125 chars (IG preview snippet).
- SEO: front-load searched keywords (backend engineer, AI, Claude Code, etc.).
- GEO (Generative Engine Optimization): quotable claims, specific numbers, definitive phrasing — so AI search engines cite the content.
- 10–15 hashtags, mix of broad + niche + Indonesia-specific.

## `topics.md` Format

Each entry records: the publish date, slug, title, and whether it has a code sample.

```markdown
| Date | Slug | Title | Code |
|---|---|---|---|
| 2025-01-01 | using-llm-for-code-review | Using LLMs for Automated Code Review | yes |
| 2025-01-02 | ai-generated-migrations | Generating Database Migrations with AI | no |
```

## Working in This Repo

Full production flow for a single day:

1. **Topic** — confirm in `topics.md` (or propose with `/generate-topics`).
2. **Slides** — run `/write-content YYYY-MM-DD_slug` to draft the carousel.
3. **Code** (if needed) — run `/write-code-sample YYYY-MM-DD_slug`, then `/code-review` (linter, security, pentest scan, logic) before commit. Same workflow we teach on Day 5.
4. **Caption + hashtags** — run `/caption YYYY-MM-DD_slug` to append the Indonesian post body. Re-run anytime to regenerate.
5. **Design prompt** — run `/design-prompt YYYY-MM-DD_slug`, paste output into Claude design/Canva to produce the 9:16 slides.

Other tasks:
- **Reorder topics**: update the date column in `topics.md` and rename the affected files.
