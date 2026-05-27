# Backend Engineer Agent Memory

Auto-maintained log of what's been covered, what's pending, and what Bryan has told us along the way. Read at session start. Update after every produced day.

## Topics covered (published)

| Date | Slug | Angle | Code |
|---|---|---|---|
| 2026-05-27 | why-delegate-to-ai | Mindset shift: treat AI like a junior dev (spec + review), not autocomplete or oracle | no |
| 2026-05-28 | claude-code-vs-web-chat | Tool choice: web chat for one-off questions, Claude Code for repo-aware multi-file work | no |
| 2026-05-29 | first-claude-md | Project setup: writing CLAUDE.md as a briefing for AI partner | yes (Go baseline) |
| 2026-05-30 | ai-migrations | Database: natural-language to up+down migration, verify rollback before merge | yes (Go + SQLite + embedded migrations) |
| 2026-05-31 | pre-pr-self-review | Code review: 4-pass /code-review (linter, security, pentest, logic) before push | no |
| 2026-06-01 | scaffold-crud | Productivity: scaffold handler + repo + validation from a resource description | yes (Go CRUD on top of Day 4) |
| 2026-06-02 | connect-db-via-mcp | MCP: connect dev DB to Claude Code so AI knows schema + index + relations | yes (Postgres + Docker + Go seed + MCP config) |

## Topics planned (in topics.md, not yet produced)

Days 8+ are listed in `topics.md`. Re-read that file at the start of any planning session — the index is canonical.

## Domain coverage

- **AI workflow mindset**: 1 (Day 1).
- **Tool choice**: 1 (Day 2).
- **Project setup**: 1 (Day 3).
- **Database**: 2 (Day 4 migrations, Day 7 MCP).
- **Code review**: 1 (Day 5).
- **CRUD scaffolding**: 1 (Day 6).

Light or unrepresented: **observability, CI/CD, testing, performance, debugging, queues/caching, deployment, security beyond /code-review**.

When proposing Days 8+, bias toward these gaps. Run `/generate-topics 5` and bias it to non-Go variety (some Python) and to `web` or `both` tools (4 of 7 days so far are `cc`-only).

## Code languages used

- **Go**: Days 3, 4, 6 (one evolving `go-service`) and Day 7 (standalone Postgres + MCP).
- **Python**: not used yet. Candidate for variety in Days 8+.

## Series in progress

- **`go-service`** — Days 3, 4, 6.
  - Day 3: minimal `net/http` service, `GET /health`, `go.mod`, example `CLAUDE.md`.
  - Day 4: + SQLite (pure-Go `modernc.org/sqlite`), embedded `migrations/` (`go:embed`), `schema_migrations` tracking table, `users` table created via migration, `GET /migrations` endpoint.
  - Day 6: + `users.go` with full CRUD (`UserRepo`, validation, 5 handlers using Go 1.22 ServeMux method-and-path routing).
  - Each day's `code/<slug>/` is a self-contained copy that extends the previous. Diff between consecutive days reads as one coherent change.

## Decisions and feedback Bryan has confirmed

- **Tone**: kamu/aku, professional but casual. Not lo/gue (too streetwise), not formal anda. Confirmed Day 1.
- **No em-dashes** anywhere in slide or caption content. They read as AI-generated. Use a period or comma.
- **Visual identity locked**: midnight navy `#0B0A1F`, cyan `#00D4FF` accent, soft lavender `#A5A3C2` muted text, ambient glows on every slide, `@vilamas.ai` footer. Source of truth: `.claude/skills/design-prompt/design-guidelines.md`. Never revert.
- **Series continuity over fresh samples**: when topics share a series tag, the code evolves rather than restarting. Confirmed during Days 2–7 planning.
- **/code-review is built-in / external**, not a skill we maintain in this repo. Day 5 just teaches how to use it.
- **Plan-review is mandatory** — never jump to generating content. Always outline → approval → write.
- **Day file = three zones**: slides → `---` → caption + hashtags → `---` → design prompt. Established Day 1, used Days 2–7.
- **SQLite pure-Go (`modernc.org/sqlite`)** is the default DB for series code. Zero CGO, zero install friction.

## Open questions / things to revisit

- Day 8+ topic mix — once Days 2–7 are scheduled, run `/generate-topics 5` and rebalance toward observability, testing, debugging, or Python.
- Whether `/code-review` should eventually become a project skill, or stay external.
- The Day 6 CRUD sample maps duplicate-email errors to 500 instead of 409. Left as a deliberate review-finding example in the README; could become a follow-up day topic ("mapping DB errors to HTTP correctly").
- Day 4 + Day 6 embed an identical migration. Future days that extend `go-service` should add new migrations (002+), never edit existing.
