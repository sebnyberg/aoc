package day23

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day23part2 int

func BenchmarkDay23Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day23part1 = Part2(ax.MustReadFineLines("input"))
	}
}

func TestDay23Part2(t *testing.T) {
	assert.Equal(t, 44169, Part2(ax.MustReadFineLines("smallpart2")))
	assert.Equal(t, 43226, Part2(ax.MustReadFineLines("inputpart2")))
}

func Part2(rows []string) int {
	return 0
}
