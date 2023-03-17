package php

import "strings"

func Chop(s string, charsToRemove ...string) string {
	removalSet := "\t\n\v\r "
	if len(charsToRemove) > 0 {
		removalSet = charsToRemove[0]
	}
	isRemovalRune := func(r rune) bool {
		for _, c := range removalSet {
			if r == c {
				return true
			}
		}
		return false
	}
	return strings.TrimRightFunc(s, isRemovalRune)
}
