package php

func Addslashes(s string) string {
	var escaped string
	for _, char := range s {
		switch char {
		case '\'', '"', '\\':
			escaped += "\\"
			escaped += string(char)
		default:
			escaped += string(char)
		}
	}
	return escaped
}
