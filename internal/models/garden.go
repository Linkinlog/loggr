package models

func NewGarden(n, l, d string, i []*Item) *Garden {
	return &Garden{
		id:          genId(),
		Name:        n,
		Location:    l,
		Description: d,
		Inventory:   i,
	}
}

type Garden struct {
	id          string
	Name        string
	Location    string
	Description string
	Inventory   []*Item
}

func (g *Garden) Id() string {
	return g.id
}

func (g *Garden) AddItem(i *Item) {
	g.Inventory = append(g.Inventory, i)
}
