package php

import (
	"fmt"
	"reflect"
)

func VarDump(args ...interface{}) {
	for _, arg := range args {
		val := reflect.ValueOf(arg)
		typ := reflect.TypeOf(arg)
		fmt.Printf("(%s) ", typ.String())

		switch val.Kind() {
		case reflect.Invalid:
			fmt.Printf("nil\n")
		case reflect.Slice, reflect.Array:
			fmt.Printf("%d) { ", val.Len())
			for i := 0; i < val.Len(); i++ {
				if i > 0 {
					fmt.Printf(", ")
				}
				VarDump(val.Index(i).Interface())
			}
			fmt.Printf(" }\n")
		case reflect.Struct:
			fmt.Printf("{ ")
			for i := 0; i < val.NumField(); i++ {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s: ", typ.Field(i).Name)
				VarDump(val.Field(i).Interface())
			}
			fmt.Printf(" }\n")
		case reflect.Map:
			keys := val.MapKeys()
			fmt.Printf("{ ")
			for i := 0; i < len(keys); i++ {
				if i > 0 {
					fmt.Printf(", ")
				}
				VarDump(keys[i].Interface())
				fmt.Printf(": ")
				VarDump(val.MapIndex(keys[i]).Interface())
			}
			fmt.Printf(" }\n")
		default:
			fmt.Printf("%v\n", val)
		}
	}
}
