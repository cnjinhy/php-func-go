package php

import (
	"path/filepath"
	"strings"
)

func Dirname(path string, levels ...int) string {
	if len(levels) == 0 {
		levels = []int{1}
	}
	dir := filepath.Dir(path)
	for i := 1; i < levels[0]; i++ {
		dir = filepath.Dir(dir)
	}
	return strings.TrimRight(dir, "/\\")
}
