package main

import (
	"github.com/Linkinlog/loggr/internal/env"
	"github.com/Linkinlog/loggr/internal/stores"
)

func main() {
	e := env.NewEnv()

	sqliteAddr := e.GetOrDefault("SQLITE_ADDR", "loggr.db")

	s := stores.NewSqliteStore(sqliteAddr)

	schema := `CREATE TABLE IF NOT EXISTS gardens (
		id TEXT PRIMARY KEY NOT NULL,
		user_id TEXT NOT NULL,
		name TEXT NOT NULL,
		location TEXT NOT NULL,
		description TEXT NOT NULL,
		image TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS items (
		id TEXT PRIMARY KEY NOT NULL,
		garden_id TEXT NOT NULL,
		name TEXT NOT NULL,
		image TEXT NOT NULL,
		type INTEGER NOT NULL,
		field1 TEXT NOT NULL,
		field2 TEXT NOT NULL,
		field3 TEXT NOT NULL,
		field4 TEXT NOT NULL,
		field5 TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS sessions (
		id TEXT PRIMARY KEY NOT NULL,
		user_id TEXT NOT NULL
	);
	`

	s.MustExec(schema)
}
