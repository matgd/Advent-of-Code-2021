package main

import (
	"testing"

	"github.com/matgd/advent2021/utils"
)

func TestTask1(t *testing.T) {
	exampleInput := utils.GetMatrixFromInputFile("input_example.txt")

	expected := 1656
	got := task1(exampleInput)
	if expected != got {
		t.Log("[Task 1]: Expected ", expected, "got", got)
		t.Fail()
	}
}

func TestTask2(t *testing.T) {
	exampleInput := utils.GetMatrixFromInputFile("input_example.txt")

	expected := 195
	got := task2(exampleInput)
	if expected != got {
		t.Log("[Task 2]: Expected ", expected, "got", got)
		t.Fail()
	}
}
