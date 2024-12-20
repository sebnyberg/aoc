package main

import (
	"fmt"
	"math"

	"github.com/sebnyberg/aoc/ax"
)

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func findDistances(state []string, i, j, startDist int) [][]int {
	n := len(state)
	minDist := make([][]int, n)
	for i := range minDist {
		minDist[i] = make([]int, n)
		for j := range minDist[i] {
			minDist[i][j] = math.MaxInt32
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n && state[i][j] != '#'
	}
	minDist[i][j] = startDist
	curr := [][2]int{{i, j}}
	next := [][2]int{}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			i := x[0]
			j := x[1]
			for _, d := range dirs {
				ii := i + d[0]
				jj := j + d[1]
				newDist := minDist[i][j] + 1
				if !ok(ii, jj) || newDist >= minDist[ii][jj] {
					continue
				}
				minDist[ii][jj] = newDist
				next = append(next, [2]int{ii, jj})
			}
		}
		curr, next = next, curr
	}
	return minDist
}

func solve(inf string, cheatTime int) any {
	lines := ax.MustReadFileLines(inf)
	var start, end [2]int
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'S' {
				start = [2]int{i, j}
			}
			if lines[i][j] == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	n := len(lines)
	ok := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n && lines[i][j] != '#'
	}
	dists := findDistances(lines, start[0], start[1], 0)
	normalTime := dists[end[0]][end[1]]
	timeSaved := map[int]int{}

	timeLeft := make([][]int, n)
	for i := range timeLeft {
		timeLeft[i] = make([]int, n)
		for j := range timeLeft[i] {
			if timeLeft[i][j] != math.MaxInt32 {
				timeLeft[i][j] = normalTime - dists[i][j]
			}
		}
	}

	for i := range dists {
		for j := range dists[i] {
			if dists[i][j] == math.MaxInt32 {
				continue
			}
			// Check whether cheating from this position saves any time
			for ii := i - cheatTime; ii <= i+cheatTime; ii++ {
				for jj := j - cheatTime; jj <= j+cheatTime; jj++ {
					di := abs(ii - i)
					dj := abs(jj - j)
					k := di + dj
					if k > cheatTime || k == 0 {
						continue
					}
					if !ok(ii, jj) {
						continue
					}
					newTime := dists[i][j] + k + timeLeft[ii][jj]
					delta := normalTime - newTime
					if delta > 0 {
						timeSaved[delta]++
					}
				}
			}
		}
	}
	var res int
	for k, v := range timeSaved {
		if k >= 100 {
			res += v
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve("input", 2))
	fmt.Printf("Result2:\n%v\n", solve("input", 20))
}
