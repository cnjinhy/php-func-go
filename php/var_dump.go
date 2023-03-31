package php

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const indent = "    "

func VarDump(v interface{}) {
	fmt.Printf("%s\n", varDump(v))
}

func varDump(v interface{}) string {
	return dumpValue(reflect.ValueOf(v), 0, false)
}

func dumpValue(v reflect.Value, level int, lastChild bool) string {
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("string(%d) \"%s\"", len(v.String()), v.String())
	case reflect.Slice, reflect.Array:
		return dumpSlice(v, level, lastChild)
	case reflect.Map:
		return dumpMap(v, level, lastChild)
	case reflect.Struct:
		return dumpStruct(v, level, lastChild)
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return "NULL"
		}
		return dumpValue(v.Elem(), level, lastChild)
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
	return fmt.Sprintf("%v", v.Interface())
}

func dumpSlice(v reflect.Value, level int, lastChild bool) string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf("slice(%d) {\n", v.Len()))

	for i := 0; i < v.Len(); i++ {
		output.WriteString(strings.Repeat(indent, level+1))

		childValue := v.Index(i)
		output.WriteString("[" + strconv.Itoa(i) + "] => ")

		if childValue.Kind() == reflect.Ptr {
			if childValue.IsNil() {
				output.WriteString("NULL")
			} else {
				output.WriteString(dumpValue(childValue.Elem(), level+1, i == v.Len()-1))
			}
		} else {
			output.WriteString(dumpValue(childValue, level+1, i == v.Len()-1))
		}

		output.WriteString("\n")
	}

	output.WriteString(strings.Repeat(indent, level))
	output.WriteString("}")
	if !lastChild {
		output.WriteString("\n")
	}
	return output.String()
}

func dumpMap(v reflect.Value, level int, lastChild bool) string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf("map(%d) {\n", v.Len()))

	for i, key := range v.MapKeys() {
		output.WriteString(strings.Repeat(indent, level+1))

		output.WriteString("[\"")
		if key.Kind() == reflect.Ptr {
			if key.IsNil() {
				output.WriteString("NULL")
			} else {
				output.WriteString(key.Elem().String())
			}
		} else {
			output.WriteString(key.String())
		}
		output.WriteString("\"] => ")

		childValue := v.MapIndex(key)
		if childValue.Kind() == reflect.Ptr {
			if childValue.IsNil() {
				output.WriteString("NULL")
			} else {
				output.WriteString(dumpValue(childValue.Elem(), level+1, i == v.Len()-1))
			}
		} else {
			output.WriteString(dumpValue(childValue, level+1, i == v.Len()-1))
		}

		output.WriteString("\n")
	}

	output.WriteString(strings.Repeat(indent, level))
	output.WriteString("}")
	if !lastChild {
		output.WriteString("\n")
	}
	return output.String()
}

func dumpStruct(v reflect.Value, level int, lastChild bool) string {
	t := v.Type()

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	var output strings.Builder
	output.WriteString(fmt.Sprintf("struct %s {\n", t.Name()))

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Name

		output.WriteString(fmt.Sprintf("%s%s: ", strings.Repeat(indent, level+1), fieldName))

		childLast := i == v.NumField()-1
		child := dumpValue(field, level+1, childLast)

		output.WriteString(child)
		output.WriteString("\n")
	}

	output.WriteString(fmt.Sprintf("%s}", strings.Repeat(indent, level)))
	if !lastChild {
		output.WriteString(",")
	}
	return output.String()
}
