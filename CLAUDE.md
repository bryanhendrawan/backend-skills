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
  YYYY-MM-DD_topic-name.md            # One file per day: slides + caption + design prompt
code/
  YYYY-MM-DD_topic-name/             # Code samples for days that need them (optional)
    main.py / main.go / ...           # Runnable, self-contained example
    README.md                         # How to run it (only if non-obvious)
output/
  claude-design-pdf/                  # PDFs exported from Canva/Claude Design
  images/                             # Per-slide PNGs from /pdf-to-images (gitignored)
  caption/                            # Plain-text captions from /export-caption (gitignored)
scripts/
  pdf_to_images.py                    # PDF → PNG converter (used by /pdf-to-images)
  upload_to_drive.py                  # Google Drive uploader (used by /upload-to-drive)
secrets/                              # Credentials — gitignored, never committed
  vilamas-gcloud-secret.json          # OAuth client secret for Google Drive
  token.json                          # Cached OAuth token (auto-generated on first run)
.claude/
  agents/
    backend-engineer.md               # Subagent with backend + AI domain knowledge
  agent-memory/
    backend-engineer/
      MEMORY.md                       # Auto-written by the agent; tracks topic coverage
  skills/
    generate-topics/SKILL.md          # /generate-topics
    write-content/SKILL.md            # /write-content
    write-code-sample/SKILL.md        # /write-code-sample
    design-prompt/                    # /design-prompt
      SKILL.md
      design-guidelines.md            # Visual style (edit to change look)
    caption/                          # /caption
      SKILL.md
      caption-guidelines.md           # Voice, SEO/GEO rules, hashtag mix
    export-caption/SKILL.md           # /export-caption
    pdf-to-images/SKILL.md            # /pdf-to-images
    upload-to-drive/SKILL.md          # /upload-to-drive — upload PNGs + captions to Drive
    wrap-up/SKILL.md                  # /wrap-up — end-of-session cleanup
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
| `write-content` | `/write-content <date_slug>` | Write the carousel slides for a given topic |
| `write-code-sample` | `/write-code-sample <date_slug>` | Generate the `code/` sample for a given day |
| `design-prompt` | `/design-prompt <date_slug>` | Emit a paste-ready Claude design prompt for the slides |
| `caption` | `/caption <date_slug>` | Generate TikTok + Instagram captions + hashtags and append to the day file |
| `export-caption` | `/export-caption <date_slug>` | Save caption + hashtags as plain text to `output/caption/` |
| `pdf-to-images` | `/pdf-to-images <YYYY-MM-DD>` | Convert an exported Canva/Claude Design PDF to per-slide PNGs |
| `upload-to-drive` | `/upload-to-drive <YYYY-MM-DD>` | Upload PNGs + caption txt files to Google Drive under `Content/{date}/` |
| `wrap-up` | `/wrap-up` | End-of-session: update CLAUDE.md + README, write memory, commit and push |

## Day File Format

Each `days/YYYY-MM-DD_topic-name.md` has **three zones**, separated by horizontal rules:

1. **Slides zone** (top) — the carousel script; parsed by `/design-prompt`.
2. **Post zone** (after first `---`) — Indonesian caption + hashtags, managed by `/caption`.
3. **Design prompt zone** (after second `---`) — paste-ready Claude design prompt, managed by `/design-prompt`.

```markdown
# [Topic Title]
<!-- Slide 1 -->
<!-- cover-terminal: ~/[context] $ [phrase] —[flag] —[flag] -->

**Theme**: [one-line hook]

## The Problem
<!-- Slide 2 -->

[Short punchy headline — 4–7 words]

<!-- terminal-block: ~/[context] — [label] -->
[ERR_01] [mistake + `code term`].
↳ [consequence ≤10 words]
[ERR_02] [mistake + `code term`].
↳ [consequence ≤10 words]
$ verdict: [conclusion ≤6 words].
<!-- /terminal-block -->

## How It Works

1. [Step one. ≤15 words.]
<!-- Slide 3 -->

2. [Step two. ≤15 words.]
<!-- Slide 4 -->

3. [Step three. ≤15 words.]
<!-- Slide 5 -->

## Key Takeaway
<!-- Slide 6 -->

**[One bold statement. ≤12 words.]**

## Try It Today
<!-- Slide 7 -->

[One specific action in the next 10 minutes. ≤15 words.]

---

## Caption
[Indonesian caption ~400 chars. Written by /caption.]

## Hashtags
[10–15 hashtags, space-separated.]

---

## Design Prompt
[Generated by /design-prompt — do not edit manually.]
```

**Slide component markers** — use HTML comments in the slides zone to flag structured visual blocks. `/design-prompt` translates these into pixel-level design specs:

