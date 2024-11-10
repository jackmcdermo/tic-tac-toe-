package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jackmcdermo/tic-tac-toe-/game"
	"github.com/jackmcdermo/tic-tac-toe-/minmax"
)

type channelData struct {
	gameInstance *game.Game
	input        string
}

const (
	NEW_GAME_PROMPT = "Enter '1' to play against a friend, '2' to play against the AI,  or '3' for two AI players to square off! Press 'q' to quit."
	USE_RANDOM_AI   = false
)

func main() {

	fmt.Println("Welcome to Tic-Tac-Toe!")
	fmt.Println(NEW_GAME_PROMPT)

	var gameInstance *game.Game = nil

	inputChannel := make(chan channelData)

	go func() {

		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			inputChannel <- channelData{gameInstance, text}
		}
	}()
	for channelData := range inputChannel {
		gameInstance = handleUserInput(channelData)
	}
}

// handleUserInput handles the user's input based on the current game state.
func handleUserInput(data channelData) *game.Game {

	// if there is no game instance, the user has the option to start a new game
	// or exit the program
	if data.gameInstance == nil {
		return handleNewGameInput(data.input)
	}

	if data.gameInstance != nil && playerNeedsAILevel(data.gameInstance.Player1) {
		return handleSetAILevel(data.gameInstance, &data.gameInstance.Player1, data.input)
	}

	if data.gameInstance != nil && playerNeedsAILevel(data.gameInstance.Player2) {
		return handleSetAILevel(data.gameInstance, &data.gameInstance.Player2, data.input)
	}

	return handlePlayerMove(data.gameInstance, data.input)
}

func playerNeedsAILevel(player game.Player) bool {
	return player.IsAI && player.AiPlayerDifficulty == -1
}

// parseMove parses a move string in the format "row, col" and returns the row and column
// values. Returns an error if the move is not in the correct format or if the row or column
// values are out of range.
func parseMove(move string) (int, int, error) {
	var row, col int
	_, err := fmt.Sscanf(move, "%d,%d", &row, &col)
	if err != nil {
		return 0, 0, err
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return 0, 0, fmt.Errorf("invalid move")
	}
	return row, col, nil
}

// initGame initializes a new game instance with the given player configuration.
func initGame(humanPlayerCount int, randomAI bool, p1ai int, p2ai int) *game.Game {

	player1 := game.NewPlayer("X", false, 0, "Player 1 (X)")
	player2 := game.NewPlayer("O", false, 0, "Player 2 (O)")

	if humanPlayerCount < 2 {
		name := fmt.Sprintf("Player 2 (O, ai: %d)", p1ai)
		player2 = game.NewPlayer("O", true, p1ai, name)
	}

	if humanPlayerCount < 1 {
		name := fmt.Sprintf("Player 1 (X, ai: %d)", p1ai)
		player1 = game.NewPlayer("X", true, p2ai, name)
	}

	gameInstance := game.NewGame([]game.Player{player1, player2}, randomAI)
	gameInstance.InitGame()

	if playerNeedsAILevel(player1) {
		printAILevelPrompt(player1.Name)
		return gameInstance
	} else if playerNeedsAILevel(player2) {
		printAILevelPrompt(player2.Name)
		return gameInstance
	}
	return startGame(gameInstance)
}

func startGame(gameInstance *game.Game) *game.Game {
	player := gameInstance.NextMovePlayer()
	fmt.Printf("%s won the coin toss. so they go first!\n", player.Name)
	gameInstance.Board.PrintBoard()
	gameInstance.PrintMovePrompt()
	return gameInstance
}

// handleNewGameInput handles the user's input when starting a new game.
func handleNewGameInput(input string) *game.Game {

	switch input {
	case "1\n": // Two human players
		return initGame(2, USE_RANDOM_AI, 0, 0)
	case "2\n": // Player vs AI
		return initGame(1, USE_RANDOM_AI, 0, -1)
	case "3\n": // AI vs AI
		return initGame(0, USE_RANDOM_AI, -1, -1)
	case "q\n": // Quit
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	default:
		fmt.Printf("Invalid input. %s\n", NEW_GAME_PROMPT)
	}
	return nil
}

func handleSetAILevel(gameInstance *game.Game, player *game.Player, input string) *game.Game {

	level, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number between 1 and 10.")
		return gameInstance
	}

	if level < 1 || level > 10 {
		fmt.Println("Please enter a valid number between 1 and 10.")
		return gameInstance
	}
	player.AiPlayerDifficulty = level - 1
	return gameInstance
}

// handlePlayerMove handles the user's move input. If the move is valid, the move is placed
// on the board and the game state is checked. If the game is over, the game state is printed
// and the user is prompted to start a new game.
func handlePlayerMove(gameInstance *game.Game, move string) *game.Game {

	// Check if the user wants to quit
	if move == "q\n" {
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	}

	var row, col int
	var err error
	nextMovePlayer := gameInstance.NextMovePlayer()
	// Have the AI player make a move if it is their turn
	if nextMovePlayer.IsAI {
		fmt.Println("AI player is making a move...")
		if gameInstance.RandomAI {
			row, col = minmax.GetBestMoveWithRandom(*gameInstance.Board, nextMovePlayer.AiPlayerDifficulty, nextMovePlayer.Token)
		} else {
			row, col = minmax.GetBestMove(*gameInstance.Board, nextMovePlayer.AiPlayerDifficulty, nextMovePlayer.Token)
		}
		// Ortherwise, parse the user's move
	} else {
		row, col, err = parseMove(move)
		if err != nil {
			fmt.Println(err)
			gameInstance.PrintMovePrompt()
		}
	}

	// Place the user's move on the board
	result := gameInstance.DoMove(row, col)

	switch result {

	case game.ValidMove:
		printNextMoveMessage(gameInstance, "")
	case game.SpaceOccupied:
		printNextMoveMessage(gameInstance, "That space is already occupied. Please try again.")
	case game.XWin:
		printGameOverMessage("Player 1 wins!", gameInstance)
		return nil
	case game.OWin:
		printGameOverMessage("Player 2 wins!", gameInstance)
		return nil
	case game.Tie:
		printGameOverMessage("It's a tie!", gameInstance)
		return nil
	}

	return gameInstance
}

// printNextMoveMessage prints the next move prompt and the current board state.
func printNextMoveMessage(gameInstance *game.Game, msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	gameInstance.Board.PrintBoard()
	gameInstance.PrintMovePrompt()
}

// printGameOverMessage prints the game over message and the new game prompt.
func printGameOverMessage(msg string, gameInstance *game.Game) {

	gameInstance.Board.PrintBoard()
	fmt.Printf("Game over! %s\n", msg)
	fmt.Println("")
	fmt.Println(NEW_GAME_PROMPT)
}

func printAILevelPrompt(playerName string) {
	fmt.Printf("Enter the AI level (1-10) for %s, 10 being the most difficult: ", playerName)
}
