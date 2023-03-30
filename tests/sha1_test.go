package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestSha1(t *testing.T) {
	result := php.Sha1("cnjinhy")
	expected := "f6b3c70ba74b94e445b0615d2f4d812386d1b647"
	if result != expected {
		t.Errorf("Sha1(cnjinhy) returned %s, expected %s", result, expected)
	}
}
