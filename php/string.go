package php

import (
	"fmt"
	"strconv"
	"strings"
)

func StrReplace(s string, old string, new string) string {
	return strings.ReplaceAll(s, old, new)
}

func Echo(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			fmt.Print(v)
		case int:
			fmt.Print(strconv.Itoa(v))
		case float64:
			fmt.Print(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			fmt.Print(v)
		}
	}
	fmt.Print("\n")
}

func SubStr(str string, start int, length ...int) (substr string) {
	strLength := len(str)
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

	if realLength == strLength {
		return str
	} else {
		end := start + realLength
		return str[start:end]
	}
}

func MbSubStr(str string, start int, length ...int) (substr string) {
	// Converting to []rune to support unicode.
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
