package php

import "strings"

func Ucwords(s string) string {
	return strings.Title(strings.ToLower(s))
}
