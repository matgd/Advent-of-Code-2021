package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/matgd/advent2021/utils"
)

// Directions
const (
	Forward string = "forward"
	Down           = "down"
	Up             = "up"
)

func task1() {
	positionY, positionZ := 0, 0

	commands := utils.GetStringsFromInputFile("input.txt")

	for _, command := range commands {
		directionValue := strings.Split(command, " ")
		direction := directionValue[0]
		value, err := strconv.Atoi(directionValue[1])

		if err != nil {
			log.Fatalf("Couldn't parse %s to int", directionValue[1])
		}

		switch direction {
		case Forward:
			positionY += value
		case Down:
			positionZ -= value
		case Up:
			positionZ += value
		}
	}
	fmt.Printf("[Task 1] Y: %d, Z: %d\n", positionY, positionZ)
	fmt.Printf("[Task 1] Result: %f\n", float64(positionY)*math.Abs(float64(positionZ)))
}

func task2() {
	positionY, positionZ, aim := 0, 0, 0

	commands := utils.GetStringsFromInputFile("input.txt")

	for _, command := range commands {
		directionValue := strings.Split(command, " ")
		direction := directionValue[0]
		value, err := strconv.Atoi(directionValue[1])

		if err != nil {
			log.Fatalf("Couldn't parse %s to int", directionValue[1])
		}

		switch direction {
		case Forward:
			positionY += value
			positionZ += value * aim
		case Down:
			aim -= value
		case Up:
			aim += value
		}
	}
	fmt.Printf("[Task 2] Y: %d, Z: %d, A: %d\n", positionY, positionZ, aim)
	fmt.Printf("[Task 2] Result: %f\n", float64(positionY)*math.Abs(float64(positionZ)))
}

func main() {
	task1()
	task2()
}
