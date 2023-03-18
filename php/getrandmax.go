package php

import (
	"math"
)

//由于math/rand包中生成的伪随机数是基于算法计算的，因此在特定的环境下，生成的随机数的最大值可能会是int32。
//On a stock installation as of version 7.4.33 / 8.1.12, this function seems to return 2147483647 (maximum 32-bit integer).
func Getrandmax() int {
	return math.MaxInt32
}
