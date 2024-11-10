package minmax

import (
	"math"

	"github.com/jackmcdermo/tic-tac-toe-/board"
)

func GetBestMove(board board.Board, maxDepth int, playerToken string) (int, int) {
	bestRow, bestCol := -1, -1
	var bestScore int

	bestScore = math.MinInt

	// iterate over the board and find the first empty space
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetToken(i, j) == " " {

				// Simulate a move for the AI player
				board.PlaceToken(i, j, playerToken)

				// Call minmax to get the score for the move
				score := minmax(board, 0, false, maxDepth, playerToken)

				// Undo the move
				board.RemoveToken(i, j)

				if score > bestScore {
					bestScore = score
					bestRow = i
					bestCol = j
				}
			}
		}
	}

	return bestRow, bestCol
}

// _minmax is a recursive function that implements the minimax algorithm
func minmax(board board.Board, depth int, isMaximizing bool, maxDepth int, playerToken string) int {

	opponentToken := "X"
	if playerToken == "X" {
		opponentToken = "O"
	}

	if board.CheckWinForPlayer(opponentToken) {
		return -1
	} else if board.CheckWinForPlayer(playerToken) {
		return 1
	} else if board.CheckTie() || depth == maxDepth {
		return 0
	}

	depth += 1
	if isMaximizing {
		maxEval := math.MinInt
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board.GetToken(i, j) == " " {
					// Simulate a move for the AI player
					board.PlaceToken(i, j, playerToken)

					// Recursively call _minmax with the new board state
					score := minmax(board, depth, false, maxDepth, playerToken)

					// Undo the move
					board.RemoveToken(i, j)
					maxEval = max(maxEval, score)
				}
			}
		}
		return maxEval

	} else {
		minEval := math.MaxInt
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board.GetToken(i, j) == " " {

					// Simulate a move for the human player
					board.PlaceToken(i, j, opponentToken)

					// Recursively call _minmax with the new board state
					score := minmax(board, depth, true, maxDepth, playerToken)

					// Undo the move
					board.RemoveToken(i, j)
					minEval = min(minEval, score)
				}
			}
		}
		return minEval
	}
}
