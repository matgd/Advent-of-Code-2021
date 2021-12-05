package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matgd/advent2021/utils"
)

type coordinate struct {
	x, y int
}

func loadCoordinates() [][]coordinate {
	input := utils.GetStringsFromInputFile("input.txt")
	coordsStrings := make([][]string, len(input))
	coords := make([][]coordinate, len(input))

	for i, line := range input {
		split := strings.Split(line, " ")
		start, end := split[0], split[len(split)-1]
		coordsStrings[i] = []string{start, end}
	}

	for i, coordsStringsSet := range coordsStrings {
		setA := strings.Split(coordsStringsSet[0], ",")
		setB := strings.Split(coordsStringsSet[1], ",")
		setAX, _ := strconv.Atoi(setA[0])
		setAY, _ := strconv.Atoi(setA[1])
		setBX, _ := strconv.Atoi(setB[0])
		setBY, _ := strconv.Atoi(setB[1])
		coords[i] = []coordinate{
			{x: setAX, y: setAY},
			{x: setBX, y: setBY},
		}
	}

	return coords
}

func isHorizontal(coordsA, coordsB coordinate) bool {
	return coordsA.y == coordsB.y
}

func isVertical(coordsA, coordsB coordinate) bool {
	return coordsA.x == coordsB.x
}

func isDiagonalSlash(coordsA, coordsB coordinate) bool {
	if coordsA.x > coordsB.x {
		coordsA, coordsB = coordsB, coordsA
	}
	return coordsB.x-coordsA.x == coordsB.y-coordsA.y
}

func isDiagonalBackslash(coordsA, coordsB coordinate) bool {
	if coordsA.x > coordsB.x {
		coordsA, coordsB = coordsB, coordsA
	}
	return coordsB.x-coordsA.x == coordsA.y-coordsB.y
}

func mark(hitMap map[int]map[int]int, c coordinate) {
	if _, exists := hitMap[c.x]; !exists {
		hitMap[c.x] = map[int]int{}
	}
	hitMap[c.x][c.y]++
}

func task1() {
	coords := loadCoordinates()
	hitMap := map[int]map[int]int{}

	for _, crd := range coords {
		start, end := crd[0], crd[1]
		if isVertical(start, end) {
			x := start.x
			if end.y < start.y {
				start, end = end, start
			}

			for y := start.y; y <= end.y; y++ {
				mark(hitMap, coordinate{x, y})
			}
		} else if isHorizontal(start, end) {
			y := start.y
			if end.x < start.x {
				start, end = end, start
			}
			for x := start.x; x <= end.x; x++ {
				mark(hitMap, coordinate{x, y})
			}
		}
	}
	moreThan2 := 0
	for _, vertical := range hitMap {
		for _, point := range vertical {
			if point >= 2 {
				moreThan2++
			}
		}
	}
	fmt.Println("[Task 1]:", moreThan2)
}

func task2() {
	coords := loadCoordinates()
	hitMap := map[int]map[int]int{}

	for _, crd := range coords {
		start, end := crd[0], crd[1]
		if isVertical(start, end) {
			x := start.x
			if end.y < start.y {
				start, end = end, start
			}

			for y := start.y; y <= end.y; y++ {
				mark(hitMap, coordinate{x, y})
			}
		} else if isHorizontal(start, end) {
			y := start.y
			if end.x < start.x {
				start, end = end, start
			}
			for x := start.x; x <= end.x; x++ {
				mark(hitMap, coordinate{x, y})
			}
		} else if isDiagonalSlash(start, end) {
			if start.x > end.x {
				start, end = end, start
			}
			y := start.y
			for x := start.x; x <= end.x; x++ {
				mark(hitMap, coordinate{x, y})
				y++
			}
		} else if isDiagonalBackslash(start, end) {
			fmt.Println(start, end)
			if start.y > end.y {
				start, end = end, start
			}
			x := start.x
			for y := start.y; y <= end.y; y++ {
				mark(hitMap, coordinate{x, y})
				x--
			}
		}
	}
	moreThan2 := 0
	for _, vertical := range hitMap {
		for _, point := range vertical {
			if point >= 2 {
				moreThan2++
			}
		}
	}
	fmt.Println("[Task 2]:", moreThan2)
}
func main() {
	task1()
	task2()
}
