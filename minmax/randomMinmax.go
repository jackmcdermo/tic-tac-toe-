package minmax

import (
	"log"
	"math"
	"math/rand"

	"github.com/jackmcdermo/tic-tac-toe-/board"
)

func GetBestMoveWithRandom(board board.Board, maxDepth int, playerToken string) (int, int) {
	bestRow, bestCol := -1, -1
	var bestScore int

	bestScore = math.MinInt
	rs := NewRandomSpot(board)
	for spot := rs.GetNextOpenMove(board); spot != nil; spot = rs.GetNextOpenMove(board) {

		// Simulate a move for the AI player
		ok := board.PlaceToken(spot.row, spot.col, playerToken)
		if !ok {
			log.Printf("trying to place a token in a non-empty spot")
			panic("trying to place a token in a non-empty spot")
		}

		// Call minmax to get the score for the move
		score := randminmax(board, 0, false, maxDepth, playerToken)

		// Undo the move
		board.RemoveToken(spot.row, spot.col)

		if score > bestScore {
			bestScore = score
			bestRow = spot.row
			bestCol = spot.col
		}
	}
	return bestRow, bestCol
}

func randminmax(board board.Board, depth int, isMaximizing bool, maxDepth int, playerToken string) int {

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
		rs := NewRandomSpot(board)
		for spot := rs.GetNextOpenMove(board); spot != nil; spot = rs.GetNextOpenMove(board) {

			// Simulate a move for the player
			ok := board.PlaceToken(spot.row, spot.col, playerToken)
			if !ok {
				log.Printf("trying to place a token in a non-empty spot")
				panic("trying to place a token in a non-empty spot")
			}

			// Recursively call _minmax with the new board state
			score := randminmax(board, depth, false, maxDepth, playerToken)

			// Undo the move
			board.RemoveToken(spot.row, spot.col)
			maxEval = max(maxEval, score)
		}
		return maxEval

	} else {
		minEval := math.MaxInt
		rs := NewRandomSpot(board)

		for spot := rs.GetNextOpenMove(board); spot != nil; spot = rs.GetNextOpenMove(board) {

			// Simulate a move for the opponent
			ok := board.PlaceToken(spot.row, spot.col, opponentToken)
			if !ok {
				log.Printf("trying to place a token in a non-empty spot")
				panic("trying to place a token in a non-empty spot")
			}

			// Recursively call _minmax with the new board state
			score := randminmax(board, depth, true, maxDepth, playerToken)

			// Undo the move
			board.RemoveToken(spot.row, spot.col)
			minEval = min(minEval, score)

		}
		return minEval
	}
}

type randomSpot struct {
	startingSpot spot
	currentRow   int
	currentCol   int
	totalMoves   int
	id           int
}

type Spot struct {
	row int
	col int
}

func (rs *randomSpot) GetNextOpenMove(board board.Board) *Spot {
	if rs.totalMoves == 0 {
		rs.totalMoves++
		return &Spot{rs.startingSpot.row, rs.startingSpot.col}
	}

	for rs.totalMoves < 9 {
		// Move over one column
		rs.currentRow++

		// If we are at the end of the row, move to the next column
		// and reset the row
		if rs.currentRow > 2 {
			rs.currentRow = 0
			rs.currentCol++
		}

		// If we are at the end of the column, start at the top
		if rs.currentCol > 2 {
			rs.currentCol = 0
		}

		rs.totalMoves++
		if board.GetToken(rs.currentRow, rs.currentCol) == " " {
			return &Spot{rs.currentRow, rs.currentCol}
		}
	}

	// There are no more open spots
	return nil
}

type spot struct {
	row int
	col int
}

var totalRandomSpots int = 0

func NewRandomSpot(board board.Board) randomSpot {

	id := totalRandomSpots
	totalRandomSpots++
	// Get all open spots on the board
	openSpots := make(map[int]spot)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetToken(i, j) == " " {
				openSpots[len(openSpots)] = spot{i, j}
			}
		}
	}
	randomIndex := rand.Intn(len(openSpots))

	startingSpot := openSpots[randomIndex]

	return randomSpot{startingSpot, startingSpot.row, startingSpot.col, 0, id}
}
