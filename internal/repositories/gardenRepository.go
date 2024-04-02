package repositories

import "github.com/Linkinlog/loggr/internal/models"

type GardenStorer interface {
	AddGarden(g *models.Garden) (string, error)
	GetGarden(id string) (*models.Garden, error)
	UpdateGarden(id string, g *models.Garden) error
	DeleteGarden(id string) error
	ListGardens() ([]*models.Garden, error)
}

func NewGardenRepository(s GardenStorer) *GardenRepository {
	return &GardenRepository{
		store: s,
	}
}

type GardenRepository struct {
	store GardenStorer
}

func (gr *GardenRepository) Add(g *models.Garden) (string, error) {
	return gr.store.AddGarden(g)
}

func (gr *GardenRepository) Get(id string) (*models.Garden, error) {
	return gr.store.GetGarden(id)
}

func (gr *GardenRepository) Update(id string, g *models.Garden) error {
	return gr.store.UpdateGarden(id, g)
}

func (gr *GardenRepository) Delete(id string) error {
	return gr.store.DeleteGarden(id)
}

func (gr *GardenRepository) List() ([]*models.Garden, error) {
	return gr.store.ListGardens()
}

func (gr *GardenRepository) ListGardenInventory(gardenID string) ([]*models.Item, error) {
	g, err := gr.Get(gardenID)
	if err != nil {
		return nil, err
	}

	return g.Inventory, nil
}

func (gr *GardenRepository) AddItemToGarden(gardenID string, item *models.Item) error {
	g, err := gr.Get(gardenID)
	if err != nil {
		return err
	}

	g.AddItem(item)
	return gr.Update(gardenID, g)
}

func (gr *GardenRepository) RemoveItemFromGarden(gardenID, itemID string) error {
	g, err := gr.Get(gardenID)
	if err != nil {
		return err
	}

	for i, item := range g.Inventory {
		if item.Id() == itemID {
			g.Inventory = append(g.Inventory[:i], g.Inventory[i+1:]...)
			return gr.Update(gardenID, g)
		}
	}

	return models.ErrNotFound
}

func (gr *GardenRepository) GetItemFromGarden(gardenID, itemID string) (*models.Item, error) {
	g, err := gr.Get(gardenID)
	if err != nil {
		return nil, err
	}

	for _, item := range g.Inventory {
		if item.Id() == itemID {
			return item, nil
		}
	}

	return nil, models.ErrNotFound
}

func (gr *GardenRepository) UpdateItemInGarden(gardenID, itemID string, item *models.Item) error {
	g, err := gr.Get(gardenID)
	if err != nil {
		return err
	}

	for i, it := range g.Inventory {
		if it.Id() == itemID {
			g.Inventory[i].Name = item.Name
			g.Inventory[i].Image = item.Image
			g.Inventory[i].Type = item.Type
			g.Inventory[i].Fields = item.Fields

			return gr.Update(gardenID, g)
		}
	}

	return models.ErrNotFound
}
