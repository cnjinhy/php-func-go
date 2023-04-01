package php

func MbStrlen(str string) int {
	var (
		runes       = []rune(str)
		runesLength = len(runes)
	)
	return runesLength
}
