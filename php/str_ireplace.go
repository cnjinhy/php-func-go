package php

import (
	"fmt"
	"reflect"
	"strings"
)

func StriReplace(search interface{}, replace string, subject interface{}) string {
	sValue := reflect.ValueOf(search)
	rValue := reflect.ValueOf(replace)
	subValue := reflect.ValueOf(subject)

	if sValue.Kind() == reflect.Slice || sValue.Kind() == reflect.Array || sValue.Kind() == reflect.Map {
		sLen := sValue.Len()

		for i := 0; i < sLen; i++ {
			searchValue := strings.ToLower(fmt.Sprintf("%v", sValue.Index(i)))
			replaceValue := strings.ToLower(rValue.String())
			subjectValue := strings.ToLower(fmt.Sprintf("%v", subValue))

			// Replace all occurrences of the lowercase search string in the lowercase subject string.
			subjectValue = strings.ReplaceAll(subjectValue, searchValue, replaceValue)

			// Replace all occurrences of the original search string in the original subject string.
			subjectValue = strings.ReplaceAll(fmt.Sprintf("%v", subValue), fmt.Sprintf("%v", sValue.Index(i)), subjectValue)

			subValue = reflect.ValueOf(subjectValue)
		}
	} else {
		searchValue := strings.ToLower(sValue.String())
		replaceValue := strings.ToLower(rValue.String())
		subjectValue := strings.ToLower(fmt.Sprintf("%v", subValue))

		// Replace all occurrences of the lowercase search string in the lowercase subject string.
		subjectValue = strings.ReplaceAll(subjectValue, searchValue, replaceValue)

		// Replace all occurrences of the original search string in the original subject string.
		subjectValue = strings.ReplaceAll(fmt.Sprintf("%v", subValue), sValue.String(), subjectValue)

		subValue = reflect.ValueOf(subjectValue)
	}

	return subValue.String()
}
