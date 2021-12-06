package p_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 909, res)
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
		dist := ax.MustParseIntBase(parts[3], 10)
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
	var f maxDistFinder
	for i := 0; i < len(adj); i++ {
		f.dfs(adj, 1<<i, 0, i, len(adj)-1)
	}
	return f.maxDist
}

type maxDistFinder struct {
	maxDist int
}

func (f *maxDistFinder) dfs(adj [][]edge, visited, dist, pos, remains int) {
	if remains == 0 {
		f.maxDist = ax.Max(f.maxDist, dist)
	}
	for _, near := range adj[pos] {
		if visited&(1<<near.target) > 0 {
			continue
		}
		f.dfs(adj, visited|(1<<near.target), dist+near.dist, near.target, remains-1)
	}
}
