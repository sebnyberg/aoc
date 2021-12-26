package ax

// Is checks whether all runes return true when passed through the provided
// function. For function examples, see the unicode package, e.g.
// unicode.IsLetter.
func Is(s string, f func(r rune) bool) bool {
	for _, ch := range s {
		if !f(ch) {
			return false
		}
	}
	return true
}

func IsExactly(s string, ch rune) bool {
	for _, c := range s {
		if c != ch {
			return false
		}
	}
	return true
}

// All checks that all values in the slice are equal to the provided value.
func All[T comparable](ts []T, t T) bool {
	for _, tt := range ts {
		if tt != t {
			return false
		}
	}
	return true
}
