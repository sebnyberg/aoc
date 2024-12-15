package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func print(state [][]byte) {
	for i := range state {
		fmt.Println(string(state[i]))
	}
	fmt.Println("")
}

func expand(state [][]byte) [][]byte {
	m := len(state)
	n := len(state[0])
	next := make([][]byte, m)
	for i := range next {
		next[i] = make([]byte, n*2)
	}
	for i := range state {
		for j, v := range state[i] {
			switch v {
			case '#', '.':
				next[i][j*2] = v
				next[i][j*2+1] = v
			case '@':
				next[i][j*2] = v
				next[i][j*2+1] = '.'
			case 'O':
				next[i][j*2] = '['
				next[i][j*2+1] = ']'
			default:
				log.Fatalln("error!")
			}
		}
	}
	return next
}

const (
	nodeTypeBox     = 'O'
	nodeTypeWideBox = '['
	nodeTypeWall    = '#'
	nodeTypeEmpty   = '.'
	nodeTypeRobot   = '@'
)

type node struct {
	i, j     int
	typ      byte
	children []*node
	pushed   bool
}

func (n *node) draw(state [][]byte) {
	if n == nil || n.typ == nodeTypeEmpty {
		return
	}
	state[n.i][n.j] = n.typ
	if n.typ == nodeTypeWideBox {
		state[n.i][n.j+1] = ']'
	}
	for _, x := range n.children {
		x.draw(state)
	}
}

func (n *node) clear(state [][]byte) {
	if n == nil || n.typ == nodeTypeEmpty {
		return
	}
	state[n.i][n.j] = '.'
	if n.typ == nodeTypeWideBox {
		state[n.i][n.j+1] = '.'
	}
	for _, x := range n.children {
		x.clear(state)
	}
}

func (n *node) push(di, dj int) {
	if n == nil || n.pushed {
		return
	}
	n.pushed = true
	n.i += di
	n.j += dj
	for _, x := range n.children {
		x.push(di, dj)
	}
}

var adjustJ = []int{
	'@': 0,
	']': -1,
	'[': 0,
	'.': 0,
	'#': 0,
	'O': 0,
}

func parseGraph(state [][]byte, i, j, di, dj int) *node {
	root := &node{i, j, nodeTypeRobot, nil, false}
	nodes := map[[2]int]*node{{i, j}: root}
	curr := []*node{root}
	next := []*node{}
	getNode := func(next *[]*node, i, j int, typ byte) *node {
		if _, ok := nodes[[2]int{i, j}]; !ok {
			x := &node{i, j, typ, nil, false}
			nodes[[2]int{i, j}] = x
			*next = append(*next, x)
		}
		return nodes[[2]int{i, j}]
	}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			if x.typ == nodeTypeWall {
				return nil // cannot move
			}
			if x.typ == nodeTypeEmpty {
				continue // no more children
			}

			ii := x.i + di
			jj := x.j + dj + adjustJ[state[x.i+di][x.j+dj]]

			// Non-wide-boxes are easy.
			if x.typ != nodeTypeWideBox {
				y := getNode(&next, ii, jj, state[ii][jj])
				x.children = append(x.children, y)
				continue
			}

			// For wide boxes and left/right, adjust jj accordingly and add to
			// children.
			if di == 0 {
				if dj == 1 {
					jj += 2
				}
				y := getNode(&next, ii, jj, state[ii][jj])
				x.children = append(x.children, y)
				continue
			}

			// For wide boxes and up/down, start by adding the left above
			left := getNode(&next, ii, jj, state[ii][jj])
			x.children = append(x.children, left)
			jjj := x.j + 1 + adjustJ[state[ii][x.j+1]]
			if jjj != jj {
				right := getNode(&next, ii, jjj, state[ii][jjj])
				x.children = append(x.children, right)
			}
		}
		curr, next = next, curr
	}
	return root
}

func solve(inf string, doExpand bool) any {
	lines := ax.MustReadFileLines(inf)
	var state [][]byte
	for lines[0] != "" {
		state = append(state, []byte(lines[0]))
		lines = lines[1:]
	}
	if doExpand {
		state = expand(state)
	}
	var pos [2]int
	for i := range state {
		for j, v := range state[i] {
			if v == '@' {
				pos = [2]int{i, j}
			}
		}
	}
	lines = lines[1:]
	moves := strings.Join(lines, "")
	for _, m := range moves {
		var di, dj int
		switch m {
		case '<':
			di, dj = 0, -1
		case '^':
			di, dj = -1, 0
		case 'v':
			di, dj = 1, 0
		case '>':
			di, dj = 0, 1
		}
		// Build a root of nodes from the starting position
		root := parseGraph(state, pos[0], pos[1], di, dj)
		if root == nil {
			continue
		}

		// Push forward
		root.clear(state)
		root.push(di, dj)
		root.draw(state)
		pos = [2]int{root.i, root.j}
	}

	var res int
	for i := range state {
		for j, v := range state[i] {
			if v == '[' || v == 'O' {
				res += 100*i + j
			}
		}
	}
	return res
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve("input", false))
	fmt.Printf("Result2:\n%v\n", solve("input", true))
}
