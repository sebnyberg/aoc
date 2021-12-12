package day12

import (
	"aoc/ax"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay12Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay12Part2(t *testing.T) {
	assert.Equal(t, "36", Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, "122880", Part2(ax.MustReadFineLines("input")))
}

func Part2(rows []string) string {
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

	// Todo: change to fixed-size array instead of map
	type key struct {
		double bool
		bm     int
		cur    int
	}
	mem := make(map[key]int, 1000)

	var dfs func(bm, cur int, double bool) int
	dfs = func(bm, cur int, double bool) int {
		if cur == end {
			return 1
		}
		if v, exists := mem[key{double, bm, cur}]; exists {
			return v
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
		mem[key{double, bm, cur}] = res
		return res
	}
	return strconv.Itoa(dfs(1<<start, start, false))
}
