package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestStrToTime(t *testing.T) {
	result := php.StrToTime("2023-03-30 12:12:12")
	var expected int64
	expected = 1680149532
	if result != expected {
		t.Errorf("StrToTime(2023-03-30 12:12:12) returned %d, expected %d", result, expected)
	}
}
