---
name: write-content
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

Pick a template variant (see Step 3) and draft the slide outline as a numbered list, with one line per slide.

**Problem-Solution template (default)**:
1. **Cover** — title (≤8 words) + subtitle (≤15 words).
2. **The Problem** — 1–2 sentences, specific pain.
3. **Step 1, 2, 3 (or 4)** — each ≤15 words.
4. **Key Takeaway** — one bold statement, ≤12 words.
5. *(optional)* **Stay Tuned** — one teaser line for an upcoming day.
6. **Try It Today** — one specific action takeable in the next 10–20 minutes.

**Showdown template**:
1. **Cover** — title (≤8 words) + subtitle (≤15 words).
2. **Masalah Lama** (or similar) — 1–2 sentences + `process-flow` showing the current painful process. Company scenario required.
3. **Dengan AI** — 1–2 sentences + `process-flow` showing the AI-assisted approach.
4. **The Diff** — short body + `comparison-window` with `[TANPA AI]` / `[DENGAN AI]` labels.
5. **The Receipt** — short body + `data-table` scorecard with concrete numbers.
6. **Close** — bold tease line (≤15 words, general AI practice) + specific Try It Today action. No separate Stay Tuned or Try It Today slides.

Present the outline to Bryan and **stop**. Wait for approval or redirection before writing the file. Skip this step only if Bryan explicitly says "skip outline" or "just write it".

## Step 3 — Write the slides

Pick the template that fits the topic, then write the slides zone. The two canonical templates are below — alternate between them across days so the feed doesn't feel repetitive.

### Template A — Problem-Solution (default)

Matches the Day 1 template:

```markdown
# [Title]
<!-- Slide 1 -->
<!-- cover-terminal: ~/[context] $ [action-phrase] —[flag] —[flag] -->

**Theme**: [one-line subtitle]

## The Problem
<!-- Slide 2 -->

[Short punchy headline — 4–7 words, bold claim]

<!-- terminal-block: ~/[context] — [label] -->
[ERR_01] [first mistake as a label + inline `code term`].
↳ [consequence in 1 line, muted]
[ERR_02] [second mistake as a label + inline `code term`].
↳ [consequence in 1 line, muted]
$ verdict: [conclusion].
<!-- /terminal-block -->

## How It Works

1. [Step one]
<!-- Slide 3 -->

2. [Step two]
<!-- Slide 4 -->

<!-- comparison-window: [Title] -->
[BAD] [bad example label with `code term`].
↳ [consequence ≤10 words]
[GOOD] [good example label with `code term`].
↳ [benefit ≤10 words]
<!-- /comparison-window -->

3. [Step three]
<!-- Slide 5 -->

<!-- data-table: [Title] -->
| [Col A] | [Col B] |
|---|---|
| [row 1a] | [row 1b] |
| [row 2a] | [row 2b] |
| [row 3a] | [row 3b] |
<!-- /data-table -->

## Key Takeaway
<!-- Slide 6 -->

**[one bold statement]**

## Try It Today
<!-- Slide 7 -->

[specific action]
```

### Template B — Showdown

Use when the topic is a direct before/after comparison of a real company process. Canonical example: `days/2026-05-28_ai-invoice-processing.md`.

