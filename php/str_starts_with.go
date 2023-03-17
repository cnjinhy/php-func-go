package php

import "strings"

func StrStartsWith(haystack string, needle string) bool {
	return strings.HasPrefix(haystack, needle)
}
