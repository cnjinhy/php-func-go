package php

import (
	"io/ioutil"
	"net/http"
	"os"
)

func FileGetContents(filename string) string {
	var content []byte
	var err error

	if _, err = os.Stat(filename); err == nil {
		// 文件存在，读取内容
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			return ""
		}
	} else {
		// 文件不存在，尝试获取 URL 内容
		resp, err := http.Get(filename)
		if err != nil {
			return ""
		}
		defer resp.Body.Close()

		content, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
	}

	return string(content)
}
