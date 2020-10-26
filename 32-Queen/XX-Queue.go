package main

import (
	"fmt"
)

func main() {
	board := getBoard()

	updatedBoard := findBlockedBoxes(3, board)
	fmt.Println(updatedBoard[4][5])
	fmt.Println(updatedBoard[2][6])
	fmt.Println(updatedBoard)
}

func getBoard() [][]int {
	boardSize := make([]int, 8)
	board := [][]int{}
	for range boardSize {
		dt := &boardSize
		board = append(board, *dt)
	}
	return board
}

func findBlockedBoxes(boxId int, board [][]int) [][]int {
	boardSize := []int{3, 3, 3, 3, 3, 3, 3, 3}
	board[3] = boardSize
	board[4][6] = 2
	board[2][6] = 6

	return board
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
