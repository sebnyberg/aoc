package day19

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day19part1 int

func BenchmarkDay19Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// day19part1 = Part1(ax.MustReadFineLines("input")[0])
	}
}

func TestDay19Part1(t *testing.T) {
	assert.Equal(t, 11195, Part1(ax.MustReadFineLines("small2")))
}

func Part1(rows []string) int {
	return 0
}
