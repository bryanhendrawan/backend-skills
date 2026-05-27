# CLAUDE.md

## Project Purpose

Extends Day 3 with a SQLite database and an embedded migration runner. Demonstrates the natural-language-to-migration workflow taught in Day 4.

## Architecture

- `net/http` standard library for the HTTP server.
- `database/sql` + pure-Go `modernc.org/sqlite` driver (no CGO).
- Migrations live in `migrations/` as `NNN_name.up.sql` / `NNN_name.down.sql`.
- The runner embeds the `migrations/` directory with `go:embed`, tracks applied versions in `schema_migrations`, applies any unapplied `.up.sql` in version order on boot.

## Commands

- `go run .` — apply migrations and start the server on `:8080`.
- `go test ./...` — run tests.
- `curl localhost:8080/health` — service + DB liveness.
- `curl localhost:8080/migrations` — list applied migrations.

## Conventions

- Same as Day 3 (handlers parameterless except for `w`/`r`, errors bubble up, JSON sets Content-Type).
- New migration: copy the highest version + 1, write both `.up.sql` and `.down.sql`. Never edit an already-applied migration; always add a new one.
- Down migrations are not run automatically. They exist for manual rollback testing in a dev branch.

## Adding a migration via AI

Describe the change in plain language to your AI: *"tambahin field phone_number ke tabel users, nullable, dengan index buat lookup by phone."* AI returns both `.up.sql` and `.down.sql`. You review naming, index choice, and rollback safety before committing.
