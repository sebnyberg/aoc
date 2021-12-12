package p_test

import (
	"aoc/ax"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 1134},
		{"input", 1135260},
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

	// Helpers
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] != 9
	}
	to1d := func(i, j int) int {
		return i*n + j
	}

	// DSU with path compression and set size
	parent := make([]int, m*n)
	size := make([]int, m*n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if x != parent[x] {
			return find(parent[x])
		}
		return x
	}
	union := func(x, y int) {
		rootX, rootY := find(x), find(y)
		if rootX != rootY {
			parent[rootY] = rootX
			size[rootX] += size[rootY]
			size[rootY] = 1
		}
	}

	// For each cell, visit all neighbours.
	seen := make([]bool, n*m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if seen[to1d(i, j)] || !ok(i, j) {
				continue
			}
			seen[to1d(i, j)] = true
			for _, near := range [][2]int{
				{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1},
			} {
				ii, jj := near[0], near[1]
				if ok(ii, jj) && !seen[to1d(ii, jj)] {
					union(to1d(i, j), to1d(ii, jj))
				}
			}
		}
	}

	// Trim away 1-size groups
	var i int
	for _, sz := range size {
		if sz > 1 {
			size[i] = sz
			i++
		}
	}
	size = size[:i]

	sort.Ints(size)
	sum := size[len(size)-1]
	for _, sz := range size[len(size)-3 : len(size)-1] {
		sum *= sz
	}

	return sum
}
