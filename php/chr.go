package php

import "unicode/utf8"

func Chr(codepoint int) string {
	// If codepoint is less than 0 or greater than 0x10FFFF, return an empty string.
	if codepoint < 0 || codepoint > 0x10FFFF {
		return ""
	}

	// If codepoint is between 0xD800 and 0xDFFF, return an empty string.
	if codepoint >= 0xD800 && codepoint <= 0xDFFF {
		return ""
	}

	// If codepoint is greater than 0x7F and the system uses ASCII as the default charset, return an empty string.
	if codepoint > 0x7F && utf8.RuneLen(rune(codepoint)) == 1 {
		return ""
	}

	// Return the single character string representing the Unicode codepoint.
	return string(codepoint)
}
