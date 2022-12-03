package rugsack

type Rugsack struct {
	items      map[byte]bool
	duplicates map[byte]bool
	Duplicates string
}

func New(compartment string) Rugsack {
	set := map[byte]bool{}

	for i := 0; i < len(compartment); i++ {
		set[compartment[i]] = true
	}

	return Rugsack{
		items:      set,
		Duplicates: "",
		duplicates: map[byte]bool{},
	}
}

func (r *Rugsack) SetDuplicates(compartment string) {
	for i := 0; i < len(compartment); i++ {
		item := compartment[i]

		_, hasItem := r.items[item]
		_, hasDuplicate := r.duplicates[item]

		if hasItem && !hasDuplicate {
			r.Duplicates += string(item)
			r.duplicates[item] = true
		}
	}
}

func (r Rugsack) GetTotalDuplicatePriority() int {
	total := 0

	for i := 0; i < len(r.Duplicates); i++ {
		item := r.Duplicates[i]
		itemValue := int(item)

		if item >= 'a' {
			total += itemValue - 96
		} else {
			total += itemValue - 38
		}
	}

	return total
}
