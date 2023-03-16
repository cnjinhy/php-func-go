package tests

import (
	"phpfunc/php"
	"testing"
)

func TestMbStrlen(t *testing.T) {
	result := php.MbStrlen("cnjinhy")
	expected := 7
	if result != expected {
		t.Errorf("TestMbStrlen(cnjinhy) returned %d, expected %d", result, expected)
	}

	result = php.MbStrlen("Hello World")
	expected = 11
	if result != expected {
		t.Errorf("TestMbStrlen(Hello World) returned %d, expected %d", result, expected)
	}

	result = php.MbStrlen("Hello 中国")
	expected = 8
	if result != expected {
		t.Errorf("TestMbStrlen(Hello 中国) returned %d, expected %d", result, expected)
	}

	result = php.MbStrlen("Hello 中国 Hi")
	expected = 11
	if result != expected {
		t.Errorf("TestMbStrlen(Hello 中国 Hi) returned %d, expected %d", result, expected)
	}
}
