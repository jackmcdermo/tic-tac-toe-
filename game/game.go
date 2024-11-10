package game

import (
	"fmt"

	"github.com/jackmcdermo/tic-tac-toe-/board"
)

// create a enum for possible move results
type MoveResult int

const (
	ValidMove MoveResult = iota
	SpaceOccupied
	XWin
	OWin
	Tie
)

type Player struct {
	Token              string // "X" or "O"
	IsAI               bool   // true if player is AI
	AiPlayerDifficulty int    // 0 - 9 (9 is hardest)
	Name               string
}

func NewPlayer(token string, isAI bool, aiPlayerDifficulty int, name string) Player {
	return Player{
		Token:              token,
		IsAI:               isAI,
		AiPlayerDifficulty: aiPlayerDifficulty,
		Name:               name,
	}
}

type Game struct {
	Board          *board.Board
	Player1        Player
	Player2        Player
	nextMovePlayer *Player
	RandomAI       bool
}

func NewGame(players []Player, randomAI bool) *Game {
	return &Game{
		Board:    board.NewBoard(),
		Player1:  players[0],
		Player2:  players[1],
		RandomAI: randomAI,
	}
}

func (g *Game) AwaitingAI() bool {
	return g.nextMovePlayer.IsAI
}

func (g *Game) NextMovePlayer() Player {
	return *g.nextMovePlayer
}

func (g *Game) InitGame() {
	g.Board.InitBoard()
	if g.getPlayerOneStartsFirst() {
		g.nextMovePlayer = &g.Player1
	} else {
		g.nextMovePlayer = &g.Player2
	}
}

func (g *Game) PrintMovePrompt() {

	if g.nextMovePlayer.IsAI {
		fmt.Println("Press enter for the AI player to go...")
	} else {
		fmt.Printf("%s, enter your move (row,col): ", g.nextMovePlayer.Name)
	}
}

func (g *Game) DoMove(row int, col int) MoveResult {
	var player Player

	// Place the move on the board. If the move was
	// successful, check if the game is over.
	if g.Board.PlaceToken(row, col, g.nextMovePlayer.Token) {
		// Check if the move resulted in a win
		if g.Board.CheckWin() {
			if player.Token == "X" {
				return XWin
			} else {
				return OWin
			}
		} else if g.Board.CheckTie() {
			return Tie
		}

		// Switch players
		if g.nextMovePlayer == &g.Player1 {
			g.nextMovePlayer = &g.Player2
		} else {
			g.nextMovePlayer = &g.Player1
		}

		return ValidMove

	} else {
		return SpaceOccupied
	}
}

func (g *Game) getPlayerOneStartsFirst() bool {
	// randomNum := rand.Float64()*2 - 1
	// return randomNum >= 0
	return true
}
