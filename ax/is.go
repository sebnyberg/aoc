package ax

import (
	"unicode"
)

// IsInt returns whether the string is a valid integer.
func IsInt(s string) bool {
	for i := range s {
		if i == 0 {
			if s[i] == '-' {
				continue
			}
			if s[i] >= '1' && s[i] <= '9' {
				continue
			}
			return false
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// IsWord returns whether s is a word (only alpha letters).
func IsWord(s string) bool {
	return StringIsFunc(s, unicode.IsLetter)
}
