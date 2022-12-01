package ax

func StringIsFunc(s string, f func(r rune) bool) bool {
	for _, ch := range s {
		if !f(ch) {
			return false
		}
	}
	return true
}

func StringIs(s string, ch rune) bool {
	for _, c := range s {
		if c != ch {
			return false
		}
	}
	return true
}
