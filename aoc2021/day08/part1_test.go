package day08

import (
	"aoc/ax"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay08Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay08Part1(t *testing.T) {
	assert.Equal(t, "26", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "375", Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) string {
	var count int
	for _, row := range rows {
		parts := strings.Split(row, "|")
		outputFields := strings.Fields(parts[1])
		for _, field := range outputFields {
			if len(field) == 2 || len(field) == 4 || len(field) == 3 || len(field) == 7 {
				count++
			}
		}
	}
	return strconv.Itoa(count)
}
