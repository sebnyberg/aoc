package day15

import (
	"container/heap"
	"math"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day15part2res int

func BenchmarkDay15Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day15part2res = Part2(ax.MustReadFileLines("input"))
	}
}

func TestDay15Part2(t *testing.T) {
	assert.Equal(t, 315, Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, 2831, Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) int {
	// Use Dijkstra's
	m := len(rows)
	n := len(rows[0])
	var grid [500][500]byte
	var seen [500][500]bool
	var dist [500][500]uint16
	for i, row := range rows {
		for j := range row {
			// Copy board + adjust values
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					ii, jj := k*m+i, l*n+j
					grid[ii][jj] = rows[i][j] - '0' + byte(k+l)
					dist[ii][jj] = math.MaxUint16
					if grid[ii][jj] > 9 {
						grid[ii][jj] %= 9
					}
				}
			}
		}
	}
	mm := 5 * m
	nn := 5 * n
	ok := func(i, j int) bool {
		return i >= 0 && i < mm && j >= 0 && j < nn
	}

	// Keep a min heap of distances
	h := make(minHeap, 1, 100)
	h[0] = pos{0, 0, 0}

	// While there are entries in the min heap (always true for this)
	var maxLen int
	for {
		maxLen = max(maxLen, len(h))
		x := heap.Pop(&h).(pos)
		seen[x.i][x.j] = true
		if x.i == mm-1 && x.j == nn-1 {
			return int(maxLen)
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

type pos struct {
	i, j int
	val  uint16
}
type minHeap []pos

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pos))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
