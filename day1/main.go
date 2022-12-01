package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/kbezuid/day1/structs"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Open file error : %v", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfCount, currentCalories = 1, 0
	var leaders = structs.NewLeaders(3)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {

			log.Println("--------------------------")
			log.Printf("Current elf : %d @ %d Calories\n", elfCount, currentCalories)
			log.Printf("Current Leaders\t\t %s\n", leaders.CalloriesToString())

			if currentCalories > leaders.MaxCalories() || currentCalories > leaders.MinCalories() {
				added, err := leaders.AddLeader(elfCount, currentCalories)

				if added {
					log.Printf("Updated Leaders\t\t %s\n", leaders.CalloriesToString())
				}

				if err != nil {
					log.Fatalf("Could not add leader : %v", err)
					os.Exit(1)
				}
			}
			log.Println("--------------------------")
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

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
	}

	log.Printf("Leaders are %s @ %s Calories", leaders.NumbersToString(), leaders.CalloriesToString())
	log.Printf("Total Calories %d", leaders.TotalCalories())
}
