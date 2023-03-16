package php

import (
	"fmt"
	"strconv"
)

func Echo(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			fmt.Print(v)
		case int:
			fmt.Print(strconv.Itoa(v))
		case float64:
			fmt.Print(strconv.FormatFloat(v, 'f', -1, 64))
		default:
			fmt.Print(v)
		}
	}
	fmt.Print("\n")
}
