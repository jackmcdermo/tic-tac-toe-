package minmax

import (
	"testing"

	"github.com/jackmcdermo/tic-tac-toe-/board"
	"github.com/stretchr/testify/suite"
)

type bestMoveSuite struct {
	suite.Suite
	board board.Board
}

func TestGetestMoveSuite(t *testing.T) {
	suite.Run(t, new(bestMoveSuite))
}

func (s *bestMoveSuite) SetupSuite() {
	s.board = *board.NewBoard()
}

func (s *bestMoveSuite) TestGetBestMoveOneMoveLeft() {

	openSpaces := 1
	spaces := [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", " "},
	}
	s.board.SetStartingBoard(spaces)
	s.evaluateMove(2, 2, openSpaces)
}

func (s *bestMoveSuite) TestGetBestMoveBlockOpponentWinningMove() {

	openSpaces := 4
	spaces := [3][3]string{
		{"X", "X", " "},
		{"O", " ", " "},
		{"O", "X", " "},
	}
	s.board.SetStartingBoard(spaces)
	s.evaluateMove(0, 2, openSpaces)
}

func (s *bestMoveSuite) TestGetBestMoveWinningMove() {

	openSpaces := 4
	spaces := [3][3]string{
		{"X", "X", " "},
		{"O", "O", " "},
		{"X", " ", " "},
	}
	s.board.SetStartingBoard(spaces)
	s.evaluateMove(1, 2, openSpaces)
}

func (s *bestMoveSuite) TestGetBestMoveOneMoveLeft_Random() {

	openSpaces := 1
	spaces := [3][3]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", " "},
	}
	s.board.SetStartingBoard(spaces)
	s.evaluateRandomMove(2, 2, openSpaces)
}

func (s *bestMoveSuite) TestGetBestMoveBlockOpponentWinningMove_Random() {

	openSpaces := 4
	spaces := [3][3]string{
		{"X", "X", " "},
		{"O", " ", " "},
		{"O", "X", " "},
	}
	s.board.SetStartingBoard(spaces)
	s.evaluateRandomMove(0, 2, openSpaces)
}

func (s *bestMoveSuite) TestGetBestMoveWinningMove_Random() {

	openSpaces := 4
	spaces := [3][3]string{
		{"X", "X", " "},
		{"O", "O", " "},
		{"X", " ", " "},
	}
	s.board.SetStartingBoard(spaces)
	s.evaluateRandomMove(1, 2, openSpaces)
}

func (s *bestMoveSuite) evaluateRandomMove(expectedRow int, expectedCol int, openSpaces int) {
	row, col := GetBestMoveWithRandom(s.board, openSpaces-1, "O")
	s.Equal(expectedRow, row)
	s.Equal(expectedCol, col)
}
func (s *bestMoveSuite) evaluateMove(expectedRow int, expectedCol int, openSpaces int) {
	row, col := GetBestMove(s.board, openSpaces-1, "O")
	s.Equal(expectedRow, row)
	s.Equal(expectedCol, col)
}
