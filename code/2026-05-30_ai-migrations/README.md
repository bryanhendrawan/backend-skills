# Day 4 — AI-generated migrations

Extends Day 3 with a SQLite store and an embedded migration runner.

## Run

```
go mod download
go run .
```

Then:

```
curl localhost:8080/health
curl localhost:8080/migrations
```

The first boot applies `001_create_users.up.sql`. Subsequent boots see it in `schema_migrations` and skip it.

## How the migrations were generated

The slides demonstrate the workflow: describe the schema change in plain language to Claude Code, review the `.up.sql` and `.down.sql` it produces, and only commit after running the rollback locally.

Example prompt:

> Tambahin tabel `users` ke service ini. Field: id auto-increment, name, email unique, created_at default now. Bikin index buat lookup by email. Tulis up dan down migration di `migrations/`.

The two files in `migrations/` are the result, lightly reviewed.

## Resetting

Delete `day4.db` to start fresh. The migration runner will re-apply from version 1.
