package php

import "unicode"

func Lcfirst(str string) string {
	if len(str) == 0 {
		return str
	}
	r := []rune(str)
	if unicode.IsLower(r[0]) {
		return str
	}
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
