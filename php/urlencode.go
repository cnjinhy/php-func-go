package php

import "net/url"

func UrlEncode(str string) string {
	return url.QueryEscape(str)
}
