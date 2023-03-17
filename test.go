package main

import "phpfunc/php"

func main() {
	php.Echo(php.Dirname(php.MAGIC__FILE__()))

}
