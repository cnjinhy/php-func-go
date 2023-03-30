package main

import "github.com/cnjinhy/php-func-go/php"

func main() {
	php.Echo(php.Date("Y-m-d H:i:s"))

	php.Echo(php.StrToTime("+1 day"))
}
