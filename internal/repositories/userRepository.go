package repositories

import (
	"errors"

	"github.com/Linkinlog/loggr/internal/models"
)

var ErrNilUser = errors.New("nil user")

type UserStorer interface {
	AddUser(u *models.User) (string, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, u *models.User) error
}

func NewUserRepository(s UserStorer) *UserRepository {
	return &UserRepository{
		store: s,
	}
}

type UserRepository struct {
	store UserStorer
}

func (ur *UserRepository) Add(u *models.User) (string, error) {
	if u == nil {
		return "", ErrNilUser
	}
	return ur.store.AddUser(u)
}

func (ur *UserRepository) Get(id string) (*models.User, error) {
	return ur.store.GetUser(id)
}

func (ur *UserRepository) Update(id string, u *models.User) error {
	return ur.store.UpdateUser(id, u)
}

func (ur *UserRepository) Save(u *models.User) (string, error) {
	if u == nil {
		return "", ErrNilUser
	}

	id, err := ur.Add(u)
	if errors.Is(err, models.ErrAlreadyExists) {
		return id, ur.Update(id, u)
	}

	return id, err
}
