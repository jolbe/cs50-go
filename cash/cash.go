package main

import (
	"cs50-go/utils"
	"fmt"
)

func main() {
	cents := getChange()

	quarters := calculateQuarters(cents)
	cents -= (quarters * 25)

	dimes := calculateDimes(cents)
	cents -= (dimes * 10)

	nickels := cents / 5
	cents -= (nickels * 5)

	pennies := cents

	sum := quarters + dimes + nickels + pennies

	fmt.Println(sum)
}

func calculateDimes(cents int) int {
	return cents / 10
}

func calculateQuarters(cents int) int {
	quarters := 0

	for cents >= 25 {
		quarters++
		cents -= 25
	}

	return quarters
}

func getChange() int {
	for {
		change := utils.GetInt("Change owed: ")
		if change >= 0 {
			return change
		}
	}
}
