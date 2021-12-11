package main

import (
	"fmt"
	"time"

	"github.com/matgd/advent2021/utils"
)

const (
	maxNeighbours int = 8
	energyToFlash     = 10
)

type octopus struct {
	energy      int
	neighbours  []*octopus
	justFlashed bool
}

func newOctopus(energy int) octopus {
	return octopus{
		energy:      energy,
		neighbours:  make([]*octopus, 0, maxNeighbours),
		justFlashed: false,
	}
}

func toOctopuses(matrix [][]int) [][]octopus {
	octopuses := make([][]octopus, 0, len(matrix))

	for i, row := range matrix {
		octopuses = append(octopuses, make([]octopus, 0, len(row)))
		for _, n := range row {
			octopuses[i] = append(octopuses[i], newOctopus(n))
		}
	}

	return octopuses
}

func assignNeighbours(octopuses [][]octopus, row, column int) {
	neighboursCoords := map[string][]int{
		"topLeft":     {row - 1, column - 1},
		"top":         {row - 1, column},
		"topRight":    {row - 1, column + 1},
		"left":        {row, column - 1},
		"right":       {row, column + 1},
		"bottomLeft":  {row + 1, column - 1},
		"bottom":      {row + 1, column},
		"bottomRight": {row + 1, column + 1},
	}

	maxRowIndex := len(octopuses[0]) - 1
	minRowindex := 0
	maxColumnIndex := len(octopuses) - 1
	minColumnIndex := minRowindex
	for _, v := range neighboursCoords {
		r := v[0]
		c := v[1]
		if (r >= minRowindex) && (r <= maxRowIndex) && (c >= minColumnIndex) && (c <= maxColumnIndex) {
			octopuses[row][column].neighbours = append(octopuses[row][column].neighbours, &octopuses[r][c])
		}
	}
}

func (o *octopus) addEnergy(preventFlash bool) {
	if preventFlash {
		o.energy++
		return
	}

	if !o.justFlashed {
		o.energy++
		if o.energy >= energyToFlash {
			o.flash()
		}
	}
}

func (o *octopus) flash() {
	if !o.justFlashed {
		o.justFlashed = true
		o.energy = 0
		for i := range o.neighbours {
			o.neighbours[i].addEnergy(false)
		}
	}
}

func (o *octopus) flashIfYouCan() {
	if o.energy >= energyToFlash {
		o.flash()
	}
}

func printMatrix(octopuses [][]octopus) {
	for i := range octopuses {
		for j := range octopuses[i] {
			fmt.Print(octopuses[i][j].energy)
		}
		fmt.Println()
	}
	fmt.Println()
}

func task1(octopusMatrix [][]int) int {
	octopuses := toOctopuses(octopusMatrix)
	for i := range octopuses {
		for j := range octopuses[i] {
			assignNeighbours(octopuses, i, j)
		}
	}

	steps := 100
	flashes := 0
	for s := 0; s < steps; s++ {
		for i := range octopuses {
			for j := range octopuses[i] {
				octopuses[i][j].addEnergy(true)
			}
		}

		for i := range octopuses {
			for j := range octopuses {
				octopuses[i][j].flashIfYouCan()
			}
		}

		for i := range octopuses {
			for j := range octopuses[i] {
				if octopuses[i][j].justFlashed {
					flashes++
					octopuses[i][j].justFlashed = false
				}
			}
		}
	}

	return flashes
}

func task2(octopusMatrix [][]int) int {
	octopuses := toOctopuses(octopusMatrix)
	for i := range octopuses {
		for j := range octopuses[i] {
			assignNeighbours(octopuses, i, j)
		}
	}

	totalOctopuses := len(octopuses) * len(octopuses[0])
	steps := 0
	fullTotalGigaFlashHappened := false

	for !fullTotalGigaFlashHappened {
		octopusesFlashed := 0
		for i := range octopuses {
			for j := range octopuses[i] {
				octopuses[i][j].addEnergy(true)
			}
		}

		for i := range octopuses {
			for j := range octopuses {
				octopuses[i][j].flashIfYouCan()
			}
		}

		for i := range octopuses {
			for j := range octopuses[i] {
				if octopuses[i][j].justFlashed {
					octopusesFlashed++
					octopuses[i][j].justFlashed = false
				}
			}
		}
		steps++
		if octopusesFlashed == totalOctopuses {
			fullTotalGigaFlashHappened = true
		}
	}

	return steps
}

func main() {
	inputMatrix := utils.GetMatrixFromInputFile("input.txt")
	start := time.Now()
	t1 := task1(inputMatrix)
	end := time.Since(start)
	fmt.Println("Flashed", t1, "times.", "Executed in", end)

	start = time.Now()
	t2 := task2(inputMatrix)
	end = time.Since(start)
	fmt.Printf("Full total giga flash happened on step %v. Executed in %v.\n", t2, end)
}
