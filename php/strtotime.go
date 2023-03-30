package php

import (
	"github.com/araddon/dateparse"
	"time"
)

// Strtotime 将日期字符串转换为时间戳（秒），支持各种格式的日期字符串和相对时间字符串
// timezone 为空时，自动根据系统或者Go的环境识别时区
func StrToTime(dateStr string) int64 {
	t, err := dateparse.ParseIn(dateStr, time.Local) // 解析为指定时区的时间
	if err != nil {
		return 0
	}
	return t.Unix()
}
