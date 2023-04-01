package main

import "github.com/cnjinhy/php-func-go/php"

func main() {
	var a int
	var b int64
	a = 0
	b = 0
	php.VarDump(a == 0)
	php.VarDump(b == 0)

	php.Echo(php.Date("Y-m-d H:i:s"))
	data := make(map[string]interface{})
	data["a"] = "ValueA"
	data["b"] = 123
	data["c"] = "Hello"
	data["f"] = "中国"
	php.Echo(php.HttpBuildQuery(data))
	ff := make(map[string]interface{})
	ss := make(map[string]interface{})
	ss["d"] = "ssValueD"
	ff["ddd"] = "ddd"
	ss["ds"] = 123
	ss["ff"] = ff
	data["s"] = ss
	php.PrintR(data)

	php.VarDump(php.Gettype(data))

	s := php.ParseUrl("http://username:password@hostname:9090/path?arg=value#anchor")
	php.VarDump(s)

	php.VarDump(php.Microtime())
	php.VarDump(php.Microtime(true))

}
