package php

import (
	"github.com/gogf/gf/v2/util/gconv"
	"math"
)

func Ceil(x float64) int64 {
	return gconv.Int64(math.Ceil(x))
}
