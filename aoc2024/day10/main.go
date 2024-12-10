package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])
	scores := make([][]int, m)
	for i := range scores {
		scores[i] = make([]int, n)
		for j := range scores[i] {
			scores[i][j] = -1 // unseen
		}
	}
	var curr [][2]int
	var next [][2]int
	var dirs = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	findTrailHeads := func(i, j int) int {
		curr = append(curr[:0], [2]int{i, j})
		// reset seen
		for k := range seen {
			for kk := range seen[k] {
				seen[k][kk] = false
			}
		}
		seen[i][j] = true
		for k := byte('1'); k <= '9' && len(curr) > 0; k++ {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if ii < 0 || ii >= m || jj < 0 || jj >= n || lines[ii][jj] != k || seen[ii][jj] {
						continue
					}
					seen[ii][jj] = true
					next = append(next, [2]int{ii, jj})
				}
			}
			curr, next = next, curr
		}
		return len(curr)
	}
	var res int
	for i := range lines {
		for j, v := range lines[i] {
			if v == '0' {
				res += findTrailHeads(i, j)
			}
		}
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])
	scores := make([][]int, m)
	for i := range scores {
		scores[i] = make([]int, n)
		for j := range scores[i] {
			scores[i][j] = -1 // unseen
		}
	}
	startPos := [][2]int{}
	for i := range lines {
		for j, v := range lines[i] {
			if v == '0' {
				startPos = append(startPos, [2]int{i, j})
			}
		}
	}
	var res int
	for _, p := range startPos {
		res += dfs(scores, lines, p[0], p[1], m, n)
	}
	return res
}

var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(scores [][]int, lines []string, i, j, m, n int) int {
	if scores[i][j] != -1 {
		return scores[i][j]
	}
	if lines[i][j] == '9' {
		return 1
	}
	var score int
	for _, d := range dirs {
		ii := i + d[0]
		jj := j + d[1]
		if ii < 0 || ii >= m || jj < 0 || jj >= n {
			continue
		}
		if lines[ii][jj]-lines[i][j] != 1 {
			continue
		}
		score += dfs(scores, lines, ii, jj, m, n)
	}
	scores[i][j] = score
	return scores[i][j]
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
