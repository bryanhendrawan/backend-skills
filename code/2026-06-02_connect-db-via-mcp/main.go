package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

const schema = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    amount_cents INTEGER NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
`

type seedUser struct {
	Name   string
	Email  string
	Orders []seedOrder
}

type seedOrder struct {
	AmountCents int
	Status      string
}

var seedData = []seedUser{
	{"Bryan", "bryan@vilamas.ai", []seedOrder{
		{12000, "paid"},
		{45000, "pending"},
	}},
	{"Ari", "ari@example.com", []seedOrder{
		{8000, "paid"},
	}},
	{"Dita", "dita@example.com", []seedOrder{
		{30000, "refunded"},
		{15000, "paid"},
		{22000, "paid"},
	}},
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5433/demo"
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer conn.Close(ctx)

	if _, err := conn.Exec(ctx, schema); err != nil {
		log.Fatalf("schema: %v", err)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatalf("begin: %v", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, "TRUNCATE users, orders RESTART IDENTITY CASCADE"); err != nil {
		log.Fatalf("truncate: %v", err)
	}

	for _, u := range seedData {
		var userID int
		err := tx.QueryRow(ctx,
			"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
			u.Name, u.Email,
		).Scan(&userID)
		if err != nil {
			log.Fatalf("insert user %s: %v", u.Email, err)
		}

		for _, o := range u.Orders {
			if _, err := tx.Exec(ctx,
				"INSERT INTO orders (user_id, amount_cents, status) VALUES ($1, $2, $3)",
				userID, o.AmountCents, o.Status,
			); err != nil {
				log.Fatalf("insert order: %v", err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("commit: %v", err)
	}

	log.Println("seeded: 3 users, 6 orders. point your MCP config at this DB and try a query.")
}
