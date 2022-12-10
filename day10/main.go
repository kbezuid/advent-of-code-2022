package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const noop = "noop"
const addx = "addx"
const screenW = 40
const spriteW = 3

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cycle := 0
	xValue := 1
	totalSignalStrength := 0
	currentPixel := 0

	cycleMap := map[string]int{
		noop: 1,
		addx: 2,
	}

	cyclesToCheck := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		cycles := cycleMap[parts[0]]

		for c := 0; c < cycles; c++ {
			cycle++

			if cyclesToCheck[cycle] {
				totalSignalStrength += cycle * xValue
			}

			drawScreen(currentPixel, xValue)
			currentPixel++

			if currentPixel == screenW {
				currentPixel = 0
			}
		}

		if strings.HasPrefix(line, addx) {
			dx, _ := strconv.Atoi(parts[1])
			xValue += dx
		}
	}

	fmt.Printf("Total Signal Strength %d", totalSignalStrength)
}

func getScreen(width int, spriteWidth int, x int) string {
	screen := ""
	spriteLeft := x - spriteWidth/2
	spriteRight := x + spriteWidth/2

	for i := 0; i < width; i++ {
		if i >= spriteLeft && i <= spriteRight {
			screen += "#"
		} else {
			screen += "."
		}
	}

	return screen
}

func drawScreen(pixel int, x int) {
	screen := getScreen(screenW, spriteW, x)

	if screen[pixel] == '#' {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if pixel == len(screen)-1 {
		fmt.Println()
	}
}
