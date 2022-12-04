package main

import (
	"fmt"
	"math/bits"
)

const input = 1362

func solve1(inf string) string {
	isWall := func(x, y int) bool {
		a := x*x + 3*x + 2*x*y + y + y*y + input
		count := bits.OnesCount(uint(a))
		return count&1 == 1
	}
	seen := make(map[[2]int]bool)
	curr := [][2]int{{1, 1}}
	next := [][2]int{}
	seen[[2]int{1, 1}] = true
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && !isWall(i, j)
	}

	for k := 1; ; k++ {
		next = next[:0]
		dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
		for _, p := range curr {
			for _, d := range dirs {
				xx := p[0] + d[0]
				yy := p[1] + d[1]
				key := [2]int{xx, yy}
				if !ok(xx, yy) || seen[key] {
					continue
				}
				if xx == 31 && yy == 39 {
					return fmt.Sprint(k)
				}
				seen[key] = true
				next = append(next, key)
			}
		}
		curr, next = next, curr

	}
}

func solve2(inf string) string {
	isWall := func(x, y int) bool {
		a := x*x + 3*x + 2*x*y + y + y*y + input
		count := bits.OnesCount(uint(a))
		return count&1 == 1
	}
	seen := make(map[[2]int]bool)
	curr := [][2]int{{1, 1}}
	next := [][2]int{}
	seen[[2]int{1, 1}] = true
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && !isWall(i, j)
	}

	for k := 1; k <= 50; k++ {
		next = next[:0]
		dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
		for _, p := range curr {
			for _, d := range dirs {
				xx := p[0] + d[0]
				yy := p[1] + d[1]
				key := [2]int{xx, yy}
				if !ok(xx, yy) || seen[key] {
					continue
				}
				seen[key] = true
				next = append(next, key)
			}
		}
		curr, next = next, curr
	}
	return fmt.Sprint(len(seen))
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
