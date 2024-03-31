package stores_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemory(t *testing.T) {
	t.Parallel()

	im := stores.NewInMemory(nil)
	assert.NotNil(t, im)
}

func TestInMemory_ListGardens(t *testing.T) {
	t.Parallel()

	im := stores.NewInMemory(nil)
	image := models.NewImage("id", "https://example.com")
	g1 := models.NewGarden("garden 1", "location 1", "description 1", image, []*models.Item{})
	g2 := models.NewGarden("garden 2", "location 2", "description 2", image, []*models.Item{})
	g3 := models.NewGarden("garden 3", "location 3", "description 3", image, []*models.Item{})
	im.Gardens = []*models.Garden{g1, g2, g3}

	got, err := im.ListGardens()

	assert.NoError(t, err)
	assert.Len(t, got, 3)
	assert.Equal(t, g1, got[0])
	assert.Equal(t, g2, got[1])
	assert.Equal(t, g3, got[2])
}

func TestInMemory_AddGarden(t *testing.T) {
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
			garden:        models.NewGarden("garden 1", "location 1", "description 1", nil, []*models.Item{}),
			expectedError: nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			im := stores.NewInMemory(nil)
			id, err := im.AddGarden(tc.garden)

			if tc.expectedError != nil {
				assert.Empty(t, id)
				assert.Error(t, err, tc.expectedError)
				return
			}
			assert.NotEmpty(t, id)
			assert.NoError(t, err)
			assert.Len(t, im.Gardens, 1)
			assert.Equal(t, tc.garden, im.Gardens[0])
		})
	}
}

func TestInMemory_GetGarden(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expectedError error
		modFunc       func(*stores.InMemory)
	}{
		"valid garden": {
			expectedError: nil,
			modFunc:       func(im *stores.InMemory) {},
		},
		"valid garden, not found": {
			expectedError: models.ErrNotFound,
			modFunc: func(im *stores.InMemory) {
				im.Gardens = []*models.Garden{}
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			im := stores.NewInMemory(nil)
			g := models.NewGarden("garden 1", "location 1", "description 1", nil, []*models.Item{})
			_, _ = im.AddGarden(g)

			tc.modFunc(im)

			got, err := im.GetGarden(g.Id())

			if tc.expectedError != nil {
				assert.Nil(t, got)
				assert.Error(t, err, tc.expectedError)
				return
			}

			assert.NotNil(t, got)
			assert.NoError(t, err)
			assert.Equal(t, g, got)
		})
	}
}

func TestInMemory_UpdateGarden(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expectedError error
		newGarden     *models.Garden
		modFunc       func(*stores.InMemory)
	}{
		"valid garden": {
			expectedError: nil,
			newGarden:     models.NewGarden("garden 2", "location 2", "description 2", nil, []*models.Item{}),
			modFunc:       func(im *stores.InMemory) {},
		},
		"valid garden, not found": {
			expectedError: models.ErrNotFound,
			newGarden:     models.NewGarden("garden 2", "location 2", "description 2", nil, []*models.Item{}),
			modFunc: func(im *stores.InMemory) {
				im.Gardens = []*models.Garden{}
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			im := stores.NewInMemory(nil)
			g := models.NewGarden("garden 1", "location 1", "description 1", nil, []*models.Item{})
			_, _ = im.AddGarden(g)

			tc.modFunc(im)

			err := im.UpdateGarden(g.Id(), tc.newGarden)

			if tc.expectedError != nil {
				assert.Error(t, err, tc.expectedError)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, im.Gardens, 1)
			assert.Equal(t, g.Id(), im.Gardens[0].Id())
			assert.Equal(t, tc.newGarden.Name, im.Gardens[0].Name)
			assert.Equal(t, tc.newGarden.Location, im.Gardens[0].Location)
			assert.Equal(t, tc.newGarden.Description, im.Gardens[0].Description)
			assert.Equal(t, tc.newGarden.Inventory, im.Gardens[0].Inventory)
		})
	}
}

func TestInMemory_DeleteGarden(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expectedError error
		modFunc       func(*stores.InMemory)
	}{
		"valid garden": {
			expectedError: nil,
			modFunc:       func(im *stores.InMemory) {},
		},
		"valid garden, not found": {
			expectedError: models.ErrNotFound,
			modFunc: func(im *stores.InMemory) {
				im.Gardens = []*models.Garden{}
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			im := stores.NewInMemory(nil)
			g := models.NewGarden("garden 1", "location 1", "description 1", nil, []*models.Item{})
			_, _ = im.AddGarden(g)

			tc.modFunc(im)

			err := im.DeleteGarden(g.Id())

			if tc.expectedError != nil {
				assert.Error(t, err, tc.expectedError)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, im.Gardens, 0)
		})
	}
}