```markdown
# [Title]
<!-- Slide 1 -->
<!-- cover-terminal: ~/[company-context] $ [phrase] —[side-a] —[side-b] -->

**Theme**: [one-line subtitle. Frame as a real company or team scenario — not a personal experiment.]

## [Nama masalah — e.g. "Masalah Lama", "Cara Lama"]
<!-- Slide 2 -->

[1–2 sentences describing the current painful process at the company.]

<!-- process-flow: [Title] -->
[STEP_01] [step description, ≤4 words]
[STEP_02] [step description, ≤4 words]
[STEP_03] [step description, ≤4 words]
[STEP_04] [step description, ≤4 words]
<!-- /process-flow -->

## Dengan AI
<!-- Slide 3 -->

[1–2 sentences describing how AI changes the approach.]

<!-- process-flow: [Title] -->
[STEP_01] [step description, ≤4 words]
[STEP_02] [step description, ≤4 words]
[STEP_03] [step description, ≤4 words]
<!-- /process-flow -->

## The Diff
<!-- Slide 4 -->

[Short body, ≤15 words. Reframe the comparison beyond speed.]

<!-- comparison-window: [Title] -->
[TANPA AI] [what "without AI" costs — time, energy, risk].
↳ [honest consequence ≤10 words]
[DENGAN AI] [what "with AI" unlocks or shifts].
↳ [honest benefit ≤10 words]
<!-- /comparison-window -->

## The Receipt
<!-- Slide 5 -->

[Short body, ≤15 words.]

<!-- data-table: [Title] -->
| Metrik | Tanpa AI | Dengan AI |
|---|---|---|
| [metric 1] | [value] | [value] |
| [metric 2] | [value] | [value] |
| [metric 3] | [value] | [value] |
<!-- /data-table -->

## Close
<!-- Slide 6 -->

**[Bold tease — one sentence, general AI practice or team pattern, ≤15 words. No brand, no day reference.]**

[Try It Today action — one specific sentence. Direct and energetic, no "Yuk!" opener.]
```

The `<!-- Slide N -->` markers let `/design-prompt` parse slide boundaries.

### Close slide (combined closing)

Both templates end with `## Close` — a single slide that combines the tease, the Try It Today action, and the CTA. This replaces the old "Stay Tuned" + "Try It Today" two-slide pattern.

```
## Close
<!-- Slide N -->

**[Bold tease — one sentence, general AI practice or team pattern, ≤15 words. No brand, no day reference.]**

[Try It Today action — one specific sentence. Direct and energetic, no "Yuk!" opener.]
```

The design step renders this as two stacked zones: tease in soft lavender italic + a divider + action in bold white H2 + CTA text below.

### Vary the CTA copy

The static CTA boilerplate ("Yuk, save, like, share, dan follow untuk content part berikutnya.") was identical on Days 1–7. The **closing CTA wording varies per day** — handled at the `/design-prompt` step. Keep the action statement specific and energetic.

### Process flow

Use `<!-- process-flow -->` when the slide's main concept is a sequence of steps — a workflow, pipeline, or company process. Renders as pill nodes connected by arrows, left to right. Prefer this over `terminal-block` for workflow slides.

```
<!-- process-flow: Title -->
[STEP_01] [step description, ≤4 words]
[STEP_02] [step description, ≤4 words]
[STEP_03] [step description, ≤4 words]
<!-- /process-flow -->
```

**Process flow constraints**: step labels `[STEP_01]` through `[STEP_05]`. Step text ≤4 words. Max 5 steps — more is unreadable at this canvas width.

### Highlight stat

Use `<!-- highlight-stat -->` when one number tells the whole story — time saved, error reduction, scale. The number dominates the slide; everything else supports it. Use sparingly — once per day maximum.

```
<!-- highlight-stat -->
[NUMBER] 5x
[LABEL] lebih cepat dari proses manual
[CONTEXT] untuk tim dengan 50 invoice per hari
<!-- /highlight-stat -->
```

**Highlight stat constraints**: `[NUMBER]` is the big stat (e.g., "5x", "80%", "3 menit"). `[LABEL]` ≤6 words. `[CONTEXT]` ≤8 words.

### Code snippet
Use `<!-- code-snippet -->` when the slide's key proof IS the code — a real config file, migration SQL, small handler. Renders as a compact syntax-highlighted dark panel. Max 6–8 lines; trim to the most informative lines. No blank lines inside the block.

```
<!-- code-snippet: language — real-context-label -->
code here (6–8 lines max, no blank lines)
<!-- /code-snippet -->
```

The label must name the real context: `CLAUDE.md tim finance`, `migration users table`, `handler POST /invoices` — not `example` or `code`. Use instead of `terminal-block` when showing actual code, not command output.

