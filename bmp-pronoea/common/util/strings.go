package util

import (
	"strings"
	"unicode"
)

func IsNotEmpty(str ...string) bool {
	for _, s := range str {
		if s == "" {
			return false
		}
	}
	return true
}

func IsEmpty(str string) bool {
	return str == ""
}

func StringsContains(str, subStr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(subStr))
}

func UnderLineToHumpCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return string(unicode.ToLower(rune(name[0]))) + name[1:]
}

func HeadToLower(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}