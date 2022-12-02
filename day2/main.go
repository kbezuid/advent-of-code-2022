package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kbezuid/day2/game"
)

func main() {
	game := game.NewGame()

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Open file error : %v", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		game.PlayRound(line[0], line[2])

	}

	fmt.Println("---------------")
	fmt.Printf("Total Score %d\n", game.TotalScore)
	fmt.Println("---------------")
}
