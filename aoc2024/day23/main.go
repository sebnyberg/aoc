package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	// Parse inputs
	lines := ax.MustReadFileLines(inf)
	var adj [521][521]byte
	var names []string
	idx := make(map[string]int)
	addName := func(name string) {
		if _, ok := idx[name]; ok {
			return
		}
		idx[name] = len(names)
		names = append(names, name)
	}
	for _, l := range lines {
		fields := strings.Split(l, "-")
		a := fields[0]
		b := fields[1]
		addName(a)
		addName(b)
		i := idx[a]
		j := idx[b]
		adj[i][j] = 1
		adj[j][i] = 1
	}
	var res int
	for i := 0; i < len(names)-2; i++ {
		a := names[i]
		for j := i + 1; j < len(names)-1; j++ {
			b := names[j]
			for k := j + 1; k < len(names); k++ {
				c := names[k]
				if a[0] != 't' && b[0] != 't' && c[0] != 't' {
					continue
				}
				if adj[i][j]+adj[i][k]+adj[j][k] == 3 {
					res++
				}
			}
		}
	}
	return res
}

func solve2(inf string) any {
	// Parse inputs
	lines := ax.MustReadFileLines(inf)
	var adj [521][521]byte
	var names []string
	idx := make(map[string]int)
	addName := func(name string) {
		if _, ok := idx[name]; ok {
			return
		}
		idx[name] = len(names)
		names = append(names, name)
	}
	for _, l := range lines {
		fields := strings.Split(l, "-")
		a := fields[0]
		b := fields[1]
		addName(a)
		addName(b)
		i := idx[a]
		j := idx[b]
		adj[i][j] = 1
		adj[j][i] = 1
	}
	n := len(names)
	set := make([]int, 0, n)
	res := dfs(&adj, names, set, 0, n)
	resStrs := make([]string, len(res))
	for i, j := range res {
		resStrs[i] = names[j]
	}
	sort.Strings(resStrs)
	return strings.Join(resStrs, ",")
}

func dfs(adj *[521][521]byte, names []string, set []int, i, n int) []int {
	if i == n {
		cpy := make([]int, len(set))
		copy(cpy, set)
		return cpy
	}
	res := dfs(adj, names, set, i+1, n) // don't add anything

	// Try to add i to the set
	var adjSum int
	for _, x := range set {
		adjSum += int(adj[i][x])
	}
	if adjSum == len(set) {
		otherRes := dfs(adj, names, append(set, i), i+1, n)
		if len(otherRes) > len(res) {
			res = otherRes
		}
	}
	return res
}

func main() {
	// fmt.Printf("Result1:\n%v\n", solve1("input"))
	fmt.Printf("Result2:\n%v\n", solve2("input"))
}
