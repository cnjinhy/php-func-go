package php

import (
	"reflect"
	"time"
)

func Gettype(any interface{}) string {
	if any == nil {
		return "nil"
	}
	switch value := any.(type) {
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case uint:
		return "uint"
	case uint8:
		return "uint8"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case string:
		return "string"
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		if value == nil {
			return ""
		}
		// Reflect checks.
		var (
			rv   = reflect.ValueOf(value)
			kind = rv.Kind()
		)
		switch kind {
		case reflect.Chan:
			return "chan"
		case
			reflect.Map:
			return "map"
		case
			reflect.Slice:
			return "slice"
		case
			reflect.Func:
			return "func"
		case
			reflect.Ptr:
			return "ptr"
		case
			reflect.Interface:
			return "interface"
		case
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
			return "unsafepointer"
		case reflect.String:
			return "string"
		}
	}
	return ""
}
