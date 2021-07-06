package main

import "fmt"

type Board [][]State
type State int

type ClientBoard [][]string

const (
	FULL State = iota
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

func NewBoard(clientBoard ClientBoard) Board {
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
