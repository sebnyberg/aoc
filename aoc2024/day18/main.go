package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

type cell struct {
	blocked bool
	minDist int
}

var dirs = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type point struct {
	x, y int
}

func solve1(inf string, n, nbytes int) any {
	points := parsePoints(inf)[:nbytes]
	dist := solve(points, n)
	return dist
}

func solve2(inf string, n int) any {
	points := parsePoints(inf)
	for i := range points {
		dist := solve(points[:i+1], n)
		if dist == -1 {
			return points[i]
		}
	}
	return 0
}

func parsePoints(inf string) []point {
	lines := ax.MustReadFileLines(inf)
	var points []point
	for _, l := range lines {
		parts := strings.Split(l, ",")
		x := ax.Atoi(parts[0])
		y := ax.Atoi(parts[1])
		points = append(points, point{x, y})
	}
	return points
}

func solve(points []point, n int) int {
	blocked := make([][]bool, n)
	for i := range blocked {
		blocked[i] = make([]bool, n)
	}
	for _, p := range points {
		blocked[p.x][p.y] = true
	}

	// Dijkstra's
	minDist := make([][]int, n)
	for x := range minDist {
		minDist[x] = make([]int, n)
		for y := range minDist[x] {
			minDist[x][y] = math.MaxInt32
		}
	}
	h := minHeap{pos{0, 0, 0}}
	minDist[0][0] = 0
	ok := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n && !blocked[i][j]
	}
	for len(h) > 0 {
		x := heap.Pop(&h).(pos)
		if x.x == n-1 && x.y == n-1 {
			return x.dist
		}
		for _, d := range dirs {
			ii := x.x + d.x
			jj := x.y + d.y
			if !ok(ii, jj) || x.dist+1 >= minDist[ii][jj] {
				continue
			}
			heap.Push(&h, pos{ii, jj, x.dist + 1})
			minDist[ii][jj] = x.dist + 1
		}
	}
	return -1
}

type pos struct {
	x, y int
	dist int
}

type minHeap []pos

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h minHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
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

func main() {
	// fmt.Printf("Result1:\n%v\n", solve1("testinput", 7, 12))
	// fmt.Printf("Result1:\n%v\n", solve1("input", 71, 1024))
	fmt.Printf("Result2:\n%v\n", solve2("testinput", 7))
	fmt.Printf("Result2:\n%v\n", solve2("input", 71))
	// fmt.Printf("Result2:\n%v\n", solve2("input"))
}
