package php

import "net/url"

func UrlDecode(str string) string {
	res, _ := url.QueryUnescape(str)
	return res
}
