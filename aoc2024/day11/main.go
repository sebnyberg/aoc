package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	return solve(inf, 25)
}

func solve2(inf string) any {
	return solve(inf, 75)
}

func solve(inf string, iterations int) any {
	curr := make(map[int]int)
	next := make(map[int]int)

	line := ax.MustReadFileLines(inf)[0]
	for _, w := range strings.Fields(line) {
		curr[ax.Atoi(w)]++
	}

	digits := func(x int) int {
		if x == 0 {
			return 1
		}
		var res int
		for x > 0 {
			res++
			x /= 10
		}
		return res
	}

	split := func(x, digits int) (int, int) {
		var r int
		mul := 1
		for i := 0; i < digits; i++ {
			r += (x % 10) * mul
			x /= 10
			mul *= 10
		}
		return x, r
	}

	for i := 0; i < iterations; i++ {
		for k := range next {
			delete(next, k)
		}
		for x, count := range curr {
			if x == 0 {
				next[1] += count
				continue
			}
			if d := digits(x); d%2 == 0 {
				l, r := split(x, d/2)
				next[l] += count
				next[r] += count
				continue
			}
			next[x*2024] += count
		}
		curr, next = next, curr
	}
	var res int
	for _, count := range curr {
		res += count
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
