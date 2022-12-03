package group

type Group struct {
	bags   bags
	badges map[byte]bool
	Badges string
}

type bags struct {
	itemsA map[byte]bool
	itemsB map[byte]bool
}

func New(rugsackA string) Group {
	g := Group{
		Badges: "",
		badges: map[byte]bool{},
		bags:   newBags(),
	}

	setBag(rugsackA, g.bags.itemsA)

	return g
}

func (g *Group) AddRugsackB(rugsackB string) {
	setBag(rugsackB, g.bags.itemsB)
}

func (g *Group) AddRugsackC(rugsackC string) {
	for i := 0; i < len(rugsackC); i++ {
		item := rugsackC[i]
		_, hasInA := g.bags.itemsA[item]
		_, hasInB := g.bags.itemsB[item]

		if hasInA && hasInB {
			_, hasBadge := g.badges[item]

			if !hasBadge {
				g.badges[item] = true
				g.Badges += string(item)
			}
		}
	}
}

func newBags() bags {
	return bags{
		itemsA: map[byte]bool{},
		itemsB: map[byte]bool{},
	}
}

func setBag(rugsack string, bag map[byte]bool) {
	for i := 0; i < len(rugsack); i++ {
		item := rugsack[i]

		_, hasItem := bag[item]

		if !hasItem {
			bag[item] = true
		}
	}
}

func (g *Group) GetTotalPriority() int {
	total := 0

	for i := 0; i < len(g.Badges); i++ {
		item := g.Badges[i]
		itemValue := int(item)

		if item >= 'a' {
			total += itemValue - 96
		} else {
			total += itemValue - 38
		}
	}

	return total
}
