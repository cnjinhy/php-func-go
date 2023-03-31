package php

import "os"

const FILE_APPEND = 8

// 写入文件并返回写入的字节数
func FilePutContents(filename string, data string, append ...int) bool {
	var flag int
	if len(append) > 0 && append[0] == FILE_APPEND {
		flag = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	} else {
		flag = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	}
	file, err := os.OpenFile(filename, flag, 0644)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return false
	}

	return true
}
