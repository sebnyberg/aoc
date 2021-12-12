package aoc2021

import (
	"aoc/ax"
	"strconv"
	"strings"
	"unicode"
)

func Day12Part1(rows []string) string {
	strToIdx := make(map[string]int, len(rows))
	adj := make([][]int, 0, len(rows))
	var pos int
	smallCave := make([]bool, 0, len(rows))

	maybeInit := func(a string) {
		if _, exists := strToIdx[a]; exists {
			return
		}
		strToIdx[a] = pos
		if ax.Is(a, unicode.IsLower) {
			smallCave = append(smallCave, true)
		} else {
			smallCave = append(smallCave, false)
		}
		adj = append(adj, make([]int, 0, len(rows)))
		pos++
	}

	for _, row := range rows {
		parts := strings.Split(row, "-")
		a, b := parts[0], parts[1]
		maybeInit(a)
		maybeInit(b)
		adj[strToIdx[a]] = append(adj[strToIdx[a]], strToIdx[b])
		adj[strToIdx[b]] = append(adj[strToIdx[b]], strToIdx[a])
	}
	start := strToIdx["start"]
	end := strToIdx["end"]
	var res int

	var dfs func(bm, cur int)
	dfs = func(bm, cur int) {
		if cur == end {
			res++
			return
		}
		for _, nei := range adj[cur] {
			if !smallCave[nei] {
				dfs(bm, nei)
				continue
			}
			if bm&(1<<nei) == 0 {
				dfs(bm|(1<<nei), nei)
			}
		}
	}
	dfs(1<<start, start)
	return strconv.Itoa(res)
}
