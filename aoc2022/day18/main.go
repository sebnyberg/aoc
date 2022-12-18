package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	cubes := make(map[[3]int]struct{})
	for _, l := range lines {
		fs := strings.Split(l, ",")
		x := ax.Atoi(fs[0])
		y := ax.Atoi(fs[1])
		z := ax.Atoi(fs[2])
		cubes[[3]int{x, y, z}] = struct{}{}
	}
	dirs := [][3]int{
		{-1, 0, 0}, {1, 0, 0},
		{0, -1, 0}, {0, 1, 0},
		{0, 0, -1}, {0, 0, 1},
	}
	var res int
	for c := range cubes {
		var sides int
		for _, d := range dirs {
			c2 := [3]int{
				c[0] + d[0],
				c[1] + d[1],
				c[2] + d[2],
			}
			if _, exists := cubes[c2]; exists {
				continue
			}
			sides++
		}
		res += sides
	}
	return res
}

var dirs3d = [][3]int{
	{-1, 0, 0}, {1, 0, 0},
	{0, -1, 0}, {0, 1, 0},
	{0, 0, -1}, {0, 0, 1},
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	inside := make(map[[3]int]struct{})
	for _, l := range lines {
		fs := strings.Split(l, ",")
		x := ax.Atoi(fs[0])
		y := ax.Atoi(fs[1])
		z := ax.Atoi(fs[2])
		inside[[3]int{x, y, z}] = struct{}{}
	}

	vis := map[[3]int]struct{}{}
	outside := map[[3]int]struct{}{}
	curr := [][3]int{}
	next := [][3]int{}
	var i int
	// Do one loop to fill spaces
	for c := range inside {
		for _, d := range dirs3d {
			c2 := [3]int{
				c[0] + d[0],
				c[1] + d[1],
				c[2] + d[2],
			}
			if _, exists := inside[c2]; exists {
				continue
			}
			if _, exists := outside[c2]; exists {
				continue
			}
			fill(inside, outside, vis, &curr, &next, c2)
		}
		i++
	}
	// Then one loop to count sides
	var res int
	for c := range inside {
		var sides int
		for _, d := range dirs3d {
			c2 := [3]int{
				c[0] + d[0],
				c[1] + d[1],
				c[2] + d[2],
			}
			if _, exists := inside[c2]; exists {
				continue
			}
			sides++
		}
		res += sides
	}
	return res
}

func fill(inside, outside, vis map[[3]int]struct{}, curr, next *[][3]int, start [3]int) bool {
	for k := range vis {
		delete(vis, k)
	}
	*curr = (*curr)[:0]
	*curr = append(*curr, start)
	vis[start] = struct{}{}
	for count := 1; len(*curr) > 0; count++ {
		if count > 100 {
			// Infinite space! put into "outside"
			for c := range vis {
				outside[c] = struct{}{}
			}
			return false
		}
		*next = (*next)[:0]
		for _, c := range *curr {
			for _, d := range dirs3d {
				c2 := [3]int{
					c[0] + d[0],
					c[1] + d[1],
					c[2] + d[2],
				}
				if _, exists := outside[c2]; exists {
					return false
				}
				if _, exists := inside[c2]; exists {
					continue
				}
				if _, exists := vis[c2]; exists {
					continue
				}
				vis[c2] = struct{}{}
				*next = append(*next, c2)
			}
		}
		curr, next = next, curr
	}
	// Space was filled! Put cubes into list of cubes
	for c := range vis {
		inside[c] = struct{}{}
	}
	return true
}

func main() {
	f := "input"
	// fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
