package php

import "os"

//该函数的原理是使用 Golang 的 os 包检查文件是否存在。
//如果文件不存在，则会在检查文件是否可读之前返回 false。否则，它将尝试打开文件以进行读取操作，如果打开失败则返回 false，否则返回 true。
//注意：由于 Golang 和 PHP 的权限模型略有不同，因此在某些情况下可能无法完全模拟 PHP 中 is_readable() 函数的行为。
//The principle of this function is to use Golang's os package to check whether a file exists.
//If the file does not exist, it returns false before checking whether the file is readable. Otherwise, it will attempt to open the file for a read operation, returning false if the opening fails, otherwise returning true.
//Note: Due to the slightly different permission models between Golang and PHP, it may not be possible to fully simulate the is in PHP in some cases_ The behavior of the readable() function.
func IsReadable(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
