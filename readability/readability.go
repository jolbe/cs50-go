package main

import (
	"cs50-go/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	// Prompt the user for input
	text := utils.GetString("Text: ")

	// Count the number of letters, words and sentences
	letters := CountLetters(text)
	words := CountWords(text)
	sentences := CountSentences(text)

	// Compute the Coleman-Liau index
	index := ComputeColemanLiau(letters, words, sentences)

	// Print the grade level
	grade := GetGrade(index)
	fmt.Println(grade)
}

func CountLetters(text string) int {
	letters := 0
	for _, ch := range text {
		if utils.IsAsciiLetter(ch) {
			letters++
		}
	}
	return letters
}

func CountWords(text string) int {
	return len(strings.Fields(text))
}

func CountSentences(text string) int {
	sentences := 0
	for _, ch := range text {
		if ch == '.' || ch == '!' || ch == '?' {
			sentences++
		}
	}
	return sentences
}

func ComputeColemanLiau(letters, words, sentences int) float64 {
	L := float64(letters) / float64(words) * 100
	S := float64(sentences) / float64(words) * 100
	return 0.0588*L - 0.296*S - 15.8
}

func GetGrade(index float64) string {
	grade := int(math.Round(index))

	switch {
	case grade < 1:
		return "Before Grade 1"
	case grade >= 16:
		return "Grade 16+"
	default:
		return fmt.Sprintf("Grade %d", grade)
	}
}
