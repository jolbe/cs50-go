package main

import (
	"cs50-go/utils"
	"fmt"
	"unicode"
)

var points = [...]int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

func main() {
	// Prompt the user for two words
	word1 := utils.GetString("Player 1: ")
	word2 := utils.GetString("Player 2: ")

	// Compute the score of each word
	score1 := computeScore(word1)
	score2 := computeScore(word2)

	// Print the winner
	winner := determineWinner(score1, score2)
	fmt.Println(winner)
}

func computeScore(word string) int {
	score := 0
	for _, r := range word {

		ch := unicode.ToUpper(r)
		// if ch >= 'A' && ch <= 'Z' {
		if utils.IsAsciiLetter(ch) {
			pos := ch - 'A'
			score += points[pos]
		}
	}

	return score
}

func determineWinner(score1, score2 int) string {
	switch {
	case score1 > score2:
		return "Player 1 wins!"
	case score2 > score1:
		return "Player 2 wins!"

	default:
		return "Tie!"
	}
}
