package models_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	t.Parallel()
	title := "field 1"
	description := "field 1 description"
	f := models.NewField(title, description)

	assert.Equal(t, title, f.Title)
	assert.Equal(t, description, f.Description)
}
