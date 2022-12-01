package main

import (
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
	"golang.org/x/exp/slices"
)

func solve1(in *input) string {
	target := in.xs[0].x
	var count []int
	for n := 1000; ; n *= 10 {
		for i := range count {
			count[i] = 0
		}
		count = slices.Grow(count, n+1)
		count = count[:cap(count)]
		for x := 2; x <= n; x++ {
			for y := x; y <= n; y += x {
				count[y] += x * 10
			}
		}
		for x := 1; x <= n; x++ {
			if count[x] >= target {
				return fmt.Sprint(x)
			}
		}
	}
}

func solve2(in *input) string {
	target := in.xs[0].x
	var count []int
	for n := 1000; ; n *= 10 {
		for i := range count {
			count[i] = 0
		}
		count = slices.Grow(count, n+1)
		count = count[:cap(count)]
		for x := 1; x <= n; x++ {
			y := x
			for k := 0; k < 50; k++ {
				if y > n {
					break
				}
				count[y] += x * 11
				y += x
			}
		}
		for x := 1; x <= n; x++ {
			if count[x] >= target {
				return fmt.Sprint(x)
			}
		}
	}
}

type inputItem struct {
	x int
}

type input struct {
	n  int
	xs []inputItem
}

func (p *input) parse(s string) {
	var x inputItem
	x.x = ax.Atoi(s)
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	f, _ := os.Open("input")
	rows := ax.ReadLines(f)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
