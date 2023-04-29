package php

import "net/url"

func Decode(str string) string {
	res, _ := url.QueryUnescape(str)
	return res
}
