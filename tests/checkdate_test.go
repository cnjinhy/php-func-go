package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestCheckdate(t *testing.T) {
	result := php.CheckDate(12, 31, 2000)
	expected := true
	if result != expected {
		t.Errorf("Checkdate(12, 31, 2000) returned %t, expected %t", result, expected)
	}

	result = php.CheckDate(2, 29, 2001)
	expected = false
	if result != expected {
		t.Errorf("Checkdate(2, 29, 2000) returned %t, expected %t", result, expected)
	}

	result = php.CheckDate(12, 31, -400)
	expected = false
	if result != expected {
		t.Errorf("Checkdate(12, 31, -400) returned %t, expected %t", result, expected)
	}
}
