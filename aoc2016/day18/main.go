package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

const safe = '.'
const trap = '^'

func solve1(inf string) string {
	l := ax.MustReadFileLines(inf)[0]
	// l = ".^^.^.^^^^"
	n := len(l)
	curr := []byte(l)
	next := make([]byte, n)
	var count int
	for i := range l {
		if l[i] == safe {
			count++
		}
	}
	for i := 0; i < 39; i++ {
		next = next[:0]
		for i := range curr {
			var left bool
			if i > 0 && curr[i-1] == trap {
				left = true
			}
			var right bool
			if i < n-1 && curr[i+1] == trap {
				right = true
			}
			center := curr[i] == trap
			if left && center && !right ||
				(center && right && !left) ||
				(left && !right && !center) ||
				(right && !left && !center) {
				next = append(next, trap)
			} else {
				count++
				next = append(next, safe)
			}
		}
		curr, next = next, curr
	}
	return fmt.Sprint(count)
}

func solve2(inf string) string {
	l := ax.MustReadFileLines(inf)[0]
	n := len(l)
	curr := []byte(l)
	next := make([]byte, n)
	var count int
	for i := range l {
		if l[i] == safe {
			count++
		}
	}
	for i := 0; i < (400000 - 1); i++ {
		next = next[:0]
		var delta int
		for i := range curr {
			var left bool
			if i > 0 && curr[i-1] == trap {
				left = true
			}
			var right bool
			if i < n-1 && curr[i+1] == trap {
				right = true
			}
			center := curr[i] == trap
			if left && center && !right ||
				(center && right && !left) ||
				(left && !right && !center) ||
				(right && !left && !center) {
				next = append(next, trap)
			} else {
				delta++
				next = append(next, safe)
			}
		}
		count += delta
		curr, next = next, curr
	}
	return fmt.Sprint(count)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
