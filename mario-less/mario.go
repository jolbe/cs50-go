package main

import (
	"cs50-go/utils"
	"fmt"
)

func main() {
	height := getHeight()

	for i := 0; i < height; i++ {
		printRow(height-i-1, i+1)
	}
}

func getHeight() int {
	for {
		height := utils.GetInt("Height? ")
		if height > 0 {
			return height
		}
	}
}

func printRow(spaces int, bricks int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}

	for i := 0; i < bricks; i++ {
		fmt.Print("#")
	}

	fmt.Println()
}
