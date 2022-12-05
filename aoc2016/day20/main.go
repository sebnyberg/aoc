package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	lines := ax.MustReadFileLines(inf)
	// lines = []string{
	// 	"5-8",
	// 	"0-2",
	// 	"4-7",
	// }
	type delta struct {
		d     int
		index int
	}
	var deltas []delta
	for _, l := range lines {
		fs := strings.Split(l, "-")
		a := ax.Atoi(fs[0])
		b := ax.Atoi(fs[1])
		deltas = append(deltas, delta{1, a}, delta{-1, b + 1})
	}
	sort.Slice(deltas, func(i, j int) bool {
		if deltas[i].index == deltas[j].index {
			return deltas[i].d > deltas[j].d
		}
		return deltas[i].index < deltas[j].index
	})
	var nranges int
	var t0 int
	var res int
	var first string
	for _, d := range deltas {
		if d.d == -1 && nranges == 1 {
			t0 = d.index
		} else if d.d == 1 && nranges == 0 {
			res += d.index - t0
		}
		nranges += d.d
		if nranges == 0 && first == "" {
			first = fmt.Sprint(d.index)
		}
	}
	return fmt.Sprintf("first:\n%s\nsecond:\n%s\n", first, fmt.Sprint(res))
}

func solve2(inf string) string {
	var res int
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
