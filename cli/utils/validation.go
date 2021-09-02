package utils

import (
	"strings"
	"unicode"
)

// https://stackoverflow.com/a/31964846
func IsValidIdentifierString(id string) bool {
	isValidRune := func(r rune) bool {
		return !unicode.IsDigit(r) && r != '-' && r != '_' && (r < 'A' || r > 'z')
	}

	return (strings.IndexFunc(id, isValidRune) == -1)
}
