package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	var res int
	xs := in.xs

	_ = xs
	return fmt.Sprint(res)
}

func solve2(in *input) string {
	var res int
	xs := in.xs

	_ = xs
	return fmt.Sprint(res)
}

type inputItem struct {
	s       string
	t       string
	a, b, c int
	x, y    int
	x1, y1  int
	x2, y2  int
}

type input struct {
	n  int
	xs []inputItem
}

var pat = regexp.MustCompile(``)

func (p *input) parse(s string) {
	var x inputItem

	// fmt.Sscanf(s, "%dx%dx%d", &item.a, &item.b, &item.c)

	ss := pat.FindStringSubmatch(s)
	// p.a = ax.Atoi[ss[1]]
	// p.a = ax.Atoi[ss[1]]
	// p.a = ax.Atoi[ss[1]]
	_ = ss

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
