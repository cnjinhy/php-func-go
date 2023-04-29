package main

import (
	"fmt"
	"github.com/cnjinhy/php-func-go/php"
	"regexp"
	"strings"
)

func main() {
	var mdStr string
	mdStr += "# php-func-go\n\n"
	mdStr += "#### Description\nSome `PHP` built-in functions implemented using Golang\n\n使用golang实现的一些PHP函数，方便开发使用。\n\n"
	mdStr += "Note: Due to the language features of `Golang` and `PHP`, some functions are not 100% compatible. For example, the maps in Golang are unordered, and when using `http_ build_query`, the output order may be inconsistent. For example, arrays in Golang cannot be accessed outside the bounds, so there are restrictions on `empty`, and so on.\n\n"
	mdStr += "注意：由于`Golang`和`PHP`的语言特性，有些函数并不是100%兼容，比如Golang中的map是无序的，当使用`http_build_query`的时候有可能输出顺序不一致，比如Golang中的数组不能越界访问，所以empty也有所限制，等等。\n\n"
	mdStr += "#### Demo\nthe `README.md` is generated using the `php-func-go` function by [doc.go](https://github.com/cnjinhy/php-func-go/blob/master/doc.go)\n\n"
	mdStr += "`README.md`文件是通过`doc.go`生成的，`doc.go`使用了php-func-go的函数写的。\n\n"
	mdStr += "#### Unit Test\nYou can run `go test` in the dir `tests`\n\n"
	mdStr += "#### Function list\n"
	mdStr += "| php function | golang function | input argvs | return type |\n"
	mdStr += "|-------------|--------------|----------------------|--------|\n"
	fileList := php.Glob(php.MAGIC__DIR__() + "/php/*")
	for _, filePath := range fileList {
		phpFuncName := php.Basename(filePath, ".go")
		fileContent := php.FileGetContents(filePath)
		fileContentLines := php.Explode("<br />", php.Nl2Br(fileContent))
		goFuncName := php.Ucfirst(snakeToCamel(phpFuncName))

		for _, fileLine := range fileContentLines {
			fileLine = php.Trim(fileLine)
			goFuncDefine := "func " + goFuncName + "("
			//匹配函数定义行
			if php.StrToLower(php.SubStr(fileLine, 0, php.Strlen(goFuncDefine))) == php.StrToLower(goFuncDefine) {
				returnType := ""
				//由于PHP的函数大多只有一个返回值,我们处理解析一下
				funcLines := php.Explode(")", fileLine)
				argvLines := php.Explode("(", funcLines[0])
				search := [...]string{"{", "}"}
				strInterface := php.StrReplace(search, "", php.Trim(funcLines[1]))
				returnType = fmt.Sprintf("%s", strInterface)
				mdStr += "| " + phpFuncName + " | " + goFuncName + " | " + argvLines[1] + " | " + returnType + " |\n"
			}
		}
	}

	php.FilePutContents(php.MAGIC__DIR__()+"/README.md", mdStr)
	return
}

func snakeToCamel(str string) string {
	re := regexp.MustCompile(`_([a-z])`)
	return php.PregReplaceCallback(re.String(), str, func(matches []string) string {
		return strings.ToUpper(matches[1])
	})
}
