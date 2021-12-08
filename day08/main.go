package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/matgd/advent2021/utils"
)

type stringByteSum int
type stringsLenSorted []string

func (s stringsLenSorted) Len() int {
	return len(s)
}

func (s stringsLenSorted) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s stringsLenSorted) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func calcStringByteSum(s string) stringByteSum {
	var sum stringByteSum = 0
	for _, r := range s {
		sum += stringByteSum(r)
	}
	return sum
}

func getDigitMapping(digits []string) map[stringByteSum]int {
	mapping := map[stringByteSum]int{}
	mappingDigit := map[int]string{}

	digitsLenSorted := stringsLenSorted(digits)
	sort.Sort(digitsLenSorted)

	// Order for finding 1, 7, 4, 8 first
	digitsLenSorted[9], digitsLenSorted[3] = digitsLenSorted[3], digitsLenSorted[9]
	digitsLenSorted[6], digitsLenSorted[9] = digitsLenSorted[9], digitsLenSorted[6]

	for _, d := range digitsLenSorted[:4] {
		dbyte := calcStringByteSum(d)
		switch len(d) {
		case 2:
			mapping[dbyte] = 1
			mappingDigit[1] = d
		case 3:
			mapping[dbyte] = 7
			mappingDigit[7] = d
		case 4:
			mapping[dbyte] = 4
			mappingDigit[4] = d
		case 7:
			mapping[dbyte] = 8
			mappingDigit[8] = d
		}
	}

	// find 6 -> 6 segments without one from 1
	for _, d := range digitsLenSorted[7:] { // 6 segments
		neededSeg := 2
		for _, r := range d {
			for _, rr := range mappingDigit[1] {
				if r == rr {
					neededSeg--
					if neededSeg == 0 {
						break
					}
				}
			}
		}
		if neededSeg == 1 {
			mapping[calcStringByteSum(d)] = 6
			mappingDigit[6] = d
		}
	}

	// find 9 -> includes segments from 4, 0 does not
	for _, d := range digitsLenSorted[7:] { // 6 segments
		if d == mappingDigit[6] {
			continue
		}

		neededSeg := 4
		for _, r := range d {
			for _, rr := range mappingDigit[4] {
				if r == rr {
					neededSeg--
					if neededSeg == 0 {
						mapping[calcStringByteSum(d)] = 9
						mappingDigit[9] = d
						break
					}
				}

			}
		}
	}

	// find 0 -> last one from 6-segmented
	for _, d := range digitsLenSorted[7:] { // 6 segments
		if d == mappingDigit[6] || d == mappingDigit[9] {
			continue
		}
		mapping[calcStringByteSum(d)] = 0
		mappingDigit[0] = d
	}

	// find 5 -> 6 includes all segments from 5
	for _, d := range digitsLenSorted[4:7] { // 5 segments
		neededSeg := 5
		for _, r := range d {
			for _, rr := range mappingDigit[6] {
				if r == rr {
					neededSeg--
					if neededSeg == 0 {
						mapping[calcStringByteSum(d)] = 5
						mappingDigit[5] = d
					}
				}
			}
		}
	}

	// find 2 -> 5 includes 3 segments from 2, 3 includes 4
	for _, d := range digitsLenSorted[4:7] { // 5 segments
		if d == mappingDigit[5] {
			continue
		}

		neededSeg := 5
		for _, r := range d {
			for _, rr := range mappingDigit[5] {
				if r == rr {
					neededSeg--
				}
			}
		}

		if neededSeg == 2 {
			mapping[calcStringByteSum(d)] = 2
			mappingDigit[2] = d
		} else if neededSeg == 1 {
			mapping[calcStringByteSum(d)] = 3
			mappingDigit[3] = d
		}
	}

	return mapping
}

func countFromMapping(readings []string, mapping map[stringByteSum]int) map[int]int {
	count := map[int]int{}
	for _, r := range readings {
		rbyte := calcStringByteSum(r)
		if v, exists := mapping[rbyte]; exists {
			count[v]++
		}
	}
	return count
}

func task1(digits, readings [][]string) {
	count := map[int]int{}
	for i := range digits {
		mapping := getDigitMapping(digits[i])
		counted := countFromMapping(readings[i], mapping)

		for k, v := range counted {
			count[k] += v
		}
	}

	fmt.Println("Sum of 1, 4, 7, 8:", count[1]+count[4]+count[7]+count[8])
}

func task2(digits, readings [][]string) {
	output := 0
	for i := range digits {
		mapping := getDigitMapping(digits[i])
		digitNumber := []int{
			mapping[calcStringByteSum(readings[i][0])],
			mapping[calcStringByteSum(readings[i][1])],
			mapping[calcStringByteSum(readings[i][2])],
			mapping[calcStringByteSum(readings[i][3])],
		}
		output += digitNumber[0]*1000 + digitNumber[1]*100 + digitNumber[2]*10 + digitNumber[3]
	}

	fmt.Println("Output:", output)
}

func main() {
	lines := utils.GetStringsFromInputFile("input.txt")
	digits, readings := make([][]string, len(lines)), make([][]string, len(lines))

	for i, l := range lines {
		split := strings.Split(l, " | ")

		digits[i] = make([]string, 10, 10)
		readings[i] = make([]string, 4)

		digits[i] = strings.Split(split[0], " ")
		readings[i] = strings.Split(split[1], " ")
	}

	task1(digits, readings)
	task2(digits, readings)
	fmt.Println("BORKED! Solution only works for example.")
}
