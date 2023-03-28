package php

import "reflect"

func ArrayKeyExists(key interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Map:
		m := reflect.ValueOf(array)
		if m.MapIndex(reflect.ValueOf(key)).IsValid() {
			return true
		}
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if i == key {
				return true
			}
		}
	}
	return false
}
