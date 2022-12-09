package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve(inf string, ropelen int) any {
	lines := ax.MustReadFileLines(inf)
	rope := make([][2]int, ropelen)
	for i := 0; i < ropelen; i++ {
		rope[i] = [2]int{0, 0}
	}
	deltas := map[string][2]int{
		"R": {0, 1},
		"L": {0, -1},
		"U": {1, 0},
		"D": {-1, 0},
	}
	seen := make(map[[2]int]bool)
	for _, l := range lines {
		fs := strings.Fields(l)
		d := deltas[fs[0]]
		steps := ax.Atoi(fs[1])
		for k := 0; k < steps; k++ {
			// Move head
			rope[0][0] += d[0]
			rope[0][1] += d[1]
			for t := 1; t < ropelen; t++ {
				// If the tail is more than 1 distance away in either direction
				dy := rope[t-1][0] - rope[t][0]
				dx := rope[t-1][1] - rope[t][1]
				if ax.Abs(dx) > 1 || ax.Abs(dy) > 1 {
					// Adjust it by at most 1 toward the head
					rope[t][0] += ax.Max(-1, ax.Min(1, dy))
					rope[t][1] += ax.Max(-1, ax.Min(1, dx))
				}
			}
			seen[rope[ropelen-1]] = true
		}
	}
	return len(seen)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, 2))
	fmt.Printf("Result2:\n%v\n", solve(f, 10))
}
