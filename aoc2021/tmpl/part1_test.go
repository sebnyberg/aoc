package dayx

import (
	"aoc/ax"
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
	assert.Equal(t, "11111", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "11111", Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) string {
	return ""
}
