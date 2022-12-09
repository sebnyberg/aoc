package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func p(snake [][2]int) {
	var maxX, maxY int
	minX, minY := math.MaxInt32, math.MaxInt32
	m := make(map[[2]int]int)
	for i, a := range snake {
		m[a] = i
		maxX = ax.Max(maxX, a[1])
		maxY = ax.Max(maxY, a[0])
		minX = ax.Min(minX, a[1])
		minY = ax.Min(minY, a[0])
	}
	for i := minY - 1; i <= maxY+1; i++ {
		for j := minX - 1; j <= maxX+1; j++ {
			if k, exists := m[[2]int{i, j}]; exists {
				if k == 0 {
					fmt.Print("H")
				} else {
					fmt.Print(k)
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func solve1(inf string, ropelen int) any {
	lines := ax.MustReadFileLines(inf)
	rope := make([][2]int, ropelen)
	for i := 0; i < ropelen; i++ {
		rope[i] = [2]int{0, 0}
	}
	seen := make(map[[2]int]bool)
	deltas := map[string][2]int{
		"R": {0, 1},
		"L": {0, -1},
		"U": {1, 0},
		"D": {-1, 0},
	}
	diags := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, l := range lines {
		fs := strings.Fields(l)
		d := deltas[fs[0]]
		steps := ax.Atoi(fs[1])
		for k := 0; k < steps; k++ {
			rope[0][0] += d[0]
			rope[0][1] += d[1]
			for t := 1; t < ropelen; t++ {
				dj := rope[t-1][1] - rope[t][1]
				di := rope[t-1][0] - rope[t][0]
				if dj == 0 {
					if di >= 2 {
						rope[t][0]++
					} else if di <= -2 {
						rope[t][0]--
					}
					continue
				}
				if di == 0 {
					if dj >= 2 {
						rope[t][1]++
					} else if dj <= -2 {
						rope[t][1]--
					}
					continue
				}
				if ax.Abs(di) <= 1 && ax.Abs(dj) <= 1 {
					continue
				}
				// Find a diagonal that puts the tail in the right spot
				for _, diag := range diags {
					i := rope[t][0] + diag[0]
					j := rope[t][1] + diag[1]
					di := rope[t-1][0] - i
					dj := rope[t-1][1] - j
					if ax.Abs(di) <= 1 && ax.Abs(dj) <= 1 {
						rope[t] = [2]int{i, j}
						break
					}
				}
			}
			// p(rope)
			seen[rope[ropelen-1]] = true
		}
	}
	return len(seen)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f, 2))
	fmt.Printf("Result2:\n%v\n", solve1(f, 10))
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)

	dirs := [][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	around := func(a, b [2]int) bool {
		if a == b {
			return true
		}
		for _, d := range dirs {
			if [2]int{a[0] + d[0], a[1] + d[1]} == b {
				return true
			}
		}
		return false
	}
	var snake [][2]int
	for i := 0; i < 10; i++ {
		snake = append(snake, [2]int{0, 0})
	}
	m := 10
	seen := make(map[[2]int]bool)
	seen[[2]int{0, 0}] = true
	diags := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, l := range lines {
		fs := strings.Fields(l)
		// If the head and tail for any pair is out of range then there are two
		// cases:
		// 1. They share column / row: move one step closer
		// 2. They do no share col / row: move prior diagonally
		switch fs[0] {
		case "R":
			for i := 0; i < ax.Atoi(fs[1]); i++ {
				snake[0][1]++
				for i := 1; i < m; i++ {
					h, t := snake[i-1], snake[i]
					dy := h[0] - t[0]
					dx := h[1] - t[1]
					if dx == 0 && dy >= 2 {
						t[0]++
						snake[i] = t
						continue
					}
					if dx == 0 && dy <= -2 {
						t[0]--
						snake[i] = t
						continue
					}
					if dy == 0 && dx >= 2 {
						t[1]++
						snake[i] = t
						continue
					}
					if dy == 0 && dx <= -2 {
						t[1]--
						snake[i] = t
						continue
					}
					if around(t, h) {
						continue
					}
					// Otherwise, we need to move diagonally
					// Try all diagonal moves and see which one works
					for _, d := range diags {
						c := [2]int{t[0] + d[0], t[1] + d[1]}
						if around(h, c) {
							snake[i] = c
							break
						}
					}
				}
				p(snake)
				seen[snake[9]] = true
			}
		case "L":
			for i := 0; i < ax.Atoi(fs[1]); i++ {
				snake[0][1]--
				for i := 1; i < m; i++ {
					h, t := snake[i-1], snake[i]
					dy := h[0] - t[0]
					dx := h[1] - t[1]
					if dx == 0 && dy >= 2 {
						t[0]++
						snake[i] = t
						continue
					}
					if dx == 0 && dy <= -2 {
						t[0]--
						snake[i] = t
						continue
					}
					if dy == 0 && dx >= 2 {
						t[1]++
						snake[i] = t
						continue
					}
					if dy == 0 && dx <= -2 {
						t[1]--
						snake[i] = t
						continue
					}
					if around(t, h) {
						continue
					}
					// If sharing col/row, just move in that direction
					if h[0] == t[0] {
						snake[i][1]--
						continue
					}
					// Otherwise, we need to move diagonally
					// Try all diagonal moves and see which one works
					for _, d := range diags {
						c := [2]int{t[0] + d[0], t[1] + d[1]}
						if around(h, c) {
							snake[i] = c
							break
						}
					}
				}
				p(snake)
				seen[snake[9]] = true
			}
		case "U":
			for i := 0; i < ax.Atoi(fs[1]); i++ {
				snake[0][0]--
				for i := 1; i < m; i++ {
					h, t := snake[i-1], snake[i]
					dy := h[0] - t[0]
					dx := h[1] - t[1]
					if dx == 0 && dy >= 2 {
						t[0]++
						snake[i] = t
						continue
					}
					if dx == 0 && dy <= -2 {
						t[0]--
						snake[i] = t
						continue
					}
					if dy == 0 && dx >= 2 {
						t[1]++
						snake[i] = t
						continue
					}
					if dy == 0 && dx <= -2 {
						t[1]--
						snake[i] = t
						continue
					}
					if around(t, h) {
						continue
					}
					// If sharing col/row, just move in that direction
					if h[1] == t[1] {
						snake[i][0]--
						continue
					}
					// Otherwise, we need to move diagonally
					// Try all diagonal moves and see which one works
					for _, d := range diags {
						c := [2]int{t[0] + d[0], t[1] + d[1]}
						if around(h, c) {
							snake[i] = c
							break
						}
					}
				}
				p(snake)
				seen[snake[9]] = true
			}
		case "D":
			for i := 0; i < ax.Atoi(fs[1]); i++ {
				snake[0][0]++
				for i := 1; i < m; i++ {
					h, t := snake[i-1], snake[i]
					dy := h[0] - t[0]
					dx := h[1] - t[1]
					if dx == 0 && dy >= 2 {
						t[0]++
						snake[i] = t
						continue
					}
					if dx == 0 && dy <= -2 {
						t[0]--
						snake[i] = t
						continue
					}
					if dy == 0 && dx >= 2 {
						t[1]++
						snake[i] = t
						continue
					}
					if dy == 0 && dx <= -2 {
						t[1]--
						snake[i] = t
						continue
					}
					if around(t, h) {
						continue
					}
					// If sharing col/row, just move in that direction
					if h[1] == t[1] {
						snake[i][0]++
						continue
					}
					// Otherwise, we need to move diagonally
					// Try all diagonal moves and see which one works
					for _, d := range diags {
						c := [2]int{t[0] + d[0], t[1] + d[1]}
						if around(h, c) {
							snake[i] = c
							break
						}
					}
				}
				p(snake)
				seen[snake[9]] = true
			}
		}
		// fmt.Println(h, t)
	}
	return len(seen)
}
