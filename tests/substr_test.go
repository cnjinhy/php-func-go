package tests

import (
	"phpfunc/php"
	"testing"
)

func TestSubstr(t *testing.T) {
	result := php.SubStr("Hello China", 0, 5)
	expected := "Hello"
	if result != expected {
		t.Errorf("TestSubstr(Hello China,0,5) returned %s, expected %s", result, expected)
	}

	result = php.SubStr("Hello China", 1, 6)
	expected = "ello C"
	if result != expected {
		t.Errorf("TestSubstr(Hello China,1,6) returned %s, expected %s", result, expected)
	}

	result = php.SubStr("Hello China", 1)
	expected = "ello China"
	if result != expected {
		t.Errorf("TestSubstr(Hello China,1) returned %s, expected %s", result, expected)
	}

	result = php.SubStr("Hello China", 1, -2)
	expected = "ello Chi"
	if result != expected {
		t.Errorf("TestSubstr(Hello China,1,-2) returned %s, expected %s", result, expected)
	}

	result = php.SubStr("Hello China", -3)
	expected = "ina"
	if result != expected {
		t.Errorf("TestSubstr(Hello China,-3) returned %s, expected %s", result, expected)
	}

	result = php.SubStr("Hello China", -3, 1)
	expected = "i"
	if result != expected {
		t.Errorf("TestSubstr(Hello China,-3,1) returned %s, expected %s", result, expected)
	}
}
