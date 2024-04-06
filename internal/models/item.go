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

func Plants(items []*Item) []*Item {
	plants := []*Item{}
	for _, item := range items {
		if item.Type == Plant {
			plants = append(plants, item)
		}
	}
	return plants
}

func NewItem(n string, i string, t ItemType, f [5]string) *Item {
	return &Item{
		Id:     genId(),
		Name:   n,
		Image:  i,
		Type:   t,
		Fields: f,
	}
}

type Item struct {
	Id     string `db:"id"`
	Name   string `db:"name"`
	Image  string `db:"image"`
	Type   ItemType
	Fields [5]string // Only 5 for now
}
