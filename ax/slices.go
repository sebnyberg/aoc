package ax

import (
	"golang.org/x/exp/constraints"
)

// All checks whether all elements pass a given predicate.
func All[S ~[]T, T comparable](ts S, f func(T) bool) bool {
	for _, tt := range ts {
		if !f(tt) {
			return false
		}
	}
	return true
}

// Sum calculates the arithmetic sum of a slice of numbers.
func Sum[S ~[]T, T constraints.Ordered](a S) T {
	var sum T
	for _, item := range a {
		sum += item
	}
	return sum
}

// Head returns the top n items in the slice.
func Head[T any](a []T, n int) []T {
	if n > len(a) {
		n = len(a)
	}
	return a[:n]
}

// Tail returns the last n items in the slice.
func Tail[T any](a []T, n int) []T {
	if n > len(a) {
		n = len(a)
	}
	return a[len(a)-n:]
}

// SliceToMapSet converts a slice to a map
func SliceToMapSet[S ~[]T, T comparable](a S) map[T]struct{} {
	m := make(map[T]struct{}, len(a))
	for _, x := range a {
		m[x] = struct{}{}
	}
	return m
}
