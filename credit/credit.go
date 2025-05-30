// https://cs50.harvard.edu/x/psets/1/credit/#problem-to-solve
package main

import (
	"cs50-go/utils"
	"fmt"
	"math"
)

func main() {
	number := utils.GetUint64("Number: ")
	fmt.Println(DetectCardType(number))
}

func DetectCardType(number uint64) string {
	// Helper functions
	isAmex := func(number uint64) bool {
		return NumDigits(number) == 15 && (TrimNumber(number, 2) == 34 || TrimNumber(number, 2) == 37) && IsValidLuhn(number)
	}
	isMastercard := func(number uint64) bool {
		return NumDigits(number) == 16 && (TrimNumber(number, 2) >= 51 && TrimNumber(number, 2) <= 55) && IsValidLuhn(number)
	}
	isVisa := func(number uint64) bool {
		return (NumDigits(number) == 13 || NumDigits(number) == 16) && TrimNumber(number, 1) == 4 && IsValidLuhn(number)
	}

	switch {
	case isAmex(number):
		return "AMEX"
	case isMastercard(number):
		return "MASTERCARD"
	case isVisa(number):
		return "VISA"
	default:
		return "INVALID"
	}
}

func TrimNumber(number uint64, digits int) uint64 {
	if NumDigits(number) <= digits {
		return number
	}

	divisor := math.Pow(10, float64(NumDigits(number)-digits))
	return number / uint64(divisor)
}

func NumDigits(number uint64) int {
	digits := 0

	for number > 0 {
		digits++
		number /= 10
	}

	return digits
}

func IsValidLuhn(number uint64) bool {
	sum := 0
	count := 0

	for number > 0 {
		digit := int(number % 10)

		if count%2 == 1 {
			digit *= 2

			if digit > 9 {
				digit = (digit / 10) + (digit % 10)
			}

		}

		sum += digit

		count++
		number /= 10
	}

	return sum%10 == 0
}
