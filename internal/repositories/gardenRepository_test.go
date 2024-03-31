package repositories_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/stores"
	"github.com/stretchr/testify/assert"
)

func TestNewGardenRepository(t *testing.T) {
	t.Parallel()

	gr := repositories.NewGardenRepository(nil)
	assert.NotNil(t, gr)
}

func TestGardenRepository_AddGarden(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		garden        *models.Garden
		expectedError error
	}{
		"empty garden": {
			garden:        &models.Garden{},
			expectedError: models.ErrEmptyID,
		},
		"valid garden": {
			garden:        models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{}),
			expectedError: nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			store := stores.NewInMemory(nil)
			gr := repositories.NewGardenRepository(store)

			id, err := gr.AddGarden(tc.garden)

			assert.Equal(t, tc.garden.Id(), id)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGardenRepository_GetGarden(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	got, err := gr.GetGarden(id)

	assert.Equal(t, g, got)
	assert.NoError(t, err)
}

func TestGardenRepository_UpdateGarden(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	g.Name = "garden 2"
	g.Location = "location 2"
	g.Description = "description 2"

	err := gr.UpdateGarden(id, g)

	got, _ := gr.GetGarden(id)

	assert.NoError(t, err)
	assert.Equal(t, g, got)
}

func TestGardenRepository_DeleteGarden(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	err := gr.DeleteGarden(id)

	assert.NoError(t, err)

	_, err = gr.GetGarden(id)

	assert.Error(t, err, models.ErrNotFound)
}

func TestGardenRepository_ListGardens(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	_, _ = gr.AddGarden(g)

	gardens, err := gr.ListGardens()

	assert.Len(t, gardens, 1)
	assert.Equal(t, g, gardens[0])
	assert.NoError(t, err)
}

func TestGardenRepository_AddItemToGarden(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	image := models.NewImage("id", "https://example.com", "", "")
	item := models.NewItem("item 1", image, models.Plant, [5]*models.Field{})

	err := gr.AddItemToGarden(id, item)

	got, _ := gr.GetGarden(id)

	assert.NoError(t, err)
	assert.Len(t, got.Inventory, 1)
	assert.Equal(t, item, got.Inventory[0])
}

func TestGardenRepository_ListGardenInventory(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	inventory, err := gr.ListGardenInventory(id)

	assert.Equal(t, g.Inventory, inventory)
	assert.NoError(t, err)
}

func TestGardenRepository_RemoveItemFromGarden(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	image := models.NewImage("id", "https://example.com", "", "")
	item := models.NewItem("item 1", image, models.Plant, [5]*models.Field{})
	_ = gr.AddItemToGarden(id, item)

	err := gr.RemoveItemFromGarden(id, item.Id())

	got, _ := gr.GetGarden(id)

	assert.NoError(t, err)
	assert.Len(t, got.Inventory, 0)
}

func TestGardenRepository_GetItemFromGarden(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	image := models.NewImage("id", "https://example.com", "", "")
	item := models.NewItem("item 1", image, models.Plant, [5]*models.Field{})
	_ = gr.AddItemToGarden(id, item)

	got, err := gr.GetItemFromGarden(id, item.Id())

	assert.Equal(t, item, got)
	assert.NoError(t, err)
}

func TestGardenRepository_UpdateItemInGarden(t *testing.T) {
	t.Parallel()

	store := stores.NewInMemory(nil)
	gr := repositories.NewGardenRepository(store)

	g := models.NewGarden("garden 1", "location 1", "description 1", models.NewImage("image 1", "url", "", ""), []*models.Item{})
	id, _ := gr.AddGarden(g)

	image := models.NewImage("id", "https://example.com", "", "")
	item := models.NewItem("item 1", image, models.Plant, [5]*models.Field{})
	_ = gr.AddItemToGarden(id, item)

	item.Name = "item 2"
	item.Image = models.NewImage("id", "https://example.com", "", "")

	err := gr.UpdateItemInGarden(id, item.Id(), item)

	got, _ := gr.GetGarden(id)

	assert.NoError(t, err)
	assert.Equal(t, item, got.Inventory[0])
}
