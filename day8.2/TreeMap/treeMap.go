package treemap

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type TreeMap struct {
	Width        int
	Height       int
	HeightMap    [][]int
	VisibleTrees map[Tree]bool
	BestTree     Tree
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
				t.BestTree = NewTree(x, y, currentScore)
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

func (t *TreeMap) ToImage(imageName string, imagePixelScale int) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{imagePixelScale * t.Width, imagePixelScale * t.Height}})

	for x := 0; x < t.Width; x++ {
		for y := 0; y < t.Height; y++ {
			r, g, b, a := linearGradient(float64(t.HeightMap[x][y]))
			color := color.RGBA{r, g, b, a}
			for i := 0; i < imagePixelScale; i++ {
				for j := 0; j < imagePixelScale; j++ {
					img.Set(x*imagePixelScale+i, y*imagePixelScale+j, color)
				}
			}
		}
	}
	f, _ := os.Create(imageName)
	png.Encode(f, img)
}

func (t *TreeMap) ToVisibleImage(imageName string, imagePixelScale int) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{imagePixelScale * t.Width, imagePixelScale * t.Height}})

	for tree, _ := range t.VisibleTrees {
		color := color.RGBA{255, 0, 0, 0xFF}
		for i := 0; i < imagePixelScale; i++ {
			for j := 0; j < imagePixelScale; j++ {
				img.Set(tree.X*imagePixelScale+i, tree.Y*imagePixelScale+j, color)
			}
		}
	}

	f, _ := os.Create(imageName)
	png.Encode(f, img)
}

func (t *TreeMap) ToBestSpotImage(imageName string, imagePixelScale int) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{imagePixelScale * t.Width, imagePixelScale * t.Height}})

	color := color.RGBA{0, 0xFF, 0, 0xFF}

	for i := 0; i < imagePixelScale; i++ {
		for j := 0; j < imagePixelScale; j++ {
			img.Set(t.BestTree.X*imagePixelScale+i, t.BestTree.Y*imagePixelScale+j, color)
		}
	}

	f, _ := os.Create(imageName)
	png.Encode(f, img)
}

func linearGradient(x float64) (uint8, uint8, uint8, uint8) {
	colorA := []float64{0x00, 0x00, 0x00}
	colorB := []float64{0xD1, 0x00, 0xFF}

	d := x / 9
	r := colorA[0] + d*(colorB[0]-colorA[0])
	g := colorA[1] + d*(colorB[1]-colorA[1])
	b := colorA[2] + d*(colorB[2]-colorA[2])
	return uint8(r), uint8(g), uint8(b), 0xFF
}
