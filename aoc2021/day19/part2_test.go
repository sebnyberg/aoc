package day19

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day19part2res int

func BenchmarkDay19Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day19part2res = Part2(ax.MustReadFineLines("input")[0])
	}
}

func TestDay19Part2(t *testing.T) {
	assert.Equal(t, 45, Part2("target area: x=20..30, y=-10..-5"))
	assert.Equal(t, 11175, Part2(ax.MustReadFineLines("input")[0]))
}

func Part2(row string) int {
	return 0
}
