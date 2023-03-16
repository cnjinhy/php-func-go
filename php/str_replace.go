package php

import "strings"

func StrReplace(search, replace, subject interface{}) interface{} {
	switch s := subject.(type) {
	case string:
		var r string
		switch t := replace.(type) {
		case string:
			r = strings.Replace(s, search.(string), t, -1)
		case []string:
			for _, v := range t {
				s = strings.Replace(s, search.(string), v, 1)
			}
			r = s
		}
		return r
	case []string:
		var r []string
		for _, v := range s {
			r = append(r, StrReplace(search, replace, v).(string))
		}
		return r
	}
	return ""
}
