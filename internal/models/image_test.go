package models_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	t.Parallel()

	image := models.NewImage("id", "https://example.com")
	assert.NotNil(t, image)
}
