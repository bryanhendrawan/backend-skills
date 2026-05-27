# Claude Code atau web chat?
<!-- Slide 1 -->

**Theme**: Pilih alat AI berdasarkan bentuk masalah, bukan kebiasaan.

## The Problem
<!-- Slide 2 -->

Kebanyakan engineer default ke satu alat dan nggak pernah ganti. Tugas yang harusnya 5 menit jadi 1 jam karena salah pilih.

## How It Works

1. Web chat menang buat: nanya konsep, debug satu function, brainstorming arsitektur tanpa konteks repo.
<!-- Slide 3 -->

2. Claude Code menang buat: edit multi-file, refactor lintas modul, MCP ke DB, code review otomatis.
<!-- Slide 4 -->

3. Rule of thumb: kalau jawaban butuh lihat dua file di repo, pakai Claude Code.
<!-- Slide 5 -->

## Key Takeaway
<!-- Slide 6 -->

**Pilih alat yang tahu konteks kamu, bukan yang paling familiar.**

## Try It Today
<!-- Slide 7 -->

Tugas berikutnya, sengaja pakai alat yang beda dari kebiasaan. Catat hasilnya.

---

## Caption

Senior backend engineer nggak pakai satu AI tool buat semua. Mereka pilih berdasarkan bentuk masalah.

Web chat buat nanya cepat, brainstorm konsep, atau debug satu function. Claude Code buat edit multi-file, refactor besar, atau langsung ngajak AI lihat schema database lewat MCP.

Kalau jawaban kamu butuh lihat dua file atau lebih, kamu udah salah pilih alat.

Save kalau kamu masih default ke satu tool. Follow buat tips backend × AI harian.

## Hashtags

#programmer #softwareengineer #coding #backend #ai #claudecode #chatgpt #aitools #aiforcoding #backendengineer #devid #programmerid #techindonesia #ngoding #softwareengineerid

---

## Design Prompt

---

**Claude Design Prompt — Claude Code atau web chat?**

Project folder name: `2026-05-28_claude-code-vs-web-chat`
Save all slides for this topic inside that folder. Each new topic gets its own folder.

