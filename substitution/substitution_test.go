package main

import (
	"strings"
	"testing"
)

// Test that parseKeyArg normalizes the key to uppercase
func TestParseKeyArg_NormalizesKeyToUpper(t *testing.T) {
	args := []string{"progname", "QwErTyUiOpAsDfGhJkLzXcVbNm"} // Mixed case
	expected := "QWERTYUIOPASDFGHJKLZXCVBNM"

	key, err := parseKeyArg(args)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if key != expected {
		t.Errorf("Expected normalized key %q, got %q", expected, key)
	}
}

func TestValidateKey(t *testing.T) {
	validKey := "QWERTYUIOPASDFGHJKLZXCVBNM"
	err := validateKey(validKey)
	if err != nil {
		t.Errorf("Expected valid key, got error: %v", err)
	}

	tests := []struct {
		name string
		key  string
	}{
		{"TooShort", "ABCDE"},
		{"TooLong", "ABCDEFGHIJKLMNOPQRSTUVWXYZX"},
		{"NonLetters", "QWERTY1234567890ZXCVBNMASDFGH"},
		{"DuplicateLetters", "AABCDEFGHIJKLMNOPQRSTUVWXY"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateKey(tc.key)
			if err == nil {
				t.Errorf("Expected error for key '%s', got nil", tc.key)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	key := "QWERTYUIOPASDFGHJKLZXCVBNM"

	tests := []struct {
		plaintext string
		expected  string
	}{
		{"HELLO", "ITSSG"},
		{"hello", "itssg"},
		{"Hello, World!", "Itssg, Vgksr!"},
		{"CS50", "EL50"},
		{"AttackAtDawn", "QzzqeaQzRqvf"},
	}

	for _, tt := range tests {
		result := encrypt(tt.plaintext, key)
		if result != tt.expected {
			t.Errorf("encrypt(%q, %q) = %q; want %q", tt.plaintext, key, result, tt.expected)
		}
	}
}

// TestEncryptEdgeCases covers additional edge cases for encrypt.
func TestEncryptEdgeCases(t *testing.T) {
	key := "QWERTYUIOPASDFGHJKLZXCVBNM"

	tests := []struct {
		name      string
		plaintext string
		expected  string
	}{
		{"EmptyString", "", ""},
		{"AllNonLetters", "1234!@#$", "1234!@#$"},
		{"SingleLetterUpper", "A", "Q"},
		{"SingleLetterLower", "z", "m"},
		{"MixedAlphaNum", "abc123XYZ", "qwe123BNM"},
		{"Whitespace", "   ", "   "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := encrypt(tt.plaintext, key)
			if result != tt.expected {
				t.Errorf("encrypt(%q, %q) = %q; want %q", tt.plaintext, key, result, tt.expected)
			}
		})
	}
}

func BenchmarkEncrypt(b *testing.B) {
	key := "QWERTYUIOPASDFGHJKLZXCVBNM"
	plaintext := strings.Repeat("HELLO world! ", 1000)

	for i := 0; i < b.N; i++ {
		_ = encrypt(plaintext, key)
	}
}
