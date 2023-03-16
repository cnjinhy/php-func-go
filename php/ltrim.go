package php

import "strings"

func Ltrim(s string, charsToRemove ...string) string {
	removalSet := "\t\n\v\r "

	if len(charsToRemove) > 0 {
		removalSet = charsToRemove[0]
	}

	return strings.TrimLeftFunc(s, func(r rune) bool {
		return strings.ContainsRune(removalSet, r)
	})
}
