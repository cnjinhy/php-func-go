package x

import (
	"github.com/elliotchance/orderedmap"
	"strings"
)

func Get(om *orderedmap.OrderedMap, path string, defaultValue ...interface{}) interface{} {
	// 将路径按点分隔为键的列表
	// a.b.c.d
	keys := strings.Split(path, ".")
	// 从第一层开始遍历
	var value interface{}
	for index, key := range keys {
		if index == 0 {
			value, _ = om.Get(key)
		} else {
			//判断类型
			if _, ok := value.(*orderedmap.OrderedMap); !ok {
				panic("Invalid type must be *orderedmap.OrderedMap")
			}
			value, _ = value.(*orderedmap.OrderedMap).Get(key)
		}
	}
	// 返回最后一层的value，如果为nil则返回默认值
	if value == nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return value
}
