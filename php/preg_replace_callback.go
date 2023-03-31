package php

import "regexp"

func PregReplaceCallback(pattern string, subject string, callback func([]string) string) string {
	re := regexp.MustCompile(pattern)

	var result []byte
	lastIndex := 0

	for _, match := range re.FindAllStringSubmatchIndex(subject, -1) {
		result = append(result, []byte(subject[lastIndex:match[0]])...)
		result = append(result, []byte(callback(re.FindStringSubmatch(subject[match[0]:match[1]])))...)
		lastIndex = match[1]
	}

	result = append(result, []byte(subject[lastIndex:])...)
	return string(result)
}
