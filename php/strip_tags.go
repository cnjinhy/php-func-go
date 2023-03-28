package php

import (
	strip "github.com/grokify/html-strip-tags-go"
)

func StripTags(s string) string {
	return strip.StripTags(s)
}
