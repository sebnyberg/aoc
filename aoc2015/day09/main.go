package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var absf = ax.Abs[float64]
var absi = ax.Abs[int]
var minf = ax.Min[float64]
var mini = ax.Min[int]
var minu = ax.Min[uint16]
var maxf = ax.Max[float64]
var maxi = ax.Max[int]
var maxu = ax.Max[uint16]
var print = fmt.Print
var printf = fmt.Printf
var println = fmt.Println
var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var sprintln = fmt.Sprintln
var tof = ax.MustParseFloat[float64]
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint16]

func pprint(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmt.Printf(fmtStr, a...)
}
func pprintln(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmtStr += "\n"
	fmt.Printf(fmtStr, a...)
}

var intr = regexp.MustCompile(`[1-9][0-9]*|0`)

func isnum(s string) bool {
	return intr.MatchString(s)
}

func Solve1(rs []parsedRow) string {
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

	return sprint(res)
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

func Solve2(rs []parsedRow) string {
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

	return sprint(res)
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

type parsedRow struct {
	s    string
	u, v string
	dist int
}

func Parse(s string) parsedRow {
	var r parsedRow
	r.s = s
	fmt.Sscanf(s, "%s to %s = %d", &r.u, &r.v, &r.dist)
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var p ax.Problem[parsedRow]
	p.HeadN = 3
	p.TailN = 3
	for sc.Scan() {
		s := sc.Text()
		p.Input = append(p.Input, s)
		p.Parsed = append(p.Parsed, Parse(s))
	}
	p.Result1 = Solve1(p.Parsed)
	p.Result2 = Solve2(p.Parsed)
	fmt.Fprint(os.Stdout, p)
}
