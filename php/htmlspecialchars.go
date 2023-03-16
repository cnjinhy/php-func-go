package php

import "strings"

func HtmlSpecialChars(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"\"", "&quot;",
		"'", "&#039;",
		"<", "&lt;",
		">", "&gt;",
	)
	return replacer.Replace(s)
}
