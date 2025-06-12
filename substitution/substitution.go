package main

import (
	"cs50-go/utils"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const alphabetLength = 26

func main() {
	key, err := parseKeyArg(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		printUsageAndExit(os.Args[0])
	}

	plaintext := utils.GetString("plaintext:  ")
	ciphertext := encrypt(plaintext, key)
	fmt.Println("ciphertext:", ciphertext)
}

func parseKeyArg(args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("expect exactly one argument, got %d", len(args)-1)
	}

	// Normalize to upper case
	key := strings.ToUpper(args[1])

	if err := validateKey(key); err != nil {
		return "", err
	}

	return key, nil
}

func printUsageAndExit(progName string) {
	fmt.Fprintf(os.Stderr, "Usage: %s key\n", progName)
	fmt.Fprintf(os.Stderr, "Key must be exactly 26 unique letters (A-Z)\n")
	os.Exit(1)
}

func validateKey(key string) error {
	if len(key) != alphabetLength {
		return fmt.Errorf("key must be exactly %d characters long, got %d", alphabetLength, len(key))
	}

	var seen [alphabetLength]bool

	for i, ch := range key {
		if !utils.IsAsciiLetter(ch) {
			return fmt.Errorf("key contains non-letter character '%c' at position %d", ch, i)
		}

		// key is normalized to upper case so it's safe to assume upper case char arithmetic
		index := int(ch - 'A')
		if seen[index] {
			return fmt.Errorf("key contains duplicate letter '%c'", ch)
		}

		seen[index] = true
	}

	return nil
}

func encrypt(plaintext, key string) string {
	if len(plaintext) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(plaintext))

	for _, ch := range plaintext {
		if utils.IsAsciiLetter(ch) {
			// Get the position in alphabet (0-25)
			index := int(unicode.ToUpper(ch) - 'A')

			// Get the cipher character from key (upper case by default because key is normalized to upper case)
			cipherCh := rune(key[index])

			// Perserve original case
			if unicode.IsLower(ch) {
				cipherCh = unicode.ToLower(cipherCh)
			}

			builder.WriteRune(cipherCh)
		} else {
			// Non-alphabetic characters pass through unchanged
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}
