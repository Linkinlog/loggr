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
	inventory := []*models.Item{
		models.NewItem("item 1", "", models.Plant, [5]string{
			"field 1 description",
			"field 2 description",
			"field 3 description",
			"field 4 description",
		}),
		models.NewItem("item 2", "", models.Plant, [5]string{}),
	}

	g := models.NewGarden(name, location, description, "", inventory)

	assert.Equal(t, name, g.Name)
	assert.Equal(t, location, g.Location)
	assert.Equal(t, description, g.Description)
}

func TestGarden_Id(t *testing.T) {
	t.Parallel()
	g := models.NewGarden("garden 1", "location 1", "description 1", "", []*models.Item{})
	g1 := models.NewGarden("garden 1", "location 1", "description 1", "", []*models.Item{})

	assert.NotEmpty(t, g.Id)
	assert.NotEmpty(t, g1.Id)
	assert.NotEqual(t, g.Id, g1.Id)
}
