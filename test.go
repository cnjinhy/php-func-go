package main

import "php/php"

func main() {
	php.Echo(php.Dirname(php.MAGIC__FILE__()))
	php.Echo(php.MAGIC__DIR__())

	php.Echo(php.Md5File(php.MAGIC__FILE__()))

	//php.Echo(php.Filectime(php.MAGIC__FILE__()))

	php.VarDump(php.Glob(php.MAGIC__DIR__() + "/php/*.go"))
}
