package day01

import (
	"aoc/ax"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkPart1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestPart1(t *testing.T) {
	assert.Equal(t, "7", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "1292", Part1(ax.MustReadFineLines("input")))
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
