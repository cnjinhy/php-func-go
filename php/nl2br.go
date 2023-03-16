package php

import "strings"

func Nl2Br(s string, useXHTML ...bool) string {
	xhtml := true
	if len(useXHTML) > 0 {
		xhtml = useXHTML[0]
	}

	var lineBreak string
	if xhtml {
		lineBreak = "<br />"
	} else {
		lineBreak = "<br>"
	}
	return strings.ReplaceAll(s, "\n", lineBreak+"\n")
}
