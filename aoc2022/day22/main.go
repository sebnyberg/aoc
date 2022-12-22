package main

import (
	"bytes"
	"fmt"
	"math"
	"unicode"

	"github.com/sebnyberg/aoc/ax"
)

var dirs = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

const (
	dirRight = 0
	dirDown  = 1
	dirLeft  = 2
	dirUp    = 3
)

func solve1(grid [][]byte, moves string) any {
	// Find leftmost open tile on first row
	var i, j int
	for grid[0][j] != '.' {
		j++
	}

	dir := 0

	var mi int
	for mi < len(moves) {
		switch moves[mi] {
		case 'L':
			dir = (dir + 3) % 4
			mi++
		case 'R':
			dir = (dir + 1) % 4
			mi++
		default:
			mj := mi
			for mj < len(moves) && unicode.IsDigit(rune(moves[mj])) {
				mj++
			}
			x := ax.Atoi(moves[mi:mj])
			if x == 0 {
				panic(moves[mi:])
			}
			mi = mj
			for kk := 0; kk < x; kk++ {
				ii, jj, nextDir := next2DPos(grid, i, j, dir)
				if grid[ii][jj] == '#' {
					break
				}
				i = ii
				j = jj
				dir = nextDir
			}
		}
	}
	res := 1000*(i+1) + 4*(j+1) + dir
	return res
}

func next2DPos(grid [][]byte, i, j int, dir int) (int, int, int) {
	d := dirs[dir]
	i += d[0]
	j += d[1]
	m := len(grid)
	n := len(grid[0])
	for i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == ' ' {
		// i and j are out of bounds, either in whitespace or outside the board
		// keep moving in the given direction until reaching the grid again
		switch {
		case i < 0:
			i = m - 1
		case i >= m:
			i = 0
		case j < 0:
			j = n - 1
		case j >= n:
			j = 0
		default:
			i += d[0]
			j += d[1]
		}
	}
	return i, j, dir
}

func solve2(
	sides [6][][]byte,
	moves string,
	sideTopLefts [6][2]int,
	sideDirs, nextSides [6][4]int,
) any {
	dirChars := []byte{'>', 'v', '<', '^'}

	printGrid := func(s, dir, ii, jj int) {
		for range sides[s][0] {
			fmt.Printf("%d", s)
		}
		fmt.Println()
		for i := range sides[s] {
			for j := range sides[s][i] {
				if i == ii && j == jj {
					fmt.Printf("%c", dirChars[dir])
				} else {
					fmt.Printf("%c", sides[s][i][j])
				}
			}
			fmt.Println()
		}
		for range sides[s][0] {
			fmt.Printf("%d", s)
		}
		fmt.Println()
		fmt.Println()
	}
	_ = printGrid

	// Find leftmost open tile on first row
	var i, j int
	var side int

	dir := 0
	var mi int
	for mi < len(moves) {
		switch moves[mi] {
		case 'L':
			dir = (dir + 3) % 4
			mi++
		case 'R':
			dir = (dir + 1) % 4
			mi++
		default:
			mj := mi
			for mj < len(moves) && unicode.IsDigit(rune(moves[mj])) {
				mj++
			}
			x := ax.Atoi(moves[mi:mj])
			if x == 0 {
				panic(moves[mi:])
			}
			mi = mj
			for kk := 0; kk < x; kk++ {
				ii, jj, nextDir, nextSide := nextCubePos(
					sides, sideDirs, nextSides, side, i, j, dir)

				if sides[nextSide][ii][jj] == '#' {
					break
				}
				i = ii
				j = jj
				dir = nextDir
				side = nextSide
			}
		}
	}
	actualI := sideTopLefts[side][0] + i + 1
	actualJ := sideTopLefts[side][1] + j + 1
	res := 1000*actualI + 4*actualJ + dir
	return res
}

