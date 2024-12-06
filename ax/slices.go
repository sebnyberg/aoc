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

// Reverse reverses a slice, returning a copy
func Reverse[T any](a []T) []T {
	cpy := make([]T, len(a))
	copy(cpy, a)
	ReverseInplace(cpy)
	return cpy
}

// ReverseInPlace reverses a slice in-place
func ReverseInplace[T any](a []T) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}

// ReverseRows reverses the rows of a 2D slice, returning a copy
func ReverseRows[T any](a [][]T) [][]T {
	cpy := make([][]T, len(a))
	for i := range cpy {
		cpy[i] = make([]T, len(a[i]))
		copy(cpy[i], a[i])
	}
	a = cpy
	ReverseRowsInplace(a)
	return cpy
}

// ReverseRows reverses the rows of a 2D slice
func ReverseRowsInplace[T any](a [][]T) {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
}
