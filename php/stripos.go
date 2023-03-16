package php

import "strings"

func Stripos(haystack string, needle string) interface{} {
	haystack = StrToLower(haystack)
	needle = StrToLower(needle)
	index := strings.Index(haystack, needle)
	if index == -1 {
		return false
	} else {
		return index
	}
}
