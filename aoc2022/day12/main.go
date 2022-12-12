package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func solve(inf string, appendAs bool) any {
	lines := ax.MustReadFileLines(inf)
	n := len(lines[0])
	m := len(lines)
	grid := make([][]byte, m)
	seen := make([][]bool, m)
	var endI, endJ int
	curr := [][2]int{}
	next := [][2]int{}
	for i := range grid {
		grid[i] = []byte(lines[i])
		seen[i] = make([]bool, n)
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				curr = append(curr, [2]int{i, j})
				grid[i][j] = 'z'
				continue
			}
			if appendAs && grid[i][j] == 'a' {
				curr = append(curr, [2]int{i, j})
				continue
			}
			if grid[i][j] == 'E' {
				endI = i
				endJ = j
				grid[i][j] = 'z'
			}
		}
	}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	for _, x := range curr {
		seen[x[0]][x[1]] = true
	}

	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			i := x[0]
			j := x[1]
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				if !ok(ii, jj) || seen[ii][jj] || (grid[ii][jj] > grid[i][j]+1) {
					continue
				}
				seen[ii][jj] = true
				if ii == endI && jj == endJ {
					return steps
				}
				next = append(next, [2]int{ii, jj})
			}
		}
		curr, next = next, curr
	}
	return "not found"
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, false))
	fmt.Printf("Result2:\n%v\n", solve(f, true))
}
