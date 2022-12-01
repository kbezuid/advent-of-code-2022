package structs

import (
	"fmt"
)

type leaders struct {
	elves []elf
	Size  int
}

func NewLeaders(numberOfLeaders int) leaders {
	return leaders{
		elves: NewElves(numberOfLeaders, 0, 0),
		Size:  numberOfLeaders,
	}
}

func (l leaders) MaxCalories() int {
	return l.elves[0].calories
}

func (l leaders) MinCalories() int {
	return l.elves[l.Size-1].calories
}

func (l leaders) AddLeader(elfNumber int, calories int) (bool, error) {

	added := false
	var err error

	for i := 0; i < l.Size; i++ {
		if l.elves[i].calories < calories {
			l.shiftLeaders(i)
			l.elves[i] = NewElf(elfNumber, calories)
			added = true
			i = l.Size
		}
	}

	if !added {
		err = fmt.Errorf("could not add leader with %d calories. Max %d Min %d", calories, l.MaxCalories(), l.MinCalories())
	}

	return added, err
}

func (l leaders) TotalCalories() int {
	total := 0

	for i := 0; i < l.Size; i++ {
		total += l.elves[i].calories
	}

	return total
}

func (l leaders) CalloriesToString() string {
	calories := ""

	for i := 0; i < l.Size; i++ {
		calories += fmt.Sprintf("%d ", l.elves[i].calories)
	}

	return calories
}

func (l leaders) NumbersToString() string {
	numbers := ""

	for i := 0; i < l.Size; i++ {
		numbers += fmt.Sprintf("%d ", l.elves[i].number)
	}

	return numbers
}

func (l leaders) shiftLeaders(insertIndex int) {
	for i := l.Size - 1; i >= insertIndex && i > 0; i-- {
		l.elves[i] = l.elves[i-1]
	}
}
