package php

import "reflect"

func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	var (
		int0   int   = 0
		int80  int8  = 0
		int160 int16 = 0
		int320 int32 = 0
		int640 int64 = 0

		uint0   uint   = 0
		uint80  uint8  = 0
		uint160 uint16 = 0
		uint320 uint32 = 0
		uint640 uint64 = 0

		float320 float32 = 0
		float640 float32 = 0
	)

	switch reflect.TypeOf(val).Kind() {
	case reflect.Slice, reflect.Array:
		return reflect.ValueOf(val).Len() == 0
	case reflect.Map:
		return reflect.ValueOf(val).Len() == 0
	case reflect.String:
		return len(val.(string)) == 0
	case reflect.Ptr:
		return Empty(reflect.ValueOf(val).Elem().Interface())
	case reflect.Struct:
		return reflect.DeepEqual(val, reflect.Zero(reflect.TypeOf(val)).Interface())
	case reflect.Bool:
		return !val.(bool)
	case reflect.Int:
		return val == int0
	case reflect.Int8:
		return val == int80
	case reflect.Int16:
		return val == int160
	case reflect.Int32:
		return val == int320
	case reflect.Int64:
		return val == int640
	case reflect.Uint:
		return val == uint0
	case reflect.Uint8:
		return val == uint80
	case reflect.Uint16:
		return val == uint160
	case reflect.Uint32:
		return val == uint320
	case reflect.Uint64:
		return val == uint640
	case reflect.Float32:
		return val == float320
	case reflect.Float64:
		return val == float640
	default:
		return reflect.ValueOf(val).IsZero()
	}
}
