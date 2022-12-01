package main

import (
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	rs := in.xs
	apply := func(s string) string {
		res := []byte{}
		val := s[0]
		var count int
		n := len(s)
		for i := 0; i <= n; i++ {
			if i == n || (i > 0 && s[i] != s[i-1]) {
				res = append(res, fmt.Sprint(count)...)
				res = append(res, val)
				if i == n {
					break
				}
				count = 0
			}
			val = s[i]
			count++
		}
		return string(res)
	}
	s := rs[0].s
	for i := 0; i < 40; i++ {
		s = apply(s)
	}
	return fmt.Sprint(len(s))
}

func solve2(in *input) string {
	rs := in.xs
	apply := func(s string) string {
		res := []byte{}
		val := s[0]
		var count int
		n := len(s)
		for i := 0; i <= n; i++ {
			if i == n || (i > 0 && s[i] != s[i-1]) {
				res = append(res, fmt.Sprint(count)...)
				res = append(res, val)
				if i == n {
					break
				}
				count = 0
			}
			val = s[i]
			count++
		}
		return string(res)
	}
	s := rs[0].s
	for i := 0; i < 50; i++ {
		s = apply(s)
	}
	return fmt.Sprint(len(s))
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
