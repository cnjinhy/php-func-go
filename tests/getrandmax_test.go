package tests

import (
	"phpfunc/php"
	"testing"
)

func TestGetRandMax(t *testing.T) {
	max := php.Getrandmax()

	if max != 2147483647 {
		t.Errorf("getRandMax() returned %d, expected 2147483647", max)
	}
}
