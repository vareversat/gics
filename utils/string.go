package utils

import "strings"

const separator = ","

// StringToStringArray Use default const.separator
func StringToStringArray(value string) []string {
	return strings.Split(value, separator)
}

// StringToStringArrayWithSeparator override the default separator
func StringToStringArrayWithSeparator(value string, separator string) []string {
	return strings.Split(value, separator)
}
