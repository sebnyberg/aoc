package main

import (
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	rs := in.xs
	m := len(rs)
	n := len(rs[0].s)
	grid := make([][]byte, m)
	next := make([][]byte, m)
	for i := range grid {
		grid[i] = []byte(rs[i].s)
		next[i] = make([]byte, n)
	}

	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for steps := 0; steps < 100; steps++ {
		for i := range next {
			next[i] = next[i][:0]
		}
		for i := range grid {
			for j, v := range grid[i] {
				var onCount int
				for _, d := range dirs {
					ii := i + d[0]
					jj := j + d[1]
					if !ok(ii, jj) {
						continue
					}
					if grid[ii][jj] == '#' {
						onCount++
					}
				}
				if v == '#' {
					if onCount == 2 || onCount == 3 {
						next[i] = append(next[i], '#')
					} else {
						next[i] = append(next[i], '.')
					}
				} else {
					if onCount == 3 {
						next[i] = append(next[i], '#')
					} else {
						next[i] = append(next[i], '.')
					}
				}
			}
		}
		grid, next = next, grid
	}
	var onCount int
	for i := range grid {
		for _, v := range grid[i] {
			if v == '#' {
				onCount++
			}
		}
	}
	return fmt.Sprint(onCount)
}

func solve2(in *input) string {
	rs := in.xs
	m := len(rs)
	n := len(rs[0].s)
	grid := make([][]byte, m)
	next := make([][]byte, m)
	for i := range grid {
		grid[i] = []byte(rs[i].s)
		next[i] = make([]byte, n)
	}
	grid[0][0] = '#'
	grid[0][n-1] = '#'
	grid[m-1][0] = '#'
	grid[m-1][n-1] = '#'

	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for steps := 0; steps < 100; steps++ {
		for i := range next {
			next[i] = next[i][:0]
		}
		for i := range grid {
			for j, v := range grid[i] {
				var onCount int
				for _, d := range dirs {
					ii := i + d[0]
					jj := j + d[1]
					if !ok(ii, jj) {
						continue
					}
					if grid[ii][jj] == '#' {
						onCount++
					}
				}
				if v == '#' {
					if onCount == 2 || onCount == 3 {
						next[i] = append(next[i], '#')
					} else {
						next[i] = append(next[i], '.')
					}
				} else {
					if onCount == 3 {
						next[i] = append(next[i], '#')
					} else {
						next[i] = append(next[i], '.')
					}
				}
			}
		}
		grid, next = next, grid
		grid[0][0] = '#'
		grid[0][n-1] = '#'
		grid[m-1][0] = '#'
		grid[m-1][n-1] = '#'

	}
	var onCount int
	for i := range grid {
		for _, v := range grid[i] {
			if v == '#' {
				onCount++
			}
		}
	}
	return fmt.Sprint(onCount)
}

type inputItem struct {
	s string
}

type input struct {
	n  int
	xs []inputItem
}

func (p *input) parse(s string) {
	var x inputItem
	x.s = s
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	rows := ax.ReadLines(os.Stdin)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
