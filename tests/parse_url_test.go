package tests

import (
	"github.com/cnjinhy/php-func-go/php"
	"testing"
)

func TestParseUrl(t *testing.T) {
	s := php.ParseUrl("https://tom:123456@hostname:9090/api?arg=value&b=1#anchor")
	if s.Scheme != "https" {
		t.Errorf("ParseUrl Port returned %s, expected %s", s.Scheme, "https")
	}
	if s.Port != "9090" {
		t.Errorf("ParseUrl Port returned %s, expected %s", s.Port, "9090")
	}
	if s.User != "tom" {
		t.Errorf("ParseUrl User returned %s, expected %s", s.User, "tom")
	}
	if s.Password != "123456" {
		t.Errorf("ParseUrl Password returned %s, expected %s", s.Password, "123456")
	}
	if s.Fragment != "anchor" {
		t.Errorf("ParseUrl Port returned %s, expected %s", s.Fragment, "anchor")
	}
	if s.Path != "/api" {
		t.Errorf("ParseUrl Path returned %s, expected %s", s.Path, "/api")
	}
	if s.Query["arg"] != "value" {
		t.Errorf("ParseUrl Query:arg returned %s, expected %s", s.Query["arg"], "value")
	}

}
