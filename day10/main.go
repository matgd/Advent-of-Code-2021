package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/matgd/advent2021/utils"
)

var lineScoring map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var counterParts map[rune]rune = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

type runeStack struct {
	lifo      []rune
	nextIndex int
}

func newRuneStack() *runeStack {
	rs := runeStack{}
	rs.lifo = make([]rune, 0, 10)
	rs.nextIndex = 0

	return &rs
}

func (rs *runeStack) push(r rune) {
	rs.lifo = append(rs.lifo, r)
	rs.nextIndex++
}

func (rs *runeStack) top() *rune {
	if rs.nextIndex > 0 {
		return &rs.lifo[rs.nextIndex-1]
	}
	return nil
}

func (rs *runeStack) empty() bool {
	return rs.nextIndex == 0
}

func (rs *runeStack) pop() rune {
	if !rs.empty() {
		t := *rs.top()
		rs.lifo = rs.lifo[:len(rs.lifo)-1]
		rs.nextIndex--
		return t
	}
	panic("Tried 'pop' on empty stack.")
}

func (rs *runeStack) clear() {
	rs.lifo = make([]rune, 0, 10)
	rs.nextIndex = 0
}

func (rs *runeStack) closeChunk(r rune) bool {
	if rs.empty() {
		return false
	}

	counterPart, _ := counterParts[r]
	popped := rs.pop()
	if popped != counterPart {
		return false
	}

	return true
}

func closingChunk(r rune) bool {
	for k := range counterParts {
		if k == r {
			return true
		}
	}
	return false
}

func task1(readings []string) int {
	stack := newRuneStack()

	score := 0
	for _, line := range readings {
		stack.clear()
		for _, r := range line {
			if !closingChunk(r) {
				stack.push(r)
			} else {
				if !stack.closeChunk(r) {
					score += lineScoring[r]
					break
				}
			}
		}
	}
	return score
}

func task2(readings []string) uint64 {
	stack := newRuneStack()
	incompleteLines := make([]string, 0, len(readings))
	incompleteStacks := make([]runeStack, 0, len(readings))
	scores := make([]uint64, 0, len(readings))

	for _, line := range readings {
		stack.clear()

		corrupted := false
		for _, r := range line {
			if !closingChunk(r) {
				stack.push(r)
			} else {
				if !stack.closeChunk(r) {
					corrupted = true
					break
				}
			}
		}
		if !stack.empty() && !corrupted {
			incompleteLines = append(incompleteLines, line)
			stackCopy := *stack
			incompleteStacks = append(incompleteStacks, stackCopy)
		}

	}
	scoring := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	for _, incompleteStack := range incompleteStacks {
		var score uint64 = 0
		for !incompleteStack.empty() {
			r := incompleteStack.pop()
			score *= 5
			score += uint64(scoring[r])
		}
		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })

	return scores[len(scores)/2]
}

func main() {
	subsystemReadings := utils.GetStringsFromInputFile("input.txt")
	start := time.Now()
	res1 := task1(subsystemReadings)
	end := time.Since(start)
	fmt.Println("[Task 1]:", res1, "| Executed in", end)

	start = time.Now()
	res2 := task2(subsystemReadings)
	end = time.Since(start)
	fmt.Println("[Task 2]:", res2, "| Executed in", end)
}
