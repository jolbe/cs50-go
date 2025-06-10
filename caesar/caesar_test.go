package main

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		letter rune
		key    int
		want   rune
	}{
		{'A', 1, 'B'},
		{'A', 25, 'Z'},
		{'A', 26, 'A'}, // wrap around
		{'Z', 1, 'A'},  // wrap around
		{'a', 1, 'b'},
		{'a', 25, 'z'},
		{'a', 26, 'a'}, // wrap around
		{'z', 1, 'a'},  // wrap around
		{'!', 5, '!'},  // non-lette{'A', -1, 'Z'},    // left shift
		{'B', -2, 'Z'}, // left shift with wrap
		{'a', -1, 'z'}, // left shift lowercase
		{'b', -2, 'z'}, // left shift lowercase with wrapr should stay the same
	}

	for _, tt := range tests {
		got := Rotate(tt.letter, tt.key)
		if got != tt.want {
			t.Errorf("rotate(%q, %d) = %q; want %q", tt.letter, tt.key, got, tt.want)
		}
	}
}

func TestEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		key       int
		want      string
	}{
		{"HELLO", 3, "KHOOR"},
		{"WORLD", 5, "BTWQI"},
		{"hello", 3, "khoor"},
		{"abc xyz", 2, "cde zab"},
		{"CS50!", 1, "DT50!"},
		{"", 4, ""}, // empty string
	}

	for _, tt := range tests {
		got := Encrypt(tt.plaintext, tt.key)
		if got != tt.want {
			t.Errorf("encrypt(%q, %d) = %q; want %q", tt.plaintext, tt.key, got, tt.want)
		}
	}
}
