package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var res int
	for _, line := range lines {
		parts := strings.Fields(line)
		var a []int
		for i := range parts {
			a = append(a, ax.MustParseInt[int](parts[i]))
		}
		if safe(a) {
			res++
		}
	}
	return res
}

func safe(a []int) bool {
	dir := 1
	if a[1]-a[0] < 0 {
		dir = -1
	}
	for i := range a {
		if i == 0 {
			continue
		}
		d := dir * (a[i] - a[i-1])
		if d <= 0 || d > 3 {
			return false
		}
	}
	return true
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var res int
outer:
	for _, line := range lines {
		parts := strings.Fields(line)
		var a []int
		for i := range parts {
			a = append(a, ax.MustParseInt[int](parts[i]))
		}
		for i := 0; i < len(a); i++ {
			var bb []int
			for j := 0; j < i; j++ {
				bb = append(bb, a[j])
			}
			for j := i + 1; j < len(a); j++ {
				bb = append(bb, a[j])
			}
			if safe(bb) {
				res++
				continue outer
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
