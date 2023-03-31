package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestStrReplace(t *testing.T) {
	search := []string{"world", "Golang"}
	replace := "Hello"
	subject := "Hello, world! Welcome to the world of Golang!"
	expected := "Hello, Hello! Welcome to the Hello of Hello!"
	result := php.StrReplace(search, replace, subject)
	if result != expected {
		t.Errorf("TestStrReplace returned %s, expected %s", result, expected)
	}

	searchStr := "abc"
	replace = "def"
	subject = "Hello, abc!"
	expected = "Hello, def!"
	result = php.StrReplace(searchStr, replace, subject)
	if result != expected {
		t.Errorf("TestStrReplace returned %s, expected %s", result, expected)
	}

	searchStr = "中国"
	replace = "美国"
	subject = "Hello, 中国!"
	expected = "Hello, 美国!"
	result = php.StrReplace(searchStr, replace, subject)
	if result != expected {
		t.Errorf("TestStrReplace returned %s, expected %s", result, expected)
	}

}
