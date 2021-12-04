package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matgd/advent2021/utils"
)

// BingoBoard represnts bingo boards as array of arrays
type BingoBoard struct {
	id         int
	board      [][]int
	marked     [][]bool
	lastMarked int
}

type foundChannel struct {
	boardID int
	found   bool
}

func (b *BingoBoard) loadRow(row []int) {
	b.board = append(b.board, row)
	b.marked = append(b.marked, make([]bool, len(row)))
}

func (b *BingoBoard) checkMark(mark int, checkChannel chan foundChannel) {
	found := false
	for i, row := range b.board {
		for j, value := range row {
			if value == mark {
				b.marked[i][j] = true
				b.lastMarked = value
				found = true
				checkChannel <- foundChannel{boardID: b.id, found: true}
				return
			}
		}
		if found {
			checkChannel <- foundChannel{boardID: b.id, found: true}
			return
		}
	}
	checkChannel <- foundChannel{boardID: b.id, found: false}
}

func (b *BingoBoard) anyDiagonalMarked() bool {
	for i := range b.marked[0] {
		if b.marked[i][i] {
			return true
		}
	}
	return false
}

func (b *BingoBoard) checkBingo(chBingo chan bool, chLastMarked chan int) {
	if b.anyDiagonalMarked() {
		for i, row := range b.marked {
			if utils.AllTrueArray(row) {
				chBingo <- true
				chLastMarked <- b.lastMarked
			}

			column := make([]bool, len(b.marked))
			for j := range b.marked {
				column[j] = b.marked[j][i]
			}
			if utils.AllTrueArray(column) {
				chBingo <- true
				chLastMarked <- b.lastMarked
			}
		}

	}
	chBingo <- false
}

// check if column or row is filled
// start by checking if anything from diagonal is filled, if yes
//       check rows and columns

func getBoardsFromInput(input []string) []BingoBoard {
	boards := []BingoBoard{}
	nextBoardID := 0
	currentBoardFilled := BingoBoard{id: nextBoardID}
	for _, boardRow := range input[2:] {
		if boardRow == "" {
			boards = append(boards, currentBoardFilled)
			nextBoardID++
			currentBoardFilled = BingoBoard{id: nextBoardID}
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

func task1() {
	input := utils.GetStringsFromInputFile("input.txt")
	boards := getBoardsFromInput(input)

	splitStringNumbers := strings.Split(input[0], ",")
	drawnNumbers := make([]int, len(splitStringNumbers))
	for i, stringNumber := range splitStringNumbers {
		drawnNumbers[i], _ = strconv.Atoi(stringNumber)

	}

	// bingoChannel := make(chan bool)
	// lastMarkedChannel := make(chan int, 1)

	// main loop
	for _, n := range drawnNumbers {
		checkMarkChannel := make(chan foundChannel, len(boards))
		for _, b := range boards {
			go b.checkMark(n, checkMarkChannel)
		}

		// foundMarkBoardsIds := []int{}
		// for v := range checkMarkChannel {
		// if v.found {
		// foundMarkBoardsIds = append(foundMarkBoardsIds, v.boardID)
		// }
		// }
		fmt.Println(len(checkMarkChannel))
	}

}

func main() {
	task1()
}
