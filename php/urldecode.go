package php

import "net/url"

func Urldecode(str string) string {
	res, _ := url.QueryUnescape(str)
	return res
}
