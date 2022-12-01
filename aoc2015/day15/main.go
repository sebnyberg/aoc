package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	rs := in.xs
	var res int
	var dfs func([5]int, int, int)
	dfs = func(stats [5]int, rem, i int) {
		if i == len(rs) {
			if rem != 0 {
				return
			}
			a := ax.Max(0, stats[0]) * ax.Max(0, stats[1]) *
				ax.Max(0, stats[2]) * ax.Max(0, stats[3])
			res = ax.Max(res, a)
			return
		}
		for x := 0; x <= rem; x++ {
			c := x * rs[i].capacity
			d := x * rs[i].durability
			f := x * rs[i].flavour
			t := x * rs[i].texture
			stats[0] += c
			stats[1] += d
			stats[2] += f
			stats[3] += t
			dfs(stats, rem-x, i+1)
			stats[0] -= c
			stats[1] -= d
			stats[2] -= f
			stats[3] -= t
		}
	}
	dfs([5]int{}, 100, 0)
	return fmt.Sprint(res)
}

func solve2(in *input) string {
	var res int
	rs := in.xs
	var dfs func([5]int, int, int)
	dfs = func(stats [5]int, rem, i int) {
		if i == len(rs) {
			if rem != 0 || stats[4] != 500 {
				return
			}
			a := ax.Max(0, stats[0]) * ax.Max(0, stats[1]) *
				ax.Max(0, stats[2]) * ax.Max(0, stats[3])
			res = ax.Max(res, a)
			return
		}
		for x := 0; x <= rem; x++ {
			stats[0] += x * rs[i].capacity
			stats[1] += x * rs[i].durability
			stats[2] += x * rs[i].flavour
			stats[3] += x * rs[i].texture
			stats[4] += x * rs[i].calories
			dfs(stats, rem-x, i+1)
			stats[0] -= x * rs[i].capacity
			stats[1] -= x * rs[i].durability
			stats[2] -= x * rs[i].flavour
			stats[3] -= x * rs[i].texture
			stats[4] -= x * rs[i].calories
		}
	}
	dfs([5]int{}, 100, 0)
	return fmt.Sprint(res)
}

type inputItem struct {
	s          string
	name       string
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

type input struct {
	n  int
	xs []inputItem
}

var rrr = regexp.MustCompile(`(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

func (p *input) parse(s string) {
	var x inputItem
	ss := rrr.FindStringSubmatch(s)
	x.name = ss[1]
	x.capacity = ax.Atoi(ss[2])
	x.durability = ax.Atoi(ss[3])
	x.flavour = ax.Atoi(ss[4])
	x.texture = ax.Atoi(ss[5])
	x.calories = ax.Atoi(ss[6])
	p.xs = append(p.xs, x)
	x.s = s
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
