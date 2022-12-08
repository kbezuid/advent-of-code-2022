package treemap

import (
	"fmt"
)

type TreeMap struct {
	Width        int
	Height       int
	HeightMap    [][]int
	VisibleTrees map[Tree]bool
}

type Tree struct {
	X int
	Y int
	H int
}

func NewTree(x int, y int, h int) Tree {
	return Tree{
		X: x,
		Y: y,
		H: h,
	}
}

func New(width int, height int) TreeMap {

	heightMap := make([][]int, width)

	for i := 0; i < width; i++ {
		heightMap[i] = make([]int, height)
	}

	return TreeMap{
		Width:        width,
		Height:       height,
		HeightMap:    heightMap,
		VisibleTrees: map[Tree]bool{},
	}
}

func (t *TreeMap) SetHeight(x int, y int, height int) {
	t.HeightMap[x][y] = height
}

func (t *TreeMap) CountVisibleFromOutside() int {
	totalVisible := 0

	for x := 0; x < t.Width; x++ {
		for y := 0; y < t.Height; y++ {
			tree := NewTree(x, y, t.HeightMap[x][y])
			if x == 0 || y == 0 || x == t.Width-1 || y == t.Height-1 {
				totalVisible++
				t.VisibleTrees[tree] = true
				continue
			}

			if t.IsVisibleFromLeft(x, y) || t.IsVisibleFromTop(x, y) || t.IsVisibleFromRight(x, y) || t.IsVisibleFromBottom(x, y) {
				if !t.VisibleTrees[tree] {
					totalVisible++
					t.VisibleTrees[tree] = true
				}
			}
		}
	}

	return totalVisible
}

func (t *TreeMap) IsVisibleFromLeft(x int, y int) bool {
	var firstHeight = t.HeightMap[0][y]
	var treeHeight = t.HeightMap[x][y]

	if firstHeight >= treeHeight {
		return false
	}

	for ix := x - 1; ix > 0; ix-- {
		if t.HeightMap[ix][y] >= treeHeight {
			return false
		}
	}

	return true
}

func (t *TreeMap) IsVisibleFromRight(x int, y int) bool {
	var firstHeight = t.HeightMap[t.Width-1][y]
	var treeHeight = t.HeightMap[x][y]

	if firstHeight >= treeHeight {
		return false
	}

	for ix := x + 1; ix < t.Width-1; ix++ {
		if t.HeightMap[ix][y] >= treeHeight {
			return false
		}
	}

	return true
}

func (t *TreeMap) IsVisibleFromTop(x int, y int) bool {
	var firstHeight = t.HeightMap[x][0]
	var treeHeight = t.HeightMap[x][y]

	if firstHeight >= treeHeight {
		return false
	}

	for iy := y - 1; iy > 0; iy-- {
		if t.HeightMap[x][iy] >= treeHeight {
			return false
		}
	}

	return true
}

func (t *TreeMap) IsVisibleFromBottom(x int, y int) bool {
	var firstHeight = t.HeightMap[x][t.Height-1]
	var treeHeight = t.HeightMap[x][y]

	if firstHeight >= treeHeight {
		return false
	}

	for iy := y + 1; iy < t.Height-1; iy++ {
		if t.HeightMap[x][iy] >= treeHeight {
			return false
		}
	}

	return true
}

func (t *TreeMap) Print() {
	for x := 0; x < t.Width; x++ {
		for y := 0; y < t.Height; y++ {
			fmt.Printf("%d", t.HeightMap[x][y])
		}
		fmt.Println()
	}
}

func (t *Tree) Print() {
	fmt.Printf("(%d,%d,%d)\n", t.X, t.Y, t.H)
}

func (t *TreeMap) GetMaxScenicScore() int {
	maxScore := 0

	for x := 0; x < t.Width; x++ {
		for y := 0; y < t.Height; y++ {
			currentScore := t.ScenicScore(x, y)

			if currentScore > maxScore {
				maxScore = currentScore
			}
		}
	}

	return maxScore
}

func (t *TreeMap) ScenicScore(x int, y int) int {
	left, right, up, down := 0, 0, 0, 0

	height := t.HeightMap[x][y]

	for ix := x - 1; ix >= 0; ix-- {
		left++

		if t.HeightMap[ix][y] >= height {
			ix = -1
		}
	}

	for ix := x + 1; ix < t.Width; ix++ {
		right++

		if t.HeightMap[ix][y] >= height {
			ix = t.Width
		}
	}

	for iy := y - 1; iy >= 0; iy-- {
		up++

		if t.HeightMap[x][iy] >= height {
			iy = -1
		}
	}

	for iy := y + 1; iy < t.Width; iy++ {
		down++

		if t.HeightMap[x][iy] >= height {
			iy = t.Height
		}
	}

	return left * right * up * down
}
