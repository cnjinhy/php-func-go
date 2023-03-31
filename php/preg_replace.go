package php

import "regexp"

func PregReplace(pattern string, replacement string, subject string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllStringFunc(subject, func(s string) string {
		return re.ReplaceAllString(s, replacement)
	})
}
