//go:build !windows
// +build !windows

package php

import (
	"os"
	"runtime"
	"syscall"
)

func FileCTime(path string) int64 {
	osType := runtime.GOOS
	if osType == "windows" {
		fileInfo, _ := os.Stat(path)
		wFileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
		tNanSeconds := wFileSys.CreationTime.Nanoseconds()
		tSec := tNanSeconds / 1e9
		return tSec
	} else {
		fileInfo, _ := os.Stat(path)
		stat_t := fileInfo.Sys().(*syscall.Stat_t)
		tCreate := int64(stat_t.Ctim.Sec)
		return tCreate
	}
}
