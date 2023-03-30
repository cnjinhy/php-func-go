package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestUcfirst(t *testing.T) {
	result := php.Ucfirst("hello World")
	expected := "Hello World"
	if result != expected {
		t.Errorf("Ucfirst(\"hello World\") returned %s, expected %s", result, expected)
	}

	result = php.Ucfirst("hello 中国")
	expected = "Hello 中国"
	if result != expected {
		t.Errorf("Ucfirst(\"hello World\") returned %s, expected %s", result, expected)
	}

}
