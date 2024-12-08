package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	pos := map[byte][][2]int{}
	m := len(lines)
	n := len(lines[0])

	// capture positions
	for i, l := range lines {
		for j, ch := range l {
			if ch == '.' {
				continue
			}
			pos[byte(ch)] = append(pos[byte(ch)], [2]int{i, j})
		}
	}
	ok := func(pos [2]int) bool {
		return pos[0] >= 0 && pos[0] < m && pos[1] >= 0 && pos[1] < n
	}

	// For each pair of points, calculate antinodes
	seen := make(map[[2]int]bool)
	for _, points := range pos {
		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				a := points[i]
				b := points[j]
				di := b[0] - a[0]
				dj := b[1] - a[1]
				candA := [2]int{b[0] + di, b[1] + dj}
				if ok(candA) {
					seen[candA] = true
				}
				candB := [2]int{a[0] - di, a[1] - dj}
				if ok(candB) {
					seen[candB] = true
				}
			}
		}
	}

	return len(seen)
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	pos := map[byte][][2]int{}
	m := len(lines)
	n := len(lines[0])

	// capture positions
	for i, l := range lines {
		for j, ch := range l {
			if ch == '.' {
				continue
			}
			pos[byte(ch)] = append(pos[byte(ch)], [2]int{i, j})
		}
	}
	ok := func(pos [2]int) bool {
		return pos[0] >= 0 && pos[0] < m && pos[1] >= 0 && pos[1] < n
	}

	// For each pair of points, calculate antinodes
	seen := make(map[[2]int]bool)
	for _, points := range pos {
		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				a := points[i]
				b := points[j]
				di := b[0] - a[0]
				dj := b[1] - a[1]
				for pp := a; ok(pp); {
					if pp != b || pp != a {
						seen[pp] = true
					}
					pp = [2]int{pp[0] + di, pp[1] + dj}
				}
				for pp := a; ok(pp); {
					if pp != b || pp != a {
						seen[pp] = true
					}
					pp = [2]int{pp[0] - di, pp[1] - dj}
				}
			}
		}
	}

	return len(seen)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
