package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
