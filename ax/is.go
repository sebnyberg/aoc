package ax

func Is(s string, f func(r rune) bool) bool {
	for _, ch := range s {
		if !f(ch) {
			return false
		}
	}
	return true
}
