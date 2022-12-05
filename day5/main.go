package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: Maybe redo this challenge with just an array of strings this seems way to inefficient
const emptySlot = "   "

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	foundInstructions := false

	columns := make([]([]string), 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			foundInstructions = true
			continue
		}

		if line[0:3] == " 1 " {
			continue
		}

		if foundInstructions {
			printColumns(columns)

			parts := strings.Split(line, " ")

			count, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])

			from--
			to--

			fmt.Printf("%d: %d -> %d\n\n", count, from, to)

			if len(columns[from]) >= count {
				items := make([]string, 0)
				items = append(items, columns[from][0:count]...)

				if len(columns[from]) == count {
					columns[from] = []string{}
				} else {
					columns[from] = columns[from][count:]
				}

				columns[to] = append(items, columns[to]...)
			}

		} else {
			colIndex := 0
			for i := 0; i < len(line); i += 4 {
				colValue := line[i : i+3]
				hasValue := colValue != emptySlot

				if colIndex > len(columns)-1 {
					columns = append(columns, []string{})
				}

				if hasValue {
					columns[colIndex] = append(columns[colIndex], colValue)
				}

				colIndex++
			}
		}
	}

	fmt.Println("---------------")
	printColumns(columns)

	fmt.Print("Top Row : ")

	for i := 0; i < len(columns); i++ {
		if len(columns[i]) > 0 {
			val := columns[i][0]
			val = strings.ReplaceAll(val, "[", "")
			val = strings.ReplaceAll(val, "]", "")
			fmt.Print(val)
		}
	}

	// fmt.Printf("Total Contained %d\n", containedCount)
	fmt.Println("\n---------------")
}

func printColumns(columns [][]string) {
	maxLen := 0

	for i := 0; i < len(columns); i++ {
		if len(columns[i]) > maxLen {
			maxLen = len(columns[i])
		}
	}

	for row := maxLen; row > 0; row-- {
		for col := 0; col < len(columns); col++ {
			currentLen := len(columns[col])
			lenDiff := maxLen - currentLen

			if row <= currentLen {
				fmt.Print(columns[col][maxLen-row-lenDiff])
			} else {
				fmt.Print(emptySlot)
			}
		}

		fmt.Println()
	}
}
