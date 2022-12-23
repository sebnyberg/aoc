package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string, rounds int) any {
	lines := ax.MustReadFileLines(inf)
	elves := make(map[[2]int]struct{})
	for i, l := range lines {
		for j, ch := range l {
			if ch == '#' {
				elves[[2]int{i, j}] = struct{}{}
			}
		}
	}
	has := func(m map[[2]int]struct{}, i, j int, dirs [][2]int) bool {
		for _, d := range dirs {
			ii := i + d[0]
			jj := j + d[1]
			if _, exists := m[[2]int{ii, jj}]; exists {
				return true
			}
		}
		return false
	}
	dirs8 := [][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	north := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}}
	south := [][2]int{{1, -1}, {1, 0}, {1, 1}}
	west := [][2]int{{-1, -1}, {0, -1}, {1, -1}}
	east := [][2]int{{-1, 1}, {0, 1}, {1, 1}}
	directions := [][][2]int{north, south, west, east}
	var dirIdx int

	next := make(map[[2]int][][2]int)
	for round := 0; round < rounds; round++ {
		for k := range next {
			delete(next, k)
		}
		l := [][2]int{}
		for e := range elves {
			l = append(l, e)
		}
		sort.Slice(l, func(i, j int) bool {
			return l[i][0] < l[j][0] || (l[i][0] == l[j][0] && l[i][1] < l[j][1])
		})
		for _, e := range l {
			i := e[0]
			j := e[1]
			if !has(elves, i, j, dirs8) {
				next[e] = append(next[e], e)
				continue
			}
			for k := 0; k < 4; k++ {
				ii := (dirIdx + k) % 4
				if !has(elves, i, j, directions[ii]) {
					mvt := directions[ii][1]
					p := [2]int{i + mvt[0], j + mvt[1]}
					next[p] = append(next[p], e)
					goto done
				}
			}
			next[e] = append(next[e], e)
		done:
		}
		dirIdx++
		for k := range elves {
			delete(elves, k)
		}
		// Any elf whose next position is not proposed by any other elf will
		// move to that position
		var moves int
		for e := range next {
			if len(next[e]) == 1 {
				if next[e][0] != e {
					moves++
				}
				elves[e] = struct{}{}
			} else {
				// stay in the same position as before
				for _, e2 := range next[e] {
					elves[e2] = struct{}{}
				}
			}
		}
		if moves == 0 {
			return round + 1
		}
	}

	var res int
	var s bytes.Buffer
	print(elves, &s)
	ss := s.String()
	for _, ch := range ss {
		if ch == '.' {
			res++
		}
	}

	return res
}

func print(elves map[[2]int]struct{}, w io.Writer) {
	minI := math.MaxInt32
	maxI := math.MinInt32
	minJ := math.MaxInt32
	maxJ := math.MinInt32
	for e := range elves {
		minI = ax.Min(minI, e[0])
		maxI = ax.Max(maxI, e[0])
		minJ = ax.Min(minJ, e[1])
		maxJ = ax.Max(maxJ, e[1])
	}
	for i := 0; i < maxI-minI+1; i++ {
		for j := 0; j < maxJ-minJ+1; j++ {
			p := [2]int{
				minI + i,
				minJ + j,
			}
			if _, exists := elves[p]; exists {
				fmt.Fprint(w, "#")
			} else {
				fmt.Fprint(w, ".")
			}
		}
		fmt.Fprint(w, "\n")
	}
	fmt.Fprint(w, "\n")
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f, 10))
	fmt.Printf("Result1:\n%v\n", solve1(f, 50000))
}
