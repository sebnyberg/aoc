package day01

import (
	"aoc/ax"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkPart2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "5", Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, "1262", Part2(ax.MustReadFineLines("input")))
}

func Part2(rows []string) string {
	nums := make([]int, len(rows))
	for i := range rows {
		nums[i] = ax.MustParseIntBase[int](rows[i], 10)
	}
	var count int
	var cur, prev int
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if i > 2 {
			cur -= nums[i-3]
			if prev < cur {
				count++
			}
		}
		prev = cur
	}
	return strconv.Itoa(count)
}
