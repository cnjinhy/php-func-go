package php

import "encoding/json"

func JsonEncode(value interface{}) string {
	jsonStr, _ := json.Marshal(value)
	return string(jsonStr)
}
