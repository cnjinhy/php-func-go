package php

import "os"

//使用 Golang 的 os 包获取文件的信息，包括文件大小，然后返回文件大小。如果文件不存在，则会返回一个错误。
//注意：由于 Golang 和 PHP 的权限模型略有不同，因此在某些情况下可能无法完全模拟 PHP 中 filesize() 函数的行为。
func Filesize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
