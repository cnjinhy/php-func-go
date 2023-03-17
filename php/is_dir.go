package php

import "os"

//该函数使用了 Go 语言的 os.Stat 函数来获取文件信息，如果出现错误则返回 false，否则判断文件信息是否为目录类型，如果是，则返回 true，否则返回 false。
//需要注意的是，在 Go 中，目录和文件都是用 os.FileInfo 类型来表示的，因此我们需要使用 IsDir 方法来判断文件类型是否为目录。
//同时，该函数的参数必须是一个有效的路径，否则会返回 false。
//This function uses the Go language's os.Stat function to obtain file information. If an error occurs, it returns false. Otherwise, it determines whether the file information is of a directory type. If so, it returns true, and otherwise, it returns false.
//Note that in Go, both directories and files are represented by the os.FileInfo type, so we need to use the IsDir method to determine whether the file type is a directory.
//At the same time, the parameter to this function must be a valid path, otherwise it will return false.
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsDir()
}
