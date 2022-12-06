package main

import (
	"fmt"
	"unicode"

	"github.com/sebnyberg/aoc/ax"
)

func solve(inf string, goback bool) int {
	lines := ax.MustReadFileLines(inf)
	grid := make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	var ii, jj int
	var maxval int
	for i := range grid {
		for j, v := range grid[i] {
			if unicode.IsDigit(rune(v)) {
				maxval = ax.Max(maxval, int(v-'0'))
			}
			if grid[i][j] == '0' {
				ii = i
				jj = j
			}
		}
	}
	initial := state{
		bm: 1,
		i:  ii,
		j:  jj,
	}
	curr := []state{initial}
	next := []state{}
	m := len(grid)
	n := len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] != '#'
	}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	seen := make(map[state]bool)
	seen[initial] = true
	want := (1 << (maxval + 1)) - 1
	// var final []state
	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, s := range curr {
			for _, d := range dirs {
				ii := s.i + d[0]
				jj := s.j + d[1]
				if !ok(ii, jj) {
					continue
				}
				bm := s.bm
				if grid[ii][jj] != '.' {
					bm |= (1 << int(grid[ii][jj]-'0'))
				}
				nextState := state{
					bm: bm,
					i:  ii,
					j:  jj,
				}
				if bm == want && (!goback || (ii == initial.i && jj == initial.j)) {
					return steps
				}
				if seen[nextState] {
					continue
				}
				seen[nextState] = true
				next = append(next, nextState)
			}
		}
		curr, next = next, curr
	}
	panic("out of places to visit")
}

type state struct {
	bm int
	i  int
	j  int
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, false))
	fmt.Printf("Result2:\n%v\n", solve(f, true))
}
