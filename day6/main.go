package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			buffer := line[i : i+4]

			if IsUnique(buffer) {
				fmt.Printf("Packet Starts @ %d : %s\n", i+4, buffer)
				i = len(line)
			}
		}
	}
}

func IsUnique(items string) bool {
	values := map[rune]bool{}

	for _, v := range items {
		if values[v] {
			return false
		}

		values[v] = true
	}

	return true
}
