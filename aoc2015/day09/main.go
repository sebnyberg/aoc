package main

import (
	"fmt"
	"math"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	rs := in.xs
	adj := make(map[string][]string)
	dist := make(map[string][]int)

	for _, r := range rs {
		a, b := r.u, r.v
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		dist[a] = append(dist[a], r.dist)
		dist[b] = append(dist[b], r.dist)
	}

	res := math.MaxInt32
	for a := range adj {
		seen := make(map[string]bool)
		seen[a] = true
		a := dfs(seen, adj, dist, a, 1)
		if a < res {
			res = a
		}
	}

	return fmt.Sprint(res)
}

func dfs(seen map[string]bool, adj map[string][]string, dist map[string][]int, a string, n int) int {
	if n == len(adj) {
		return 0
	}
	res := math.MaxInt32
	for i, b := range adj[a] {
		if seen[b] {
			continue
		}
		seen[b] = true
		x := dist[a][i]
		y := dfs(seen, adj, dist, b, n+1)
		if x+y < res {
			res = x + y
		}
		seen[b] = false
	}
	return res
}

func solve2(in *input) string {
	rs := in.xs
	adj := make(map[string][]string)
	dist := make(map[string][]int)

	for _, r := range rs {
		a, b := r.u, r.v
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		dist[a] = append(dist[a], r.dist)
		dist[b] = append(dist[b], r.dist)
	}

	var res int
	for a := range adj {
		seen := make(map[string]bool)
		seen[a] = true
		a := dfs2(seen, adj, dist, a, 1)
		if a > res {
			res = a
		}
	}

	return fmt.Sprint(res)
}

func dfs2(seen map[string]bool, adj map[string][]string, dist map[string][]int, a string, n int) int {
	if n == len(adj) {
		return 0
	}
	var res int
	for i, b := range adj[a] {
		if seen[b] {
			continue
		}
		seen[b] = true
		x := dist[a][i]
		y := dfs2(seen, adj, dist, b, n+1)
		if x+y > res {
			res = x + y
		}
		seen[b] = false
	}
	return res
}

type inputItem struct {
	s    string
	u, v string
	dist int
}

type input struct {
	n  int
	xs []inputItem
}

var pat = regexp.MustCompile(``)

func (p *input) parse(s string) {
	var x inputItem
	x.s = s
	fmt.Sscanf(s, "%s to %s = %d", &x.u, &x.v, &x.dist)
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	rows := ax.ReadLines(os.Stdin)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
