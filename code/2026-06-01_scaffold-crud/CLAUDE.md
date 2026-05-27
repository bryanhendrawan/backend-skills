# CLAUDE.md

## Project Purpose

Extends Day 4 with full CRUD endpoints for the `users` resource. Demonstrates the AI scaffolding workflow taught in Day 6: describe the resource once, AI generates handler + repository + route registration + validation, then review.

## Architecture

- Go 1.22 `net/http` ServeMux with method-and-path routing (`POST /users`, `GET /users/{id}`, etc).
- `database/sql` + pure-Go `modernc.org/sqlite` driver.
- Migrations embedded from `migrations/` directory (carried over from Day 4).
- Repository layer (`UserRepo`) sits between handlers and `database/sql`. Handlers don't see SQL.
- Validation lives next to the payload type, runs before the repo call.

## Commands

- `go run .` — apply migrations and start the server on `:8080`.
- `curl localhost:8080/health` — service liveness.
- `curl localhost:8080/users` — list users.
- `curl -X POST localhost:8080/users -d '{"name":"Bryan","email":"bryan@example.com"}'` — create.

## Conventions

- Handler responsibilities: parse + validate input, call repo, map errors to HTTP status, write response. No SQL in handlers.
- Repo responsibilities: SQL only. Returns `ErrNotFound` for missing rows, raw `error` for everything else.
- Validation lives in `users.go` next to the type, not in a separate validators package. Keep small repos small.
- `writeJSON` / `writeError` helpers centralize response shape so every endpoint matches.
- New resource: copy the `users.go` pattern (struct + repo + validation + 5 handlers) and register the routes in `main.go`.

## Scaffolding via AI

Describe the resource: name, fields, validation rules, business constraints. AI returns the full file in this pattern (struct, repo, validation, handlers). You review naming, error mapping, and edge cases (empty list, missing ID, duplicate email).

The bar: the code reads like the rest of this repo. Same helpers, same error patterns, same conventions.
