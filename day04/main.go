package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/matgd/advent2021/utils"
)

// BingoBoard represnts bingo boards as array of arrays
type BingoBoard struct {
	id     int
	board  [][]int
	marked [][]bool
}

type foundChannel struct {
	boardID int
	found   bool
}

func (b *BingoBoard) loadRow(row []int) {
	b.board = append(b.board, row)
	b.marked = append(b.marked, make([]bool, len(row)))
}

func (b *BingoBoard) checkMark(mark int, checkChannel chan foundChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	for i, row := range b.board {
		for j, value := range row {
			if value == mark {
				b.marked[i][j] = true
				checkChannel <- foundChannel{boardID: b.id, found: true}
				return
			}
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

func (b *BingoBoard) checkBingo(bingoChannel chan foundChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	if b.anyDiagonalMarked() {
		for i, row := range b.marked {
			if utils.AllTrueArray(row) {
				bingoChannel <- foundChannel{boardID: b.id, found: true}
				return
			}

			column := make([]bool, len(b.marked))
			for j := range b.marked {
				column[j] = b.marked[j][i]
			}
			if utils.AllTrueArray(column) {
				bingoChannel <- foundChannel{boardID: b.id, found: true}
				return
			}
		}

	}
	bingoChannel <- foundChannel{boardID: b.id, found: false}
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
	foundMarkBoardsIds := []int{}
	for _, n := range drawnNumbers {
		var checkChannels []chan foundChannel
		for range boards {
			checkChannels = append(checkChannels, make(chan foundChannel, 1))
		}

		var wg sync.WaitGroup
		wg.Add(len(boards))
		for i, b := range boards {
			b := b // https://stackoverflow.com/a/57080138
			// if the assignment is not present we 'b' will change in scope later
			go b.checkMark(n, checkChannels[i], &wg)
		}
		wg.Wait()

		for i := range checkChannels {
			result := <-checkChannels[i]
			if result.found {
				foundMarkBoardsIds = append(foundMarkBoardsIds, result.boardID)
			}
		}

		var bingoChannels []chan foundChannel
		for range foundMarkBoardsIds {
			bingoChannels = append(bingoChannels, make(chan foundChannel, 1))
		}
		var wgBingo sync.WaitGroup
		wgBingo.Add(len(foundMarkBoardsIds))
		for i, fm := range foundMarkBoardsIds {
			fm := fm
			go boards[fm].checkBingo(bingoChannels[i], &wgBingo)
		}

		wgBingo.Wait()
		for i := range bingoChannels {
			result := <-bingoChannels[i]
			if result.found {
				winningBoard := boards[result.boardID]
				fmt.Println("Winner!", result.boardID, winningBoard)

				sumUnmarked := 0
				for m, rowMarkings := range winningBoard.marked {
					for n, v := range rowMarkings {
						if !v {
							sumUnmarked += winningBoard.board[m][n]
						}
					}
				}
				fmt.Println("[Task 1]:", sumUnmarked)
				fmt.Println("[Task 1]:", sumUnmarked*n)
				return
			}
		}
	}

}

func task2() {
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
	foundMarkBoardsIds := []int{}
	winningBoards := map[int]bool{}
	for drawnNumbersIndex, n := range drawnNumbers {
		var checkChannels []chan foundChannel
		for range boards {
			checkChannels = append(checkChannels, make(chan foundChannel, 1))
		}

		var wg sync.WaitGroup
		wg.Add(len(boards))
		for i, b := range boards {
			b := b // https://stackoverflow.com/a/57080138
			// if the assignment is not present we 'b' will change in scope later
			go b.checkMark(n, checkChannels[i], &wg)
		}
		wg.Wait()

		for i := range checkChannels {
			result := <-checkChannels[i]
			if result.found {
				foundMarkBoardsIds = append(foundMarkBoardsIds, result.boardID)
			}
		}

		var bingoChannels []chan foundChannel
		for range foundMarkBoardsIds {
			bingoChannels = append(bingoChannels, make(chan foundChannel, 1))
		}
		var wgBingo sync.WaitGroup
		wgBingo.Add(len(foundMarkBoardsIds))
		for i, fm := range foundMarkBoardsIds {
			fm := fm
			go boards[fm].checkBingo(bingoChannels[i], &wgBingo)
		}

		wgBingo.Wait()
		for i := range bingoChannels {
			result := <-bingoChannels[i]
			if result.found {
				winningBoard := boards[result.boardID]
				if len(winningBoards) < len(boards)-1 {
					winningBoards[winningBoard.id] = true
					continue
				}
				for q := range boards {
					if _, ok := winningBoards[q]; !ok {
						winningBoard = boards[q]
					}
				}
				for leftNumbersIndex := drawnNumbersIndex + 1; leftNumbersIndex < len(drawnNumbers); leftNumbersIndex++ {
					lastBingo := make(chan foundChannel, 1)
					wg.Add(1)
					winningBoard.checkMark(drawnNumbers[leftNumbersIndex], lastBingo, &wg)
					bingo := <-lastBingo
					if bingo.found {
						n = drawnNumbers[leftNumbersIndex]
						break
					}
				}

				fmt.Println("Last winner:", winningBoard.id, winningBoard)

				sumUnmarked := 0
				for m, rowMarkings := range winningBoard.marked {
					for n, v := range rowMarkings {
						if !v {
							sumUnmarked += winningBoard.board[m][n]
						}
					}
				}
				fmt.Println("[Task 2]:", sumUnmarked)
				fmt.Println("[Task 2]:", n)
				fmt.Println("[Task 2]:", sumUnmarked*n)
				return
			}
		}
	}

}

func main() {
	task1()
	task2()
}
