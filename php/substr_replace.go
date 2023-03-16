package php

func SubstrReplace(str, replace string, start int, length ...int) string {
	strLen := len(str)
	if start < 0 {
		start = strLen + start
		if start < 0 {
			start = 0
		}
	}
	if start > strLen {
		start = strLen
	}
	var end int
	if len(length) > 0 {
		end = start + length[0]

		if length[0] < 0 {
			end = strLen + length[0]
			if end < start {
				end = start
			}
		}

		if end > strLen {
			end = strLen
		}
	} else {
		end = strLen
	}
	str = str[:start] + replace + str[end:]
	return str
}
