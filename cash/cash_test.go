package main

import (
	"fmt"
	"testing"
)

func TestGetChange(t *testing.T) {
	const prompt = "Change owed: "
	const readerResult = 42

	var receivedPrompt string
	reader := IntReaderFunc(func(prompt string) int {
		receivedPrompt = prompt
		return readerResult
	})

	got := GetChange(prompt, reader)

	if got != readerResult {
		t.Errorf("got %d; want %d", got, readerResult)
	}

	if receivedPrompt != prompt {
		t.Errorf("prompt = %q; want %q", receivedPrompt, prompt)
	}
}

func TestGetChangeRetriesNegativeValues(t *testing.T) {
	inputs := []int{-5, -1, 42}
	calls := 0

	reader := IntReaderFunc(func(prompt string) int {
		input := inputs[calls]
		calls++
		return input
	})

	if got := GetChange("Change owed: ", reader); got != 42 {
		t.Errorf("got %d; want 42", got)
	}

	if calls != 3 {
		t.Errorf("GetInt should be called 3 times but instead got called %d times", calls)
	}
}

func TestCountCoins(t *testing.T) {
	cases := []struct {
		input int
		want  int
	}{
		{0, 0},
		{41, 4},
		{1, 1},
		{5, 1},
		{10, 1},
		{15, 2},
		{25, 1},
		{160, 7},
		{2300, 92},
	}

	t.Run("normal version", func(t *testing.T) {
		for _, tt := range cases {
			t.Run(fmt.Sprintf("CountCoins(%v) = %v", tt.input, tt.want), func(t *testing.T) {
				got := CountCoins(tt.input)
				if got != tt.want {
					t.Errorf("got %v; want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("recursive version", func(t *testing.T) {
		for _, tt := range cases {
			t.Run(fmt.Sprintf("CountCoinsRecursive(%v) = %v", tt.input, tt.want), func(t *testing.T) {
				got := CountCoinsRecursive(tt.input)
				if got != tt.want {
					t.Errorf("got %v; want %v", got, tt.want)
				}
			})
		}
	})
}
