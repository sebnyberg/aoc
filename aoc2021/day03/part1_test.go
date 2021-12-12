package day03

import (
	"aoc/ax"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay03Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay03Part1(t *testing.T) {
	assert.Equal(t, "198", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "775304", Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) string {
	n := len(rows[0])
	m := len(rows)
	oneCount := make([]int, n)
	for _, row := range rows {
		for i := 0; i < len(row); i++ {
			if row[i] == '1' {
				oneCount[i]++
			}
		}
	}
	var gamma, eps int
	for _, count := range oneCount {
		gamma <<= 1
		eps <<= 1
		if count*2 > m {
			gamma += 1
		} else {
			eps += 1
		}
	}
	return strconv.Itoa(gamma * eps)
}