Platform: Instagram (post/reels) + TikTok (image/video)
Aspect ratio: 4:5 portrait — 1080 × 1350 px per slide
Total slides: 7

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
- Left: slide counter `01 / 07` — monospace, 16–18px, soft lavender (#A5A3C2)
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
- Cover: cover-terminal prompt above H1 title; H1 left-aligned, key word in cyan italic; theme subheading below in soft lavender; `geser buat baca →→→` bottom-left in small soft lavender
- Problem slide: section label + H2 headline + terminal window block below
- Step slides: filled cyan circle badge (48–56px diameter, solid cyan #00D4FF, white step number centered inside) to the left of the step text; body left-aligned right of badge; optional component block below
- Key Takeaway: H2 headline left-aligned; bold statement; short cyan horizontal rule below as accent
- CTA (Try It Today): stronger cyan ambient glow; bold H2 action statement left-aligned; CTA row of 4 chips below: Simpan | Suka | Bagikan | Follow @vilamas.ai

## Branding
- Series label: top-right every slide — see Slide Structure above
- No stock photos — icons and typography only
- Footer: bottom center every slide — see Slide Structure above

## Slide Component Types

### Cover Terminal Prompt
- Format: `~/[context] $ [action-phrase-as-flags]`
- Styling: path in soft lavender, `$` in cyan, command text in white — all monospace, small
- Placement: left-aligned above H1 title
- Em-dashes as flag characters are intentional — preserve exactly

### Terminal Window Block
- Panel: #141128, rounded 8px, border #1E1B3A, left 2px cyan accent bar
- Chrome: three dots top-left — amber #E8955A, red #E85A5A, green #5AE87B
- Header: path/context in soft lavender monospace
- Rows: [ERR_XX] in amber, backtick terms in cyan, ↳ in cyan + muted lavender, $ verdict muted lavender + bold white value

### Comparison Window
- Panel: same as Terminal Window Block
- Header: title in soft lavender monospace
- Left col [BAD] in amber; right col [GOOD] in cyan; thin column divider #1E1B3A
- Row labels in column-header color; ↳ explanations in soft lavender monospace

### Data Table
- Panel: #141128, rounded 8px, border #1E1B3A
- Header row: cyan bold uppercase ~14px, bottom border #1E1B3A
- Data rows: white ~16px; alternate row banding #141128 / #1A1735; generous cell padding

## Tone
- Bold and modern — premium AI product feel
- Terminal/window elements are used selectively — they signal "engineer's context", not decoration
- High contrast, vibrant but not noisy — one strong color moment per slide
- Ambient glows give depth; keep type clean and uncluttered on top

---

**SLIDES:**

Note: every slide uses the global frame — counter top-left, series label top-right, @vilamas.ai footer bottom-center. All content is left-aligned. Full spec in the Slide Structure section above.

**Slide 1 — Cover**
Frame: 01 / 07 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Heading: Claude Code atau web chat? — H1, 72–80px bold white, left-aligned
- Subheading: Pilih alat AI berdasarkan bentuk masalah, bukan kebiasaan. — 18–22px soft lavender
- Navigation hint: geser buat baca →→→ — small soft lavender, bottom-left above footer

**Slide 2 — The Problem**
Frame: 02 / 07 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — THE PROBLEM (13–14px uppercase, letter-spaced, cyan, short cyan rule to the left)
- Headline: Kebanyakan engineer default ke satu alat dan nggak pernah ganti. — H2, 44–52px bold white, left-aligned
- Body: Tugas yang harusnya 5 menit jadi 1 jam karena salah pilih. — 24–28px white

**Slide 3 — Step 1**
Frame: 03 / 07 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — HOW IT WORKS (uppercase cyan, short cyan rule)
- Step badge: filled cyan circle 48–56px, solid #00D4FF, white "1" centered inside; sits to the left of the body text
- Body: Web chat menang buat: nanya konsep, debug satu function, brainstorming arsitektur tanpa konteks repo. — 24–28px white, left-aligned to the right of the badge

**Slide 4 — Step 2**
Frame: 04 / 07 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Step badge: filled cyan circle 48–56px, solid #00D4FF, white "2" centered inside; sits to the left of the body text
- Body: Claude Code menang buat: edit multi-file, refactor lintas modul, MCP ke DB, code review otomatis. — 24–28px white, left-aligned to the right of the badge

**Slide 5 — Step 3**
Frame: 05 / 07 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Step badge: filled cyan circle 48–56px, solid #00D4FF, white "3" centered inside; sits to the left of the body text
- Body: Rule of thumb: kalau jawaban butuh lihat dua file di repo, pakai Claude Code. — 24–28px white, left-aligned to the right of the badge

**Slide 6 — Key Takeaway**
Frame: 06 / 07 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — KEY TAKEAWAY (uppercase cyan, short cyan rule)
- Statement: Pilih alat yang tahu konteks kamu, bukan yang paling familiar. — H2, 44–52px bold white, left-aligned
- Short cyan horizontal rule below the statement as decorative accent

**Slide 7 — Try It Today**
Frame: 07 / 07 top-left | • Backend x AI top-right | @vilamas.ai footer
Content block vertically centered between top bar and footer.
- Section label: — TRY IT TODAY (uppercase cyan, short cyan rule)
- Action: Tugas berikutnya, sengaja pakai alat yang beda dari kebiasaan. Catat hasilnya. — H2, 44–52px bold white, left-aligned
- CTA row (below the action text, left-aligned, equal spacing between chips):
  - 🔖 Simpan — icon ~24px + "Simpan" label ~12px below, soft lavender, subtle #141128 chip rounded 8px
  - ❤️ Suka — icon ~24px + "Suka" label ~12px below, soft lavender, subtle #141128 chip rounded 8px
  - 📤 Bagikan — icon ~24px + "Bagikan" label ~12px below, soft lavender, subtle #141128 chip rounded 8px
  - ➕ Follow @vilamas.ai — icon ~24px + "Follow @vilamas.ai" label ~12px below, cyan (#00D4FF) to make it pop, subtle #141128 chip rounded 8px
- Stronger cyan ambient glow on this slide vs. others

---
