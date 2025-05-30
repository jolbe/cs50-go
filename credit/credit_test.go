package main

// Test suite for credit card detection functionality
// Covers DetectCardType and all helper functions (TrimNumber, NumDigits, IsValidLuhn)
// Based on CS50 problem requirements for detecting AMEX, MASTERCARD, VISA, and INVALID cards

/*
CS50 test cases reference:
:) credit.c exists
:) credit.c compiles
:) identifies 378282246310005 as AMEX
:) identifies 371449635398431 as AMEX
:) identifies 5555555555554444 as MASTERCARD
:) identifies 5105105105105100 as MASTERCARD
:) identifies 4111111111111111 as VISA
:) identifies 4012888888881881 as VISA
:) identifies 4222222222222 as VISA
:) identifies 1234567890 as INVALID (invalid length, checksum, identifying digits)
:) identifies 369421438430814 as INVALID (invalid identifying digits)
:) identifies 4062901840 as INVALID (invalid length)
:) identifies 5673598276138003 as INVALID (invalid identifying digits)
:) identifies 4111111111111113 as INVALID (invalid checksum)
:) identifies 4222222222223 as INVALID (invalid checksum)
:) identifies 3400000000000620 as INVALID (AMEX identifying digits, VISA/Mastercard length)
:) identifies 430000000000000 as INVALID (VISA identifying digits, AMEX length)
*/

import (
	"testing"
)

func TestDetectCardType(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		number uint64
		want   string
	}{
		// AMEX test cases
		{"AMEX - 378282246310005", 378282246310005, "AMEX"},
		{"AMEX - 371449635398431", 371449635398431, "AMEX"},

		// MASTERCARD test cases
		{"MASTERCARD - 5555555555554444", 5555555555554444, "MASTERCARD"},
		{"MASTERCARD - 5105105105105100", 5105105105105100, "MASTERCARD"},

		// VISA test cases
		{"VISA - 4111111111111111", 4111111111111111, "VISA"},
		{"VISA - 4012888888881881", 4012888888881881, "VISA"},
		{"VISA - 4222222222222", 4222222222222, "VISA"},

		// INVALID test cases
		{"INVALID - short number", 1234567890, "INVALID"},
		{"INVALID - wrong identifying digits", 369421438430814, "INVALID"},
		{"INVALID - wrong length", 4062901840, "INVALID"},
		{"INVALID - wrong identifying digits mastercard", 5673598276138003, "INVALID"},
		{"INVALID - wrong checksum visa", 4111111111111113, "INVALID"},
		{"INVALID - wrong checksum visa short", 4222222222223, "INVALID"},
		{"INVALID - amex digits visa length", 3400000000000620, "INVALID"},
		{"INVALID - visa digits amex length", 430000000000000, "INVALID"},

		// Additional edge cases
		{"INVALID - zero", 0, "INVALID"},
		{"INVALID - single digit", 1, "INVALID"},
		{"INVALID - mastercard wrong range low", 5055555555554444, "INVALID"},
		{"INVALID - mastercard wrong range high", 5655555555554444, "INVALID"},
		{"INVALID - visa wrong length 14", 41111111111111, "INVALID"},
		{"INVALID - visa wrong length 15", 411111111111111, "INVALID"},
		{"INVALID - visa wrong length 17", 41111111111111111, "INVALID"},
		{"INVALID - amex wrong length 14", 37828224631000, "INVALID"},
		{"INVALID - amex wrong length 16", 3782822463100050, "INVALID"},
		{"INVALID - amex wrong prefix 35", 351449635398431, "INVALID"},
		{"INVALID - amex wrong prefix 38", 381449635398431, "INVALID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectCardType(tt.number)
			if got != tt.want {
				t.Errorf("DetectCardType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimNumber(t *testing.T) {
	tests := []struct {
		name   string
		number uint64
		digits int
		want   uint64
	}{
		{"Trim 2 digits from 16-digit number", 4111111111111111, 2, 41},
		{"Trim 1 digit from 16-digit number", 4111111111111111, 1, 4},
		{"Trim 2 digits from 15-digit number", 378282246310005, 2, 37},
		{"Trim more digits than available", 123, 5, 123},
		{"Trim same number of digits", 123, 3, 123},
		{"Trim from single digit", 5, 1, 5},
		{"Trim zero digits", 123456, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimNumber(tt.number, tt.digits)
			if got != tt.want {
				t.Errorf("TrimNumber(%d, %d) = %d, want %d", tt.number, tt.digits, got, tt.want)
			}
		})
	}
}

func TestNumDigits(t *testing.T) {
	tests := []struct {
		name   string
		number uint64
		want   int
	}{
		{"Single digit", 5, 1},
		{"Two digits", 42, 2},
		{"Three digits", 123, 3},
		{"16-digit VISA", 4111111111111111, 16},
		{"15-digit AMEX", 378282246310005, 15},
		{"13-digit VISA", 4222222222222, 13},
		{"Zero", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumDigits(tt.number)
			if got != tt.want {
				t.Errorf("NumDigits(%d) = %d, want %d", tt.number, got, tt.want)
			}
		})
	}
}

func TestIsValidLuhn(t *testing.T) {
	tests := []struct {
		name   string
		number uint64
		want   bool
	}{
		// Valid card numbers
		{"Valid AMEX", 378282246310005, true},
		{"Valid AMEX 2", 371449635398431, true},
		{"Valid MASTERCARD", 5555555555554444, true},
		{"Valid MASTERCARD 2", 5105105105105100, true},
		{"Valid VISA 16", 4111111111111111, true},
		{"Valid VISA 16 2", 4012888888881881, true},
		{"Valid VISA 13", 4222222222222, true},

		// Invalid card numbers (wrong checksum)
		{"Invalid VISA checksum", 4111111111111113, false},
		{"Invalid VISA 13 checksum", 4222222222223, false},

		// Simple test cases
		{"Simple valid", 1234567890123452, true},    // This should pass Luhn
		{"Simple invalid", 1234567890123456, false}, // This should fail Luhn
		{"Single digit", 0, true},                   // 0 passes Luhn algorithm
		{"Single digit 5", 5, false},                // 5 fails Luhn algorithm
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidLuhn(tt.number)
			if got != tt.want {
				t.Errorf("IsValidLuhn(%d) = %t, want %t", tt.number, got, tt.want)
			}
		})
	}
}
