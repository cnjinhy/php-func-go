package php

import "strings"

func StrEndsWith(haystack string, needle string) bool {
	return strings.HasSuffix(haystack, needle)
}
