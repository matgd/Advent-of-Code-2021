package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matgd/advent2021/utils"
)

func task1() {
	reportNumbers := utils.GetStringsFromInputFile("input.txt")
	count := map[int]int{}

	for _, binaryNumber := range reportNumbers {
		for i, char := range binaryNumber {
			if bit, _ := strconv.ParseBool(string(char)); bit {
				count[i]++
			}
		}
	}

	gammaResultBuilder := strings.Builder{}
	epsilonResultBuilder := strings.Builder{}
	for i := 0; i < len(count); i++ {
		if count[i] > len(reportNumbers)/2 {
			gammaResultBuilder.WriteRune('1')
			epsilonResultBuilder.WriteRune('0')
		} else {
			gammaResultBuilder.WriteRune('0')
			epsilonResultBuilder.WriteRune('1')
		}
	}

	gamma := gammaResultBuilder.String()
	epsilon := epsilonResultBuilder.String()

	fmt.Println("[Task 1]: Gamma:", gamma, " Epsilon:", epsilon)
	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println("[Task 1]: Result:", gammaDecimal*epsilonDecimal)
}

type device string

const (
	oxygenGenerator device = "oxygenGenerator"
	co2Scrubber            = "co2Scrubber"
)

func getRating(reports []string, deviceType device, checkIndex int, ch chan string) {
	if len(reports) <= 1 {
		ch <- reports[0]
		return
	}

	filteredReportsZeros := make([]string, len(reports))
	zeros := 0

	filteredReportsOnes := make([]string, len(reports))
	ones := 0

	for _, binaryNumber := range reports {
		if bit, _ := strconv.ParseBool(string(binaryNumber[checkIndex])); bit {
			filteredReportsOnes[ones] = binaryNumber
			ones++
		} else {
			filteredReportsZeros[zeros] = binaryNumber
			zeros++
		}
	}

	checkIndex++
	if deviceType == oxygenGenerator {
		if ones >= zeros {
			getRating(filteredReportsOnes[:ones], deviceType, checkIndex, ch)
		} else {
			getRating(filteredReportsZeros[:zeros], deviceType, checkIndex, ch)
		}
	}
	if deviceType == co2Scrubber {
		if ones >= zeros {
			getRating(filteredReportsZeros[:zeros], deviceType, checkIndex, ch)
		} else {
			getRating(filteredReportsOnes[:ones], deviceType, checkIndex, ch)
		}
	}
}

func task2() {
	reportNumbers := utils.GetStringsFromInputFile("input.txt")

	reportNumbersGenerator := make([]string, len(reportNumbers))
	reportNumbersGeneratorNextIndex := 0

	reportNumbersScrubber := make([]string, len(reportNumbers))
	reportNumbersScrubberNextIndex := 0

	// First partitioning
	for _, binaryNumber := range reportNumbers {
		if firstBit, _ := strconv.ParseBool(string(binaryNumber[0])); firstBit {
			reportNumbersGenerator[reportNumbersGeneratorNextIndex] = binaryNumber
			reportNumbersGeneratorNextIndex++
		} else {
			reportNumbersScrubber[reportNumbersScrubberNextIndex] = binaryNumber
			reportNumbersScrubberNextIndex++
		}
	}

	generatorChannel := make(chan string, 1)
	scrubberChannel := make(chan string, 1)

	go getRating(reportNumbersGenerator[:reportNumbersGeneratorNextIndex], oxygenGenerator, 1, generatorChannel)
	go getRating(reportNumbersScrubber[:reportNumbersScrubberNextIndex], co2Scrubber, 1, scrubberChannel)

	generatorRating := <-generatorChannel
	scrubberRating := <-scrubberChannel

	fmt.Println("[Task 2]: Oxygen:", generatorRating, " CO2:", scrubberRating)
	oxygen, _ := strconv.ParseInt(generatorRating, 2, 64)
	co2, _ := strconv.ParseInt(scrubberRating, 2, 64)

	fmt.Println("[Task 2]: Result: ", oxygen*co2)
}

func main() {
	task1()
	task2()
}
