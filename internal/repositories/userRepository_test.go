package repositories_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/stores"
	"github.com/stretchr/testify/assert"
)

func TestNewUserRepository(t *testing.T) {
	t.Parallel()
	store := stores.NewInMemory(nil, nil)
	ur := repositories.NewUserRepository(store)

	assert.NotNil(t, ur)
}

func TestUserRepository_Add(t *testing.T) {
	t.Parallel()
	store := stores.NewInMemory(nil, nil)
	ur := repositories.NewUserRepository(store)

	u, _ := models.NewUser("Batman", "idk@lol.com", "password123")

	id, err := ur.Add(u)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	_, err = ur.Add(nil)
	assert.Error(t, err)
}

func TestUserRepository_Get(t *testing.T) {
	t.Parallel()
	store := stores.NewInMemory(nil, nil)
	ur := repositories.NewUserRepository(store)

	_, err := ur.Get("id")
	assert.Error(t, err)
}

func TestUserRepository_Update(t *testing.T) {
	t.Parallel()
	store := stores.NewInMemory(nil, nil)
	ur := repositories.NewUserRepository(store)

	err := ur.Update("id", nil)
	assert.Error(t, err)
}

func TestUserRepository_Delete(t *testing.T) {
	t.Parallel()
	store := stores.NewInMemory(nil, nil)
	ur := repositories.NewUserRepository(store)

	err := ur.Delete("id")
	assert.Error(t, err)
}

func TestUserRepository_List(t *testing.T) {
	t.Parallel()
	store := stores.NewInMemory(nil, nil)
	ur := repositories.NewUserRepository(store)

	users, err := ur.List()
	assert.NoError(t, err)
	assert.Empty(t, users)
}
