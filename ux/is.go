package ux

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
