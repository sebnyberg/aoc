package day3part2

import (
	"aoc/ax"
	"strings"
)

const (
	Problem = 3
	Part    = 2
)

func Run(lines []string) int {
	n := len(strings.TrimSpace(lines[0]))
	m := len(lines)
	bitLines := make([]int, m)
	for i, line := range lines {
		bitLines[i] = ax.MustParseIntBase[int](line, 2)
	}
	oxyLines := make(map[int]struct{})
	coLines := make(map[int]struct{})
	for _, k := range bitLines {
		oxyLines[k] = struct{}{}
		coLines[k] = struct{}{}
	}
	var coRes, oxyRes int
	for bit := 1 << (n - 1); bit > 0; bit >>= 1 {
		var oxyCount int
		for oxyLine := range oxyLines {
			if oxyLine&bit > 0 {
				oxyCount++
			}
		}
		if oxyCount*2 >= len(oxyLines) {
			for k := range oxyLines {
				if k&bit == 0 {
					delete(oxyLines, k)
				}
			}
		} else {
			for k := range oxyLines {
				if k&bit > 0 {
					delete(oxyLines, k)
				}
			}
		}
		if oxyCount == 1 {
			for k := range oxyLines {
				oxyRes = k
			}
		}
		var coCount int
		for coLine := range coLines {
			if coLine&bit > 0 {
				coCount++
			}
		}
		if coCount*2 >= len(coLines) {
			for k := range coLines {
				if k&bit > 0 {
					delete(coLines, k)
				}
			}
		} else {
			for k := range coLines {
				if k&bit == 0 {
					delete(coLines, k)
				}
			}
		}
		if coCount == 1 {
			for k := range coLines {
				coRes = k
			}
		}
	}
	return oxyRes * coRes
}
