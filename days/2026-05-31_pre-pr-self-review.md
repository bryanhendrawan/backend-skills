# Review PR kamu, sebelum reviewer kamu.
<!-- Slide 1 -->

**Theme**: 4 pass /code-review sebelum push: linter, security, pentest, logic.

## The Problem
<!-- Slide 2 -->

Reviewer nemuin typo, null check yang lupa, atau SQL injection di feature baru. Jam-jam yang terbuang buat hal yang harusnya ketangkep duluan.

## How It Works

1. Pass linter: format, naming, dead code. AI tahu konvensi tim kalau ada di CLAUDE.md.
<!-- Slide 3 -->

2. Pass security: secrets di kode, hardcoded keys, auth check yang lupa di endpoint.
<!-- Slide 4 -->

3. Pass pentest mindset: input validation, rate limit, SSRF, IDOR di endpoint baru.
<!-- Slide 5 -->

4. Pass logic: edge case, race condition, error path yang nggak ke-handle.
<!-- Slide 6 -->

## Key Takeaway
<!-- Slide 7 -->

**Reviewer suka PR yang lewat 4 pass duluan.**

## Try It Today
<!-- Slide 8 -->

Branch terakhir kamu, run /code-review sebelum push.

---

## Caption

Reviewer kamu nggak suka nemuin typo atau SQL injection di feature baru. Itu jam-jam yang harusnya kamu nangkep duluan.

Pre-PR self-review = 4 pass: linter (format, naming), security (secrets, auth), pentest mindset (input validation, SSRF, IDOR), dan logic (edge case, race condition). Semua sebelum reviewer kamu buka tab.

Bukan soal nggak percaya tim. Soal hormatin waktu mereka.

Save kalau kamu mau PR yang reviewer kamu suka buka. Follow buat tips backend × AI harian.

## Hashtags

#programmer #softwareengineer #coding #backend #ai #claudecode #aitools #aiforcoding #backendengineer #codereview #devsecops #devid #programmerid #techindonesia #softwareengineerid

---

## Design Prompt

---

**Claude Design Prompt — Review PR kamu, sebelum reviewer kamu.**

Project folder name: `2026-05-31_pre-pr-self-review`
Save all slides for this topic inside that folder. Each new topic gets its own folder.

Platform: Instagram (post/reels) + TikTok (image/video)
Aspect ratio: 9:16 portrait — 1080 × 1920 px per slide
Total slides: 8

## Theme
- Background: deep midnight navy (#0B0A1F) — dark with a purple lean, not plain black
- Ambient glow: faint cyan radial glow from top-right, faint violet radial glow from bottom-left — apply on every slide for premium depth
- Grid backdrop: subtle grid overlay, slightly dialed down so gradients read clearly
- Primary text: pure white (#FFFFFF)
- Accent: cyan (#00D4FF) — step numbers, keywords, highlights, badges — pops against the navy base
- Muted text: soft lavender (#A5A3C2) — subtitles, labels, secondary info
- Body text: clean sans-serif (Inter or similar)

## Layout
- One idea per slide — generous whitespace, nothing crowded
- Cover slide: navy base with ambient glows prominent, large bold white title centered, cyan accent on key word(s), muted subtitle below
- Step slides: filled cyan circle badge with white step number left side, white body text right side
- Key Takeaway slide: cyan bold statement centered on navy, large — poster feel
- CTA slide (Try It Today): stronger cyan glow, white bold text, action-first

## Branding
- Series label: "Backend x AI" — small, top-right corner of every slide, soft lavender (#A5A3C2)
- No stock photos — icons and typography only
- Slide footer: account handle "@vilamas.ai" bottom center, small, soft lavender (#A5A3C2)

## Tone
- Bold and modern — premium AI product feel, not a terminal
- High contrast, vibrant but not noisy — one strong color moment per slide
- Ambient glows give depth; keep type clean and uncluttered on top

---

**SLIDES:**

**Slide 1 — Cover**
- Heading: Review PR kamu, sebelum reviewer kamu.
- Subheading: 4 pass /code-review sebelum push: linter, security, pentest, logic.

**Slide 2 — The Problem**
Reviewer nemuin typo, null check yang lupa, atau SQL injection di feature baru. Jam-jam yang terbuang buat hal yang harusnya ketangkep duluan.

**Slide 3 — Step 1 (Linter)**
Pass linter: format, naming, dead code. AI tahu konvensi tim kalau ada di CLAUDE.md.

**Slide 4 — Step 2 (Security)**
Pass security: secrets di kode, hardcoded keys, auth check yang lupa di endpoint.

**Slide 5 — Step 3 (Pentest)**
Pass pentest mindset: input validation, rate limit, SSRF, IDOR di endpoint baru.

**Slide 6 — Step 4 (Logic)**
Pass logic: edge case, race condition, error path yang nggak ke-handle.

**Slide 7 — Key Takeaway**
Reviewer suka PR yang lewat 4 pass duluan.

**Slide 8 — Try It Today**
Branch terakhir kamu, run /code-review sebelum push.

---
