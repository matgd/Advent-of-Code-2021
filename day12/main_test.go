package main

import (
	"testing"

	"github.com/matgd/advent2021/utils"
)

func TestTask1(t *testing.T) {
	exampleInput := utils.GetStringsFromInputFile("input_example.txt")

	expected := 10
	got := task1(exampleInput)
	if expected != got {
		t.Log("[Task 1]: Expected ", expected, "got", got)
		t.Fail()
	}
}
