package stores

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func NewSqliteStore(addr string) *SqliteStore {
	return &SqliteStore{
		DB: sqlx.MustOpen("libsql", addr),
	}
}

type SqliteStore struct {
	*sqlx.DB
}
