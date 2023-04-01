package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestGettype(t *testing.T) {
	var test1 int
	test1 = 1
	result := php.Gettype(test1)
	expected := "int"
	if result != expected {
		t.Errorf("Gettype(test1) returned %s, expected %s", result, expected)
	}

	var test2 int64
	test2 = 123
	result = php.Gettype(test2)
	expected = "int64"
	if result != expected {
		t.Errorf("Gettype(test2) returned %s, expected %s", result, expected)
	}

	test3 := "123"
	result = php.Gettype(test3)
	expected = "string"
	if result != expected {
		t.Errorf("Gettype(test3) returned %s, expected %s", result, expected)
	}

	test4 := make(map[string]interface{})
	test4["a"] = "ValueA"
	test4["b"] = 123
	test4["c"] = "Hello"
	result = php.Gettype(test4)
	expected = "map"
	if result != expected {
		t.Errorf("Gettype(test4) returned %s, expected %s", result, expected)
	}

	test5 := 123
	result = php.Gettype(test5)
	expected = "int"
	if result != expected {
		t.Errorf("Gettype(test5) returned %s, expected %s", result, expected)
	}
}
