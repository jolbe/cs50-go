package main

import (
	"fmt"

	"cs50-go/utils"
)

func main() {
	height := getHeigth("Height: ", 42)

	for i := 0; i < height; i++ {
		rightAlignedPyramid(height-i-1, i+1)
		printChars(" ", 2)
		leftAlignedPyramid(height-i-1, i+1)
		fmt.Println()
	}
}

func leftAlignedPyramid(spaces, bricks int) {
	printChars("#", bricks)
	printChars(" ", spaces)
}

func rightAlignedPyramid(spaces, bricks int) {
	printChars(" ", spaces)
	printChars("#", bricks)
}

func printChars(ch string, n int) {
	for range n {
		fmt.Print(ch)
	}
}

func getHeigth(prompt string, max int) int {
	for {
		n := utils.GetInt(prompt)
		if n > 0 && n <= max {
			return n
		}
	}
}
