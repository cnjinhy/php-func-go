package php

import (
	"path/filepath"
	"strings"
)

func Basename(path string, suffix ...string) string {
	// Use the filepath.Base function to get the last element of the path.
	basename := filepath.Base(path)

	// If a suffix is provided, remove it from the basename.
	if len(suffix) > 0 && suffix[0] != "" {
		s := suffix[0]
		if strings.HasSuffix(basename, s) {
			basename = basename[:len(basename)-len(s)]
		}
	}

	return basename
}
