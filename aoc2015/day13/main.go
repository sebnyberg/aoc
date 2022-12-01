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
	names := []string{}
	idx := map[string]int{}
	for _, r := range rs {
		if _, exists := idx[r.a]; !exists {
			idx[r.a] = len(names)
			names = append(names, r.a)
		}
		if _, exists := idx[r.b]; !exists {
			idx[r.b] = len(names)
			names = append(names, r.b)
		}
	}
	n := len(names)
	deltas := make([][]int, n)
	for i := range deltas {
		deltas[i] = make([]int, n)
	}
	for _, r := range rs {
		deltas[idx[r.a]][idx[r.b]] = r.delta
	}

	res := math.MinInt32
	var tried int
	cost := func(arr []int) {
		var x int
		for i, j := range arr {
			left := arr[(i-1+n)%n]
			right := arr[(i+1)%n]
			x += deltas[j][left] + deltas[j][right]
		}
		tried++
		if x > res {
			res = x
		}
	}

	var dfs func([]bool, []int, int)
	dfs = func(seen []bool, arr []int, j int) {
		if j == n {
			cost(arr)
			return
		}
		for i := range seen {
			if seen[i] {
				continue
			}
			seen[i] = true
			arr[j] = i
			dfs(seen, arr, j+1)
			seen[i] = false
		}
	}

	seen := make([]bool, n)
	arr := make([]int, n)
	dfs(seen, arr, 0)

	return fmt.Sprint(res)
}

func solve2(in *input) string {
	in.xs = append(in.xs, inputItem{
		a:     "Me",
		b:     "Carol",
		delta: 0,
	})
	return solve1(in)
}

type inputItem struct {
	s     string
	a     string
	b     string
	delta int
}

type input struct {
	n  int
	xs []inputItem
}

var rrr = regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)`)

func (p *input) parse(s string) {
	var x inputItem
	ss := rrr.FindStringSubmatch(s)
	x.a = ss[1]
	x.b = ss[4]
	x.delta = ax.Atoi(ss[3])
	if ss[2] == "lose" {
		x.delta *= -1
	}
	x.s = s
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
