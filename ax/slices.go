package ax

import "golang.org/x/exp/constraints"

func All[T comparable](ts []T, t T) bool {
	for _, tt := range ts {
		if tt != t {
			return false
		}
	}
	return true
}

func Sum[T constraints.Ordered](a []T) T {
	var sum T
	for _, item := range a {
		sum += item
	}
	return sum
}

func SumIf[T constraints.Ordered](a []T, f func(T) bool) T {
	var sum T
	for _, item := range a {
		if f(item) {
			sum += item
		}
	}
	return sum
}

func SumMatrix[T constraints.Ordered](a [][]T) T {
	var sum T
	for i := range a {
		for _, item := range a[i] {
			sum += item
		}
	}
	return sum
}

func SumMatrixIf[T constraints.Ordered](a [][]T, f func(int, int, T) bool) T {
	var sum T
	for i := range a {
		for j, item := range a[i] {
			if f(i, j, item) {
				sum += item
			}
		}
	}
	return sum
}

func CountIf[T any](a []T, f func(int, T) bool) int {
	var count int
	for i, el := range a {
		if f(i, el) {
			count++
		}
	}
	return count
}

func CountIf2D[T any](a [][]T, f func(int, int, T) bool) int {
	var count int
	for i := range a {
		for j, el := range a[i] {
			if f(i, j, el) {
				count++
			}
		}
	}
	return count
}
