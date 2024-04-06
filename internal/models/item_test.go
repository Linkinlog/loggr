package models_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	t.Parallel()
	name := "item 1"
	fields := [5]string{
		"field 1 description",
		"field 2 description",
		"field 3 description",
		"field 4 description",
		"field 5 description",
	}

	i := models.NewItem(name, "", models.Plant, fields)

	assert.Equal(t, name, i.Name)
	assert.Equal(t, "", i.Image)
	assert.Equal(t, fields, i.Fields)
}
