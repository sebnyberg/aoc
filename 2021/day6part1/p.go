package day6part1

import (
	"aoc/ax"
	"strings"
)

const (
	Problem = 6
	Part    = 1
)

func Run(rows []string) int {
	var fishCount [9]int
	for _, valStr := range strings.Split(rows[0], ",") {
		val := ax.MustParseInt[int](valStr)
		fishCount[val]++
	}
	var nextCount [9]int
	for day := 0; day < 80; day++ {
		for i := 0; i < 8; i++ {
			nextCount[i] = fishCount[(i+1)%9]
		}
		nextCount[6] += fishCount[0]
		nextCount[8] = fishCount[0]
		nextCount, fishCount = fishCount, nextCount
	}
	res := ax.Sum(fishCount[:])
	return res
}
