package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	curr := [][2]int{}
	next := [][2]int{}
	var dirs = [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	bfs := func(i, j int) (int, int) {
		seen[i][j] = true
		curr = append(curr, [2]int{i, j})
		var perim int
		area := 1
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if !ok(ii, jj) || lines[ii][jj] != lines[i][j] {
						perim++
						continue
					}
					if seen[ii][jj] {
						continue
					}
					seen[ii][jj] = true
					next = append(next, [2]int{ii, jj})
					area++
				}
			}
			curr, next = next, curr
		}
		return area, perim
	}
	var res int
	for i := range lines {
		for j := range lines[i] {
			if seen[i][j] {
				continue
			}
			area, perim := bfs(i, j)
			res += area * perim
		}
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])
	seen := make([][]bool, m)
	region := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
		region[i] = make([]bool, n)
	}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	var left = [2]int{0, -1}
	var up = [2]int{-1, 0}
	var right = [2]int{0, 1}
	var down = [2]int{1, 0}
	var dirs = [][2]int{right, down, left, up}

	curr := [][2]int{}
	next := [][2]int{}
	bfs := func(start [2]int) (int, int, int) {
		// reset region
		for i := range region {
			for j := range region[i] {
				region[i][j] = false
			}
		}
		var perims [][2][2]int
		seen[start[0]][start[1]] = true
		curr = append(curr, start)
		area := 1
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if !ok(ii, jj) || lines[ii][jj] != lines[start[0]][start[1]] {
						perims = append(perims, [2][2]int{x, {ii, jj}})
						continue
					}
					if seen[ii][jj] {
						continue
					}
					seen[ii][jj] = true
					next = append(next, [2]int{ii, jj})
					area++
				}
			}
			curr, next = next, curr
		}
		// For each perimiter pair, eliminate pairs sharing the same side
		sidePairs := make(map[[2][2]int]struct{})
		for _, p := range perims {
			sidePairs[p] = struct{}{}
		}
		var sides int
		move := func(p, dir [2]int) [2]int {
			return [2]int{
				p[0] + dir[0],
				p[1] + dir[1],
			}
		}
		for p := range sidePairs {
			sides++
			delete(sidePairs, p)
			// figure out direction based on pair
			dir := [2]int{p[1][0] - p[0][0], p[1][1] - p[0][1]}
			// move perpendicular, removing pairs until there are no more
			var dir1, dir2 [2]int
			switch dir {
			case left, right:
				dir1 = up
				dir2 = down
			case up, down:
				dir1 = right
				dir2 = left
			}
			for _, d := range [][2]int{dir1, dir2} {
				cur1 := p
				for {
					cur1 = [2][2]int{move(cur1[0], d), move(cur1[1], d)}
					_, exists := sidePairs[cur1]
					if !exists {
						break
					}
					delete(sidePairs, cur1)
				}
			}
		}
		return area, len(perims), sides
	}
	var res int
	for i := range lines {
		for j := range lines[i] {
			if seen[i][j] {
				continue
			}
			area, _, sides := bfs([2]int{i, j})
			res += area * sides
		}
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
