package php

import (
	"fmt"
	"reflect"
)

func VarDump(val interface{}) {
	fmt.Printf("(%v) %v = ", reflect.TypeOf(val), reflect.ValueOf(val))
	switch reflect.ValueOf(val).Kind() {
	case reflect.Slice:
		fmt.Printf("[]")
		for i := 0; i < reflect.ValueOf(val).Len(); i++ {
			VarDump(reflect.ValueOf(val).Index(i).Interface())
			fmt.Printf(", ")
		}
	case reflect.Map:
		fmt.Printf("map[")
		for _, key := range reflect.ValueOf(val).MapKeys() {
			VarDump(key.Interface())
			fmt.Printf(": ")
			VarDump(reflect.ValueOf(val).MapIndex(key).Interface())
			fmt.Printf(", ")
		}
		fmt.Printf("]")
	default:
		fmt.Printf("%v", reflect.ValueOf(val))
	}
}
