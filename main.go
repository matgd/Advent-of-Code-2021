package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matgd/advent2021/utils"
)

// BingoBoard represnts bingo boards as array of arrays
type BingoBoard struct {
	board [][]int
}

func (b *BingoBoard) loadRow(row []int) {
	b.board = append(b.board, row)
}

func main() {
	input := utils.GetStringsFromInputFile("input.txt")

	splitStringNumbers := strings.Split(input[0], ",")
	drawnNumbers := make([]int, len(splitStringNumbers))
	for i, stringNumber := range splitStringNumbers {
		drawnNumbers[i], _ = strconv.Atoi(stringNumber)
	}

	fmt.Println(drawnNumbers)
	boards := []BingoBoard{}

	currentBoardFilled := BingoBoard{}
	for _, boardRow := range input[2:] {
		if boardRow == "" {
			boards = append(boards, currentBoardFilled)
			currentBoardFilled = BingoBoard{}
			continue
		}

		row := strings.Split(boardRow, " ")
		rowNumbers := make([]int, len(row))
		for i, number := range row {
			rowNumbers[i], _ = strconv.Atoi(number)
		}
		currentBoardFilled.loadRow(rowNumbers)
	}

	fmt.Println(boards)
}
