package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var res int
	for k := 0; k*3 < len(lines); k++ {
		left := parseList(lines[k*3])
		right := parseList(lines[k*3+1])
		d := left.cmp(right)
		if d <= 0 {
			res += (k + 1)
		}
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var packets []*listNode
	for k := 0; k*3 < len(lines); k++ {
		packets = append(packets, parseList(lines[k*3]))
		packets = append(packets, parseList(lines[k*3+1]))
	}
	packets = append(packets, parseList("[[2]]"))
	packets = append(packets, parseList("[[6]]"))
	n := len(packets)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return packets[idx[i]].cmp(packets[idx[j]]) <= 0
	})
	res := 1
	for _, i := range idx {
		if idx[i] >= n-2 {
			res *= (i + 1)
		}
	}
	return res
}

type listNode struct {
	val   int
	isval bool
	list  []*listNode
}

func (n *listNode) islist() bool {
	return !n.isval
}

func (l *listNode) cmp(r *listNode) int {
	if l.isval || r.isval {
		if l.isval && r.isval {
			return l.val - r.val
		}
		if r.isval {
			r = &listNode{list: []*listNode{r}}
			return l.cmp(r)
		}
		l = &listNode{list: []*listNode{l}}
		return l.cmp(r)
	}

	n := ax.Min(len(l.list), len(r.list))
	for i := 0; i < n; i++ {
		d := l.list[i].cmp(r.list[i])
		if d != 0 {
			return d
		}
	}
	return len(l.list) - len(r.list)
}

func parseList(s string) *listNode {
	stack := []*listNode{{}}
	var npar int
	for i := 0; i < len(s); {
		if s[i] == ',' {
			i++
			continue
		}
		if s[i] == '[' {
			npar++
			stack = append(stack, &listNode{})
			i++
			continue
		}
		if s[i] == ']' {
			stack[npar-1].list = append(stack[npar-1].list, stack[npar])
			stack = stack[:npar]
			npar--
			i++
			continue
		}
		var x int
		for i < len(s) && s[i] <= '9' && s[i] >= '0' {
			x *= 10
			x += int(s[i] - '0')
			i++
		}
		stack[npar].list = append(stack[npar].list,
			&listNode{
				val:   x,
				isval: true,
			},
		)
	}
	return stack[0].list[0]
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
