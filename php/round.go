package php

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

func Round(num float64, point int) (newNum float64) {
	newNum, _ = strconv.ParseFloat(fmt.Sprintf("%."+gconv.String(point)+"f", num), 64)
	return newNum
}
