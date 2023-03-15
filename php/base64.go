package php

import "encoding/base64"

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(encoded string) string {
	data, _ := base64.StdEncoding.DecodeString(encoded)
	return string(data)
}
