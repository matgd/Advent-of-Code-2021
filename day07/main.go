package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	"github.com/matgd/advent2021/utils"
)

type fuelCalcFunc func([]int, int) int

func calcFuel(crabs []int, selectedCrab int) int {
	fuel := 0
	for _, crab := range crabs {
		fuel += int(math.Abs(float64(crab) - float64(selectedCrab)))
	}

	return fuel
}

func calcFuelExpensive(crabs []int, selectedCrab int) int {
	fuel := 0
	for _, crab := range crabs {
		distance := int(math.Abs(float64(crab) - float64(selectedCrab)))
		fuel += distance * (distance + 1) / 2
	}

	return fuel
}

func getLeastFuel(crabs, positions []int, fn fuelCalcFunc, leastFuelCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	leastFuel := math.MaxInt
	for _, p := range positions {
		if f := fn(crabs, p); f < leastFuel {
			leastFuel = f
		}
	}
	leastFuelCh <- leastFuel
}

func task(crabsPos []int, fn fuelCalcFunc) {
	min := math.MaxInt
	max := math.MinInt

	for _, crab := range crabsPos {
		if crab < min {
			min = crab
		}
		if crab > max {
			max = crab
		}

	}

	crabPositions := make([]int, 0, len(crabsPos))
	for i := min; i <= max; i++ {
		crabPositions = append(crabPositions, i)
	}

	cpuCores := runtime.NumCPU()
	if cpuCores > len(crabPositions) {
		cpuCores = len(crabPositions)
	}

	fuelCh := make(chan int, cpuCores)
	wg := sync.WaitGroup{}
	wg.Add(cpuCores)
	for core := 0; core < cpuCores; core++ {
		from := (len(crabPositions) / cpuCores) * core
		to := (len(crabPositions) / cpuCores) * (core + 1)
		go getLeastFuel(crabsPos, crabPositions[from:to], fn, fuelCh, &wg)
	}
	wg.Wait()

	minRes := math.MaxInt
	for ch := 0; ch < cpuCores; ch++ {
		if res := <-fuelCh; res < minRes {
			minRes = res
		}
	}

	fmt.Println("Least fuel:", minRes)
}

func main() {
	input := utils.GetIntegersFromCSV("input.txt")
	start := time.Now()
	task(input, calcFuel)
	fmt.Println("Executed in:", time.Since(start))
	start = time.Now()
	task(input, calcFuelExpensive)
	fmt.Println("Executed in:", time.Since(start))
}
