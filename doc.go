package main

import (
	"fmt"
	"github.com/cnjinhy/php-func-go/php"
	"regexp"
	"strings"
)

func main() {
	var mdStr string
	mdStr += "#### php-func-go\nSome PHP built-in functions implemented using Golang\n"
	mdStr += "#### description\nthe README.md is generated using the php-func-go function through doc.go\n"
	mdStr += "#### function list\n"
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
			if php.SubStr(fileLine, 0, php.Strlen(goFuncDefine)) == goFuncDefine {
				returnType := ""
				//由于PHP的函数大多只有一个返回值,我们处理解析一下
				funcLines := php.Explode(")", fileLine)
				argvLines := php.Explode("(", funcLines[0])
				strInterface := php.StrReplace("", "", php.Trim(funcLines[1]))
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