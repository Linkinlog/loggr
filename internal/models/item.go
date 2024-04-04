package models

import "strings"

func SearchItems(items []*Item, name string) []*Item {
	found := []*Item{}
	for _, item := range items {
		iName := strings.ToLower(item.Name)
		name = strings.ToLower(name)
		if strings.Contains(iName, name) {
			found = append(found, item)
		}
	}
	return found
}

func NewItem(n string, i *Image, t ItemType, f [5]*Field) *Item {
	return &Item{
		id:     genId(),
		Name:   n,
		Image:  i,
		Type:   t,
		Fields: f,
	}
}

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
