package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

const (
	air  = 0
	rock = 1
	sand = 2
)

const sz = 700

func solve(inf string, withFloor bool) any {
	var grid [sz][sz]byte
	var pts [][2]int
	var maxY, maxX int
	minX := math.MaxInt32

	draw := func(a, b [2]int) {
		maxX = max(maxX, max(a[0], b[0]))
		maxY = max(maxY, max(a[1], b[1]))
		minX = min(minX, min(a[0], b[0]))
		for y := min(a[1], b[1]); y <= max(a[1], b[1]); y++ {
			for x := min(a[0], b[0]); x <= max(a[0], b[0]); x++ {
				grid[y][x] = rock
			}
		}
	}

	lines := ax.MustReadFileLines(inf)
	for _, l := range lines {
		pts = pts[:0]
		for _, p := range strings.Split(l, "->") {
			fs := strings.Split(p, ",")
			x := ax.Atoi(strings.Trim(fs[0], " "))
			y := ax.Atoi(strings.Trim(fs[1], " "))
			pts = append(pts, [2]int{x, y})
		}
		for i := 1; i < len(pts); i++ {
			draw(pts[i-1], pts[i])
		}
	}

	floor := sz - 1
	if withFloor { // Part 2
		floor = 2 + maxY
	}
	for k := 0; ; k++ {
		x, y := 500, 0
	fallLoop:
		for {
			switch {
			case y+1 == floor:
				break fallLoop // sorry, Dijkstra
			case y+1 > maxY+2:
				return k
			case grid[y+1][x] == air:
				y++
			case grid[y+1][x-1] == air:
				y++
				x--
			case grid[y+1][x+1] == air:
				y++
				x++
			default:
				break fallLoop // sorry, Dijkstra
			}
		}
		if withFloor && x == 500 && y == 0 { // Part 2
			return k + 1
		}
		grid[y][x] |= sand
	}
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, false))
	fmt.Printf("Result2:\n%v\n", solve(f, true))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
