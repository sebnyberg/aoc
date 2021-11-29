package ax

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Min16(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
