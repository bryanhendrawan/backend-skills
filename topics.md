# Topics

Master index of every topic — planned and published. Always check this file for uniqueness (by angle, not just title) before proposing new topics.

**Columns**
- **Date** — intended publish date; also drives the `days/YYYY-MM-DD_<slug>.md` filename.
- **Slug** — kebab-case identifier used in filenames.
- **Title** — the display title shown in the day file.
- **Code** — `yes` if a `code/YYYY-MM-DD_<slug>/` sample exists or is planned.
- **Stack** — language for the code sample (`go`, `py`, `—`). Connected series share a stack.
- **Tool** — primary AI tool featured (`cc` = Claude Code CLI, `web` = Claude.ai / ChatGPT web chat, `both`).
- **Series** — short tag if the day belongs to a multi-day connected sequence.
- **Status** — content pipeline state: `draft` (no day file yet) / `written` (slides + caption + design prompt done, not yet posted) / `posted` (live on social media).
- **Post URL** — URL to the live post (Instagram, TikTok, or other platform). Fill in when status becomes `posted`; leave `—` otherwise.

---

## Batch 1: Days 1–31 (2026-05-27 → 2026-06-26)

| Date       | Slug                          | Title                                                                  | Code | Stack | Tool | Series        | Status  | Post URL |
|------------|-------------------------------|------------------------------------------------------------------------|------|-------|------|---------------|---------|----------|
| 2026-05-27 | why-delegate-to-ai            | Why backend engineers should learn to delegate to AI                   | no   | —     | both | —             | posted  | [IG](https://www.instagram.com/p/DY2LTxbgaw1/?utm_source=ig_web_copy_link&igsh=MzRlODBiNWFlZA==) [TT](https://www.tiktok.com/@vilamas.ai/photo/7644591137868778773?is_from_webapp=1&sender_device=pc&web_id=7635681994010330645) [FB](https://web.facebook.com/share/p/1J53brjXvn/) |
| 2026-05-28 | ai-invoice-processing         | Invoice ratusan, tim cuma 2 orang: gimana AI bantu                     | no   | —     | both | —             | draft   | —        |
| 2026-05-29 | first-claude-md               | Your first CLAUDE.md for a backend repo                                | yes  | go    | cc   | go-service    | written | —        |
| 2026-05-30 | ai-migrations                 | Generating database migrations from a natural-language schema change   | yes  | go    | cc   | go-service    | written | —        |
| 2026-05-31 | pre-pr-self-review            | Pre-PR self-review with /code-review: linter, security, pentest, logic | no   | —     | cc   | —             | written | —        |
| 2026-06-01 | scaffold-crud                 | Scaffolding CRUD endpoints by describing the resource                  | yes  | go    | cc   | go-service    | written | —        |
| 2026-06-02 | connect-db-via-mcp            | Connecting your dev database to Claude Code via MCP                    | yes  | go    | cc   | —             | written | —        |
| 2026-06-03 | realistic-fixtures            | Generating realistic test fixtures (and the edge cases you'd miss)     | yes  | go    | cc   | go-service    | draft   | —        |
| 2026-06-04 | explore-unfamiliar-codebase   | Exploring an unfamiliar codebase with web chat before opening the IDE  | no   | —     | web  | —             | draft   | —        |
| 2026-06-05 | table-driven-tests            | Table-driven tests at scale with AI                                    | yes  | go    | cc   | go-service    | draft   | —        |
| 2026-06-06 | decode-stack-trace            | Decoding a 500-line stack trace with AI                                | no   | —     | web  | —             | draft   | —        |
| 2026-06-07 | explain-your-own-code         | Decoding a service your team inherited when the author quit            | no   | —     | both | —             | draft   | —        |
| 2026-06-08 | triage-prod-logs              | Triaging production logs with AI: noise into signal                    | yes  | py    | both | —             | draft   | —        |
| 2026-06-09 | reproduce-flaky-test          | Reproducing a flaky test with AI's help                                | no   | —     | web  | —             | draft   | —        |
| 2026-06-10 | complex-sql-queries           | Writing the SQL you're embarrassed to write by hand                    | yes  | py    | web  | —             | draft   | —        |
| 2026-06-11 | ai-architecture-rubber-duck   | Using AI as a rubber duck for architecture decisions                   | no   | —     | web  | —             | draft   | —        |
| 2026-06-12 | api-contract-first            | Drafting an API contract before writing a line of code                 | yes  | py    | both | —             | draft   | —        |
| 2026-06-13 | capacity-estimation           | Capacity estimation with AI: "how many RPS can this handle?"           | no   | —     | web  | —             | draft   | —        |
| 2026-06-14 | spot-n-plus-one               | Spotting N+1 queries with AI before prod does                          | yes  | py    | cc   | —             | draft   | —        |
| 2026-06-15 | production-dockerfile         | Writing a production-ready Dockerfile with AI                          | yes  | go    | cc   | go-deploy     | draft   | —        |
| 2026-06-16 | k8s-from-dockerfile           | Generating Kubernetes manifests from a Dockerfile                      | yes  | go    | cc   | go-deploy     | draft   | —        |
| 2026-06-17 | github-actions-backend        | GitHub Actions for a backend service in ten minutes                    | yes  | go    | cc   | go-deploy     | draft   | —        |
| 2026-06-18 | better-pr-descriptions        | Generating PR descriptions that don't suck                             | no   | —     | cc   | —             | draft   | —        |
| 2026-06-19 | rfc-with-ai                   | Writing an RFC with AI as a co-author                                  | no   | —     | web  | —             | draft   | —        |
| 2026-06-20 | api-docs-from-spec            | Generating API docs from OpenAPI specs and code                        | yes  | py    | cc   | —             | draft   | —        |
| 2026-06-21 | first-custom-skill            | Your first custom Claude Code skill: a project-specific workflow       | yes  | go    | cc   | cc-customize  | draft   | —        |
| 2026-06-22 | custom-code-review-subagent   | Building a subagent for code review tailored to your codebase          | yes  | go    | cc   | cc-customize  | draft   | —        |
| 2026-06-23 | hooks-for-formatting          | Hooks: automating formatting and linting after every AI edit           | yes  | go    | cc   | cc-customize  | draft   | —        |
| 2026-06-24 | security-review-injection     | Security review with AI: catching injection risks in PRs               | yes  | py    | cc   | —             | draft   | —        |
| 2026-06-25 | ai-failure-modes              | When NOT to trust AI: the failure modes every backend engineer needs   | no   | —     | both | —             | draft   | —        |
| 2026-06-26 | workflow-retrospective        | Your AI-augmented daily workflow: what to keep, drop, double down on   | no   | —     | both | —             | draft   | —        |

---

## Series notes

- **`go-service`** (Days 3, 4, 6, 8, 10): one Go HTTP service built across five non-consecutive days. Each entry adds to the same repo — CLAUDE.md, migrations, CRUD handlers, fixtures, tests. Days 4–10 in this series are interleaved with non-Go topics so a daily reader gets variety, but the codebase is cumulative: Day 10's code depends on Day 8's, which depends on Day 6's, etc.
- **`go-deploy`** (Days 20–22): containerize and deploy the `go-service`. Dockerfile → K8s manifests → CI/CD. Consecutive because the build artifacts chain tightly.
- **`cc-customize`** (Days 26–28): customize Claude Code for the `go-service` project — skill, subagent, hooks. Consecutive for the same reason.

The three Go series share the same toy service so a reader following along builds one real thing across the month.

## Workflow note: the `/code-review` skill (Day 5)

Day 5 introduces `/code-review` early on purpose. Once published, we apply that same workflow to our own work: every generated `code/` sample is reviewed via `/code-review` before it lands. The skill covers linter issues, security/injection risks, pentest-style scans, and logic correctness — the same checks we want backend readers to apply to AI-generated code in their own projects.

## Coverage breakdown (Batch 1)

| Bucket                          | Days                                  |
|---------------------------------|---------------------------------------|
| Mindset / decision frameworks   | 1, 2, 30, 31                          |
| Daily coding (Go service)       | 3, 4, 6, 8, 10                        |
| Code review (reusable workflow) | 5                                     |
| MCP integration                 | 7                                     |
| Understanding existing code     | 9, 11, 12                             |
| Debugging & ops                 | 13, 14, 15                            |
| Architecture & design           | 16, 17, 18                            |
| Performance review              | 19                                    |
| Deployment (Go)                 | 20, 21, 22                            |
| Documentation & communication   | 23, 24, 25                            |
| Claude Code customization       | 26, 27, 28                            |
| Security                        | 29                                    |
