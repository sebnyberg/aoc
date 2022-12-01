package ax

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	var i int
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

func Items[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, len(m))
	values := make([]V, len(m))
	var i int
	for k, v := range m {
		values[i] = v
		keys[i] = k
		i++
	}
	return keys, values
}
