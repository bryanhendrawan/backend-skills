# Invoice ratusan, tim cuma 2 orang
<!-- Slide 1 -->
<!-- cover-terminal: ~/finance $ process-invoices —parse —validate —report -->

**Theme**: Bagaimana AI mengubah proses invoice manual yang melelahkan jadi workflow yang mudah.

## Masalah Lama
<!-- Slide 2 -->

Puluhan email invoice masuk setiap hari, format beda-beda, semua harus diinput manual ke sistem.

<!-- process-flow: Proses manual sekarang -->
[STEP_01] Email invoice masuk
[STEP_02] Baca dan identifikasi field
[STEP_03] Input manual ke sistem ERP
[STEP_04] Cek tanggal jatuh tempo
[STEP_05] Buat laporan pembayaran
<!-- /process-flow -->

## Dengan AI
<!-- Slide 3 -->

AI parse setiap invoice, map otomatis ke field sistem. Tim hanya validasi dan konfirmasi.

<!-- process-flow: Proses dengan AI -->
[STEP_01] Email invoice masuk
[STEP_02] AI parse dan map semua field
[STEP_03] Tim validasi dalam hitungan menit
[STEP_04] AI generate laporan deadline otomatis
<!-- /process-flow -->

## The Diff
<!-- Slide 4 -->

Bukan cuma soal kecepatan. Risiko error dan deadline terlewat jauh berkurang.

<!-- comparison-window: Yang berubah -->
[TANPA AI] format beda tiap vendor, input manual sering error-prone.
↳ invoice deadline terlewat, relasi vendor rusak.
[DENGAN AI] AI yang mapping, kita cukup ambil keputusan.
↳ tim fokus validasi, bukan data entry sepanjang hari.
<!-- /comparison-window -->

## The Receipt
<!-- Slide 5 -->

Estimasi dampak nyata untuk tim finance dengan 50 invoice per hari.

<!-- data-table: Impact per hari -->
| Metrik | Tanpa AI | Dengan AI |
|---|---|---|
| Waktu per invoice | 15 menit | 3 menit |
| Error rate input | ~8% | ~1% |
| Laporan deadline | manual | otomatis |
<!-- /data-table -->

## Close
<!-- Slide 6 -->

**Kalau satu orang di company udah bisa kayak gini. Bayangkan jika semua tim bisa melakukannya**

Cari satu proses di tim kamu yang paling banyak copy-paste manual. Kita mulai dari sana.

---

## Caption

**TikTok**
50 invoice sehari, 2 orang tim finance, format beda tiap vendor.
AI yang baca, map, dan siapkan laporannya.
Tim tinggal validasi.
Follow kalau tim kamu masih input manual.

**Instagram**
Tim finance 2 orang, 50 invoice sehari, format beda-beda tiap vendor. Ini bukan masalah kecepatan, ini masalah workflow yang salah desain.

Dengan AI, 15 menit per invoice jadi 3 menit. Error rate turun dari 8% ke 1%. Laporan deadline pembayaran muncul otomatis.

AI tidak gantikan orangnya, AI yang handle bagian repetitif. Kita tinggal fokus ke keputusan yang butuh konteks.

Simpan post ini kalau ada proses di tim kamu yang masih jalan manual. Follow untuk lanjutannya.

## Hashtags

**TikTok**
#fyp #FYPIndonesia #aitools #aiforcoding #backendengineer #devid #techindonesia #workflowautomation

**Instagram**
#programmer #softwareengineer #backend #ai #coding #backendengineer #aitools #aiforcoding #devtools #productivity #devid #programmerid #techindonesia #aiindonesia #workflowautomation

---

## Design Prompt

---

**Claude Design Prompt — Invoice ratusan, tim cuma 2 orang**

Project folder name: `2026-05-28_ai-invoice-processing`
Save all slides for this topic inside that folder. Each new topic gets its own folder.

