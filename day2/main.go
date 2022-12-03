package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kbezuid/day2/game"
)

func main() {
	os.Exit(run())
}

func run() int {
	g := game.NewGame()

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Open file error : %v", err)
		return 1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		g.PlayRoundWithOutcome(line[0], line[2])
	}

	fmt.Println("---------------")
	fmt.Printf("Total Score %d\n", g.TotalScore)
	fmt.Println("---------------")

	return 0
}
