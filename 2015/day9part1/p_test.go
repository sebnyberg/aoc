package p_test

import (
	"aoc/ax"
	"math"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 117, res)
}

type edge struct {
	target int
	dist   int
}

var linePat = regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)

func run(lines chan string) int {
	// This is a brute-force graph traversal exercise. It's not too bad - the
	// input dataset is small. Simply do DFS to exhaust all alternatives.
	names := make([]string, 0)
	nameToIdx := make(map[string]int)
	adj := make([][]edge, 0)
	// Parse graph
	for line := range lines {
		parts := linePat.FindStringSubmatch(line)
		from := parts[1]
		to := parts[2]
		dist := ax.MustParseInt(parts[3], 10)
		if _, exists := nameToIdx[from]; !exists {
			nameToIdx[from] = len(names)
			names = append(names, from)
			adj = append(adj, []edge{})
		}
		if _, exists := nameToIdx[to]; !exists {
			nameToIdx[to] = len(names)
			names = append(names, to)
			adj = append(adj, []edge{})
		}
		a, b := nameToIdx[from], nameToIdx[to]
		adj[a] = append(adj[a], edge{b, dist})
		adj[b] = append(adj[b], edge{a, dist})
	}

	// Perform DFS to find the shortest distance that visits all places
	res := math.MaxInt32
	for i := 0; i < len(adj); i++ {
		res = ax.Min(res, dfs(adj, 0, 0, 0, len(adj)-1))
	}
	return res
}

func dfs(adj [][]edge, visited, dist, pos, remains int) int {
	if remains == 0 {
		return dist
	}
	res := math.MaxInt32
	for _, near := range adj[pos] {
		if visited&(1<<near.target) > 0 {
			continue
		}
		res = ax.Min(res, dfs(adj, visited|(1<<near.target), dist+near.dist, near.target, remains-1))
	}
	return res
}
