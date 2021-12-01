package ax

// Min returns the minimum value of the two input values, or the first value if
// the values are equal.
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Min16 returns the maximum value of the two input values, or the first value
// if the values are equal.
func Min16(a, b int16) int16 {
	if a <= b {
		return a
	}
	return b
}

// Max returns the maximum value of the two input values, or the first value if
// the values are equal.
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Max16 returns the maximum value of the two input values, or the first value
// if the values are equal.
func Max16(a, b int16) int16 {
	if a >= b {
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
