package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	lines := ax.MustReadFileLines(inf)
	nodes := make(map[[2]int]node)
	for _, l := range lines[2:] {
		fs := strings.Fields(l)
		name := strings.Split(fs[0], "/")[3]
		xa := strings.Split(name, "-")[1]
		x := ax.Atoi(xa[1:])
		ya := strings.Split(name, "-")[2]
		y := ax.Atoi(ya[1:])
		var n node
		n.x = x
		n.y = y
		n.size = ax.Atoi(fs[1][:len(fs[1])-1])
		n.used = ax.Atoi(fs[2][:len(fs[2])-1])
		n.avail = ax.Atoi(fs[3][:len(fs[3])-1])
		n.perc = ax.Atoi(fs[4][:len(fs[4])-1])
		nodes[[2]int{x, y}] = n
	}
	// seen := make(map[[2]int]bool)
	var res int
	for _, a := range nodes {
		for _, b := range nodes {
			if a.x == b.x && a.y == b.y {
				continue
			}
			if a.used == 0 {
				continue
			}
			if b.avail >= a.used {
				res++
			}
		}
		// dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
		// for _, d := range dirs {
		// 	xx := a.x + d[0]
		// 	yy := a.y + d[1]
		// 	kk := [2]int{xx, yy}
		// 	if v, exists := nodes[kk]; !exists {
		// 		continue
		// 	}
		// }
	}
	return fmt.Sprint(res)
}

type node struct {
	x, y  int
	size  int
	used  int
	avail int
	perc  int
}

func solve2(inf string) string {
	lines := ax.MustReadFileLines(inf)
	m := 38
	n := 24
	// m = 3
	// n = 3
	nodes := make([][]node, m)
	invalid := make([][]bool, m)
	for i := range nodes {
		nodes[i] = make([]node, n)
		invalid[i] = make([]bool, n)
	}
	var maxx int
	minSize := math.MaxInt32
	var maxUsed int
	var startX, startY int
	for _, l := range lines[2:] {
		fs := strings.Fields(l)
		name := strings.Split(fs[0], "/")[3]
		xa := strings.Split(name, "-")[1]
		x := ax.Atoi(xa[1:])
		ya := strings.Split(name, "-")[2]
		y := ax.Atoi(ya[1:])
		size := ax.Atoi(fs[1][:len(fs[1])-1])
		if size > 100 {
			invalid[x][y] = true
		}
		if size <= 100 {
			minSize = ax.Min(minSize, size)
		}
		used := ax.Atoi(fs[2][:len(fs[2])-1])
		if used <= 100 {
			maxUsed = ax.Max(maxUsed, used)
		}
		if used == 0 {
			startX = x
			startY = y
		}
		maxx = ax.Max(maxx, x)
	}
	// Based on looking at the data, it appears that nodes are either valid or
	// not in terms of moving data, and there is only one position that is empty.

	// A winning state is found when:
	//
	// 1. The cursor (empty node) is in the goal position (0,0)
	// 2. The wanted data is next to (0,0)
	//
	// The current state of the board can be summarized as:
	//
	// 1. current empty position
	// 2. current wanted data position
	//
	// We can use BFS to find the optimal number of steps
	type state struct {
		emptyX  int
		emptyY  int
		wantedX int
		wantedY int
	}
	seen := make(map[state]bool)
	initial := state{
		emptyX:  startX,
		emptyY:  startY,
		wantedX: maxx,
		wantedY: 0,
	}
	seen[initial] = true
	curr := []state{initial}
	next := []state{}
	ok := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < m && y < n
	}
	win := func(s *state) bool {
		return s.emptyX == 0 && s.emptyY == 0 &&
			(s.wantedX+s.wantedY == 1)
	}

	for steps := 1; ; steps++ {
		next = next[:0]
		for _, x := range curr {
			dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
			for _, d := range dirs {
				nextState := x
				xx := x.emptyX + d[0]
				yy := x.emptyY + d[1]
				if !ok(xx, yy) || invalid[xx][yy] {
					continue
				}
				if x.wantedX == xx && x.wantedY == yy {
					nextState.wantedX = x.emptyX
					nextState.wantedY = x.emptyY
				}
				nextState.emptyX = xx
				nextState.emptyY = yy
				if seen[nextState] {
					continue
				}
				if win(&nextState) {
					return fmt.Sprint(steps + 1)
				}
				seen[nextState] = true
				next = append(next, nextState)
			}
		}
		curr, next = next, curr
	}
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
