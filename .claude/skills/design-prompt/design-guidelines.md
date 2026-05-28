# Design Guidelines

Visual style for the "Backend Engineer × AI" carousel series.
Edit this file to change the look — the skill picks it up automatically.

## Theme
- Background: deep midnight navy (#0B0A1F) — dark with a purple lean, not plain black
- Ambient glow: faint cyan radial glow from top-right, faint violet radial glow from bottom-left — apply on every slide for premium depth
- Grid backdrop: subtle grid overlay, slightly dialed down so gradients read clearly
- Primary text: pure white (#FFFFFF)
- Accent: cyan (#00D4FF) — step numbers, keywords, highlights, badges — pops against the navy base
- Muted text: soft lavender (#A5A3C2) — subtitles, labels, secondary info
- Body text: clean sans-serif (Inter or similar)
- Code / monospace: Fira Code, JetBrains Mono, or similar

## Typography Scale

All sizes for 1080 × 1350 px canvas. Never compress type — use whitespace instead.

| Role | Size | Weight | Color | Font |
|---|---|---|---|---|
| Cover title (H1) | 72–80px | Bold | White | Sans-serif |
| Slide headline (H2) | 44–52px | Bold | White | Sans-serif |
| Body / step text | 24–28px | Regular | White | Sans-serif |
| Section label | 13–14px | Uppercase, letter-spaced | Cyan (#00D4FF) | Sans-serif |
| Subheading / theme line | 18–22px | Regular | Soft lavender | Sans-serif |
| Terminal / code text | 16–20px | Regular | Varies by role | Monospace |
| Slide counter | 16–18px | Regular | Soft lavender | Monospace |
| Footer handle | 14–16px | Regular | Soft lavender | Sans-serif |

Clear hierarchy rule: H1 is ~3× body. H2 is ~2× body. Labels and footer are visually recessed — never competing with body text.

## Slide Structure (Global Frame)

Every slide — without exception — uses this three-zone layout:

**Top bar**
- Left: slide counter `01 / 07` — monospace, 16–18px, soft lavender (#A5A3C2)
- Right: series label `• Backend x AI` — 16–18px, soft lavender (#A5A3C2)

**Content zone**
- All text and components are left-aligned to a consistent left margin (~80px from edge)
- The content block as a whole is **vertically centered** on the canvas between the top bar and footer — equal whitespace above and below the block
- Section label (non-cover slides): ~14px uppercase cyan, short cyan horizontal rule to its left, above the slide headline

**Footer**
- `@vilamas.ai` — bottom center, 14–16px, soft lavender (#A5A3C2)
- Present on every slide, no exceptions

## Layout

- One idea per slide — generous whitespace around the content block; the slide breathes
- **Component panels are compact**: internal padding inside panels (terminal-block, comparison-window, data-table, process-flow, code-snippet, chat-bubble, metric-card) is tight — no blank lines within a panel, no excess top/bottom padding. The whitespace lives outside the panel, not inside it.
- **Cover**: cover-terminal prompt above H1 title; H1 left-aligned, key word in cyan italic; theme subheading below in soft lavender; `geser buat baca →→→` bottom-left in small soft lavender
- **Problem slide**: section label + H2 headline + terminal window block below
- **Step slides**: filled cyan circle badge (48–56px diameter, solid cyan #00D4FF, white step number centered inside) to the left of the step text; body text right of badge, left-aligned; optional component block below
- **Key Takeaway**: H2 headline left-aligned; large bold statement; short cyan horizontal rule below as accent
- **CTA / Try It Today** (Problem-Solution template): stronger cyan ambient glow; bold H2 action statement left-aligned; CTA text below — soft lavender, 16–18px, varied per day (not boilerplate). No chips or icons.
- **Close slide** (Showdown template): two stacked zones separated by a short cyan horizontal rule. Top zone: section label "WHAT'S NEXT" + tease line in soft lavender italic, 18–22px. Bottom zone: section label "TRY IT TODAY" + bold white H2 action + CTA text in soft lavender, 16–18px. Stronger cyan ambient glow.

## Branding
- Series label: top-right every slide — see Slide Structure above
- No stock photos — icons and typography only
- Footer handle: bottom center every slide — see Slide Structure above

## Slide Component Types

Use these purpose-built components when content benefits from structured visuals. Not every slide needs one.

### Cover Terminal Prompt
A monospace command line above the cover title, framing the topic as an engineering context.
- Format: `~/[context] $ [action-phrase-as-flags]`
- Styling: path in soft lavender, `$` in cyan, command text in white — all monospace, small
- Placement: left-aligned above the H1 title
- Em-dashes as flag characters (`—before —you —type`) are intentional — preserve exactly
- Use on: Cover slide only

### Terminal Window Block
A floating dark panel for errors, logs, or structured step info.
- Panel: background #141128, rounded corners (8px), border (#1E1B3A)
- Left edge: thin 2px cyan vertical accent bar
- Window chrome: three dots top-left — amber (#E8955A), red (#E85A5A), green (#5AE87B) — 10–12px
- Header: path/context string in soft lavender monospace, e.g. `~/ai-usage — debugging`
- Content rows (monospace, 16–18px):
  - Tag labels `[ERR_01]`, `[ERR_02]` — amber/orange
  - Backtick-wrapped terms — cyan (#00D4FF)
  - Continuation: `↳` in cyan, description in muted lavender, indented
  - Conclusion: `$ verdict:` in muted lavender, value in bold white
- Use on: The Problem slide, or any slide with 2–3 discrete structured errors

### Comparison Window
A two-column floating panel for before/after or bad/good contrasts.
- Panel: same background and chrome as Terminal Window Block
- Header: comparison title in soft lavender monospace
- Left column header: `[BAD]` in amber/orange
- Right column header: `[GOOD]` in cyan (#00D4FF)
- Column divider: thin vertical line (#1E1B3A)
- Row content: label in column-header color, explanation below in soft lavender, ~16px monospace
- Use on: Step slides where a bad/good contrast sharpens the concept

### Data Table
A minimal styled table for structured comparisons — 2–4 rows, 2–3 columns max.
- Panel: background #141128, rounded corners, border (#1E1B3A)
- Header row: cyan (#00D4FF) text, bold, ~14px uppercase, bottom border in #1E1B3A
- Data rows: white text ~16px; alternate rows use #141128 and #1A1735 for subtle banding
- Cell padding: generous — never cramped
- Use on: Any slide comparing 2–4 distinct values across consistent dimensions

### Process Flow
A horizontal flow diagram for visualizing sequential steps — workflows, pipelines, company processes.
- Panel: background #141128, rounded corners (8px), border (#1E1B3A)
- Step nodes: #1A1735 background, 8px radius, thin border (#1E1B3A)
- Step label `[STEP_XX]`: small cyan monospace, above or inside the node
- Step text: white, 14–16px, max 3–4 words per node
- Arrows between nodes: thin cyan (#00D4FF) right-pointing arrow, vertically centered between nodes
- Flow title: soft lavender monospace, above the panel
- Max 5 nodes — more becomes unreadable at 1080px canvas width
- Use on: Any slide where a sequence of steps is the main concept

### Highlight Stat
A large floating stat callout — no panel, the number dominates the slide.
- `[NUMBER]`: 90–110px, bold, cyan (#00D4FF) — the dominant visual element
- `[LABEL]`: 24–28px, bold white, directly below the number
- `[CONTEXT]`: 16–18px, soft lavender, below the label
- Optional: faint cyan glow halo around the number for depth
- Use on: Any slide where one number tells the whole story (savings, error reduction, scale)
- One per day maximum — use sparingly so it stays impactful

### Code Snippet
A compact syntax-highlighted code panel. Use when the slide's proof IS the code itself.
- Panel: background #141128, rounded corners (8px), border (#1E1B3A), window chrome dots top-left
- Label: soft lavender monospace, above the panel — must name the real context (e.g., "CLAUDE.md tim finance", "migration users table")
- Code text: monospace 15–18px; keywords in cyan (#00D4FF); strings in amber (#E8955A); comments in soft lavender; rest in white
- **Compact**: panel height is tight to the content — no excess vertical padding; no blank lines rendered between code lines
- Max 6–8 visible lines — trim to the most informative lines; the panel should feel dense and purposeful
- Use instead of Terminal Window Block when showing actual source code, config, or SQL (not command output)

### Metric Card
A horizontal row of 2–3 compact impact cards, each showing one metric.
- Layout: 2–3 equal-width cards in a tight row, stretching to fill content width; ~20px gap between cards
- Each card: background #1A1735, rounded 12px, thin border (#1E1B3A); compact internal padding
- Label (text before `|`): 13–15px white, positioned above the value
- Value (text after `|`): 32–40px bold cyan (#00D4FF) — dominant element inside each card
- Use when 2–3 numbers carry equal weight (e.g., time saved + error rate + cycle time); distinct from Highlight Stat which features one dominant number

### Chat Bubble
A two-row panel showing an AI prompt/response exchange. Use when "this is what you type → this is what you get" is the motivating insight.
- Panel: background #141128, rounded corners (8px), border (#1E1B3A)
- [PROMPT] row: "You" label in amber (#E8955A) 12px monospace; prompt text in white 16px; 2px amber vertical left accent bar; no extra padding above
- [RESPONSE] row: "AI" label in cyan (#00D4FF) 12px monospace; response text in white 16px; 2px cyan vertical left accent bar
- Gap between rows: 8px; tight overall vertical padding — panel height matches 2 rows + minimal spacing
- Exactly 1 [PROMPT] + 1 [RESPONSE] per block; panel stays compact, not sprawling
- Use on Step slides where the key insight is the directness of the interaction

## Tone
- Bold and modern — premium AI product feel
- Terminal/window elements are used selectively — they signal "engineer's context", not decoration
- High contrast, vibrant but not noisy — one strong color moment per slide
- Ambient glows give depth; keep type clean and uncluttered on top
