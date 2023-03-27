package php

import "fmt"

func Bin2hex(input string) string {
	var result string
	for i := 0; i < len(input); i++ {
		result += fmt.Sprintf("%02x", input[i])
	}
	return result
}
