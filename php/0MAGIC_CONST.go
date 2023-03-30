package php

import (
	"path/filepath"
	"reflect"
	"runtime"
)

func MAGIC__FILE__() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filename
}

func MAGIC__LINE__() int {
	_, _, line, ok := runtime.Caller(1)
	if !ok {
		return 0
	}
	return line
}

func MAGIC__FUNCTION__() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return runtime.FuncForPC(pc).Name()
}

func MAGIC__DIR__() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Dir(file)
}

func MAGIC__NAMESPACE__() string {
	pkgPath := reflect.TypeOf(0).PkgPath()
	return pkgPath
}
