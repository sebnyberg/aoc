package day03

import (
	"aoc/ax"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay03Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay03Part2(t *testing.T) {
	assert.Equal(t, "230", Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, "1370737", Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) string {
	n := len(strings.TrimSpace(rows[0]))
	m := len(rows)
	nums := make([]int, m)
	for i, row := range rows {
		nums[i] = ax.MustParseIntBase[int](row, 2)
	}

	calculate := func(nums map[int]struct{}, inv bool) int {
		for bit := 1 << (n - 1); bit > 0; bit >>= 1 {
			var count int
			for num := range nums {
				if num&bit > 0 {
					count++
				}
			}
			var want int
			if inv {
				want = bit
			}
			if count*2 >= len(nums) {
				want ^= bit
			}
			for k := range nums {
				if k&bit != want {
					delete(nums, k)
				}
			}
			if len(nums) == 1 {
				for k := range nums {
					return k
				}
			}
		}
		panic("not found")
	}

	oxyLines := make(map[int]struct{})
	coLines := make(map[int]struct{})
	for _, num := range nums {
		oxyLines[num] = struct{}{}
		// Invert numbers for CO2
		coLines[num] = struct{}{}
	}

	coRes := calculate(coLines, true)
	oxyRes := calculate(oxyLines, false)
	return strconv.Itoa(oxyRes * coRes)
}
