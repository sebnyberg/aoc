package main

import (
	"fmt"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var res int
	for k := 0; k*3 < len(lines); k++ {
		left := parseList(lines[k*3])
		right := parseList(lines[k*3+1])
		if left.less(right) {
			res += (k + 1)
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

func (left *listNode) less(right *listNode) bool {
	if left.isval || right.isval {
		if left.isval && right.isval {
			return left.val <= right.val
		}
		if right.isval {
			a := &listNode{list: []*listNode{right}}
			return left.less(a)
		}
		a := &listNode{list: []*listNode{left}}
		return a.less(right)
	}

	n := ax.Min(len(left.list), len(right.list))
	for i := 0; i < n; i++ {
		if !left.list[i].less(right.list[i]) {
			return false
		}
	}
	return len(left.list) <= len(right.list)
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
}
