package models

func NewItem(n string, i *Image, t ItemType, f [5]*Field) *Item {
	return &Item{
		id:     genId(),
		Name:   n,
		Image:  i,
		Type:   t,
		Fields: f,
	}
}

type ItemType int

const (
	_ ItemType = iota
	Plant
	Tool
	Seed
)

type Item struct {
	id     string
	Name   string
	Image  *Image
	Type   ItemType
	Fields [5]*Field // Only 5 for now
}

func (i *Item) Id() string {
	return i.id
}
