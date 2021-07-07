package main

/*
 * O = max player
 * X = min player
 *
 * MaxAlphaBeta determines the optimal move for O
 * MinAlphaBeta determines the optimal move for X
 *
 * alpha = lowest value assured for max player
 * beta = highest value assured for min player
 *
 * -1 = X wins
 *  0 = draw
 *  1 = O wins
 */

func (board Board) MaxAlphaBeta(alpha, beta int) (int, int, int) {
	maxValue := -2
	var x int
	var y int

	result := board.DetermineState()

	if result == X {
		return -1, 0, 0
	} else if result == O {
		return 1, 0, 0
	} else if result == DRAW {
		return 0, 0, 0
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == EMPTY {
				// place O here
				board[row][col] = O
				// determine min value if we made this move (min value = best for X)
				minValue, _, _ := board.MinAlphaBeta(alpha, beta)
				// if min value is greater than current max value, it is currently the best move
				if minValue > maxValue {
					maxValue = minValue
					x = row
					y = col
				}
				// reset O placement
				board[row][col] = EMPTY

				if maxValue > alpha {
					alpha = maxValue
				}

				// because alpha is greater than beta, min player would never choose it, so we stop searching
				if alpha >= beta {
					return maxValue, x, y
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
	} else if result == DRAW {
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

				if minValue < beta {
					beta = minValue
				}

				if beta <= alpha {
					return minValue, x, y
				}
			}
		}
	}

	return minValue, x, y
}
