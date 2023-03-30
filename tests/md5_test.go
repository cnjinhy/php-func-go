package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestMd5(t *testing.T) {
	result := php.Md5("cnjinhy")
	expected := "018a9db227ad427a1020e995bd9ef241"
	if result != expected {
		t.Errorf("Md5(cnjinhy) returned %s, expected %s", result, expected)
	}
}
