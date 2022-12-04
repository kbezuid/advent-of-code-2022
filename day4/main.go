package main

import (
	"bufio"
	"fmt"
	"github.com/kbezuid/day4/sections"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	containedCount := 0
	overlapCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		rangeA := strings.Split(assignments[0], "-")
		rangeB := strings.Split(assignments[1], "-")

		s := sections.New(getMinMax(rangeA[0], rangeA[1]))

		min, max := getMinMax(rangeB[0], rangeB[1])

		if s.Contains(min, max) {
			containedCount++
			fmt.Printf("Containment found : %s \t %s\n", rangeA, rangeB)
		}

		if s.Overlaps(min, max) {
			overlapCount++
			fmt.Printf("Overlap found : %s \t %s\n", rangeA, rangeB)
		}
	}

	fmt.Println("---------------")
	fmt.Printf("Total Contained %d\n", containedCount)
	fmt.Printf("Total Overlap %d\n", overlapCount)
	fmt.Println("---------------")
}

func getMinMax(min string, max string) (int, int) {
	a, _ := strconv.Atoi(min)
	b, _ := strconv.Atoi(max)

	return a, b
}
