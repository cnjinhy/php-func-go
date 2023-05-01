package php

import (
	"github.com/gogf/gf/v2/util/gconv"
)

func Intval(any interface{}) int {
	return gconv.Int(any)
}

func Intval64(any interface{}) int64 {
	return gconv.Int64(any)
}
