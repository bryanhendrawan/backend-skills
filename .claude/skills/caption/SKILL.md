---
name: caption
description: Generates an Indonesian Instagram/TikTok caption and hashtag block for a day file, with SEO + GEO best practices baked in. Use after the day's carousel slides are written.
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Read `days/$ARGUMENTS.md` to understand the day's content — title, theme, problem, steps, takeaway, CTA.

Also read `caption-guidelines.md` in this skill directory for voice, structure, SEO, GEO, and hashtag rules.

Then write the caption + hashtags into the same day file. Behavior:

- If the file already has a `---` separator followed by `## Caption` and `## Hashtags`, **replace** that entire block.
- If it does not, **append** the block to the end of the file, preceded by `---`.

Generate **two separate captions** — one for TikTok and one for Instagram. They share the same topic but differ in length, rhythm, and CTA style. Do not reuse the same text for both.

Use this exact format:

```markdown
---

## Caption

**TikTok**
[TikTok caption — 150–300 chars, punchy, one idea per line, CTA invites comment or save]

**Instagram**
[Instagram caption — 150–400 chars, can be more narrative, CTA focuses on save/follow]

## Hashtags

**TikTok**
[8–10 hashtags space-separated, must include #fyp #FYPIndonesia, rest are niche + Indonesia geo]

**Instagram**
[10–12 hashtags space-separated, broad + niche tech + Indonesia-specific + topic-specific]
```

**Character rules (strict)**:
- Allowed punctuation: period, comma, question mark, exclamation mark, straight quote, hashtag symbol.
- Forbidden: em-dash, pipe, underscore, asterisk, parentheses, decorative symbols of any kind.
- The `**TikTok**` and `**Instagram**` labels above are markdown bold for display only — they do not appear in the caption text itself.

After writing, output:
1. Confirmation: "Caption replaced in days/$ARGUMENTS.md" (or "appended").
2. Both captions in full (for review in chat).
3. Both hashtag blocks.
4. Character count for each caption (excluding hashtags).

If iteration is needed, run the skill again — it replaces cleanly.
