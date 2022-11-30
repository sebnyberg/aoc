package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint]
var tof = ax.MustParseFloat[float64]
var mini = ax.Min[int]
var minf = ax.Min[float64]
var minu = ax.Min[uint]

func solve1(input []parsedRow) string {
	var total int
	for _, r := range input {
		lw := r.l * r.w
		wh := r.w * r.h
		hl := r.h * r.l
		extra := mini(mini(lw, wh), hl)
		total += extra + 2*(lw+wh+hl)
	}
	return sprint(total)
}

func solve2(input []parsedRow) string {
	var total int
	for _, r := range input {
		lw := 2 * (r.l + r.w)
		wh := 2 * (r.w + r.h)
		hl := 2 * (r.h + r.l)
		ribbon := mini(lw, mini(wh, hl))
		vol := r.l * r.w * r.h
		total += ribbon + vol
	}
	return sprint(total)
}

type parsedRow struct {
	l int
	w int
	h int
}

func parse(s string) parsedRow {
	var r parsedRow
	var l, w, h int
	fmt.Sscanf(s, "%dx%dx%d", &l, &w, &h)
	r.l = l
	r.w = w
	r.h = h
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var p ax.Problem[parsedRow]
	p.HeadN = 3
	p.TailN = 3
	for sc.Scan() {
		s := sc.Text()
		p.Input = append(p.Input, s)
		p.Parsed = append(p.Parsed, parse(s))
	}
	p.Result1 = solve1(p.Parsed)
	p.Result2 = solve2(p.Parsed)
	fmt.Fprint(os.Stdout, p)
}
