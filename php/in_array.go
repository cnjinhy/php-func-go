package php

import "reflect"

func InArray(needle interface{}, haystack interface{}, strict ...bool) bool {
	// 获取 haystack 的反射值和类型
	h := reflect.ValueOf(haystack)
	kind := h.Kind()

	// 判断 haystack 是否为切片或数组
	if kind != reflect.Slice && kind != reflect.Array {
		if kind == reflect.Map {
			// 遍历 map 查找 key 是否存在
			for _, key := range h.MapKeys() {
				if reflect.DeepEqual(key.Interface(), needle) {
					return true
				}
			}
		}
		return false
	}

	// 判断是否需要进行严格类型检查
	isStrict := false
	if len(strict) > 0 && strict[0] {
		isStrict = true
	}

	// 遍历切片或数组查找元素是否存在
	for i := 0; i < h.Len(); i++ {
		item := h.Index(i).Interface()
		if isStrict {
			if reflect.DeepEqual(item, needle) {
				return true
			}
		} else {
			if reflect.TypeOf(item) == reflect.TypeOf(needle) && item == needle {
				return true
			}
		}
	}
	return false
}
