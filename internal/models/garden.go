package models

import "strings"

func SearchGardens(gardens []*Garden, name string) []*Garden {
	found := []*Garden{}
	for _, garden := range gardens {
		gName := strings.ToLower(garden.Name)
		name = strings.ToLower(name)

		if strings.Contains(gName, name) {
			found = append(found, garden)
		}
	}
	return found
}

func NewGarden(n, l, d string, img string, i []*Item) *Garden {
	return &Garden{
		Id:          genId(),
		Name:        n,
		Location:    l,
		Description: d,
		Image:       img,
	}
}

type Garden struct {
	Id          string `db:"id"`
	Image       string `db:"image"`
	Name        string `db:"name"`
	Location    string `db:"location"`
	Description string `db:"description"`
}
