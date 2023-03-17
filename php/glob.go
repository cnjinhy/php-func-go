package php

import "path/filepath"

//我们使用 Go 中的 filepath.Glob() 函数来搜索与指定模式匹配的文件和目录，并返回一个包含所有匹配项的字符串切片。
//由于 Go 中的 filepath.Glob() 函数不会返回错误，我们忽略了第二个返回值，并假定函数总是会返回匹配项的切片。
//We use the filepath. Glob() function in Go to search for files and directories that match the specified pattern, and return a string slice containing all the matches.
//Since the filepath. Glob() function in Go does not return an error, we ignore the second return value and assume that the function always returns a slice of the matching item.

func Glob(pattern string) []string {
	matches, _ := filepath.Glob(pattern)
	return matches
}
