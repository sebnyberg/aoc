package day18

import (
	"testing"

	"github.com/sebnyberg/aoc/ax"
)

var day18part2res int

func BenchmarkDay18Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day18part2res = Part2(ax.MustReadFileLines("input")[0])
	}
}

func TestDay18Part2(t *testing.T) {
}

func Part2(row string) int {
	return 0
}
