package php

import "strings"

func Trim(str string, charlist ...string) string {
	var charsToRemove string
	if len(charlist) == 0 {
		charsToRemove = " \t\n\r\x00\x0B"
	} else {
		charsToRemove = charlist[0]
	}
	return strings.Trim(str, charsToRemove)
}
