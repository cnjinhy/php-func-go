package php

func MbSubStr(str string, start int, length ...int) string {
	var (
		runes       = []rune(str)
		runesLength = len(runes)
	)
	strLength := runesLength
	if start < 0 {
		if -start > strLength {
			start = 0
		} else {
			start = strLength + start
		}
	} else if start > strLength {
		return ""
	}
	realLength := 0
	if len(length) > 0 {
		realLength = length[0]
		if realLength < 0 {
			if -realLength > strLength-start {
				realLength = 0
			} else {
				realLength = strLength - start + realLength
			}
		} else if realLength > strLength-start {
			realLength = strLength - start
		}
	} else {
		realLength = strLength - start
	}
	end := start + realLength
	if end > runesLength {
		end = runesLength
	}
	return string(runes[start:end])
}
