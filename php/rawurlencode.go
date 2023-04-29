package php

import (
	"net/url"
	"strings"
)

func RawUrlEncode(str string) string {
	return strings.ReplaceAll(url.QueryEscape(str), "+", "%20")
}
