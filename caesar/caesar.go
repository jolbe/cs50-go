package main

import (
	"cs50-go/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	key := parseKeyArg(os.Args)

	plaintext := utils.GetString("plaintext:  ")
	ciphertext := Encrypt(plaintext, key)

	fmt.Println("ciphertext:", ciphertext)
}

func parseKeyArg(args []string) int {
	if len(args) != 2 {
		printUsageAndExit(args[0])
	}

	key, err := strconv.Atoi(args[1])
	if err != nil {
		printUsageAndExit(args[0])
	}

	return key
}

func printUsageAndExit(progName string) {
	fmt.Printf("Usage: %s key\n", progName)
	os.Exit(1)
}

func Encrypt(plaintext string, key int) string {
	var builder strings.Builder
	builder.Grow(len(plaintext))

	for _, ch := range plaintext {
		if utils.IsAsciiLetter(ch) {
			builder.WriteRune(Rotate(ch, key))
		} else {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}

func Rotate(ch rune, key int) rune {
	var base rune
	switch {
	case unicode.IsUpper(ch):
		base = 'A'
	case unicode.IsLower(ch):
		base = 'a'
	default:
		return ch
	}

	// Normalize key to range [-25, 25] with wrap-around support.
	key = key % 26

	alphaPos := int(ch) - int(base)
	offset := (alphaPos + key) % 26
	if offset < 0 {
		offset += 26
	}

	return base + rune(offset)
}
