package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		return strings.TrimSpace(input)
	}
}

func GetString(prompt string) string {
	return getInput(prompt)
}

func GetInt(prompt string) int {
	for {
		input := getInput(prompt)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid integer. Please try again.")
			continue
		}
		return number
	}
}

func GetUint64(prompt string) uint64 {
	for {
		input := getInput(prompt)
		number, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			fmt.Println("Invalid integer. Please try again.")
			continue
		}
		return uint64(number)
	}
}

func IsAsciiLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
