package models_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	t.Parallel()
	name := "item 1"
	image := "image 1"
	fields := [5]*models.Field{
		models.NewField("field 1", "field 1 description"),
		models.NewField("field 2", "field 2 description"),
		models.NewField("field 3", "field 3 description"),
		models.NewField("field 4", "field 4 description"),
		models.NewField("field 5", "field 5 description"),
	}

	i := models.NewItem(name, image, fields)

	assert.Equal(t, name, i.Name)
	assert.Equal(t, image, i.Image)
	assert.Equal(t, fields, i.Fields)
}
