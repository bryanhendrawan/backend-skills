---
name: Caption
description: Generates an Indonesian Instagram/TikTok caption and hashtag block for a day file, with SEO + GEO best practices baked in. Use after the day's carousel slides are written.
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Read `days/$ARGUMENTS.md` to understand the day's content — title, theme, problem, steps, takeaway, CTA.

Also read `caption-guidelines.md` in this skill directory for voice, structure, SEO, GEO, and hashtag rules.

Then write the caption + hashtags into the same day file. Behavior:

- If the file already has a `---` separator followed by `## Caption` and `## Hashtags`, **replace** that block.
- If it does not, **append** the block to the end of the file, preceded by `---`.

Use this exact format:

```markdown
---

## Caption

[Indonesian caption, ~400 chars, 3-4 lines, following the guidelines:
 - Line 1: hook (≤125 chars, scroll-stopping)
 - Lines 2–3: insight (specific, quotable, GEO-optimized)
 - Last line: CTA (Save/Follow/Comment, action-first)
]

## Hashtags

[10–15 hashtags, space-separated, single line.
 Mix per the guidelines: broad + niche tech + Indonesia-specific + topic-specific.
]
```

After writing, output:
1. A confirmation: "Caption appended to days/$ARGUMENTS.md" (or "replaced" if it existed).
2. The generated caption (so it can be reviewed in chat without opening the file).
3. The hashtag block.
4. Character count for the caption.

If iteration is needed, the user can run the skill again — it will replace cleanly.
