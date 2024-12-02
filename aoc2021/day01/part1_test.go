package day01

import (
	"strconv"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay01Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay01Part1(t *testing.T) {
	assert.Equal(t, "7", Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, "1292", Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) string {
	nums := make([]int, len(rows))
	for i := range rows {
		nums[i] = ax.MustParseIntBase[int](rows[i], 10)
	}
	var count int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			count++
		}
	}
	return strconv.Itoa(count)
}
