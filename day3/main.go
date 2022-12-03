package main

import (
	"bufio"
	"fmt"
	"github.com/kbezuid/day3/group"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	totalPriorities := 0
	scanner := bufio.NewScanner(file)
	lineCount := 0

	var g group.Group

	for scanner.Scan() {
		line := scanner.Text()

		if lineCount == 0 {
			g = group.New(line)
		}

		if lineCount == 1 {
			g.AddRugsackB(line)
		}

		if lineCount == 2 {
			g.AddRugsackC(line)
			currentTotalPriority := g.GetTotalPriority()
			totalPriorities += currentTotalPriority

			fmt.Printf("Badges: %s\tBadges Priority: %d\tTotal Priority: %d\n", g.Badges, currentTotalPriority, totalPriorities)
		}

		lineCount++

		if lineCount == 3 {
			lineCount = 0
		}
	}

	fmt.Println("---------------")
	fmt.Printf("Total Priority %d\n", totalPriorities)
	fmt.Println("---------------")
}
