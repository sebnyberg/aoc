package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

type monkey struct {
	idx           int
	startingItems []int
	items         []int
	op            func(int) int
	mod           int
	next          map[bool]int
}

func solve(inf string, div, nrounds int) any {
	lines := ax.MustReadFileLines(inf)
	var ms []monkey
	for i := 0; i < len(lines); i += 7 {
		var m monkey
		idxs := strings.Fields(lines[i])[1]
		m.idx = ax.Atoi(idxs[:len(idxs)-1])
		itemsa := strings.Fields(strings.ReplaceAll(lines[i+1], ",", ""))[2:]
		m.items = make([]int, len(itemsa))
		for i := range itemsa {
			m.items[i] = ax.Atoi(itemsa[i])
		}
		op := strings.Join(strings.Fields(lines[i+2])[4:], "")
		if op[1] == 'o' {
			m.op = func(a int) int {
				return a * a
			}
		} else if op[0] == '*' {
			m.op = func(a int) int {
				return ax.Atoi(op[1:]) * a
			}
		} else {
			m.op = func(a int) int {
				return ax.Atoi(op[1:]) + a
			}
		}
		m.mod = ax.Atoi(strings.Fields(lines[i+3])[3])
		m.next = make(map[bool]int, 2)
		m.next[true] = ax.Atoi(strings.Fields(lines[i+4])[5])
		m.next[false] = ax.Atoi(strings.Fields(lines[i+5])[5])
		ms = append(ms, m)
	}
	count := make(map[int]int)

	mod := 1
	for _, m := range ms {
		mod *= m.mod
	}

	for round := 0; round < nrounds; round++ {
		for i := range ms {
			for _, x := range ms[i].items {
				count[i]++
				x = (ms[i].op(x) / div) % mod
				var j int
				j = ms[i].next[x%ms[i].mod == 0]
				ms[j].items = append(ms[j].items, x)
			}
			ms[i].items = ms[i].items[:0]
		}
	}
	var counts []int
	for _, c := range count {
		counts = append(counts, c)
	}

	sort.Ints(counts)
	k := len(counts)
	return counts[k-1] * counts[k-2]
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, 3, 20))
	fmt.Printf("Result2:\n%v\n", solve(f, 1, 10000))
}
