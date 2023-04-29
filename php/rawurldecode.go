package php

import (
	"net/url"
	"strings"
)

func Rawurldecode(str string) (string, error) {
	return url.QueryUnescape(strings.ReplaceAll(str, "%20", "+"))
}
