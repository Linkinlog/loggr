package repositories_test

import (
	"testing"

	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/repositories"
	"github.com/Linkinlog/loggr/internal/stores"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func setupItemRepository() (*repositories.ItemRepository, *models.Item) {
	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	ir := repositories.NewItemRepository(store)

	schema := `CREATE TABLE IF NOT EXISTS gardens (
		id TEXT PRIMARY KEY NOT NULL,
		user_id TEXT NOT NULL,
		name TEXT NOT NULL,
		location TEXT NOT NULL,
		description TEXT NOT NULL,
		image TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS items (
		id TEXT PRIMARY KEY NOT NULL,
		garden_id TEXT NOT NULL,
		name TEXT NOT NULL,
		image TEXT NOT NULL,
		type INTEGER NOT NULL,
		field1 TEXT NOT NULL,
		field2 TEXT NOT NULL,
		field3 TEXT NOT NULL,
		field4 TEXT NOT NULL,
		field5 TEXT NOT NULL
	);
	`

	store.MustExec(schema)

	fields := [5]string{"field1", "field2", "field3", "field4", "field5"}
	i := models.NewItem("item 1", "description 1", models.Plant, fields)

	store.MustExec("INSERT INTO items (id, garden_id, name, image, type, field1, field2, field3, field4, field5) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", i.Id, "1", i.Name, i.Image, i.Type, i.Fields[0], i.Fields[1], i.Fields[2], i.Fields[3], i.Fields[4])

	return ir, i
}

func TestNewItemRepository(t *testing.T) {
	t.Parallel()

	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	repo := repositories.NewItemRepository(store)

	assert.NotNil(t, repo)
}

func TestNewItemRepository_GetItem(t *testing.T) {
	t.Parallel()

	repo, _ := setupItemRepository()

	assert.NotNil(t, repo)

	_, err := repo.Get("1")
	assert.NotNil(t, err)
}

func TestNewItemRepository_AddItem(t *testing.T) {
	t.Parallel()

	repo, _ := setupItemRepository()

	assert.NotNil(t, repo)

	i := &models.Item{
		Id:     "1",
		Name:   "item 1",
		Image:  "image 1",
		Type:   models.ItemType(1),
		Fields: [5]string{"field 1", "field 2", "field 3", "field 4", "field 5"},
	}

	id, err := repo.Add(i)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func TestNewItemRepository_UpdateItem(t *testing.T) {
	t.Parallel()

	repo, i := setupItemRepository()

	assert.NotNil(t, repo)

	err := repo.Update(i)
	assert.Nil(t, err)
}

func TestNewItemRepository_DeleteItem(t *testing.T) {
	t.Parallel()

	repo, i := setupItemRepository()

	assert.NotNil(t, repo)

	err := repo.Delete(i.Id)
	assert.Nil(t, err)
}
