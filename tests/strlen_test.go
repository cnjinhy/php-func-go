package tests

import (
	"phpfunc/php"
	"testing"
)

func TestStrlen(t *testing.T) {
	result := php.Strlen("cnjinhy")
	expected := 7
	if result != expected {
		t.Errorf("TestStrlen(cnjinhy) returned %d, expected %d", result, expected)
	}

	result = php.Strlen("Hello World")
	expected = 11
	if result != expected {
		t.Errorf("TestStrlen(Hello World) returned %d, expected %d", result, expected)
	}

	result = php.Strlen("Hello 中国")
	expected = 12
	if result != expected {
		t.Errorf("TestStrlen(Hello 中国) returned %d, expected %d", result, expected)
	}

	result = php.Strlen("Hello 中国 Hi")
	expected = 15
	if result != expected {
		t.Errorf("TestStrlen(Hello 中国) returned %d, expected %d", result, expected)
	}
}
