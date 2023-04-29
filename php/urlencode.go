package php

import "net/url"

func Urlencode(str string) string {
	return url.QueryEscape(str)
}
