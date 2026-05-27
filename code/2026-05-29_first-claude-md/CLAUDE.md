# CLAUDE.md

## Project Purpose

A minimal HTTP service for the backend-skills content library. Demonstrates a starter Go service so the CLAUDE.md, scaffolding, and conventions are visible end-to-end. This same service is extended on Day 4 (migrations) and Day 6 (CRUD).

## Architecture

- Single binary, `net/http` standard library only.
- One handler so far: `GET /health` returns `{"status":"ok","service":"day3"}`.
- No external dependencies, no database, no framework.

## Commands

- `go run .` — run the service locally on `:8080`.
- `go test ./...` — run tests.
- `go vet ./...` — static analysis.
- `curl localhost:8080/health` — smoke test.

## Conventions

- Errors bubble up via `error` return values, no panics in handlers.
- Handler functions take `(w http.ResponseWriter, r *http.Request)` only, no global state.
- One file per resource. When a resource grows beyond 200 lines, split by concern.
- Logs use the standard `log` package. No structured logging until we actually need it.
- JSON responses set `Content-Type: application/json` before encoding.
