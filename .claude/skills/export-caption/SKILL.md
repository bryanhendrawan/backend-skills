---
description: Extracts the caption and hashtags from a day file and saves them as plain text to output/caption/<slug>.txt — ready to paste into Instagram or TikTok. Usage: /export-caption YYYY-MM-DD_slug
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Extract the caption and hashtags from the day file and write to a plain text file.

1. Read `days/$ARGUMENTS.md`.

2. Find the **post zone** — the content between the first and second horizontal rules (`---`). It contains `## Caption` and `## Hashtags` sections.

3. Extract:
   - Caption text: everything under `## Caption`, excluding the header line itself
   - Hashtags: everything under `## Hashtags`, excluding the header line itself

4. Write to `output/caption/$ARGUMENTS.txt` in this exact format — plain text, no markdown headers:
   ```
   <caption text>

   <hashtags>
   ```

5. Create `output/caption/` if it does not exist.

6. Confirm the file path and show the first line of the caption as a preview.
