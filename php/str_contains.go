package php

import "strings"

func StrContains(haystack string, needle string) bool {
	return strings.Contains(haystack, needle)
}
