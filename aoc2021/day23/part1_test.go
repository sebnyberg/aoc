package day23

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day23part1 int

func BenchmarkDay23Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day23part1 = Part1(ax.MustReadFineLines("input"))
	}
}

func TestDay23Part1(t *testing.T) {
	assert.Equal(t, 12521, Part1(ax.MustReadFineLines("small")))
}

func Part1(rows []string) int {
	// Finding a greedy solution may be possible but will be very difficult.

	// We need to find a way to describe the state of the board so that outcomes
	// (costs) can be memoized

	return 0
}
