package php

import "strings"

func Strstr(s, substr string, options ...bool) string {
	beforeNeedle := false
	if len(options) > 0 {
		beforeNeedle = options[0]
	}
	index := strings.Index(s, substr)
	if index == -1 {
		return ""
	}
	if beforeNeedle {
		return s[:index]
	}
	return s[index:]
}
