package main

import "fmt"

type Board [][]State
type State int

type ClientBoard [][]string

const (
	DRAW State = iota
	EMPTY
	X
	O
)

func (board Board) Draw() {
	stateToStringMap := make(map[State]string)
	stateToStringMap[EMPTY] = " "
	stateToStringMap[X] = "X"
	stateToStringMap[O] = "O"

	for row := range board {
		for col := range board {
			state := board[row][col]
			fmt.Printf("%s|", stateToStringMap[state])
		}
		fmt.Println()
	}
}

func NewBoard(size int) Board {
	board := make(Board, size)
	for row := range board {
		board[row] = make([]State, size)

		for col := range board[row] {
			board[row][col] = EMPTY
		}
	}

	return board
}

func NewClientBoard(board Board) ClientBoard {
	stateToStringMap := make(map[State]string)
	stateToStringMap[EMPTY] = " "
	stateToStringMap[X] = "X"
	stateToStringMap[O] = "O"

	value := make(ClientBoard, len(board))
	for row := range board {
		value[row] = make([]string, len(board))
		for col := range board {
			value[row][col] = stateToStringMap[board[row][col]]
		}
	}

	return value
}

func NewBoardFromClient(clientBoard ClientBoard) Board {
	stringToStateMap := make(map[string]State)
	stringToStateMap[" "] = EMPTY
	stringToStateMap["X"] = X
	stringToStateMap["O"] = O

	value := make(Board, len(clientBoard))
	for row := range clientBoard {
		value[row] = make([]State, len(clientBoard))
		for col := range clientBoard {
			value[row][col] = stringToStateMap[clientBoard[row][col]]
		}
	}

	return value
}
