package php

import (
	"net/url"
	"strings"
)

func Rawurlencode(str string) string {
	return strings.ReplaceAll(url.QueryEscape(str), "+", "%20")
}
