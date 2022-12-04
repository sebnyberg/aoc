package main

import (
	"container/ring"
	"fmt"
)

func solve1() string {
	input := 3017957
	return fmt.Sprint(g(input, 2) + 1)
}

func g(n, k int) int {
	if n == 1 {
		return 0
	}
	return (g(n-1, k) + k) % n
}

func solve2() string {
	// With pen and paper I realised that there is a pattern of pairs of elves
	// getting removed from the ring until it only has one elf left.
	n := 3017957
	r := ring.New(n - 1)
	for x := 1; x <= n; x++ {
		if x == n/2+1 {
			continue
		}
		r.Value = x
		r = r.Next()
	}
	for r.Value != n/2+2 {
		r = r.Next()
	}
	n--
	for n > 1 {
		if n > 2 {
			r.Unlink(2)
			n -= 2
		} else {
			r.Unlink(1)
			n--
		}
		r = r.Next()
	}
	return fmt.Sprint(r.Value)
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve1())
	fmt.Printf("Result2:\n%v\n\n", solve2())
}
