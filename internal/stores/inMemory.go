package stores

import "github.com/Linkinlog/loggr/internal/models"

func NewInMemory(u []*models.User) *InMemory {
	return &InMemory{
		Users: u,
	}
}

type InMemory struct {
	Users []*models.User
}

func (im *InMemory) ListUsers() ([]*models.User, error) {
	return im.Users, nil
}

func (im *InMemory) AddUser(u *models.User) (string, error) {
	if u.Id() == "" {
		return "", models.ErrEmptyID
	}
	if _, err := im.GetUser(u.Email); err == nil {
		return "", models.ErrAlreadyExists
	}
	im.Users = append(im.Users, u)
	return u.Id(), nil
}

func (im *InMemory) GetUser(email string) (*models.User, error) {
	for _, u := range im.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, models.ErrNotFound
}

func (im *InMemory) UpdateUser(id string, u *models.User) error {
	for i, user := range im.Users {
		if user.Id() == id {
			im.Users[i].FirstName = u.FirstName
			im.Users[i].LastName = u.LastName
			im.Users[i].Email = u.Email
			im.Users[i].Gardens = u.Gardens
			return nil
		}
	}
	return models.ErrNotFound
}

func (im *InMemory) DeleteUser(id string) error {
	for i, u := range im.Users {
		if u.Id() == id {
			im.Users = append(im.Users[:i], im.Users[i+1:]...)
			return nil
		}
	}
	return models.ErrNotFound
}
