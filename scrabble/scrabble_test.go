package main

import "testing"

// TestComputeScore covers basic scoring, case, punctuation, and edge cases.
func TestComputeScore(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected int
	}{
		// Basic letter scoring tests
		{"single letter A", "A", 1},
		{"single letter Z", "Z", 10},
		{"single letter Q", "Q", 10},
		{"single letter X", "X", 8},

		// Case sensitivity tests
		{"lowercase a", "a", 1},
		{"lowercase z", "z", 10},
		{"mixed case Word", "Word", 8}, // W(4) + o(1) + r(1) + d(2) = 8
		{"all uppercase WORD", "WORD", 8},
		{"all lowercase word", "word", 8},

		// Punctuation and symbols
		{"word with punctuation Question?", "Question?", 17},
		{"word with punctuation Question!", "Question!", 17},
		{"word with comma Oh,", "Oh,", 5},
		{"word with exclamation hai!", "hai!", 6},
		{"word with dash", "co-op", 8}, // c(3)+o(1)+o(1)+p(3)=8
		{"word with underscore", "foo_bar", 11}, // f(4)+o(1)+o(1)+b(3)+a(1)+r(1)=11, but _ ignored

		// Tie scenarios
		{"drawing", "drawing", 12},
		{"illustration", "illustration", 12},

		// Winner comparison words
		{"COMPUTER", "COMPUTER", 14},
		{"science", "science", 11},
		{"Scrabble", "Scrabble", 14},
		{"wiNNeR", "wiNNeR", 9},
		{"pig", "pig", 6},
		{"dog", "dog", 5},
		{"Skating!", "Skating!", 12},
		{"figure?", "figure?", 10},

		// Edge cases
		{"empty string", "", 0},
		{"only punctuation", "!@#$%", 0},
		{"numbers and letters", "test123", 4},
		{"spaces", "a b c", 7},
		{"non-English letters", "façade", 9}, // f(4)+a(1)+a(1)+d(2)+e(1), ç ignored
		{"emoji", "hello🙂", 8}, // h(4)+e(1)+l(1)+l(1)+o(1), emoji ignored
		{"very long word", "supercalifragilisticexpialidocious", 56},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeScore(tt.word)
			if result != tt.expected {
				t.Errorf("computeScore(%q) = %d, want %d", tt.word, result, tt.expected)
			}
		})
	}
}

// TestScrabbleScenarios covers two-player comparisons and ties.
func TestScrabbleScenarios(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		winner string // "player1", "player2", or "tie"
	}{
		{"Question tie", "Question?", "Question!", "tie"},
		{"drawing illustration tie", "drawing", "illustration", "tie"},
		{"hai beats Oh", "hai!", "Oh,", "player1"},
		{"COMPUTER beats science", "COMPUTER", "science", "player1"},
		{"Scrabble beats wiNNeR", "Scrabble", "wiNNeR", "player1"},
		{"pig beats dog", "pig", "dog", "player1"},
		{"Skating beats figure", "Skating!", "figure?", "player1"},
		{"dog beats pig", "dog", "pig", "player2"},
		{"empty vs nonempty", "", "a", "player2"},
		{"punctuation tie", "!!!", "???", "tie"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score1 := computeScore(tt.word1)
			score2 := computeScore(tt.word2)

			winnerMsg := determineWinner(score1, score2)
			var result string
			switch winnerMsg {
			case "Player 1 wins!":
				result = "player1"
			case "Player 2 wins!":
				result = "player2"
			case "Tie!":
				result = "tie"
			default:
				result = "unknown"
			}

			if result != tt.winner {
				t.Errorf("Game %q vs %q: got %s (scores: %d vs %d), want %s",
					tt.word1, tt.word2, result, score1, score2, tt.winner)
			}
		})
	}
}

// TestDetermineWinner directly tests the determineWinner function.
func TestDetermineWinner(t *testing.T) {
	tests := []struct {
		score1  int
		score2  int
		want    string
	}{
		{10, 5, "Player 1 wins!"},
		{5, 10, "Player 2 wins!"},
		{7, 7, "Tie!"},
		{0, 0, "Tie!"},
		{100, 99, "Player 1 wins!"},
		{99, 100, "Player 2 wins!"},
	}

	for _, tt := range tests {
		got := determineWinner(tt.score1, tt.score2)
		if got != tt.want {
			t.Errorf("determineWinner(%d, %d) = %q, want %q", tt.score1, tt.score2, got, tt.want)
		}
	}
}

