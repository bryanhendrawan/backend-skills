---
name: Design Prompt
description: Generates a ready-to-paste Claude design prompt for a day carousel. Use when preparing a day file for Instagram and TikTok visual production.
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Read `days/$ARGUMENTS.md` and parse every slide from the file — use the `<!-- Slide N -->` markers to identify slide boundaries, and the `##` sections and numbered list items as slide content.

**Important**: stop parsing at the first horizontal rule (`---`) that appears at column 0 after the slides zone. Everything between that separator and the `## Design Prompt` section (if one exists) is the caption and hashtag block, managed by the `/caption` skill. If a `## Design Prompt` section already exists in the file, replace it entirely. If it does not exist, append it after the hashtags block.

Also read `design-guidelines.md` in this skill directory for the visual style specs.

**Em-dash rule**: replace every em-dash (—) in slide prose with a comma or period. Exception: em-dashes inside `<!-- cover-terminal: ... -->` markers are intentional flag characters — preserve them exactly.

**Special slide components** — detect these markers in the slides zone and translate them into explicit visual specs in the SLIDES section:

1. `<!-- cover-terminal: ~/path $ command —flags -->` — Cover slide terminal prompt. Output as:
   ```
   - Cover terminal prompt (left-aligned above title, monospace): ~/path $ command —flags
     Style: path in soft lavender, $ in cyan, command flags in white
   ```

2. `<!-- terminal-block: ~/path — label --> ... <!-- /terminal-block -->` — Terminal window block. Output as:
   ```
   - Terminal window block (left-aligned):
     - Panel: #141128 background, rounded 8px, border #1E1B3A, left 2px cyan accent bar
     - Window chrome: dots top-left (amber #E8955A, red #E85A5A, green #5AE87B)
     - Header: ~/path — label in soft lavender monospace
     For each row inside the block:
     - [TAG] rows: tag in amber/orange, backtick terms in cyan (#00D4FF), rest in white
     - ↳ rows: ↳ in cyan, text in muted lavender, indented
     - $ verdict/result rows: label in muted lavender, value in bold white
   ```

3. `<!-- comparison-window: Title --> ... <!-- /comparison-window -->` — Two-column bad/good panel. Output as:
   ```
   - Comparison window (left-aligned):
     - Panel: #141128, rounded 8px, border #1E1B3A, chrome dots top-left
     - Header: [Title] in soft lavender monospace
     - Left column header [BAD] in amber/orange; right column header [GOOD] in cyan
     - Column divider: thin line #1E1B3A
     For each [BAD] / [GOOD] row pair inside the block:
     - Label line: in column-header color (amber or cyan)
     - ↳ explanation: soft lavender, indented, ~16px monospace
   ```

4. `<!-- data-table: Title --> ... <!-- /data-table -->` — Styled data table. Output as:
   ```
   - Data table (left-aligned):
     - Panel: #141128, rounded 8px, border #1E1B3A
     - Title row: [Title] in soft lavender monospace, above the table
     - Header row: cyan (#00D4FF) bold uppercase, ~14px, bottom border #1E1B3A
     - Data rows: white ~16px; alternate row banding #141128 / #1A1735; generous cell padding
     For each markdown table row inside the block, render as a table row
   ```

**Global frame rule**: every slide in the SLIDES section must include this line:
```
Frame: [0N / total] top-left (monospace, soft lavender) | [• Backend x AI] top-right (soft lavender) | [@vilamas.ai] footer bottom-center (soft lavender)
```
Replace `0N` with the zero-padded slide number and `total` with the total slide count.

Generate the formatted prompt below, then **append or replace** the `## Design Prompt` section in `days/$ARGUMENTS.md`. Write directly to the file — do not output to terminal.

---

**Claude Design Prompt — [extract title from the # heading]**

Project folder name: `[ARGUMENTS value]`
Save all slides for this topic inside that folder. Each new topic gets its own folder.

Platform: Instagram (post/reels) + TikTok (image/video)
Aspect ratio: 4:5 portrait — 1080 × 1350 px per slide
Total slides: [count the slides]

[Insert the full contents of design-guidelines.md here verbatim]

---

**SLIDES:**

Note: every slide uses the global frame — counter top-left, series label top-right, @vilamas.ai footer bottom-center. All content is left-aligned and the content block is vertically centered on the canvas. Full spec in the Slide Structure section above.

**Slide 1 — Cover**
Frame: 01 / [total] top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
[cover-terminal spec if present]
- Heading: [title from # heading] — left-aligned, key word(s) in cyan italic
- Subheading: [theme from **Theme** line] — soft lavender, below heading
- Navigation hint: geser buat baca →→→ — small soft lavender, bottom-left above footer

**Slide 2 — The Problem**
Frame: 02 / [total] top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — THE PROBLEM (uppercase cyan, short cyan rule to the left)
- Headline: [content headline from ## The Problem] — H2, bold white, left-aligned
[terminal-block spec if present]

[For ## How It Works — each numbered step becomes its own slide:]
**Slide 3 — Step 1**
Frame: 03 / [total] top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — HOW IT WORKS (uppercase cyan, short cyan rule)
- Step badge: filled cyan circle 48–56px, solid #00D4FF, white "1" centered inside; badge sits to the left of the step text
- Body: [step 1 content] — 24–28px white, left-aligned to the right of the badge
[component spec if present]

**Slide 4 — Step 2**
Frame: 04 / [total] top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Step badge: filled cyan circle 48–56px, solid #00D4FF, white "2" centered inside
- Body: [step 2 content] — 24–28px white, left-aligned to the right of the badge
[component spec if present]

**Slide 5 — Step 3**
Frame: 05 / [total] top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Step badge: filled cyan circle 48–56px, solid #00D4FF, white "3" centered inside
- Body: [step 3 content] — 24–28px white, left-aligned to the right of the badge
[component spec if present]

[Continue for any additional steps]

**Slide [N-1] — Key Takeaway**
Frame: 0[N-1] / [total] top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — KEY TAKEAWAY (uppercase cyan, short cyan rule)
- Statement: [content from ## Key Takeaway] — H2, bold white, left-aligned
- Short cyan horizontal rule below the statement as decorative accent

**Slide [N] — Try It Today**
Frame: 0[N] / [total] top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — TRY IT TODAY (uppercase cyan, short cyan rule)
- Action: [content from ## Try It Today] — H2, bold white, left-aligned
- CTA row (below the action text, left-aligned):
  4 action chips in a horizontal row, equal spacing:
  - 🔖 Simpan — bookmark icon, label below, soft lavender
  - ❤️ Suka — heart icon, label below, soft lavender
  - 📤 Bagikan — share icon, label below, soft lavender
  - ➕ Follow @vilamas.ai — plus icon, label below, cyan (#00D4FF) to make it pop
  Each chip: subtle #141128 panel, rounded 8px, icon ~24px + label ~12px below
- Stronger cyan ambient glow on this slide vs. others

---
