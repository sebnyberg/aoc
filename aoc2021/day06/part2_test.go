package day06

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay06Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay06Part2(t *testing.T) {
	assert.Equal(t, "26984457539", Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, "1767323539209", Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) string {
	const ndays = 256
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
