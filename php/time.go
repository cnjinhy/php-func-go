package php

import (
	"time"
)

//在 Golang 中，我们可以使用 time 包来实现类似 PHP 中的 time 函数。
func Time() int64 {
	return time.Now().Unix()
}
