package php

import (
	"fmt"
	"reflect"
	"strings"
)

func StrReplace(search interface{}, replace string, subject interface{}) string {
	sValue := reflect.ValueOf(search)
	rValue := reflect.ValueOf(replace)
	subValue := reflect.ValueOf(subject)

	if sValue.Kind() == reflect.Slice || sValue.Kind() == reflect.Array || sValue.Kind() == reflect.Map {
		sLen := sValue.Len()

		for i := 0; i < sLen; i++ {
			searchValue := fmt.Sprintf("%v", sValue.Index(i))
			replaceValue := rValue.String()
			subjectValue := fmt.Sprintf("%v", subValue)

			subjectValue = strings.Replace(subjectValue, searchValue, replaceValue, -1)

			subValue = reflect.ValueOf(subjectValue)
		}
	} else {
		searchValue := sValue.String()
		replaceValue := rValue.String()
		subjectValue := fmt.Sprintf("%v", subValue)

		subjectValue = strings.Replace(subjectValue, searchValue, replaceValue, -1)

		subValue = reflect.ValueOf(subjectValue)
	}

	return subValue.String()
}
