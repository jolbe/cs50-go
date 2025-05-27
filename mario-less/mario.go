package main

import (
	"fmt"

	"cs50-go/utils"
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
	for range spaces {
		fmt.Print(" ")
	}

	for range bricks {
		fmt.Print("#")
	}

	fmt.Println()
}
