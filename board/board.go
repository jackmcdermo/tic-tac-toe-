package board

import (
	"fmt"
	"log"
)

const emptySpace = " "

type Board struct {
	spaces [3][3]string
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) SetStartingBoard(postions [3][3]string) {
	b.spaces = postions
}

func (b *Board) InitBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.spaces[i][j] = emptySpace
		}
	}
}

// RemoveToken removes the token at the given row and column.
func (b *Board) RemoveToken(row int, col int) {
	b.spaces[row][col] = emptySpace
}

// PlaceToken places a move on the board at the given row and column
// with the given value. Returns true if the move was successfully placed,
// false if the space is already taken.
func (b *Board) PlaceToken(row int, col int, playerToken string) bool {

	if row < 0 || row > 2 || col < 0 || col > 2 {
		log.Printf("row and col must be between 0 and 2")
		return false
	}

	if b.spaces[row][col] != emptySpace {
		return false
	}

	b.spaces[row][col] = playerToken
	return true
}

func (b *Board) PrintBoard() {

	fmt.Println("")
	fmt.Println("   0|1|2")
	b.printHorizontalLine()
	b.printRow(0)
	b.printHorizontalLine()
	b.printRow(1)
	b.printHorizontalLine()
	b.printRow(2)
	fmt.Println("")
}

// GetToken returns the token at the given row and column.
func (b *Board) GetToken(row int, col int) string {
	return b.spaces[row][col]
}

// CheckWinForPlayer checks if the game has been won by a specific player.
func (b *Board) CheckWinForPlayer(playerToken string) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if b.compareSpaces(b.spaces[i][0], playerToken) && b.compareSpaces(b.spaces[i][1], playerToken) && b.compareSpaces(b.spaces[i][2], playerToken) {
			return true
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if b.compareSpaces(b.spaces[0][i], playerToken) && b.compareSpaces(b.spaces[1][i], playerToken) && b.compareSpaces(b.spaces[2][i], playerToken) {
			return true
		}
	}
	// Check diagonals
	if b.compareSpaces(b.spaces[0][0], playerToken) && b.compareSpaces(b.spaces[1][1], playerToken) && b.compareSpaces(b.spaces[2][2], playerToken) {
		return true
	}

	if b.compareSpaces(b.spaces[0][2], playerToken) && b.compareSpaces(b.spaces[1][1], playerToken) && b.compareSpaces(b.spaces[2][0], playerToken) {
		return true
	}
	return false
}

// CheckWin checks if the game has been won by a player.
func (b *Board) CheckWin() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if b.compareSpaces(b.spaces[i][0], b.spaces[i][1]) && b.compareSpaces(b.spaces[i][1], b.spaces[i][2]) {
			return true
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if b.compareSpaces(b.spaces[0][i], b.spaces[1][i]) && b.compareSpaces(b.spaces[1][i], b.spaces[2][i]) {
			return true
		}
	}
	// Check diagonals
	if b.compareSpaces(b.spaces[0][0], b.spaces[1][1]) && b.compareSpaces(b.spaces[1][1], b.spaces[2][2]) {
		return true
	}
	if b.compareSpaces(b.spaces[0][2], b.spaces[1][1]) && b.compareSpaces(b.spaces[1][1], b.spaces[2][0]) {
		return true
	}
	return false
}

// CheckTie checks if the game is a tie.
func (b *Board) CheckTie() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.spaces[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func (b *Board) compareSpaces(s1 string, s2 string) bool {
	return (s1 == s2) && (s1 != " ") && (s2 != " ")
}

func (b *Board) printHorizontalLine() {
	fmt.Println("   -----")
}

func (b *Board) printRow(row int) {
	fmt.Printf(" %d %s|%s|%s \n", row, b.spaces[row][0], b.spaces[row][1], b.spaces[row][2])
}
