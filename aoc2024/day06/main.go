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
	var pos [2]int
	for i, l := range lines {
		for j, ch := range l {
			if ch == '^' {
				pos = [2]int{i, j}
			}
		}
	}
	ii := pos[0]
	jj := pos[1]
	dirs := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dir := 0
	var res int
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	for ok(ii, jj) {
		if !seen[ii][jj] {
			res++
			seen[ii][jj] = true
		}

		nextII := ii + dirs[dir][0]
		nextJJ := jj + dirs[dir][1]
		if !ok(nextII, nextJJ) {
			break
		}
		if ok(nextII, nextJJ) && lines[nextII][nextJJ] != '#' {
			ii = nextII
			jj = nextJJ
			continue
		}

		dir = (dir + 1) % 4 // turn
	}
	return res
}

func hasLoop(lines [][]byte, i, j int) bool {
	m := len(lines)
	n := len(lines[0])
	ii := i
	jj := j
	dirs := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dir := 0
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	dup := make(map[[3]int]struct{})
	for ok(ii, jj) {
		k := [3]int{ii, jj, dir}
		if _, exists := dup[k]; exists {
			return true
		}
		dup[k] = struct{}{}

		nextII := ii + dirs[dir][0]
		nextJJ := jj + dirs[dir][1]
		if !ok(nextII, nextJJ) {
			break
		}
		if ok(nextII, nextJJ) && lines[nextII][nextJJ] != '#' {
			ii = nextII
			jj = nextJJ
			continue
		}

		dir = (dir + 1) % 4 // turn
	}
	return false
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var pos [2]int
	for i, l := range lines {
		for j, ch := range l {
			if ch == '^' {
				pos = [2]int{i, j}
			}
		}
	}
	var res int
	m := len(lines)
	bs := make([][]byte, m)
	for i := range bs {
		bs[i] = []byte(lines[i])
	}
	for i := range bs {
		for j := range bs[i] {
			if bs[i][j] == '#' || bs[i][j] == '^' {
				continue
			}
			bs[i][j] = '#'
			if hasLoop(bs, pos[0], pos[1]) {
				res++
			}
			bs[i][j] = '.'
		}
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