Platform: Instagram (post/reels) + TikTok (image/video)
Aspect ratio: 4:5 portrait — 1080 × 1350 px per slide
Total slides: 6

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
- Left: slide counter `01 / 06` — monospace, 16–18px, soft lavender (#A5A3C2)
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
- **CTA / Try It Today** (Problem-Solution template): stronger cyan ambient glow; bold H2 action statement left-aligned; CTA text below — soft lavender, 16–18px, varied per day (not boilerplate). No chips or icons.
- **Close slide** (Showdown template): two stacked zones separated by a short cyan horizontal rule. Top zone: section label "WHAT'S NEXT" + tease line in soft lavender italic, 18–22px. Bottom zone: section label "TRY IT TODAY" + bold white H2 action + CTA text in soft lavender, 16–18px. Stronger cyan ambient glow.

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
- Header: path/context string in soft lavender monospace
- Content rows (monospace, 16–18px): tag labels in amber/orange, backtick terms in cyan, ↳ continuation in muted lavender, $ verdict in muted lavender + bold white value

### Comparison Window
A two-column floating panel for before/after contrasts.
- Panel: same background and chrome as Terminal Window Block
- Header: comparison title in soft lavender monospace
- Column divider: thin vertical line (#1E1B3A)
- Row content: label in column-header color, explanation below in soft lavender, ~16px monospace

### Data Table
A minimal styled table for structured comparisons — 2–4 rows, 2–3 columns max.
- Panel: background #141128, rounded corners, border (#1E1B3A)
- Header row: cyan (#00D4FF) text, bold, ~14px uppercase, bottom border in #1E1B3A
- Data rows: white text ~16px; alternate rows use #141128 and #1A1735 for subtle banding
- Cell padding: generous — never cramped

### Process Flow
A horizontal flow diagram for visualizing sequential steps — workflows, pipelines, company processes.
- Panel: background #141128, rounded corners (8px), border (#1E1B3A)
- Step nodes: #1A1735 background, 8px radius, thin border (#1E1B3A)
- Step label `[STEP_XX]`: small cyan monospace, above or inside the node
- Step text: white, 14–16px, max 3–4 words per node
- Arrows between nodes: thin cyan (#00D4FF) right-pointing arrow, vertically centered between nodes
- Flow title: soft lavender monospace, above the panel
- Max 5 nodes — more becomes unreadable at 1080px canvas width

### Highlight Stat
A large floating stat callout — no panel, the number dominates the slide.
- `[NUMBER]`: 90–110px, bold, cyan (#00D4FF) — the dominant visual element
- `[LABEL]`: 24–28px, bold white, directly below the number
- `[CONTEXT]`: 16–18px, soft lavender, below the label
- Optional: faint cyan glow halo around the number for depth

## Tone
- Bold and modern — premium AI product feel
- Terminal/window elements are used selectively — they signal "engineer's context", not decoration
- High contrast, vibrant but not noisy — one strong color moment per slide
- Ambient glows give depth; keep type clean and uncluttered on top

---

**SLIDES:**

Note: every slide uses the global frame — counter top-left, series label top-right, @vilamas.ai footer bottom-center. All content is left-aligned and the content block is vertically centered on the canvas. Full spec in the Slide Structure section above.

**Slide 1 — Cover**
Frame: 01 / 06 top-left (monospace, soft lavender) | • Backend x AI top-right (soft lavender) | @vilamas.ai footer bottom-center (soft lavender)
Content block vertically centered between top bar and footer.
- Cover terminal prompt (left-aligned above title, monospace): ~/finance $ process-invoices —parse —validate —report
  Style: path in soft lavender, $ in cyan, command flags in white
- Heading: Invoice ratusan, tim cuma 2 orang — H1, 72–80px bold white, left-aligned; "2 orang" in cyan italic
- Subheading: Bagaimana AI mengubah proses invoice manual yang melelahkan jadi workflow yang mudah. — 18–22px soft lavender, below heading
- Navigation hint: geser buat baca →→→ — small soft lavender, bottom-left above footer

**Slide 2 — Masalah Lama**
Frame: 02 / 06 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — MASALAH LAMA (13–14px uppercase, letter-spaced, cyan, short cyan rule to the left)
- Headline: Puluhan email invoice masuk setiap hari, format beda-beda, semua harus diinput manual ke sistem. — H2, 44–52px bold white, left-aligned
- Process flow panel (left-aligned below headline):
  - Panel: #141128, rounded 8px, border #1E1B3A
  - Title: Proses manual sekarang — soft lavender monospace, above the panel
  - 5 pill nodes, left-to-right, connected by thin cyan (#00D4FF) right-pointing arrows:
    - Node 1: label [STEP_01] in small cyan monospace; text "Email invoice masuk" in white 14–16px; #1A1735 background, thin border #1E1B3A
    - Node 2: label [STEP_02]; text "Baca dan identifikasi field"
    - Node 3: label [STEP_03]; text "Input manual ke sistem ERP"
    - Node 4: label [STEP_04]; text "Cek tanggal jatuh tempo"
    - Node 5: label [STEP_05]; text "Buat laporan pembayaran"

**Slide 3 — Dengan AI**
Frame: 03 / 06 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — DENGAN AI (uppercase cyan, short cyan rule)
- Headline: AI parse setiap invoice, map otomatis ke field sistem. Tim hanya validasi dan konfirmasi. — H2, 44–52px bold white, left-aligned
- Process flow panel (left-aligned below headline):
  - Panel: #141128, rounded 8px, border #1E1B3A
  - Title: Proses dengan AI — soft lavender monospace, above the panel
  - 4 pill nodes, left-to-right, connected by thin cyan (#00D4FF) right-pointing arrows:
    - Node 1: label [STEP_01]; text "Email invoice masuk"
    - Node 2: label [STEP_02]; text "AI parse dan map semua field"
    - Node 3: label [STEP_03]; text "Tim validasi dalam hitungan menit"
    - Node 4: label [STEP_04]; text "AI generate laporan deadline otomatis"

**Slide 4 — The Diff**
Frame: 04 / 06 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — THE DIFF (uppercase cyan, short cyan rule)
- Headline: Bukan cuma soal kecepatan. Risiko error dan deadline terlewat jauh berkurang. — H2, 44–52px bold white, left-aligned
- Comparison window (left-aligned below headline):
  - Panel: #141128, rounded 8px, border #1E1B3A, chrome dots top-left (amber #E8955A, red #E85A5A, green #5AE87B)
  - Header: Yang berubah — soft lavender monospace
  - Left column header: [TANPA AI] — amber/orange
  - Right column header: [DENGAN AI] — cyan (#00D4FF)
  - Column divider: thin vertical line #1E1B3A
  - Row 1 left: "format beda tiap vendor, input manual sering error-prone." — amber monospace
  - Row 1 left ↳: "invoice deadline terlewat, relasi vendor rusak." — soft lavender, indented
  - Row 1 right: "AI yang mapping, kita cukup ambil keputusan." — cyan monospace
  - Row 1 right ↳: "tim fokus validasi, bukan data entry sepanjang hari." — soft lavender, indented

**Slide 5 — The Receipt**
Frame: 05 / 06 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — THE RECEIPT (uppercase cyan, short cyan rule)
- Headline: Estimasi dampak nyata untuk tim finance dengan 50 invoice per hari. — H2, 44–52px bold white, left-aligned
- Data table (left-aligned below headline):
  - Panel: #141128, rounded 8px, border #1E1B3A
  - Title: Impact per hari — soft lavender monospace, above the table
  - Header row: METRIK | TANPA AI | DENGAN AI — cyan bold uppercase ~14px, bottom border #1E1B3A
  - Row 1: "Waktu per invoice" | "15 menit" | "3 menit" — white ~16px, #141128 background
  - Row 2: "Error rate input" | "~8%" | "~1%" — white ~16px, #1A1735 background
  - Row 3: "Laporan deadline" | "manual" | "otomatis" — white ~16px, #141128 background
  - Generous cell padding; no outer border radius overlap with panel

**Slide 6 — Close**
Frame: 06 / 06 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — WHAT'S NEXT (uppercase cyan, short cyan rule)
- Tease: Kalau satu orang di company udah bisa kayak gini. Bayangkan jika semua tim bisa melakukannya — 18–22px soft lavender italic, left-aligned
- Short cyan horizontal rule as visual divider between tease and action
- Section label: — TRY IT TODAY (uppercase cyan, short cyan rule)
- Action: Cari satu proses di tim kamu yang paling banyak copy-paste manual. Kita mulai dari sana. — H2, 44–52px bold white, left-aligned
- CTA text (below action, left-aligned): "Jangan lupa like, save, share, and follow buat tips berikutnya" — 16–18px soft lavender, no chips or icons
- Stronger cyan ambient glow on this slide vs. others

---
