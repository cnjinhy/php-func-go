package php

import "strings"

func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}
