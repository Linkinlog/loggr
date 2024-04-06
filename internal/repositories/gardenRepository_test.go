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

func setupGardenRepository() (*repositories.GardenRepository, *models.Garden) {
	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	gr := repositories.NewGardenRepository(store)

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
	g := models.NewGarden("garden 1", "location 1", "description 1", "", []*models.Item{i})

	store.MustExec("INSERT INTO gardens (id, user_id, name, location, description, image) VALUES (?, ?, ?, ?, ?, ?)", g.Id, 0, g.Name, g.Location, g.Description, g.Image)
	store.MustExec("INSERT INTO items (id, garden_id, name, image, type, field1, field2, field3, field4, field5) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", i.Id, g.Id, i.Name, i.Image, i.Type, i.Fields[0], i.Fields[1], i.Fields[2], i.Fields[3], i.Fields[4])

	return gr, g
}

func TestNewGardenRepository(t *testing.T) {
	t.Parallel()

	store := &stores.SqliteStore{
		DB: sqlx.MustOpen("sqlite", ":memory:"),
	}
	gr := repositories.NewGardenRepository(store)

	assert.NotNil(t, gr)
}

func TestGardenRepository_Add(t *testing.T) {
	t.Parallel()
	gr, _ := setupGardenRepository()

	g := models.NewGarden("garden 2", "location 2", "description 2", "", []*models.Item{})
	id, err := gr.Add(g)
	assert.NoError(t, err)

	assert.NotEmpty(t, id)
}

func TestGardenRepository_Get(t *testing.T) {
	t.Parallel()
	gr, g := setupGardenRepository()

	got, err := gr.Get(g.Id)
	assert.NoError(t, err)
	assert.NotNil(t, got)
}

func TestGardenRepository_GetInventory(t *testing.T) {
	t.Parallel()
	gr, g := setupGardenRepository()

	items, err := gr.GetInventory(g.Id)
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 1, len(items))
}

func TestGardenRepository_Update(t *testing.T) {
	t.Parallel()
	gr, g := setupGardenRepository()

	err := gr.Update(g)
	assert.NoError(t, err)
}

func TestGardenRepository_Delete(t *testing.T) {
	t.Parallel()
	gr, g := setupGardenRepository()

	err := gr.Delete(g.Id)
	assert.NoError(t, err)
}

func TestGardenRepository_GetItemsForGarden(t *testing.T) {
	t.Parallel()
	gr, g := setupGardenRepository()

	items, err := gr.GetItemsForGarden(g.Id)
	assert.NoError(t, err)
	assert.NotNil(t, items)
}