| Marker | Visual output | Use on |
|---|---|---|
| `<!-- cover-terminal: ~/path $ cmd —flags -->` | Monospace prompt line above Cover title | Cover only |
| `<!-- terminal-block: ~/path — label --> ... <!-- /terminal-block -->` | Dark panel with `[ERR_XX]` rows, `↳` continuations, `$ verdict` | Problem slide |
| `<!-- comparison-window: Title --> ... <!-- /comparison-window -->` | Two-column bad/good panel | Step slides |
| `<!-- data-table: Title --> ... <!-- /data-table -->` | Styled alternating-row table | Step or Key Takeaway |
| `<!-- process-flow: Title --> ... <!-- /process-flow -->` | Pill-node flow diagram, left to right (max 5 nodes) | Workflow/sequence slides |
| `<!-- highlight-stat --> ... <!-- /highlight-stat -->` | Large single-number callout, no panel | Impact slides; once per day max |
| `<!-- code-snippet: lang — label --> ... <!-- /code-snippet -->` | Compact syntax-highlighted code panel (6–8 lines) | Steps proving concept with code |
| `<!-- metric-card --> ... <!-- /metric-card -->` | 2–3 compact impact cards side-by-side | Impact or Key Takeaway slides |
| `<!-- chat-bubble --> ... <!-- /chat-bubble -->` | AI prompt/response pair (1 prompt + 1 response) | Steps showing prompt → output |

**Component compactness rule**: panels are tight to their content — no blank lines inside a component block, no excess internal padding. Whitespace lives around the panel, not inside it.

**Slides-zone constraints** (non-negotiable):
- **Language**: Bahasa Indonesia, `kamu/aku` voice — professional but casual, not too streetwise. Technical terms stay in English: `backend`, `AI`, `Claude Code`, `prompt`, `code review`, `PR`, `migration`, etc. — don't force translation.
- **Company scenarios**: examples must be grounded in company or team work — not personal side projects or solo experiments. Pick a scenario a practicing engineer would recognise from their own job: a sprint task, a team process, an ops incident.
- **Title**: ≤8 words. Hook, not a chapter heading.
- **Slide body**: ≤2 sentences OR ≤3 bullets per slide — one idea per slide.
- **Steps**: ≤15 words each. Split a long step into two rather than going over.
- **No em-dashes** (—) anywhere in prose. Use a period or comma instead. Em-dashes inside `<!-- cover-terminal: ... -->` are intentional flag characters — preserve those.
- **No raw tables, nested lists, or blockquotes** in prose — use the component markers above for structured content instead.
- **Inline code**: ≤8 lines, illustrative only. Full code lives in `code/`.
- **No jargon without a one-word gloss** inline — e.g. "N+1 query (loop database tersembunyi)".
- **Closing slide**: Problem-Solution template ends with `## Try It Today`. Showdown template ends with `## Close` (combines the tease, action, and CTA in one slide — no separate Stay Tuned or Try It Today slides).

**Post-zone constraints**:
- Language: Bahasa Indonesia, `kamu/aku` tone — see `caption-guidelines.md`.
- **Two captions per day**: `**TikTok**` (150–300 chars) and `**Instagram**` (150–400 chars), each under the same `## Caption` header with its own `## Hashtags` block under `## Hashtags`.
- No special characters in captions: only period, comma, question mark, exclamation mark. No em-dash, pipe, asterisk, or parentheses.
- TikTok: punchy, one idea per line, must include `#fyp #FYPIndonesia`, 8–10 hashtags.
- Instagram: first line ≤125 chars (IG preview), 10–12 hashtags, save/follow CTA.
- SEO: front-load searched keywords (backend engineer, AI, Claude Code, etc.).
- GEO (Generative Engine Optimization): quotable claims, specific numbers, definitive phrasing — so AI search engines cite the content.
- **Fixed CTA text**: "Jangan lupa like, save, share, and follow buat tips berikutnya" — use this exact wording in design prompts.

## `topics.md` Format

Each entry records the full pipeline state for a topic.

```markdown
| Date | Slug | Title | Code | Stack | Tool | Series | Status | Post URL |
|---|---|---|---|---|---|---|---|---|
| 2026-05-27 | why-delegate-to-ai | Why backend engineers should learn to delegate to AI | no | — | both | — | posted | [IG](...) [TT](...) [FB](...) |
| 2026-05-28 | ai-invoice-processing | Invoice ratusan, tim cuma 2 orang | no | — | both | — | draft | — |
```

- **Post URL**: all platform links in one cell, formatted as `[IG](...) [TT](...) [FB](...)`. Never split into separate columns.
- **Status**: `draft` → `written` → `posted`.

## Working in This Repo

Full production flow for a single day:

1. **Topic** — confirm in `topics.md` (or propose with `/generate-topics`).
2. **Slides** — run `/write-content YYYY-MM-DD_slug` to draft the carousel.
3. **Code** (if needed) — run `/write-code-sample YYYY-MM-DD_slug`, then `/code-review` (linter, security, pentest scan, logic) before commit. Same workflow we teach on Day 5.
4. **Caption + hashtags** — run `/caption YYYY-MM-DD_slug` to append the Indonesian post body. Re-run anytime to regenerate.
5. **Design prompt** — run `/design-prompt YYYY-MM-DD_slug`, paste output into Claude design/Canva to produce the 4:5 slides.

Other tasks:
- **Reorder topics**: update the date column in `topics.md` and rename the affected files.
