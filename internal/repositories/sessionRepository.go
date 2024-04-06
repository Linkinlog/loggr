package repositories

import (
	"errors"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
)

var ErrNilSession = errors.New("nil session")

func NewSessionRepository(db *stores.SqliteStore) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

type SessionRepository struct {
	db *stores.SqliteStore
}

func (sr *SessionRepository) Add(s *models.Session) error {
	if s == nil {
		return ErrNilSession
	}
	query := `INSERT INTO sessions (id, user_id) VALUES (?, ?)`

	_, err := sr.db.Exec(query, s.Id, s.User.Id)

	return err
}

func (sr *SessionRepository) Get(sessionId string) (*models.Session, error) {
	query := `SELECT sessions.id as s_id, users.* FROM sessions JOIN users ON sessions.user_id = users.id WHERE sessions.id = ?`

	var dbSession struct {
		Id string `db:"s_id"`
		*models.User
	}
	err := sr.db.Get(&dbSession, query, sessionId)
	if err != nil {
		return nil, err
	}

	s := models.NewSession(dbSession.User)
	s.Id = dbSession.Id

	return s, nil
}

func (sr *SessionRepository) Delete(sessionId string) error {
	query := `DELETE FROM sessions WHERE id = ?`

	_, err := sr.db.Exec(query, sessionId)

	return err
}
