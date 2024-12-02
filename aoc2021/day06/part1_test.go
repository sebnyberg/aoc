package day06

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay06Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay06Part1(t *testing.T) {
	assert.Equal(t, "5934", Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, "395627", Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) string {
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
	return strconv.Itoa(res)
}
