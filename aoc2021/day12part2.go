package aoc2021

import (
	"aoc/ax"
	"strconv"
	"strings"
	"unicode"
)

func Day12Part2(rows []string) string {
	strToIdx := make(map[string]int, len(rows))
	strs := make([]string, 0, len(rows))
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
		strs = append(strs, a)
		pos++
	}

	for _, row := range rows {
		parts := strings.Split(row, "-")
		a, b := parts[0], parts[1]
		maybeInit(a)
		maybeInit(b)
		if b != "start" {
			adj[strToIdx[a]] = append(adj[strToIdx[a]], strToIdx[b])
		}
		if a != "start" {
			adj[strToIdx[b]] = append(adj[strToIdx[b]], strToIdx[a])
		}
	}
	start := strToIdx["start"]
	end := strToIdx["end"]

	var dfs func(bm, cur int, double bool) int
	dfs = func(bm, cur int, double bool) int {
		if cur == end {
			return 1
		}
		var res int
		for _, nei := range adj[cur] {
			if !smallCave[nei] {
				res += dfs(bm, nei, double)
				continue
			}
			if bm&(1<<nei) == 0 {
				res += dfs(bm|(1<<nei), nei, double)
			} else if !double {
				res += dfs(bm|(1<<nei), nei, true)
			}
		}
		return res
	}
	return strconv.Itoa(dfs(1<<start, start, false))
}
