package structs

type elf struct {
	number   int
	calories int
}

func NewElf(number int, calories int) elf {
	return elf{
		number,
		calories,
	}
}

func NewElves(numberOfElves int, initialNumber int, initialCalories int) []elf {
	var elves = make([]elf, numberOfElves)

	for i := 0; i < numberOfElves; i++ {
		elves[i] = NewElf(0, -1*i)
	}

	return elves
}
