package php

func MbStrlen(str string) (strlen int) {
	var (
		runes       = []rune(str)
		runesLength = len(runes)
	)
	return runesLength
}
