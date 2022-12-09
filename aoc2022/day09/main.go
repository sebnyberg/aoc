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
	seen := make(map[[2]int]bool)
	deltas := map[string][2]int{
		"R": {0, 1},
		"L": {0, -1},
		"U": {1, 0},
		"D": {-1, 0},
	}
	diags := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, l := range lines {
		fs := strings.Fields(l)
		d := deltas[fs[0]]
		steps := ax.Atoi(fs[1])
		for k := 0; k < steps; k++ {
			rope[0][0] += d[0]
			rope[0][1] += d[1]
			for t := 1; t < ropelen; t++ {
				dj := rope[t-1][1] - rope[t][1]
				di := rope[t-1][0] - rope[t][0]
				if dj == 0 {
					if di >= 2 {
						rope[t][0]++
					} else if di <= -2 {
						rope[t][0]--
					}
					continue
				}
				if di == 0 {
					if dj >= 2 {
						rope[t][1]++
					} else if dj <= -2 {
						rope[t][1]--
					}
					continue
				}
				if ax.Abs(di) <= 1 && ax.Abs(dj) <= 1 {
					continue
				}
				// Find a diagonal that puts the tail in the right spot
				for _, diag := range diags {
					i := rope[t][0] + diag[0]
					j := rope[t][1] + diag[1]
					di := rope[t-1][0] - i
					dj := rope[t-1][1] - j
					if ax.Abs(di) <= 1 && ax.Abs(dj) <= 1 {
						rope[t] = [2]int{i, j}
						break
					}
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
