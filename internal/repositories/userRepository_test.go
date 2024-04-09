package repositories_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/stores"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func setupUserRepository() (*repositories.UserRepository, *models.User, *models.Garden) {
	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	ur := repositories.NewUserRepository(store)

	schema := `CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		reset_code TEXT NOT NULL
	);
CREATE TABLE IF NOT EXISTS gardens (
		id TEXT PRIMARY KEY NOT NULL,
		user_id TEXT NOT NULL,
		name TEXT NOT NULL,
		location TEXT NOT NULL,
		description TEXT NOT NULL,
		image TEXT NOT NULL
	);
	`

	store.MustExec(schema)

	u, _ := models.NewUser("Batman", "idk@lol.com", "password123")
	store.MustExec("INSERT INTO users (id, first_name, last_name, email, password, reset_code) VALUES (?, ?, ?, ?, ?, ?)", u.Id, u.FirstName, u.LastName, u.Email, string(u.Password), u.ResetCode)

	g := models.NewGarden("garden 1", "location 1", "description 1", "https://image.com", []*models.Item{})

	store.MustExec("INSERT INTO gardens (id, user_id, name, location, description, image) VALUES (?, ?, ?, ?, ?, ?)", g.Id, u.Id, g.Name, g.Location, g.Description, g.Image)

	return ur, u, g
}

func TestNewUserRepository(t *testing.T) {
	t.Parallel()

	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	ur := repositories.NewUserRepository(store)

	assert.NotNil(t, ur)
}

func TestUserRepository_Add(t *testing.T) {
	t.Parallel()
	ur, _, _ := setupUserRepository()

	u, _ := models.NewUser("Robin", "idk@lol.com", "password123")
	id, err := ur.Add(u)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	_, err = ur.Add(nil)
	assert.Error(t, err)
}

func TestUserRepository_Get(t *testing.T) {
	t.Parallel()
	ur, u, _ := setupUserRepository()

	_, err := ur.Get("id")
	assert.Error(t, err)

	user, err := ur.Get(u.Id)

	assert.NoError(t, err)
	assert.Equal(t, u.FirstName, user.FirstName)
	assert.Equal(t, u.LastName, user.LastName)
	assert.Equal(t, u.Email, user.Email)
	assert.Equal(t, u.Password, user.Password)
}

func TestUserRepository_Update(t *testing.T) {
	t.Parallel()
	ur, u, _ := setupUserRepository()

	err := ur.Update("id", nil)
	assert.Error(t, err)

	u.FirstName = "Bruce"
	err = ur.Update(u.Id, u)
	assert.NoError(t, err)

	user, err := ur.Get(u.Id)
	assert.NoError(t, err)
	assert.Equal(t, u.FirstName, user.FirstName)
}

func TestUserRepository_Delete(t *testing.T) {
	t.Parallel()
	ur, u, _ := setupUserRepository()

	err := ur.Delete("id")
	assert.Error(t, err)

	err = ur.Delete(u.Id)

	assert.NoError(t, err)
}

func TestUserRepository_AssociateGardens(t *testing.T) {
	t.Parallel()
	ur, u, g := setupUserRepository()

	_, err := ur.GetGardensForUser("id")
	assert.Error(t, err)

	err = ur.AssociateGarden(u.Id, g)
	assert.NoError(t, err)
}

func TestUserRepository_GetGardens(t *testing.T) {
	t.Parallel()
	ur, u, g := setupUserRepository()

	_, err := ur.GetGardensForUser("id")
	assert.Error(t, err)

	err = ur.AssociateGarden(u.Id, g)
	assert.NoError(t, err)

	gardens, err := ur.GetGardensForUser(u.Id)

	assert.NoError(t, err)
	assert.Len(t, gardens, 1)
	assert.Equal(t, g.Id, gardens[0].Id)
	assert.Equal(t, g.Name, gardens[0].Name)
	assert.Equal(t, g.Location, gardens[0].Location)
	assert.Equal(t, g.Description, gardens[0].Description)
	assert.Equal(t, g.Image, gardens[0].Image)
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	t.Parallel()
	ur, _, _ := setupUserRepository()

	_, err := ur.GetByEmail("idk@lol.com")
	assert.NoError(t, err)

	_, err = ur.GetByEmail("wrong")
	assert.Error(t, err)
}
