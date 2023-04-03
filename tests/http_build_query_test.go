package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"github.com/elliotchance/orderedmap"
	"testing"
)

func TestHttpBuildQuery(t *testing.T) {
	data := orderedmap.NewOrderedMap()
	data.Set("a", "ValueA")
	data.Set("b", 123)
	data.Set("c", "Hello")
	data.Set("f", "中国")

	//php.VarDump(data)
	result := php.HttpBuildQuery(data)
	expected := "a=ValueA&b=123&c=Hello&f=%E4%B8%AD%E5%9B%BD"
	if result != expected {
		t.Errorf("HttpBuildQuery returned %s, expected %s", result, expected)
	}
}
