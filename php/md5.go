package php

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5(args ...interface{}) string {
	switch len(args) {
	case 0:
		panic("md5() expects at least 1 parameter, 0 given")
	case 1:
		str, ok := args[0].(string)
		if !ok {
			panic("md5() expects parameter 1 to be string, " + fmt.Sprintf("%T", args[0]) + " given")
		}
		return md5String(str)
	case 2:
		str, ok := args[0].(string)
		if !ok {
			panic("md5() expects parameter 1 to be string, " + fmt.Sprintf("%T", args[0]) + " given")
		}
		raw, ok := args[1].(bool)
		if !ok {
			panic("md5() expects parameter 2 to be bool, " + fmt.Sprintf("%T", args[1]) + " given")
		}
		if raw {
			return md5StringRaw(str)
		}
		return md5String(str)
	default:
		panic("md5() expects at most 2 parameters, " + fmt.Sprintf("%d", len(args)) + " given")
	}
}

func md5String(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func md5StringRaw(input string) string {
	hash := md5.Sum([]byte(input))
	return string(hash[:])
}
