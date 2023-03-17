package php

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Sha1File(filename string, binary ...bool) string {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	// 创建SHA1哈希对象
	hash := sha1.New()

	// 将文件内容写入哈希对象
	_, err = io.Copy(hash, file)
	if err != nil {
		return ""
	}

	// 获取散列值
	hashValue := hash.Sum(nil)

	// 如果指定了binary为true，则返回原始二进制格式摘要
	if len(binary) > 0 && binary[0] {
		return string(hashValue)
	}

	// 将散列值转换为16进制字符串
	return hex.EncodeToString(hashValue)
}
