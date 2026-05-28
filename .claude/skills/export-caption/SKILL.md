---
description: Extracts the caption and hashtags from a day file and saves them as plain text to output/caption/<slug>.txt — ready to paste into Instagram or TikTok. Usage: /export-caption YYYY-MM-DD_slug
argument-hint: <YYYY-MM-DD_slug>
disable-model-invocation: false
---

Extract the caption and hashtags from the day file and write to plain text files.

1. Read `days/$ARGUMENTS.md`.

2. Find the **post zone** — the content between the first and second horizontal rules (`---`). It contains `## Caption` and `## Hashtags` sections.

3. Detect the caption format:
   - **New format** (two platforms): the `## Caption` section contains `**TikTok**` and `**Instagram**` sub-labels. Extract each platform's caption and hashtags separately.
   - **Legacy format** (single caption): the `## Caption` section has no sub-labels. Extract as a single block.

4. For **new format**, write two files:

   `output/caption/$ARGUMENTS_tiktok.txt`:
   ```
   <TikTok caption text>

   <TikTok hashtags>
   ```

   `output/caption/$ARGUMENTS_instagram.txt`:
   ```
   <Instagram caption text>

   <Instagram hashtags>
   ```

   For **legacy format**, write one file:

   `output/caption/$ARGUMENTS.txt`:
   ```
   <caption text>

   <hashtags>
   ```

5. Create `output/caption/` if it does not exist.

6. Confirm the file path(s) and show the first line of each caption as a preview.
