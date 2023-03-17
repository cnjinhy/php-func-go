package php

import "strings"

func SubstrCount(s, sub string, args ...int) int {
	start, length := 0, 0
	switch len(args) {
	case 1:
		start = args[0]
	case 2:
		start, length = args[0], args[1]
	}
	if start < 0 {
		start = len(s) + start
	}
	if length == 0 {
		length = len(s) - start
	}
	count := 0
	for i := start; i < len(s) && i < start+length; i++ {
		if strings.HasPrefix(s[i:], sub) {
			count++
			i += len(sub) - 1
		}
	}
	return count
}
