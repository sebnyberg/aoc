package main

import (
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	m := ax.Set[[2]int]{}
	var x, y int
	m.Add([2]int{x, y})
	for _, s := range in.xs[0].s {
		switch s {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		m.Add([2]int{x, y})
	}
	return fmt.Sprint(len(m))
}

func solve2(in *input) string {
	m := ax.Set[[2]int]{}
	var xy [2][2]int
	x := 0
	y := 1
	m.Add([2]int{0, 0})
	for i := 0; i < len(in.xs[0].s); i += 2 {
		for j := 0; j < 2; j++ {
			switch in.xs[0].s[i+j] {
			case '^':
				xy[j][y]++
			case 'v':
				xy[j][y]--
			case '>':
				xy[j][x]++
			case '<':
				xy[j][x]--
			}
			m.Add([2]int{xy[j][x], xy[j][y]})
		}
	}
	return fmt.Sprint(len(m))
}

type inputItem struct {
	s string
}

type input struct {
	n  int
	xs []inputItem
}

func (p *input) parse(s string) {
	var x inputItem
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
