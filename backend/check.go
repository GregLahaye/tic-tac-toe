package main

func (board Board) DetermineState() State {
	if state, win := board.CheckVerticalWin(); win {
		return state
	}

	if state, win := board.CheckHorizontalWin(); win {
		return state
	}

	if state, win := board.CheckFirstDiagonalWin(); win {
		return state
	}

	if state, win := board.CheckSecondDiagonalWin(); win {
		return state
	}

	return board.CheckBoardState()
}

func (board Board) CheckVerticalWin() (State, bool) {
	for col := 0; col < len(board); col++ {
		state := board[0][col]
		filled := state != EMPTY
		for row := 0; row < len(board); row++ {
			filled = filled && board[row][col] == state
		}

		if filled {
			return state, true
		}
	}

	return EMPTY, false
}

func (board Board) CheckHorizontalWin() (State, bool) {
	for row := 0; row < len(board); row++ {
		state := board[row][0]
		filled := state != EMPTY
		for col := 0; col < len(board); col++ {
			filled = filled && board[row][col] == state
		}

		if filled {
			return state, true
		}
	}

	return EMPTY, false
}

func (board Board) CheckFirstDiagonalWin() (State, bool) {
	state := board[0][len(board)-1]
	filled := state != EMPTY

	for i := 0; i < len(board); i++ {
		filled = filled && board[i][len(board)-1-i] == state
	}

	if filled {
		return state, true
	}

	return EMPTY, false
}

func (board Board) CheckSecondDiagonalWin() (State, bool) {
	state := board[0][0]
	filled := state != EMPTY

	for i := 0; i < len(board); i++ {
		filled = filled && board[i][i] == state
	}

	if filled {
		return state, true
	}

	return EMPTY, false
}

func (board Board) CheckBoardState() State {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == EMPTY {
				return EMPTY
			}
		}
	}

	return DRAW
}
