package repositories_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/stores"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func setupSessionRepository() (*repositories.SessionRepository, *models.Session) {
	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	sr := repositories.NewSessionRepository(store)

	schema := `CREATE TABLE IF NOT EXISTS sessions (
		id TEXT PRIMARY KEY NOT NULL,
		user_id TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	);
	`

	store.MustExec(schema)

	u, _ := models.NewUser("Batman", "idk@aol.com", "password123")
	s := models.NewSession(u)
	store.MustExec("INSERT INTO sessions (id, user_id) VALUES (?, ?)", s.Id, s.User.Id)
	store.MustExec("INSERT INTO users (id, first_name, last_name, email, password) VALUES (?, ?, ?, ?, ?)", u.Id, u.FirstName, u.LastName, u.Email, string(u.Password))

	return sr, s
}

func TestNewSessionRepository(t *testing.T) {
	t.Parallel()

	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	sr := repositories.NewSessionRepository(store)

	assert.NotNil(t, sr)
}

func TestSessionRepository_Add(t *testing.T) {
	t.Parallel()
	sr, _ := setupSessionRepository()

	u, _ := models.NewUser("Robin", "", "")
	s := models.NewSession(u)
	err := sr.Add(s)
	assert.NoError(t, err)
}

func TestSessionRepository_Get(t *testing.T) {
	t.Parallel()
	sr, s := setupSessionRepository()

	got, err := sr.Get(s.Id)
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, s.User.Id, got.User.Id)
}

func TestSessionRepository_Delete(t *testing.T) {
	t.Parallel()
	sr, _ := setupSessionRepository()

	err := sr.Delete("session-id")
	assert.NoError(t, err)
}
