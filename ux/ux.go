package ux

import "constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Abs[T constraints.Signed](a T) T {
	if a < 0 {
		return -a
	}
	return a
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

func CountMatrixIf[T any](a [][]T, f func(int, int, T) bool) int {
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

type Stack[T any] []T

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		var nilT T
		return nilT
	}
	it := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return it
}

func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		var nilT T
		return nilT
	}
	return (*s)[len(*s)-1]
}

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

func MapSet[T comparable](a []T) map[T]struct{} {
	m := make(map[T]struct{}, len(a))
	for _, el := range a {
		m[el] = struct{}{}
	}
	return m
}
