package main

import (
	"testing"

	"github.com/matgd/advent2021/utils"
)

func TestTask1(t *testing.T) {
	exampleInput := utils.GetStringsFromInputFile("input_example.txt")

	expectedScore := 26397
	gotScore := task1(exampleInput)
	if expectedScore != gotScore {
		t.Log("[Task 1]: Expected ", expectedScore, "got", gotScore)
		t.Fail()
	}
}

func TestTask2(t *testing.T) {
	exampleInput := utils.GetStringsFromInputFile("input_example.txt")

	var expectedScore uint64 = 288957
	gotScore := task2(exampleInput)
	if expectedScore != gotScore {
		t.Log("[Task 2]: Expected ", expectedScore, "got", gotScore)
		t.Fail()
	}
}
