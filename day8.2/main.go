package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	treemap "github.com/kbezuid/day8.2/TreeMap"
)

const w = 99
const h = 99
const imagePixelScale = 10

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	treeMap := treemap.New(w, h)
	scanner := bufio.NewScanner(file)
	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		trees := strings.Split(line, "")

		for col, t := range trees {
			h, _ := strconv.Atoi(t)
			treeMap.SetHeight(col, row, h)
		}
		row++
	}

	fmt.Printf("Visible Trees %d\n", treeMap.CountVisibleFromOutside())
	fmt.Printf("Sceninc Score %d\n", treeMap.GetMaxScenicScore())

	treeMap.ToImage("Raw.png", imagePixelScale)
	treeMap.ToVisibleImage("Visible.png", imagePixelScale)
	treeMap.ToBestSpotImage("Best.png", imagePixelScale)

	fmt.Printf("Done")
}
