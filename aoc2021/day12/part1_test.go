package day12

import (
	"aoc/ax"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func BenchmarkPart1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestPart1(t *testing.T) {
	assert.Equal(t, "10", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "3510", Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) string {
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

	type key struct{ bm, cur int }
	mem := make(map[key]int, 1000)

	var dfs func(bm, cur int) int
	dfs = func(bm, cur int) int {
		if cur == end {
			return 1
		}
		if v, exists := mem[key{bm, cur}]; exists {
			return v
		}
		var res int
		for _, nei := range adj[cur] {
			if !smallCave[nei] {
				res += dfs(bm, nei)
				continue
			}
			if bm&(1<<nei) == 0 {
				res += dfs(bm|(1<<nei), nei)
			}
		}
		mem[key{bm, cur}] = res
		return res
	}
	return strconv.Itoa(dfs(1<<start, start))
}
