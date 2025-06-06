package main

import (
	"math"
	"testing"
)

func TestCountLetters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Only letters",
			input:    "Hello",
			expected: 5,
		},
		{
			name:     "Letters with spaces",
			input:    "Hello World",
			expected: 10,
		},
		{
			name:     "Letters with punctuation",
			input:    "Hello, World!",
			expected: 10,
		},
		{
			name:     "Mixed case",
			input:    "HeLLo WoRld",
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountLetters(tt.input)
			if result != tt.expected {
				t.Errorf("CountLetters(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Single word",
			input:    "Hello",
			expected: 1,
		},
		{
			name:     "Multiple words",
			input:    "Hello World",
			expected: 2,
		},
		{
			name:     "Multiple spaces",
			input:    "Hello   World",
			expected: 2,
		},
		{
			name:     "Words with punctuation",
			input:    "Hello, World!",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.input)
			if result != tt.expected {
				t.Errorf("CountWords(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountSentences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "No sentences",
			input:    "Hello World",
			expected: 0,
		},
		{
			name:     "Single sentence",
			input:    "Hello World.",
			expected: 1,
		},
		{
			name:     "Multiple sentences",
			input:    "Hello. World! How are you?",
			expected: 3,
		},
		// {
		// 	name:     "Multiple punctuation",
		// 	input:    "Hello... World!?!",
		// 	expected: 2,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountSentences(tt.input)
			if result != tt.expected {
				t.Errorf("CountSentences(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestComputeColemanLiau(t *testing.T) {
	tests := []struct {
		name      string
		letters   int
		words     int
		sentences int
		expected  float64
	}{
		{
			name:      "Grade 1 example",
			letters:   336,
			words:     100,
			sentences: 10,
			expected:  1.0,
		},
		{
			name:      "Grade 2 example",
			letters:   343,
			words:     100,
			sentences: 8,
			expected:  2.0,
		},
		{
			name:      "Grade 3 example",
			letters:   355,
			words:     100,
			sentences: 7,
			expected:  3.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ComputeColemanLiau(tt.letters, tt.words, tt.sentences)
			if math.Abs(result-tt.expected) > 0.1 {
				t.Errorf("ComputeColemanLiau(%d, %d, %d) = %.2f; want %.2f",
					tt.letters, tt.words, tt.sentences, result, tt.expected)
			}
		})
	}
}

func TestGetGrade(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{
			name:     "Before Grade 1",
			input:    0.49,
			expected: "Before Grade 1",
		},
		{
			name:     "Grade 1",
			input:    1.0,
			expected: "Grade 1",
		},
		{
			name:     "Grade 5",
			input:    5.0,
			expected: "Grade 5",
		},
		{
			name:     "Grade 16+",
			input:    16.0,
			expected: "Grade 16+",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetGrade(tt.input)
			if result != tt.expected {
				t.Errorf("GetGrade(%.1f) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReadingLevels(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "Before Grade 1 - Simple sentences",
			text:     "One fish. Two fish. Red fish. Blue fish.",
			expected: "Before Grade 1",
		},
		{
			name:     "Grade 2 - Questions and statements",
			text:     "Would you like them here or there? I would not like them here or there. I would not like them anywhere.",
			expected: "Grade 2",
		},
		{
			name:     "Grade 3 - Exclamations and statements",
			text:     "Congratulations! Today is your day. You're off to Great Places! You're off and away!",
			expected: "Grade 3",
		},
		{
			name:     "Grade 5 - Complex narrative",
			text:     "Harry Potter was a highly unusual boy in many ways. For one thing, he hated the summer holidays more than any other time of year. For another, he really wanted to do his homework, but was forced to do it in secret, in the dead of the night. And he also happened to be a wizard.",
			expected: "Grade 5",
		},
		{
			name:     "Grade 7 - Complex sentence",
			text:     "In my younger and more vulnerable years my father gave me some advice that I've been turning over in my mind ever since.",
			expected: "Grade 7",
		},
		{
			name:     "Grade 8 - Complex narrative with dialogue",
			text:     "Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, \"and what is the use of a book,\" thought Alice \"without pictures or conversation?\"",
			expected: "Grade 8",
		},
		{
			name:     "Grade 8 - Detailed description",
			text:     "When he was nearly thirteen, my brother Jem got his arm badly broken at the elbow. When it healed, and Jem's fears of never being able to play football were assuaged, he was seldom self-conscious about his injury. His left arm was somewhat shorter than his right; when he stood or walked, the back of his hand was at right angles to his body, his thumb parallel to his thigh.",
			expected: "Grade 8",
		},
		{
			name:     "Grade 9 - Philosophical quote",
			text:     "There are more things in Heaven and Earth, Horatio, than are dreamt of in your philosophy.",
			expected: "Grade 9",
		},
		{
			name:     "Grade 10 - Complex narrative",
			text:     "It was a bright cold day in April, and the clocks were striking thirteen. Winston Smith, his chin nuzzled into his breast in an effort to escape the vile wind, slipped quickly through the glass doors of Victory Mansions, though not quickly enough to prevent a swirl of gritty dust from entering along with him.",
			expected: "Grade 10",
		},
		{
			name:     "Grade 16+ - Technical text",
			text:     "A large class of computational problems involve the determination of properties of graphs, digraphs, integers, arrays of integers, finite families of finite sets, boolean formulas and elements of other countable domains.",
			expected: "Grade 16+",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			letters := CountLetters(tt.text)
			words := CountWords(tt.text)
			sentences := CountSentences(tt.text)
			index := ComputeColemanLiau(letters, words, sentences)
			result := GetGrade(index)

			if result != tt.expected {
				t.Errorf("Text: %q\nLetters: %d, Words: %d, Sentences: %d\nIndex: %.2f\nGot: %q, Want: %q",
					tt.text, letters, words, sentences, index, result, tt.expected)
			}
		})
	}
}
