package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(
	`Valve (\w+) has flow rate=(\d+); (tunnel|tunnels) (lead|leads) to (valve|valves) (.*)`)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var names []string
	nameIdx := make(map[string]int)
	flow := []int{}
	adj := [][]int{}
	for _, l := range lines {
		g := pat.FindStringSubmatch(l)
		valve := g[1]
		leadsTo := strings.Split(g[6], ",")
		if _, exists := nameIdx[valve]; !exists {
			nameIdx[valve] = len(names)
			names = append(names, valve)
			flow = append(flow, 0)
			adj = append(adj, []int{})
		}
		for _, v := range leadsTo {
			v = strings.Trim(v, " ")
			if _, exists := nameIdx[v]; !exists {
				nameIdx[v] = len(names)
				names = append(names, v)
				flow = append(flow, 0)
				adj = append(adj, []int{})
			}
		}
		flow[nameIdx[valve]] = ax.Atoi(g[2])
		for _, v := range leadsTo {
			v = strings.Trim(v, " ")
			i := nameIdx[valve]
			j := nameIdx[v]
			adj[i] = append(adj[i], j)
		}
	}
	ii := nameIdx["AA"]
	mem := make(map[state1]int)
	res := dfs1(mem, adj, flow, ii, 0, 29)
	return res
}

type state1 struct {
	pos  int
	open int
	t    int
}

func dfs1(mem map[state1]int, adj [][]int, flow []int, i, bm, t int) int {
	if t == 0 {
		return 0
	}
	s := state1{
		pos:  i,
		open: bm,
		t:    t,
	}
	if v, exists := mem[s]; exists {
		return v
	}
	n := len(flow)
	if bm == (1<<n)-1 {
		return 0
	}
	// If on a valve and it has pressure, open it
	if bm&(1<<i) == 0 && flow[i] > 0 {
		res := flow[i]*t + dfs1(mem, adj, flow, i, bm|(1<<i), t-1)
		mem[s] = res
		return res
	}

	// Otherwise, move to a new valve
	var res int
	for _, nei := range adj[i] {
		var x int
		res = ax.Max(res, x+dfs1(mem, adj, flow, nei, bm, t-1))
	}
	mem[s] = res
	return mem[s]
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var names []string
	nameIdx := make(map[string]int)
	flow := []int{}
	adj := [][]int{}
	for _, l := range lines {
		g := pat.FindStringSubmatch(l)
		valve := g[1]
		leadsTo := strings.Split(g[6], ",")
		if _, exists := nameIdx[valve]; !exists {
			nameIdx[valve] = len(names)
			names = append(names, valve)
			flow = append(flow, 0)
			adj = append(adj, []int{})
		}
		for _, v := range leadsTo {
			v = strings.Trim(v, " ")
			if _, exists := nameIdx[v]; !exists {
				nameIdx[v] = len(names)
				names = append(names, v)
				flow = append(flow, 0)
				adj = append(adj, []int{})
			}
		}
		flow[nameIdx[valve]] = ax.Atoi(g[2])
		for _, v := range leadsTo {
			v = strings.Trim(v, " ")
			i := nameIdx[valve]
			j := nameIdx[v]
			adj[i] = append(adj[i], j)
		}
	}
	// We need to calculate the length between valves that have a positive flow
	//
	// Start by compressing the graph to only hold nodes with positive pressure
	n := len(names)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	var ii int
	for i, v := range names {
		if v == "AA" {
			ii = i
			break
		}
	}
	sort.Slice(idx, func(i, j int) bool {
		if idx[i] == ii {
			return true
		}
		if idx[j] == ii {
			return false
		}
		return flow[idx[i]] > flow[idx[j]]
	})
	inv := make([]int, n)
	for i, j := range idx {
		inv[j] = i
	}
	adj2 := make([][]int, n)
	flow2 := make([]int, n)
	names2 := make([]string, n)
	for i := range idx {
		adj2[i] = adj[idx[i]]
		for j := range adj2[i] {
			adj2[i][j] = inv[adj2[i][j]]
		}
		flow2[i] = flow[idx[i]]
		names2[i] = names[idx[i]]
	}
	adj = adj2
	flow = flow2
	names = names2
	var m int
	for i := 1; i < n; i++ {
		if flow[i] == 0 {
			m = i
			break
		}
	}

	calcDist := func(k int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[k] = 0
		curr := []int{k}
		next := []int{}
		for steps := 1; len(curr) > 0; steps++ {
			next = next[:0]
			for _, x := range curr {
				for _, nei := range adj[x] {
					if dist[nei] != math.MaxInt32 {
						continue
					}
					dist[nei] = steps
					next = append(next, nei)
				}
			}
			curr, next = next, curr
		}
		return dist
	}
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = calcDist(i)[:m]
	}
	flow = flow[:m]
	names = names[:m]
	adj = adj[:m]

	mem := make(map[state2]int)
	res := dfs2(mem, dist, flow, [2]int{0, 0}, [2]int{0, 0}, 0)
	return res
}

type state2 struct {
	pos  [2]int
	open int
	t    [2]int
}

func dfs2(mem map[state2]int, dist [][]int, flow []int, pos, t [2]int, bm int) int {
	if t[0] > 26 || t[1] > 26 {
		return math.MinInt32
	}
	s := state2{
		pos:  pos,
		open: bm,
		t:    t,
	}
	if v, exists := mem[s]; exists {
		return v
	}
	n := len(flow)
	var p int
	if t[1] < t[0] {
		p = 1 // elephant's turn to move
	}
	var res int

	// Try to move to all unseen valves
	k := pos[p]
	for j := 1; j < n; j++ {
		if bm&(1<<j) > 0 {
			continue
		}
		next := pos
		next[p] = j
		nextT := t
		nextT[p] += dist[k][j] + 1
		pressure := flow[j] * (26 - nextT[p])
		x := dfs2(mem, dist, flow, next, nextT, bm|(1<<j))
		res = ax.Max(res, pressure+x)
	}

	mem[s] = res
	return mem[s]
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
