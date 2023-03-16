package php

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
