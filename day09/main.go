package main

import (
	"fmt"
	"strconv"

	"github.com/matgd/advent2021/utils"
)

func isLowestPoint(checkedPoint int, surroundingPoints ...int) bool {
	for _, p := range surroundingPoints {
		if p < checkedPoint {
			return false
		}
	}
	return true
}

func checkRow(heatmap [][]int, row int) []int {
	lowestPoints := make([]int, 0, len(heatmap[row]))
	maxHeight := 9

	checkTop := row != 0
	checkBottom := row != len(heatmap)-1

	var topLeft, top, topRight, left, right, bottomLeft, bottom, bottomRight int

	//[0]

	right = heatmap[row][1]
	if checkTop {
		top, topRight = heatmap[row-1][0], heatmap[row-1][1]
	} else {
		top, topRight = maxHeight, maxHeight
	}
	if checkBottom {
		bottom, bottomRight = heatmap[row+1][0], heatmap[row+1][1]
	} else {
		bottom, bottomRight = maxHeight, maxHeight
	}

	currentPoint := heatmap[row][0]
	surroundedLeft := make([]int, 0, 5)
	surroundedLeft = append(surroundedLeft, right)
	if checkTop {
		surroundedLeft = append(surroundedLeft, top)
		surroundedLeft = append(surroundedLeft, topRight)
	}
	if checkBottom {
		surroundedLeft = append(surroundedLeft, bottom)
		surroundedLeft = append(surroundedLeft, bottomRight)
	}

	if isLowestPoint(currentPoint, surroundedLeft...) {
		lowestPoints = append(lowestPoints, currentPoint)
	}

	// [1:-1]
	for x := 1; x < len(heatmap[row])-1; x++ {
		if checkTop {
			topLeft, top, topRight = heatmap[row-1][x-1], heatmap[row-1][x], heatmap[row-1][x+1]
		} else {
			topLeft, top, topRight = maxHeight, maxHeight, maxHeight
		}

		if checkBottom {
			bottomLeft, bottom, bottomRight = heatmap[row+1][x-1], heatmap[row+1][x], heatmap[row+1][x+1]
		} else {
			bottomLeft, bottom, bottomRight = maxHeight, maxHeight, maxHeight
		}
		left = heatmap[row][x-1]
		right = heatmap[row][x+1]

		surroundingPoints := []int{
			topLeft, top, topRight,
			left, right,
			bottomLeft, bottom, bottomRight,
		}

		currentPoint = heatmap[row][x]
		if isLowestPoint(currentPoint, surroundingPoints...) {
			lowestPoints = append(lowestPoints, currentPoint)
		}
	}

	//[-1]
	lastIndex := len(heatmap[row]) - 1
	left = heatmap[row][lastIndex-1]
	if checkTop {
		topLeft, top = heatmap[row-1][lastIndex-1], heatmap[row-1][lastIndex]
	} else {
		topLeft, top = maxHeight, maxHeight

	}
	if checkBottom {
		bottomLeft, bottom = heatmap[row+1][lastIndex-1], heatmap[row+1][lastIndex]
	} else {
		bottomLeft, bottom = maxHeight, maxHeight
	}

	currentPoint = heatmap[row][lastIndex]
	surroundedRight := make([]int, 0, 5)
	surroundedRight = append(surroundedRight, left)
	if checkTop {
		surroundedRight = append(surroundedRight, top)
		surroundedRight = append(surroundedRight, topLeft)
	}
	if checkBottom {
		surroundedRight = append(surroundedRight, bottom)
		surroundedRight = append(surroundedRight, bottomLeft)
	}

	if isLowestPoint(currentPoint, surroundedRight...) {
		lowestPoints = append(lowestPoints, currentPoint)
	}

	return lowestPoints
}

func task(heatmap []string) {
	heatmapMatrix := make([][]int, len(heatmap))
	for i, row := range heatmap {
		heatmapMatrix[i] = make([]int, len(row))
		for j, r := range row {
			number, _ := strconv.Atoi(string(r))
			heatmapMatrix[i][j] = number
		}
	}

	sum := 0
	for ri := range heatmapMatrix {
		for _, v := range checkRow(heatmapMatrix, ri) {
			sum += v + 1
		}
	}

	fmt.Println("Sum of risk levels:", sum)
}

func main() {
	// inputFile := "input_example.txt"
	inputFile := "input.txt"
	input := utils.GetStringsFromInputFile(inputFile)
	task(input)

	// Task 2 - BFS
}
