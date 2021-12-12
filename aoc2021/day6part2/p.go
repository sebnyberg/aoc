package day6part2

import (
	"aoc/ax"
	"strconv"
	"strings"
)

const (
	Problem = 6
	Part    = 1
	ndays   = 256
)

func Run(rows []string) string {
	var fishCount [9]int
	for _, valStr := range strings.Split(rows[0], ",") {
		val := ax.MustParseInt[int](valStr)
		fishCount[val]++
	}
	var nextCount [9]int
	for day := 0; day < ndays; day++ {
		for i := 0; i < 8; i++ {
			nextCount[i] = fishCount[(i+1)%9]
		}
		nextCount[6] += fishCount[0]
		nextCount[8] = fishCount[0]
		nextCount, fishCount = fishCount, nextCount
	}
	return strconv.Itoa(ax.Sum(fishCount[:]))
}
