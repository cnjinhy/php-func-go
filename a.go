package main

import (
	"fmt"
)

type (
	// 定义一个PHP风格的map
	PHPMap map[interface{}]interface{}
)

func main() {
	// 定义一个PHP风格的map
	var m PHPMap = PHPMap{}

	// 设置键值对
	m["foo"] = "bar"
	m[123] = true
	m[3.14] = []string{"a", "b", "c"}

	// 访问键值对
	fmt.Println(m["foo"])   // 输出: bar
	fmt.Println(m[123])     // 输出: true
	fmt.Println(m[3.14][0]) // 输出: [a b c]
}
