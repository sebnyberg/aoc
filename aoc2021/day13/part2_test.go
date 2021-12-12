package day13

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay13Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay13Part2(t *testing.T) {
	assert.Equal(t, "36", Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, "132880", Part2(ax.MustReadFineLines("input")))
}

func Part2(rows []string) string {
	return ""
}
