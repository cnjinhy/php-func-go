package php

import (
	"io"
	"os"
)

func Copy(src string, dest string) bool {
	srcFile, err := os.Open(src)
	if err != nil {
		return false
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return false
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return false
	}

	err = destFile.Sync()
	if err != nil {
		return false
	}

	return true
}
