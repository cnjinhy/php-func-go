package php

import "strings"

func Explode(delimiter string, str string) []string {
	return strings.Split(str, delimiter)
}
