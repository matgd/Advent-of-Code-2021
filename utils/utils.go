package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInputFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Couldn't open file '%s' due to error: %s\n", path, err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Couldn't parse contents of file '%s' due to error: %s\n", path, scanner.Err())
	}
	return lines
}

// GetStringsFromInputFile returns input from file in from of strings in array
func GetStringsFromInputFile(path string) []string {
	return getInputFromFile(path)
}

// GetIntegersFromInputFile returns input from file in from of integers in array
func GetIntegersFromInputFile(path string) []int {
	var lines []int
	for _, v := range getInputFromFile(path) {
		integer, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Couldn't parse contents of file '%s' due to error: %s\n", path, err)
		}
		lines = append(lines, integer)
	}
	return lines
}

// GetIntegersFromCSV returns array of integers from CSV file
func GetIntegersFromCSV(path string) []int {
	var numbers []int
	for _, line := range getInputFromFile(path) {
		stringNumbers := strings.Split(line, ",")
		for _, n := range stringNumbers {
			nInt, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("Couldn't parse contents of file '%s' due to error: %s\n", path, err)
			}
			numbers = append(numbers, nInt)
		}
	}
	return numbers
}

// GetMatrixFromInputFile returns array of arrays from input file
// which includes one-digit numbers, non-seperated
// in multiple lines
func GetMatrixFromInputFile(path string) [][]int {
	matrix := make([][]int, 0, 10)
	for i, line := range getInputFromFile(path) {
		matrix = append(matrix, []int{})
		for _, n := range line {
			nInt, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatalf("Couldn't parse contents of file '%s' due to error: %s\n", path, err)
			}
			matrix[i] = append(matrix[i], nInt)
		}
	}
	return matrix
}

// GetFloatsFromInputFile returns input from file in from of integers in array
func GetFloatsFromInputFile(path string) []float64 {
	var lines []float64
	for _, v := range getInputFromFile(path) {
		floatNumber, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatalf("Couldn't parse contents of file '%s' due to error: %s\n", path, err)
		}
		lines = append(lines, floatNumber)
	}
	return lines
}

// AllTrueArray checks if all values in array are true
func AllTrueArray(arr []bool) bool {
	for _, value := range arr {
		if !value {
			return false
		}
	}
	return true
}
