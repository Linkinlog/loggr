package models_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewGarden(t *testing.T) {
	t.Parallel()
	name := "garden 1"
	location := "location 1"
	description := "description 1"
	image := models.NewImage("id", "https://example.com")
	inventory := []*models.Item{
		models.NewItem("item 1", image, [5]*models.Field{
			models.NewField("field 1", "field 1 description"),
			models.NewField("field 2", "field 2 description"),
			models.NewField("field 3", "field 3 description"),
			models.NewField("field 4", "field 4 description"),
		}),
		models.NewItem("item 2", image, [5]*models.Field{}),
	}

	g := models.NewGarden(name, location, description, image, inventory)

	assert.Equal(t, name, g.Name)
	assert.Equal(t, location, g.Location)
	assert.Equal(t, description, g.Description)
	assert.Equal(t, inventory, g.Inventory)
}

func TestGarden_Id(t *testing.T) {
	t.Parallel()
	image := models.NewImage("id", "https://example.com")
	g := models.NewGarden("garden 1", "location 1", "description 1", image, []*models.Item{})
	g1 := models.NewGarden("garden 1", "location 1", "description 1", image, []*models.Item{})

	assert.NotEmpty(t, g.Id())
	assert.NotEmpty(t, g1.Id())
	assert.NotEqual(t, g.Id(), g1.Id())
}

func TestGarden_AddItem(t *testing.T) {
	t.Parallel()
	image := models.NewImage("id", "https://example.com")
	g := models.NewGarden("garden 1", "location 1", "description 1", image, []*models.Item{})
	item := models.NewItem("item 1", image, [5]*models.Field{})

	g.AddItem(item)

	assert.Len(t, g.Inventory, 1)
	assert.Equal(t, item, g.Inventory[0])
}
