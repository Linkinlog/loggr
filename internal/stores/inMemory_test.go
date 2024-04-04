package stores_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInMemory(t *testing.T) {
	t.Parallel()

	im := stores.NewInMemory(nil)
	assert.NotNil(t, im)
}

func TestInMemory_ListUsers(t *testing.T) {
	t.Parallel()

	users := []*models.User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@hotmail.com",
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane@gmail.com",
		},
	}

	im := stores.NewInMemory(users)
	u, err := im.ListUsers()
	require.NoError(t, err)
	assert.Equal(t, users, u)
}

func TestInMemory_AddUser(t *testing.T) {
	t.Parallel()

	users := []*models.User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@hotmail.com",
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane@gmail.com",
		},
	}

	im := stores.NewInMemory(users)
	u, _ := models.NewUser("Alice", "Doe", "alice@aol.com")
	id, err := im.AddUser(u)
	require.NoError(t, err)
	assert.Equal(t, u.Id(), id)
}
