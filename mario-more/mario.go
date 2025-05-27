package main

import (
	"cs50-go/utils"
	"fmt"
)

func main() {
	height := getHeigth()

	for i := range height {
		printRow(height-i-1, i+1)
	}
}

func printRow(spaces int, bricks int) {
	// right-sided pyramide
	for range spaces {
		fmt.Print(" ")
	}
	for range bricks {
		fmt.Print("#")
	}

	// middle gap
	fmt.Print("  ")

	// left-sided pyramide
	for range bricks {
		fmt.Print("#")
	}

	// go to next row
	fmt.Println()
}

func getHeigth() int {
	for {
		height := utils.GetInt("Height? ")
		if height >= 1 && height <= 8 {
			return height
		}
	}
}
