package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestBaseName(t *testing.T) {
	result := php.Basename("/etc/sudoers.d")
	expected := "sudoers.d"
	if result != expected {
		t.Errorf("TestBaseName returned %s, expected %s", result, expected)
	}

	result = php.Basename("/etc/sudoers.d", ".d")
	expected = "sudoers"
	if result != expected {
		t.Errorf("TestBaseName returned %s, expected %s", result, expected)
	}

	result = php.Basename("/etc/password")
	expected = "password"
	if result != expected {
		t.Errorf("TestBaseName returned %s, expected %s", result, expected)
	}

	result = php.Basename(".")
	expected = "."
	if result != expected {
		t.Errorf("TestBaseName returned %s, expected %s", result, expected)
	}

}
