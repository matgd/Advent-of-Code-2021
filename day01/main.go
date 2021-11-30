package main

import (
	"fmt"

	utils "github.com/matgd/advent2021/utils"
)

func task1() {
	counter := 0
	measurements := utils.GetIntegersFromInputFile("input.txt")
	measurementsShifted := measurements[1:]

	for i, valueB := range measurementsShifted {
		valueA := measurements[i]
		if valueB > valueA {
			counter++
		}
	}

	fmt.Println("[Task 1] Larger than previous count: ", counter)
}

func task2() {
	counter := 0
	measurementsA := utils.GetIntegersFromInputFile("input.txt")
	measurementsB := measurementsA[1:]
	measurementsC := measurementsB[1:]

	for i := 0; i < len(measurementsC)-1; i++ {
		windowSumA := measurementsA[i] + measurementsB[i] + measurementsC[i]
		j := i + 1
		windowSumB := measurementsA[j] + measurementsB[j] + measurementsC[j]

		if windowSumB > windowSumA {
			counter++
		}
	}

	fmt.Println("[Task 2] Larger than previous count: ", counter)
}

func main() {
	task1()
	task2()
}
