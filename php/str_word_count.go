package php

import (
	"strings"
	"unicode"
)

func StrWordCount(s string, format ...int) interface{} {
	// 根据不同的格式返回不同的结果
	switch len(format) {
	case 0:
		return len(strings.Fields(s))
	case 1:
		switch format[0] {
		case 0:
			return strings.Fields(s)
		case 1:
			// 以数组形式返回每个单词和它在字符串中出现的位置
			var words []string
			var positions []int
			for _, v := range strings.FieldsFunc(s, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			}) {
				words = append(words, v)
				positions = append(positions, strings.Index(s, v))
			}
			return []interface{}{words, positions}
		case 2:
			// 返回所有单词的首字母大写形式
			words := strings.Fields(s)
			for i, v := range words {
				words[i] = strings.Title(strings.ToLower(v))
			}
			return strings.Join(words, " ")
		}
	}
	return nil
}
