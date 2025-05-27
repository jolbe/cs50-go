package main

import (
	"fmt"

	"cs50-go/utils"
)

func main() {
	change := GetChange("Change owed: ", IntReaderFunc(utils.GetInt))

	coinsRecursive := CountCoinsRecursive(change)
	fmt.Println(coinsRecursive)

	coins := CountCoins(change)
	fmt.Println(coins)
}

type IntReader interface {
	GetInt(prompt string) int
}

type IntReaderFunc func(prompt string) int

func (r IntReaderFunc) GetInt(prompt string) int {
	return r(prompt)
}

func GetChange(prompt string, reader IntReader) int {
	for {
		change := reader.GetInt(prompt)
		if change >= 0 {
			return change
		}
	}
}

func CountCoins(cents int) int {
	quarters := calculateQuarters(cents)
	cents -= (quarters * 25)

	dimes := cents / 10
	cents -= (dimes * 10)

	nickels := cents / 5
	cents -= (nickels * 5)

	pennies := cents

	return quarters + dimes + nickels + pennies
}

func calculateQuarters(cents int) int {
	quarters := 0
	for cents >= 25 {
		quarters++
		cents -= 25
	}
	return quarters
}

func CountCoinsRecursive(cents int) int {
	if cents == 0 {
		return 0
	}

	var coin int
	if cents >= 25 {
		coin = 25
	} else if cents >= 10 {
		coin = 10
	} else if cents >= 5 {
		coin = 5
	} else {
		coin = 1
	}

	return 1 + CountCoinsRecursive(cents-coin)
}
