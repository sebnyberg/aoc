package ax

// All checks that all values in the slice are equal to the provided value.
func As[U, T ~uint8](ts []U) []T {
	res := make([]T, len(ts))
	for i, t := range ts {
		res[i] = T(t)
	}
	return res
}
