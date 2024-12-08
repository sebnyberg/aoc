package ax

import "golang.org/x/exp/constraints"

func ModInverse[T constraints.Integer](a, mod T) T {
	return ModPow(a, mod-2, mod)
}

func ModPow[T constraints.Integer](a, b, mod T) T {
	if b == 0 {
		return 1
	}
	p := ModPow(a, b/2, mod) % mod
	p = p * p % mod
	if b%2 == 0 {
		return p
	}
	return (a * p) % mod
}

func Pow[T constraints.Integer](a, b T) T {
	if b == 0 {
		return 1
	}
	p := Pow(a, b/2)
	p = p * p
	if b%2 == 0 {
		return p
	}
	return a * p
}
