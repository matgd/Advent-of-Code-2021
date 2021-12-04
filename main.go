package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matgd/advent2021/utils"
)

// BingoBoard represnts bingo boards as array of arrays
type BingoBoard struct {
	board  [][]int
	marked [][]bool
}

func (b *BingoBoard) loadRow(row []int) {
	b.board = append(b.board, row)
	b.marked = append(b.marked, make([]bool, len(row)))
}

func (b *BingoBoard) checkMark(mark int) {
	found := false
	for i, row := range b.board {
		for j, value := range row {
			if value == mark {
				b.marked[i][j] = true
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}

func (b *BingoBoard) anyDiagonalMarked() bool {
	for i := range b.marked[0] {
		if b.marked[i][i] {
			return true
		}
	}
	return false
}

func (b *BingoBoard) checkBingo() bool {
	if b.anyDiagonalMarked() {
		for i, row := range b.marked {
			if utils.AllTrueArray(row) {
				return true
			}

			column := make([]bool, len(b.marked))
			for j := range b.marked {
				column[j] = b.marked[j][i]
			}
			if utils.AllTrueArray(column) {
				return true
			}
		}

	}
	return false
}

// check if column or row is filled
// start by checking if anything from diagonal is filled, if yes
//       check rows and columns

func getBoardsFromInput(input []string) []BingoBoard {
	splitStringNumbers := strings.Split(input[0], ",")
	drawnNumbers := make([]int, len(splitStringNumbers))
	for i, stringNumber := range splitStringNumbers {
		drawnNumbers[i], _ = strconv.Atoi(stringNumber)
	}

	boards := []BingoBoard{}
	currentBoardFilled := BingoBoard{}
	for _, boardRow := range input[2:] {
		if boardRow == "" {
			boards = append(boards, currentBoardFilled)
			currentBoardFilled = BingoBoard{}
			continue
		}

		rowWithEmptyStrings := strings.Split(boardRow, " ")

		rowNumbers := []int{}
		for _, value := range rowWithEmptyStrings {
			if value != "" {
				number, _ := strconv.Atoi(value)
				rowNumbers = append(rowNumbers, number)
			}
		}
		currentBoardFilled.loadRow(rowNumbers)
	}

	return boards
}

func main() {
	input := utils.GetStringsFromInputFile("input.txt")
	boards := getBoardsFromInput(input)
	fmt.Println(boards[0])
	boards[0].checkMark(27)
	boards[0].checkMark(2)
	boards[0].checkMark(57)
	boards[0].checkMark(4)
	boards[0].checkMark(69)
	fmt.Println(boards[0])
	fmt.Println(boards[0].checkBingo())
}
