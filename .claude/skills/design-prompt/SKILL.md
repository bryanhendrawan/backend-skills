---
name: Design Prompt
description: Generates a ready-to-paste Claude design prompt for a day carousel. Use when preparing a day file for Instagram and TikTok visual production.
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Read `days/$ARGUMENTS.md` and parse every slide from the file — use the `<!-- Slide N -->` markers to identify slide boundaries, and the `##` sections and numbered list items as slide content.

**Important**: stop parsing at the first horizontal rule (`---`) that appears at column 0 after the slides zone. Everything between that separator and the `## Design Prompt` section (if one exists) is the caption and hashtag block, managed by the `/caption` skill. If a `## Design Prompt` section already exists in the file, replace it entirely. If it does not exist, append it after the hashtags block.

Also read `design-guidelines.md` in this skill directory for the visual style specs.

**Em-dash rule**: before writing any slide content, replace every em-dash (—) in the text with a comma or period — whichever reads more naturally. Em-dashes must not appear anywhere in the SLIDES section of the output.

Generate the formatted prompt below, then **append or replace** the `## Design Prompt` section in `days/$ARGUMENTS.md`. If a `## Design Prompt` section already exists, replace it entirely. If it does not exist, append it after the hashtags block. Do not output it to the terminal — write it directly to the file.

---

**Claude Design Prompt — [extract title from the # heading]**

Project folder name: `[ARGUMENTS value]`
Save all slides for this topic inside that folder. Each new topic gets its own folder.

Platform: Instagram (post/reels) + TikTok (image/video)
Aspect ratio: 9:16 portrait — 1080 × 1920 px per slide
Total slides: [count the slides]

[Insert the full contents of design-guidelines.md here verbatim]

---

**SLIDES:**

**Slide 1 — Cover**
- Heading: [title from # heading]
- Subheading: [theme from **Theme** line]

**Slide 2 — The Problem**
[content from ## The Problem section]

[For ## How It Works — each numbered step becomes its own slide:]
**Slide 3 — Step 1**
[step 1 content]

**Slide 4 — Step 2**
[step 2 content]

**Slide 5 — Step 3**
[step 3 content]

[Continue for any additional steps]

**Slide [N-1] — Key Takeaway**
[content from ## Key Takeaway]

**Slide [N] — Try It Today**
[content from ## Try It Today]

---
