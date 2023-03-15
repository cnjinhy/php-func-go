package main

import "phpfunc/php"

func main() {
	php.Echo(php.Date("Y年m月d日 H:i:s"))
	php.Echo("呵呵")
	php.Echo(php.Time())
	php.Echo(php.Microtime(true))
	php.Echo(php.Sha1("hehe"))
	php.Echo(php.MbSubStr("12345", 0, 4))

	php.Echo(php.Base64Encode("haha"))

	php.Echo(php.Base64Decode(php.Base64Encode("haha")))

}
