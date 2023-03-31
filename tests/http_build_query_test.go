package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestHttpBuildQuery(t *testing.T) {
	data := make(map[string]interface{})
	data["a"] = "ValueA"
	data["b"] = 123
	data["c"] = "Hello"
	data["f"] = "中国"

	result := php.HttpBuildQuery(data)
	expected := "a=ValueA&b=123&c=Hello&f=%E4%B8%AD%E5%9B%BD"
	if result != expected {
		t.Errorf("TestMbStrlen(cnjinhy) returned %s, expected %s", result, expected)
	}

}