func nextCubePos(
	sides [6][][]byte,
	sideDirs, nextSides [6][4]int,
	side, i, j, dir int,
) (jj int, ii int, nextDir int, nextSide int) {
	d := dirs[dir]
	ii = i + d[0]
	jj = j + d[1]
	n := len(sides[side][0])
	if ii >= 0 && jj >= 0 && ii < n && jj < n {
		// Still inside the same side
		return ii, jj, dir, side
	}
	if ii == i && jj == j {
		panic("wtf")
	}

	// We just moved to a new side
	nextSide = nextSides[side][dir]
	nextDir = sideDirs[side][dir]

	// Depending on how we entered this side, the location will be different.
	switch dir {
	case dirRight:
		switch nextDir {
		case dirRight:
			return ii, 0, nextDir, nextSide
		case dirLeft:
			return n - 1 - ii, n - 1, nextDir, nextSide
		case dirDown:
			return 0, n - 1 - ii, nextDir, nextSide
		case dirUp:
			return n - 1, ii, nextDir, nextSide
		}
	case dirDown:
		switch nextDir {
		case dirRight:
			return n - 1 - jj, 0, nextDir, nextSide
		case dirLeft:
			return jj, n - 1, nextDir, nextSide
		case dirDown:
			return 0, jj, nextDir, nextSide
		case dirUp:
			return n - 1, n - 1 - jj, nextDir, nextSide
		}
	case dirLeft:
		switch nextDir {
		case dirRight:
			return n - 1 - ii, 0, nextDir, nextSide
		case dirLeft:
			return ii, n - 1, nextDir, nextSide
		case dirDown:
			return 0, ii, nextDir, nextSide
		case dirUp:
			return n - 1, n - 1 - ii, nextDir, nextSide
		}
	case dirUp:
		switch nextDir {
		case dirRight:
			return jj, 0, nextDir, nextSide
		case dirLeft:
			return n - 1 - jj, n - 1, nextDir, nextSide
		case dirDown:
			return 0, n - 1 - jj, nextDir, nextSide
		case dirUp:
			return n - 1, jj, nextDir, nextSide
		}
	}

	panic("should never happen")
}

func parse(fname string) (grid [][]byte, sides [6][][]byte, sideTopLefts [6][2]int, moves string) {
	lines := ax.MustReadFileLines(fname)
	var maxWidth int
	for i, l := range lines {
		if l == "" {
			moves = lines[i+1]
			break
		}
		grid = append(grid, []byte(l))
		maxWidth = ax.Max(maxWidth, len(l))
	}
	// Adjust grid to have same width everywhere
	padding := bytes.Repeat([]byte{' '}, maxWidth)
	// sides = make([][]byte, len(grid))
	for i := range grid {
		// sides[i] = make([]byte, maxWidth)
		m := maxWidth - len(grid[i])
		grid[i] = append(grid[i], padding[:m]...)
	}

	// Parse side length
	var count int
	for i := range grid {
		for _, v := range grid[i] {
			if v != ' ' {
				count++
			}
		}
	}
	width := int(math.Sqrt(float64(count / 6)))

	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	// Parse sides
	var side int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == ' ' || seen[i][j] {
				continue
			}
			// This position has not been covered (yet)
			// Parse the next 50x50 as this side number
			sideTopLefts[side] = [2]int{i, j}
			sides[side] = make([][]byte, width)
			for di := 0; di < width; di++ {
				sides[side][di] = make([]byte, width)
				for dj := 0; dj < width; dj++ {
					seen[i+di][j+dj] = true
					sides[side][di][dj] = grid[i+di][j+dj]
				}
			}
			side++
		}
	}
	return
}

func main() {
	f := "input"
	grid, sides, sideTopLefts, moves := parse(f)
	fmt.Printf("Result1:\n%v\n", solve1(grid, moves))

	// This was painstakingly, manually constructed
	// When going out of bounds, this is the new direction to head.
	// We also need to adjust the actual position
	nextSides := [6][4]int{
		0: {1, 2, 3, 5},
		1: {4, 2, 0, 5},
		2: {1, 4, 3, 0},
		3: {4, 5, 0, 2},
		4: {1, 5, 3, 2},
		5: {4, 1, 0, 3},
	}
	sideDirs := [6][4]int{
		// {right, down, left, up}
		0: {dirRight, dirDown, dirRight, dirRight},
		1: {dirLeft, dirLeft, dirLeft, dirUp},
		2: {dirUp, dirDown, dirDown, dirUp},
		3: {dirRight, dirDown, dirRight, dirRight},
		4: {dirLeft, dirLeft, dirLeft, dirUp},
		5: {dirUp, dirDown, dirDown, dirUp},
	}

	fmt.Printf("Result2:\n%v\n", solve2(sides, moves, sideTopLefts, sideDirs, nextSides))
}
