package day15

import (
	"container/heap"
	"math"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day15part1 int

func BenchmarkDay15Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day15part1 = Part1(ax.MustReadFileLines("input"))
	}
}

func TestDay15Part1(t *testing.T) {
	assert.Equal(t, 40, Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, 456, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	// Use Dijkstra's
	m := len(rows)
	n := len(rows[0])
	var grid [100][100]byte
	var seen [100][100]bool
	var dist [100][100]uint16
	for i, row := range rows {
		for j := range row {
			grid[i][j] = rows[i][j] - '0'
			dist[i][j] = math.MaxUint16
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	// Keep a min heap of distances
	h := make(minHeap, 1, 100)
	h[0] = pos{0, 0, 0}

	// While there are entries in the min heap (always true for this)
	for {
		x := heap.Pop(&h).(pos)
		seen[x.i][x.j] = true
		if x.i == m-1 && x.j == n-1 {
			return int(x.val)
		}
		for _, nei := range [][2]int{
			{x.i + 1, x.j}, {x.i - 1, x.j}, {x.i, x.j - 1}, {x.i, x.j + 1},
		} {
			ii, jj := nei[0], nei[1]
			if !ok(ii, jj) || seen[ii][jj] {
				continue
			}
			risk := x.val + uint16(grid[ii][jj])
			if risk >= dist[ii][jj] {
				continue
			}
			dist[ii][jj] = risk
			heap.Push(&h, pos{ii, jj, risk})
		}
	}
}
