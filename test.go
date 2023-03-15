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

	d := php.Explode(",", "1,2,3,4,5")
	php.Echo(php.Implode(":", d))
	//php.Echo(php.FileGetContents("https://www.baidu.com"))
	php.Echo(php.StrToUpper("jjjffff"))

	php.Echo(php.MbStrlen("您好1"))
	php.Echo(php.Strlen("您好1"))
	php.Echo(php.CheckDate(12, 100, 2022))
	php.Echo(php.Trim("\nffff\n"))
}
