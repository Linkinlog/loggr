package repositories

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
)

var (
	ErrNilUser   = errors.New("nil user")
	ErrNilGarden = errors.New("nil garden")
	ErrNilItem   = errors.New("nil item")
)

func NewUserRepository(db *stores.SqliteStore) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *stores.SqliteStore
}

func (ur *UserRepository) Add(u *models.User) (string, error) {
	if u == nil {
		return "", ErrNilUser
	}
	query := `INSERT INTO users (id, first_name, last_name, email, password, image, reset_code) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := ur.db.Exec(query, u.Id, u.FirstName, u.LastName, u.Email, string(u.Password), u.Image, u.ResetCode)

	return u.Id, err
}

func (ur *UserRepository) GenerateResetCode(email string) (string, error) {
	query := `UPDATE users SET reset_code = ? WHERE email = ?`

	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	code := fmt.Sprintf("%X", bytes)

	_, err := ur.db.Exec(query, code, email)

	return code, err
}

func (ur *UserRepository) ClearResetCode(email string) error {
	query := `UPDATE users SET reset_code = ? WHERE email = ?`

	_, err := ur.db.Exec(query, "", email)

	return err
}

func (ur *UserRepository) Get(id string) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = ?`

	u := &models.User{}
	err := ur.db.Get(u, query, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = ?`

	u := &models.User{}
	err := ur.db.Get(u, query, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) GetByResetCode(code string) (*models.User, error) {
	query := `SELECT * FROM users WHERE reset_code = ?`

	u := &models.User{}
	err := ur.db.Get(u, query, code)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) Update(id string, u *models.User) error {
	if u == nil {
		return ErrNilUser
	}
	query := `UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ?, image = ? WHERE id = ?`

	_, err := ur.db.Exec(query, u.FirstName, u.LastName, u.Email, u.Password, u.Image, id)

	return err
}

func (ur *UserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := ur.db.Exec(query, id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (ur *UserRepository) GetGardensForUser(id string) ([]*models.Garden, error) {
	query := `SELECT id, name, location,image, description FROM gardens WHERE user_id = ?`

	gardens := []*models.Garden{}
	err := ur.db.Select(&gardens, query, id)
	if err != nil {
		return nil, err
	}

	if len(gardens) == 0 {
		return nil, models.ErrNotFound
	}

	return gardens, nil
}

func (ur *UserRepository) AssociateGarden(id string, g *models.Garden) error {
	if g == nil {
		return ErrNilGarden
	}
	query := `UPDATE gardens SET user_id = ? WHERE id = ?`

	_, err := ur.db.Exec(query, id, g.Id)

	return err
}
