package day01

import (
	"aoc/ax"
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
	assert.Equal(t, "11111", Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, "11111", Part2(ax.MustReadFineLines("input")))
}

func Part2(lines []string) string {
	return ""
}
