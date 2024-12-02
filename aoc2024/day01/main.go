package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var a, b []int
	for _, line := range lines {
		parts := strings.Fields(line)
		aa := ax.MustParseInt[int](parts[0])
		bb := ax.MustParseInt[int](parts[1])
		a = append(a, aa)
		b = append(b, bb)
	}
	sort.Ints(a)
	sort.Ints(b)
	var res int
	for i := range a {
		res += abs(a[i] - b[i])
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	b := map[int]int{}
	var a []int
	for _, line := range lines {
		parts := strings.Fields(line)
		aa := ax.MustParseInt[int](parts[0])
		a = append(a, aa)
		bb := ax.MustParseInt[int](parts[1])
		b[bb]++
	}
	var res int
	for i := range a {
		res += a[i] * b[a[i]]
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
