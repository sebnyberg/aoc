package main

import (
	"fmt"
	"io"

	"github.com/sebnyberg/aoc/ax"
)

type key struct {
	stepIdx int
	i, j    int
}

const (
	tornadoDown  = 1 << 0
	tornadoLeft  = 1 << 1
	tornadoUp    = 1 << 2
	tornadoRight = 1 << 3
)

var dirs = [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

var chars = map[byte]byte{
	tornadoDown:  'v',
	tornadoLeft:  '<',
	tornadoUp:    '^',
	tornadoRight: '>',
}

var bits = map[byte]byte{
	'.': 0,
	'v': tornadoDown,
	'<': tornadoLeft,
	'^': tornadoUp,
	'>': tornadoRight,
}

func parseGrid(fname string) [][]byte {
	lines := ax.MustReadFileLines(fname)
	var grid [][]byte
	for i, l := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}
		grid = append(grid, []byte(l[1:len(l)-1]))
		for j := range grid[i-1] {
			grid[i-1][j] = bits[grid[i-1][j]]
		}
	}
	return grid
}

func findPath(grid [][]byte, start [2]int, finish [2]int) (res int, finalGrid [][]byte) {
	m := len(grid)
	n := len(grid[0])
	curr := [][2]int{start}
	next := [][2]int{}
	seen := make(map[key]bool)
	nextGrid := make([][]byte, m)
	for i := range grid {
		nextGrid[i] = make([]byte, n)
	}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	lcm := m * n / gcd(m, n)

	for steps := 0; len(curr) > 0; steps++ {
		next = next[:0]
		for i := range nextGrid {
			for j := range nextGrid[i] {
				nextGrid[i][j] = 0
			}
		}

		// Move tornadoes
		for i := range grid {
			for j := range grid[i] {
				for k := 0; k < 4; k++ {
					if grid[i][j]&(1<<k) == 0 {
						continue
					}
					d := dirs[k]
					ii := i + d[0]
					jj := j + d[1]
					if ii >= m {
						ii = 0
					}
					if jj >= n {
						jj = 0
					}
					if ii < 0 {
						ii = m - 1
					}
					if jj < 0 {
						jj = n - 1
					}
					nextGrid[ii][jj] |= (1 << k)
				}
			}
		}

		// A next position is valid iff the target position is not occupied by a
		// tornado.
		//
		// A state is a unique combination of tornado positions and player
		// position. A unique tornado state is given by the LCM of the width and
		// height. In other words, the grid state repeats itself when the
		// horizontal and vertical movements happen to coincide.
		//
		// The worst-case storage is therefore O(n*m*n*m).
		//
		stepIdx := steps % lcm
		for _, x := range curr {
			i := x[0]
			j := x[1]
			for _, d := range dirs {
				ii := i + d[0]
				jj := j + d[1]
				p := [2]int{ii, jj}
				if p == finish {
					return steps + 1, nextGrid
				}
				kk := key{stepIdx, ii, jj}
				_, exists := seen[kk]
				if !ok(ii, jj) || nextGrid[ii][jj] > 0 || exists {
					continue
				}
				next = append(next, p)
				seen[kk] = true
			}
			// We may also stand still
			kk := key{stepIdx, i, j}
			_, exists := seen[kk]
			p := [2]int{i, j}
			if !exists && (p == start || ok(i, j) && nextGrid[i][j] == 0) {
				next = append(next, p)
				seen[kk] = true
			}
		}
		curr, next = next, curr
		grid, nextGrid = nextGrid, grid
	}

	return 0, nil
}

func printGrid(w io.Writer, grid [][]byte, playerPos [2]int) {
	if playerPos == [2]int{-1, -1} {
		playerPos = [2]int{-10, -10}
	}
	n := len(grid[0])
	for i := 0; i < n+2; i++ {
		if i == 1 {
			if playerPos[0] == -1 {
				fmt.Fprint(w, "E")
			} else {
				fmt.Fprint(w, ".")
			}
		} else {
			fmt.Fprint(w, "#")
		}
	}
	fmt.Fprintln(w)
	for i := range grid {
		fmt.Fprint(w, "#")
		for j := range grid[i] {
			var count int
			var b byte = '.'
			for k := 0; k < 4; k++ {
				if grid[i][j]&(1<<k) > 0 {
					count++
					b = chars[(1 << k)]
				}
			}
			if count > 1 {
				b = byte('0' + count)
			}
			if playerPos[0] == i && playerPos[1] == j {
				if b != '.' {
					panic("player in invalid pos")
				}
				b = 'E'
			}
			fmt.Fprintf(w, "%c", b)
		}
		fmt.Fprint(w, "#")
		fmt.Fprintln(w)
	}
	for i := 0; i < n+2; i++ {
		if i == n {
			fmt.Fprint(w, ".")
		} else {
			fmt.Fprint(w, "#")
		}
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	f := "input"
	grid := parseGrid(f)
	m := len(grid)
	n := len(grid[0])
	var part2 int
	steps, grid := findPath(grid, [2]int{-1, 0}, [2]int{m, n - 1})
	fmt.Printf("Result1:\n%v\n", steps)
	part2 += steps
	steps, grid = findPath(grid, [2]int{m, n - 1}, [2]int{-1, 0})
	part2 += steps
	steps, grid = findPath(grid, [2]int{-1, 0}, [2]int{m, n - 1})
	part2 += steps
	fmt.Printf("Result2:\n%v\n", part2)
}
