package ax

import "golang.org/x/exp/constraints"

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

type TNegative interface {
	constraints.Signed | constraints.Float
}

func Abs[T TNegative](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
