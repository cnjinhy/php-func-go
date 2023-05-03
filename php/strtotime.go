package php

import (
	fun "github.com/x-funs/go-fun"
)

func StrToTime(args ...any) int64 {
	if len(args) == 0 {
		return fun.StrToTime()
	}
	if len(args) == 1 {
		return fun.StrToTime(args[0])
	}
	if len(args) == 2 {
		return fun.StrToTime(args[0], args[1])
	}
	return 0
}
