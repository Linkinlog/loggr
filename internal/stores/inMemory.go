package stores

import "github.com/Linkinlog/loggr/internal/models"

func NewInMemory(g []*models.Garden, u []*models.User) *InMemory {
	return &InMemory{
		Gardens: g,
		Users:   u,
	}
}

type InMemory struct {
	Gardens []*models.Garden
	Users   []*models.User
}

func (im *InMemory) ListGardens() ([]*models.Garden, error) {
	return im.Gardens, nil
}

func (im *InMemory) ListUsers() ([]*models.User, error) {
	return im.Users, nil
}

func (im *InMemory) AddGarden(g *models.Garden) (string, error) {
	if g.Id() == "" {
		return "", models.ErrEmptyID
	}
	im.Gardens = append(im.Gardens, g)
	return g.Id(), nil
}

func (im *InMemory) AddUser(u *models.User) (string, error) {
	if u.Id() == "" {
		return "", models.ErrEmptyID
	}
	im.Users = append(im.Users, u)
	return u.Id(), nil
}

func (im *InMemory) GetGarden(id string) (*models.Garden, error) {
	for _, g := range im.Gardens {
		if g.Id() == id {
			return g, nil
		}
	}
	return nil, models.ErrNotFound
}

func (im *InMemory) GetUser(email string) (*models.User, error) {
	for _, u := range im.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, models.ErrNotFound
}

func (im *InMemory) UpdateGarden(id string, g *models.Garden) error {
	for i, garden := range im.Gardens {
		if garden.Id() == id {
			im.Gardens[i].Name = g.Name
			im.Gardens[i].Location = g.Location
			im.Gardens[i].Description = g.Description
			im.Gardens[i].Inventory = g.Inventory
			return nil
		}
	}
	return models.ErrNotFound
}

func (im *InMemory) UpdateUser(id string, u *models.User) error {
	for i, user := range im.Users {
		if user.Id() == id {
			im.Users[i].FirstName = u.FirstName
			im.Users[i].LastName = u.LastName
			im.Users[i].Email = u.Email
			return nil
		}
	}
	return models.ErrNotFound
}

func (im *InMemory) DeleteGarden(id string) error {
	for i, g := range im.Gardens {
		if g.Id() == id {
			im.Gardens = append(im.Gardens[:i], im.Gardens[i+1:]...)
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
