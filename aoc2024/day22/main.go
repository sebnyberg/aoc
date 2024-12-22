package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

const mod = 16777216

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var res int
	for _, l := range lines {
		x := ax.MustParseInt[int](l)
		for k := 0; k < 2000; k++ {
			x = ((x * 64) ^ x) % mod
			x = ((x / 32) ^ x) % mod
			x = ((x * 2048) ^ x) % mod
		}
		res += x
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var combs [][4]int
	for first := -9; first <= 9; first++ {
		for second := -9; second <= 9; second++ {
			for third := -9; third <= 9; third++ {
				for fourth := -9; fourth <= 9; fourth++ {
					combs = append(combs, [4]int{first, second, third, fourth})
				}
			}
		}
	}
	seqResult := make(map[[4]int]int)
	seen := make(map[[4]int]bool)
	for _, l := range lines {
		x := ax.MustParseInt[int](l)
		for k := range seen {
			delete(seen, k)
		}
		var changes [4]int
		for k := 0; k < 2000; k++ {
			next := x
			next = ((next * 64) ^ next) % mod
			next = ((next / 32) ^ next) % mod
			next = ((next * 2048) ^ next) % mod
			for i := 0; i < 3; i++ {
				changes[i] = changes[i+1]
			}
			changes[3] = next%10 - x%10
			if k >= 4 && !seen[changes] {
				seen[changes] = true
				seqResult[changes] += next % 10
			}
			x = next
		}
	}
	var res int
	for _, v := range seqResult {
		res = max(res, v)
	}
	return res
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve1("input"))
	fmt.Printf("Result2:\n%v\n", solve2("input"))
}
