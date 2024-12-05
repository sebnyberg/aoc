package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var j int
	edges := make(map[[2]int]struct{})
	for i, l := range lines {
		if l == "" {
			j = i + 1
			break
		}
		parts := strings.Split(l, "|")
		a := ax.MustParseInt[int](parts[0])
		b := ax.MustParseInt[int](parts[1])
		edges[[2]int{a, b}] = struct{}{}
	}

	var res int
outer:
	for _, l := range lines[j:] {
		var updatePages []int
		for _, x := range strings.Split(l, ",") {
			updatePages = append(updatePages, ax.MustParseInt[int](x))
		}
		for i := 0; i < len(updatePages)-1; i++ {
			for j := i + 1; j < len(updatePages); j++ {
				a := updatePages[i]
				b := updatePages[j]
				if _, exists := edges[[2]int{b, a}]; exists {
					continue outer
				}
			}
		}
		res += updatePages[len(updatePages)/2]
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var j int

	indeg := make(map[int]int)
	adj := make(map[int][]int)
	edges := make(map[[2]int]struct{})
	for i, l := range lines {
		if l == "" {
			j = i + 1
			break
		}
		parts := strings.Split(l, "|")
		a := ax.MustParseInt[int](parts[0])
		b := ax.MustParseInt[int](parts[1])
		indeg[b]++
		adj[a] = append(adj[a], b)
		edges[[2]int{a, b}] = struct{}{}
	}
	ok := func(a []int) bool {
		for i := range a {
			for j := i + 1; j < len(a); j++ {
				if _, ok := edges[[2]int{a[j], a[i]}]; ok {
					return false
				}
			}
		}
		return true
	}

	var sort func(a []int) []int
	sort = func(a []int) []int {
		for i := range a {
			for j := i + 1; j < len(a); j++ {
				if _, ok := edges[[2]int{a[j], a[i]}]; ok {
					a[i], a[j] = a[j], a[i]
					return sort(a)
				}
			}
		}
		return a
	}

	var res int
	for _, l := range lines[j:] {
		var updatePages []int
		for _, x := range strings.Split(l, ",") {
			updatePages = append(updatePages, ax.MustParseInt[int](x))
		}
		if ok(updatePages) {
			continue
		}
		sorted := sort(updatePages)
		res += sorted[len(sorted)/2]
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
