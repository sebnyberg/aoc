package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) string {
	var res int
	lines := ax.MustReadFileLines(inf)
	parseDelta := func(s string) []int {
		ab := strings.Split(s, "-")
		return []int{ax.Atoi(ab[0]), ax.Atoi(ab[1])}
	}
	for _, l := range lines {
		ab := strings.Split(l, ",")
		first := parseDelta(ab[0])
		second := parseDelta(ab[1])
		if second[0] >= first[0] && second[1] <= first[1] ||
			(first[0] >= second[0] && first[1] <= second[1]) {
			res++
		}
	}
	return fmt.Sprint(res)
}

func solve2(inf string) string {
	lines := ax.MustReadFileLines(inf)
	parseDelta := func(s string) []int {
		ab := strings.Split(s, "-")
		return []int{ax.Atoi(ab[0]), ax.Atoi(ab[1])}
	}
	type delta struct {
		i      int
		change int
	}
	var res int
	for _, l := range lines {
		ab := strings.Split(l, ",")
		first := parseDelta(ab[0])
		second := parseDelta(ab[1])
		var deltas []delta
		deltas = append(deltas, delta{first[0], 1}, delta{first[1] + 1, -1})
		deltas = append(deltas, delta{second[0], 1}, delta{second[1] + 1, -1})
		sort.Slice(deltas, func(i, j int) bool {
			if deltas[i].i == deltas[j].i {
				return deltas[i].change < deltas[j].change
			}
			return deltas[i].i < deltas[j].i
		})
		var ninterval int
		for _, d := range deltas {
			ninterval += d.change
			if ninterval >= 2 {
				res++
				break
			}
		}
	}
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
