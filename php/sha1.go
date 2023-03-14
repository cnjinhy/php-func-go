package php

import (
	"crypto/sha1"
	"fmt"
)

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1Hash := h.Sum(nil)
	return fmt.Sprintf("%x", sha1Hash)
}
