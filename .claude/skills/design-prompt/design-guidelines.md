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

- One idea per slide — generous whitespace, nothing crowded
- **Cover**: cover-terminal prompt above H1 title; H1 left-aligned, key word in cyan italic; theme subheading below in soft lavender; `geser buat baca →→→` bottom-left in small soft lavender
- **Problem slide**: section label + H2 headline + terminal window block below
- **Step slides**: filled cyan circle badge (48–56px diameter, solid cyan #00D4FF, white step number centered inside) to the left of the step text; body text right of badge, left-aligned; optional component block below
- **Key Takeaway**: H2 headline left-aligned; large bold statement; short cyan horizontal rule below as accent
- **CTA (Try It Today)**: stronger cyan ambient glow; bold H2 action statement left-aligned; CTA row below — 4 horizontal action chips: Simpan | Suka | Bagikan | Follow @vilamas.ai (each chip: icon + label, soft lavender, equal spacing)

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

## Tone
- Bold and modern — premium AI product feel
- Terminal/window elements are used selectively — they signal "engineer's context", not decoration
- High contrast, vibrant but not noisy — one strong color moment per slide
- Ambient glows give depth; keep type clean and uncluttered on top
