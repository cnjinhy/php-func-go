package php

import "strings"

func Strpos(haystack string, needle string) interface{} {
	index := strings.Index(haystack, needle)
	if index == -1 {
		return false
	} else {
		return index
	}
}
