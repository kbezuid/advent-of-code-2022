package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const lsCommand = "$ ls"
const commandToken = "$"
const upCommand = "$ cd .."
const rootCommand = "$ cd /"
const cdCommand = "$ cd"

const diskSize = 70000000
const spaceNeeded = 30000000

type dir struct {
	name     string
	size     int
	children []*dir
	parent   *dir
}

func newDir(parent *dir, name string) dir {
	return dir{
		size:     0,
		name:     name,
		children: []*dir{},
		parent:   parent,
	}
}

func (d *dir) hasChild(childName string) bool {

	if d.children == nil {
		return false
	}

	for _, v := range d.children {
		if v.name == childName {
			return true
		}
	}

	return false
}

func (d *dir) getTotalDirSize() int {
	dirSize := d.size

	for _, c := range d.children {
		dirSize += c.getTotalDirSize()
	}

	return dirSize
}

func (d *dir) getChild(childName string) *dir {
	if d.children == nil {
		return nil
	}

	for _, v := range d.children {
		if v.name == childName {
			return v
		}
	}

	return nil
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lsStarted := false
	root := newDir(nil, "/")
	currentDir := &root
	currentDirSize := 0

	for scanner.Scan() {
		line := scanner.Text()

		if lsStarted && strings.HasPrefix(line, commandToken) {
			lsStarted = false
			currentDir.size = currentDirSize
			currentDirSize = 0
		}

		if strings.HasPrefix(line, cdCommand) {
			if line == upCommand {
				currentDir = currentDir.parent
				continue
			}

			if line != rootCommand {
				parts := strings.Split(line, " ")
				dirName := parts[2]

				if !currentDir.hasChild(dirName) {
					child := newDir(currentDir, dirName)
					currentDir.children = append(currentDir.children, &child)
					currentDir = &child
				} else {
					currentDir = currentDir.getChild(dirName)
				}
			}
		}

		if lsStarted {
			parts := strings.Split(line, " ")
			fileSize, err := strconv.Atoi(parts[0])

			if err == nil {
				currentDirSize += fileSize
			}
		}

		if line == lsCommand {
			lsStarted = true
		}
	}

	if lsStarted {
		lsStarted = false
		currentDir.size = currentDirSize
		currentDirSize = 0
	}

	freeSpace := diskSize - root.getTotalDirSize()
	neededSpace := spaceNeeded - freeSpace

	channels := []chan int{make(chan int), make(chan int)}
	done := []chan bool{make(chan bool), make(chan bool)}

	go partA(100000, channels[0], done[0])
	go partB(neededSpace, channels[1], done[1])

	root.process(channels)

	for _, c := range channels {
		close(c)
	}

	<-done[0]
	<-done[1]
	fmt.Println("Done")
}

func (d *dir) process(channels []chan int) {
	for _, c := range channels {
		c <- d.getTotalDirSize()
	}

	for _, child := range d.children {
		child.process(channels)
	}
}

func partA(maxSize int, data <-chan int, done chan<- bool) {
	totalSize := 0

	for v := range data {
		if v <= maxSize {
			totalSize += v
		}
	}

	fmt.Printf("Part A : %d\n", totalSize)
	close(done)
}

func partB(requiredSpace int, data <-chan int, done chan<- bool) {
	minSize := math.MaxInt

	for v := range data {
		if v >= requiredSpace && v < minSize {
			minSize = v
		}
	}

	fmt.Printf("Part B : %d\n", minSize)
	close(done)
}
