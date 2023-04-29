package php

import "strings"

func HtmlSpecialCharsDecode(s string) string {
	return strings.NewReplacer(
		"&amp;", "&",
		"&lt;", "<",
		"&gt;", ">",
		"&quot;", `"`,
		"&#34;", `"`,
		"&#39;", "'",
		"&qpos;", "'",
	).Replace(s)
}
