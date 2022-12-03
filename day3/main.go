package main

import (
	"bufio"
	"fmt"
	"github.com/kbezuid/day3/rugsack"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	totalPriorities := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		totalItems := len(line)

		compartmentA := line[0 : totalItems/2]
		compartmentB := line[totalItems/2 : totalItems]

		r := rugsack.New(compartmentA)
		r.SetDuplicates(compartmentB)

		currentTotalPriority := r.GetTotalDuplicatePriority()
		totalPriorities += currentTotalPriority

		fmt.Printf("Duplicates: %s\tBag Priority: %d\tTotal Priority: %d\n", r.Duplicates, currentTotalPriority, totalPriorities)
	}

	fmt.Println("---------------")
	fmt.Printf("Total Priority %d\n", totalPriorities)
	fmt.Println("---------------")
}
