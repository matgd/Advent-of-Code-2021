package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/matgd/advent2021/utils"
)

// Fold stores axis and value
type Fold struct {
	axis  string
	value int
}

func readInput(path string) (dots [][]int, folds []Fold, maxX int, maxY int) {
	input := utils.GetStringsFromInputFile(path)
	maxX = math.MinInt
	maxY = math.MinInt

	dots = make([][]int, 0, len(input))
	folds = make([]Fold, 0, len(input))

	processingDots := true
	for _, line := range input {
		if line == "" {
			processingDots = false
			continue
		}

		if processingDots {
			split := strings.Split(line, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}

			dots = append(dots, []int{x, y})
		} else {
			foldString := strings.Split(line, "fold along ")
			split := strings.Split(foldString[1], "=")
			val, _ := strconv.Atoi(split[1])

			folds = append(folds, Fold{axis: split[0], value: val})
		}
	}

	return
}

func printMatrix(matrix [][]uint8) {
	for i := range matrix {
		for j := range matrix[i] {
			v := matrix[i][j]
			if v > 0 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func countDots(matrix [][]uint8) int {
	dots := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > 0 {
				dots++
			}
		}
	}
	return dots
}

func task(firstTask bool) {
	dots, fold, maxX, maxY := readInput("input.txt")

	paperMatrix := make([][]uint8, maxY+1)
	for i := range paperMatrix {
		paperMatrix[i] = make([]uint8, maxX+1)
	}

	for _, dot := range dots {
		paperMatrix[dot[1]][dot[0]]++
	}

	var foldRange []Fold
	if firstTask {
		foldRange = fold[:1]
	} else {
		foldRange = fold
	}

	for _, f := range foldRange {
		if f.axis == "y" {
			cutOff := paperMatrix[f.value+1:]
			foldedArr := make([][]uint8, len(cutOff))
			foldedArrIndex := 0
			for i := len(cutOff) - 1; i >= 0; i-- {
				foldedArr[foldedArrIndex] = make([]uint8, len(cutOff[i]))
				for j := range cutOff[i] {
					foldedArr[foldedArrIndex][j] = cutOff[i][j]
				}
				foldedArrIndex++
			}

			// cut paperMatrix
			paperMatrix = paperMatrix[:f.value]

			// Apply from bottom
			for i := range foldedArr {
				for j := range foldedArr[i] {
					if foldedArr[i][j] > 0 {
						paperMatrix[f.value-len(foldedArr)+i][j]++
					}
				}
			}
		} else if f.axis == "x" {
			cutOff := make([][]uint8, len(paperMatrix))
			foldedArr := make([][]uint8, len(cutOff))

			for i := range paperMatrix {
				cutOff[i] = make([]uint8, f.value)
				foldedArr[i] = make([]uint8, len(paperMatrix[i])-f.value)

				cutOff[i] = paperMatrix[i][:f.value]
				foldedArr[i] = paperMatrix[i][f.value:]
			}

			// cut
			paperMatrix = cutOff

			for i := range foldedArr {
				for j := range foldedArr[i] {
					if foldedArr[i][j] > 0 {
						paperMatrix[i][len(paperMatrix[i])-j]++
					}
				}
			}

		}
	}

	if firstTask {
		fmt.Printf("[Task 1]: After 1 fold there are %v dots.\n", countDots(paperMatrix))
	} else {
		fmt.Println("[Task 2]:")
		printMatrix(paperMatrix)
	}
}

func main() {
	task(true)
	task(false)
}
