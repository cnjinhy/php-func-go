package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestFilesize(t *testing.T) {
	result, _ := php.Filesize(php.Dirname(php.MAGIC__FILE__()) + "/test.txt")
	expected := int64(7)
	if result != expected {
		t.Errorf("Filesize("+php.Dirname(php.MAGIC__FILE__())+"/test.txt"+") returned %d, expected %d", result, expected)
	}
}
