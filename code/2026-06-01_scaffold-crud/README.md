# Day 6 — Scaffolding CRUD by describing the resource

Extends Day 4 with full CRUD endpoints for `users`.

## Run

```
go mod download
go run .
```

Then try the endpoints:

```
curl localhost:8080/health
curl -X POST localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Bryan","email":"bryan@example.com"}'
curl localhost:8080/users
curl localhost:8080/users/1
curl -X PUT localhost:8080/users/1 \
  -H 'Content-Type: application/json' \
  -d '{"name":"Bryan H","email":"bryan@vilamas.ai"}'
curl -X DELETE localhost:8080/users/1
```

## How the CRUD was scaffolded

The slides show the workflow: describe the `users` resource (fields, validation, behavior) to Claude Code, and it produces `users.go` in one shot — struct + repo + validation + 5 handlers — matching the conventions in `CLAUDE.md`.

Example prompt:

> Tambahin resource `users` ke service ini. Field: id, name (required), email (required, valid format, unique), created_at (auto). Bikin handler POST/GET-list/GET-one/PUT/DELETE pakai ServeMux Go 1.22, plus repository layer pakai database/sql. Validation di package yang sama. Ikutin pattern `writeJSON` / `writeError` di main.go.

You review: error mapping (404 vs 500 vs 400), naming consistency, missing edge cases (e.g. duplicate email returns 500 in this sample — a follow-up improvement to map to 409).

## Resetting

Delete `day6.db` to start fresh. Migrations re-apply from version 1.
