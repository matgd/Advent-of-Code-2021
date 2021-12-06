package main

import (
	"fmt"
	"time"

	"github.com/matgd/advent2021/utils"
)

func task(days int, initialFish []int) {
	var fish [9]*uint64 // Group fish by value

	// Create pointers to zero values
	for i := 0; i < len(fish); i++ {
		fish[i] = new(uint64)
	}

	for _, f := range initialFish {
		*fish[f]++
	}

	for currentDay := 0; currentDay < days; currentDay++ {
		resetFish := *fish[0]

		fish[0] = fish[1]
		fish[1] = fish[2]
		fish[2] = fish[3]
		fish[3] = fish[4]
		fish[4] = fish[5]
		fish[5] = fish[6]

		fish[6] = fish[7]
		*fish[6] += resetFish

		fish[7] = fish[8]

		fish[8] = new(uint64)
		*fish[8] = resetFish
	}

	var totalFish uint64 = 0
	for i := range fish {
		totalFish += *fish[i]
	}
	fmt.Println(days, "days: There are", totalFish, "fish.")
}

func main() {
	initialFish := utils.GetIntegersFromCSV("input.txt")

	start := time.Now()
	task(80, initialFish)
	fmt.Println("Executed in", time.Since(start))

	start = time.Now()
	task(256, initialFish)
	fmt.Println("Executed in", time.Since(start))
}