**Code snippet constraints**: 6–8 lines max. No blank lines inside the block. Label names real context. Language is the syntax (e.g., `go`, `sql`, `markdown`).

### Metric card
Use `<!-- metric-card -->` when 2–3 numbers together tell the impact story. Cards sit side-by-side. Use this over `highlight-stat` when multiple metrics carry equal weight.

```
<!-- metric-card -->
[METRIC] label text | value
[METRIC] label text | value
[METRIC] label text | value
<!-- /metric-card -->
```

Real company example: `[METRIC] PR review cycle | dari 2 hari ke 4 jam`

**Metric card constraints**: 2–3 cards max. Label (before `|`) ≤5 words. Value (after `|`) ≤5 words. Use on impact or key-takeaway slides.

### Chat bubble
Use `<!-- chat-bubble -->` when the motivating insight is a direct prompt-to-output exchange — "this is literally what you type, this is what AI gives you." One [PROMPT] + one [RESPONSE] only. Keep both lines short so the panel stays compact.

```
<!-- chat-bubble -->
[PROMPT] one-line prompt — ≤15 words
[RESPONSE] one-line AI output — ≤15 words
<!-- /chat-bubble -->
```

Real company example: `[PROMPT] Tambah kolom email NOT NULL UNIQUE ke tabel users, lengkap dengan rollback.`

**Chat bubble constraints**: exactly 1 [PROMPT] + 1 [RESPONSE]. Each ≤15 words. Use on Step slides where the core insight is showing what you type and what you get.

### Cover terminal prompt
Every Cover slide gets a `<!-- cover-terminal: ... -->` marker. Choose a path and command phrase that frames the topic — e.g. `~/senior-eng $ think —before —you —type`, `~/prod $ diagnose —slow —query`, `~/ci $ review —diff —context`. The phrase should read like a shell command that expresses the core idea.

### Terminal window block
Use `<!-- terminal-block -->` when the content fits a 2–3 row labeled list. Each row gets a bracketed label, a short description with a backtick-wrapped keyword, and a `↳` consequence line. Close with `$ verdict:`. If the content is a single continuous idea, use plain prose and omit the block.

Pick the label that matches the row's semantic:
- `[ERR_01]`, `[ERR_02]` — for the Problem slide (mistakes, anti-patterns).
- `[LOG_01]`, `[LOG_02]` — for Showdown's Round slides (stats, receipts).
- Other label families (e.g. `[STAT_XX]`, `[RUN_XX]`) are fine as long as both rows in the block use the same family.

**Terminal block constraints**: backtick-wrapped terms stay in English. `↳` lines ≤10 words. Verdict ≤6 words.

### Comparison window
Use `<!-- comparison-window -->` when the concept is sharpest as a two-side contrast. Each side gets a label line and one `↳` explanation line. Limit to 2 pairs max per block.

Pick labels that match the semantic:
- `[BAD]` / `[GOOD]` — for the default Problem-Solution template (anti-pattern vs. recommended).
- `[TANPA AI]` / `[DENGAN AI]`, `[BEFORE]` / `[AFTER]`, `[SIDE_A]` / `[SIDE_B]` — for Showdown or any head-to-head framing where one side isn't strictly "wrong". Never use a brand name as a label.

**Comparison window constraints**: labels ≤8 words. `↳` lines ≤10 words.

### Data table
Use `<!-- data-table -->` on a step or Key Takeaway slide when comparing 2–4 items across consistent dimensions. 2–3 columns max, 2–4 data rows max.

**Data table constraints**: cell text ≤6 words. Header labels ≤3 words. Use English for column headers.

## Constraints (non-negotiable)

- **Company scenarios**: examples and scenarios must be grounded in company or team work — not personal side projects, weekend experiments, or solo hobbies. When choosing a concrete scenario, pick something a practicing backend engineer at a company would recognise from their own job: a sprint task, a team process, an ops problem, a production incident.

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
