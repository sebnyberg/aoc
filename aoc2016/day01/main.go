package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var input = "L1, L5, R1, R3, L4, L5, R5, R1, L2, L2, L3, R4, L2, R3, R1, L2, R5, R3, L4, R4, L3, R3, R3, L2, R1, L3, R2, L1, R4, L2, R4, L4, R5, L3, R1, R1, L1, L3, L2, R1, R3, R2, L1, R4, L4, R2, L189, L4, R5, R3, L1, R47, R4, R1, R3, L3, L3, L2, R70, L1, R4, R185, R5, L4, L5, R4, L1, L4, R5, L3, R2, R3, L5, L3, R5, L1, R5, L4, R1, R2, L2, L5, L2, R4, L3, R5, R1, L5, L4, L3, R4, L3, L4, L1, L5, L5, R5, L5, L2, L1, L2, L4, L1, L2, R3, R1, R1, L2, L5, R2, L3, L5, L4, L2, L1, L2, R3, L1, L4, R3, R3, L2, R5, L1, L3, L3, L3, L5, R5, R1, R2, L3, L2, R4, R1, R1, R3, R4, R3, L3, R3, L5, R2, L2, R4, R5, L4, L3, L1, L5, L1, R1, R2, L1, R3, R4, R5, R2, R3, L2, L1, L5"

func solve1(inf string) string {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dir int
	cmds := strings.Split(input, ", ")
	var x, y int
	for _, c := range cmds {
		if c[0] == 'L' {
			dir = (dir + 3) % 4
		} else if c[0] == 'R' {
			dir = (dir + 1) % 4
		}
		dist := ax.Atoi(c[1:])
		x += dirs[dir][0] * dist
		y += dirs[dir][1] * dist
	}
	return fmt.Sprint(ax.Abs(x + y))
}

func solve2(inf string) string {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dir int
	cmds := strings.Split(input, ", ")
	seen := make(ax.Set[[2]int])
	var x, y int
	seen.Add([2]int{x, y})
	for _, c := range cmds {
		if c[0] == 'L' {
			dir = (dir + 3) % 4
		} else if c[0] == 'R' {
			dir = (dir + 1) % 4
		}
		dist := ax.Atoi(c[1:])
		for d := 0; d < dist; d++ {
			x += dirs[dir][0]
			y += dirs[dir][1]
			if seen.Has([2]int{x, y}) {
				return fmt.Sprint(ax.Abs(x + y))
			}
			seen.Add([2]int{x, y})
		}
	}
	return fmt.Sprint(ax.Abs(x + y))
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
