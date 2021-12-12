package day13

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay13Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay13Part1(t *testing.T) {
	assert.Equal(t, "10", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "3510", Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) string {
	return ""
}
