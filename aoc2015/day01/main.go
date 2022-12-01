package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	var floor int
	for _, ch := range in.xs[0].s {
		switch ch {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return fmt.Sprint(floor)
}

func solve2(in *input) string {
	var floor int
	for i, ch := range in.xs[0].s {
		switch ch {
		case '(':
			floor++
		case ')':
			if floor == 0 {
				return fmt.Sprint(i + 1)
			}
			floor--
		}
	}
	return ""
}

type inputItem struct {
	s string
}

type input struct {
	n  int
	xs []inputItem
}

var pat = regexp.MustCompile(``)

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
