package php

func Ord(s string) int {
	// If s is an empty string, return 0.
	if s == "" {
		return 0
	}

	// Extract the first character of the string and convert it to a rune.
	r := []rune(s)[0]

	// If r is a surrogate code point, return 0.
	if r >= 0xD800 && r <= 0xDFFF {
		return 0
	}

	// Return the Unicode code point of the character.
	return int(r)
}
