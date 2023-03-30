package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestMd5File(t *testing.T) {
	result := php.Md5File(php.Dirname(php.MAGIC__FILE__()) + "/test.txt")
	expected := "018a9db227ad427a1020e995bd9ef241"
	if result != expected {
		t.Errorf("Md5File("+php.Dirname(php.MAGIC__FILE__())+"/test.txt"+") returned %s, expected %s", result, expected)
	}
}
