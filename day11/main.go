package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	number        int
	startingItems []int
	items         []int
	test          int
	operation     string
	pass          int
	fail          int
	worryLevel    int
}

const rounds = 20

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	monkeys := map[int]*Monkey{}
	var currentMonkey *Monkey

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Monkey") {
			parts := strings.Split(line, " ")
			parts[1] = strings.ReplaceAll(parts[1], ":", "")

			number, _ := strconv.Atoi(parts[1])

			m := newMonkey(number)
			currentMonkey = &m
			monkeys[currentMonkey.number] = currentMonkey
		}

		if strings.HasPrefix(line, "  Starting items: ") {
			parts := strings.Split(line, ",")
			parts[0] = strings.ReplaceAll(parts[0], "  Starting items: ", "")
			items := make([]int, len(parts))

			for i, v := range parts {
				items[i], _ = strconv.Atoi(strings.Trim(v, " "))
			}

			currentMonkey.setStartingItems(items)
		}

		if strings.HasPrefix(line, "  Operation:") {
			operation := strings.ReplaceAll(line, "  Operation: new = ", "")
			currentMonkey.setOperation(operation)
		}

		if strings.HasPrefix(line, "  Test: ") {
			test := strings.ReplaceAll(line, "  Test: divisible by ", "")
			currentMonkey.setTest(test)
		}

		if strings.HasPrefix(line, "    If true:") {
			passMonkey := strings.ReplaceAll(line, "    If true: throw to monkey ", "")
			currentMonkey.setPassMonkey(passMonkey)
		}

		if strings.HasPrefix(line, "    If false:") {
			failMonkey := strings.ReplaceAll(line, "    If false: throw to monkey ", "")
			currentMonkey.setFailMonkey(failMonkey)
		}
	}

	inspections := playRounds(rounds, monkeys)

	sort.Ints(inspections)
	monkeyBusiness := inspections[len(inspections)-1] * inspections[len(inspections)-2]

	fmt.Printf("Rounds %d : Monkey Business %d\n", rounds, monkeyBusiness)
	fmt.Println("Done")
}

func playRounds(rounds int, monkeys map[int]*Monkey) []int {

	inspections := make([]int, len(monkeys))

	for r := 0; r < rounds; r++ {
		for m := 0; m < len(monkeys); m++ {
			worryLevel := 0
			inspecting := len(monkeys[m].items) > 0
			currentItem := 0

			for inspecting {
				inspections[m] = inspections[m] + 1
				wl := float64(monkeys[m].items[currentItem])

				opParts := strings.Split(monkeys[m].operation, " ")

				left := wl

				if opParts[0] != "old" {
					left, _ = strconv.ParseFloat(opParts[0], 64)
				}

				right := wl

				if opParts[2] != "old" {
					right, _ = strconv.ParseFloat(opParts[2], 64)
				}

				switch opParts[1] {
				case "+":
					wl = left + right
				case "-":
					wl = left - right
				case "*":
					wl = left * right
				case "/":
					wl = left / right
				default:
					fmt.Printf("Unknown operation %s\n", opParts[1])

				}

				worryLevel = int(math.Floor(wl / 3))

				next := monkeys[m].fail

				if worryLevel%monkeys[m].test == 0 {
					next = monkeys[m].pass
				}

				throwItem(currentItem, worryLevel, monkeys[m], monkeys[next])

				if len(monkeys[m].items) == 0 {
					inspecting = false
				}
			}
		}
	}

	return inspections
}

func throwItem(item int, worryLevel int, from *Monkey, to *Monkey) {
	from.items[item] = worryLevel
	to.items = append(to.items, from.items[item])
	from.items = append(from.items[:item], from.items[item+1:]...)
}

func newMonkey(number int) Monkey {
	return Monkey{
		number:        number,
		items:         []int{},
		startingItems: []int{},
		test:          0,
		operation:     "",
		pass:          -1,
		fail:          -1,
		worryLevel:    0,
	}
}

func (m *Monkey) setPassMonkey(passMonkey string) {
	pass, _ := strconv.Atoi(passMonkey)
	m.pass = pass
}

func (m *Monkey) setFailMonkey(failMonkey string) {
	fail, _ := strconv.Atoi(failMonkey)
	m.fail = fail
}

func (m *Monkey) setTest(test string) {
	t, _ := strconv.Atoi(test)
	m.test = t
}

func (m *Monkey) setOperation(operation string) {
	m.operation = strings.Trim(operation, " ")
}

func (m *Monkey) setStartingItems(items []int) {
	m.items = items
	m.startingItems = items
}
