package aoc2021

import (
	"aoc/ax"
	"strconv"
)

func Day01Part1(lines []string) string {
	lineInts := make([]int, len(lines))
	for i := range lines {
		lineInts[i] = ax.MustParseIntBase[int](lines[i], 10)
	}
	var count int
	for i := 1; i < len(lineInts); i++ {
		if lineInts[i-1] < lineInts[i] {
			count++
		}
	}
	return strconv.Itoa(count)
}
