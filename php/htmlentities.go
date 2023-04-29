package php

import "html"

func HtmlEntities(s string) string {
	return html.EscapeString(s)
}
