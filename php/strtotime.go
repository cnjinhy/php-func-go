package php

import (
	fun "github.com/x-funs/go-fun"
)

// Strtotime 将日期字符串转换为时间戳（秒），支持各种格式的日期字符串和相对时间字符串
// timezone 为空时，自动根据系统或者Go的环境识别时区
func StrToTime(args ...any) int64 {
	return fun.StrToTime(args)
}
