package php

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// md5_file returns the MD5 hash of the contents of the file at the given path.
// If the second parameter is provided, the hash will be appended to it.
func Md5File(path string, prevHash ...string) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()

	h := md5.New()
	if len(prevHash) > 0 {
		io.WriteString(h, prevHash[0])
	}

	if _, err := io.Copy(h, f); err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
