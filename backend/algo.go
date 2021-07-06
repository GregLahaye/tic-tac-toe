package main

func (board Board) MaxAlphaBeta(alpha, beta int) (int, int, int) {
	maxValue := -2
	var x int
	var y int

	result := board.DetermineState()

	if result == X {
		return -1, 0, 0
	} else if result == O {
		return 1, 0, 0
	} else if result == FULL {
		return 0, 0, 0
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == EMPTY {
				board[row][col] = O
				minValue, _, _ := board.MinAlphaBeta(alpha, beta)
				if minValue > maxValue {
					maxValue = minValue
					x = row
					y = col
				}
				board[row][col] = EMPTY

				if maxValue >= beta {
					return maxValue, x, y
				}

				if maxValue > alpha {
					alpha = maxValue
				}
			}
		}
	}

	return maxValue, x, y
}

func (board Board) MinAlphaBeta(alpha, beta int) (int, int, int) {
	minValue := 2
	var x int
	var y int

	result := board.DetermineState()

	if result == X {
		return -1, 0, 0
	} else if result == O {
		return 1, 0, 0
	} else if result == FULL {
		return 0, 0, 0
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == EMPTY {
				board[row][col] = X
				maxValue, _, _ := board.MaxAlphaBeta(alpha, beta)
				if maxValue < minValue {
					minValue = maxValue
					x = row
					y = col
				}
				board[row][col] = EMPTY

				if minValue <= alpha {
					return minValue, x, y
				}

				if minValue < beta {
					beta = minValue
				}
			}
		}
	}

	return minValue, x, y
}
