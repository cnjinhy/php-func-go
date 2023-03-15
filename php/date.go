package php

import (
	"fmt"
	"time"
)

func Date(format string, timestamp ...interface{}) string {
	var timest int64
	var err error

	// 根据传入参数的数量，决定如何解析时间戳
	switch len(timestamp) {
	case 0:
		// 没有传入时间戳，则使用当前时间
		timest = time.Now().Unix()
	default:
		// 传入一个参数，则解析为时间戳
		timest, err = parseTimestamp(timestamp[0])
		if err != nil {
			return ""
		}
	}

	t := time.Unix(timest, 0)
	year, month, day := t.Date()
	hour, min, sec := t.Hour(), t.Minute(), t.Second()

	// 处理格式中的转义字符
	var result string
	//处理中文字符
	// Converting to []rune to support unicode.
	rune := []rune(format)
	for i := 0; i < len(rune); i++ {
		if rune[i] == '\\' {
			if i+1 < len(rune) {
				result += string(rune[i+1])
				i++
				continue
			}
		}

		switch rune[i] {
		case 'd':
			result += fmt.Sprintf("%02d", day)
		case 'D':
			result += t.Weekday().String()[:3]
		case 'j':
			result += fmt.Sprintf("%d", day)
		case 'l':
			result += t.Weekday().String()
		case 'S':
			switch day {
			case 1, 21, 31:
				result += "st"
			case 2, 22:
				result += "nd"
			case 3, 23:
				result += "rd"
			default:
				result += "th"
			}
		case 'w':
			result += fmt.Sprintf("%d", t.Weekday())
		case 'z':
			result += fmt.Sprintf("%d", t.YearDay()-1)
		case 'W':
			_, week := t.ISOWeek()
			result += fmt.Sprintf("%02d", week)
		case 'F':
			result += month.String()
		case 'm':
			result += fmt.Sprintf("%02d", int(month))
		case 'M':
			result += month.String()[:3]
		case 'n':
			result += fmt.Sprintf("%d", int(month))
		case 't':
			days := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
			result += fmt.Sprintf("%d", days)
		case 'L':
			if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
				result += "1"
			} else {
				result += "0"
			}
		case 'o':
			_, week := t.ISOWeek()
			if week == 1 && month == 12 {
				result += fmt.Sprintf("%d", year+1)
			} else if week >= 52 && month == 1 {
				result += fmt.Sprintf("%d", year-1)
			} else {
				result += fmt.Sprintf("%d", year)
			}
		case 'Y':
			result += fmt.Sprintf("%04d", year)
		case 'y':
			result += fmt.Sprintf("%02d", year%100)
		case 'a':
			if hour < 12 {
				result += "am"
			} else {
				result += "pm"
			}
		case 'A':
			if hour < 12 {
				result += "AM"
			} else {
				result += "PM"
			}
		case 'g':
			result += fmt.Sprintf("%d", hour%12)
		case 'G':
			result += fmt.Sprintf("%d", hour)
		case 'h':
			result += fmt.Sprintf("%02d", hour%12)
		case 'H':
			result += fmt.Sprintf("%02d", hour)
		case 'i':
			result += fmt.Sprintf("%02d", min)
		case 's':
			result += fmt.Sprintf("%02d", sec)
		case 'O':
			tz := t.Format("-0700")
			result += tz[:3] + ":" + tz[3:]
		case 'P':
			tz := t.Format("-07:00")
			result += tz[:3] + ":" + tz
		case 'T':
			result += t.Format("MST")
		case 'c':
			result += t.Format("2006-01-02T15:04:05-07:00")
		case 'r':
			result += t.Format("Mon, 02 Jan 2006 15:04:05 MST")
		case 'U':
			result += fmt.Sprintf("%d", timestamp)
		default:
			result += string(rune[i])
		}
	}

	return result
}

// 解析传入的时间戳
func parseTimestamp(arg interface{}) (int64, error) {
	switch v := arg.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		t, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
		if err != nil {
			return 0, err
		}
		return t.Unix(), nil
	default:
		return 0, fmt.Errorf("unsupported argument type %T", arg)
	}
}

func CheckDate(month, day, year int) bool {
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}
	if year < 1 || year > 9999 {
		return false
	}
	if month == 2 {
		if isLeapYear(year) {
			return day <= 29
		}
		return day <= 28
	}
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return day <= 30
	}
	return true
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}
