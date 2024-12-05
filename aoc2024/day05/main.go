package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)

	edges := make(map[[2]int]struct{})
	for lines[0] != "" {
		parts := strings.Split(lines[0], "|")
		a := ax.MustParseInt[int](parts[0])
		b := ax.MustParseInt[int](parts[1])
		edges[[2]int{a, b}] = struct{}{}
		lines = lines[1:] // pop current line
	}
	lines = lines[1:] // pop current line

	var res int
outer:
	for _, l := range lines {
		var pages []int
		for _, x := range strings.Split(l, ",") {
			pages = append(pages, ax.MustParseInt[int](x))
		}
		for i := 0; i < len(pages)-1; i++ {
			for j := i + 1; j < len(pages); j++ {
				a := pages[i]
				b := pages[j]
				if _, exists := edges[[2]int{b, a}]; exists {
					continue outer
				}
			}
		}
		res += pages[len(pages)/2]
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)

	// Capture edges
	edges := make(map[[2]int]struct{})
	for lines[0] != "" {
		parts := strings.Split(lines[0], "|")
		a := ax.MustParseInt[int](parts[0])
		b := ax.MustParseInt[int](parts[1])
		edges[[2]int{a, b}] = struct{}{}
		lines = lines[1:] // pop current line
	}
	lines = lines[1:] // pop current line

	var sort func(a []int) (didSort bool)
	sort = func(a []int) (didSort bool) {
		for i := range a {
			for j := i + 1; j < len(a); j++ {
				if _, ok := edges[[2]int{a[j], a[i]}]; ok {
					a[i], a[j] = a[j], a[i]
					sort(a)
					return true
				}
			}
		}
		return false
	}

	// Try-sort pages and add to result if sorterd
	var result int
	for _, l := range lines {
		var pages []int
		for _, x := range strings.Split(l, ",") {
			pages = append(pages, ax.MustParseInt[int](x))
		}
		if !sort(pages) {
			continue
		}
		result += pages[len(pages)/2]
	}
	return result
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
