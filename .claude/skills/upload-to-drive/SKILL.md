---
name: upload-to-drive
description: Upload per-slide PNGs and caption txt files for a given date to Google Drive under Content/{date}/. Requires output/images/{date}/ to exist (run /pdf-to-images first). Caption files from /export-caption are included automatically if present. First run opens a browser for OAuth consent; subsequent runs use the cached token in secrets/token.json.
argument-hint: <YYYY-MM-DD>
---

Upload slides and captions to Google Drive for the given date.

## Step 1 — Resolve the date slug

Extract `YYYY-MM-DD` from the argument. If the argument contains a full date (e.g. "2026-05-28" or "topic 2026-05-28"), parse it.

## Step 2 — Verify images exist

Check that `output/images/{date_slug}/` exists and contains at least one `slide-*.png`. If not, tell the user to run `/pdf-to-images {date_slug}` first and stop.

## Step 3 — Run the upload script

```
uv run --with google-auth-oauthlib --with google-api-python-client \
    python scripts/upload_to_drive.py {date_slug}
```

The script:
- On first run: opens a browser for Google OAuth consent; saves token to `secrets/token.json`
- On subsequent runs: uses the cached token (auto-refreshes if expired)
- Finds or creates `Content/{date_slug}/` folder in Google Drive
- Uploads each `slide-NN.png`; overwrites if file with same name already exists
- Also uploads `output/caption/{date_slug}*_tiktok.txt` and `*_instagram.txt` if they exist
- If caption files are missing, reports it and suggests running `/export-caption` first

## Step 4 — Report results

Report back:
- How many slides and caption files were uploaded
- The Google Drive folder path (`Content/{date_slug}/`)
- If captions were missing, remind the user to run `/export-caption {date_slug}` then re-run this skill
- Any errors, with a suggested fix
