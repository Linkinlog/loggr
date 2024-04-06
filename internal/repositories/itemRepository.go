package repositories

import (
	"github.com/Linkinlog/loggr/internal/models"
	"github.com/Linkinlog/loggr/internal/stores"
)

func NewItemRepository(db *stores.SqliteStore) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

type ItemRepository struct {
	db *stores.SqliteStore
}

func (ir *ItemRepository) Add(i *models.Item) (string, error) {
	if i == nil {
		return "", ErrNilItem
	}
	query := `INSERT INTO items (id, garden_id, name, image, type, field1, field2, field3, field4, field5) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := ir.db.Exec(query, i.Id, 0, i.Name, i.Image, int(i.Type), i.Fields[0], i.Fields[1], i.Fields[2], i.Fields[3], i.Fields[4])

	return i.Id, err
}

func (ir *ItemRepository) Get(id string) (*models.Item, error) {
	query := `SELECT id, name, image, type, field1, field2, field3, field4, field5 FROM items WHERE id = ?`

	var dbModel struct {
		*models.Item
		Type   int    `db:"type"`
		Field1 string `db:"field1"`
		Field2 string `db:"field2"`
		Field3 string `db:"field3"`
		Field4 string `db:"field4"`
		Field5 string `db:"field5"`
	}
	err := ir.db.Get(&dbModel, query, id)
	if err != nil {
		return nil, err
	}

	i := dbModel.Item
	i.Type = models.ItemType(dbModel.Type)
	i.Fields = [5]string{dbModel.Field1, dbModel.Field2, dbModel.Field3, dbModel.Field4, dbModel.Field5}

	return i, nil
}

func (ir *ItemRepository) Update(i *models.Item) error {
	if i == nil {
		return ErrNilItem
	}
	query := `UPDATE items SET name = ?, image = ?, type = ?, field1 = ?, field2 = ?, field3 = ?, field4 = ?, field5 = ? WHERE id = ?`

	_, err := ir.db.Exec(query, i.Name, i.Image, i.Type, i.Fields[0], i.Fields[1], i.Fields[2], i.Fields[3], i.Fields[4], i.Id)

	return err
}

func (ir *ItemRepository) Delete(id string) error {
	query := `DELETE FROM items WHERE id = ?`

	_, err := ir.db.Exec(query, id)

	return err
}
