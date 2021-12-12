package day1part2

import (
	"aoc/ax"
	"strconv"
)

const (
	Problem = 1
	Part    = 2
)

func Run(lines []string) string {
	lineInts := make([]int, len(lines))
	for i := range lines {
		lineInts[i] = ax.MustParseIntBase[int](lines[i], 10)
	}
	var count int
	var cur, prev int
	for i := 0; i < len(lineInts); i++ {
		cur += lineInts[i]
		if i > 2 {
			cur -= lineInts[i-3]
			if prev < cur {
				count++
			}
		}
		prev = cur
	}
	return strconv.Itoa(count)
}
