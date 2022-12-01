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

// Map maps a slice to a new slice using a function f.
func Map[T any, U any](a []T, f func(T) U) []U {
	res := make([]U, len(a))
	for i := range a {
		res[i] = f(a[i])
	}
	return res
}

// ForEach calls the function for each element in the slice
func ForEach[S ~[]T, T any](a S, f func(T)) {
	for i := range a {
		f(a[i])
	}
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

// MapInplace does an in-place mapping of elements. This function should not be
// used - it exists merely as a signal that .Map is not actually an in-place
// operation.
func MapInplace[S ~[]T, T any](a S, f func(T) T) {
	for i := range a {
		a[i] = f(a[i])
	}
}
