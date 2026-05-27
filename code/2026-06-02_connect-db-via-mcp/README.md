# Day 7 — Connect your dev database to Claude Code via MCP

Standalone demo. Spins up a Postgres in Docker, seeds it with users + orders, and wires Claude Code to query it via MCP.

## Prerequisites

- Docker + Docker Compose.
- Go 1.22+.
- Node.js (for `npx`, which runs the official MCP Postgres server).
- Claude Code.

## Step 1 — Start Postgres

```
docker compose up -d
```

Postgres is exposed on `localhost:5433` (5432 is often taken by a system Postgres).

## Step 2 — Seed the demo data

```
go mod download
go run .
```

This creates the `users` + `orders` schema and inserts 3 users and 6 orders.

## Step 3 — Register the MCP server with Claude Code

Copy the contents of `mcp-config.json` into your Claude Code MCP settings (the location depends on how you run Claude Code — see the Claude Code docs for the per-project `.claude/mcp.json` or user-level config path).

The config tells Claude Code to launch `@modelcontextprotocol/server-postgres` pointed at the demo DB. First run downloads the package via `npx`.

## Step 4 — Ask Claude Code

Open Claude Code in this directory and ask:

- *"Schema apa aja yang ada di DB ini?"*
- *"Berapa total revenue (sum of amount_cents) per user, urut dari yang terbesar?"*
- *"User mana yang pernah punya order dengan status refunded?"*

Claude Code uses the MCP server to introspect the schema and run read queries directly. No more pasting `\d users` into chat.

## Cleanup

```
docker compose down -v
```

Removes the container + volume.

## Notes on safety

The official MCP Postgres server connects in **read-only** mode by default. For dev DBs this is fine; never point an MCP server at a prod DB without explicit allow-listing and access review.
