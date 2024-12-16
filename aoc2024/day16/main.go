package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/sebnyberg/aoc/ax"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pos struct {
	i, j  int
	dir   int
	score int
}

type minHeap []pos

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	a := h[i].score
	b := h[j].score
	if a == b {
		if h[i].i == h[j].i {
			return h[i].j < h[j].j
		}
		return h[i].i < h[j].i
	}
	return h[i].score < h[j].score
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

func print(state [][]byte) {
	for i := range state {
		fmt.Println(string(state[i]))
	}
	fmt.Println("")
}

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	state := make([][]byte, len(lines))
	var start, end [2]int
	for i, l := range lines {
		state[i] = []byte(l)
		for j, v := range l {
			if v == 'S' {
				start = [2]int{i, j}
			} else if v == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	m := len(state)
	n := len(state[0])
	minScore := make([][][4]int, m)
	for i := range minScore {
		minScore[i] = make([][4]int, n)
		for j := range minScore[i] {
			for k := range minScore[i][j] {
				minScore[i][j][k] = math.MaxInt32
			}
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && state[i][j] != '#'
	}

	h := minHeap{pos{start[0], start[1], 0, 0}}
	minScore[start[0]][start[1]][0] = 0
	fmt.Println(m * n * 4)
	for i := 0; len(h) > 0; i++ {
		x := heap.Pop(&h).(pos)
		if x.i == end[0] && x.j == end[1] {
			return x.score
		}
		// move forward
		ii := x.i + dirs[x.dir][0]
		jj := x.j + dirs[x.dir][1]
		if ok(ii, jj) && minScore[ii][jj][x.dir] > x.score+1 {
			minScore[ii][jj][x.dir] = x.score + 1
			heap.Push(&h, pos{ii, jj, x.dir, x.score + 1})
		}
		// turn left
		left := (x.dir + 3) % 4
		if minScore[x.i][x.j][left] > x.score+1000 {
			minScore[x.i][x.j][left] = x.score + 1000
			heap.Push(&h, pos{x.i, x.j, left, x.score + 1000})
		}
		// turn right
		right := (x.dir + 1) % 4
		if minScore[x.i][x.j][right] > x.score+1000 {
			minScore[x.i][x.j][right] = x.score + 1000
			heap.Push(&h, pos{x.i, x.j, right, x.score + 1000})
		}
	}
	return -1
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	state := make([][]byte, len(lines))
	var start, end [2]int
	for i, l := range lines {
		state[i] = []byte(l)
		for j, v := range l {
			if v == 'S' {
				start = [2]int{i, j}
			} else if v == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	m := len(state)
	n := len(state[0])
	minScore := make([][][4]int, m)
	prev := make([][][4][][3]int, m)
	for i := range minScore {
		minScore[i] = make([][4]int, n)
		prev[i] = make([][4][][3]int, n)
		for j := range minScore[i] {
			for k := range minScore[i][j] {
				minScore[i][j][k] = math.MaxInt32
			}
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && state[i][j] != '#'
	}

	countPaths := func(i, j, dir int) int {
		seen := make([][]bool, m)
		seenDir := make([][][4]bool, m)
		for i := range seen {
			seen[i] = make([]bool, n)
			seenDir[i] = make([][4]bool, n)
		}
		count := 1
		seen[i][j] = true
		seenDir[i][j][dir] = true
		curr := [][3]int{{i, j, dir}}
		next := [][3]int{}
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, y := range prev[x[0]][x[1]][x[2]] {
					if seenDir[y[0]][y[1]][y[2]] {
						continue
					}
					seenDir[y[0]][y[1]][y[2]] = true
					if !seen[y[0]][y[1]] {
						count++
					}
					seen[y[0]][y[1]] = true
					next = append(next, y)
				}
			}
			curr, next = next, curr
		}
		return count
	}

	h := minHeap{pos{start[0], start[1], 0, 0}}
	minScore[start[0]][start[1]][0] = 0
	fmt.Println(m * n * 4)
	for i := 0; len(h) > 0; i++ {
		x := heap.Pop(&h).(pos)
		if x.i == end[0] && x.j == end[1] {
			res := countPaths(x.i, x.j, x.dir)
			return res
		}
		// turn left
		left := (x.dir + 3) % 4
		if minScore[x.i][x.j][left] == x.score+1000 {
			prev[x.i][x.j][left] = append(prev[x.i][x.j][left], [3]int{x.i, x.j, x.dir})
		} else if minScore[x.i][x.j][left] > x.score+1000 {
			// reset prev
			prev[x.i][x.j][left] = append(prev[x.i][x.j][left][:0], [3]int{x.i, x.j, x.dir})
			minScore[x.i][x.j][left] = x.score + 1000
			heap.Push(&h, pos{x.i, x.j, left, x.score + 1000})
		}
		// turn right
		right := (x.dir + 1) % 4
		if minScore[x.i][x.j][right] == x.score+1000 {
			prev[x.i][x.j][right] = append(prev[x.i][x.j][right], [3]int{x.i, x.j, x.dir})
		} else if minScore[x.i][x.j][right] > x.score+1000 {
			// reset prev
			prev[x.i][x.j][right] = append(prev[x.i][x.j][right][:0], [3]int{x.i, x.j, x.dir})
			minScore[x.i][x.j][right] = x.score + 1000
			heap.Push(&h, pos{x.i, x.j, right, x.score + 1000})
		}

		// move forward
		ii := x.i + dirs[x.dir][0]
		jj := x.j + dirs[x.dir][1]
		if !ok(ii, jj) {
			continue
		}
		if minScore[ii][jj][x.dir] == x.score+1 {
			prev[ii][jj][x.dir] = append(prev[ii][jj][x.dir], [3]int{x.i, x.j, x.dir})
		} else if minScore[ii][jj][x.dir] > x.score+1 {
			// reset prev
			prev[ii][jj][x.dir] = append(prev[ii][jj][x.dir][:0], [3]int{x.i, x.j, x.dir})
			minScore[ii][jj][x.dir] = x.score + 1
			heap.Push(&h, pos{ii, jj, x.dir, x.score + 1})
		}
	}
	return -1
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve1("input"))
	fmt.Printf("Result2:\n%v\n", solve2("input"))
}
