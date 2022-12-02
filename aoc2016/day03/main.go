package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	var res int
	for _, l := range ax.MustReadFileLines(inf) {
		parts := strings.Fields(l)
		a := ax.Atoi(parts[0])
		b := ax.Atoi(parts[1])
		c := ax.Atoi(parts[2])
		xs := []int{a, b, c}
		sort.Ints(xs)
		if xs[0]+xs[1] <= xs[2] {
			// invalid
		} else {
			res++
		}
	}
	return fmt.Sprint(res)
}

func solve2(inf string) string {
	var res int
	lines := ax.MustReadFileLines(inf)
	valid := func(as, bs, cs string) bool {
		a := ax.Atoi(as)
		b := ax.Atoi(bs)
		c := ax.Atoi(cs)
		xs := []int{a, b, c}
		sort.Ints(xs)
		if xs[0]+xs[1] <= xs[2] {
			return false
		} else {
			return true
		}
	}
	for i := 0; i < len(lines); i += 3 {
		parts1 := strings.Fields(lines[i])
		parts2 := strings.Fields(lines[i+1])
		parts3 := strings.Fields(lines[i+2])
		if valid(parts1[0], parts2[0], parts3[0]) {
			res++
		}
		if valid(parts1[1], parts2[1], parts3[1]) {
			res++
		}
		if valid(parts1[2], parts2[2], parts3[2]) {
			res++
		}
	}
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
