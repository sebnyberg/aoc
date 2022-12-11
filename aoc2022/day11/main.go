package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

type monkey struct {
	idx           int
	startingItems []int
	items         []int
	op            func(int) int
	mod           int
	trueTarget    int
	falseTarget   int
}

func solve1(inf string) any {
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
		m.trueTarget = ax.Atoi(strings.Fields(lines[i+4])[5])
		m.falseTarget = ax.Atoi(strings.Fields(lines[i+5])[5])
		ms = append(ms, m)
	}
	count := make(map[int]int)

	for round := 0; round < 20; round++ {
		for i, m := range ms {
			for _, x := range m.items {
				count[i]++
				x = m.op(x) / 3
				var j int
				if x%m.mod == 0 {
					j = m.trueTarget
				} else {
					j = m.falseTarget
				}
				// fmt.Printf("mokey %v throws item %v to %v\n", i, x, j)
				ms[j].items = append(ms[j].items, x)
			}
			ms[i].items = ms[i].items[:0]
		}
	}
	fmt.Println(count)
	return ""
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
}
