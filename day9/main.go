package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	visited := map[image.Point]bool{
		{0, 0}: true,
	}

	directions := map[string]image.Point{
		"L": {-1, 0},
		"R": {1, 0},
		"U": {0, -1},
		"D": {0, 1},
	}

	knots := make([]image.Point, 10)
	tailIndex := len(knots) - 1

	for i := range knots {
		knots[i] = image.Point{0, 0}
	}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		dir := directions[parts[0]]
		count, _ := strconv.Atoi(parts[1])

		for i := 0; i < count; i++ {
			Move(&knots[0], dir)
			for j := 0; j < len(knots)-1; j++ {
				Update(&knots[j], &knots[j+1])
			}
			visited[knots[tailIndex]] = true
		}

	}

	fmt.Printf("%d\n", len(visited))
	fmt.Printf("Done")
}

func Move(p *image.Point, dir image.Point) {
	r := p.Add(dir)
	p.X = r.X
	p.Y = r.Y
}

func Update(h *image.Point, t *image.Point) {
	diff := h.Sub(*t)
	dist := Abs(diff)

	if dist > 1 {
		unit := Unit(*t, *h)
		newT := t.Add(unit)
		t.X = newT.X
		t.Y = newT.Y
	}
}

func Abs(p image.Point) int {
	a := math.Pow(float64(p.X), 2)
	b := math.Pow(float64(p.Y), 2)

	return int(math.Sqrt(a + b))
}

func Unit(from image.Point, to image.Point) image.Point {
	diff := to.Sub(from)

	size := float64(Abs(diff))
	dx := float64(diff.X) / size
	dy := float64(diff.Y) / size

	x := math.Round(dx)
	y := math.Round(dy)

	return image.Point{int(x), int(y)}
}
