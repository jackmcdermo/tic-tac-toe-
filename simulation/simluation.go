package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jackmcdermo/tic-tac-toe-/game"
	"github.com/jackmcdermo/tic-tac-toe-/minmax"
)

type Results struct {
	TotalRounds        int
	Player1Wins        int
	Player1Difficulty  int
	Player2Wins        int
	Player2Difficulty  int
	Ties               int
	Player1StartsFirst int
	TotalDuration      float64
	Player1Duration    float64
	Player2Duration    float64
}

type Simulation struct {
	Player1Difficulty int
	Player2Difficulty int
	TotalRounds       int
	RandomAI          bool
}

func NewSimulation(player1Difficulty int, player2Difficulty int, totalRounds int, randomAI bool) Simulation {
	return Simulation{
		Player1Difficulty: player1Difficulty,
		Player2Difficulty: player2Difficulty,
		TotalRounds:       totalRounds,
		RandomAI:          randomAI,
	}
}

func (s *Simulation) RunSimulation() Results {
	results := Results{
		TotalRounds:        s.TotalRounds,
		Player1Difficulty:  s.Player1Difficulty,
		Player2Difficulty:  s.Player2Difficulty,
		Player1Wins:        0,
		Player2Wins:        0,
		Ties:               0,
		Player1StartsFirst: 0,
		TotalDuration:      0,
		Player1Duration:    0,
		Player2Duration:    0,
	}

	start := time.Now()
	for i := 0; i < s.TotalRounds; i++ {
		s.simulateGame(&results)
	}
	results.TotalDuration = time.Since(start).Seconds()

	return results
}

func (s *Simulation) simulateGame(results *Results) {

	// Create a new game instance
	gameInstance := game.NewGame([]game.Player{
		game.NewPlayer("X", true, s.Player1Difficulty, "Player 1"),
		game.NewPlayer("O", true, s.Player2Difficulty, "Player 2"),
	}, s.RandomAI)

	gameInstance.InitGame()
	if gameInstance.NextMovePlayer().Token == "X" {
		results.Player1StartsFirst++
	}

	// Play the game until it is over
	for !gameInstance.Board.CheckWin() && !gameInstance.Board.CheckTie() {
		moveStart := time.Now()
		player := gameInstance.NextMovePlayer()
		row, col := minmax.GetBestMove(*gameInstance.Board, player.AiPlayerDifficulty, player.Token)
		moveDuration := time.Since(moveStart).Seconds()
		if player.Token == "X" {
			results.Player1Duration += moveDuration
		} else {
			results.Player2Duration += moveDuration
		}
		gameInstance.DoMove(row, col)
	}

	// Update the results based on the outcome of the game
	if gameInstance.Board.CheckWinForPlayer("X") {
		results.Player1Wins++
	} else if gameInstance.Board.CheckWinForPlayer("O") {
		results.Player2Wins++
	} else {
		results.Ties++
	}
}

func main() {

	player1 := flag.Int("p1", 5, "The AI Difficulty for player 1")
	player2 := flag.Int("p2", 5, "The AI Difficulty for player 2")
	rounds := flag.Int("r", 100, "The number of rounds to simulate")
	randomAI := flag.String("rai", "no", "Use random AI for player 2 (use 'yes' or 'no')")
	matrix := flag.String("matrix", "no", "Run full matrix of simulations")
	flag.Parse()

	if *matrix == "yes" {
		runTestMatrix(*rounds)
	} else {

		randomAIBool := false
		if *randomAI == "yes" {
			randomAIBool = true
		}

		simulation := NewSimulation(*player1, *player2, *rounds, randomAIBool)

		results := simulation.RunSimulation()

		fmt.Println("Simulation Results")
		fmt.Printf("Player 1 wins: %d\n", results.Player1Wins)
		fmt.Printf("Player 2 wins: %d\n", results.Player2Wins)
		fmt.Printf("Ties: %d\n", results.Ties)
		fmt.Printf("Duration (seconds): %f\n", results.TotalDuration)
		fmt.Println("CSV Results")

		writeHeaders()
		resultsAsCSV(results, randomAIBool)
	}
}

func runTestMatrix(rounds int) {
	writeHeaders()
	// for i := 1; i <= 9; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		for k := 0; k <= 1; k++ {
	// 			useRandomAI := k == 1
	// 			simulation := NewSimulation(i, j, rounds, useRandomAI)
	// 			results := simulation.RunSimulation()
	// 			resultsAsCSV(results, useRandomAI)
	// 		}
	// 	}
	// }

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			simulation := NewSimulation(i, j, rounds, false)
			results := simulation.RunSimulation()
			resultsAsCSV(results, false)
		}
	}

	fmt.Println("Matrix simulation complete")
}
func writeHeaders() {
	// Print out the headers
	fmt.Println("Total Rounds,Player 1 Wins,Player 1 Difficulty,Player 2 Wins,Player 2 Difficulty,Ties,Player 1 Starts First,Player1Duration,Player2Duration,TotalDuration,Random AI")
}

func resultsAsCSV(results Results, randomAI bool) {
	// Print out the results
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%f,%f,%f,%t\n", results.TotalRounds, results.Player1Wins, results.Player1Difficulty, results.Player2Wins, results.Player2Difficulty, results.Ties, results.Player1StartsFirst, results.Player1Duration, results.Player2Duration, results.TotalDuration, randomAI)
}
