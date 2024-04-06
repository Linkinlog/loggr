package repositories

import (
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
)

func NewGardenRepository(db *stores.SqliteStore) *GardenRepository {
	return &GardenRepository{
		db: db,
	}
}

type GardenRepository struct {
	db *stores.SqliteStore
}

func (gr *GardenRepository) Add(g *models.Garden) (string, error) {
	if g == nil {
		return "", ErrNilGarden
	}
	query := `INSERT INTO gardens (id, user_id, name, location, description, image) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := gr.db.Exec(query, g.Id, 0, g.Name, g.Location, g.Description, g.Image)

	return g.Id, err
}

func (gr *GardenRepository) Get(id string) (*models.Garden, error) {
	query := `SELECT id, name, location, description, image FROM gardens WHERE id = ?`

	g := &models.Garden{}
	err := gr.db.Get(g, query, id)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (gr *GardenRepository) GetInventory(gardenId string) ([]*models.Item, error) {
	query := `SELECT id, name, image, type, field1, field2, field3, field4, field5 FROM items WHERE garden_id = ?`

	rows, err := gr.db.Queryx(query, gardenId)
	if err != nil {
		return nil, err
	}

	items := []*models.Item{}
	for rows.Next() {
		var dbItem struct {
			*models.Item
			Type   int    `db:"type"`
			Field1 string `db:"field1"`
			Field2 string `db:"field2"`
			Field3 string `db:"field3"`
			Field4 string `db:"field4"`
			Field5 string `db:"field5"`
		}
		err = rows.StructScan(&dbItem)
		if err != nil {
			return nil, err
		}

		item := &models.Item{
			Id:    dbItem.Id,
			Name:  dbItem.Name,
			Image: dbItem.Image,
			Type:  models.ItemType(dbItem.Type),
		}
		fields := [5]string{
			dbItem.Field1,
			dbItem.Field2,
			dbItem.Field3,
			dbItem.Field4,
			dbItem.Field5,
		}
		item.Fields = fields

		items = append(items, item)
	}

	return items, nil
}

func (gr *GardenRepository) Update(g *models.Garden) error {
	if g == nil {
		return ErrNilGarden
	}
	query := `UPDATE gardens SET name = ?, location = ?, description = ?, image = ? WHERE id = ?`

	_, err := gr.db.Exec(query, g.Name, g.Location, g.Description, g.Image, g.Id)

	return err
}

func (gr *GardenRepository) Delete(id string) error {
	query := `DELETE FROM gardens WHERE id = ?`

	result, err := gr.db.Exec(query, id)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return models.ErrNotFound
	}

	query = `DELETE FROM items WHERE garden_id = ?`

	_, err = gr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (gr *GardenRepository) GetItemsForGarden(id string) ([]*models.Item, error) {
	query := `SELECT id, name, image, type, field1, field2, field3, field4, field5 FROM items WHERE garden_id = ?`

	rows, err := gr.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}

	items := []*models.Item{}
	for rows.Next() {
		var dbItem struct {
			*models.Item
			Type   int    `db:"type"`
			Field1 string `db:"field1"`
			Field2 string `db:"field2"`
			Field3 string `db:"field3"`
			Field4 string `db:"field4"`
			Field5 string `db:"field5"`
		}
		err = rows.StructScan(&dbItem)
		if err != nil {
			return nil, err
		}

		item := &models.Item{
			Id:    dbItem.Id,
			Name:  dbItem.Name,
			Image: dbItem.Image,
			Type:  models.ItemType(dbItem.Type),
		}
		fields := [5]string{
			dbItem.Field1,
			dbItem.Field2,
			dbItem.Field3,
			dbItem.Field4,
			dbItem.Field5,
		}
		item.Fields = fields

		items = append(items, item)
	}

	return items, nil
}

func (gr *GardenRepository) AssociateItem(g string, i *models.Item) (string, error) {
	if i == nil {
		return "", ErrNilItem
	}
	query := `UPDATE items SET garden_id = ? WHERE id = ?`

	_, err := gr.db.Exec(query, g, i.Id)

	return i.Id, err
}
