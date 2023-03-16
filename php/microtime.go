package php

import (
	"fmt"
	"math"
	"time"
)

//在 Golang 中，可以使用 time 包和 math 包来实现类似于 PHP 中的 microtime 函数。
//对于 PHP 中的 microtime 函数，其参数是可选的。如果不传递参数，则返回一个字符串，其中包含当前时间的微秒数和秒数；如果传递 true 参数，则返回当前时间的浮点数表示，精确到微秒。
func Microtime(getAsFloat ...bool) interface{} {
	now := time.Now()
	sec := float64(now.Unix())
	usec := float64(now.Nanosecond()) / 1000000.0
	result := sec + usec

	if len(getAsFloat) > 0 && getAsFloat[0] {
		return result
	}
	return fmt.Sprintf("%0.8f %d", math.Floor(result), int64(usec))
}
