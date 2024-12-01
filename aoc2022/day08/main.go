package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])
	grid := make([][]byte, m)
	vis := make([][]uint8, m)
	for i := range grid {
		grid[i] = []byte(lines[i])
		vis[i] = make([]uint8, n)
	}

	// Scan lines
	for i := range grid {
		var maxSize byte
		for j := range grid[i] {
			if grid[i][j] > maxSize {
				vis[i][j] = 1
			}
			maxSize = ax.Max(maxSize, grid[i][j])
		}
		maxSize = 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] > maxSize {
				vis[i][j] = 1
			}
			maxSize = ax.Max(maxSize, grid[i][j])
		}
	}

	// Scan columns
	for j := range grid[0] {
		var maxSize byte
		for i := range grid {
			if grid[i][j] > maxSize {
				vis[i][j] = 1
			}
			maxSize = ax.Max(maxSize, grid[i][j])
		}
		maxSize = 0
		for i := m - 1; i >= 0; i-- {
			if grid[i][j] > maxSize {
				vis[i][j] = 1
			}
			maxSize = ax.Max(maxSize, grid[i][j])
		}
	}

	// Count visible
	var res int
	for i := range vis {
		for _, v := range vis[i] {
			res += int(v)
		}
	}
	return res
}

func solve2(inf string) any {
	// Note: this is of course not how I solved the problem live. For such a
	// small problem, O(n^3) scanning is going to be practically instant.
	//
	// However, I just had to go for the O(n^2) solution using monotonic stacks.
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])
	grid := make([][]byte, m)
	score := make([][]int, m)
	for i := range grid {
		grid[i] = []byte(lines[i])
		score[i] = make([]int, n)
		for j := range score {
			score[i][j] = 1
		}
	}

	// Calculate the score
	// For each row....
	for i := range grid {
		forward := calcView(grid, i, 0, 0, 1)
		for j, v := range forward {
			score[i][j] *= v
		}
		backward := calcView(grid, i, n-1, 0, -1)
		for j, v := range backward {
			score[i][n-1-j] *= v
		}
	}
	// For each col...
	for j := range grid[0] {
		down := calcView(grid, 0, j, 1, 0)
		for i, v := range down {
			score[i][j] *= v
		}
		up := calcView(grid, m-1, j, -1, 0)
		for i, v := range up {
			score[m-1-i][j] *= v
		}
	}

	var res int
	for i := range score {
		for _, v := range score[i] {
			res = ax.Max(res, v)
		}
	}
	return res
}

func calcView(grid [][]byte, i0, j0, di, dj int) []int {
	// Starting at (i, j) and incrementing by (di, dj), while (i, j) are
	// within bounds, calculate the view of each tree.
	//
	// To do so, use an index stack where corresponding trees are
	// decreasing. When finding the view for a tree, find the first tree
	// which is >= that trees hight and calculate the distance to that tree.
	dist := func(i1, j1, i2, j2 int) int {
		return ax.Abs(i2-i1) + ax.Abs(j2-j1)
	}
	m := len(grid)
	n := len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	var res []int
	stack := [][2]int{}
	i := i0
	j := j0
	for ok(i, j) {
		var view int
		for k := len(stack) - 1; k >= 0; k-- {
			x := stack[k]
			view = dist(i, j, x[0], x[1])
			if grid[x[0]][x[1]] >= grid[i][j] {
				break
			}
			if k == 0 {
				// All of this will be yours one day, Simba!
				view = dist(i, j, i0, j0)
				break
			}
		}
		res = append(res, view)
		// Pop from the stack while it includes elements that would be obscured
		// by this tree
		for len(stack) > 0 {
			x := stack[len(stack)-1]
			if grid[x[0]][x[1]] > grid[i][j] {
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{i, j})
		i += di
		j += dj
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result1:\n%v\n", solve2(f))
}
