package p_test

import (
	"aoc/ax"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 15},
		{"input", 425},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

func run(rows []string) int {
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
	return sum
}
