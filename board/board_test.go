package board

import (
	"testing"
)

func TestCheckWinForPlayer(t *testing.T) {
	tests := []struct {
		name        string
		board       [3][3]string
		playerToken string
		expected    bool
	}{
		{
			name: "Player wins with a row",
			board: [3][3]string{
				{"X", "X", "X"},
				{" ", " ", " "},
				{" ", " ", " "},
			},
			playerToken: "X",
			expected:    true,
		},
		{
			name: "Player wins with a column",
			board: [3][3]string{
				{"X", " ", " "},
				{"X", " ", " "},
				{"X", " ", " "},
			},
			playerToken: "X",
			expected:    true,
		},
		{
			name: "Player wins with a diagonal",
			board: [3][3]string{
				{"X", " ", " "},
				{" ", "X", " "},
				{" ", " ", "X"},
			},
			playerToken: "X",
			expected:    true,
		},
		{
			name: "Player does not win",
			board: [3][3]string{
				{"X", "O", "X"},
				{"O", "X", "O"},
				{"O", "X", "O"},
			},
			playerToken: "X",
			expected:    false,
		},
		// {
		// 	name: "Empty board",
		// 	board: [3][3]string{
		// 		{" ", " ", " "},
		// 		{" ", " ", " "},
		// 		{" ", " ", " "},
		// 	},
		// 	playerToken: "X",
		// 	expected:    false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{spaces: tt.board}
			result := b.CheckWinForPlayer(tt.playerToken)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
func TestGetToken(t *testing.T) {
	tests := []struct {
		name     string
		board    [3][3]string
		row      int
		col      int
		expected string
	}{
		{
			name: "Get token from top-left corner",
			board: [3][3]string{
				{"X", " ", " "},
				{" ", " ", " "},
				{" ", " ", " "},
			},
			row:      0,
			col:      0,
			expected: "X",
		},
		{
			name: "Get token from center",
			board: [3][3]string{
				{" ", " ", " "},
				{" ", "O", " "},
				{" ", " ", " "},
			},
			row:      1,
			col:      1,
			expected: "O",
		},
		{
			name: "Get token from bottom-right corner",
			board: [3][3]string{
				{" ", " ", " "},
				{" ", " ", " "},
				{" ", " ", "X"},
			},
			row:      2,
			col:      2,
			expected: "X",
		},
		{
			name: "Get token from empty space",
			board: [3][3]string{
				{" ", " ", " "},
				{" ", " ", " "},
				{" ", " ", " "},
			},
			row:      1,
			col:      1,
			expected: " ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{spaces: tt.board}
			result := b.GetToken(tt.row, tt.col)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
