package models

func NewItem(n string, i *Image, f [5]*Field) *Item {
	return &Item{
		id:     genId(),
		Name:   n,
		Image:  i,
		Fields: f,
	}
}

type Item struct {
	id     string
	Name   string
	Image  *Image
	Fields [5]*Field // Only 5 for now
}

func (i *Item) Id() string {
	return i.id
}
