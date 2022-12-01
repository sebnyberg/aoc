package main

import (
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	var total int
	for _, r := range in.xs {
		lw := r.l * r.w
		wh := r.w * r.h
		hl := r.h * r.l
		extra := ax.Min(ax.Min(lw, wh), hl)
		total += extra + 2*(lw+wh+hl)
	}
	return fmt.Sprint(total)
}

func solve2(in *input) string {
	var total int
	for _, r := range in.xs {
		lw := 2 * (r.l + r.w)
		wh := 2 * (r.w + r.h)
		hl := 2 * (r.h + r.l)
		ribbon := ax.Min(lw, ax.Min(wh, hl))
		vol := r.l * r.w * r.h
		total += ribbon + vol
	}
	return fmt.Sprint(total)
}

type inputItem struct {
	w, l, h int
}

type input struct {
	n  int
	xs []inputItem
}

func (p *input) parse(s string) {
	var x inputItem
	fmt.Sscanf(s, "%dx%dx%d", &x.w, &x.l, &x.h)
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
