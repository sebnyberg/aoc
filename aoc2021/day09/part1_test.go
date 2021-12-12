package day09

import (
	"aoc/ax"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay09Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay09Part1(t *testing.T) {
	assert.Equal(t, "15", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "425", Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) string {
	m := len(rows)
	n := len(rows[0])
	grid := make([][]int, m)
	for i, row := range rows {
		grid[i] = make([]int, n)
		for j := range row {
			grid[i][j] = int(row[j] - '0')
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	minPoint := func(i, j int) bool {
		for _, p := range [][2]int{{i + 1, j}, {i - 1, j}, {i, j - 1}, {i, j + 1}} {
			ii, jj := p[0], p[1]
			if ok(ii, jj) && grid[i][j] >= grid[ii][jj] {
				return false
			}
		}
		return true
	}

	var sum int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if minPoint(i, j) {
				sum += grid[i][j] + 1
			}
		}
	}
	return strconv.Itoa(sum)
}
