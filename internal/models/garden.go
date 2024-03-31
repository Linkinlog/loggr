package models

func NewGarden(n, l, d string, img *Image, i []*Item) *Garden {
	return &Garden{
		id:          genId(),
		Name:        n,
		Location:    l,
		Description: d,
		Image:       img,
		Inventory:   i,
	}
}

type Garden struct {
	id          string
	Image       *Image
	Name        string
	Location    string
	Description string
	Inventory   []*Item
}

func (g *Garden) Id() string {
	return g.id
}

func (g *Garden) Plants() []*Item {
	var plants []*Item
	for _, i := range g.Inventory {
		if i.Type == Plant {
			plants = append(plants, i)
		}
	}
	return plants
}

func (g *Garden) Tools() []*Item {
	var tools []*Item
	for _, i := range g.Inventory {
		if i.Type == Tool {
			tools = append(tools, i)
		}
	}
	return tools
}

func (g *Garden) Seeds() []*Item {
	var seeds []*Item
	for _, i := range g.Inventory {
		if i.Type == Seed {
			seeds = append(seeds, i)
		}
	}
	return seeds
}

func (g *Garden) AddItem(i *Item) {
	g.Inventory = append(g.Inventory, i)
}
