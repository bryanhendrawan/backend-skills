# Migration database dari satu kalimat.
<!-- Slide 1 -->

**Theme**: Schema change yang aman = AI generate, kamu yang verify rollback.

## The Problem
<!-- Slide 2 -->

Nulis migration manual berarti lupa naming convention, lupa down migration, atau lupa index pendukung. Semua jadi technical debt.

## How It Works

1. Deskripsikan perubahan dalam bahasa biasa: "tambahin field email unique ke tabel users."
<!-- Slide 3 -->

2. AI generate up dan down migration sekaligus. Review file, jangan langsung apply.
<!-- Slide 4 -->

3. Run di branch baru, test rollback. Baru merge ke main.
<!-- Slide 5 -->

## Key Takeaway
<!-- Slide 6 -->

**Migration yang aman = AI generate, kamu yang verify rollback.**

## Try It Today
<!-- Slide 7 -->

Pilih schema change yang udah lama ditunda. Coba prompt AI hari ini.

---

## Caption

Migration database paling males ditulis manual. Padahal ini bagian yang AI paling jago, kalau kamu kasih intent yang jelas.

Deskripsikan perubahan schema dalam bahasa biasa: "tambahin field email unique ke tabel users." AI generate up dan down migration sekaligus. Kamu tinggal review naming, index, dan test rollback sebelum apply.

Yang bikin migration aman bukan AI-nya. Tapi kamu yang verify rollback di branch dev sebelum merge.

Save buat referensi waktu kamu nulis migration berikutnya. Follow buat tips backend × AI harian.

## Hashtags

#programmer #softwareengineer #coding #backend #ai #claudecode #aitools #aiforcoding #backendengineer #database #sqltips #golang #devid #programmerid #techindonesia

---

## Design Prompt

---

**Claude Design Prompt — Migration database dari satu kalimat.**

Project folder name: `2026-05-30_ai-migrations`
Save all slides for this topic inside that folder. Each new topic gets its own folder.

Platform: Instagram (post/reels) + TikTok (image/video)
Aspect ratio: 9:16 portrait — 1080 × 1920 px per slide
Total slides: 7

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
- Heading: Migration database dari satu kalimat.
- Subheading: Schema change yang aman = AI generate, kamu yang verify rollback.

**Slide 2 — The Problem**
Nulis migration manual berarti lupa naming convention, lupa down migration, atau lupa index pendukung. Semua jadi technical debt.

**Slide 3 — Step 1**
Deskripsikan perubahan dalam bahasa biasa: "tambahin field email unique ke tabel users."

**Slide 4 — Step 2**
AI generate up dan down migration sekaligus. Review file, jangan langsung apply.

**Slide 5 — Step 3**
Run di branch baru, test rollback. Baru merge ke main.

**Slide 6 — Key Takeaway**
Migration yang aman = AI generate, kamu yang verify rollback.

**Slide 7 — Try It Today**
Pilih schema change yang udah lama ditunda. Coba prompt AI hari ini.

---
