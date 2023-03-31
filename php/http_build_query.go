package php

import (
	"fmt"
	"reflect"
	"strings"
)

func HttpBuildQuery(data interface{}) string {
	values := reflect.ValueOf(data)
	if values.Kind() != reflect.Map && values.Kind() != reflect.Struct {
		panic("http_build_query: data must be a map or struct")
	}

	var buf strings.Builder
	switch values.Kind() {
	case reflect.Map:
		for _, k := range values.MapKeys() {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(urlencode(fmt.Sprintf("%v", k)))
			buf.WriteByte('=')
			buf.WriteString(urlencode(fmt.Sprintf("%v", values.MapIndex(k))))
		}
	case reflect.Struct:
		for i := 0; i < values.NumField(); i++ {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(urlencode(values.Type().Field(i).Name))
			buf.WriteByte('=')
			buf.WriteString(urlencode(fmt.Sprintf("%v", values.Field(i))))
		}
	}

	return buf.String()
}

func urlencode(str string) string {
	var buf strings.Builder
	for _, b := range []byte(str) {
		if shouldEscape(b) {
			buf.WriteString(fmt.Sprintf("%%%02X", b))
		} else {
			buf.WriteByte(b)
		}
	}
	return buf.String()
}

func shouldEscape(c byte) bool {
	if c > 126 || c < 33 {
		return true
	}
	return false
}
