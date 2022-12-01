package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Open file error : %v", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfCount, maxCalories, maxElf, currentCalories = 1, 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
				maxElf = elfCount
			}

			log.Printf("Current Elf %d @ %d Calories. Max Elf %d @ %d Calories", elfCount, currentCalories, maxCalories, maxElf)
			elfCount++
			currentCalories = 0
		} else {
			calories, err := strconv.Atoi(line)

			if err != nil {
				log.Printf("Could not parse %s to integer. %v", line, err)
				continue
			}

			currentCalories += calories
		}
	}

	log.Printf("Winner is %d @ %d Calories", maxElf, maxCalories)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
	}
}
