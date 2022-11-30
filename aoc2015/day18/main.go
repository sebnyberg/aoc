package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var absf = ax.Abs[float64]
var absi = ax.Abs[int]
var minf = ax.Min[float64]
var mini = ax.Min[int]
var minu = ax.Min[uint16]
var maxf = ax.Max[float64]
var maxi = ax.Max[int]
var maxu = ax.Max[uint16]
var print = fmt.Print
var printf = fmt.Printf
var println = fmt.Println
var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var sprintln = fmt.Sprintln
var tof = ax.MustParseFloat[float64]
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint16]

func pprint(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmt.Printf(fmtStr, a...)
}
func pprintln(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmtStr += "\n"
	fmt.Printf(fmtStr, a...)
}

var intr = regexp.MustCompile(`[1-9][0-9]*|0`)

func isnum(s string) bool {
	return intr.MatchString(s)
}

func Solve1(rs []parsedRow) string {
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
	return sprint(onCount)
}

func Solve2(rs []parsedRow) string {
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
	return sprint(onCount)
}

type parsedRow struct {
	s string
}

func Parse(s string) parsedRow {
	var r parsedRow
	r.s = s
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var p ax.Problem[parsedRow]
	p.HeadN = 3
	p.TailN = 3
	p.Trunc = true
	for sc.Scan() {
		s := sc.Text()
		p.Input = append(p.Input, s)
		p.Parsed = append(p.Parsed, Parse(s))
	}
	p.Result1 = Solve1(p.Parsed)
	p.Result2 = Solve2(p.Parsed)
	fmt.Fprint(os.Stdout, p)
}
