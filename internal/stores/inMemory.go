package stores

import "github.com/Linkinlog/loggr/internal/models"

func NewInMemory(g []*models.Garden) *InMemory {
	return &InMemory{
		Gardens: g,
	}
}

type InMemory struct {
	Gardens []*models.Garden
}

func (im *InMemory) ListGardens() ([]*models.Garden, error) {
	return im.Gardens, nil
}

func (im *InMemory) AddGarden(g *models.Garden) (string, error) {
	if g.Id() == "" {
		return "", models.ErrEmptyID
	}
	im.Gardens = append(im.Gardens, g)
	return g.Id(), nil
}

func (im *InMemory) GetGarden(id string) (*models.Garden, error) {
	for _, g := range im.Gardens {
		if g.Id() == id {
			return g, nil
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

func (im *InMemory) DeleteGarden(id string) error {
	for i, g := range im.Gardens {
		if g.Id() == id {
			im.Gardens = append(im.Gardens[:i], im.Gardens[i+1:]...)
			return nil
		}
	}
	return models.ErrNotFound
}
